// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TagRuleInitParameters struct {

	// Whether AAD logs should be sent to the Monitor resource?
	SendAADLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs,omitempty"`

	// Whether activity logs from Azure resources should be sent to the Monitor resource?
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Whether subscription logs should be sent to the Monitor resource?
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`

	// One or more (up to 10) tag_filter blocks as defined below.
	TagFilter []TagRuleTagFilterInitParameters `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type TagRuleObservation struct {

	// The ID of the logz Tag Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Logz Monitor. Changing this forces a new logz Tag Rule to be created.
	LogzMonitorID *string `json:"logzMonitorId,omitempty" tf:"logz_monitor_id,omitempty"`

	// Whether AAD logs should be sent to the Monitor resource?
	SendAADLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs,omitempty"`

	// Whether activity logs from Azure resources should be sent to the Monitor resource?
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Whether subscription logs should be sent to the Monitor resource?
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`

	// One or more (up to 10) tag_filter blocks as defined below.
	TagFilter []TagRuleTagFilterObservation `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type TagRuleParameters struct {

	// The ID of the Logz Monitor. Changing this forces a new logz Tag Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/logz/v1beta1.Monitor
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LogzMonitorID *string `json:"logzMonitorId,omitempty" tf:"logz_monitor_id,omitempty"`

	// Reference to a Monitor in logz to populate logzMonitorId.
	// +kubebuilder:validation:Optional
	LogzMonitorIDRef *v1.Reference `json:"logzMonitorIdRef,omitempty" tf:"-"`

	// Selector for a Monitor in logz to populate logzMonitorId.
	// +kubebuilder:validation:Optional
	LogzMonitorIDSelector *v1.Selector `json:"logzMonitorIdSelector,omitempty" tf:"-"`

	// Whether AAD logs should be sent to the Monitor resource?
	// +kubebuilder:validation:Optional
	SendAADLogs *bool `json:"sendAadLogs,omitempty" tf:"send_aad_logs,omitempty"`

	// Whether activity logs from Azure resources should be sent to the Monitor resource?
	// +kubebuilder:validation:Optional
	SendActivityLogs *bool `json:"sendActivityLogs,omitempty" tf:"send_activity_logs,omitempty"`

	// Whether subscription logs should be sent to the Monitor resource?
	// +kubebuilder:validation:Optional
	SendSubscriptionLogs *bool `json:"sendSubscriptionLogs,omitempty" tf:"send_subscription_logs,omitempty"`

	// One or more (up to 10) tag_filter blocks as defined below.
	// +kubebuilder:validation:Optional
	TagFilter []TagRuleTagFilterParameters `json:"tagFilter,omitempty" tf:"tag_filter,omitempty"`
}

type TagRuleTagFilterInitParameters struct {

	// The action for a filtering tag. Possible values are Include and Exclude is allowed. Note that the Exclude takes priority over the Include.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name of this tag_filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of this tag_filter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagRuleTagFilterObservation struct {

	// The action for a filtering tag. Possible values are Include and Exclude is allowed. Note that the Exclude takes priority over the Include.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name of this tag_filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of this tag_filter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagRuleTagFilterParameters struct {

	// The action for a filtering tag. Possible values are Include and Exclude is allowed. Note that the Exclude takes priority over the Include.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// The name of this tag_filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of this tag_filter.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// TagRuleSpec defines the desired state of TagRule
type TagRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagRuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TagRuleInitParameters `json:"initProvider,omitempty"`
}

// TagRuleStatus defines the observed state of TagRule.
type TagRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TagRule is the Schema for the TagRules API. Manages a logz Tag Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type TagRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagRuleSpec   `json:"spec"`
	Status            TagRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagRuleList contains a list of TagRules
type TagRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagRule `json:"items"`
}

// Repository type metadata.
var (
	TagRule_Kind             = "TagRule"
	TagRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TagRule_Kind}.String()
	TagRule_KindAPIVersion   = TagRule_Kind + "." + CRDGroupVersion.String()
	TagRule_GroupVersionKind = CRDGroupVersion.WithKind(TagRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TagRule{}, &TagRuleList{})
}
