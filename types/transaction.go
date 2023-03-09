package types

import (
	"crypto/sha256"

	"github.com/vicolby/blocker/crypto"
	"github.com/vicolby/blocker/proto"
	pb "google.golang.org/protobuf/proto"
)

func SignTransaction(tx *proto.Transaction, pk *crypto.PrivateKey) *crypto.Signature {
	hash := HashTransaction(tx)
	sig := pk.Sign(hash)
	return sig
}

func HashTransaction(tx *proto.Transaction) []byte {
	b, err := pb.Marshal(tx)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}

func VerifyTransaction(tx *proto.Transaction) bool {
	for _, input := range tx.Inputs {
		var (
			sig    = crypto.SignatureFromBytes(input.Signature)
			pubKey = crypto.PublicKeyFromBytes(input.PublicKey)
		)
		input.Signature = nil
		if !sig.Verify(HashTransaction(tx), pubKey) {
			return false
		}
	}
	return true
}
