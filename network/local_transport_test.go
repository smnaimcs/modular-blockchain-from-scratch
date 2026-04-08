package network

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.consumeCh
	b, err := io.ReadAll(rpc.Payload)
	assert.Nil(t, err)

	assert.Equal(t, b, msg)
	assert.Equal(t, rpc.From, tra.addr)
}

func TestBroadcast(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	trc := NewLocalTransport("C")

	tra.Connect(trb)
	tra.Connect(trc)

	msg := []byte("foo")
	assert.Nil(t, tra.Broadcast(msg))

	rpcB := <-trb.Consume()
	b, err := io.ReadAll(rpcB.Payload)
	assert.Nil(t, err)
	assert.Equal(t, b, msg)

	rpcC := <-trc.Consume()
	c, err := io.ReadAll(rpcC.Payload)
	assert.Nil(t, err)
	assert.Equal(t, c, msg)

}
