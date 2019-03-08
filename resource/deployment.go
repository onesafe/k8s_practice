package resource

import (
	"time"
	"fmt"
	apiv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/informers"
	"github.com/onesafe/k8s_practice/client"
	listerV1beta1 "k8s.io/client-go/listers/extensions/v1beta1"
	"gitlab.4pd.io/pht3/aol/pkg/signals"
)

var (
	deploymentLister listerV1beta1.DeploymentLister
)

func init() {
	fmt.Println("init k8s client")
	clientset, err := client.GetK8sClientSet();
	if err != nil {
		fmt.Print("Get k8s client error: " + err.Error())
	}

	deploymentInformerFactory := informers.NewSharedInformerFactory(clientset, time.Minute*10)

	deployInformer := deploymentInformerFactory.Extensions().V1beta1().Deployments()

	stopCh := signals.SetupSignalHandler()
	deploymentInformerFactory.Start(stopCh)

	deploymentLister = deployInformer.Lister()
}

func GetDeployment(namespace string, deploymentName string) (*apiv1beta1.Deployment, error) {
	deployment, err := deploymentLister.Deployments(namespace).Get(deploymentName)
	if err != nil {
		return nil, err
	}
	return deployment, nil
}