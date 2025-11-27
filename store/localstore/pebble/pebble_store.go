package pebble

import "github.com/cockroachdb/pebble"

type PebbleLocalStore struct {
	db *pebble.DB
}

func NewPebbleLocalStore(path string) (*PebbleLocalStore, error) {
	db, err := pebble.Open(path, nil)
	if err != nil {
		return nil, err
	}
	return &PebbleLocalStore{db: db}, nil
}

func (s *PebbleLocalStore) Close() error {
	return s.db.Close()
}
