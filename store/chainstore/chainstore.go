package chainstore

import (
	"github.com/cockroachdb/pebble"
	accountv1 "github.com/sonata-labs/sonata/gen/account/v1"
	"google.golang.org/protobuf/proto"
)

type ChainStore struct {
	db *pebble.DB
}

func NewChainStore(path string) (*ChainStore, error) {
	db, err := pebble.Open(path, nil)
	if err != nil {
		return nil, err
	}
	return &ChainStore{db: db}, nil
}

func (c *ChainStore) StartBatch() (*pebble.Batch, error) {
	return c.db.NewBatch(), nil
}

func (c *ChainStore) StoreAccount(batch *pebble.Batch, account *accountv1.Account) error {
	address := account.Address
	accountBytes, err := proto.Marshal(account)
	if err != nil {
		return err
	}
	return batch.Set([]byte(address), accountBytes, nil)
}
