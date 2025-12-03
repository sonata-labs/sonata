package localstore

import "github.com/cockroachdb/pebble"

type LocalStore struct {
	db *pebble.DB
}

func NewLocalStore(path string) (*LocalStore, error) {
	db, err := pebble.Open(path, nil)
	if err != nil {
		return nil, err
	}
	return &LocalStore{db: db}, nil
}
