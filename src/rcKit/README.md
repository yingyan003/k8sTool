# 背景
公司机房迁移，k8s集群从1.4.9升级到1.9.7，需要用到一键扩缩容应用。需求是：

1. 写一个二进制工具，直接在机器上使用。
2. 可修改指定namespace下某个应用的副本数
3. 可修改所有namespace下所有应用的副本数
3. 可连接旧集群（不需要TLS证书认证）和新集群（需要TLS证书认证）


# 使用方法说明：

将scaler的二进制文件拷贝到目标机器。
配置环境变量：
APISERVER: k8s的apiserver,形如：ip:端口。 必须配置
LOGLEVEL: 日志级别，默认是“ERROR”。 可选
在scaler的二进制文件的同级目录存放以下3个证书文件
· ca.crt：对应k8s证书ca.crt（证书名可能不同）
· client.crt:  对应k8s证书apiserver-kubelet-client.crt
· client.key:  对应k8s证书apiserver-kubelet-client.key
使用方式：
1）脚本运行参数查看方式： ./scaler  -h
2）参数详解：
-auth int
是否访问需要tls认证的集群，非必填，0表示不认证，1表示认证，默认值为0

-caPath string
ca对应证书的全路径名，非必填，无默认值。不提供则从当前路径取

-certPath string
crt对应证书的全路径名，非必填，无默认值。不提供则从当前路径取

-keyPath string
key对应证书的全路径名，非必填，无默认值。不提供则从当前路径取

-ns int
用于测试集群连接情况，非必填，默认为0，为1则返回namespace列表

-orgName string
组织名，必填，无默认值。为“all"表示更新所有namespace下的所有应用（kube-system和kube-public除外）

-dpName string
应用名，非必填，无默认值。不提供时默认更新指定组织下的所有应用

-replicas int
应用副本数，非必填，默认值为1 (default 1)
3) 如果连接需要证书的集群，证书配置可在运行参数指定，也可在配置文件配置。
如果在运行参数指定了证书的路径，必须3个证书同时提供，否则将从scaler二进制文件所在的同级目录下读取指定文件名的证书。

测试集群的连接情况：
./scaler -auth=0/1 -ns=1
返回结果：nil表示连接失败，否则返回namespace列表表示连接成功。


# 常见用法：

1.测试集群连接

source env.sh

TLS集群：./scaler -auth=1 -ns=1 -caPath=/etc/kubernetes/pki/ca.crt -certPath=/etc/kubernetes/pki/apiserver-kubelet-client.crt -keyPath=/etc/kubernetes/pki/apiserver-kubelet-client.key

非TLS集群：./scaler  -ns=1



2. 调整指定组织下的指定应用

source env.sh

TLS集群：./scaler -auth=1 -orgName XXX -dpName XXX -replicas XXX -caPath=/etc/kubernetes/pki/ca.crt -certPath=/etc/kubernetes/pki/apiserver-kubelet-client.crt -keyPath=/etc/kubernetes/pki/apiserver-kubelet-client.key

非TLS集群：./scaler  -orgName XXX -dpName XXX -replicas XXX



3. 调整指定组织下的全部应用

source env.sh

TLS集群：./scaler -auth=1 -orgName XXX -replicas XXX  -caPath=/etc/kubernetes/pki/ca.crt -certPath=/etc/kubernetes/pki/apiserver-kubelet-client.crt -keyPath=/etc/kubernetes/pki/apiserver-kubelet-client.key

非TLS集群：./scaler  -orgName XXX -replicas XXX



4. 调制全部组织（除kube-system和kube-public外）下的全部应用

source env.sh

TLS集群：./scaler -auth=1 -orgName all -replicas XXX -caPath=/etc/kubernetes/pki/ca.crt -certPath=/etc/kubernetes/pki/apiserver-kubelet-client.crt -keyPath=/etc/kubernetes/pki/apiserver-kubelet-client.key

非TLS集群：./scaler -orgName all -replicas XXX



