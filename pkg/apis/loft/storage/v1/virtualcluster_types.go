package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualCluster holds the information
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type VirtualCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterSpec   `json:"spec,omitempty"`
	Status VirtualClusterStatus `json:"status,omitempty"`
}

// GetConditions returns the set of conditions for this object.
func (in *VirtualCluster) GetConditions() Conditions {
	return in.Status.Conditions
}

// SetConditions sets the conditions on this object.
func (in *VirtualCluster) SetConditions(conditions Conditions) {
	in.Status.Conditions = conditions
}

type VirtualClusterSpec struct {
	// Access defines the access of users and teams to the virtual cluster.
	// +optional
	Access *VirtualClusterAccess `json:"access,omitempty"`

	// The helm release configuration for the virtual cluster. This is optional, but
	// when filled, loft will deploy the specified chart for the given
	// +optional
	HelmRelease *VirtualClusterHelmRelease `json:"helmRelease,omitempty"`

	// A label selector to select the virtual cluster pod to route
	// incoming requests to.
	// +optional
	Pod *PodSelector `json:"pod,omitempty"`

	// A reference to the cluster admin kube config. This is needed for
	// the cli & ui to access the virtual clusters
	// +optional
	KubeConfigRef *SecretRef `json:"kubeConfigRef,omitempty"`

	// Objects are Kubernetes style yamls that should get deployed into the virtual cluster
	// +optional
	Objects string `json:"objects,omitempty"`
}

type VirtualClusterAccess struct {
	// If enabled, service account tokens will not be allowed to access this virtual cluster
	// through the Loft gateway.
	// +optional
	DisableServiceAccountsAuth bool `json:"disableServiceAccountsAuth,omitempty"`

	// Specifies which cluster role should get applied to users or teams that do not
	// match a rule below.
	// +optional
	DefaultClusterRole string `json:"defaultClusterRole,omitempty"`

	// Rules defines which users and teams should have which access to the virtual
	// cluster. If no rule matches an authenticated incoming user, the user will get cluster admin
	// access.
	// +optional
	Rules []VirtualClusterAccessRule `json:"rules,omitempty"`
}

type VirtualClusterAccessRule struct {
	// Users this rule matches. * means all users.
	// +optional
	Users []string `json:"users,omitempty"`

	// Teams that this rule matches.
	// +optional
	Teams []string `json:"teams,omitempty"`

	// ClusterRole is the cluster role that should be assigned to the
	// +optional
	ClusterRole string `json:"clusterRole,omitempty"`
}

// SecretRef is the reference to a secret containing the user password
type SecretRef struct {
	// +optional
	SecretName string `json:"secretName,omitempty"`
	// +optional
	SecretNamespace string `json:"secretNamespace,omitempty"`
	// +optional
	Key string `json:"key,omitempty"`
}

type VirtualClusterHelmRelease struct {
	// infos about what chart to deploy
	// +optional
	Chart VirtualClusterHelmChart `json:"chart,omitempty"`

	// the values for the given chart
	// +optional
	Values string `json:"values,omitempty"`
}

type VirtualClusterHelmChart struct {
	// the name of the helm chart
	// +optional
	Name string `json:"name,omitempty"`

	// the repo of the helm chart
	// +optional
	Repo string `json:"repo,omitempty"`

	// the version of the helm chart to use
	// +optional
	Version string `json:"version,omitempty"`
}

type PodSelector struct {
	// A label selector to select the virtual cluster pod to route
	// incoming requests to.
	// +optional
	Selector metav1.LabelSelector `json:"podSelector,omitempty"`

	// The port of the pod to route to
	// +optional
	Port *int `json:"port,omitempty"`
}

// VirtualClusterStatus holds the status of a virtual cluster
type VirtualClusterStatus struct {
	// Phase describes the current phase the virtual cluster is in
	// +optional
	Phase VirtualClusterPhase `json:"phase,omitempty"`

	// Reason describes the reason in machine readable form why the cluster is in the current
	// phase
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message describes the reason in human readable form why the cluster is in the currrent
	// phase
	// +optional
	Message string `json:"message,omitempty"`

	// ControlPlaneReady defines if the virtual cluster control plane is ready.
	// +optional
	ControlPlaneReady bool `json:"controlPlaneReady,omitempty"`

	// Conditions holds several conditions the virtual cluster might be in
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`

	// ObservedGeneration is the latest generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// DEPRECATED: do not use anymore
	// the status of the helm release that was used to deploy the virtual cluster
	// +optional
	HelmRelease *VirtualClusterHelmReleaseStatus `json:"helmRelease,omitempty"`
}

type VirtualClusterHelmReleaseStatus struct {
	// +optional
	Phase string `json:"phase,omitempty"`

	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// +optional
	Reason string `json:"reason,omitempty"`

	// +optional
	Message string `json:"message,omitempty"`

	// the release that was deployed
	// +optional
	Release VirtualClusterHelmRelease `json:"release,omitempty"`
}

// VirtualClusterPhase describes the phase of a virtual cluster
type VirtualClusterPhase string

// These are the valid admin account types
const (
	VirtualClusterUnknown  VirtualClusterPhase = ""
	VirtualClusterPending  VirtualClusterPhase = "Pending"
	VirtualClusterDeployed VirtualClusterPhase = "Deployed"
	VirtualClusterFailed   VirtualClusterPhase = "Failed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualClusterList contains a list of User
type VirtualClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualCluster{}, &VirtualClusterList{})
}
