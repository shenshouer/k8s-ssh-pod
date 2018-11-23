package k8s

import (
	"github.com/shenshouer/k8s-ssh-pod/config"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewK8SClient() (*kubernetes.Clientset, *rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", config.GetConfig().KubeConfig)
	if err != nil {
		return nil, nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, config, err
}

func GetNamespaces(client *kubernetes.Clientset) (*v1.NamespaceList, error) {
	return client.CoreV1().Namespaces().List(metav1.ListOptions{})
}

func GetPods(client *kubernetes.Clientset, namespace string) (*v1.PodList, error) {
	return client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}
