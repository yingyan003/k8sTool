package k8s

import (
	"k8s.io/api/extensions/v1beta1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	mylog "github.com/maxwell92/gokits/log"
	"scaler/common/sysinit"
)

var log = mylog.Log

type Deployment struct{}

func (dp *Deployment) Get(namespace, deploymentName string) *v1beta1.Deployment {
	deployment, err := sysinit.KubeClient.ExtensionsV1beta1().Deployments(namespace).Get(deploymentName, meta_v1.GetOptions{})
	if err != nil {
		log.Errorf("k8s.Deployment Get failed: error=%s, namespace=%s, dpName=%s\n", err, namespace,deploymentName)
		return nil
	}
	log.Infof("k8s.Deployment Get Success: namespace=%s, dpName=%s", namespace,deploymentName)
	return deployment
}

func (dp *Deployment) Scale(namespace string, deployment *v1beta1.Deployment) *v1beta1.Deployment {
	deployment, err := sysinit.KubeClient.ExtensionsV1beta1().Deployments(namespace).Update(deployment)
	if err != nil {
		log.Errorf("k8s.Deployment Scale failed: error=%s, namespace=%s, dpName=%s", err, namespace, deployment.Name)
		return nil
	}
	log.Infof("k8s.Deployment Scale Success: namespace=%s, dpName=%s", namespace, deployment.Name)
	return deployment
}

func (dp *Deployment) List(namespace string) *v1beta1.DeploymentList {
	DeploymentList, err := sysinit.KubeClient.ExtensionsV1beta1().Deployments(namespace).List(meta_v1.ListOptions{})
	if err != nil {
		log.Errorf("k8s.Deployment List failed: error=%s, namespace=%s", err, namespace)
		return nil
	}
	log.Infof("k8s.Deployment List Success: namespace=%s", namespace)
	return DeploymentList
}

//当给定的namespace不存在
//List时：并不报错，返回的len(result.item)=0
//Get时：不报namespace不存在，报指定的资源名不存在
