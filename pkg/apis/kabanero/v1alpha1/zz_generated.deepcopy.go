// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CliStatus) DeepCopyInto(out *CliStatus) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CliStatus.
func (in *CliStatus) DeepCopy() *CliStatus {
	if in == nil {
		return nil
	}
	out := new(CliStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collection) DeepCopyInto(out *Collection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collection.
func (in *Collection) DeepCopy() *Collection {
	if in == nil {
		return nil
	}
	out := new(Collection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Collection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionList) DeepCopyInto(out *CollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Collection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionList.
func (in *CollectionList) DeepCopy() *CollectionList {
	if in == nil {
		return nil
	}
	out := new(CollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionSpec) DeepCopyInto(out *CollectionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionSpec.
func (in *CollectionSpec) DeepCopy() *CollectionSpec {
	if in == nil {
		return nil
	}
	out := new(CollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionStatus) DeepCopyInto(out *CollectionStatus) {
	*out = *in
	if in.ActiveAssets != nil {
		in, out := &in.ActiveAssets, &out.ActiveAssets
		*out = make([]RepositoryAssetStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionStatus.
func (in *CollectionStatus) DeepCopy() *CollectionStatus {
	if in == nil {
		return nil
	}
	out := new(CollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceCollectionConfig) DeepCopyInto(out *InstanceCollectionConfig) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]RepositoryConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceCollectionConfig.
func (in *InstanceCollectionConfig) DeepCopy() *InstanceCollectionConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceCollectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kabanero) DeepCopyInto(out *Kabanero) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kabanero.
func (in *Kabanero) DeepCopy() *Kabanero {
	if in == nil {
		return nil
	}
	out := new(Kabanero)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kabanero) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroInstanceStatus) DeepCopyInto(out *KabaneroInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroInstanceStatus.
func (in *KabaneroInstanceStatus) DeepCopy() *KabaneroInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroList) DeepCopyInto(out *KabaneroList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kabanero, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroList.
func (in *KabaneroList) DeepCopy() *KabaneroList {
	if in == nil {
		return nil
	}
	out := new(KabaneroList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KabaneroList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroSpec) DeepCopyInto(out *KabaneroSpec) {
	*out = *in
	in.Collections.DeepCopyInto(&out.Collections)
	out.Tekton = in.Tekton
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroSpec.
func (in *KabaneroSpec) DeepCopy() *KabaneroSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroStatus) DeepCopyInto(out *KabaneroStatus) {
	*out = *in
	out.KabaneroInstance = in.KabaneroInstance
	out.KnativeEventing = in.KnativeEventing
	out.KnativeServing = in.KnativeServing
	out.Tekton = in.Tekton
	in.Cli.DeepCopyInto(&out.Cli)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroStatus.
func (in *KabaneroStatus) DeepCopy() *KabaneroStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeEventingStatus) DeepCopyInto(out *KnativeEventingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeEventingStatus.
func (in *KnativeEventingStatus) DeepCopy() *KnativeEventingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeEventingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServingStatus) DeepCopyInto(out *KnativeServingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServingStatus.
func (in *KnativeServingStatus) DeepCopy() *KnativeServingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeServingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryAssetStatus) DeepCopyInto(out *RepositoryAssetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryAssetStatus.
func (in *RepositoryAssetStatus) DeepCopy() *RepositoryAssetStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryAssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryConfig) DeepCopyInto(out *RepositoryConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryConfig.
func (in *RepositoryConfig) DeepCopy() *RepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(RepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonCustomizationSpec) DeepCopyInto(out *TektonCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonCustomizationSpec.
func (in *TektonCustomizationSpec) DeepCopy() *TektonCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(TektonCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonStatus) DeepCopyInto(out *TektonStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonStatus.
func (in *TektonStatus) DeepCopy() *TektonStatus {
	if in == nil {
		return nil
	}
	out := new(TektonStatus)
	in.DeepCopyInto(out)
	return out
}
