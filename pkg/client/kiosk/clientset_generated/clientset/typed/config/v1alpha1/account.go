// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/loft-sh/agentapi/v2/pkg/apis/kiosk/config/v1alpha1"
	scheme "github.com/loft-sh/agentapi/v2/pkg/client/kiosk/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AccountsGetter has a method to return a AccountInterface.
// A group's client should implement this interface.
type AccountsGetter interface {
	Accounts() AccountInterface
}

// AccountInterface has methods to work with Account resources.
type AccountInterface interface {
	Create(ctx context.Context, account *v1alpha1.Account, opts v1.CreateOptions) (*v1alpha1.Account, error)
	Update(ctx context.Context, account *v1alpha1.Account, opts v1.UpdateOptions) (*v1alpha1.Account, error)
	UpdateStatus(ctx context.Context, account *v1alpha1.Account, opts v1.UpdateOptions) (*v1alpha1.Account, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Account, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AccountList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Account, err error)
	AccountExpansion
}

// accounts implements AccountInterface
type accounts struct {
	client rest.Interface
}

// newAccounts returns a Accounts
func newAccounts(c *ConfigV1alpha1Client) *accounts {
	return &accounts{
		client: c.RESTClient(),
	}
}

// Get takes name of the account, and returns the corresponding account object, and an error if there is any.
func (c *accounts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Account, err error) {
	result = &v1alpha1.Account{}
	err = c.client.Get().
		Resource("accounts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Accounts that match those selectors.
func (c *accounts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AccountList{}
	err = c.client.Get().
		Resource("accounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested accounts.
func (c *accounts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("accounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a account and creates it.  Returns the server's representation of the account, and an error, if there is any.
func (c *accounts) Create(ctx context.Context, account *v1alpha1.Account, opts v1.CreateOptions) (result *v1alpha1.Account, err error) {
	result = &v1alpha1.Account{}
	err = c.client.Post().
		Resource("accounts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(account).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a account and updates it. Returns the server's representation of the account, and an error, if there is any.
func (c *accounts) Update(ctx context.Context, account *v1alpha1.Account, opts v1.UpdateOptions) (result *v1alpha1.Account, err error) {
	result = &v1alpha1.Account{}
	err = c.client.Put().
		Resource("accounts").
		Name(account.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(account).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *accounts) UpdateStatus(ctx context.Context, account *v1alpha1.Account, opts v1.UpdateOptions) (result *v1alpha1.Account, err error) {
	result = &v1alpha1.Account{}
	err = c.client.Put().
		Resource("accounts").
		Name(account.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(account).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the account and deletes it. Returns an error if one occurs.
func (c *accounts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("accounts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *accounts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("accounts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched account.
func (c *accounts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Account, err error) {
	result = &v1alpha1.Account{}
	err = c.client.Patch(pt).
		Resource("accounts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
