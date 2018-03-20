package node

import (
	"sync"

	"github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/operatorkit/framework"
	"github.com/giantswarm/operatorkit/informer"
	"k8s.io/client-go/kubernetes"

	"github.com/giantswarm/azure-operator/service/azureconfig/v1"
)

type FrameworkConfig struct {
	G8sClient versioned.Interface
	K8sClient kubernetes.Interface
	Logger    micrologger.Logger

	ProjectName string
}

type Framework struct {
	logger micrologger.Logger

	framework *framework.Framework
	bootOnce  sync.Once
}

func NewFramework(config FrameworkConfig) (*framework.Framework, error) {
	var err error

	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.ProjectName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ProjectName must not be empty", config)
	}

	var newInformer *informer.Informer
	{
		c := informer.Config{
			ResyncPeriod: informer.DefaultResyncPeriod,
			Watcher:      config.K8sClient.CoreV1().Nodes(),
		}

		newInformer, err = informer.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var v1ResourceSet *framework.ResourceSet
	{
		c := v1.ResourceSetConfig{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,

			ProjectName: config.ProjectName,
		}

		v1ResourceSet, err = v1.NewResourceSet(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var resourceRouter *framework.ResourceRouter
	{
		c := framework.ResourceRouterConfig{
			Logger: config.Logger,

			ResourceSets: []*framework.ResourceSet{
				v1ResourceSet,
			},
		}

		resourceRouter, err = framework.NewResourceRouter(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var f *framework.Framework
	{
		c := framework.Config{
			Informer:       newInformer,
			Logger:         config.Logger,
			ResourceRouter: resourceRouter,
		}

		f, err = framework.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	return f, nil
}
