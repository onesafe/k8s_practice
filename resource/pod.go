package resource

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)


func RegisterPodLister() {
	PodLister = KubeInformerFactory.Core().V1().Pods().Lister()
}

func GetPod(namespace string, podname string) (*v1.Pod, error) {
	pod, err := PodLister.Pods(namespace).Get(podname)
	if err != nil {
		return nil, err
	}
	return pod, nil
}

func ListPod(selector labels.Selector) ([]*v1.Pod, error) {
	pods, err := PodLister.List(selector)
	if err != nil {
		return nil, err
	}
	return pods, nil
}