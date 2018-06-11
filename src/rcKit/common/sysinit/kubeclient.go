package sysinit

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"rcKit/common/utils"
	"rcKit/common/constant"
	"sync"
	mylog "github.com/maxwell92/gokits/log"
	"strings"
	"io/ioutil"
)

var log = mylog.Log
var once sync.Once
var KubeClient *kubernetes.Clientset

func NewKubeClient(auth int, caPath, certPath, keyPath string) error {
	var err error
	//非tls连接
	if auth == 0 {
		err = getKubeClinet()
	} else {
		err = getTLSKubeClinet(caPath, certPath, keyPath)
	}
	return err
}

func getKubeClinet() error {
	var err error
	config := new(rest.Config)
	config.Host = "http://" + utils.LoadEnvVar(constant.ENV_APISERVER, constant.APISERVER)
	KubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorf("NewKubeClient failed: error=%s, apiServer=%s\n", err, config.Host)
	}
	return err
}

func getTLSKubeClinet(caPath, certPath, keyPath string) error {
	var err error
	config := new(rest.Config)
	config.Host = "https://" + utils.LoadEnvVar(constant.ENV_APISERVER, constant.APISERVER)
	//tls连接
	if !strings.EqualFold(caPath, "") && !strings.EqualFold(certPath, "") && !strings.EqualFold(keyPath, "") {
		//从指定的路径下读文件
		config.CAData, err = ioutil.ReadFile(caPath)
		if err != nil {
			log.Errorf("NewKubeClient failed: Read ca File error. err=%v", err)
			return err
		}
		config.CertData, err = ioutil.ReadFile(certPath)
		if err != nil {
			log.Errorf("NewKubeClient failed: Read cert File error. err=%v", err)
			return err
		}
		config.KeyData, err = ioutil.ReadFile(keyPath)
		if err != nil {
			log.Errorf("NewKubeClient failed: Read key File error. err=%v", err)
			return err
		}
	} else {
		//如果不指定文件路径，默认从当前路径下读文件
		config.CAData, err = ioutil.ReadFile(constant.CA_FILE)
		if err != nil {
			log.Errorf("NewKubeClient failed: Read ca File error. err=%v", err)
			return err
		}
		config.CertData, err = ioutil.ReadFile(constant.CERT_FILE)
		if err != nil {
			log.Errorf("NewKubeClient failed: Read cert File error. err=%v", err)
			return err
		}
		config.KeyData, err = ioutil.ReadFile(constant.KEY_FILE)
		if err != nil {
			log.Errorf("NewKubeClient failed: Read key File error. err=%v", err)
			return err
		}
	}

	KubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorf("GetTLSKubeClient failed: error=%s, apiServer=%s\n", err, config.Host)
	}
	return err
}
