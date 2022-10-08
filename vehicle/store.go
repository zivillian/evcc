package vehicle

import (
	"github.com/evcc-io/evcc/server/db/settings"
)

type Store struct {
	key string
}

func NewStore(key string) *Store {
	return &Store{key: key}
}

func (s *Store) Load(res any) error {
	return settings.Json(s.key, &res)
}

func (s *Store) Save(val any) error {
	return settings.SetJson(s.key, val)
}
