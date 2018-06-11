package main

import (
	"flag"
	"scaler/controller"
	"scaler/common/sysinit"
	"scaler/k8s"
	mylog "github.com/maxwell92/gokits/log"
	"strings"
)

var log = mylog.Log

func main() {
	//为接口赋值
	dpc := interfaceImpl(&k8s.Deployment{}, &k8s.Namespace{})
	readArgsAndScale(dpc)
}

func init() {
	sysinit.InitLog()
}

func interfaceImpl(ideployment k8s.IDeployment, inamespace k8s.INamespace) *controller.DeploymentController {
	return &controller.DeploymentController{Deployment: ideployment, Namespace: inamespace}
}

func readArgsAndScale(dpc *controller.DeploymentController) {
	orgName := flag.String("orgName", "", "组织名，必填，无默认值。")
	dpName := flag.String("dpName", "", "应用名，非必填，无默认值。不提供时默认更新指定组织下的所有应用")
	replicas := flag.Int("replicas", 1, "应用副本数，非必填，默认值为1")
	auth := flag.Int("auth", 0, "是否访问需要tls认证的集群，非必填，0表示不认证，1表示认证，默认值为0")
	ca := flag.String("caPath", "", "ca对应证书的全路径名，非必填，无默认值。不提供则从当前路径取")
	cert := flag.String("certPath", "", "crt对应证书的全路径名，非必填，无默认值。不提供则从当前路径取")
	key := flag.String("keyPath", "", "key对应证书的全路径名，非必填，无默认值。不提供则从当前路径取")
	conn := flag.Int("ns", 0, "用于测试集群连接情况，非必填，默认为0，为1则返回namespace列表")
	flag.Parse()

	opt := new(controller.DeployOptions)
	opt.OrgName = *orgName
	opt.DpName = *dpName
	opt.Replicas = *replicas
	opt.Auth = *auth
	opt.Ca = *ca
	opt.Cert = *cert
	opt.Key = *key
	opt.Conn = *conn

	if !validate(opt) {
		return
	}

	err:=sysinit.NewKubeClient(opt.Auth, opt.Ca, opt.Cert, opt.Key)
	if err!=nil{
		log.Errorf("k8s connecton failed")
		return
	}

	if opt.Conn == 1 {
		list:=dpc.Namespace.List()
		nss:=make([]string,0)
		for _,l:=range list.Items{
			nss=append(nss,l.Name)
		}
		log.Infof("namespace count=%d, list=%v",len(list.Items),nss)
		return
	}
	dpc.Scale(opt)
}

func validate(opt *controller.DeployOptions) bool {
	//非连接测试
	if opt.Conn != 1 {
		if opt.Replicas > 20 {
			log.Errorf("replicas can't be greater than 20")
			return false
		}
		if strings.EqualFold(opt.OrgName, "") {
			log.Errorf("orgName must provided")
			return false
		}
	}
	if opt.Auth != 0 && opt.Auth != 1 {
		log.Errorf("auth is illegal")
		return false
	}
	return true
}
