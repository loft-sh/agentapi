// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	clusterv1 "github.com/loft-sh/agentapi/v3/pkg/apis/loft/cluster/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSpaces implements SpaceInterface
type FakeSpaces struct {
	Fake *FakeClusterV1
}

var spacesResource = schema.GroupVersionResource{Group: "cluster.loft.sh", Version: "v1", Resource: "spaces"}

var spacesKind = schema.GroupVersionKind{Group: "cluster.loft.sh", Version: "v1", Kind: "Space"}

// Get takes name of the space, and returns the corresponding space object, and an error if there is any.
func (c *FakeSpaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *clusterv1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(spacesResource, name), &clusterv1.Space{})
	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.Space), err
}

// List takes label and field selectors, and returns the list of Spaces that match those selectors.
func (c *FakeSpaces) List(ctx context.Context, opts v1.ListOptions) (result *clusterv1.SpaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(spacesResource, spacesKind, opts), &clusterv1.SpaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clusterv1.SpaceList{ListMeta: obj.(*clusterv1.SpaceList).ListMeta}
	for _, item := range obj.(*clusterv1.SpaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spaces.
func (c *FakeSpaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(spacesResource, opts))
}

// Create takes the representation of a space and creates it.  Returns the server's representation of the space, and an error, if there is any.
func (c *FakeSpaces) Create(ctx context.Context, space *clusterv1.Space, opts v1.CreateOptions) (result *clusterv1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(spacesResource, space), &clusterv1.Space{})
	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.Space), err
}

// Update takes the representation of a space and updates it. Returns the server's representation of the space, and an error, if there is any.
func (c *FakeSpaces) Update(ctx context.Context, space *clusterv1.Space, opts v1.UpdateOptions) (result *clusterv1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(spacesResource, space), &clusterv1.Space{})
	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.Space), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpaces) UpdateStatus(ctx context.Context, space *clusterv1.Space, opts v1.UpdateOptions) (*clusterv1.Space, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(spacesResource, "status", space), &clusterv1.Space{})
	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.Space), err
}

// Delete takes name of the space and deletes it. Returns an error if one occurs.
func (c *FakeSpaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(spacesResource, name, opts), &clusterv1.Space{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(spacesResource, listOpts)

	_, err := c.Fake.Invokes(action, &clusterv1.SpaceList{})
	return err
}

// Patch applies the patch and returns the patched space.
func (c *FakeSpaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *clusterv1.Space, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(spacesResource, name, pt, data, subresources...), &clusterv1.Space{})
	if obj == nil {
		return nil, err
	}
	return obj.(*clusterv1.Space), err
}
