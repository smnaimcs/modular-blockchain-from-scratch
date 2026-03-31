package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/smnaimcs/projectx/crypto"
	"github.com/smnaimcs/projectx/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	h := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp: time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("hello world"),
	}

	return NewBlock(h, []Transaction{tx})
}

func TestBlockHash(t *testing.T) {
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	b := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(0)
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, b.Sign(privKey))

	assert.Nil(t, b.Verify())
	anotherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = anotherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())
}