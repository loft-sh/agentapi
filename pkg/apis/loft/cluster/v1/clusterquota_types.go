package v1

import (
	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterQuota holds the virtual cluster information
// +k8s:openapi-gen=true
// +resource:path=clusterquotas,rest=ClusterQuotaREST
type ClusterQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterQuotaSpec   `json:"spec,omitempty"`
	Status ClusterQuotaStatus `json:"status,omitempty"`
}

type ClusterQuotaSpec struct {
	storagev1.ClusterQuotaSpec `json:",inline"`
}

type ClusterQuotaStatus struct {
	storagev1.ClusterQuotaStatus `json:",inline"`

	// Owner describes the owner of the space. This can be either empty (nil), be a team or
	// an loft user. If the space has an account that does not belong to an user / team in loft
	// this is empty
	// +optional
	Owner *UserOrTeam `json:"owner,omitempty"`
}
