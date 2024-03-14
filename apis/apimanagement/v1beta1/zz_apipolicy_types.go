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

type APIPolicyInitParameters struct {

	// The XML Content for this Policy as a string.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type APIPolicyObservation struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// The ID of the API Management API Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The XML Content for this Policy as a string.
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

type APIPolicyParameters struct {

	// The name of the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Management
	// +kubebuilder:validation:Optional
	APIManagementName *string `json:"apiManagementName,omitempty" tf:"api_management_name,omitempty"`

	// Reference to a Management to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameRef *v1.Reference `json:"apiManagementNameRef,omitempty" tf:"-"`

	// Selector for a Management to populate apiManagementName.
	// +kubebuilder:validation:Optional
	APIManagementNameSelector *v1.Selector `json:"apiManagementNameSelector,omitempty" tf:"-"`

	// The ID of the API Management API within the API Management Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=API
	// +kubebuilder:validation:Optional
	APIName *string `json:"apiName,omitempty" tf:"api_name,omitempty"`

	// Reference to a API to populate apiName.
	// +kubebuilder:validation:Optional
	APINameRef *v1.Reference `json:"apiNameRef,omitempty" tf:"-"`

	// Selector for a API to populate apiName.
	// +kubebuilder:validation:Optional
	APINameSelector *v1.Selector `json:"apiNameSelector,omitempty" tf:"-"`

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

	// The XML Content for this Policy as a string.
	// +kubebuilder:validation:Optional
	XMLContent *string `json:"xmlContent,omitempty" tf:"xml_content,omitempty"`

	// A link to a Policy XML Document, which must be publicly available.
	// +kubebuilder:validation:Optional
	XMLLink *string `json:"xmlLink,omitempty" tf:"xml_link,omitempty"`
}

// APIPolicySpec defines the desired state of APIPolicy
type APIPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIPolicyParameters `json:"forProvider"`
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
	InitProvider APIPolicyInitParameters `json:"initProvider,omitempty"`
}

// APIPolicyStatus defines the observed state of APIPolicy.
type APIPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// APIPolicy is the Schema for the APIPolicys API. Manages an API Management API Policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type APIPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APIPolicySpec   `json:"spec"`
	Status            APIPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIPolicyList contains a list of APIPolicys
type APIPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIPolicy `json:"items"`
}

// Repository type metadata.
var (
	APIPolicy_Kind             = "APIPolicy"
	APIPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: APIPolicy_Kind}.String()
	APIPolicy_KindAPIVersion   = APIPolicy_Kind + "." + CRDGroupVersion.String()
	APIPolicy_GroupVersionKind = CRDGroupVersion.WithKind(APIPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&APIPolicy{}, &APIPolicyList{})
}
