// This file was automatically generated by informer-gen

package v1

import (
	route_v1 "github.com/openshift/api/route/v1"
	versioned "github.com/openshift/client-go/route/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/route/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/route/listers/route/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// RouteInformer provides access to a shared informer and lister for
// Routes.
type RouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RouteLister
}

type routeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRouteInformer constructs a new informer for Route type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRouteInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRouteInformer constructs a new informer for Route type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RouteV1().Routes(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RouteV1().Routes(namespace).Watch(options)
			},
		},
		&route_v1.Route{},
		resyncPeriod,
		indexers,
	)
}

func (f *routeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRouteInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *routeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&route_v1.Route{}, f.defaultInformer)
}

func (f *routeInformer) Lister() v1.RouteLister {
	return v1.NewRouteLister(f.Informer().GetIndexer())
}
