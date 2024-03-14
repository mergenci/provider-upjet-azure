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

type ApplicationInsightsInitParameters struct {
}

type ApplicationInsightsObservation struct {
}

type ApplicationInsightsParameters struct {

	// The instrumentation key used to push data to Application Insights.
	// +kubebuilder:validation:Required
	InstrumentationKeySecretRef v1.SecretKeySelector `json:"instrumentationKeySecretRef" tf:"-"`
}

type EventHubInitParameters struct {

	// The name of an EventHub.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EventHubObservation struct {

	// The name of an EventHub.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EventHubParameters struct {

	// The connection string of an EventHub Namespace.
	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// The name of an EventHub.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type LoggerInitParameters struct {

	// An application_insights block as documented below. Changing this forces a new resource to be created.
	ApplicationInsights []ApplicationInsightsInitParameters `json:"applicationInsights,omitempty" tf:"application_insights,omitempty"`

	// Specifies whether records should be buffered in the Logger prior to publishing. Defaults to true.
	Buffered *bool `json:"buffered,omitempty" tf:"buffered,omitempty"`

	// A description of this Logger.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An eventhub block as documented below. Changing this forces a new resource to be created.
	EventHub []EventHubInitParameters `json:"eventhub,omitempty" tf:"eventhub,omitempty"`

	// The target resource id which will be linked in the API-Management portal page. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

type LoggerObservation struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// An application_insights block as documented below. Changing this forces a new resource to be created.
	ApplicationInsights []ApplicationInsightsParameters `json:"applicationInsights,omitempty" tf:"application_insights,omitempty"`

	// Specifies whether records should be buffered in the Logger prior to publishing. Defaults to true.
	Buffered *bool `json:"buffered,omitempty" tf:"buffered,omitempty"`

	// A description of this Logger.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An eventhub block as documented below. Changing this forces a new resource to be created.
	EventHub []EventHubObservation `json:"eventhub,omitempty" tf:"eventhub,omitempty"`

	// The ID of the API Management Logger.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The target resource id which will be linked in the API-Management portal page. Changing this forces a new resource to be created.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`
}

type LoggerParameters struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// An application_insights block as documented below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ApplicationInsights []ApplicationInsightsParameters `json:"applicationInsights,omitempty" tf:"application_insights,omitempty"`

	// Specifies whether records should be buffered in the Logger prior to publishing. Defaults to true.
	// +kubebuilder:validation:Optional
	Buffered *bool `json:"buffered,omitempty" tf:"buffered,omitempty"`

	// A description of this Logger.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An eventhub block as documented below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	EventHub []EventHubParameters `json:"eventhub,omitempty" tf:"eventhub,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The target resource id which will be linked in the API-Management portal page. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`
}

// LoggerSpec defines the desired state of Logger
type LoggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoggerParameters `json:"forProvider"`
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
	InitProvider LoggerInitParameters `json:"initProvider,omitempty"`
}

// LoggerStatus defines the observed state of Logger.
type LoggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Logger is the Schema for the Loggers API. Manages a Logger within an API Management Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Logger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggerSpec   `json:"spec"`
	Status            LoggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoggerList contains a list of Loggers
type LoggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Logger `json:"items"`
}

// Repository type metadata.
var (
	Logger_Kind             = "Logger"
	Logger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Logger_Kind}.String()
	Logger_KindAPIVersion   = Logger_Kind + "." + CRDGroupVersion.String()
	Logger_GroupVersionKind = CRDGroupVersion.WithKind(Logger_Kind)
)

func init() {
	SchemeBuilder.Register(&Logger{}, &LoggerList{})
}
