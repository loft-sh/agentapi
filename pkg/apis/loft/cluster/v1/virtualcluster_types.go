package v1

import (
	storagev1 "github.com/loft-sh/agentapi/v2/pkg/apis/loft/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualCluster holds the virtual cluster information
// +k8s:openapi-gen=true
// +resource:path=virtualclusters,rest=VirtualClusterREST
type VirtualCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterSpec   `json:"spec,omitempty"`
	Status VirtualClusterStatus `json:"status,omitempty"`
}

type VirtualClusterSpec struct {
	storagev1.VirtualClusterSpec `json:",inline"`
}

type VirtualClusterStatus struct {
	storagev1.VirtualClusterStatus `json:",inline"`

	// SleepModeConfig is the sleep mode config of the space
	// +optional
	SleepModeConfig *SleepModeConfig `json:"sleepModeConfig,omitempty"`
}
