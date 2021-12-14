// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Code generated by controller-gen. DO NOT EDIT.

package v2alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APMFeatureConfig) DeepCopyInto(out *APMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.HostPortConfig != nil {
		in, out := &in.HostPortConfig, &out.HostPortConfig
		*out = new(HostPortConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.UnixDomainSocketConfig != nil {
		in, out := &in.UnixDomainSocketConfig, &out.UnixDomainSocketConfig
		*out = new(UnixDomainSocketConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APMFeatureConfig.
func (in *APMFeatureConfig) DeepCopy() *APMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(APMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerFeatureConfig) DeepCopyInto(out *AdmissionControllerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MutateUnlabelled != nil {
		in, out := &in.MutateUnlabelled, &out.MutateUnlabelled
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerFeatureConfig.
func (in *AdmissionControllerFeatureConfig) DeepCopy() *AdmissionControllerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSPMFeatureConfig) DeepCopyInto(out *CSPMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigMapConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSPMFeatureConfig.
func (in *CSPMFeatureConfig) DeepCopy() *CSPMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(CSPMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CWSFeatureConfig) DeepCopyInto(out *CWSFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigMapConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EnableSyscallMonitor != nil {
		in, out := &in.EnableSyscallMonitor, &out.EnableSyscallMonitor
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CWSFeatureConfig.
func (in *CWSFeatureConfig) DeepCopy() *CWSFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(CWSFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterChecksRunnerFeatureConfig) DeepCopyInto(out *ClusterChecksRunnerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterChecksRunnerFeatureConfig.
func (in *ClusterChecksRunnerFeatureConfig) DeepCopy() *ClusterChecksRunnerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterChecksRunnerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapConfig) DeepCopyInto(out *ConfigMapConfig) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.KeyToPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapConfig.
func (in *ConfigMapConfig) DeepCopy() *ConfigMapConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigMapConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerCollectionFeatureConfig) DeepCopyInto(out *ContainerCollectionFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerCollectionFeatureConfig.
func (in *ContainerCollectionFeatureConfig) DeepCopy() *ContainerCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(ContainerCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomConfig) DeepCopyInto(out *CustomConfig) {
	*out = *in
	if in.ConfigData != nil {
		in, out := &in.ConfigData, &out.ConfigData
		*out = new(string)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigMapConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomConfig.
func (in *CustomConfig) DeepCopy() *CustomConfig {
	if in == nil {
		return nil
	}
	out := new(CustomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgent) DeepCopyInto(out *DatadogAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgent.
func (in *DatadogAgent) DeepCopy() *DatadogAgent {
	if in == nil {
		return nil
	}
	out := new(DatadogAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatadogAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentGenericContainer) DeepCopyInto(out *DatadogAgentGenericContainer) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HealthPort != nil {
		in, out := &in.HealthPort, &out.HealthPort
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentGenericContainer.
func (in *DatadogAgentGenericContainer) DeepCopy() *DatadogAgentGenericContainer {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentGenericContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentList) DeepCopyInto(out *DatadogAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatadogAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentList.
func (in *DatadogAgentList) DeepCopy() *DatadogAgentList {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatadogAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentPodTemplateOverride) DeepCopyInto(out *DatadogAgentPodTemplateOverride) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]DatadogAgentGenericContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Kubelet != nil {
		in, out := &in.Kubelet, &out.Kubelet
		*out = new(KubeletConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentPodTemplateOverride.
func (in *DatadogAgentPodTemplateOverride) DeepCopy() *DatadogAgentPodTemplateOverride {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentPodTemplateOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentResourceOverride) DeepCopyInto(out *DatadogAgentResourceOverride) {
	*out = *in
	if in.DatadogAgentPodTemplateOverride != nil {
		in, out := &in.DatadogAgentPodTemplateOverride, &out.DatadogAgentPodTemplateOverride
		*out = new(DatadogAgentPodTemplateOverride)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentResourceOverride.
func (in *DatadogAgentResourceOverride) DeepCopy() *DatadogAgentResourceOverride {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentResourceOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentSpec) DeepCopyInto(out *DatadogAgentSpec) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(DatadogFeatures)
		(*in).DeepCopyInto(*out)
	}
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(GlobalConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = make(map[ResourceName]DatadogAgentResourceOverride, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentSpec.
func (in *DatadogAgentSpec) DeepCopy() *DatadogAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentStatus) DeepCopyInto(out *DatadogAgentStatus) {
	*out = *in
	if in.DefaultOverride != nil {
		in, out := &in.DefaultOverride, &out.DefaultOverride
		*out = new(DatadogAgentSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentStatus.
func (in *DatadogAgentStatus) DeepCopy() *DatadogAgentStatus {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogCredentials) DeepCopyInto(out *DatadogCredentials) {
	*out = *in
	if in.APISecret != nil {
		in, out := &in.APISecret, &out.APISecret
		*out = new(Secret)
		**out = **in
	}
	if in.AppSecret != nil {
		in, out := &in.AppSecret, &out.AppSecret
		*out = new(Secret)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogCredentials.
func (in *DatadogCredentials) DeepCopy() *DatadogCredentials {
	if in == nil {
		return nil
	}
	out := new(DatadogCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogFeatures) DeepCopyInto(out *DatadogFeatures) {
	*out = *in
	if in.LogCollection != nil {
		in, out := &in.LogCollection, &out.LogCollection
		*out = new(LogCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ProcessCollection != nil {
		in, out := &in.ProcessCollection, &out.ProcessCollection
		*out = new(ProcessCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerCollection != nil {
		in, out := &in.ContainerCollection, &out.ContainerCollection
		*out = new(ContainerCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.APM != nil {
		in, out := &in.APM, &out.APM
		*out = new(APMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CSPM != nil {
		in, out := &in.CSPM, &out.CSPM
		*out = new(CSPMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CWS != nil {
		in, out := &in.CWS, &out.CWS
		*out = new(CWSFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NPM != nil {
		in, out := &in.NPM, &out.NPM
		*out = new(NPMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.USM != nil {
		in, out := &in.USM, &out.USM
		*out = new(USMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OrchestratorExplorer != nil {
		in, out := &in.OrchestratorExplorer, &out.OrchestratorExplorer
		*out = new(OrchestratorExplorerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeStateMetricsCore != nil {
		in, out := &in.KubeStateMetricsCore, &out.KubeStateMetricsCore
		*out = new(KubeStateMetricsCoreFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AdmissionController != nil {
		in, out := &in.AdmissionController, &out.AdmissionController
		*out = new(AdmissionControllerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalMetricsServer != nil {
		in, out := &in.ExternalMetricsServer, &out.ExternalMetricsServer
		*out = new(ExternalMetricsServerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterChecksRunner != nil {
		in, out := &in.ClusterChecksRunner, &out.ClusterChecksRunner
		*out = new(ClusterChecksRunnerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusScrape != nil {
		in, out := &in.PrometheusScrape, &out.PrometheusScrape
		*out = new(PrometheusScrapeFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DatadogMonitor != nil {
		in, out := &in.DatadogMonitor, &out.DatadogMonitor
		*out = new(DatadogMonitorFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogFeatures.
func (in *DatadogFeatures) DeepCopy() *DatadogFeatures {
	if in == nil {
		return nil
	}
	out := new(DatadogFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogMonitorFeatureConfig) DeepCopyInto(out *DatadogMonitorFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogMonitorFeatureConfig.
func (in *DatadogMonitorFeatureConfig) DeepCopy() *DatadogMonitorFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(DatadogMonitorFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(DatadogCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricsServerFeatureConfig) DeepCopyInto(out *ExternalMetricsServerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricsServerFeatureConfig.
func (in *ExternalMetricsServerFeatureConfig) DeepCopy() *ExternalMetricsServerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricsServerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalConfig) DeepCopyInto(out *GlobalConfig) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(DatadogCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetworkPolicy != nil {
		in, out := &in.NetworkPolicy, &out.NetworkPolicy
		*out = new(NetworkPolicyConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalService != nil {
		in, out := &in.LocalService, &out.LocalService
		*out = new(LocalService)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalConfig.
func (in *GlobalConfig) DeepCopy() *GlobalConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortConfig) DeepCopyInto(out *HostPortConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortConfig.
func (in *HostPortConfig) DeepCopy() *HostPortConfig {
	if in == nil {
		return nil
	}
	out := new(HostPortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfig) DeepCopyInto(out *ImageConfig) {
	*out = *in
	if in.PullPolicy != nil {
		in, out := &in.PullPolicy, &out.PullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	if in.PullSecrets != nil {
		in, out := &in.PullSecrets, &out.PullSecrets
		*out = new([]corev1.LocalObjectReference)
		if **in != nil {
			in, out := *in, *out
			*out = make([]corev1.LocalObjectReference, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfig.
func (in *ImageConfig) DeepCopy() *ImageConfig {
	if in == nil {
		return nil
	}
	out := new(ImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStateMetricsCoreFeatureConfig) DeepCopyInto(out *KubeStateMetricsCoreFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStateMetricsCoreFeatureConfig.
func (in *KubeStateMetricsCoreFeatureConfig) DeepCopy() *KubeStateMetricsCoreFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(KubeStateMetricsCoreFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeletConfig) DeepCopyInto(out *KubeletConfig) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(corev1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
	if in.TLSVerify != nil {
		in, out := &in.TLSVerify, &out.TLSVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeletConfig.
func (in *KubeletConfig) DeepCopy() *KubeletConfig {
	if in == nil {
		return nil
	}
	out := new(KubeletConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalService) DeepCopyInto(out *LocalService) {
	*out = *in
	if in.ForceEnableLocalService != nil {
		in, out := &in.ForceEnableLocalService, &out.ForceEnableLocalService
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalService.
func (in *LocalService) DeepCopy() *LocalService {
	if in == nil {
		return nil
	}
	out := new(LocalService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogCollectionFeatureConfig) DeepCopyInto(out *LogCollectionFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ContainerCollectAll != nil {
		in, out := &in.ContainerCollectAll, &out.ContainerCollectAll
		*out = new(bool)
		**out = **in
	}
	if in.ContainerCollectUsingFiles != nil {
		in, out := &in.ContainerCollectUsingFiles, &out.ContainerCollectUsingFiles
		*out = new(bool)
		**out = **in
	}
	if in.ContainerLogsPath != nil {
		in, out := &in.ContainerLogsPath, &out.ContainerLogsPath
		*out = new(string)
		**out = **in
	}
	if in.PodLogsPath != nil {
		in, out := &in.PodLogsPath, &out.PodLogsPath
		*out = new(string)
		**out = **in
	}
	if in.ContainerSymlinksPath != nil {
		in, out := &in.ContainerSymlinksPath, &out.ContainerSymlinksPath
		*out = new(string)
		**out = **in
	}
	if in.TempStoragePath != nil {
		in, out := &in.TempStoragePath, &out.TempStoragePath
		*out = new(string)
		**out = **in
	}
	if in.OpenFilesLimit != nil {
		in, out := &in.OpenFilesLimit, &out.OpenFilesLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogCollectionFeatureConfig.
func (in *LogCollectionFeatureConfig) DeepCopy() *LogCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(LogCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NPMFeatureConfig) DeepCopyInto(out *NPMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NPMFeatureConfig.
func (in *NPMFeatureConfig) DeepCopy() *NPMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(NPMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyConfig) DeepCopyInto(out *NetworkPolicyConfig) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(bool)
		**out = **in
	}
	if in.DNSSelectorEndpoints != nil {
		in, out := &in.DNSSelectorEndpoints, &out.DNSSelectorEndpoints
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyConfig.
func (in *NetworkPolicyConfig) DeepCopy() *NetworkPolicyConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestratorExplorerFeatureConfig) DeepCopyInto(out *OrchestratorExplorerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ScrubContainers != nil {
		in, out := &in.ScrubContainers, &out.ScrubContainers
		*out = new(bool)
		**out = **in
	}
	if in.ExtraTags != nil {
		in, out := &in.ExtraTags, &out.ExtraTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestratorExplorerFeatureConfig.
func (in *OrchestratorExplorerFeatureConfig) DeepCopy() *OrchestratorExplorerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(OrchestratorExplorerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessCollectionFeatureConfig) DeepCopyInto(out *ProcessCollectionFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessCollectionFeatureConfig.
func (in *ProcessCollectionFeatureConfig) DeepCopy() *ProcessCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(ProcessCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusScrapeFeatureConfig) DeepCopyInto(out *PrometheusScrapeFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EnableServiceEndpoints != nil {
		in, out := &in.EnableServiceEndpoints, &out.EnableServiceEndpoints
		*out = new(bool)
		**out = **in
	}
	if in.AdditionalConfigs != nil {
		in, out := &in.AdditionalConfigs, &out.AdditionalConfigs
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusScrapeFeatureConfig.
func (in *PrometheusScrapeFeatureConfig) DeepCopy() *PrometheusScrapeFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(PrometheusScrapeFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *USMFeatureConfig) DeepCopyInto(out *USMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new USMFeatureConfig.
func (in *USMFeatureConfig) DeepCopy() *USMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(USMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnixDomainSocketConfig) DeepCopyInto(out *UnixDomainSocketConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnixDomainSocketConfig.
func (in *UnixDomainSocketConfig) DeepCopy() *UnixDomainSocketConfig {
	if in == nil {
		return nil
	}
	out := new(UnixDomainSocketConfig)
	in.DeepCopyInto(out)
	return out
}
