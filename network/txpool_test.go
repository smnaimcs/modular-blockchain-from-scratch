package network

import (
	"testing"

	"github.com/smnaimcs/projectx/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAddTx(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)

	tx := core.NewTransaction([]byte("foo"))
	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}
