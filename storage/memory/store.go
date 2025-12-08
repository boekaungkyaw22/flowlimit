package memory

import (
	"github.com/boekaungkyaw22/flowlimit/storage"
)

type Store struct{}

func NewMemoryStore() *Store {
	return &Store{}
}

func (s *Store) Get(key string) (storage.Entry, bool) {
	return storage.Entry{}, false
}

func (s *Store) Set(key string, entry storage.Entry) {
}
