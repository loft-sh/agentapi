// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	clusterv1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/cluster/v1"
	versioned "github.com/loft-sh/agentapi/v4/pkg/clientset/versioned"
	internalinterfaces "github.com/loft-sh/agentapi/v4/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/loft-sh/agentapi/v4/pkg/listers/cluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SleepModeConfigInformer provides access to a shared informer and lister for
// SleepModeConfigs.
type SleepModeConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SleepModeConfigLister
}

type sleepModeConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSleepModeConfigInformer constructs a new informer for SleepModeConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSleepModeConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSleepModeConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSleepModeConfigInformer constructs a new informer for SleepModeConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSleepModeConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1().SleepModeConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1().SleepModeConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&clusterv1.SleepModeConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *sleepModeConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSleepModeConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sleepModeConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterv1.SleepModeConfig{}, f.defaultInformer)
}

func (f *sleepModeConfigInformer) Lister() v1.SleepModeConfigLister {
	return v1.NewSleepModeConfigLister(f.Informer().GetIndexer())
}
