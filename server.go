package main

import (
	"context"
	"github.com/ShootingCoin/api-server/entity"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/ShootingCoin/api-server/core/handler"
	"github.com/ShootingCoin/api-server/utils"
)

func main() {
	e := echo.New()

	txConfig, err := entity.SetTxConfig()
	if err != nil {
		log.Fatalf("Failed to set transaction config: %v", err)
		return
	}

	infuraURL, err := utils.GetInfuraUrl()
	if err != nil {
		return
	}

	ethClient, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Initialize WebSocket upgrader
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections by default
			return true
		},
	}

	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check the connection to Redis
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Errorln("Failed to connect to Redis: %v", err)
	}

	webSocketHandler := handler.NewWebSocketHandler(upgrader)
	matchHandler := handler.NewMatchHandler(rdb, ethClient, txConfig)
	resultHandler := handler.NewResultHandler()

	// Handler for WebSocket connections
	e.GET("/v1/ws", webSocketHandler.ConnectWebSocket)

	// Handler for matching game requests
	e.POST("/v1/match", matchHandler.MatchGame)

	// Handler for getting game results
	e.POST("/v1/result/:gameUuid", resultHandler.SaveGameResult)

	// Add CORS middleware
	e.Use(middleware.CORS(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}),
		middleware.AddTrailingSlash(),
		middleware.Recover(),
	)

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
