package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func main() {

	listNamespaces()
	fmt.Println("---------------------------")

}

func listNamespaces() {

	err, clientset := Config.GetClientSet()
	//fmt.Println(clientset)
	if err != nil {
		panic(err.Error())
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get name space:", err)
		panic(err.Error())
	}

	for i, ns := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, ns.GetName())
	}
}
