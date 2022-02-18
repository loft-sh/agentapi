package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LocalSpaceConstraint holds the space constraints
// +k8s:openapi-gen=true
type LocalSpaceConstraint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocalSpaceConstraintSpec   `json:"spec,omitempty"`
	Status LocalSpaceConstraintStatus `json:"status,omitempty"`
}

type LocalSpaceConstraintSpec struct {
	// DisplayName is the name that should be shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description is the description of this object in
	// human-readable text.
	// +optional
	Description string `json:"description,omitempty"`

	// SpaceTemplate holds the space configuration
	// +optional
	SpaceTemplate ConstraintSpaceTemplate `json:"spaceTemplate,omitempty"`

	// Sync specifies if spaces that were created through this space constraint
	// object should get synced with this object.
	// +optional
	Sync bool `json:"sync,omitempty"`
}

// ConstraintSpaceTemplate defines properties how many spaces can be owned by the account and how they should be created
type ConstraintSpaceTemplate struct {
	// The enforced metadata of the space to create. Currently, only annotations and labels are supported
	// +kubebuilder:pruning:PreserveUnknownFields
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// This defines the cluster role that will be used for the rolebinding when
	// creating a new space for the selected subjects
	// +optional
	ClusterRole *string `json:"clusterRole,omitempty"`

	// Objects are Kubernetes style yamls that should get deployed into the space
	// +optional
	Objects string `json:"objects,omitempty"`
}

// LocalSpaceConstraintStatus holds the status of a user access
type LocalSpaceConstraintStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LocalSpaceConstraintList contains a list of LocalSpaceConstraint objects
type LocalSpaceConstraintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalSpaceConstraint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LocalSpaceConstraint{}, &LocalSpaceConstraintList{})
}
