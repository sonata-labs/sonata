package pebble

import (
	"github.com/cockroachdb/pebble"
	"github.com/sonata-labs/sonata/store/chainstore"
)

type PebbleChainStore struct {
	db *pebble.DB
}

var _ chainstore.ChainStore = (*PebbleChainStore)(nil)

func NewPebbleChainStore(path string) (*PebbleChainStore, error) {
	db, err := pebble.Open(path, nil)
	if err != nil {
		return nil, err
	}
	return &PebbleChainStore{db: db}, nil
}
