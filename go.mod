module github.com/loft-sh/agentapi/v2

go 1.13

require (
	github.com/go-openapi/spec v0.20.1
	github.com/loft-sh/apiserver v0.0.0-20210607160412-10c99558fdeb
	k8s.io/api v0.21.1
	k8s.io/apimachinery v0.21.1
	k8s.io/apiserver v0.21.1
	k8s.io/client-go v0.21.1
	k8s.io/kube-openapi v0.0.0-20210527164424-3c818078ee3d
	sigs.k8s.io/controller-runtime v0.9.0
)

replace (
	github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.3
	github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.3
	github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.5
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/kubernetes-incubator/reference-docs => github.com/kubernetes-sigs/reference-docs v0.0.0-20170929004150-fcf65347b256
	github.com/markbates/inflect => github.com/markbates/inflect v1.0.4
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
)
