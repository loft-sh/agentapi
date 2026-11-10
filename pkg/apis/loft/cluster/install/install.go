package install

import (
	"github.com/loft-sh/agentapi/v5/pkg/apis/loft/cluster"
	clusterv1 "github.com/loft-sh/agentapi/v5/pkg/apis/loft/cluster/v1"
	"github.com/loft-sh/apiserver/pkg/builders"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

func init() {
	InstallOptions(builders.Scheme)
	InstallOptions(builders.ParameterScheme)
	utilruntime.Must(clusterv1.RegisterConversions(builders.ParameterScheme))
}

func InstallOptions(scheme *runtime.Scheme) {
	utilruntime.Must(clusterv1.InstallOptions(scheme))
	utilruntime.Must(addKnownOptionsTypes(scheme))
}

func addKnownOptionsTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		cluster.SchemeGroupVersion,
		&cluster.Metal3BMCConsoleOptions{},
	)
	return nil
}
