// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/cluster/v1"
	"github.com/loft-sh/agentapi/v2/pkg/client/loft/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ClusterV1Interface interface {
	RESTClient() rest.Interface
	ChartInfosGetter
	ClusterQuotasGetter
	HelmReleasesGetter
	LocalClusterAccessesGetter
	SleepModeConfigsGetter
	SpacesGetter
	VirtualClustersGetter
}

// ClusterV1Client is used to interact with features provided by the cluster.loft.sh group.
type ClusterV1Client struct {
	restClient rest.Interface
}

func (c *ClusterV1Client) ChartInfos() ChartInfoInterface {
	return newChartInfos(c)
}

func (c *ClusterV1Client) ClusterQuotas() ClusterQuotaInterface {
	return newClusterQuotas(c)
}

func (c *ClusterV1Client) HelmReleases(namespace string) HelmReleaseInterface {
	return newHelmReleases(c, namespace)
}

func (c *ClusterV1Client) LocalClusterAccesses() LocalClusterAccessInterface {
	return newLocalClusterAccesses(c)
}

func (c *ClusterV1Client) SleepModeConfigs(namespace string) SleepModeConfigInterface {
	return newSleepModeConfigs(c, namespace)
}

func (c *ClusterV1Client) Spaces() SpaceInterface {
	return newSpaces(c)
}

func (c *ClusterV1Client) VirtualClusters(namespace string) VirtualClusterInterface {
	return newVirtualClusters(c, namespace)
}

// NewForConfig creates a new ClusterV1Client for the given config.
func NewForConfig(c *rest.Config) (*ClusterV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterV1Client{client}, nil
}

// NewForConfigOrDie creates a new ClusterV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ClusterV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ClusterV1Client for the given RESTClient.
func New(c rest.Interface) *ClusterV1Client {
	return &ClusterV1Client{c}
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
func (c *ClusterV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
