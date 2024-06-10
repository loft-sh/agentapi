//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by defaulter-gen. DO NOT EDIT.

package cluster

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&VirtualCluster{}, func(obj interface{}) { SetObjectDefaults_VirtualCluster(obj.(*VirtualCluster)) })
	scheme.AddTypeDefaultingFunc(&VirtualClusterList{}, func(obj interface{}) { SetObjectDefaults_VirtualClusterList(obj.(*VirtualClusterList)) })
	return nil
}

func SetObjectDefaults_VirtualCluster(in *VirtualCluster) {
	if in.Status.SyncerPod != nil {
		for i := range in.Status.SyncerPod.Spec.InitContainers {
			a := &in.Status.SyncerPod.Spec.InitContainers[i]
			for j := range a.Ports {
				b := &a.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.LivenessProbe != nil {
				if a.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.ReadinessProbe != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.StartupProbe != nil {
				if a.StartupProbe.ProbeHandler.GRPC != nil {
					if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
		for i := range in.Status.SyncerPod.Spec.Containers {
			a := &in.Status.SyncerPod.Spec.Containers[i]
			for j := range a.Ports {
				b := &a.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.LivenessProbe != nil {
				if a.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.ReadinessProbe != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.StartupProbe != nil {
				if a.StartupProbe.ProbeHandler.GRPC != nil {
					if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
		for i := range in.Status.SyncerPod.Spec.EphemeralContainers {
			a := &in.Status.SyncerPod.Spec.EphemeralContainers[i]
			for j := range a.EphemeralContainerCommon.Ports {
				b := &a.EphemeralContainerCommon.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.EphemeralContainerCommon.LivenessProbe != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.EphemeralContainerCommon.ReadinessProbe != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.EphemeralContainerCommon.StartupProbe != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
	}
	if in.Status.ClusterPod != nil {
		for i := range in.Status.ClusterPod.Spec.InitContainers {
			a := &in.Status.ClusterPod.Spec.InitContainers[i]
			for j := range a.Ports {
				b := &a.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.LivenessProbe != nil {
				if a.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.ReadinessProbe != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.StartupProbe != nil {
				if a.StartupProbe.ProbeHandler.GRPC != nil {
					if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
		for i := range in.Status.ClusterPod.Spec.Containers {
			a := &in.Status.ClusterPod.Spec.Containers[i]
			for j := range a.Ports {
				b := &a.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.LivenessProbe != nil {
				if a.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.ReadinessProbe != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.StartupProbe != nil {
				if a.StartupProbe.ProbeHandler.GRPC != nil {
					if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
		for i := range in.Status.ClusterPod.Spec.EphemeralContainers {
			a := &in.Status.ClusterPod.Spec.EphemeralContainers[i]
			for j := range a.EphemeralContainerCommon.Ports {
				b := &a.EphemeralContainerCommon.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.EphemeralContainerCommon.LivenessProbe != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.EphemeralContainerCommon.ReadinessProbe != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.EphemeralContainerCommon.StartupProbe != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
	}
}

func SetObjectDefaults_VirtualClusterList(in *VirtualClusterList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_VirtualCluster(a)
	}
}
