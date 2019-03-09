package resource

import (
	"fmt"
	apiv1beta1 "k8s.io/api/extensions/v1beta1"
	"github.com/onesafe/k8s_practice/client"
	"k8s.io/apimachinery/pkg/labels"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func RegisterDeploymentLister() {
	DeploymentLister = KubeInformerFactory.Extensions().V1beta1().Deployments().Lister()
}

func GetDeployment(namespace string, deploymentName string) (*apiv1beta1.Deployment, error) {
	deployment, err := DeploymentLister.Deployments(namespace).Get(deploymentName)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

func ListDeployment(namespace string, selector labels.Selector) ([]*apiv1beta1.Deployment, error){
	deployments, err := DeploymentLister.Deployments(namespace).List(selector)
	if err != nil {
		return nil, err
	}
	return deployments, nil
}

func TestDeploymentConn(namespace string) {
	Clientset, err := client.GetK8sClientSet();
	DeploysClient := Clientset.ExtensionsV1beta1().Deployments(namespace)
	listOptions := metaV1.ListOptions{}
	deploys, err := DeploysClient.List(listOptions)
	if err != nil {
		fmt.Println("Test Conn error: ", err.Error())
	}

	for _, d := range deploys.Items {
		fmt.Println(d.Name)
	}
}