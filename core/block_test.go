package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/smnaimcs/projectx/crypto"
	"github.com/smnaimcs/projectx/types"
	"github.com/stretchr/testify/assert"
)

func TestBlockHash(t *testing.T) {
	b := randomBlock(0, types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T) {
	b := randomBlock(0, types.Hash{})
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	b := randomBlock(0, types.Hash{})
	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, b.Sign(privKey))

	assert.Nil(t, b.Verify())
	anotherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = anotherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	h := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}

	return NewBlock(h, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevHash types.Hash) *Block {
	b := randomBlock(height, prevHash)

	privKey := crypto.GeneratePrivateKey()
	assert.Nil(t, b.Sign(privKey))

	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)

	return b
}
