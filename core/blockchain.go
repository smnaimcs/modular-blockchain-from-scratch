package core

import (
	"fmt"
	"sync"
)

type Blockchain struct {
	store     Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

func NewBlockChain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWithoutValidation(genesis)
	if err != nil {
		return nil, err
	}

	return bc, nil
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given block height (%d) too high, Blockchain height (%d)", height, bc.Height())
	}

	bc.lock.Lock()
	defer bc.lock.Unlock()

	return bc.headers[height], nil
}


func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	bc.addBlockWithoutValidation(b)
	return nil
}

func (bc *Blockchain) Height() uint32 {
	bc.lock.RLock()
	bc.lock.RUnlock()
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.lock.Lock()
	bc.headers = append(bc.headers, b.Header)
	bc.lock.Unlock()
	return bc.store.Put(b)
}
