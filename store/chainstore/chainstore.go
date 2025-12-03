package chainstore

import (
	"fmt"

	"github.com/cockroachdb/pebble"
)

type ChainStore struct {
	db    *pebble.DB
	batch *pebble.Batch

	writer pebble.Writer
	reader pebble.Reader
}

func NewChainStore(path string) (*ChainStore, error) {
	db, err := pebble.Open(path, nil)
	if err != nil {
		return nil, err
	}
	return &ChainStore{db: db, writer: db, reader: db}, nil
}

// Returns a new chain store instance with the writer and reader set to a new batch.
func (c *ChainStore) Batch() *ChainStore {
	batch := c.db.NewBatch()
	return &ChainStore{db: c.db, batch: batch, writer: batch, reader: batch}
}

func (c *ChainStore) Commit() error {
	if c.batch == nil {
		return fmt.Errorf("batch not started")
	}
	return c.batch.Commit(nil)
}

func (c *ChainStore) RequireBatch() error {
	if c.batch == nil {
		return fmt.Errorf("batch not started")
	}
	return nil
}

func (c *ChainStore) Close() error {
	return c.db.Close()
}
