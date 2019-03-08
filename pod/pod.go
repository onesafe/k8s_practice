package pod

import (
	"k8s.io/client-go/informers"
	"time"
	"fmt"
	apiV1 "k8s.io/client-go/pkg/api/v1"
	listerV1 "k8s.io/client-go/listers/core/v1"
	"github.com/onesafe/k8s_practice/client"
)

var (
	podLister listerV1.PodLister
)

func init() {
	clientset, err := client.GetK8sClientSet();
	if err != nil {
		fmt.Print("Get k8s client error" + err.Error())
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(clientset, time.Minute*10)

	stopCh := make(chan struct{})
	sharedInformerFactory.Start(stopCh)

	podLister = sharedInformerFactory.Core().V1().Pods().Lister()
}

func GetPod(namespace string, podname string) *apiV1.Pod {
	pod, err := podLister.Pods(namespace).Get(podname)
	if err != nil {
		fmt.Println("Get Pod error" + err.Error())
	}
	return pod
}