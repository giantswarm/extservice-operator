package service

import (
	"context"

	"github.com/giantswarm/microendpoint/service/version"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/viper"

	"github.com/giantswarm/extservice-operator/flag"
)

type Config struct {
	Flag   *flag.Flag
	Logger micrologger.Logger
	Viper  *viper.Viper

	Description string
	GitCommit   string
	ProjectName string
	Source      string
}

type Service struct {
	Version *version.Service
}

func New(config Config) (*Service, error) {
	return &Service{}, nil
}

func (s *Service) Boot(ctx context.Context) {
}
