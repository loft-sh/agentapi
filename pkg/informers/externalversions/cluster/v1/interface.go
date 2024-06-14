// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/loft-sh/agentapi/v4/pkg/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ChartInfos returns a ChartInfoInformer.
	ChartInfos() ChartInfoInformer
	// Features returns a FeatureInformer.
	Features() FeatureInformer
	// HelmReleases returns a HelmReleaseInformer.
	HelmReleases() HelmReleaseInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ChartInfos returns a ChartInfoInformer.
func (v *version) ChartInfos() ChartInfoInformer {
	return &chartInfoInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Features returns a FeatureInformer.
func (v *version) Features() FeatureInformer {
	return &featureInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// HelmReleases returns a HelmReleaseInformer.
func (v *version) HelmReleases() HelmReleaseInformer {
	return &helmReleaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
