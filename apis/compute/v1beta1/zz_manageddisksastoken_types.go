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

type ManagedDiskSASTokenInitParameters struct {

	// The level of access required on the disk. Supported are Read, Write. Changing this forces a new resource to be created.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// The duration for which the export should be allowed. Should be between 30 & 4294967295 seconds. Changing this forces a new resource to be created.
	DurationInSeconds *float64 `json:"durationInSeconds,omitempty" tf:"duration_in_seconds,omitempty"`

	// The ID of an existing Managed Disk which should be exported. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Reference to a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDRef *v1.Reference `json:"managedDiskIdRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDSelector *v1.Selector `json:"managedDiskIdSelector,omitempty" tf:"-"`
}

type ManagedDiskSASTokenObservation struct {

	// The level of access required on the disk. Supported are Read, Write. Changing this forces a new resource to be created.
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// The duration for which the export should be allowed. Should be between 30 & 4294967295 seconds. Changing this forces a new resource to be created.
	DurationInSeconds *float64 `json:"durationInSeconds,omitempty" tf:"duration_in_seconds,omitempty"`

	// The ID of the Disk Export resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of an existing Managed Disk which should be exported. Changing this forces a new resource to be created.
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`
}

type ManagedDiskSASTokenParameters struct {

	// The level of access required on the disk. Supported are Read, Write. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level,omitempty"`

	// The duration for which the export should be allowed. Should be between 30 & 4294967295 seconds. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DurationInSeconds *float64 `json:"durationInSeconds,omitempty" tf:"duration_in_seconds,omitempty"`

	// The ID of an existing Managed Disk which should be exported. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.ManagedDisk
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedDiskID *string `json:"managedDiskId,omitempty" tf:"managed_disk_id,omitempty"`

	// Reference to a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDRef *v1.Reference `json:"managedDiskIdRef,omitempty" tf:"-"`

	// Selector for a ManagedDisk in compute to populate managedDiskId.
	// +kubebuilder:validation:Optional
	ManagedDiskIDSelector *v1.Selector `json:"managedDiskIdSelector,omitempty" tf:"-"`
}

// ManagedDiskSASTokenSpec defines the desired state of ManagedDiskSASToken
type ManagedDiskSASTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedDiskSASTokenParameters `json:"forProvider"`
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
	InitProvider ManagedDiskSASTokenInitParameters `json:"initProvider,omitempty"`
}

// ManagedDiskSASTokenStatus defines the observed state of ManagedDiskSASToken.
type ManagedDiskSASTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedDiskSASTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ManagedDiskSASToken is the Schema for the ManagedDiskSASTokens API. Manages a Disk SAS Token.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedDiskSASToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accessLevel) || (has(self.initProvider) && has(self.initProvider.accessLevel))",message="spec.forProvider.accessLevel is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.durationInSeconds) || (has(self.initProvider) && has(self.initProvider.durationInSeconds))",message="spec.forProvider.durationInSeconds is a required parameter"
	Spec   ManagedDiskSASTokenSpec   `json:"spec"`
	Status ManagedDiskSASTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedDiskSASTokenList contains a list of ManagedDiskSASTokens
type ManagedDiskSASTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedDiskSASToken `json:"items"`
}

// Repository type metadata.
var (
	ManagedDiskSASToken_Kind             = "ManagedDiskSASToken"
	ManagedDiskSASToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedDiskSASToken_Kind}.String()
	ManagedDiskSASToken_KindAPIVersion   = ManagedDiskSASToken_Kind + "." + CRDGroupVersion.String()
	ManagedDiskSASToken_GroupVersionKind = CRDGroupVersion.WithKind(ManagedDiskSASToken_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedDiskSASToken{}, &ManagedDiskSASTokenList{})
}
