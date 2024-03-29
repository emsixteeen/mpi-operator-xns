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

package v1

import (
	informers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/client-go/informers/core/v1"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespaces       informers.NamespaceSet
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespaces informers.NamespaceSet, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1.Interface {
	return &version{factory: f, namespaces: namespaces, tweakListOptions: tweakListOptions}
}

// ComponentStatuses returns a ComponentStatusInformer.
func (v *version) ComponentStatuses() v1.ComponentStatusInformer {
	return &componentStatusInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConfigMaps returns a ConfigMapInformer.
func (v *version) ConfigMaps() v1.ConfigMapInformer {
	return &configMapInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Endpoints returns a EndpointsInformer.
func (v *version) Endpoints() v1.EndpointsInformer {
	return &endpointsInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Events returns a EventInformer.
func (v *version) Events() v1.EventInformer {
	return &eventInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// LimitRanges returns a LimitRangeInformer.
func (v *version) LimitRanges() v1.LimitRangeInformer {
	return &limitRangeInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Namespaces returns a NamespaceInformer.
func (v *version) Namespaces() v1.NamespaceInformer {
	return &namespaceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Nodes returns a NodeInformer.
func (v *version) Nodes() v1.NodeInformer {
	return &nodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PersistentVolumes returns a PersistentVolumeInformer.
func (v *version) PersistentVolumes() v1.PersistentVolumeInformer {
	return &persistentVolumeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PersistentVolumeClaims returns a PersistentVolumeClaimInformer.
func (v *version) PersistentVolumeClaims() v1.PersistentVolumeClaimInformer {
	return &persistentVolumeClaimInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Pods returns a PodInformer.
func (v *version) Pods() v1.PodInformer {
	return &podInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// PodTemplates returns a PodTemplateInformer.
func (v *version) PodTemplates() v1.PodTemplateInformer {
	return &podTemplateInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ReplicationControllers returns a ReplicationControllerInformer.
func (v *version) ReplicationControllers() v1.ReplicationControllerInformer {
	return &replicationControllerInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ResourceQuotas returns a ResourceQuotaInformer.
func (v *version) ResourceQuotas() v1.ResourceQuotaInformer {
	return &resourceQuotaInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Secrets returns a SecretInformer.
func (v *version) Secrets() v1.SecretInformer {
	return &secretInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() v1.ServiceInformer {
	return &serviceInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ServiceAccounts returns a ServiceAccountInformer.
func (v *version) ServiceAccounts() v1.ServiceAccountInformer {
	return &serviceAccountInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}
