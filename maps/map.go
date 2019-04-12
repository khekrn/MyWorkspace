package maps

import (
	"algorithms/iter"
	"errors"
)

// Error codes
var (
	ErrInvalidKey  = errors.New("invalid key, key cannot be empty")
	ErrKeyNotFound = errors.New("key does not exist")
)

// Map interface
type Map interface {
	Length() int

	Put(key string, value interface{}) error

	Get(key string) (interface{}, error)

	Delete(key string) (bool, error)

	Iter() iter.Iterator
}
