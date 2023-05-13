package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
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

func (h *MatchHandler) MatchGames(c echo.Context) error {
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
	bucket, err := utils.GetBucket(matchReq.Price)
	if err != nil {
		log.Errorln(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid price: %v", err))
	}
	reqUuid := matchReq.Uuid

	// Start background worker to check for matching requests
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
		defer cancel()

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
				if h.rdb.LLen(ctx, requestsKeyPrefix+bucket).Val() == 0 {
					// Add original request back to Redis since there are no more matching requests
					pushResult := h.rdb.LPush(ctx, requestsKeyPrefix+bucket, reqUuid+":"+bucket)
					log.Infoln(fmt.Sprintf("No matching requests, adding original request back to Redis: %s", reqUuid+":"+bucket))
					if pushResult.Err() != nil {
						log.Errorln(pushResult.Err())
					}
					return
				}

				// Pop matching request from Redis
				_, err := h.rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
					popResult := pipe.RPop(ctx, requestsKeyPrefix+bucket)
					connections := common.ListConnections()

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
					matchUuid := parts[0]
					bucket2 := parts[1]

					// Check if matched client is still connected
					if !common.IsConnected(matchUuid) {
						log.Errorln(fmt.Sprintf("Matched client is no longer connected: %s", matchUuid))
						return fmt.Errorf("matched client is no longer connected")
					}

					// Send matching result to matched client
					if err := connections[matchUuid].WriteMessage(websocket.TextMessage, []byte(bucket2)); err != nil {
						log.Errorln(err)
						return err
					}
					log.Infof("Matched request: %s", bucket2)

					// Check if requesting client is still connected
					if !common.IsConnected(reqUuid) {
						log.Errorln(fmt.Sprintf("Requesting client is no longer connected: %s", matchUuid))
						return fmt.Errorf("requesting client is no longer connected")
					}

					// Send matching result to requesting client
					if err := connections[reqUuid].WriteMessage(websocket.TextMessage, []byte(bucket)); err != nil {
						log.Errorln(err)
						return err
					}

					return nil
				})

				if err != nil {
					log.Errorln(err)
					continue
				}

			}
		}
	}()

	// Return success response
	return c.NoContent(http.StatusOK)
}

func NewMatchHandler(rdb *redis.Client) *MatchHandler {
	return &MatchHandler{
		rdb: rdb,
	}
}
