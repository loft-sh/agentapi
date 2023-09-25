// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/agentapi/v3/pkg/apis/loft/cluster/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLocalClusterAccesses implements LocalClusterAccessInterface
type FakeLocalClusterAccesses struct {
	Fake *FakeClusterV1
}

var localclusteraccessesResource = v1.SchemeGroupVersion.WithResource("localclusteraccesses")

var localclusteraccessesKind = v1.SchemeGroupVersion.WithKind("LocalClusterAccess")

// Get takes name of the localClusterAccess, and returns the corresponding localClusterAccess object, and an error if there is any.
func (c *FakeLocalClusterAccesses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LocalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localclusteraccessesResource, name), &v1.LocalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LocalClusterAccess), err
}

// List takes label and field selectors, and returns the list of LocalClusterAccesses that match those selectors.
func (c *FakeLocalClusterAccesses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LocalClusterAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localclusteraccessesResource, localclusteraccessesKind, opts), &v1.LocalClusterAccessList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.LocalClusterAccessList{ListMeta: obj.(*v1.LocalClusterAccessList).ListMeta}
	for _, item := range obj.(*v1.LocalClusterAccessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localClusterAccesses.
func (c *FakeLocalClusterAccesses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localclusteraccessesResource, opts))
}

// Create takes the representation of a localClusterAccess and creates it.  Returns the server's representation of the localClusterAccess, and an error, if there is any.
func (c *FakeLocalClusterAccesses) Create(ctx context.Context, localClusterAccess *v1.LocalClusterAccess, opts metav1.CreateOptions) (result *v1.LocalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localclusteraccessesResource, localClusterAccess), &v1.LocalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LocalClusterAccess), err
}

// Update takes the representation of a localClusterAccess and updates it. Returns the server's representation of the localClusterAccess, and an error, if there is any.
func (c *FakeLocalClusterAccesses) Update(ctx context.Context, localClusterAccess *v1.LocalClusterAccess, opts metav1.UpdateOptions) (result *v1.LocalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localclusteraccessesResource, localClusterAccess), &v1.LocalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LocalClusterAccess), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalClusterAccesses) UpdateStatus(ctx context.Context, localClusterAccess *v1.LocalClusterAccess, opts metav1.UpdateOptions) (*v1.LocalClusterAccess, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localclusteraccessesResource, "status", localClusterAccess), &v1.LocalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LocalClusterAccess), err
}

// Delete takes name of the localClusterAccess and deletes it. Returns an error if one occurs.
func (c *FakeLocalClusterAccesses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(localclusteraccessesResource, name, opts), &v1.LocalClusterAccess{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalClusterAccesses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localclusteraccessesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.LocalClusterAccessList{})
	return err
}

// Patch applies the patch and returns the patched localClusterAccess.
func (c *FakeLocalClusterAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocalClusterAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localclusteraccessesResource, name, pt, data, subresources...), &v1.LocalClusterAccess{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.LocalClusterAccess), err
}
