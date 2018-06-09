package controller

import (
	"scaler/k8s"
	"strings"
	mylog "github.com/maxwell92/gokits/log"
	"k8s.io/api/extensions/v1beta1"
)

var log = mylog.Log

type DeployOptions struct {
	DpName  string
	OrgName string
	Replicas int
	Auth    int
	Ca string
	Cert string
	Key string
	Conn int
}

type DeploymentController struct {
	Deployment k8s.IDeployment
	Namespace  k8s.INamespace
}

type FailedNsAndDp struct {
	Namespace string
	dpNames   []string
}

func (dp *DeploymentController) Scale(opt *DeployOptions) {
	replicaInt32 := int32(opt.Replicas)
	//更新单个应用
	if !strings.EqualFold(opt.DpName, "") {
		deploy := dp.Deployment.Get(opt.OrgName, opt.DpName)
		if !dp.validate(opt, deploy) {
			return
		}
		deploy.Spec.Replicas = &replicaInt32
		result := dp.Deployment.Scale(opt.OrgName, deploy)
		if result == nil {
			log.Errorf("Scale: single app scale failed. orgName=%s, dpName=%s, preReplicas=%d, newReplicas=%d", opt.OrgName, opt.DpName)
			return
		}
		log.Infof("Scale: single app scale success. orgName=%s, dpName=%s", opt.OrgName, opt.DpName)
		return
	}

	if !dp.validate(opt, nil) {
		return
	}

	//更新所有组织，但忽略kube-public和kube-system
	if strings.EqualFold(opt.OrgName, "all") {
		faileds := make([]FailedNsAndDp, 0)
		nss := dp.Namespace.List()
		if nss == nil {
			log.Errorf("Scale: apps of all org scale failed. list namespace failed")
			return
		}
		log.Infof("Scale: apps of all namespace. count=%d", len(nss.Items))
		for _, ns := range nss.Items {
			if !strings.EqualFold(ns.Name, "kube-system") && !strings.EqualFold(ns.Name, "kube-public") {
				ok, data := dp.scaleAppOfOneOrg(ns.Name, replicaInt32)
				if !ok {
					failed := new(FailedNsAndDp)
					failed.Namespace = ns.Name
					failed.dpNames = data
					faileds = append(faileds, *failed)
				}
			}
		}
		log.Infof("Scale all org result: updata:%d, success:%d, failed:%d",len(nss.Items),len(nss.Items)-len(faileds),len(faileds))

		if len(faileds)!=0{
			log.Infof("failed namespace and app:")
			for _, failed := range faileds {
				log.Infof("namespace: ", failed.Namespace)
				log.Infof("appName: ", failed.dpNames)
			}
		}
	} else {
		//更新指定组织下的所有应用
		dp.scaleAppOfOneOrg(opt.OrgName, replicaInt32)
	}
}

func (dp *DeploymentController) validate(opt *DeployOptions, deploy *v1beta1.Deployment) bool {
	if strings.EqualFold(opt.OrgName, "all") {
		return true
	}
	if dp.Namespace.Get(opt.OrgName) == nil {
		log.Errorf("orgName isn't existed：orgName=%s", opt.OrgName)
		return false
	}
	//更新指定应用时，校验应用是否存在
	if !strings.EqualFold(opt.DpName, "") && deploy == nil {
		log.Errorf("app isn't existed：dpName=%s", opt.DpName)
		return false
	}
	return true
}

func (dp *DeploymentController) scaleAppOfOneOrg(orgName string, replicas int32) (bool, []string) {
	i := 0
	failedApps := make([]string, 0)
	dpList := dp.Deployment.List(orgName)
	if dpList == nil {
		log.Errorf("ScaleAppOfOneOrg: list dp failed. orgName=%s", orgName)
		return false, nil
	}
	for _, deploy := range dpList.Items {
		deploy.Spec.Replicas = &replicas
		result := dp.Deployment.Scale(orgName, &deploy)
		if result == nil {
			i++
			log.Errorf("ScaleAppOfOneOrg: scale app failed. orgName=%s, dpName=%s", orgName, deploy.Name)
		}
		failedApps = append(failedApps, deploy.Name)
	}

	log.Infof("ScaleAppOfOneOrg resut: update: %d, succes: %d, failed: %d", len(dpList.Items), len(dpList.Items)-i, i)

	if i == 0 {
		log.Infof("ScaleAppOfOneOrg: scale app success. orgName=%s", orgName)
		return true, nil
	}

	return false, failedApps
}
