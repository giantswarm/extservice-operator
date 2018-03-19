package service

import (
	"github.com/giantswarm/extservice-operator/flag/service/kubernetes"
	"github.com/giantswarm/extservice-operator/flag/service/extservice"
)

type Service struct {
	ExtService extservice.ExtService
	Kubernetes kubernetes.Kubernetes
}
