package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//pods, err := clientset.CoreV1().Pods("namespace").List(metav1.ListOptions{LabelSelector: "app=my-app"})
	pods, err := clientset.CoreV1().Pods("namespace").List(nil, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.Name, pod.Status.Phase)
		//err := clientset.CoreV1().Pods("namespace").Delete(pod.Name, &metav1.DeleteOptions{})
		//if err != nil {
		//	panic(fmt.Errorf("failed to delete pod: %s", err))
		//}
	}
}
