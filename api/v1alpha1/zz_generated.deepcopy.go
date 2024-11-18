//go:build !ignore_autogenerated

/*
Copyright 2024.

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
	"github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackup) DeepCopyInto(out *NonAdminBackup) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackup.
func (in *NonAdminBackup) DeepCopy() *NonAdminBackup {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NonAdminBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupList) DeepCopyInto(out *NonAdminBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NonAdminBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupList.
func (in *NonAdminBackupList) DeepCopy() *NonAdminBackupList {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NonAdminBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupSpec) DeepCopyInto(out *NonAdminBackupSpec) {
	*out = *in
	if in.BackupSpec != nil {
		in, out := &in.BackupSpec, &out.BackupSpec
		*out = new(v1.BackupSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupSpec.
func (in *NonAdminBackupSpec) DeepCopy() *NonAdminBackupSpec {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupStatus) DeepCopyInto(out *NonAdminBackupStatus) {
	*out = *in
	if in.VeleroBackup != nil {
		in, out := &in.VeleroBackup, &out.VeleroBackup
		*out = new(VeleroBackup)
		(*in).DeepCopyInto(*out)
	}
	if in.VeleroDeleteBackupRequest != nil {
		in, out := &in.VeleroDeleteBackupRequest, &out.VeleroDeleteBackupRequest
		*out = new(VeleroDeleteBackupRequest)
		(*in).DeepCopyInto(*out)
	}
	if in.QueueInfo != nil {
		in, out := &in.QueueInfo, &out.QueueInfo
		*out = new(QueueInfo)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupStatus.
func (in *NonAdminBackupStatus) DeepCopy() *NonAdminBackupStatus {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupStorageLocation) DeepCopyInto(out *NonAdminBackupStorageLocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupStorageLocation.
func (in *NonAdminBackupStorageLocation) DeepCopy() *NonAdminBackupStorageLocation {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupStorageLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NonAdminBackupStorageLocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupStorageLocationList) DeepCopyInto(out *NonAdminBackupStorageLocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NonAdminBackupStorageLocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupStorageLocationList.
func (in *NonAdminBackupStorageLocationList) DeepCopy() *NonAdminBackupStorageLocationList {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupStorageLocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NonAdminBackupStorageLocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupStorageLocationSpec) DeepCopyInto(out *NonAdminBackupStorageLocationSpec) {
	*out = *in
	in.BackupStorageLocationSpec.DeepCopyInto(&out.BackupStorageLocationSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupStorageLocationSpec.
func (in *NonAdminBackupStorageLocationSpec) DeepCopy() *NonAdminBackupStorageLocationSpec {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupStorageLocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminBackupStorageLocationStatus) DeepCopyInto(out *NonAdminBackupStorageLocationStatus) {
	*out = *in
	if in.VeleroBackupStorageLocation != nil {
		in, out := &in.VeleroBackupStorageLocation, &out.VeleroBackupStorageLocation
		*out = new(VeleroBackupStorageLocation)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminBackupStorageLocationStatus.
func (in *NonAdminBackupStorageLocationStatus) DeepCopy() *NonAdminBackupStorageLocationStatus {
	if in == nil {
		return nil
	}
	out := new(NonAdminBackupStorageLocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminRestore) DeepCopyInto(out *NonAdminRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminRestore.
func (in *NonAdminRestore) DeepCopy() *NonAdminRestore {
	if in == nil {
		return nil
	}
	out := new(NonAdminRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NonAdminRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminRestoreList) DeepCopyInto(out *NonAdminRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NonAdminRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminRestoreList.
func (in *NonAdminRestoreList) DeepCopy() *NonAdminRestoreList {
	if in == nil {
		return nil
	}
	out := new(NonAdminRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NonAdminRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminRestoreSpec) DeepCopyInto(out *NonAdminRestoreSpec) {
	*out = *in
	if in.RestoreSpec != nil {
		in, out := &in.RestoreSpec, &out.RestoreSpec
		*out = new(v1.RestoreSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminRestoreSpec.
func (in *NonAdminRestoreSpec) DeepCopy() *NonAdminRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(NonAdminRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonAdminRestoreStatus) DeepCopyInto(out *NonAdminRestoreStatus) {
	*out = *in
	if in.VeleroRestore != nil {
		in, out := &in.VeleroRestore, &out.VeleroRestore
		*out = new(VeleroRestore)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonAdminRestoreStatus.
func (in *NonAdminRestoreStatus) DeepCopy() *NonAdminRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(NonAdminRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueInfo) DeepCopyInto(out *QueueInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueInfo.
func (in *QueueInfo) DeepCopy() *QueueInfo {
	if in == nil {
		return nil
	}
	out := new(QueueInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroBackup) DeepCopyInto(out *VeleroBackup) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1.BackupStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroBackup.
func (in *VeleroBackup) DeepCopy() *VeleroBackup {
	if in == nil {
		return nil
	}
	out := new(VeleroBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroBackupStorageLocation) DeepCopyInto(out *VeleroBackupStorageLocation) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1.BackupStorageLocationStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroBackupStorageLocation.
func (in *VeleroBackupStorageLocation) DeepCopy() *VeleroBackupStorageLocation {
	if in == nil {
		return nil
	}
	out := new(VeleroBackupStorageLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroDeleteBackupRequest) DeepCopyInto(out *VeleroDeleteBackupRequest) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1.DeleteBackupRequestStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroDeleteBackupRequest.
func (in *VeleroDeleteBackupRequest) DeepCopy() *VeleroDeleteBackupRequest {
	if in == nil {
		return nil
	}
	out := new(VeleroDeleteBackupRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VeleroRestore) DeepCopyInto(out *VeleroRestore) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(v1.RestoreStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VeleroRestore.
func (in *VeleroRestore) DeepCopy() *VeleroRestore {
	if in == nil {
		return nil
	}
	out := new(VeleroRestore)
	in.DeepCopyInto(out)
	return out
}
