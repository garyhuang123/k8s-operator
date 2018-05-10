/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	reboot_v1alpha1 "k8s-operator/operator/reboot/pkg/apis/reboot/v1alpha1"
	versioned "k8s-operator/operator/reboot/pkg/client/clientset/versioned"
	internalinterfaces "k8s-operator/operator/reboot/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s-operator/operator/reboot/pkg/client/listers/reboot/v1alpha1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RebootInformer provides access to a shared informer and lister for
// Reboots.
type RebootInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RebootLister
}

type rebootInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRebootInformer constructs a new informer for Reboot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRebootInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRebootInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRebootInformer constructs a new informer for Reboot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRebootInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RebootV1alpha1().Reboots(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RebootV1alpha1().Reboots(namespace).Watch(options)
			},
		},
		&reboot_v1alpha1.Reboot{},
		resyncPeriod,
		indexers,
	)
}

func (f *rebootInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRebootInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *rebootInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&reboot_v1alpha1.Reboot{}, f.defaultInformer)
}

func (f *rebootInformer) Lister() v1alpha1.RebootLister {
	return v1alpha1.NewRebootLister(f.Informer().GetIndexer())
}
