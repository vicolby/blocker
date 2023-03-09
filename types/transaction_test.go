package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicolby/blocker/crypto"
	"github.com/vicolby/blocker/proto"
	"github.com/vicolby/blocker/util"
)

func TestNewTransacti(t *testing.T) {
	fromPrivKey := crypto.NewPrivateKey()
	fromAddress := fromPrivKey.PublicKey().Address().Bytes()

	toPrivKey := crypto.NewPrivateKey()
	toAddress := toPrivKey.PublicKey().Address().Bytes()

	input := &proto.TxInput{
		PrevTxHash:   util.RandomHash(),
		PrevOutIndex: 0,
		PublicKey:    fromPrivKey.PublicKey().Bytes(),
	}

	output1 := &proto.TxOutput{
		Amount:  5,
		Address: toAddress,
	}

	output2 := &proto.TxOutput{
		Amount:  95,
		Address: fromAddress,
	}

	tx := &proto.Transaction{
		Version: 1,
		Inputs:  []*proto.TxInput{input},
		Outputs: []*proto.TxOutput{output1, output2},
	}

	sig := SignTransaction(tx, fromPrivKey)
	input.Signature = sig.Bytes()

	assert.True(t, VerifyTransaction(tx))

}
