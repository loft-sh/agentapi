// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLocalUsers implements LocalUserInterface
type FakeLocalUsers struct {
	Fake *FakeStorageV1
}

var localusersResource = schema.GroupVersionResource{Group: "storage.loft.sh", Version: "v1", Resource: "localusers"}

var localusersKind = schema.GroupVersionKind{Group: "storage.loft.sh", Version: "v1", Kind: "LocalUser"}

// Get takes name of the localUser, and returns the corresponding localUser object, and an error if there is any.
func (c *FakeLocalUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *storagev1.LocalUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(localusersResource, name), &storagev1.LocalUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.LocalUser), err
}

// List takes label and field selectors, and returns the list of LocalUsers that match those selectors.
func (c *FakeLocalUsers) List(ctx context.Context, opts v1.ListOptions) (result *storagev1.LocalUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(localusersResource, localusersKind, opts), &storagev1.LocalUserList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1.LocalUserList{ListMeta: obj.(*storagev1.LocalUserList).ListMeta}
	for _, item := range obj.(*storagev1.LocalUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested localUsers.
func (c *FakeLocalUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(localusersResource, opts))
}

// Create takes the representation of a localUser and creates it.  Returns the server's representation of the localUser, and an error, if there is any.
func (c *FakeLocalUsers) Create(ctx context.Context, localUser *storagev1.LocalUser, opts v1.CreateOptions) (result *storagev1.LocalUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(localusersResource, localUser), &storagev1.LocalUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.LocalUser), err
}

// Update takes the representation of a localUser and updates it. Returns the server's representation of the localUser, and an error, if there is any.
func (c *FakeLocalUsers) Update(ctx context.Context, localUser *storagev1.LocalUser, opts v1.UpdateOptions) (result *storagev1.LocalUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(localusersResource, localUser), &storagev1.LocalUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.LocalUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocalUsers) UpdateStatus(ctx context.Context, localUser *storagev1.LocalUser, opts v1.UpdateOptions) (*storagev1.LocalUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(localusersResource, "status", localUser), &storagev1.LocalUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.LocalUser), err
}

// Delete takes name of the localUser and deletes it. Returns an error if one occurs.
func (c *FakeLocalUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(localusersResource, name), &storagev1.LocalUser{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocalUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(localusersResource, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1.LocalUserList{})
	return err
}

// Patch applies the patch and returns the patched localUser.
func (c *FakeLocalUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1.LocalUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(localusersResource, name, pt, data, subresources...), &storagev1.LocalUser{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1.LocalUser), err
}
