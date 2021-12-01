// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/cluster/v1"
	scheme "github.com/loft-sh/agentapi/v2/pkg/client/loft/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualClustersGetter has a method to return a VirtualClusterInterface.
// A group's client should implement this interface.
type VirtualClustersGetter interface {
	VirtualClusters(namespace string) VirtualClusterInterface
}

// VirtualClusterInterface has methods to work with VirtualCluster resources.
type VirtualClusterInterface interface {
	Create(ctx context.Context, virtualCluster *v1.VirtualCluster, opts metav1.CreateOptions) (*v1.VirtualCluster, error)
	Update(ctx context.Context, virtualCluster *v1.VirtualCluster, opts metav1.UpdateOptions) (*v1.VirtualCluster, error)
	UpdateStatus(ctx context.Context, virtualCluster *v1.VirtualCluster, opts metav1.UpdateOptions) (*v1.VirtualCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualCluster, err error)
	VirtualClusterExpansion
}

// virtualClusters implements VirtualClusterInterface
type virtualClusters struct {
	client rest.Interface
	ns     string
}

// newVirtualClusters returns a VirtualClusters
func newVirtualClusters(c *ClusterV1Client, namespace string) *virtualClusters {
	return &virtualClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualCluster, and returns the corresponding virtualCluster object, and an error if there is any.
func (c *virtualClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualCluster, err error) {
	result = &v1.VirtualCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualClusters that match those selectors.
func (c *virtualClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualClusters.
func (c *virtualClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualCluster and creates it.  Returns the server's representation of the virtualCluster, and an error, if there is any.
func (c *virtualClusters) Create(ctx context.Context, virtualCluster *v1.VirtualCluster, opts metav1.CreateOptions) (result *v1.VirtualCluster, err error) {
	result = &v1.VirtualCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualCluster and updates it. Returns the server's representation of the virtualCluster, and an error, if there is any.
func (c *virtualClusters) Update(ctx context.Context, virtualCluster *v1.VirtualCluster, opts metav1.UpdateOptions) (result *v1.VirtualCluster, err error) {
	result = &v1.VirtualCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(virtualCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualClusters) UpdateStatus(ctx context.Context, virtualCluster *v1.VirtualCluster, opts metav1.UpdateOptions) (result *v1.VirtualCluster, err error) {
	result = &v1.VirtualCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(virtualCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualCluster and deletes it. Returns an error if one occurs.
func (c *virtualClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualCluster.
func (c *virtualClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualCluster, err error) {
	result = &v1.VirtualCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
