package chain

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const contractABI = `[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "player",
        "type": "address"
      }
    ],
    "name": "enterGame",
    "type": "event"
  }
]`

func checkEventEmissions(client *ethclient.Client, contractAddress common.Address, player1Address common.Address, player2Address common.Address) (bool, error) {
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return false, fmt.Errorf("Failed to parse contract ABI: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return false, fmt.Errorf("Failed to filter logs: %v", err)
	}

	enterGameEvent := parsedABI.Events["enterGame"]

	player1Entered := false
	player2Entered := false
	for _, log := range logs {
		event := new(struct {
			Player common.Address
		})

		err := enterGameEvent.ParseLog(*log, event)
		if err == nil {
			if event.Player == player1Address {
				player1Entered = true
			} else if event.Player == player2Address {
				player2Entered = true
			}
		}
	}

	if !player1Entered || !player2Entered {
		return false, nil
	}

	return true, nil
}
