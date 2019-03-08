package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	KUBECONFIGPATH = "~/.kube/devops.kubeconfig"
	K8SAPIINIT bool
	K8SAPI *kubernetes.Clientset
)

func SetConfig(ConfigPath string) {
	if ConfigPath == "" {
		return
	}

	KUBECONFIGPATH = ConfigPath
}

func GetK8sClientSet() (*kubernetes.Clientset, error) {
	if K8SAPIINIT {
		return K8SAPI, nil
	}

	conf, err := rest.InClusterConfig()
	if err != nil {
		if KUBECONFIGPATH != "" {
			conf, err = clientcmd.BuildConfigFromFlags("", KUBECONFIGPATH)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	K8SAPI, err = kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	K8SAPIINIT = true
	return K8SAPI, err
}
