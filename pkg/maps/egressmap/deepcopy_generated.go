// +build !ignore_autogenerated

// Copyright 2017-2021 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package egressmap

import (
	bpf "github.com/cilium/cilium/pkg/bpf"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressInfo4) DeepCopyInto(out *EgressInfo4) {
	*out = *in
	in.EgressIP.DeepCopyInto(&out.EgressIP)
	in.TunnelEndpoint.DeepCopyInto(&out.TunnelEndpoint)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressInfo4.
func (in *EgressInfo4) DeepCopy() *EgressInfo4 {
	if in == nil {
		return nil
	}
	out := new(EgressInfo4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapValue is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapValue.
func (in *EgressInfo4) DeepCopyMapValue() bpf.MapValue {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Key4) DeepCopyInto(out *Key4) {
	*out = *in
	in.SourceIP.DeepCopyInto(&out.SourceIP)
	in.DestCIDR.DeepCopyInto(&out.DestCIDR)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Key4.
func (in *Key4) DeepCopy() *Key4 {
	if in == nil {
		return nil
	}
	out := new(Key4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyMapKey is an autogenerated deepcopy function, copying the receiver, creating a new bpf.MapKey.
func (in *Key4) DeepCopyMapKey() bpf.MapKey {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}