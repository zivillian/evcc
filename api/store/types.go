package store

import "github.com/evcc-io/evcc/api/store/internal"

var Key internal.ContextKey

// Provider creates a Persister for given string key
type Provider func(string) Store

// Store can load and store data
type Store interface {
	Load(any) error
	Save(any) error
}
