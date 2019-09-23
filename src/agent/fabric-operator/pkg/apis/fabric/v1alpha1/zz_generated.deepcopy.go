// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CA) DeepCopyInto(out *CA) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CA.
func (in *CA) DeepCopy() *CA {
	if in == nil {
		return nil
	}
	out := new(CA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CA) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CACerts) DeepCopyInto(out *CACerts) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CACerts.
func (in *CACerts) DeepCopy() *CACerts {
	if in == nil {
		return nil
	}
	out := new(CACerts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CAList) DeepCopyInto(out *CAList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CA, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CAList.
func (in *CAList) DeepCopy() *CAList {
	if in == nil {
		return nil
	}
	out := new(CAList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CAList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASpec) DeepCopyInto(out *CASpec) {
	*out = *in
	if in.Certs != nil {
		in, out := &in.Certs, &out.Certs
		*out = new(CACerts)
		**out = **in
	}
	in.NodeSpec.DeepCopyInto(&out.NodeSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASpec.
func (in *CASpec) DeepCopy() *CASpec {
	if in == nil {
		return nil
	}
	out := new(CASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certs) DeepCopyInto(out *Certs) {
	*out = *in
	if in.Msp != nil {
		in, out := &in.Msp, &out.Msp
		*out = new(Msp)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSCerts != nil {
		in, out := &in.TLSCerts, &out.TLSCerts
		*out = new(TLSCerts)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certs.
func (in *Certs) DeepCopy() *Certs {
	if in == nil {
		return nil
	}
	out := new(Certs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParam) DeepCopyInto(out *ConfigParam) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParam.
func (in *ConfigParam) DeepCopy() *ConfigParam {
	if in == nil {
		return nil
	}
	out := new(ConfigParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Msp) DeepCopyInto(out *Msp) {
	*out = *in
	if in.AdminCerts != nil {
		in, out := &in.AdminCerts, &out.AdminCerts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CaCerts != nil {
		in, out := &in.CaCerts, &out.CaCerts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IntermediateCerts != nil {
		in, out := &in.IntermediateCerts, &out.IntermediateCerts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSCacerts != nil {
		in, out := &in.TLSCacerts, &out.TLSCacerts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSIntermediatecerts != nil {
		in, out := &in.TLSIntermediatecerts, &out.TLSIntermediatecerts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Msp.
func (in *Msp) DeepCopy() *Msp {
	if in == nil {
		return nil
	}
	out := new(Msp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	if in.ConfigParams != nil {
		in, out := &in.ConfigParams, &out.ConfigParams
		*out = make([]ConfigParam, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Orderer) DeepCopyInto(out *Orderer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Orderer.
func (in *Orderer) DeepCopy() *Orderer {
	if in == nil {
		return nil
	}
	out := new(Orderer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Orderer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererList) DeepCopyInto(out *OrdererList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Orderer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererList.
func (in *OrdererList) DeepCopy() *OrdererList {
	if in == nil {
		return nil
	}
	out := new(OrdererList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrdererList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrdererSpec) DeepCopyInto(out *OrdererSpec) {
	*out = *in
	if in.Certs != nil {
		in, out := &in.Certs, &out.Certs
		*out = new(Certs)
		(*in).DeepCopyInto(*out)
	}
	in.NodeSpec.DeepCopyInto(&out.NodeSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererSpec.
func (in *OrdererSpec) DeepCopy() *OrdererSpec {
	if in == nil {
		return nil
	}
	out := new(OrdererSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Peer) DeepCopyInto(out *Peer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Peer.
func (in *Peer) DeepCopy() *Peer {
	if in == nil {
		return nil
	}
	out := new(Peer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Peer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerList) DeepCopyInto(out *PeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Peer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerList.
func (in *PeerList) DeepCopy() *PeerList {
	if in == nil {
		return nil
	}
	out := new(PeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerSpec) DeepCopyInto(out *PeerSpec) {
	*out = *in
	if in.Certs != nil {
		in, out := &in.Certs, &out.Certs
		*out = new(Certs)
		(*in).DeepCopyInto(*out)
	}
	in.NodeSpec.DeepCopyInto(&out.NodeSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerSpec.
func (in *PeerSpec) DeepCopy() *PeerSpec {
	if in == nil {
		return nil
	}
	out := new(PeerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSCerts) DeepCopyInto(out *TLSCerts) {
	*out = *in
	if in.TLSRootcas != nil {
		in, out := &in.TLSRootcas, &out.TLSRootcas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSCerts.
func (in *TLSCerts) DeepCopy() *TLSCerts {
	if in == nil {
		return nil
	}
	out := new(TLSCerts)
	in.DeepCopyInto(out)
	return out
}