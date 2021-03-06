package kubernetes

import (
	"github.com/giantswarm/extservice-operator/flag/service/kubernetes/tls"
)

type Kubernetes struct {
	Address   string
	InCluster string
	TLS       tls.TLS
}
