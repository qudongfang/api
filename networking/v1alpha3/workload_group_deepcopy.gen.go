// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/workload_group.proto

// `WorkloadGroup` describes a collection of workload instances.
// It provides a specification that the workload instances can use to bootstrap
// their proxies, including the metadata and identity. It is only intended to
// be used with non-k8s workloads like Virtual Machines, and is meant to mimic
// the existing sidecar injection and deployment specification model used for
// Kubernetes workloads to bootstrap Istio proxies.
//
// The following example declares a workload group representing a collection
// of workloads that will be registered under `reviews` in namespace
// `bookinfo`. The set of labels will be associated with each workload
// instance during the bootstrap process, and the ports 3550 and 8080
// will be associated with the workload group and use service account `default`.
// `app.kubernetes.io/version` is just an arbitrary example of a label.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadGroup
// metadata:
//   name: reviews
//   namespace: bookinfo
// spec:
//   metadata:
//     labels:
//       app.kubernetes.io/name: reviews
//       app.kubernetes.io/version: "1.3.4"
//   template:
//     ports:
//       grpc: 3550
//       http: 8080
//     serviceAccount: default
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using WorkloadGroup within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadGroup) DeepCopyInto(out *WorkloadGroup) {
	p := proto.Clone(in).(*WorkloadGroup)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGroup. Required by controller-gen.
func (in *WorkloadGroup) DeepCopy() *WorkloadGroup {
	if in == nil {
		return nil
	}
	out := new(WorkloadGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto supports using WorkloadGroup_ObjectMeta within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadGroup_ObjectMeta) DeepCopyInto(out *WorkloadGroup_ObjectMeta) {
	p := proto.Clone(in).(*WorkloadGroup_ObjectMeta)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGroup_ObjectMeta. Required by controller-gen.
func (in *WorkloadGroup_ObjectMeta) DeepCopy() *WorkloadGroup_ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(WorkloadGroup_ObjectMeta)
	in.DeepCopyInto(out)
	return out
}