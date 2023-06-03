package chain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ShootingCoin/api-server/entity"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	log "github.com/sirupsen/logrus"
)

func CheckEventEmissions(client *ethclient.Client, txConfig entity.TxConfig) (bool, error) {
	// Derive the caller (transaction sender) public address
	publicKey := txConfig.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the next nonce for the transaction
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Get the suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create an authorized transactor
	auth := bind.NewKeyedTransactor(txConfig.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Read the ABI
	data, err := ioutil.ReadFile("./abi/abi.json")
	if err != nil {
		log.Fatalf("Failed to read ABI: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	eventEnter := parsedABI.Events["Entered"]

	log.Printf("Event Enter: %+v\n", eventEnter)

	// Check 'Entered' event
	filter, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Addresses: []common.Address{txConfig.ContractAddress},
		Topics: [][]common.Hash{
			{eventEnter.ID},
		},
	})
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	foundEvent := false

	for _, vLog := range filter {
		event := struct {
			User    common.Address
			BetInfo struct {
				CoinAddress common.Address
				BetAmount   *big.Int
				NftSkinId   [5]*big.Int
			}
		}{}
		err := parsedABI.UnpackIntoInterface(&event, "Entered", vLog.Data)
		if err != nil {
			log.Fatalf("Failed to unpack: %v", err)
		}

		log.Printf("User: %s, CoinAddress: %s, BetAmount: %s\n", event.User.Hex(), event.BetInfo.CoinAddress.Hex(), event.BetInfo.BetAmount.String())
		for i, skinId := range event.BetInfo.NftSkinId {
			log.Printf("NftSkinId[%d]: %s\n", i, skinId.String())
		}

		foundEvent = true
	}

	return foundEvent, nil
}

func StartGame(client *ethclient.Client, txConfig entity.TxConfig, gameId uuid.UUID, user1Address common.Address, user2Address common.Address) error {
	// Derive the caller (transaction sender) public address
	publicKey := txConfig.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the next nonce for the transaction
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Get the suggested gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Get the chain id
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain id: %v", err)
	}

	// Create an authorized transactor
	auth := bind.NewKeyedTransactor(txConfig.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Read the ABI
	data, err := ioutil.ReadFile("./abi/abi.json")
	if err != nil {
		log.Fatalf("Failed to read ABI: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	// Create input data payload
	input, err := parsedABI.Pack("startGame", new(big.Int).SetBytes(gameId[:]), user1Address, user2Address)
	if err != nil {
		log.Fatalf("Failed to pack data for StartGame: %v", err)
	}

	// Create new transaction (not yet signed)
	tx := types.NewTransaction(nonce, txConfig.ContractAddress, big.NewInt(0), uint64(300000), gasPrice, input)

	// Sign transaction using EIP-155 signer
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), txConfig.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // print the transaction hash

	return nil
}
