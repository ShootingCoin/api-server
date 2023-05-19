package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"os"

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

	log.Printf("Private Key: %s\n", privateKeyHex)
	log.Printf("Address: %s\n", address)

	return privateKey, address, nil
}
