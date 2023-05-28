package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	// Build the .keystore content
	envContent := strings.Join([]string{
		"PRIVATE_KEY=" + privateKeyHex,
		"ADDRESS=" + address,
	}, "\n")

	// Save to .keystore file
	err = ioutil.WriteFile(".keystore", []byte(envContent), 0600)
	if err != nil {
		log.Fatal(err)
	}
}
