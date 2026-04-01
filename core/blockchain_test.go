package core

import (
	"testing"

	"github.com/smnaimcs/projectx/types"
	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(3))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlock := 1
	for i := 0; i < lenBlock; i++ {
		block := randomBlockWithSignature(t,uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock+1)
	assert.NotNil(t, bc.AddBlock(randomBlock(89, types.Hash{})))
}

func TestAddToHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3, types.Hash{})))
}
func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockChain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)
	return bc
}
