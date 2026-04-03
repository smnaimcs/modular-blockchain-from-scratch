package network

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/smnaimcs/projectx/core"
	"github.com/smnaimcs/projectx/crypto"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime  time.Duration
}

type Server struct {
	ServerOpts
	mempool     *TxPool
	blockTime   time.Duration
	isValidator bool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}
	return &Server{
		ServerOpts:  opts,
		mempool:     NewTxPool(),
		blockTime:   opts.BlockTime,
		isValidator: opts.PrivateKey != nil,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			s.createNewBlock()
		}
	}

	fmt.Println("Server shutdown")
}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	hash := tx.Hash(core.TxHasher{})

	if s.mempool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("transaction already in mempool")

		return nil
	}

	logrus.WithFields(logrus.Fields{
		"hash": hash,
	}).Info("adding new tx to the mempool")

	return s.mempool.Add(tx)
}

func (s *Server) createNewBlock() error {
	fmt.Printf("creating new block")
	return nil
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
