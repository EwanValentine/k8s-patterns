package patterns

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var singleK8sClientInstance *kubernetes.Clientset
var once sync.Once

// NewClient returns a new Kubernetes client
func NewClient() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	inCluster := os.Getenv("IN_CLUSTER")
	if inCluster == "true" {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		home := homedir.HomeDir()
		p := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", p)
		if err != nil {
			panic(err.Error())
		}
	}

	return kubernetes.NewForConfig(config)
}

// NewClientInstance -
func NewClientInstance() *kubernetes.Clientset {
	once.Do(func() {
		var err error
		singleK8sClientInstance, err = NewClient()
		if err != nil {
			log.Panic(err)
		}
	})
	return singleK8sClientInstance
}
