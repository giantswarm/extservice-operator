// Package server provides a server implementation to connect network transport
// protocols and service business logic by defining server endpoints.
package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/giantswarm/microerror"
	microserver "github.com/giantswarm/microkit/server"
	"github.com/giantswarm/micrologger"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/giantswarm/extservice-operator/server/endpoint"
	"github.com/giantswarm/extservice-operator/service"
)

type Config struct {
	Service *service.Service

	MicroServerConfig microserver.Config
}

func New(config Config) (microserver.Server, error) {
	var err error

	if config.Service == nil {
		return nil, microerror.Maskf(invalidConfigError, "config.Service must not be empty")
	}

	var endpointCollection *endpoint.Endpoint
	{
		c := endpoint.Config{
			Logger:  config.MicroServerConfig.Logger,
			Service: config.Service,
		}

		endpointCollection, err = endpoint.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newServer := &server{
		logger: config.MicroServerConfig.Logger,

		bootOnce:     sync.Once{},
		config:       config.MicroServerConfig,
		shutdownOnce: sync.Once{},
	}

	// Apply internals to the micro server config.
	newServer.config.Endpoints = []microserver.Endpoint{
		endpointCollection.Version,
	}
	newServer.config.ErrorEncoder = newServer.newErrorEncoder()

	return newServer, nil
}

type server struct {
	logger micrologger.Logger

	bootOnce     sync.Once
	config       microserver.Config
	shutdownOnce sync.Once
}

func (s *server) Boot() {
	s.bootOnce.Do(func() {
		// Here goes your custom boot logic for your server/endpoint if
		// any.
	})
}

func (s *server) Config() microserver.Config {
	return s.config
}

func (s *server) Shutdown() {
	s.shutdownOnce.Do(func() {
		// Here goes your custom shutdown logic for your
		// server/endpoint if any.
	})
}

func (s *server) newErrorEncoder() kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		rErr := err.(microserver.ResponseError)
		uErr := rErr.Underlying()

		rErr.SetCode(microserver.CodeInternalError)
		rErr.SetMessage(uErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
