package main

import (
	"fmt"
	"github.com/onesafe/k8s_practice/resource"
	"k8s.io/apimachinery/pkg/labels"
	"time"
)

func main() {
	fmt.Println("starting...")

	// 构造informer Factory
	resource.Factory()

	// 通过informerFactory注册 Lister
	resource.RegisterDeploymentLister()
	resource.RegisterNamespaceLister()
	resource.RegisterPodLister()

	// informer开始跑，[这一步必须在注册Lister之后，这样才能找到Lister，不然的话找不到Lister直接返回了]
	resource.Start()

	// 这里等待几秒钟(具体时间看情况)，等待cache中同步数据，[这是个坑啊，一开始一直查不到数据]
	time.Sleep(time.Second*3)

	pasPod, err := resource.GetDeployment("prophet-ee-dev", "pas")
	if err != nil {
		fmt.Println("Get Deployment error: ", err.Error())
	} else {
		fmt.Println(pasPod.CreationTimestamp)
	}

	deployments, err := resource.ListDeployment("prophet-ee-dev", labels.Everything()	)
	if err != nil {
		fmt.Println("List Deployments error: ", err.Error())
	} else {
		for i, d := range deployments {
			fmt.Println(i, d.Name)
		}
	}

	ns, err := resource.GetNamespace("kube-system")
	if err != nil {
		fmt.Println("Get Namespace error: ", err.Error())
	} else {
		fmt.Println(ns.Name)
	}

	nss, err := resource.ListNamespace(labels.Everything())
	if err != nil {
		fmt.Println("List Namespace error: ", err.Error())
	} else {
		for i, ns := range nss {
			fmt.Println(i, ns.Name)
		}
	}

	pod, err := resource.GetPod("kube-system", "core-dns")
	if err != nil {
		fmt.Println("Get Pod error: ", err.Error())
	} else {
		fmt.Println(pod)
	}

	pods, err := resource.ListPod(labels.Everything())
	if err != nil {
		fmt.Println("List Pod error: ", err.Error())
	} else {
		for i:=0; i<5; i++ {
			fmt.Println(pods[i].Name)
		}
	}

}
