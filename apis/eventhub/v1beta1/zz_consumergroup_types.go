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

type ConsumerGroupInitParameters struct {

	// Specifies the user metadata.
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

type ConsumerGroupObservation struct {

	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// The ID of the EventHub Consumer Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the user metadata.
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

type ConsumerGroupParameters struct {

	// Specifies the name of the EventHub. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=EventHub
	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// Reference to a EventHub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameRef *v1.Reference `json:"eventhubNameRef,omitempty" tf:"-"`

	// Selector for a EventHub to populate eventhubName.
	// +kubebuilder:validation:Optional
	EventHubNameSelector *v1.Selector `json:"eventhubNameSelector,omitempty" tf:"-"`

	// Specifies the name of the grandparent EventHub Namespace. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=EventHubNamespace
	// +kubebuilder:validation:Optional
	NamespaceName *string `json:"namespaceName,omitempty" tf:"namespace_name,omitempty"`

	// Reference to a EventHubNamespace to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameRef *v1.Reference `json:"namespaceNameRef,omitempty" tf:"-"`

	// Selector for a EventHubNamespace to populate namespaceName.
	// +kubebuilder:validation:Optional
	NamespaceNameSelector *v1.Selector `json:"namespaceNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which the EventHub Consumer Group's grandparent Namespace exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the user metadata.
	// +kubebuilder:validation:Optional
	UserMetadata *string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

// ConsumerGroupSpec defines the desired state of ConsumerGroup
type ConsumerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConsumerGroupParameters `json:"forProvider"`
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
	InitProvider ConsumerGroupInitParameters `json:"initProvider,omitempty"`
}

// ConsumerGroupStatus defines the observed state of ConsumerGroup.
type ConsumerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConsumerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ConsumerGroup is the Schema for the ConsumerGroups API. Manages a Event Hubs Consumer Group as a nested resource within an Event Hub.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConsumerGroupSpec   `json:"spec"`
	Status            ConsumerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroupList contains a list of ConsumerGroups
type ConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsumerGroup `json:"items"`
}

// Repository type metadata.
var (
	ConsumerGroup_Kind             = "ConsumerGroup"
	ConsumerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConsumerGroup_Kind}.String()
	ConsumerGroup_KindAPIVersion   = ConsumerGroup_Kind + "." + CRDGroupVersion.String()
	ConsumerGroup_GroupVersionKind = CRDGroupVersion.WithKind(ConsumerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}
