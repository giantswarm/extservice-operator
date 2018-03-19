package service

import (
	"github.com/giantswarm/extservice-operator/flag/service/extservice"
	"github.com/giantswarm/extservice-operator/flag/service/kubernetes"
)

type Service struct {
	ExtService extservice.ExtService
	Kubernetes kubernetes.Kubernetes
}
