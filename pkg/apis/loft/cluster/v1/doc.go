// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/loft-sh/agentapi/v2/pkg/apis/loft/cluster
// +k8s:defaulter-gen=TypeMeta
// +groupName=cluster.loft.sh
package v1 // import "github.com/loft-sh/agentapi/v2/pkg/apis/loft/cluster/v1"
