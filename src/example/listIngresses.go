package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {
	listIngresses()
}

func listIngresses() {
	err, clientset := Config.GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	ingresses, err := clientset.ExtensionsV1beta1().Ingresses("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get ingress:", err)
	}
	for i, ingress := range ingresses.Items {
		fmt.Printf("[%d] %s\n", i, ingress.GetName())
	}
}
