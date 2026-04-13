// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:openapi-model-package=com.github.loft-sh.agentapi.v4.pkg.apis.loft.cluster.v1
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/skevetter/agentapi/pkg/apis/devsy/cluster
// +k8s:defaulter-gen=TypeMeta
// +groupName=cluster.devsy.sh
package v1 // import "github.com/skevetter/agentapi/pkg/apis/devsy/cluster/v1"
