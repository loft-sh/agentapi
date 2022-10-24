package v1

// VirtualClusterCommonSpec holds common attributes for virtual clusters and virtual cluster templates
type VirtualClusterCommonSpec struct {
	// Apps specifies the apps that should get deployed by this template
	// +optional
	Apps []AppReference `json:"apps,omitempty"`

	// Charts are helm charts that should get deployed
	// +optional
	Charts []TemplateHelmChart `json:"charts,omitempty"`

	// Objects are Kubernetes style yamls that should get deployed into the virtual cluster
	// +optional
	Objects string `json:"objects,omitempty"`

	// Access defines the access of users and teams to the virtual cluster.
	// +optional
	Access *InstanceAccess `json:"access,omitempty"`

	// HelmRelease is the helm release configuration for the virtual cluster.
	// +optional
	HelmRelease VirtualClusterHelmRelease `json:"helmRelease,omitempty"`

	// AccessPoint defines settings to expose the virtual cluster directly via an ingress rather than
	// through the (default) Loft proxy
	// +optional
	AccessPoint VirtualClusterAccessPoint `json:"accessPoint,omitempty"`
}

type VirtualClusterAccessPoint struct {
	// Ingress defines virtual cluster access via ingress
	// +optional
	Ingress VirtualClusterAccessPointIngressSpec `json:"ingress,omitempty"`
}

type VirtualClusterAccessPointIngressSpec struct {
	// Enabled defines if the virtual cluster access point (via ingress) is enabled or not; requires
	// the connected cluster to have the `loft.sh/ingress-suffix` annotation set to define the domain
	// name suffix used for the ingress.
	Enabled bool `json:"enabled,omitempty"`
}
