//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackage) DeepCopyInto(out *GlobalPythonPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackage.
func (in *GlobalPythonPackage) DeepCopy() *GlobalPythonPackage {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalPythonPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageCollection) DeepCopyInto(out *GlobalPythonPackageCollection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageCollection.
func (in *GlobalPythonPackageCollection) DeepCopy() *GlobalPythonPackageCollection {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalPythonPackageCollection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageCollectionList) DeepCopyInto(out *GlobalPythonPackageCollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalPythonPackageCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageCollectionList.
func (in *GlobalPythonPackageCollectionList) DeepCopy() *GlobalPythonPackageCollectionList {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageCollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalPythonPackageCollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageCollectionSpec) DeepCopyInto(out *GlobalPythonPackageCollectionSpec) {
	*out = *in
	in.PersistentVolumeClaimSpec.DeepCopyInto(&out.PersistentVolumeClaimSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageCollectionSpec.
func (in *GlobalPythonPackageCollectionSpec) DeepCopy() *GlobalPythonPackageCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageCollectionStatus) DeepCopyInto(out *GlobalPythonPackageCollectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageCollectionStatus.
func (in *GlobalPythonPackageCollectionStatus) DeepCopy() *GlobalPythonPackageCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageList) DeepCopyInto(out *GlobalPythonPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalPythonPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageList.
func (in *GlobalPythonPackageList) DeepCopy() *GlobalPythonPackageList {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalPythonPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackagePythonLibrarySpec) DeepCopyInto(out *GlobalPythonPackagePythonLibrarySpec) {
	*out = *in
	out.PythonSpec = in.PythonSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackagePythonLibrarySpec.
func (in *GlobalPythonPackagePythonLibrarySpec) DeepCopy() *GlobalPythonPackagePythonLibrarySpec {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackagePythonLibrarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackagePythonSpec) DeepCopyInto(out *GlobalPythonPackagePythonSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackagePythonSpec.
func (in *GlobalPythonPackagePythonSpec) DeepCopy() *GlobalPythonPackagePythonSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackagePythonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageSpec) DeepCopyInto(out *GlobalPythonPackageSpec) {
	*out = *in
	out.PythonLibrarySpec = in.PythonLibrarySpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageSpec.
func (in *GlobalPythonPackageSpec) DeepCopy() *GlobalPythonPackageSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPythonPackageStatus) DeepCopyInto(out *GlobalPythonPackageStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPythonPackageStatus.
func (in *GlobalPythonPackageStatus) DeepCopy() *GlobalPythonPackageStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalPythonPackageStatus)
	in.DeepCopyInto(out)
	return out
}
