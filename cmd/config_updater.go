package cmd

import (
	"fmt"

	"github.com/evcc-io/evcc/util"
	"github.com/spf13/viper"
)

type configUpdater struct {
	conf config
}

// Reads the Sponsortoken
func (cu configUpdater) SponsorToken() string {
	fmt.Printf("viper.GetString(\"sponsortoken\"): %v\n", viper.GetString("sponsortoken"))
	return cu.conf.SponsorToken
}

func (cu configUpdater) SiteTitle() (string, error) {
	site := &struct {
		Title string
	}{}
	if err := util.DecodeOther(cu.conf.Site, site); err != nil {
		return "", err
	}
	return site.Title, nil
}

func (cu configUpdater) SetSiteTitle(title string) error {
	viper.Set("site.title", title)
	viper.WriteConfigAs("/tmp/evcc.yaml")
	return nil
}

func NewConfigUpdater(conf config) configUpdater {
	return configUpdater{
		conf: conf,
	}
}
