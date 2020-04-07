package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	listSecrets()

}

func listSecrets() {
	err, clientset := Config.GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	secrets, err := clientset.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})

	for i, secret := range secrets.Items {
		fmt.Printf("[%d] %s\n", i, secret.GetName())
	}
}
