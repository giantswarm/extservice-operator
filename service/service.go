package service

import (
	"context"
	"sync"

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

	bootOnce *sync.Once
}

func New(config Config) (*Service, error) {
	s := &Service{
		bootOnce: new(sync.Once),
	}

	return s, nil
}

func (s *Service) Boot(ctx context.Context) {
	s.bootOnce.Do(func() {
	})
}
