"# k8s-api" 
# 开发前准备工作
# k8s增删改查

## 一.获取clientset对象

### 1.1创建admin账户

```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: k8s-authorize
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: k8s-authorize
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: k8s-authorize
  namespace: kube-system
```

### 1.2 获取admin的token

```bash
 #获取token,对api进行操作
 kubectl describe secrets $(kubectl get secrets -n kube-system |grep admin |cut -f1 -d ' ') -n kube-system |grep -E '^token' |cut -f2 -d':'|tr -d '\t'|tr -d ' '
```

### 1.3创建代理并在k8s-master运行

```go
package main
import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)
func main() {
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true, //忽略证书验证
	}
	var transport http.RoundTripper = &http.Transport{
		Proxy:                  nil,
		DialContext:            nil,
		Dial:                   nil,
		DialTLSContext:         nil,
		DialTLS:                nil,
		TLSClientConfig:        tlsConfig,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     true,
		MaxIdleConns:           0,
		MaxIdleConnsPerHost:    0,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        0,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//server,_ := url.Parse("https://10.206.16.18:16443")
		server, _ := url.Parse("https://10.0.12.9:8443")
		log.Println(request.URL.Path)
		p := httputil.NewSingleHostReverseProxy(server)
		p.Transport = transport
		p.ServeHTTP(writer, request)

	})
	log.Println("开始反向代理k8sapi")
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
```



