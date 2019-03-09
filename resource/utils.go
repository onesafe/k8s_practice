package resource

import (
	listerV1 "k8s.io/client-go/listers/core/v1"
	listerV1beta1 "k8s.io/client-go/listers/extensions/v1beta1"
)

var (
	DeploymentLister listerV1beta1.DeploymentLister
	NamespaceLister listerV1.NamespaceLister
	PodLister listerV1.PodLister
)