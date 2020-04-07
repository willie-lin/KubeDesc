package main

import (
	"KubeDesc/src/config"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

func main() {


	err, clientset := Config.GetClientSet()


	if err != nil{
		panic(err.Error())
	}

	//fmt.Println(aaa)
	watchPods(clientset)

}





func watchPods(clientset *kubernetes.Clientset) {
	//err, clientset := Config.GetClientSet()
	//if err != nil {
	//	panic(err.Error())
	//}

	for {

		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		namespace := "default"
		pod := "example-xxx"

		_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})

		if errors.IsNotFound(err) {
			fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %s in namespace %s: %v\n",
				pod, namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod %s in namespaces %s\n", pod, namespace)
		}
		time.Sleep(10 * time.Second)
	}
}







//func configString() string{
//
//	str, _ := os.Getwd()
//
//	kubeconfig := filepath.Join(str, "src/config")
//
//
//	return kubeconfig
//
//
//
//
//
//}

