package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {

	listNodes()

}

func listNodes() {
	err, clientset := Config.GetClientSet()

	if err != nil {
		panic(err.Error())
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
		log.Fatalln("failed to get nodes:", err)
	}

	for i, node := range nodes.Items {
		fmt.Printf("[%d} %s\n", i, node.GetName())
	}
}
