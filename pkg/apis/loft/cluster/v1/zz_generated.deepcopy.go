//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedMetadata) DeepCopyInto(out *AppliedMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedMetadata.
func (in *AppliedMetadata) DeepCopy() *AppliedMetadata {
	if in == nil {
		return nil
	}
	out := new(AppliedMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedObject) DeepCopyInto(out *AppliedObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedObject.
func (in *AppliedObject) DeepCopy() *AppliedObject {
	if in == nil {
		return nil
	}
	out := new(AppliedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bash) DeepCopyInto(out *Bash) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bash.
func (in *Bash) DeepCopy() *Bash {
	if in == nil {
		return nil
	}
	out := new(Bash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfo) DeepCopyInto(out *ChartInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfo.
func (in *ChartInfo) DeepCopy() *ChartInfo {
	if in == nil {
		return nil
	}
	out := new(ChartInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChartInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfoList) DeepCopyInto(out *ChartInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChartInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfoList.
func (in *ChartInfoList) DeepCopy() *ChartInfoList {
	if in == nil {
		return nil
	}
	out := new(ChartInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChartInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfoSpec) DeepCopyInto(out *ChartInfoSpec) {
	*out = *in
	out.Chart = in.Chart
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfoSpec.
func (in *ChartInfoSpec) DeepCopy() *ChartInfoSpec {
	if in == nil {
		return nil
	}
	out := new(ChartInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartInfoStatus) DeepCopyInto(out *ChartInfoStatus) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(Metadata)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartInfoStatus.
func (in *ChartInfoStatus) DeepCopy() *ChartInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ChartInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQuota) DeepCopyInto(out *ClusterQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuota.
func (in *ClusterQuota) DeepCopy() *ClusterQuota {
	if in == nil {
		return nil
	}
	out := new(ClusterQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQuotaList) DeepCopyInto(out *ClusterQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuotaList.
func (in *ClusterQuotaList) DeepCopy() *ClusterQuotaList {
	if in == nil {
		return nil
	}
	out := new(ClusterQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQuotaSpec) DeepCopyInto(out *ClusterQuotaSpec) {
	*out = *in
	in.ClusterQuotaSpec.DeepCopyInto(&out.ClusterQuotaSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuotaSpec.
func (in *ClusterQuotaSpec) DeepCopy() *ClusterQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterQuotaStatus) DeepCopyInto(out *ClusterQuotaStatus) {
	*out = *in
	in.ClusterQuotaStatus.DeepCopyInto(&out.ClusterQuotaStatus)
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(UserOrTeam)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuotaStatus.
func (in *ClusterQuotaStatus) DeepCopy() *ClusterQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EntityInfo) DeepCopyInto(out *EntityInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntityInfo.
func (in *EntityInfo) DeepCopy() *EntityInfo {
	if in == nil {
		return nil
	}
	out := new(EntityInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EpochInfo) DeepCopyInto(out *EpochInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EpochInfo.
func (in *EpochInfo) DeepCopy() *EpochInfo {
	if in == nil {
		return nil
	}
	out := new(EpochInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmRelease) DeepCopyInto(out *HelmRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmRelease.
func (in *HelmRelease) DeepCopy() *HelmRelease {
	if in == nil {
		return nil
	}
	out := new(HelmRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseApp) DeepCopyInto(out *HelmReleaseApp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseApp.
func (in *HelmReleaseApp) DeepCopy() *HelmReleaseApp {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseConfig) DeepCopyInto(out *HelmReleaseConfig) {
	*out = *in
	out.Chart = in.Chart
	if in.Bash != nil {
		in, out := &in.Bash, &out.Bash
		*out = new(Bash)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseConfig.
func (in *HelmReleaseConfig) DeepCopy() *HelmReleaseConfig {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseList) DeepCopyInto(out *HelmReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseList.
func (in *HelmReleaseList) DeepCopy() *HelmReleaseList {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseSpec) DeepCopyInto(out *HelmReleaseSpec) {
	*out = *in
	in.HelmReleaseConfig.DeepCopyInto(&out.HelmReleaseConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseSpec.
func (in *HelmReleaseSpec) DeepCopy() *HelmReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseStatus) DeepCopyInto(out *HelmReleaseStatus) {
	*out = *in
	if in.Info != nil {
		in, out := &in.Info, &out.Info
		*out = new(Info)
		(*in).DeepCopyInto(*out)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(Metadata)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseStatus.
func (in *HelmReleaseStatus) DeepCopy() *HelmReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Info) DeepCopyInto(out *Info) {
	*out = *in
	in.FirstDeployed.DeepCopyInto(&out.FirstDeployed)
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
	in.Deleted.DeepCopyInto(&out.Deleted)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Info.
func (in *Info) DeepCopy() *Info {
	if in == nil {
		return nil
	}
	out := new(Info)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastActivityInfo) DeepCopyInto(out *LastActivityInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastActivityInfo.
func (in *LastActivityInfo) DeepCopy() *LastActivityInfo {
	if in == nil {
		return nil
	}
	out := new(LastActivityInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterAccess) DeepCopyInto(out *LocalClusterAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterAccess.
func (in *LocalClusterAccess) DeepCopy() *LocalClusterAccess {
	if in == nil {
		return nil
	}
	out := new(LocalClusterAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalClusterAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterAccessList) DeepCopyInto(out *LocalClusterAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalClusterAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterAccessList.
func (in *LocalClusterAccessList) DeepCopy() *LocalClusterAccessList {
	if in == nil {
		return nil
	}
	out := new(LocalClusterAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalClusterAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterAccessSpec) DeepCopyInto(out *LocalClusterAccessSpec) {
	*out = *in
	in.LocalClusterAccessSpec.DeepCopyInto(&out.LocalClusterAccessSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterAccessSpec.
func (in *LocalClusterAccessSpec) DeepCopy() *LocalClusterAccessSpec {
	if in == nil {
		return nil
	}
	out := new(LocalClusterAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterAccessStatus) DeepCopyInto(out *LocalClusterAccessStatus) {
	*out = *in
	out.LocalClusterAccessStatus = in.LocalClusterAccessStatus
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]*UserOrTeam, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(UserOrTeam)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]*EntityInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EntityInfo)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterAccessStatus.
func (in *LocalClusterAccessStatus) DeepCopy() *LocalClusterAccessStatus {
	if in == nil {
		return nil
	}
	out := new(LocalClusterAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Maintainer) DeepCopyInto(out *Maintainer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Maintainer.
func (in *Maintainer) DeepCopy() *Maintainer {
	if in == nil {
		return nil
	}
	out := new(Maintainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]*Maintainer, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Maintainer)
				**out = **in
			}
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Urls != nil {
		in, out := &in.Urls, &out.Urls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfig) DeepCopyInto(out *SleepModeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfig.
func (in *SleepModeConfig) DeepCopy() *SleepModeConfig {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SleepModeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfigList) DeepCopyInto(out *SleepModeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SleepModeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfigList.
func (in *SleepModeConfigList) DeepCopy() *SleepModeConfigList {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SleepModeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfigSpec) DeepCopyInto(out *SleepModeConfigSpec) {
	*out = *in
	if in.ForceSleepDuration != nil {
		in, out := &in.ForceSleepDuration, &out.ForceSleepDuration
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfigSpec.
func (in *SleepModeConfigSpec) DeepCopy() *SleepModeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SleepModeConfigStatus) DeepCopyInto(out *SleepModeConfigStatus) {
	*out = *in
	if in.LastActivityInfo != nil {
		in, out := &in.LastActivityInfo, &out.LastActivityInfo
		*out = new(LastActivityInfo)
		**out = **in
	}
	if in.CurrentEpoch != nil {
		in, out := &in.CurrentEpoch, &out.CurrentEpoch
		*out = new(EpochInfo)
		**out = **in
	}
	if in.LastEpoch != nil {
		in, out := &in.LastEpoch, &out.LastEpoch
		*out = new(EpochInfo)
		**out = **in
	}
	if in.SleptLastThirtyDays != nil {
		in, out := &in.SleptLastThirtyDays, &out.SleptLastThirtyDays
		*out = new(float64)
		**out = **in
	}
	if in.SleptLastSevenDays != nil {
		in, out := &in.SleptLastSevenDays, &out.SleptLastSevenDays
		*out = new(float64)
		**out = **in
	}
	if in.ScheduledSleep != nil {
		in, out := &in.ScheduledSleep, &out.ScheduledSleep
		*out = new(int64)
		**out = **in
	}
	if in.ScheduledWakeup != nil {
		in, out := &in.ScheduledWakeup, &out.ScheduledWakeup
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SleepModeConfigStatus.
func (in *SleepModeConfigStatus) DeepCopy() *SleepModeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(SleepModeConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Space) DeepCopyInto(out *Space) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Space.
func (in *Space) DeepCopy() *Space {
	if in == nil {
		return nil
	}
	out := new(Space)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Space) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpaceConstraintNamespaceStatus) DeepCopyInto(out *SpaceConstraintNamespaceStatus) {
	*out = *in
	if in.AppliedClusterRole != nil {
		in, out := &in.AppliedClusterRole, &out.AppliedClusterRole
		*out = new(string)
		**out = **in
	}
	in.AppliedMetadata.DeepCopyInto(&out.AppliedMetadata)
	if in.AppliedObjects != nil {
		in, out := &in.AppliedObjects, &out.AppliedObjects
		*out = make([]AppliedObject, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpaceConstraintNamespaceStatus.
func (in *SpaceConstraintNamespaceStatus) DeepCopy() *SpaceConstraintNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(SpaceConstraintNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpaceList) DeepCopyInto(out *SpaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Space, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpaceList.
func (in *SpaceList) DeepCopy() *SpaceList {
	if in == nil {
		return nil
	}
	out := new(SpaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpaceObjectsNamespaceStatus) DeepCopyInto(out *SpaceObjectsNamespaceStatus) {
	*out = *in
	if in.AppliedObjects != nil {
		in, out := &in.AppliedObjects, &out.AppliedObjects
		*out = make([]AppliedObject, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpaceObjectsNamespaceStatus.
func (in *SpaceObjectsNamespaceStatus) DeepCopy() *SpaceObjectsNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(SpaceObjectsNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpaceSpec) DeepCopyInto(out *SpaceSpec) {
	*out = *in
	if in.Finalizers != nil {
		in, out := &in.Finalizers, &out.Finalizers
		*out = make([]corev1.FinalizerName, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpaceSpec.
func (in *SpaceSpec) DeepCopy() *SpaceSpec {
	if in == nil {
		return nil
	}
	out := new(SpaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpaceStatus) DeepCopyInto(out *SpaceStatus) {
	*out = *in
	if in.SleepModeConfig != nil {
		in, out := &in.SleepModeConfig, &out.SleepModeConfig
		*out = new(SleepModeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(UserOrTeam)
		(*in).DeepCopyInto(*out)
	}
	if in.SpaceObjectsStatus != nil {
		in, out := &in.SpaceObjectsStatus, &out.SpaceObjectsStatus
		*out = new(SpaceObjectsNamespaceStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.TemplateSyncStatus != nil {
		in, out := &in.TemplateSyncStatus, &out.TemplateSyncStatus
		*out = new(TemplateSyncStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpaceStatus.
func (in *SpaceStatus) DeepCopy() *SpaceStatus {
	if in == nil {
		return nil
	}
	out := new(SpaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSyncStatus) DeepCopyInto(out *TemplateSyncStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSyncStatus.
func (in *TemplateSyncStatus) DeepCopy() *TemplateSyncStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateSyncStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserOrTeam) DeepCopyInto(out *UserOrTeam) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(EntityInfo)
		**out = **in
	}
	if in.Team != nil {
		in, out := &in.Team, &out.Team
		*out = new(EntityInfo)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserOrTeam.
func (in *UserOrTeam) DeepCopy() *UserOrTeam {
	if in == nil {
		return nil
	}
	out := new(UserOrTeam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualCluster) DeepCopyInto(out *VirtualCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualCluster.
func (in *VirtualCluster) DeepCopy() *VirtualCluster {
	if in == nil {
		return nil
	}
	out := new(VirtualCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterList) DeepCopyInto(out *VirtualClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterList.
func (in *VirtualClusterList) DeepCopy() *VirtualClusterList {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterSpec) DeepCopyInto(out *VirtualClusterSpec) {
	*out = *in
	in.VirtualClusterSpec.DeepCopyInto(&out.VirtualClusterSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterSpec.
func (in *VirtualClusterSpec) DeepCopy() *VirtualClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterStatus) DeepCopyInto(out *VirtualClusterStatus) {
	*out = *in
	in.VirtualClusterStatus.DeepCopyInto(&out.VirtualClusterStatus)
	if in.SyncerPod != nil {
		in, out := &in.SyncerPod, &out.SyncerPod
		*out = new(corev1.Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterPod != nil {
		in, out := &in.ClusterPod, &out.ClusterPod
		*out = new(corev1.Pod)
		(*in).DeepCopyInto(*out)
	}
	if in.SleepModeConfig != nil {
		in, out := &in.SleepModeConfig, &out.SleepModeConfig
		*out = new(SleepModeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TemplateSyncStatus != nil {
		in, out := &in.TemplateSyncStatus, &out.TemplateSyncStatus
		*out = new(TemplateSyncStatus)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterStatus.
func (in *VirtualClusterStatus) DeepCopy() *VirtualClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterStatus)
	in.DeepCopyInto(out)
	return out
}
