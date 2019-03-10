package resource

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/api/core/v1"
)

func RegisterConfigMapLister() {
	ConfigMapLister = KubeInformerFactory.Core().V1().ConfigMaps().Lister()
}


func GetConfigMap(namespace string, configmapname string) (*v1.ConfigMap, error) {
	configmap, err := ConfigMapLister.ConfigMaps(namespace).Get(configmapname)
	if err != nil {
		return nil, err
	}
	return configmap, nil
}

func ListConfigMap(selector labels.Selector) ([]*v1.ConfigMap, error){
	configmaps, err := ConfigMapLister.List(selector)
	if err != nil {
		return nil, err
	}
	return configmaps, nil
}