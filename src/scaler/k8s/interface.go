package k8s

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/api/core/v1"
)

//校验Deployment是否实现了接口Deployment。若没实现，静态编译过不去
var _ IDeployment = &Deployment{}

type IDeployment interface {
	Get(namespace, deploymentName string) *v1beta1.Deployment
	Scale(namespace string, deployment *v1beta1.Deployment) *v1beta1.Deployment
	List(namespace string) *v1beta1.DeploymentList
}

var _ INamespace = &Namespace{}

type INamespace interface {
	Get(name string) *v1.Namespace
	List()*v1.NamespaceList
}
