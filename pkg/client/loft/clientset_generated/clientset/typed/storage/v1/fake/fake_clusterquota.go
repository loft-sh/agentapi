// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/loft-sh/agentapi/v3/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterQuotas implements ClusterQuotaInterface
type FakeClusterQuotas struct {
	Fake *FakeStorageV1
}

var clusterquotasResource = v1.SchemeGroupVersion.WithResource("clusterquotas")

var clusterquotasKind = v1.SchemeGroupVersion.WithKind("ClusterQuota")

// Get takes name of the clusterQuota, and returns the corresponding clusterQuota object, and an error if there is any.
func (c *FakeClusterQuotas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterquotasResource, name), &v1.ClusterQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterQuota), err
}

// List takes label and field selectors, and returns the list of ClusterQuotas that match those selectors.
func (c *FakeClusterQuotas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterquotasResource, clusterquotasKind, opts), &v1.ClusterQuotaList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterQuotaList{ListMeta: obj.(*v1.ClusterQuotaList).ListMeta}
	for _, item := range obj.(*v1.ClusterQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterQuotas.
func (c *FakeClusterQuotas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterquotasResource, opts))
}

// Create takes the representation of a clusterQuota and creates it.  Returns the server's representation of the clusterQuota, and an error, if there is any.
func (c *FakeClusterQuotas) Create(ctx context.Context, clusterQuota *v1.ClusterQuota, opts metav1.CreateOptions) (result *v1.ClusterQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterquotasResource, clusterQuota), &v1.ClusterQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterQuota), err
}

// Update takes the representation of a clusterQuota and updates it. Returns the server's representation of the clusterQuota, and an error, if there is any.
func (c *FakeClusterQuotas) Update(ctx context.Context, clusterQuota *v1.ClusterQuota, opts metav1.UpdateOptions) (result *v1.ClusterQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterquotasResource, clusterQuota), &v1.ClusterQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterQuotas) UpdateStatus(ctx context.Context, clusterQuota *v1.ClusterQuota, opts metav1.UpdateOptions) (*v1.ClusterQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterquotasResource, "status", clusterQuota), &v1.ClusterQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterQuota), err
}

// Delete takes name of the clusterQuota and deletes it. Returns an error if one occurs.
func (c *FakeClusterQuotas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterquotasResource, name, opts), &v1.ClusterQuota{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterQuotas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterquotasResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterQuotaList{})
	return err
}

// Patch applies the patch and returns the patched clusterQuota.
func (c *FakeClusterQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterquotasResource, name, pt, data, subresources...), &v1.ClusterQuota{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ClusterQuota), err
}
