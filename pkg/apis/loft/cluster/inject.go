package cluster

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Config will be injected during startup and then passed to the rest storages
var Config *rest.Config

// CachedClient will be injected during startup and then passed to the rest storages
var CachedClient client.Client

// UncachedClient will be injected during startup and then passed to the rest storages
var UncachedClient client.Client

// ManagementClient will be injected during startup and then passed to the rest storages
var ManagementClient client.Client
