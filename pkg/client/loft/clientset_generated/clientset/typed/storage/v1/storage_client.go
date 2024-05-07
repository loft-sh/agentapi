// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/storage/v1"
	"github.com/loft-sh/agentapi/v4/pkg/client/loft/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type StorageV1Interface interface {
	RESTClient() rest.Interface
	ClusterQuotasGetter
	LocalClusterAccessesGetter
	LocalTeamsGetter
	LocalUsersGetter
	VirtualClustersGetter
}

// StorageV1Client is used to interact with features provided by the storage.loft.sh group.
type StorageV1Client struct {
	restClient rest.Interface
}

func (c *StorageV1Client) ClusterQuotas() ClusterQuotaInterface {
	return newClusterQuotas(c)
}

func (c *StorageV1Client) LocalClusterAccesses() LocalClusterAccessInterface {
	return newLocalClusterAccesses(c)
}

func (c *StorageV1Client) LocalTeams() LocalTeamInterface {
	return newLocalTeams(c)
}

func (c *StorageV1Client) LocalUsers() LocalUserInterface {
	return newLocalUsers(c)
}

func (c *StorageV1Client) VirtualClusters(namespace string) VirtualClusterInterface {
	return newVirtualClusters(c, namespace)
}

// NewForConfig creates a new StorageV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*StorageV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new StorageV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*StorageV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &StorageV1Client{client}, nil
}

// NewForConfigOrDie creates a new StorageV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *StorageV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new StorageV1Client for the given RESTClient.
func New(c rest.Interface) *StorageV1Client {
	return &StorageV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *StorageV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
