package main

import (
	Config "KubeDesc/src/config"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"time"
)

func main() {

	err, clientset := Config.GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	watchList := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		string(v1.ResourceServices),
		v1.NamespaceAll,
		fields.Everything(),
		)

	_, controller := cache.NewInformer(
		watchList,
		&v1.Service{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Println("Service 以添加.")
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Println("Service 以更改！")
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Println("Service 已删除！")
			},
		},
		)
	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}


	
}
