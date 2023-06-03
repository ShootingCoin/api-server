package handler

import (
	"context"
	"fmt"
	"github.com/ShootingCoin/api-server/core/chain"
	"github.com/ethereum/go-ethereum/ethclient"
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
	"github.com/ShootingCoin/api-server/entity"
	"github.com/ShootingCoin/api-server/utils"
)

const (
	requestsKeyPrefix = "requests:"
	timeoutDuration   = 15 * time.Second
)

type MatchHandler struct {
	rdb       *redis.Client
	ethClient *ethclient.Client
	txConfig  entity.TxConfig
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

	// Check 'Entered' event
	entered, err := chain.CheckEventEmissions(h.ethClient, h.txConfig)
	if err != nil {
		log.Errorln("Failed to check event emissions: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to check event emissions: %v", err))
	}
	if !entered {
		return c.JSON(http.StatusBadRequest, "Unable to match game: 'Entered' event has not been emitted")
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
				popResult := h.rdb.RPop(ctx, requestsKeyPrefix+reqBucket)

				if popResult.Err() != nil {
					log.Errorln(popResult.Err())
					continue
				}

				req, err := popResult.Result()
				if err != nil {
					log.Errorln(err)
					continue
				}

				// Extract UUID and request from Redis value
				parts := strings.Split(req, ":")
				if len(parts) != 2 {
					log.Errorf("Invalid Redis value: %v", req)
					continue
				}
				matchUuid = parts[0]

				// parse message type
				msgType, err := entity.ParseMessageType("gameId")
				if err != nil {
					log.Errorln(err)
					continue
				}

				// Mark connections as matched
				if err := common.SetMatched(reqUuid, matchUuid); err != nil {
					log.Errorln(err)
					return
				}
				log.Infof("Match completed: %s:%s", reqUuid, matchUuid)

				// Send matching result to requesting client
				if err := common.WriteMessage(reqUuid, entity.Message{
					Type:    msgType.String(),
					Content: gameUuid.String(),
				}); err != nil {
					log.Errorln(err)
					h.pushRequest(ctx, requestsKeyPrefix+reqBucket, reqUuid+":"+reqBucket)

					return
				}

				// Send matching result to matched client
				if err := common.WriteMessage(matchUuid, entity.Message{
					Type:    msgType.String(),
					Content: gameUuid.String(),
				}); err != nil {
					log.Errorln(err)

					return
				}

				//err = chain.StartGame(h.ethClient, h.txConfig, gameUuid, reqUuid, matchUuid)
				//if err != nil {
				//	log.Fatalf("Failed to start the game: %v", err)
				//}

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

func NewMatchHandler(rdb *redis.Client, ethClient *ethclient.Client, txConfig entity.TxConfig) *MatchHandler {
	return &MatchHandler{
		rdb:       rdb,
		ethClient: ethClient,
		txConfig:  txConfig,
	}
}
