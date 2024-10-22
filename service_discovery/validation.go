package main

import (
	"errors"
	"net/url"

	"github.com/roh4nyh/service_discovery/model"
)

func validateData(data model.Server) error {
	if data.Name == "" {
		return errors.New("name is required")
	}

	if data.Addr.IP == nil || (data.Addr.Port == 0 || data.Addr.Port >= 65535) {
		return errors.New("invalid service address")
	}

	if url, err := url.Parse(data.HealthCheck); err != nil || url.Scheme == "" || url.Host == "" {
		return errors.New("invalid health check endpoint")
	}

	return nil
}
