package core

import (
	"fmt"
	"testing"

	"github.com/smnaimcs/projectx/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("hello world"),
	}
	assert.Nil(t, tx.Sign(privKey))
	fmt.Println(tx.Signature)
	assert.NotNil(t, tx.Signature)
}

func TestVerifySignature(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("hello world"),
	}
	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())
	anotherPrivKey := crypto.GeneratePrivateKey()
	tx.From = anotherPrivKey.PublicKey()
	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	tx := &Transaction{
		Data: []byte("hello world"),
	}

	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, tx.Sign(privKey))
	return tx
}