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

type GatewayInitParameters_2 struct {

	// The ID of the API Management Resource in which the gateway will be created. Changing this forces a new API Management Gateway resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDRef *v1.Reference `json:"apiManagementIdRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDSelector *v1.Selector `json:"apiManagementIdSelector,omitempty" tf:"-"`

	// The description of the API Management Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A location_data block as documented below.
	LocationData []LocationDataInitParameters `json:"locationData,omitempty" tf:"location_data,omitempty"`
}

type GatewayObservation_2 struct {

	// The ID of the API Management Resource in which the gateway will be created. Changing this forces a new API Management Gateway resource to be created.
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// The description of the API Management Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the API Management Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A location_data block as documented below.
	LocationData []LocationDataObservation `json:"locationData,omitempty" tf:"location_data,omitempty"`
}

type GatewayParameters_2 struct {

	// The ID of the API Management Resource in which the gateway will be created. Changing this forces a new API Management Gateway resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDRef *v1.Reference `json:"apiManagementIdRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDSelector *v1.Selector `json:"apiManagementIdSelector,omitempty" tf:"-"`

	// The description of the API Management Gateway.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A location_data block as documented below.
	// +kubebuilder:validation:Optional
	LocationData []LocationDataParameters `json:"locationData,omitempty" tf:"location_data,omitempty"`
}

type LocationDataInitParameters struct {

	// The city or locality where the resource is located.
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// The district, state, or province where the resource is located.
	District *string `json:"district,omitempty" tf:"district,omitempty"`

	// A canonical name for the geographic or physical location.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The country or region where the resource is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type LocationDataObservation struct {

	// The city or locality where the resource is located.
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// The district, state, or province where the resource is located.
	District *string `json:"district,omitempty" tf:"district,omitempty"`

	// A canonical name for the geographic or physical location.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The country or region where the resource is located.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type LocationDataParameters struct {

	// The city or locality where the resource is located.
	// +kubebuilder:validation:Optional
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// The district, state, or province where the resource is located.
	// +kubebuilder:validation:Optional
	District *string `json:"district,omitempty" tf:"district,omitempty"`

	// A canonical name for the geographic or physical location.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The country or region where the resource is located.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters_2 `json:"forProvider"`
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
	InitProvider GatewayInitParameters_2 `json:"initProvider,omitempty"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Gateway is the Schema for the Gateways API. Manages an API Management Gateway.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.locationData) || (has(self.initProvider) && has(self.initProvider.locationData))",message="spec.forProvider.locationData is a required parameter"
	Spec   GatewaySpec   `json:"spec"`
	Status GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
