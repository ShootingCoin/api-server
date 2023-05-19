package utils

import (
	"encoding/hex"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// GetAddress returns the address and the private key in .env
func GetAddress() (string, error) {
	// Load env file
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	// Get private key and address from .env
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	address := os.Getenv("ADDRESS")

	if privateKeyHex == "" || address == "" {
		return "", err
	}

	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", err
	}

	// Validate the address
	if address != crypto.PubkeyToAddress(privateKey.PublicKey).Hex() {
		return "", err
	}

	log.Printf("Private Key: %s\n", privateKeyHex)
	log.Printf("Address: %s\n", address)

	return address, nil
}
