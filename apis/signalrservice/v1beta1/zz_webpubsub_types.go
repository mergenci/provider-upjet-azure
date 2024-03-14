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

type WebPubsubIdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Web PubSub.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Web PubSub. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WebPubsubIdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Web PubSub.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Web PubSub. Possible values are SystemAssigned, UserAssigned.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WebPubsubIdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Web PubSub.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Web PubSub. Possible values are SystemAssigned, UserAssigned.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type WebPubsubInitParameters struct {

	// Whether to enable AAD auth? Defaults to true.
	AADAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled,omitempty"`

	// Specifies the number of units associated with this Web PubSub resource. Valid values are: Free: 1, Standard: 1, 2, 5, 10, 20, 50, 100.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// An identity block as defined below.
	Identity []WebPubsubIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// A live_trace block as defined below.
	LiveTrace []WebPubsubLiveTraceInitParameters `json:"liveTrace,omitempty" tf:"live_trace,omitempty"`

	// Whether to enable local auth? Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the Web PubSub service exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Web PubSub service. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to enable public network access? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Web PubSub service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies which SKU to use. Possible values are Free_F1, Standard_S1, and Premium_P1.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Whether to request client certificate during TLS handshake? Defaults to false.
	TLSClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type WebPubsubLiveTraceInitParameters struct {

	// Whether the log category ConnectivityLogs is enabled? Defaults to true
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// Whether the live trace is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the log category HttpRequestLogs is enabled? Defaults to true
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// Whether the log category MessagingLogs is enabled? Defaults to true
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`
}

type WebPubsubLiveTraceObservation struct {

	// Whether the log category ConnectivityLogs is enabled? Defaults to true
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// Whether the live trace is enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the log category HttpRequestLogs is enabled? Defaults to true
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// Whether the log category MessagingLogs is enabled? Defaults to true
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`
}

type WebPubsubLiveTraceParameters struct {

	// Whether the log category ConnectivityLogs is enabled? Defaults to true
	// +kubebuilder:validation:Optional
	ConnectivityLogsEnabled *bool `json:"connectivityLogsEnabled,omitempty" tf:"connectivity_logs_enabled,omitempty"`

	// Whether the live trace is enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Whether the log category HttpRequestLogs is enabled? Defaults to true
	// +kubebuilder:validation:Optional
	HTTPRequestLogsEnabled *bool `json:"httpRequestLogsEnabled,omitempty" tf:"http_request_logs_enabled,omitempty"`

	// Whether the log category MessagingLogs is enabled? Defaults to true
	// +kubebuilder:validation:Optional
	MessagingLogsEnabled *bool `json:"messagingLogsEnabled,omitempty" tf:"messaging_logs_enabled,omitempty"`
}

type WebPubsubObservation struct {

	// Whether to enable AAD auth? Defaults to true.
	AADAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled,omitempty"`

	// Specifies the number of units associated with this Web PubSub resource. Valid values are: Free: 1, Standard: 1, 2, 5, 10, 20, 50, 100.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The publicly accessible IP of the Web PubSub service.
	ExternalIP *string `json:"externalIp,omitempty" tf:"external_ip,omitempty"`

	// The FQDN of the Web PubSub service.
	HostName *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The ID of the Web PubSub service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []WebPubsubIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// A live_trace block as defined below.
	LiveTrace []WebPubsubLiveTraceObservation `json:"liveTrace,omitempty" tf:"live_trace,omitempty"`

	// Whether to enable local auth? Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the Web PubSub service exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Web PubSub service. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to enable public network access? Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The publicly accessible port of the Web PubSub service which is designed for browser/client use.
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`

	// The name of the resource group in which to create the Web PubSub service. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The publicly accessible port of the Web PubSub service which is designed for customer server side use.
	ServerPort *float64 `json:"serverPort,omitempty" tf:"server_port,omitempty"`

	// Specifies which SKU to use. Possible values are Free_F1, Standard_S1, and Premium_P1.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Whether to request client certificate during TLS handshake? Defaults to false.
	TLSClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type WebPubsubParameters struct {

	// Whether to enable AAD auth? Defaults to true.
	// +kubebuilder:validation:Optional
	AADAuthEnabled *bool `json:"aadAuthEnabled,omitempty" tf:"aad_auth_enabled,omitempty"`

	// Specifies the number of units associated with this Web PubSub resource. Valid values are: Free: 1, Standard: 1, 2, 5, 10, 20, 50, 100.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []WebPubsubIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// A live_trace block as defined below.
	// +kubebuilder:validation:Optional
	LiveTrace []WebPubsubLiveTraceParameters `json:"liveTrace,omitempty" tf:"live_trace,omitempty"`

	// Whether to enable local auth? Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the Web PubSub service exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Web PubSub service. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to enable public network access? Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// The name of the resource group in which to create the Web PubSub service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies which SKU to use. Possible values are Free_F1, Standard_S1, and Premium_P1.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// Whether to request client certificate during TLS handshake? Defaults to false.
	// +kubebuilder:validation:Optional
	TLSClientCertEnabled *bool `json:"tlsClientCertEnabled,omitempty" tf:"tls_client_cert_enabled,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// WebPubsubSpec defines the desired state of WebPubsub
type WebPubsubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WebPubsubParameters `json:"forProvider"`
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
	InitProvider WebPubsubInitParameters `json:"initProvider,omitempty"`
}

// WebPubsubStatus defines the observed state of WebPubsub.
type WebPubsubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WebPubsubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// WebPubsub is the Schema for the WebPubsubs API. Manages an Azure Web PubSub service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WebPubsub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   WebPubsubSpec   `json:"spec"`
	Status WebPubsubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebPubsubList contains a list of WebPubsubs
type WebPubsubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebPubsub `json:"items"`
}

// Repository type metadata.
var (
	WebPubsub_Kind             = "WebPubsub"
	WebPubsub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WebPubsub_Kind}.String()
	WebPubsub_KindAPIVersion   = WebPubsub_Kind + "." + CRDGroupVersion.String()
	WebPubsub_GroupVersionKind = CRDGroupVersion.WithKind(WebPubsub_Kind)
)

func init() {
	SchemeBuilder.Register(&WebPubsub{}, &WebPubsubList{})
}
