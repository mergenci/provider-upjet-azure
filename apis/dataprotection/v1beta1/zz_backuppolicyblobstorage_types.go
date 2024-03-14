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

type BackupPolicyBlobStorageInitParameters struct {

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`
}

type BackupPolicyBlobStorageObservation struct {

	// The ID of the Backup Policy Blob Storage.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type BackupPolicyBlobStorageParameters struct {

	// Duration of deletion after given timespan. It should follow ISO 8601 duration format. Changing this forces a new Backup Policy Blob Storage to be created.
	// +kubebuilder:validation:Optional
	RetentionDuration *string `json:"retentionDuration,omitempty" tf:"retention_duration,omitempty"`

	// The ID of the Backup Vault within which the Backup Policy Blob Storage should exist. Changing this forces a new Backup Policy Blob Storage to be created.
	// +crossplane:generate:reference:type=BackupVault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`

	// Reference to a BackupVault to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDRef *v1.Reference `json:"vaultIdRef,omitempty" tf:"-"`

	// Selector for a BackupVault to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDSelector *v1.Selector `json:"vaultIdSelector,omitempty" tf:"-"`
}

// BackupPolicyBlobStorageSpec defines the desired state of BackupPolicyBlobStorage
type BackupPolicyBlobStorageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyBlobStorageParameters `json:"forProvider"`
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
	InitProvider BackupPolicyBlobStorageInitParameters `json:"initProvider,omitempty"`
}

// BackupPolicyBlobStorageStatus defines the observed state of BackupPolicyBlobStorage.
type BackupPolicyBlobStorageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyBlobStorageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackupPolicyBlobStorage is the Schema for the BackupPolicyBlobStorages API. Manages a Backup Policy Blob Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyBlobStorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionDuration) || (has(self.initProvider) && has(self.initProvider.retentionDuration))",message="spec.forProvider.retentionDuration is a required parameter"
	Spec   BackupPolicyBlobStorageSpec   `json:"spec"`
	Status BackupPolicyBlobStorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyBlobStorageList contains a list of BackupPolicyBlobStorages
type BackupPolicyBlobStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyBlobStorage `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyBlobStorage_Kind             = "BackupPolicyBlobStorage"
	BackupPolicyBlobStorage_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyBlobStorage_Kind}.String()
	BackupPolicyBlobStorage_KindAPIVersion   = BackupPolicyBlobStorage_Kind + "." + CRDGroupVersion.String()
	BackupPolicyBlobStorage_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyBlobStorage_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyBlobStorage{}, &BackupPolicyBlobStorageList{})
}
