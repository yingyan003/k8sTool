package k8s

import (
	"k8s.io/api/core/v1"
	"scaler/common/sysinit"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespace struct{}

func (ns *Namespace) Get(name string) *v1.Namespace {
	namespace, err := sysinit.KubeClient.CoreV1().Namespaces().Get(name, meta_v1.GetOptions{})
	if err != nil {
		log.Errorf("Namespace Get failed: error=%s, namespace=%s", err, name)
		return nil
	}
	log.Infof("Namespace Get success: namespace=%s", name)
	return namespace
}

func (ns *Namespace) List() *v1.NamespaceList {
	namespaces, err := sysinit.KubeClient.CoreV1().Namespaces().List(meta_v1.ListOptions{})
	if err != nil {
		log.Errorf("Namespace Get failed: error=%s", err)
		return nil
	}
	log.Infof("Namespace Get success")
	return namespaces
}
