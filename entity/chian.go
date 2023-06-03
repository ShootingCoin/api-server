package entity

import (
	"crypto/ecdsa"
	"github.com/ShootingCoin/api-server/utils"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common"
)

type TxConfig struct {
	PrivateKey      *ecdsa.PrivateKey `json:"privateKey"`
	ServerAddress   string            `json:"address"`
	ContractAddress common.Address    `json:"contractAddress"`
}

func SetTxConfig() (TxConfig, error) {
	var txConfig TxConfig
	// Retrieve the private key and address from the keystore file
	privateKey, address, err := utils.GetAddress()
	if err != nil {
		log.Errorln("Failed to get address: %v", err)
		return txConfig, err
	}
	txConfig.PrivateKey = privateKey
	txConfig.ServerAddress = address

	// Retrieve the contract address from the keystore file
	scAddress, err := utils.GetContractAddress()
	if err != nil {
		log.Errorln("Failed to get contract address: %v", err)
		return txConfig, err
	}
	contractAddress, err := utils.ConvertToAddress(scAddress)
	if err != nil {
		log.Errorln("Failed to convert to address: %v", err)
		return txConfig, err
	}
	txConfig.ContractAddress = contractAddress
	return TxConfig{
		PrivateKey:      privateKey,
		ServerAddress:   address,
		ContractAddress: contractAddress,
	}, nil
}
