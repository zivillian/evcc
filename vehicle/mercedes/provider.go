package mercedes

import (
	"time"

	"github.com/evcc-io/evcc/provider"
)

// Provider implements the evcc vehicle api
type Provider struct {
	containerG func() (EVResponse, error)
}

// NewProvider provides the evcc vehicle api provider
func NewProvider(api *API, vin string, cache time.Duration) *Provider {
	impl := &Provider{
		containerG: provider.Cached(func() (EVResponse, error) {
			return api.Container(vin)
		}, cache),
	}
	return impl
}

// SoC implements the api.Vehicle interface
func (v *Provider) SoC() (float64, error) {
	res, err := v.containerG()
	if err == nil {
		return float64(res.SoC.Value), nil
	}

	return 0, err
}

// Range implements the api.VehicleRange interface
func (v *Provider) Range() (rng int64, err error) {
	res, err := v.containerG()
	if err == nil {
		return int64(res.RangeElectric.Value), nil
	}

	return 0, err
}
