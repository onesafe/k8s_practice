package k8s_practice

import (
	"k8s.io/client-go/informers"
	"time"
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	listerV1 "k8s.io/client-go/listers/core/v1"
)

var (
	podLister listerV1.PodLister
)

func init() {
	clientset, err := GetK8sClientSet();
	if err != nil {
		fmt.Print("Get k8s client error" + err.Error())
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(clientset, time.Minute*10)

	stopCh := make(chan struct{})
	sharedInformerFactory.Start(stopCh)

	podLister = sharedInformerFactory.Core().V1().Pods().Lister()
}

func GetPod(namespace string, podname string) *coreV1.Pod {
	pod, err := podLister.Pods(namespace).Get(podname)
	if err != nil {
		fmt.Println("Get Pod error" + err.Error())
	}
	return pod
}