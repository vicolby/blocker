package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
)

const (
	privKeyLength = 64
	pubKeyLength  = 32
	addressLength = 20
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKey() *PrivateKey {
	_, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{key: key}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(message []byte) *Signature {
	return &Signature{ed25519.Sign(p.key, message)}
}

func (p *PrivateKey) PublicKey() *PublicKey {
	b := make([]byte, pubKeyLength)
	copy(b, p.key[privKeyLength-pubKeyLength:])
	return &PublicKey{key: b}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func (p *PublicKey) Address() *Address {
	return &Address{p.key[pubKeyLength-addressLength:]}
}

type Signature struct {
	value []byte
}

func (s *Signature) Bytes() []byte {
	return s.value
}

func (s *Signature) Verify(message []byte, pubKey *PublicKey) bool {
	return ed25519.Verify(pubKey.key, message, s.value)
}

type Address struct {
	value []byte
}

func (a *Address) String() string {
	return hex.EncodeToString(a.value)
}

func (a *Address) Bytes() []byte {
	return a.value
}
