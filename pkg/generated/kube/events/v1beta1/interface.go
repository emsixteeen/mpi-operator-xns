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

package v1beta1

import (
	informers "github.com/maistra/xns-informer/pkg/informers"
	v1beta1 "k8s.io/client-go/informers/events/v1beta1"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespaces       informers.NamespaceSet
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespaces informers.NamespaceSet, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.Interface {
	return &version{factory: f, namespaces: namespaces, tweakListOptions: tweakListOptions}
}

// Events returns a EventInformer.
func (v *version) Events() v1beta1.EventInformer {
	return &eventInformer{factory: v.factory, namespaces: v.namespaces, tweakListOptions: v.tweakListOptions}
}
