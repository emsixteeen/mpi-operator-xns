package main

import (
	kubeinformersGen "github.com/emsixteeen/mpi-operator-xns/pkg/generated/kube"
	kubeinformers "k8s.io/client-go/informers"
	kubeclientset "k8s.io/client-go/kubernetes"

	mpijobinformersGen "github.com/emsixteeen/mpi-operator-xns/pkg/generated/mpijob"
	mpijobclientset "github.com/kubeflow/mpi-operator/pkg/client/clientset/versioned"
	mpijobinformers "github.com/kubeflow/mpi-operator/pkg/client/informers/externalversions"

	schedinformersGen "github.com/emsixteeen/mpi-operator-xns/pkg/generated/scheduling"
	schedclientset "sigs.k8s.io/scheduler-plugins/pkg/generated/clientset/versioned"
	schedinformers "sigs.k8s.io/scheduler-plugins/pkg/generated/informers/externalversions"

	volcanoinformersGen "github.com/emsixteeen/mpi-operator-xns/pkg/generated/volcano"
	volcanoclient "volcano.sh/apis/pkg/client/clientset/versioned"
	volcanoinformers "volcano.sh/apis/pkg/client/informers/externalversions"
)

func MultiKubeInformer(namespaces []string, kubeClient kubeclientset.Interface) kubeinformers.SharedInformerFactory {
	return kubeinformersGen.NewSharedInformerFactoryWithOptions(
		kubeClient,
		0,
		kubeinformersGen.WithNamespaces(namespaces...))
}

func MultiMpiJobInformer(namespaces []string, mpiJobClient mpijobclientset.Interface) mpijobinformers.SharedInformerFactory {
	return mpijobinformersGen.NewSharedInformerFactoryWithOptions(
		mpiJobClient,
		0,
		mpijobinformersGen.WithNamespaces(namespaces...))
}

func MultiVolcanoInformer(namespaces []string, volcanoClient volcanoclient.Interface) volcanoinformers.SharedInformerFactory {
	return volcanoinformersGen.NewSharedInformerFactoryWithOptions(
		volcanoClient,
		0,
		volcanoinformersGen.WithNamespaces(namespaces...))
}

func MultiSchedulerPluginsInformer(namespaces []string, schedClient schedclientset.Interface) schedinformers.SharedInformerFactory {
	return schedinformersGen.NewSharedInformerFactoryWithOptions(
		schedClient,
		0,
		schedinformersGen.WithNamespaces(namespaces...))
}
