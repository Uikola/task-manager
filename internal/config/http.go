package config

import (
	"errors"
	"net"
	"os"
)

const (
	httpConfigHostKey = "HTTP_HOST"
	httpConfigPortKey = "HTTP_PORT"
)

var (
	errHTTPHostNotSet = errors.New("http host not set")
	errHTTPPortNotSet = errors.New("http port not set")
)

type HTTP interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (HTTP, error) {
	host := os.Getenv(httpConfigHostKey)
	if host == "" {
		return nil, errHTTPHostNotSet
	}

	port := os.Getenv(httpConfigPortKey)
	if port == "" {
		return nil, errHTTPPortNotSet
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
