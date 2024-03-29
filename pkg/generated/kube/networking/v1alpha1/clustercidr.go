/*
Copyright Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by xns-informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	networkingv1alpha1 "k8s.io/api/networking/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/networking/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterCIDRInformer provides access to a shared informer and lister for
// ClusterCIDRs.
type ClusterCIDRInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterCIDRLister
}

type clusterCIDRInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterCIDRInformer constructs a new informer for ClusterCIDR type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterCIDRInformer(client kubernetes.Interface,
	resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterCIDRInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterCIDRInformer constructs a new informer for ClusterCIDR type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterCIDRInformer(client kubernetes.Interface,
	resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().ClusterCIDRs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha1().ClusterCIDRs().Watch(context.TODO(), options)
			},
		},
		&networkingv1alpha1.ClusterCIDR{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterCIDRInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterCIDRInformer(client, resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterCIDRInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha1.ClusterCIDR{}, f.defaultInformer)
}

func (f *clusterCIDRInformer) Lister() v1alpha1.ClusterCIDRLister {
	return v1alpha1.NewClusterCIDRLister(f.Informer().GetIndexer())
}
