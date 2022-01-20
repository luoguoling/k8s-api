package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"k8s-api/dao/clientset"
	"k8s-api/logger"
	"k8s-api/settings"
)

func main() {
	//1.加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("配置文件初始化失败,err:%v\n", err)
		return
	} else {
		fmt.Println("配置文件加载成功!!!!")
	}
	//2.初始化配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("日志始化失败,err:%v\n", err)
		return
	}
	zap.L().Sync() //缓存区日志追加到日志中
	deployName := "myapp-deploy"
	namespace := "default"
	//image := "nginx:1.15-alpine"
	var replicas int32 = 2
	if err := clientset.Init(settings.Conf.K8sInfo); err != nil {
		fmt.Println("初始化失败!!!")
		return
	}
	ctx := context.Background()

	fmt.Println(clientset.ClientSet)
	//----------对pod操作---------------------------------
	//获取pod
	clientset.GetPods(clientset.ClientSet, ctx, "default")
	//删除pod

	// -------------对deployment操作-----------------------

	//获取deployment
	clientset.GetDeploy(clientset.ClientSet, ctx, deployName, namespace)
	//删除deployment
	//clientset.DeleteDeploy(clientset.ClientSet, ctx, deployName, namespace)
	//更新deployment镜像
	//clientset.UpdateDeployImage(clientset.ClientSet, ctx, deployName, namespace, image)
	//更新deployment副本数量
	clientset.UpdateDeployReplica(clientset.ClientSet, ctx, deployName, namespace, replicas)
	//创建deployment和service
	clientset.CreateDeploy(clientset.ClientSet, ctx, namespace)

	//---------对service操作----------------------
	//获取service
	clientset.GetService(clientset.ClientSet, ctx, "default")
	//创建service
	clientset.CreateService(clientset.ClientSet, ctx, "default")
	//修改service
	clientset.UpdateService(clientset.ClientSet, ctx, "default", "nginx", 30030)
	//删除service
	clientset.DeleteService(clientset.ClientSet, ctx, "default", "nginx")

	//----------对ingress操作--------------
	//获取ingress
	clientset.GetIngress(clientset.ClientSet, ctx, "default")
	//删除ingress
	clientset.DeleteIngress(clientset.ClientSet, ctx, "default", "nginx")
	//增加ingress对象
	clientset.CreateIngress(clientset.ClientSet, ctx, "default")
	//修改ingress对象
	clientset.UpdateIngress(clientset.ClientSet, ctx, "default", "nginx")
}
