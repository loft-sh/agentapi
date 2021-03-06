//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessQuota) DeepCopyInto(out *AccessQuota) {
	*out = *in
	if in.Hard != nil {
		in, out := &in.Hard, &out.Hard
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessQuota.
func (in *AccessQuota) DeepCopy() *AccessQuota {
	if in == nil {
		return nil
	}
	out := new(AccessQuota)
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
	in.Quota.DeepCopyInto(&out.Quota)
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
	in.Total.DeepCopyInto(&out.Total)
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make(ClusterQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
func (in *ClusterQuotaStatusByNamespace) DeepCopyInto(out *ClusterQuotaStatusByNamespace) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuotaStatusByNamespace.
func (in *ClusterQuotaStatusByNamespace) DeepCopy() *ClusterQuotaStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ClusterQuotaStatusByNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ClusterQuotasStatusByNamespace) DeepCopyInto(out *ClusterQuotasStatusByNamespace) {
	{
		in := &in
		*out = make(ClusterQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterQuotasStatusByNamespace.
func (in ClusterQuotasStatusByNamespace) DeepCopy() ClusterQuotasStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ClusterQuotasStatusByNamespace)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleRef) DeepCopyInto(out *ClusterRoleRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleRef.
func (in *ClusterRoleRef) DeepCopy() *ClusterRoleRef {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRoleTemplateTemplate) DeepCopyInto(out *ClusterRoleTemplateTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AggregationRule != nil {
		in, out := &in.AggregationRule, &out.AggregationRule
		*out = new(rbacv1.AggregationRule)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRoleTemplateTemplate.
func (in *ClusterRoleTemplateTemplate) DeepCopy() *ClusterRoleTemplateTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterRoleTemplateTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConstraintSpaceTemplate) DeepCopyInto(out *ConstraintSpaceTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ClusterRole != nil {
		in, out := &in.ClusterRole, &out.ClusterRole
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConstraintSpaceTemplate.
func (in *ConstraintSpaceTemplate) DeepCopy() *ConstraintSpaceTemplate {
	if in == nil {
		return nil
	}
	out := new(ConstraintSpaceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterAccess) DeepCopyInto(out *LocalClusterAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
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
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]UserOrTeam, len(*in))
		copy(*out, *in)
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterRoles != nil {
		in, out := &in.ClusterRoles, &out.ClusterRoles
		*out = make([]ClusterRoleRef, len(*in))
		copy(*out, *in)
	}
	if in.SpaceConstraintsRef != nil {
		in, out := &in.SpaceConstraintsRef, &out.SpaceConstraintsRef
		*out = new(string)
		**out = **in
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(AccessQuota)
		(*in).DeepCopyInto(*out)
	}
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
func (in *LocalClusterRoleTemplate) DeepCopyInto(out *LocalClusterRoleTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterRoleTemplate.
func (in *LocalClusterRoleTemplate) DeepCopy() *LocalClusterRoleTemplate {
	if in == nil {
		return nil
	}
	out := new(LocalClusterRoleTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalClusterRoleTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterRoleTemplateList) DeepCopyInto(out *LocalClusterRoleTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalClusterRoleTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterRoleTemplateList.
func (in *LocalClusterRoleTemplateList) DeepCopy() *LocalClusterRoleTemplateList {
	if in == nil {
		return nil
	}
	out := new(LocalClusterRoleTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalClusterRoleTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterRoleTemplateSpec) DeepCopyInto(out *LocalClusterRoleTemplateSpec) {
	*out = *in
	in.ClusterRoleTemplate.DeepCopyInto(&out.ClusterRoleTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterRoleTemplateSpec.
func (in *LocalClusterRoleTemplateSpec) DeepCopy() *LocalClusterRoleTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(LocalClusterRoleTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalClusterRoleTemplateStatus) DeepCopyInto(out *LocalClusterRoleTemplateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalClusterRoleTemplateStatus.
func (in *LocalClusterRoleTemplateStatus) DeepCopy() *LocalClusterRoleTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(LocalClusterRoleTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpaceConstraint) DeepCopyInto(out *LocalSpaceConstraint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpaceConstraint.
func (in *LocalSpaceConstraint) DeepCopy() *LocalSpaceConstraint {
	if in == nil {
		return nil
	}
	out := new(LocalSpaceConstraint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalSpaceConstraint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpaceConstraintList) DeepCopyInto(out *LocalSpaceConstraintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalSpaceConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpaceConstraintList.
func (in *LocalSpaceConstraintList) DeepCopy() *LocalSpaceConstraintList {
	if in == nil {
		return nil
	}
	out := new(LocalSpaceConstraintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalSpaceConstraintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpaceConstraintSpec) DeepCopyInto(out *LocalSpaceConstraintSpec) {
	*out = *in
	in.SpaceTemplate.DeepCopyInto(&out.SpaceTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpaceConstraintSpec.
func (in *LocalSpaceConstraintSpec) DeepCopy() *LocalSpaceConstraintSpec {
	if in == nil {
		return nil
	}
	out := new(LocalSpaceConstraintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSpaceConstraintStatus) DeepCopyInto(out *LocalSpaceConstraintStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSpaceConstraintStatus.
func (in *LocalSpaceConstraintStatus) DeepCopy() *LocalSpaceConstraintStatus {
	if in == nil {
		return nil
	}
	out := new(LocalSpaceConstraintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalTeam) DeepCopyInto(out *LocalTeam) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalTeam.
func (in *LocalTeam) DeepCopy() *LocalTeam {
	if in == nil {
		return nil
	}
	out := new(LocalTeam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalTeam) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalTeamList) DeepCopyInto(out *LocalTeamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalTeam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalTeamList.
func (in *LocalTeamList) DeepCopy() *LocalTeamList {
	if in == nil {
		return nil
	}
	out := new(LocalTeamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalTeamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalTeamSpec) DeepCopyInto(out *LocalTeamSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalTeamSpec.
func (in *LocalTeamSpec) DeepCopy() *LocalTeamSpec {
	if in == nil {
		return nil
	}
	out := new(LocalTeamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalTeamStatus) DeepCopyInto(out *LocalTeamStatus) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalTeamStatus.
func (in *LocalTeamStatus) DeepCopy() *LocalTeamStatus {
	if in == nil {
		return nil
	}
	out := new(LocalTeamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalUser) DeepCopyInto(out *LocalUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalUser.
func (in *LocalUser) DeepCopy() *LocalUser {
	if in == nil {
		return nil
	}
	out := new(LocalUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalUserList) DeepCopyInto(out *LocalUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalUserList.
func (in *LocalUserList) DeepCopy() *LocalUserList {
	if in == nil {
		return nil
	}
	out := new(LocalUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalUserSpec) DeepCopyInto(out *LocalUserSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalUserSpec.
func (in *LocalUserSpec) DeepCopy() *LocalUserSpec {
	if in == nil {
		return nil
	}
	out := new(LocalUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalUserStatus) DeepCopyInto(out *LocalUserStatus) {
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalUserStatus.
func (in *LocalUserStatus) DeepCopy() *LocalUserStatus {
	if in == nil {
		return nil
	}
	out := new(LocalUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSelector) DeepCopyInto(out *PodSelector) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSelector.
func (in *PodSelector) DeepCopy() *PodSelector {
	if in == nil {
		return nil
	}
	out := new(PodSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserOrTeam) DeepCopyInto(out *UserOrTeam) {
	*out = *in
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
func (in *VirtualClusterAccess) DeepCopyInto(out *VirtualClusterAccess) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]VirtualClusterAccessRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterAccess.
func (in *VirtualClusterAccess) DeepCopy() *VirtualClusterAccess {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterAccessRule) DeepCopyInto(out *VirtualClusterAccessRule) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterAccessRule.
func (in *VirtualClusterAccessRule) DeepCopy() *VirtualClusterAccessRule {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterAccessRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterHelmChart) DeepCopyInto(out *VirtualClusterHelmChart) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterHelmChart.
func (in *VirtualClusterHelmChart) DeepCopy() *VirtualClusterHelmChart {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterHelmChart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterHelmRelease) DeepCopyInto(out *VirtualClusterHelmRelease) {
	*out = *in
	out.Chart = in.Chart
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterHelmRelease.
func (in *VirtualClusterHelmRelease) DeepCopy() *VirtualClusterHelmRelease {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterHelmRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterHelmReleaseStatus) DeepCopyInto(out *VirtualClusterHelmReleaseStatus) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	out.Release = in.Release
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterHelmReleaseStatus.
func (in *VirtualClusterHelmReleaseStatus) DeepCopy() *VirtualClusterHelmReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterHelmReleaseStatus)
	in.DeepCopyInto(out)
	return out
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
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = new(VirtualClusterAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.HelmRelease != nil {
		in, out := &in.HelmRelease, &out.HelmRelease
		*out = new(VirtualClusterHelmRelease)
		**out = **in
	}
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(PodSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeConfigRef != nil {
		in, out := &in.KubeConfigRef, &out.KubeConfigRef
		*out = new(SecretRef)
		**out = **in
	}
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
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HelmRelease != nil {
		in, out := &in.HelmRelease, &out.HelmRelease
		*out = new(VirtualClusterHelmReleaseStatus)
		(*in).DeepCopyInto(*out)
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
