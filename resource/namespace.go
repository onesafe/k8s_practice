package resource

import (
	"github.com/onesafe/k8s_practice/client"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
)

func TestNamespaceConn() {
	Clientset, err := client.GetK8sClientSet();
	NSClient := Clientset.CoreV1().Namespaces()
	listOptions := metaV1.ListOptions{}
	ns, err := NSClient.List(listOptions)
	if err != nil {
		fmt.Println("Test Conn error: ", err.Error())
	}
	for n, e := range ns.Items {
		fmt.Println(n)
		fmt.Println(e.Name)
		fmt.Println()
	}
}