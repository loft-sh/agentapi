// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/cluster/v1"
	clusterv1 "github.com/loft-sh/agentapi/v4/pkg/clientset/versioned/typed/cluster/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeFeatures implements FeatureInterface
type fakeFeatures struct {
	*gentype.FakeClientWithList[*v1.Feature, *v1.FeatureList]
	Fake *FakeClusterV1
}

func newFakeFeatures(fake *FakeClusterV1) clusterv1.FeatureInterface {
	return &fakeFeatures{
		gentype.NewFakeClientWithList[*v1.Feature, *v1.FeatureList](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("features"),
			v1.SchemeGroupVersion.WithKind("Feature"),
			func() *v1.Feature { return &v1.Feature{} },
			func() *v1.FeatureList { return &v1.FeatureList{} },
			func(dst, src *v1.FeatureList) { dst.ListMeta = src.ListMeta },
			func(list *v1.FeatureList) []*v1.Feature { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.FeatureList, items []*v1.Feature) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
