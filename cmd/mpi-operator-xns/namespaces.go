package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeclientset "k8s.io/client-go/kubernetes"
	"regexp"
)

func NamespaceParser(namespace string, kubeClient kubeclientset.Interface) ([]string, error) {
	nsList, err := kubeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing namespaces: %s", err.Error())
	}

	r, err := regexp.Compile(namespace)
	if err != nil {
		return nil, fmt.Errorf("error compiling namespace regex %s: %s", namespace, err.Error())
	}

	var namespaces []string
	for _, ns := range nsList.Items {
		if r.MatchString(ns.Name) {
			namespaces = append(namespaces, ns.Name)
		}
	}

	if len(namespaces) == 0 {
		return nil, fmt.Errorf("no matching namespaces found for regex %s", namespace)
	}

	return namespaces, nil
}
