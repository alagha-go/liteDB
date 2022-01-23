package database

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

type ADDRESS struct {
	PrivateKey						ecdsa.PrivateKey
	PublicKey						[]byte
}

const (
	checksumLength = 4
	version =   byte(0x00)
)


func GenerateId() string {
	address := MakeAddress()
	id := address.Address()

	fmt.Println(string(id))

	return string(id)
}


func NewKeyPair() (ecdsa.PrivateKey, []byte){
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	HandleError(err)

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}


func Base58Encode(input []byte) []byte{
	encode := base58.Encode(input)

	return []byte(encode)
}



func MakeAddress() *ADDRESS {
	private, public := NewKeyPair()
	wallet := ADDRESS{PrivateKey: private, PublicKey: public}

	return &wallet
}


func PubKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	HandleError(err)

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}


func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}


func (address *ADDRESS) Address() []byte {
	pubHash := PubKeyHash(address.PublicKey)

	vasionedHash := append([]byte{version}, pubHash...)

	checksum := Checksum(vasionedHash)

	fullHash := append(vasionedHash, checksum...)

	Address := Base58Encode(fullHash)

	return Address
}



func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}