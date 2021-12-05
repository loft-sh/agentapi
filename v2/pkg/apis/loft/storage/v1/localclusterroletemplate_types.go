package v1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LocalClusterRoleTemplate holds a cluster role configuration
// +k8s:openapi-gen=true
type LocalClusterRoleTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LocalClusterRoleTemplateSpec   `json:"spec,omitempty"`
	Status LocalClusterRoleTemplateStatus `json:"status,omitempty"`
}

type LocalClusterRoleTemplateSpec struct {
	// DisplayName is the name that should be shown in the UI
	// +optional
	DisplayName string `json:"displayName,omitempty"`

	// Description is the description of this object in
	// human-readable text.
	// +optional
	Description string `json:"description,omitempty"`

	// ClusterRoleTemplate holds the cluster role template
	// +optional
	ClusterRoleTemplate ClusterRoleTemplateTemplate `json:"clusterRoleTemplate,omitempty"`
}

type ClusterRoleTemplateTemplate struct {
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Rules holds all the PolicyRules for this ClusterRole
	// +optional
	Rules []rbacv1.PolicyRule `json:"rules" protobuf:"bytes,2,rep,name=rules"`

	// AggregationRule is an optional field that describes how to build the Rules for this ClusterRole.
	// If AggregationRule is set, then the Rules are controller managed and direct changes to Rules will be
	// stomped by the controller.
	// +optional
	AggregationRule *rbacv1.AggregationRule `json:"aggregationRule,omitempty" protobuf:"bytes,3,opt,name=aggregationRule"`
}

// LocalClusterRoleTemplateStatus holds the status of a user access
type LocalClusterRoleTemplateStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LocalClusterRoleTemplateList contains a list of LocalClusterRoleTemplate objects
type LocalClusterRoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalClusterRoleTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LocalClusterRoleTemplate{}, &LocalClusterRoleTemplateList{})
}
