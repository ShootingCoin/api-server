package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// GetAddress returns the address and the private key in .keystore
func GetAddress() (*ecdsa.PrivateKey, string, error) {
	// Load env file
	err := godotenv.Load(".keystore")
	if err != nil {
		return nil, "", err
	}

	// Get private key and address from .keystore
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	address := os.Getenv("ADDRESS")

	if privateKeyHex == "" || address == "" {
		return nil, "", err
	}

	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, "", err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, "", err
	}

	// Validate the address
	if address != crypto.PubkeyToAddress(privateKey.PublicKey).Hex() {
		return nil, "", err
	}

	log.Printf("Address: %s\n", address)

	return privateKey, address, nil
}

func GetContractAddress() (string, error) {
	// Load env file
	err := godotenv.Load(".keystore")
	if err != nil {
		return "", err
	}

	// Get private key and address from .keystore
	address := os.Getenv("CONTRACT_ADDRESS")

	if address == "" {
		return "", err
	}

	return address, nil
}

func GetInfuraUrl() (string, error) {
	// Load env file
	err := godotenv.Load(".keystore")
	if err != nil {
		return "", err
	}

	// Get private key and address from .keystore
	url := os.Getenv("INFURA_URL")

	if url == "" {
		return "", err
	}

	return url, nil
}

func ConvertToAddress(hexString string) (common.Address, error) {
	// Step 1: Remove the "0x" prefix
	hexString = hexString[2:]

	// Step 2: Decode the hexadecimal string into a byte slice
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("Failed to decode hexadecimal string:", err)
		return common.Address{}, err
	}

	// Step 3: Convert the byte slice to common.Address
	address := common.BytesToAddress(bytes)

	return address, nil
}
