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

type SharedImageGalleryInitParameters struct {

	// A description for this Shared Image Gallery.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the Shared Image Gallery.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SharedImageGalleryObservation struct {

	// A description for this Shared Image Gallery.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Shared Image Gallery.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the Shared Image Gallery.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Unique Name for this Shared Image Gallery.
	UniqueName *string `json:"uniqueName,omitempty" tf:"unique_name,omitempty"`
}

type SharedImageGalleryParameters struct {

	// A description for this Shared Image Gallery.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the Shared Image Gallery. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Shared Image Gallery.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SharedImageGallerySpec defines the desired state of SharedImageGallery
type SharedImageGallerySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedImageGalleryParameters `json:"forProvider"`
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
	InitProvider SharedImageGalleryInitParameters `json:"initProvider,omitempty"`
}

// SharedImageGalleryStatus defines the observed state of SharedImageGallery.
type SharedImageGalleryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedImageGalleryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SharedImageGallery is the Schema for the SharedImageGallerys API. Manages a Shared Image Gallery.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SharedImageGallery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   SharedImageGallerySpec   `json:"spec"`
	Status SharedImageGalleryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedImageGalleryList contains a list of SharedImageGallerys
type SharedImageGalleryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedImageGallery `json:"items"`
}

// Repository type metadata.
var (
	SharedImageGallery_Kind             = "SharedImageGallery"
	SharedImageGallery_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedImageGallery_Kind}.String()
	SharedImageGallery_KindAPIVersion   = SharedImageGallery_Kind + "." + CRDGroupVersion.String()
	SharedImageGallery_GroupVersionKind = CRDGroupVersion.WithKind(SharedImageGallery_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedImageGallery{}, &SharedImageGalleryList{})
}
