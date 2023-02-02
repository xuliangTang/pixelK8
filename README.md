# PixelK8

kubernetes management platform  API. Front-end project is [pixelK8-element](https://github.com/xuliangTang/pixelK8-element)



## **Configuration**

example application.yml

```
port: 8081

logging:
  requestLogEnable: true

errorCache:
  enable: true

cors:
  enable: true

k8s:
  host: 
  port: 
  apiServer: "https://192.168.0.111:6443"
  defaultNs: default
  kubeConfigPath: "kubeconfig"
  nodes:
    - node1:
        username: 
        password: 
        host: 
        port: 22
    - node2:
        username: 
        password: 
        host: 
        port: 22
  caCrtPath: "k8s-ca.crt"
  cAKeyPath: "k8s-ca.key"
  userAccountPath: "./ua"

i18n:
  enable: true
  defaultLanguage: "zh"
```