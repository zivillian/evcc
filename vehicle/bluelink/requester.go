package bluelink

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Requester interface {
	Request(*http.Request) error
}

type requester struct {
	deviceID string
	config   Config
	ts       oauth2.TokenSource
}

func NewRequester(config Config, ts oauth2.TokenSource, deviceID string) *requester {
	return &requester{
		config:   config,
		ts:       ts,
		deviceID: deviceID,
	}
}

// Request decorates requests with authorization headers
func (v *requester) Request(req *http.Request) error {
	stamp, err := Stamps[v.config.CCSPApplicationID].Get()
	if err != nil {
		return err
	}

	token, err := v.ts.Token()
	if err != nil {
		return err
	}

	for k, v := range map[string]string{
		"Authorization":       "Bearer " + token.AccessToken,
		"ccsp-device-id":      v.deviceID,
		"ccsp-application-id": v.config.CCSPApplicationID,
		"offset":              "1",
		"User-Agent":          "okhttp/3.10.0",
		"Stamp":               stamp,
	} {
		req.Header.Set(k, v)
	}

	return nil
}
