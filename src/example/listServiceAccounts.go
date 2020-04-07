package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	listServiceAccounts()

}

func listServiceAccounts() {
	err, clientset := Config.GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	serviceAccounts, err := clientset.CoreV1().ServiceAccounts("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i, serviceAccount := range serviceAccounts.Items {
		fmt.Printf("[%d] %s\n", i, serviceAccount.GetName())

	}
}
