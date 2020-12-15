package utils

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func getKubeClient(opt TestOptions, isHub bool) kubernetes.Interface {
	clientKube := NewKubeClient(
		opt.HubCluster.MasterURL,
		opt.KubeConfig,
		opt.HubCluster.KubeContext)
	if !isHub && len(opt.ManagedClusters) > 0 {
		clientKube = NewKubeClient(
			opt.ManagedClusters[0].MasterURL,
			opt.ManagedClusters[0].KubeConfig,
			"")
		// use the default context as workaround
		//			opt.ManagedClusters[0].KubeContext)
	}
	return clientKube
}

func getKubeClientDynamic(opt TestOptions, isHub bool) dynamic.Interface {
	url := opt.HubCluster.MasterURL
	kubeConfig := opt.KubeConfig
	kubeContext := opt.HubCluster.KubeContext
	if !isHub && len(opt.ManagedClusters) > 0 {
		url = opt.ManagedClusters[0].MasterURL
		kubeConfig = opt.ManagedClusters[0].KubeConfig
		kubeContext = ""
	}

	config, err := LoadConfig(url, kubeConfig, kubeContext)
	if err != nil {
		panic(err)
	}

	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}
