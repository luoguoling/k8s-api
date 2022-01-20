package clientset

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"log"
)

//查询service
func GetService(client *kubernetes.Clientset, ctx context.Context, namespace string) {
	serviceList, err := client.CoreV1().Services(namespace).List(ctx, metaV1.ListOptions{})
	if err != nil {
		fmt.Println("获取serviceList出错!!!")
	}
	for _, service := range serviceList.Items {
		fmt.Println(service.ObjectMeta.Name, service.Namespace)
	}
}

//创建service
func CreateService(client *kubernetes.Clientset, ctx context.Context, namespace string) {
	var targetPort int32 = 80
	intString := intstr.IntOrString{
		IntVal: targetPort,
	}
	service := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "nginx",
			Labels: map[string]string{
				"app": "nginx",
			},
			Namespace: namespace,
		},
		Spec: coreV1.ServiceSpec{
			Type: coreV1.ServiceTypeNodePort,
			Ports: []coreV1.ServicePort{
				{
					Name:       "nginx",
					Port:       80,
					TargetPort: intString,
					NodePort:   30088,
					Protocol:   coreV1.ProtocolTCP,
				},
			},
			Selector: map[string]string{
				"app": "nginx",
			},
		},
	}
	service, err := client.CoreV1().Services(namespace).Create(ctx, service, metaV1.CreateOptions{})
	if err != nil {
		log.Println("err ===> ", err)
	}
}

//删除service
func DeleteService(client *kubernetes.Clientset, ctx context.Context, namespace string, servicename string) {
	err := client.CoreV1().Services(namespace).Delete(ctx, servicename, metaV1.DeleteOptions{})
	if err != nil {
		fmt.Println("servicename删除出错!!!")
	}
	zap.L().Info("dao.clientset.service 删除成功!!!")
}

//修改service,这个例子展示了service nodeport端口修改
func UpdateService(client *kubernetes.Clientset, ctx context.Context, namespace string, servicename string, port int32) {
	serviceobj, err := client.CoreV1().Services(namespace).Get(ctx, servicename, metaV1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	serviceobj.Spec.Ports[0].NodePort = port
	serviceobj, err = client.CoreV1().Services(namespace).Update(ctx, serviceobj, metaV1.UpdateOptions{})
	if err != nil {
		fmt.Println("更新service出现问题!!!")
	}
}
