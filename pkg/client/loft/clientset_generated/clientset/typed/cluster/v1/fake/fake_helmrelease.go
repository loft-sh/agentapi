// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	clusterv1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/cluster/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHelmReleases implements HelmReleaseInterface
type FakeHelmReleases struct {
	Fake *FakeClusterV1
	ns   string
}

var helmreleasesResource = schema.GroupVersionResource{Group: "cluster.loft.sh", Version: "v1", Resource: "helmreleases"}

var helmreleasesKind = schema.GroupVersionKind{Group: "cluster.loft.sh", Version: "v1", Kind: "HelmRelease"}

// Get takes name of the helmRelease, and returns the corresponding helmRelease object, and an error if there is any.
func (c *FakeHelmReleases) Get(ctx context.Context, name string, options v1.GetOptions) (result *clusterv1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(helmreleasesResource, c.ns, name), &clusterv1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.HelmRelease), err
}

// List takes label and field selectors, and returns the list of HelmReleases that match those selectors.
func (c *FakeHelmReleases) List(ctx context.Context, opts v1.ListOptions) (result *clusterv1.HelmReleaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(helmreleasesResource, helmreleasesKind, c.ns, opts), &clusterv1.HelmReleaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clusterv1.HelmReleaseList{ListMeta: obj.(*clusterv1.HelmReleaseList).ListMeta}
	for _, item := range obj.(*clusterv1.HelmReleaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helmReleases.
func (c *FakeHelmReleases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(helmreleasesResource, c.ns, opts))

}

// Create takes the representation of a helmRelease and creates it.  Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *FakeHelmReleases) Create(ctx context.Context, helmRelease *clusterv1.HelmRelease, opts v1.CreateOptions) (result *clusterv1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(helmreleasesResource, c.ns, helmRelease), &clusterv1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.HelmRelease), err
}

// Update takes the representation of a helmRelease and updates it. Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *FakeHelmReleases) Update(ctx context.Context, helmRelease *clusterv1.HelmRelease, opts v1.UpdateOptions) (result *clusterv1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(helmreleasesResource, c.ns, helmRelease), &clusterv1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.HelmRelease), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelmReleases) UpdateStatus(ctx context.Context, helmRelease *clusterv1.HelmRelease, opts v1.UpdateOptions) (*clusterv1.HelmRelease, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(helmreleasesResource, "status", c.ns, helmRelease), &clusterv1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.HelmRelease), err
}

// Delete takes name of the helmRelease and deletes it. Returns an error if one occurs.
func (c *FakeHelmReleases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(helmreleasesResource, c.ns, name), &clusterv1.HelmRelease{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelmReleases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(helmreleasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &clusterv1.HelmReleaseList{})
	return err
}

// Patch applies the patch and returns the patched helmRelease.
func (c *FakeHelmReleases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clusterv1.HelmRelease, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(helmreleasesResource, c.ns, name, pt, data, subresources...), &clusterv1.HelmRelease{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.HelmRelease), err
}
