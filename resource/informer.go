package resource

import (
	"k8s.io/client-go/informers"
	"time"
	"fmt"
	"github.com/onesafe/k8s_practice/client"
)

var (
	KubeInformerFactory informers.SharedInformerFactory
)

func Factory() {
	Clientset, err := client.GetK8sClientSet();
	if err != nil {
		fmt.Print("Get k8s client error: " + err.Error())
	}

	KubeInformerFactory = informers.NewSharedInformerFactory(Clientset, time.Second*10)
}


func Start() {
	stopper := make(chan struct{})
	KubeInformerFactory.Start(stopper)
}