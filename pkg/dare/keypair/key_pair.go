package keypair

import (
	"crypto/rand"

	"golang.org/x/crypto/nacl/box"
)

// KeyPair ...
type KeyPair struct {
	Public  *[32]byte
	Private *[32]byte
}

// New ...
func New(public *[32]byte, private *[32]byte) KeyPair {
	return KeyPair{Public: public, Private: private}
}

// Random ...
func Random() (KeyPair, error) {
	publicKey, privateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return KeyPair{}, err
	}
	return KeyPair{Public: publicKey, Private: privateKey}, nil
}

// SharedKey ...
func (k *KeyPair) SharedKey(peer KeyPair) *[32]byte {
	sharedKey := new([32]byte)
	box.Precompute(sharedKey, peer.Public, k.Private)
	return sharedKey
}
