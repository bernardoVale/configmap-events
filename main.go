package main

import (
	"flag"
	"log"

	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	debug      = flag.Bool("debug", false, "enable debug logs")
)

func main() {
	flag.Parse()

	cfg, err := getConfig(*kubeconfig)

	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(cfg)
	events, err := clientset.CoreV1().ConfigMaps("default").Watch(v1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case ev, ok := <-events.ResultChan():
			if !ok {
				break
			}
			log.Printf("Object:%s, Type: %s", ev.Object, ev.Type)
		default:
		}
	}

}

func getConfig(cfg string) (*rest.Config, error) {
	if *kubeconfig == "" {
		return rest.InClusterConfig()
	}
	return clientcmd.BuildConfigFromFlags("", cfg)
}
