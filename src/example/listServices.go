package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	listServices()
}

func listServices() {
	err, clienteset := Config.GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	services, err := clienteset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i, service := range services.Items {
		fmt.Printf("[%d] %s\n", i, service.GetName())
	}
}
