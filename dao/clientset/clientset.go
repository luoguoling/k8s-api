//获取clientset对象
package clientset

import (
	"fmt"
	"k8s-api/settings"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var ClientSet *kubernetes.Clientset

func Init(cfg *settings.K8sInfo) (err error) {
	config := rest.Config{
		//Host:        "http://121.5.106.77:9090", //代理地址
		Host: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		//BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IkNTaDRNUk1aSEs4YnBEVm5fZGw4RFZoN3VZQ3pkdV9mRHVmOGctWEVhVGsifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrOHMtYXV0aG9yaXplLXRva2VuLWdqNWJtIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6Ims4cy1hdXRob3JpemUiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiIzYTcwNzI4Ni1mZDMyLTQ4ZDEtYmU2Yi03YjAxOGFmNWUxMmIiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZS1zeXN0ZW06azhzLWF1dGhvcml6ZSJ9.VAnQsm2oLxNIab0SpmAkKO3FgaSGjSWs24LZ_gh08nXcsps40_DDTJzUG2jFjOCAluOOUz2EzbuVbud7EN9wOSbkA7-DaBDe6v009HrFWZ0mWt3MUG2uEzFJCRP7v5ySYMtNGb8ORX-68UvVvOCGHrN0dHH2IAwtke6U9npg_sWU_wHX835C-NF05qWGk2n3dlVBFsCq6U6ntVFhEJnq48vAZA3RfMPHkEha8xKroSERSVQkbi28EVKaepimF9-LV5RBY4bzbjz8fCcC9ikvW2goggcQx4getIC9DR0NmB3qybfPdZ7ltWCOiE3lFWwELk0Rd4geb9CpWdbLojn_ug",
		BearerToken: cfg.Token,
	}
	ClientSet, err = kubernetes.NewForConfig(&config)
	fmt.Println(ClientSet)
	if err != nil {
		fmt.Println("clientset出错!!!")
		panic(err.Error())

	}
	return
}
