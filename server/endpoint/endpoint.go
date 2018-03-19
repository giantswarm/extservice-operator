package endpoint

import (
	versionendpoint "github.com/giantswarm/microendpoint/endpoint/version"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/extservice-operator/service"
)

type Config struct {
	Logger  micrologger.Logger
	Service *service.Service
}

func New(config Config) (*Endpoint, error) {
	var err error

	var versionEndpoint *versionendpoint.Endpoint
	{
		c := versionendpoint.Config{
			Logger:  config.Logger,
			Service: config.Service.Version,
		}

		versionEndpoint, err = versionendpoint.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	newEndpoint := &Endpoint{
		Version: versionEndpoint,
	}
	return newEndpoint, nil
}

// Endpoint is the endpoint collection.
type Endpoint struct {
	Version *versionendpoint.Endpoint
}
