/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OutputSynapseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputSynapseParameters struct {

	// +kubebuilder:validation:Required
	Database *string `json:"database" tf:"database,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Server *string `json:"server" tf:"server,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/streamanalytics/v1beta1.Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`

	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// OutputSynapseSpec defines the desired state of OutputSynapse
type OutputSynapseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputSynapseParameters `json:"forProvider"`
}

// OutputSynapseStatus defines the observed state of OutputSynapse.
type OutputSynapseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputSynapseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputSynapse is the Schema for the OutputSynapses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputSynapse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputSynapseSpec   `json:"spec"`
	Status            OutputSynapseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputSynapseList contains a list of OutputSynapses
type OutputSynapseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputSynapse `json:"items"`
}

// Repository type metadata.
var (
	OutputSynapse_Kind             = "OutputSynapse"
	OutputSynapse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputSynapse_Kind}.String()
	OutputSynapse_KindAPIVersion   = OutputSynapse_Kind + "." + CRDGroupVersion.String()
	OutputSynapse_GroupVersionKind = CRDGroupVersion.WithKind(OutputSynapse_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputSynapse{}, &OutputSynapseList{})
}
