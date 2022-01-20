package clientset

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

//获取pods
func GetPods(client *kubernetes.Clientset, ctx context.Context, namespace string) {
	// get pod
	pods, err := client.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Println("获取pod出错!!!")
		log.Panic(err)
	}
	fmt.Println("pod Name ===> ", pods.Items[0].Status.ContainerStatuses[0].Name)
	fmt.Println("pod Image ===> ", pods.Items[0].Status.ContainerStatuses[0].Image)
	fmt.Println("pod State ===> ", pods.Items[0].Status.ContainerStatuses[0].State.Running)
}

//
