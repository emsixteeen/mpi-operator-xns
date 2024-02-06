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

package v1alpha2

import (
	informers "github.com/maistra/xns-informer/pkg/informers"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	v1alpha2 "k8s.io/client-go/informers/resource/v1alpha2"
)

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespaces       informers.NamespaceSet
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespaces informers.NamespaceSet, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1alpha2.Interface {
	return &version{factory: f, namespaces: namespaces, tweakListOptions: tweakListOptions}
}

// PodSchedulingContexts returns a PodSchedulingContextInformer.
func (v *version) PodSchedulingContexts() v1alpha2.PodSchedulingContextInformer {
	return &podSchedulingContextInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ResourceClaims returns a ResourceClaimInformer.
func (v *version) ResourceClaims() v1alpha2.ResourceClaimInformer {
	return &resourceClaimInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ResourceClaimTemplates returns a ResourceClaimTemplateInformer.
func (v *version) ResourceClaimTemplates() v1alpha2.ResourceClaimTemplateInformer {
	return &resourceClaimTemplateInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}

// ResourceClasses returns a ResourceClassInformer.
func (v *version) ResourceClasses() v1alpha2.ResourceClassInformer {
	return &resourceClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
