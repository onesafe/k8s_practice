package resource

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)


func RegisterNodeLister() {
	NodeLister = KubeInformerFactory.Core().V1().Nodes().Lister()
}

func GetNode(nodename string) (*v1.Node, error) {
	node, err := NodeLister.Get(nodename)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func ListNode(selector labels.Selector) ([]*v1.Node, error) {
	nodes, err := NodeLister.List(selector)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}