// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat Community of Practice.

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
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIDRAssignment) DeepCopyInto(out *CIDRAssignment) {
	*out = *in
	if in.ReservedIPs != nil {
		in, out := &in.ReservedIPs, &out.ReservedIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRAssignment.
func (in *CIDRAssignment) DeepCopy() *CIDRAssignment {
	if in == nil {
		return nil
	}
	out := new(CIDRAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressIPAM) DeepCopyInto(out *EgressIPAM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressIPAM.
func (in *EgressIPAM) DeepCopy() *EgressIPAM {
	if in == nil {
		return nil
	}
	out := new(EgressIPAM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgressIPAM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressIPAMList) DeepCopyInto(out *EgressIPAMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EgressIPAM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressIPAMList.
func (in *EgressIPAMList) DeepCopy() *EgressIPAMList {
	if in == nil {
		return nil
	}
	out := new(EgressIPAMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EgressIPAMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressIPAMSpec) DeepCopyInto(out *EgressIPAMSpec) {
	*out = *in
	if in.CIDRAssignments != nil {
		in, out := &in.CIDRAssignments, &out.CIDRAssignments
		*out = make([]CIDRAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressIPAMSpec.
func (in *EgressIPAMSpec) DeepCopy() *EgressIPAMSpec {
	if in == nil {
		return nil
	}
	out := new(EgressIPAMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressIPAMStatus) DeepCopyInto(out *EgressIPAMStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressIPAMStatus.
func (in *EgressIPAMStatus) DeepCopy() *EgressIPAMStatus {
	if in == nil {
		return nil
	}
	out := new(EgressIPAMStatus)
	in.DeepCopyInto(out)
	return out
}