package vehicle

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/evcc-io/evcc/server/db/settings"
)

type Store struct {
	key string
}

func NewStore(key string, val ...string) *Store {
	s := &Store{key: key}

	if len(val) > 0 {
		s.hash(val...)
	}

	return s
}

func (s *Store) hash(val ...string) {
	mac := hmac.New(sha256.New, []byte(s.key))
	for _, val := range val {
		mac.Write([]byte(val))
	}
	s.key += "." + hex.EncodeToString(mac.Sum(nil))
}

func (s *Store) Load(res any) error {
	return settings.Json(s.key, &res)
}

func (s *Store) Save(val any) error {
	return settings.SetJson(s.key, val)
}
