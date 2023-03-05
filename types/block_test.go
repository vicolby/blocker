package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicolby/blocker/crypto"
	"github.com/vicolby/blocker/util"
)

func TestSignBlock(t *testing.T) {
	var (
		block   = util.RandomBlock()
		privKey = crypto.NewPrivateKey()
		pubKey  = privKey.PublicKey()
	)

	sig := SignBlock(privKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(HashBlock(block), pubKey))
}

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	assert.Equal(t, 32, len(hash))
}
