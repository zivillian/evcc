package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigUpdater struct {
	conf config
}

// Reads the Sponsortoken
func (cu *ConfigUpdater) SponsorToken(name string) string {
	fmt.Printf("viper.GetString(\"sponsortoken\"): %v\n", viper.GetString("sponsortoken"))
	return cu.conf.SponsorToken
}

func NewConfigUpdater(conf config) ConfigUpdater {
	return ConfigUpdater{
		conf: conf,
	}
}
