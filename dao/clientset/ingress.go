package clientset

import (
	"context"
	"fmt"
	//appV1 "k8s.io/api/apps/v1"
	//coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//列出ingress
func GetIngress(client *kubernetes.Clientset, ctx context.Context, namespace string) {
	ingressList, err := client.ExtensionsV1beta1().Ingresses(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		fmt.Println("dao.client.ingress getIngress获取错误!!!")
	}
	for _, ingress := range ingressList.Items {
		fmt.Println(ingress.ObjectMeta.Name, ingress.Namespace)
	}
}

//删除ingress
func DeleteIngress(client *kubernetes.Clientset, ctx context.Context, namespace string, ingressName string) {
	err := client.ExtensionsV1beta1().Ingresses(namespace).Delete(ctx, ingressName, metaV1.DeleteOptions{})
	if err != nil {
		fmt.Println("dao.clientset.ingress 删除失败!!!")
	}

}

//创建ingress
func CreateIngress(client *kubernetes.Clientset, ctx context.Context, namespace string) {

}

//更新ingress
func UpdateIngress(client *kubernetes.Clientset, ctx context.Context, namespace string, ingressName string) {

}
