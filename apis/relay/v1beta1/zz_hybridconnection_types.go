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

type HybridConnectionInitParameters struct {

	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.EventRelayNamespace
	RelayNamespaceName *string `json:"relayNamespaceName,omitempty" tf:"relay_namespace_name,omitempty"`

	// Reference to a EventRelayNamespace in relay to populate relayNamespaceName.
	// +kubebuilder:validation:Optional
	RelayNamespaceNameRef *v1.Reference `json:"relayNamespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventRelayNamespace in relay to populate relayNamespaceName.
	// +kubebuilder:validation:Optional
	RelayNamespaceNameSelector *v1.Selector `json:"relayNamespaceNameSelector,omitempty" tf:"-"`

	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created. Defaults to true.
	RequiresClientAuthorization *bool `json:"requiresClientAuthorization,omitempty" tf:"requires_client_authorization,omitempty"`

	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

type HybridConnectionObservation struct {

	// The ID of the Relay Hybrid Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	RelayNamespaceName *string `json:"relayNamespaceName,omitempty" tf:"relay_namespace_name,omitempty"`

	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created. Defaults to true.
	RequiresClientAuthorization *bool `json:"requiresClientAuthorization,omitempty" tf:"requires_client_authorization,omitempty"`

	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

type HybridConnectionParameters struct {

	// Specifies the name of the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the Azure Relay in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/relay/v1beta1.EventRelayNamespace
	// +kubebuilder:validation:Optional
	RelayNamespaceName *string `json:"relayNamespaceName,omitempty" tf:"relay_namespace_name,omitempty"`

	// Reference to a EventRelayNamespace in relay to populate relayNamespaceName.
	// +kubebuilder:validation:Optional
	RelayNamespaceNameRef *v1.Reference `json:"relayNamespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventRelayNamespace in relay to populate relayNamespaceName.
	// +kubebuilder:validation:Optional
	RelayNamespaceNameSelector *v1.Selector `json:"relayNamespaceNameSelector,omitempty" tf:"-"`

	// Specify if client authorization is needed for this hybrid connection. True by default. Changing this forces a new resource to be created. Defaults to true.
	// +kubebuilder:validation:Optional
	RequiresClientAuthorization *bool `json:"requiresClientAuthorization,omitempty" tf:"requires_client_authorization,omitempty"`

	// The name of the resource group in which to create the Azure Relay Hybrid Connection. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The usermetadata is a placeholder to store user-defined string data for the hybrid connection endpoint. For example, it can be used to store descriptive data, such as a list of teams and their contact information. Also, user-defined configuration settings can be stored.
	// +kubebuilder:validation:Optional
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

// HybridConnectionSpec defines the desired state of HybridConnection
type HybridConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HybridConnectionParameters `json:"forProvider"`
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
	InitProvider HybridConnectionInitParameters `json:"initProvider,omitempty"`
}

// HybridConnectionStatus defines the observed state of HybridConnection.
type HybridConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HybridConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HybridConnection is the Schema for the HybridConnections API. Manages an Azure Relay Hybrid Connection.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HybridConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HybridConnectionSpec   `json:"spec"`
	Status HybridConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HybridConnectionList contains a list of HybridConnections
type HybridConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HybridConnection `json:"items"`
}

// Repository type metadata.
var (
	HybridConnection_Kind             = "HybridConnection"
	HybridConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HybridConnection_Kind}.String()
	HybridConnection_KindAPIVersion   = HybridConnection_Kind + "." + CRDGroupVersion.String()
	HybridConnection_GroupVersionKind = CRDGroupVersion.WithKind(HybridConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&HybridConnection{}, &HybridConnectionList{})
}
