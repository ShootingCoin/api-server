package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/ShootingCoin/api-server/core/common"
	"github.com/ShootingCoin/api-server/utils"
)

const (
	requestsKeyPrefix = "requests:"
	timeoutDuration   = 15 * time.Second
)

type MatchHandler struct {
	rdb *redis.Client
}

type MatchGameReq struct {
	Price int64  `json:"price" validate:"required"`
	Uuid  string `json:"uuid" validate:"required"`
}

func (r *MatchGameReq) Validate() error {
	return validator.New().Struct(r)
}

func (h *MatchHandler) MatchGame(c echo.Context) error {
	// Get the request body
	var matchReq MatchGameReq
	if err := c.Bind(&matchReq); err != nil {
		log.Errorln(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request: %v", err))
	}

	// Validate the request
	if err := matchReq.Validate(); err != nil {
		log.Errorln(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request: %v", err))
	}

	// Convert price to bucket value
	reqBucket, err := utils.GetBucket(matchReq.Price)
	if err != nil {
		log.Errorln(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid price: %v", err))
	}
	reqUuid := matchReq.Uuid

	// Start background worker to check for matching requests
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
		defer cancel()

		// Generate UUID for matched game
		gameUuid, err := uuid.NewRandom()
		if err != nil {
			log.Errorln(err)
			return
		}

		var matchUuid string
		for {
			select {
			// Time out after 15 seconds
			case <-ctx.Done():
				// Send message to requesting player that no matches were made
				if conn, ok := common.ListConnections()[reqUuid]; ok {
					err := conn.WriteMessage(websocket.TextMessage, []byte("No matches found"))
					if err != nil {
						log.Errorln(err)
					}
				}
				return
			default:
				// No matching requests, add original request to Redis
				if h.rdb.LLen(ctx, requestsKeyPrefix+reqBucket).Val() == 0 {
					// Add original request back to Redis since there are no more matching requests
					h.pushRequest(ctx, requestsKeyPrefix+reqBucket, reqUuid+":"+reqBucket)

					return
				}

				// Pop matching request from Redis
				if _, err := h.rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
					popResult := pipe.RPop(ctx, requestsKeyPrefix+reqBucket)

					if popResult.Err() != nil {
						log.Errorln(popResult.Err())
						return popResult.Err()
					}

					req, err := popResult.Result()
					if err != nil {
						log.Errorln(err)
						return err
					}

					// Extract UUID and request from Redis value
					parts := strings.Split(req, ":")
					if len(parts) != 2 {
						log.Errorf("Invalid Redis value: %v", req)
						return fmt.Errorf("invalid Redis value")
					}
					matchUuid = parts[0]

					return nil
				}); err != nil {
					log.Errorln(err)
					continue
				}

				// Send matching result to requesting client
				if err := common.WriteMessage(reqUuid, gameUuid.String()); err != nil {
					log.Errorln(err)
					h.pushRequest(ctx, requestsKeyPrefix+reqBucket, reqUuid+":"+reqBucket)

					return
				}

				// Send matching result to matched client
				if err := common.WriteMessage(matchUuid, gameUuid.String()); err != nil {
					log.Errorln(err)

					return
				}

				// Mark connections as matched
				if err := common.SetMatched(reqUuid, matchUuid); err != nil {
					log.Errorln(err)
					return
				}
				log.Infof("Match completed: %s:%s", reqUuid, matchUuid)

				return
			}
		}
	}()

	// Return success response
	return c.NoContent(http.StatusOK)
}

func (h *MatchHandler) pushRequest(ctx context.Context, key string, value string) error {
	pushResult := h.rdb.LPush(ctx, key, value)
	if pushResult.Err() != nil {
		log.Errorln(pushResult.Err())
		return pushResult.Err()
	}

	log.Infoln(fmt.Sprintf("Pushed request to Redis - key: %s value: %s", key, value))

	return nil
}

func NewMatchHandler(rdb *redis.Client) *MatchHandler {
	return &MatchHandler{
		rdb: rdb,
	}
}
