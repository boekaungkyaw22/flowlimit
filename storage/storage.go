package storage

import "time"


type Entry struct {
	State any
	ExpiresAt time.Time
}


type Store interface {
	Get(key string) (Entry, bool)
	Set(key string, entry Entry) 
}
