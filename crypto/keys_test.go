package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrivateKey(t *testing.T) {
	privKey := NewPrivateKey()
	pubKey := privKey.PublicKey()

	assert.Equal(t, privKeyLength, len(privKey.Bytes()))
	assert.Equal(t, pubKeyLength, len(pubKey.Bytes()))
}

func TestPrivKeySign(t *testing.T) {
	msg := []byte("hello world")
	privKey := NewPrivateKey()
	pubKey := privKey.PublicKey()
	sig := privKey.Sign(msg)

	assert.True(t, sig.Verify(msg, pubKey))
	// Test with invalid message
	assert.False(t, sig.Verify([]byte("hello world!"), pubKey))
	// Test with invalid public key
	assert.False(t, sig.Verify(msg, NewPrivateKey().PublicKey()))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := NewPrivateKey()
	pubKey := privKey.PublicKey()
	addr := pubKey.Address()

	assert.Equal(t, addressLength, len(addr.Bytes()))
}
