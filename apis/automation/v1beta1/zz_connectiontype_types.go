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

type ConnectionTypeObservation struct {

	// The the Automation Connection Type ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ConnectionTypeParameters struct {

	// The name of the automation account in which the Connection is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/automation/v1beta1.Account
	// +kubebuilder:validation:Optional
	AutomationAccountName *string `json:"automationAccountName,omitempty" tf:"automation_account_name,omitempty"`

	// Reference to a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameRef *v1.Reference `json:"automationAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in automation to populate automationAccountName.
	// +kubebuilder:validation:Optional
	AutomationAccountNameSelector *v1.Selector `json:"automationAccountNameSelector,omitempty" tf:"-"`

	// One or more field blocks as defined below. Changing this forces a new Automation to be created.
	// +kubebuilder:validation:Required
	Field []FieldParameters `json:"field" tf:"field,omitempty"`

	// Whether the connection type is global. Changing this forces a new Automation to be created.
	// +kubebuilder:validation:Optional
	IsGlobal *bool `json:"isGlobal,omitempty" tf:"is_global,omitempty"`

	// The name which should be used for this Automation Connection Type. Changing this forces a new Automation to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of the Resource Group where the Automation should exist. Changing this forces a new Automation to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

type FieldObservation struct {
}

type FieldParameters struct {

	// Whether to set the isEncrypted flag of the connection field definition.
	// +kubebuilder:validation:Optional
	IsEncrypted *bool `json:"isEncrypted,omitempty" tf:"is_encrypted,omitempty"`

	// Whether to set the isOptional flag of the connection field definition.
	// +kubebuilder:validation:Optional
	IsOptional *bool `json:"isOptional,omitempty" tf:"is_optional,omitempty"`

	// The name which should be used for this connection field definition.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The type of the connection field definition.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// ConnectionTypeSpec defines the desired state of ConnectionType
type ConnectionTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionTypeParameters `json:"forProvider"`
}

// ConnectionTypeStatus defines the observed state of ConnectionType.
type ConnectionTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionType is the Schema for the ConnectionTypes API. Manages an Automation Connection Type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ConnectionType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionTypeSpec   `json:"spec"`
	Status            ConnectionTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionTypeList contains a list of ConnectionTypes
type ConnectionTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionType `json:"items"`
}

// Repository type metadata.
var (
	ConnectionType_Kind             = "ConnectionType"
	ConnectionType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionType_Kind}.String()
	ConnectionType_KindAPIVersion   = ConnectionType_Kind + "." + CRDGroupVersion.String()
	ConnectionType_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionType_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionType{}, &ConnectionTypeList{})
}