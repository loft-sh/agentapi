// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/loft-sh/agentapi/pkg/apis/kiosk/tenancy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccounts implements AccountInterface
type FakeAccounts struct {
	Fake *FakeTenancyV1alpha1
}

var accountsResource = schema.GroupVersionResource{Group: "tenancy.kiosk.sh", Version: "v1alpha1", Resource: "accounts"}

var accountsKind = schema.GroupVersionKind{Group: "tenancy.kiosk.sh", Version: "v1alpha1", Kind: "Account"}

// Get takes name of the account, and returns the corresponding account object, and an error if there is any.
func (c *FakeAccounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(accountsResource, name), &v1alpha1.Account{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// List takes label and field selectors, and returns the list of Accounts that match those selectors.
func (c *FakeAccounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(accountsResource, accountsKind, opts), &v1alpha1.AccountList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountList{ListMeta: obj.(*v1alpha1.AccountList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accounts.
func (c *FakeAccounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(accountsResource, opts))
}

// Create takes the representation of a account and creates it.  Returns the server's representation of the account, and an error, if there is any.
func (c *FakeAccounts) Create(ctx context.Context, account *v1alpha1.Account, opts v1.CreateOptions) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(accountsResource, account), &v1alpha1.Account{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// Update takes the representation of a account and updates it. Returns the server's representation of the account, and an error, if there is any.
func (c *FakeAccounts) Update(ctx context.Context, account *v1alpha1.Account, opts v1.UpdateOptions) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(accountsResource, account), &v1alpha1.Account{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccounts) UpdateStatus(ctx context.Context, account *v1alpha1.Account, opts v1.UpdateOptions) (*v1alpha1.Account, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(accountsResource, "status", account), &v1alpha1.Account{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}

// Delete takes name of the account and deletes it. Returns an error if one occurs.
func (c *FakeAccounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(accountsResource, name), &v1alpha1.Account{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(accountsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountList{})
	return err
}

// Patch applies the patch and returns the patched account.
func (c *FakeAccounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Account, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(accountsResource, name, pt, data, subresources...), &v1alpha1.Account{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Account), err
}
