package main

import (
	Config "KubeDesc/src/config"
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	listConfigMaps()

}

func listConfigMaps() {
	err, clientset := Config.GetClientSet()
	if err != nil {
		panic(err.Error())
	}
	configMaps, err := clientset.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i, configMap := range configMaps.Items {
		fmt.Printf("[%d] %s\n", i, configMap.GetName())
	}
}
