package clientset

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	appV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"log"
)

//获取deployment信息
func GetDeploy(client *kubernetes.Clientset, ctx context.Context, deployName, namespace string) {
	// get deploy
	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployName, metaV1.GetOptions{})
	if err != nil {
		log.Println(err)
		zap.L().Error("dao.clientset.deployment", zap.Error(err))
	}
	fmt.Println("deployment name ===> ", deployment.Name)
}

//更新deployment数量
func UpdateDeployReplica(client *kubernetes.Clientset, ctx context.Context, deployName, namespace string, replicas int32) {
	// 1 方法一：更新deployment 副本数量
	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployName, metaV1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	// 设置副本数量
	deployment.Spec.Replicas = &replicas
	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metaV1.UpdateOptions{})

	// 2 方法二：更新副本数量的另一种方法
	replica, err := client.AppsV1().Deployments(namespace).GetScale(ctx, deployName, metaV1.GetOptions{})
	replica.Spec.Replicas = replicas
	replica, err = client.AppsV1().Deployments(namespace).UpdateScale(ctx, deployName, replica, metaV1.UpdateOptions{})
	fmt.Println("replica name ====>", replica.Name)
}

//更新deployment镜像
func UpdateDeployImage(client *kubernetes.Clientset, ctx context.Context, deployName, namespace, image string) {
	deployment, err := client.AppsV1().Deployments(namespace).Get(ctx, deployName, metaV1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	deployment.Spec.Template.Spec.Containers[0].Image = image
	deployment, err = client.AppsV1().Deployments(namespace).Update(ctx, deployment, metaV1.UpdateOptions{})
}

//删除deployment
func DeleteDeploy(client *kubernetes.Clientset, ctx context.Context, deployName, namespace string) {
	// 删除deployment
	err := client.AppsV1().Deployments(namespace).Delete(ctx, deployName, metaV1.DeleteOptions{})
	if err != nil {
		log.Println(err)
	}
}

//创建deployment和service
func CreateDeploy(client *kubernetes.Clientset, ctx context.Context, namespace string) {
	var replicas int32 = 3
	var targetPort int32 = 80
	intString := intstr.IntOrString{
		IntVal: targetPort,
	}
	deployment := &appV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "nginx",
			Labels: map[string]string{
				"app": "nginx",
			},
			Namespace: namespace,
		},
		Spec: appV1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.16.1",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
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

	deployment, err := client.AppsV1().Deployments(namespace).Create(ctx, deployment, metaV1.CreateOptions{})
	if err != nil {
		log.Println("err ===> ", err)
	}
	service, err = client.CoreV1().Services(namespace).Create(ctx, service, metaV1.CreateOptions{})
}
