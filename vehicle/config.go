package vehicle

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/api/store"
	"github.com/evcc-io/evcc/server/db"
	"github.com/evcc-io/evcc/server/db/settings"
	"github.com/evcc-io/evcc/util"
)

const (
	expiry   = 5 * time.Minute  // maximum response age before refresh
	interval = 15 * time.Minute // refresh interval when charging
)

type (
	vehicleRegistry map[string]factoryFunc
	factoryFunc     func(context.Context, map[string]any) (api.Vehicle, error)
)

func withContext(f func(map[string]any) (api.Vehicle, error)) factoryFunc {
	return func(_ context.Context, other map[string]any) (api.Vehicle, error) {
		return f(other)
	}
}

func (r vehicleRegistry) Add(name string, factory factoryFunc) {
	if _, exists := r[name]; exists {
		panic(fmt.Sprintf("cannot register duplicate vehicle type: %s", name))
	}
	r[name] = factory
}

func (r vehicleRegistry) Get(name string) (factoryFunc, error) {
	factory, exists := r[name]
	if !exists {
		return nil, fmt.Errorf("vehicle type not registered: %s", name)
	}
	return factory, nil
}

var registry vehicleRegistry = make(map[string]factoryFunc)

// Types returns the list of vehicle types
func Types() []string {
	var res []string
	for typ := range registry {
		res = append(res, typ)
	}
	return res
}

// NewFromConfig creates vehicle from configuration
func NewFromConfig(typ string, other map[string]interface{}) (v api.Vehicle, err error) {
	var cc struct {
		Cloud bool
		Other map[string]interface{} `mapstructure:",remain"`
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Cloud {
		cc.Other["brand"] = typ
		typ = "cloud"
	}

	ctx := context.Background()
	if db.Instance == nil {
		ctx = context.WithValue(ctx, store.Key, store.Provider(func(string) store.Store { return nil }))
	} else {
		ctx = context.WithValue(ctx, store.Key, store.Provider(settings.NewStore))
	}

	factory, err := registry.Get(strings.ToLower(typ))
	if err == nil {
		if v, err = factory(ctx, cc.Other); err != nil {
			err = fmt.Errorf("cannot create vehicle '%s': %w", typ, err)
		}
	} else {
		err = fmt.Errorf("invalid vehicle type: %s", typ)
	}

	return
}
