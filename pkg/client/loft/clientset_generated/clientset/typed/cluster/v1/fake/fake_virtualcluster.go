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

// FakeVirtualClusters implements VirtualClusterInterface
type FakeVirtualClusters struct {
	Fake *FakeClusterV1
	ns   string
}

var virtualclustersResource = schema.GroupVersionResource{Group: "cluster.loft.sh", Version: "v1", Resource: "virtualclusters"}

var virtualclustersKind = schema.GroupVersionKind{Group: "cluster.loft.sh", Version: "v1", Kind: "VirtualCluster"}

// Get takes name of the virtualCluster, and returns the corresponding virtualCluster object, and an error if there is any.
func (c *FakeVirtualClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *clusterv1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualclustersResource, c.ns, name), &clusterv1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.VirtualCluster), err
}

// List takes label and field selectors, and returns the list of VirtualClusters that match those selectors.
func (c *FakeVirtualClusters) List(ctx context.Context, opts v1.ListOptions) (result *clusterv1.VirtualClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualclustersResource, virtualclustersKind, c.ns, opts), &clusterv1.VirtualClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clusterv1.VirtualClusterList{ListMeta: obj.(*clusterv1.VirtualClusterList).ListMeta}
	for _, item := range obj.(*clusterv1.VirtualClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualClusters.
func (c *FakeVirtualClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualclustersResource, c.ns, opts))

}

// Create takes the representation of a virtualCluster and creates it.  Returns the server's representation of the virtualCluster, and an error, if there is any.
func (c *FakeVirtualClusters) Create(ctx context.Context, virtualCluster *clusterv1.VirtualCluster, opts v1.CreateOptions) (result *clusterv1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualclustersResource, c.ns, virtualCluster), &clusterv1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.VirtualCluster), err
}

// Update takes the representation of a virtualCluster and updates it. Returns the server's representation of the virtualCluster, and an error, if there is any.
func (c *FakeVirtualClusters) Update(ctx context.Context, virtualCluster *clusterv1.VirtualCluster, opts v1.UpdateOptions) (result *clusterv1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualclustersResource, c.ns, virtualCluster), &clusterv1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.VirtualCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualClusters) UpdateStatus(ctx context.Context, virtualCluster *clusterv1.VirtualCluster, opts v1.UpdateOptions) (*clusterv1.VirtualCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualclustersResource, "status", c.ns, virtualCluster), &clusterv1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.VirtualCluster), err
}

// Delete takes name of the virtualCluster and deletes it. Returns an error if one occurs.
func (c *FakeVirtualClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualclustersResource, c.ns, name, opts), &clusterv1.VirtualCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &clusterv1.VirtualClusterList{})
	return err
}

// Patch applies the patch and returns the patched virtualCluster.
func (c *FakeVirtualClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clusterv1.VirtualCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualclustersResource, c.ns, name, pt, data, subresources...), &clusterv1.VirtualCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.VirtualCluster), err
}
