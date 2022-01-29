package database


import (
	"github.com/ethereum/go-ethereum/crypto"
)


func PrivateKey() string {
	privKey, _ := crypto.GenerateKey()
	return crypto.PubkeyToAddress(privKey.PublicKey).Hex()
}