package chain

import (
	"context"
	"crypto/ecdsa"
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

const (
	infuraURL = "https://polygon-mumbai.infura.io/v3/64606310f8694fff946ad5510df0b40e"
	keyString = "64606310f8694fff946ad5510df0b40e"
)

func CheckEventEmissions(client *ethclient.Client, contractAddress common.Address, privateKey *ecdsa.PrivateKey, player1Address common.Address, player2Address common.Address) (bool, error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Read the ABI
	data, err := ioutil.ReadFile("../../../abi/abi.json")
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
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{eventEnter.ID},
		},
	})
	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

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
	}

	return true, nil
}
