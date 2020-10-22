// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAccessKey) DeepCopyInto(out *IAMAccessKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAccessKey.
func (in *IAMAccessKey) DeepCopy() *IAMAccessKey {
	if in == nil {
		return nil
	}
	out := new(IAMAccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMAccessKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAccessKeyList) DeepCopyInto(out *IAMAccessKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMAccessKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAccessKeyList.
func (in *IAMAccessKeyList) DeepCopy() *IAMAccessKeyList {
	if in == nil {
		return nil
	}
	out := new(IAMAccessKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMAccessKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAccessKeyObservation) DeepCopyInto(out *IAMAccessKeyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAccessKeyObservation.
func (in *IAMAccessKeyObservation) DeepCopy() *IAMAccessKeyObservation {
	if in == nil {
		return nil
	}
	out := new(IAMAccessKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAccessKeyParameters) DeepCopyInto(out *IAMAccessKeyParameters) {
	*out = *in
	if in.IAMUsernameRef != nil {
		in, out := &in.IAMUsernameRef, &out.IAMUsernameRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.IAMUsernameSelector != nil {
		in, out := &in.IAMUsernameSelector, &out.IAMUsernameSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAccessKeyParameters.
func (in *IAMAccessKeyParameters) DeepCopy() *IAMAccessKeyParameters {
	if in == nil {
		return nil
	}
	out := new(IAMAccessKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAccessKeySpec) DeepCopyInto(out *IAMAccessKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAccessKeySpec.
func (in *IAMAccessKeySpec) DeepCopy() *IAMAccessKeySpec {
	if in == nil {
		return nil
	}
	out := new(IAMAccessKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMAccessKeyStatus) DeepCopyInto(out *IAMAccessKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMAccessKeyStatus.
func (in *IAMAccessKeyStatus) DeepCopy() *IAMAccessKeyStatus {
	if in == nil {
		return nil
	}
	out := new(IAMAccessKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroup) DeepCopyInto(out *IAMGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroup.
func (in *IAMGroup) DeepCopy() *IAMGroup {
	if in == nil {
		return nil
	}
	out := new(IAMGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupList) DeepCopyInto(out *IAMGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupList.
func (in *IAMGroupList) DeepCopy() *IAMGroupList {
	if in == nil {
		return nil
	}
	out := new(IAMGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupObservation) DeepCopyInto(out *IAMGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupObservation.
func (in *IAMGroupObservation) DeepCopy() *IAMGroupObservation {
	if in == nil {
		return nil
	}
	out := new(IAMGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupParameters) DeepCopyInto(out *IAMGroupParameters) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupParameters.
func (in *IAMGroupParameters) DeepCopy() *IAMGroupParameters {
	if in == nil {
		return nil
	}
	out := new(IAMGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupPolicyAttachment) DeepCopyInto(out *IAMGroupPolicyAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupPolicyAttachment.
func (in *IAMGroupPolicyAttachment) DeepCopy() *IAMGroupPolicyAttachment {
	if in == nil {
		return nil
	}
	out := new(IAMGroupPolicyAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMGroupPolicyAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupPolicyAttachmentList) DeepCopyInto(out *IAMGroupPolicyAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMGroupPolicyAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupPolicyAttachmentList.
func (in *IAMGroupPolicyAttachmentList) DeepCopy() *IAMGroupPolicyAttachmentList {
	if in == nil {
		return nil
	}
	out := new(IAMGroupPolicyAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMGroupPolicyAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupPolicyAttachmentObservation) DeepCopyInto(out *IAMGroupPolicyAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupPolicyAttachmentObservation.
func (in *IAMGroupPolicyAttachmentObservation) DeepCopy() *IAMGroupPolicyAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(IAMGroupPolicyAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupPolicyAttachmentParameters) DeepCopyInto(out *IAMGroupPolicyAttachmentParameters) {
	*out = *in
	if in.PolicyARNRef != nil {
		in, out := &in.PolicyARNRef, &out.PolicyARNRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.PolicyARNSelector != nil {
		in, out := &in.PolicyARNSelector, &out.PolicyARNSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupNameRef != nil {
		in, out := &in.GroupNameRef, &out.GroupNameRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.GroupNameSelector != nil {
		in, out := &in.GroupNameSelector, &out.GroupNameSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupPolicyAttachmentParameters.
func (in *IAMGroupPolicyAttachmentParameters) DeepCopy() *IAMGroupPolicyAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(IAMGroupPolicyAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupPolicyAttachmentSpec) DeepCopyInto(out *IAMGroupPolicyAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupPolicyAttachmentSpec.
func (in *IAMGroupPolicyAttachmentSpec) DeepCopy() *IAMGroupPolicyAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(IAMGroupPolicyAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupPolicyAttachmentStatus) DeepCopyInto(out *IAMGroupPolicyAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupPolicyAttachmentStatus.
func (in *IAMGroupPolicyAttachmentStatus) DeepCopy() *IAMGroupPolicyAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(IAMGroupPolicyAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupSpec) DeepCopyInto(out *IAMGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupSpec.
func (in *IAMGroupSpec) DeepCopy() *IAMGroupSpec {
	if in == nil {
		return nil
	}
	out := new(IAMGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupStatus) DeepCopyInto(out *IAMGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupStatus.
func (in *IAMGroupStatus) DeepCopy() *IAMGroupStatus {
	if in == nil {
		return nil
	}
	out := new(IAMGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupUserMembership) DeepCopyInto(out *IAMGroupUserMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupUserMembership.
func (in *IAMGroupUserMembership) DeepCopy() *IAMGroupUserMembership {
	if in == nil {
		return nil
	}
	out := new(IAMGroupUserMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMGroupUserMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupUserMembershipList) DeepCopyInto(out *IAMGroupUserMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMGroupUserMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupUserMembershipList.
func (in *IAMGroupUserMembershipList) DeepCopy() *IAMGroupUserMembershipList {
	if in == nil {
		return nil
	}
	out := new(IAMGroupUserMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMGroupUserMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupUserMembershipObservation) DeepCopyInto(out *IAMGroupUserMembershipObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupUserMembershipObservation.
func (in *IAMGroupUserMembershipObservation) DeepCopy() *IAMGroupUserMembershipObservation {
	if in == nil {
		return nil
	}
	out := new(IAMGroupUserMembershipObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupUserMembershipParameters) DeepCopyInto(out *IAMGroupUserMembershipParameters) {
	*out = *in
	if in.GroupNameRef != nil {
		in, out := &in.GroupNameRef, &out.GroupNameRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.GroupNameSelector != nil {
		in, out := &in.GroupNameSelector, &out.GroupNameSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.UserNameRef != nil {
		in, out := &in.UserNameRef, &out.UserNameRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.UserNameSelector != nil {
		in, out := &in.UserNameSelector, &out.UserNameSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupUserMembershipParameters.
func (in *IAMGroupUserMembershipParameters) DeepCopy() *IAMGroupUserMembershipParameters {
	if in == nil {
		return nil
	}
	out := new(IAMGroupUserMembershipParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupUserMembershipSpec) DeepCopyInto(out *IAMGroupUserMembershipSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupUserMembershipSpec.
func (in *IAMGroupUserMembershipSpec) DeepCopy() *IAMGroupUserMembershipSpec {
	if in == nil {
		return nil
	}
	out := new(IAMGroupUserMembershipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMGroupUserMembershipStatus) DeepCopyInto(out *IAMGroupUserMembershipStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMGroupUserMembershipStatus.
func (in *IAMGroupUserMembershipStatus) DeepCopy() *IAMGroupUserMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(IAMGroupUserMembershipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicy) DeepCopyInto(out *IAMPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicy.
func (in *IAMPolicy) DeepCopy() *IAMPolicy {
	if in == nil {
		return nil
	}
	out := new(IAMPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyList) DeepCopyInto(out *IAMPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyList.
func (in *IAMPolicyList) DeepCopy() *IAMPolicyList {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyObservation) DeepCopyInto(out *IAMPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyObservation.
func (in *IAMPolicyObservation) DeepCopy() *IAMPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyParameters) DeepCopyInto(out *IAMPolicyParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyParameters.
func (in *IAMPolicyParameters) DeepCopy() *IAMPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicySpec) DeepCopyInto(out *IAMPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicySpec.
func (in *IAMPolicySpec) DeepCopy() *IAMPolicySpec {
	if in == nil {
		return nil
	}
	out := new(IAMPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMPolicyStatus) DeepCopyInto(out *IAMPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMPolicyStatus.
func (in *IAMPolicyStatus) DeepCopy() *IAMPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(IAMPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUser) DeepCopyInto(out *IAMUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUser.
func (in *IAMUser) DeepCopy() *IAMUser {
	if in == nil {
		return nil
	}
	out := new(IAMUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserList) DeepCopyInto(out *IAMUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserList.
func (in *IAMUserList) DeepCopy() *IAMUserList {
	if in == nil {
		return nil
	}
	out := new(IAMUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserObservation) DeepCopyInto(out *IAMUserObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserObservation.
func (in *IAMUserObservation) DeepCopy() *IAMUserObservation {
	if in == nil {
		return nil
	}
	out := new(IAMUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserParameters) DeepCopyInto(out *IAMUserParameters) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.PermissionsBoundary != nil {
		in, out := &in.PermissionsBoundary, &out.PermissionsBoundary
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserParameters.
func (in *IAMUserParameters) DeepCopy() *IAMUserParameters {
	if in == nil {
		return nil
	}
	out := new(IAMUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserPolicyAttachment) DeepCopyInto(out *IAMUserPolicyAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserPolicyAttachment.
func (in *IAMUserPolicyAttachment) DeepCopy() *IAMUserPolicyAttachment {
	if in == nil {
		return nil
	}
	out := new(IAMUserPolicyAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMUserPolicyAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserPolicyAttachmentList) DeepCopyInto(out *IAMUserPolicyAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IAMUserPolicyAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserPolicyAttachmentList.
func (in *IAMUserPolicyAttachmentList) DeepCopy() *IAMUserPolicyAttachmentList {
	if in == nil {
		return nil
	}
	out := new(IAMUserPolicyAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IAMUserPolicyAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserPolicyAttachmentObservation) DeepCopyInto(out *IAMUserPolicyAttachmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserPolicyAttachmentObservation.
func (in *IAMUserPolicyAttachmentObservation) DeepCopy() *IAMUserPolicyAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(IAMUserPolicyAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserPolicyAttachmentParameters) DeepCopyInto(out *IAMUserPolicyAttachmentParameters) {
	*out = *in
	if in.PolicyARNRef != nil {
		in, out := &in.PolicyARNRef, &out.PolicyARNRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.PolicyARNSelector != nil {
		in, out := &in.PolicyARNSelector, &out.PolicyARNSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.UserNameRef != nil {
		in, out := &in.UserNameRef, &out.UserNameRef
		*out = new(corev1alpha1.Reference)
		**out = **in
	}
	if in.UserNameSelector != nil {
		in, out := &in.UserNameSelector, &out.UserNameSelector
		*out = new(corev1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserPolicyAttachmentParameters.
func (in *IAMUserPolicyAttachmentParameters) DeepCopy() *IAMUserPolicyAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(IAMUserPolicyAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserPolicyAttachmentSpec) DeepCopyInto(out *IAMUserPolicyAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserPolicyAttachmentSpec.
func (in *IAMUserPolicyAttachmentSpec) DeepCopy() *IAMUserPolicyAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(IAMUserPolicyAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserPolicyAttachmentStatus) DeepCopyInto(out *IAMUserPolicyAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserPolicyAttachmentStatus.
func (in *IAMUserPolicyAttachmentStatus) DeepCopy() *IAMUserPolicyAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(IAMUserPolicyAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserSpec) DeepCopyInto(out *IAMUserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserSpec.
func (in *IAMUserSpec) DeepCopy() *IAMUserSpec {
	if in == nil {
		return nil
	}
	out := new(IAMUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IAMUserStatus) DeepCopyInto(out *IAMUserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IAMUserStatus.
func (in *IAMUserStatus) DeepCopy() *IAMUserStatus {
	if in == nil {
		return nil
	}
	out := new(IAMUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
