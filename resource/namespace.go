package resource

import (
	"github.com/onesafe/k8s_practice/client"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func InitNamespaceLister() {
	NamespaceLister = KubeInformerFactory.Core().V1().Namespaces().Lister()
}

func GetNamespace(name string) (*v1.Namespace, error) {
	ns, err := NamespaceLister.Get(name)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func ListNamespace(selector labels.Selector) ([]*v1.Namespace, error) {
	nss, err := NamespaceLister.List(selector)
	if err != nil {
		return nil, err
	}
	return nss, nil
}

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