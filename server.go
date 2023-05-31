package main

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/ShootingCoin/api-server/core/chain"
	"github.com/ShootingCoin/api-server/core/handler"
	"github.com/ShootingCoin/api-server/utils"
)

func main() {
	e := echo.New()

	// Generate Polygon client
	client, err := ethclient.Dial("https://rpc-mumbai.matic.today")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Retrieve the private key and address from the keystore file
	privateKey, _, err := utils.GetAddress()
	if err != nil {
		log.Errorln("Failed to get address: %v", err)
		return
	}

	contractAddress, err := utils.ConvertToAddress("0xa78f62a91c7872dc3151529739e309cdf6aED887")
	if err != nil {
		log.Errorln("Failed to convert to address: %v", err)
		return
	}
	player1Address, err := utils.ConvertToAddress("0x1bc1447a4f9D24FE355523B771E32282F1Ca5ceC")
	if err != nil {
		log.Errorln("Failed to convert to address: %v", err)
		return
	}
	player2Address, err := utils.ConvertToAddress("0x1bc1447a4f9D24FE355523B771E32282F1Ca5ceC")
	if err != nil {
		log.Errorln("Failed to convert to address: %v", err)
		return
	}

	_, err = chain.CheckEventEmissions(client, contractAddress, privateKey, player1Address, player2Address)
	if err != nil {
		log.Errorln("Failed to check event emissions: %v", err)
		return
	}

	return

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
	matchHandler := handler.NewMatchHandler(rdb)
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
