// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	storagev1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/storage/v1"
	versioned "github.com/loft-sh/agentapi/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/agentapi/v4/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/loft-sh/agentapi/v4/pkg/listers/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LocalUserInformer provides access to a shared informer and lister for
// LocalUsers.
type LocalUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.LocalUserLister
}

type localUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLocalUserInformer constructs a new informer for LocalUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLocalUserInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLocalUserInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLocalUserInformer constructs a new informer for LocalUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLocalUserInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().LocalUsers().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StorageV1().LocalUsers().Watch(context.TODO(), options)
			},
		},
		&storagev1.LocalUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *localUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLocalUserInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *localUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&storagev1.LocalUser{}, f.defaultInformer)
}

func (f *localUserInformer) Lister() v1.LocalUserLister {
	return v1.NewLocalUserLister(f.Informer().GetIndexer())
}
