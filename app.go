/*
* Program prints out information of all running pods
* every 10 seconds.
 */

package main

import (
	"context"
	"fmt"  // Used for printing
	"time" // Used for sleeping

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func handleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// Create the in-cluster config
	config, err := rest.InClusterConfig()
	handleError(err)

	// Creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	handleError(err)

	for {
		// Omitting namespace (""), gets us pods from all namespaces
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		handleError(err)

		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		for _, pod := range pods.Items {
			fmt.Printf("Pod Name: %s\n", pod.Name)
			fmt.Printf("Namespace: %s\n", pod.Namespace)
			if pod.Status.StartTime != nil {
				fmt.Printf("Pod start Time: %s\n", pod.Status.StartTime.Time)
			}

			fmt.Println("Containers: ")
			for _, container := range pod.Spec.Containers {
				fmt.Printf("- %s\n", container.Name)
			}

			// Print all metadata
			fmt.Println("Metadata:")
			for key, value := range pod.ObjectMeta.Annotations {
				fmt.Printf("  Annotation: %s = %s\n", key, value)
			}
			for key, value := range pod.ObjectMeta.Labels {
				fmt.Printf("  Label: %s = %s\n", key, value)
			}

			fmt.Println("-------------------------------")
		}

		time.Sleep(10 * time.Second)

	}

}
