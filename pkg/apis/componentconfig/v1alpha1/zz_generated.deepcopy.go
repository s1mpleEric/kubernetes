// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientConnectionConfiguration) DeepCopyInto(out *ClientConnectionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientConnectionConfiguration.
func (in *ClientConnectionConfiguration) DeepCopy() *ClientConnectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(ClientConnectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerConfiguration) DeepCopyInto(out *KubeControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.NodeSyncPeriod = in.NodeSyncPeriod
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	out.MinResyncPeriod = in.MinResyncPeriod
	out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
	out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
	out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
	out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
	out.PodEvictionTimeout = in.PodEvictionTimeout
	out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
	out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.ClusterSigningDuration = in.ClusterSigningDuration
	if in.ConfigureCloudRoutes != nil {
		in, out := &in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	in.VolumeConfiguration.DeepCopyInto(&out.VolumeConfiguration)
	out.ControllerStartInterval = in.ControllerStartInterval
	if in.EnableGarbageCollector != nil {
		in, out := &in.EnableGarbageCollector, &out.EnableGarbageCollector
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.GCIgnoredResources != nil {
		in, out := &in.GCIgnoredResources, &out.GCIgnoredResources
		*out = make([]GroupResource, len(*in))
		copy(*out, *in)
	}
	out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
	if in.EnableTaintManager != nil {
		in, out := &in.EnableTaintManager, &out.EnableTaintManager
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.HorizontalPodAutoscalerUseRESTClients != nil {
		in, out := &in.HorizontalPodAutoscalerUseRESTClients, &out.HorizontalPodAutoscalerUseRESTClients
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerConfiguration.
func (in *KubeControllerManagerConfiguration) DeepCopy() *KubeControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.AlgorithmSource.DeepCopyInto(&out.AlgorithmSource)
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	out.ClientConnection = in.ClientConnection
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfiguration.
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopyInto(out *KubeSchedulerLeaderElectionConfiguration) {
	*out = *in
	in.LeaderElectionConfiguration.DeepCopyInto(&out.LeaderElectionConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerLeaderElectionConfiguration.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopy() *KubeSchedulerLeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerLeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	if in.LeaderElect != nil {
		in, out := &in.LeaderElect, &out.LeaderElect
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeRecyclerConfiguration) DeepCopyInto(out *PersistentVolumeRecyclerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeRecyclerConfiguration.
func (in *PersistentVolumeRecyclerConfiguration) DeepCopy() *PersistentVolumeRecyclerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeRecyclerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerAlgorithmSource) DeepCopyInto(out *SchedulerAlgorithmSource) {
	*out = *in
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		if *in == nil {
			*out = nil
		} else {
			*out = new(SchedulerPolicySource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerAlgorithmSource.
func (in *SchedulerAlgorithmSource) DeepCopy() *SchedulerAlgorithmSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerAlgorithmSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyConfigMapSource) DeepCopyInto(out *SchedulerPolicyConfigMapSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyConfigMapSource.
func (in *SchedulerPolicyConfigMapSource) DeepCopy() *SchedulerPolicyConfigMapSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyConfigMapSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicyFileSource) DeepCopyInto(out *SchedulerPolicyFileSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicyFileSource.
func (in *SchedulerPolicyFileSource) DeepCopy() *SchedulerPolicyFileSource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicyFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerPolicySource) DeepCopyInto(out *SchedulerPolicySource) {
	*out = *in
	if in.File != nil {
		in, out := &in.File, &out.File
		if *in == nil {
			*out = nil
		} else {
			*out = new(SchedulerPolicyFileSource)
			**out = **in
		}
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		if *in == nil {
			*out = nil
		} else {
			*out = new(SchedulerPolicyConfigMapSource)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerPolicySource.
func (in *SchedulerPolicySource) DeepCopy() *SchedulerPolicySource {
	if in == nil {
		return nil
	}
	out := new(SchedulerPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConfiguration) DeepCopyInto(out *VolumeConfiguration) {
	*out = *in
	if in.EnableHostPathProvisioning != nil {
		in, out := &in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.EnableDynamicProvisioning != nil {
		in, out := &in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	out.PersistentVolumeRecyclerConfiguration = in.PersistentVolumeRecyclerConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConfiguration.
func (in *VolumeConfiguration) DeepCopy() *VolumeConfiguration {
	if in == nil {
		return nil
	}
	out := new(VolumeConfiguration)
	in.DeepCopyInto(out)
	return out
}
