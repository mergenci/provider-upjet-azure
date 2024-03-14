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

type BackupInstancePostgreSQLInitParameters struct {

	// The ID of the Backup Policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dataprotection/v1beta1.BackupPolicyPostgreSQL
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// Reference to a BackupPolicyPostgreSQL in dataprotection to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDRef *v1.Reference `json:"backupPolicyIdRef,omitempty" tf:"-"`

	// Selector for a BackupPolicyPostgreSQL in dataprotection to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDSelector *v1.Selector `json:"backupPolicyIdSelector,omitempty" tf:"-"`

	// The ID or versionless ID of the key vault secret which stores the connection string of the database.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_id",true)
	DatabaseCredentialKeyVaultSecretID *string `json:"databaseCredentialKeyVaultSecretId,omitempty" tf:"database_credential_key_vault_secret_id,omitempty"`

	// Reference to a Secret in keyvault to populate databaseCredentialKeyVaultSecretId.
	// +kubebuilder:validation:Optional
	DatabaseCredentialKeyVaultSecretIDRef *v1.Reference `json:"databaseCredentialKeyVaultSecretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in keyvault to populate databaseCredentialKeyVaultSecretId.
	// +kubebuilder:validation:Optional
	DatabaseCredentialKeyVaultSecretIDSelector *v1.Selector `json:"databaseCredentialKeyVaultSecretIdSelector,omitempty" tf:"-"`

	// The ID of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbforpostgresql/v1beta1.Database
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// Reference to a Database in dbforpostgresql to populate databaseId.
	// +kubebuilder:validation:Optional
	DatabaseIDRef *v1.Reference `json:"databaseIdRef,omitempty" tf:"-"`

	// Selector for a Database in dbforpostgresql to populate databaseId.
	// +kubebuilder:validation:Optional
	DatabaseIDSelector *v1.Selector `json:"databaseIdSelector,omitempty" tf:"-"`

	// The location of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type BackupInstancePostgreSQLObservation struct {

	// The ID of the Backup Policy.
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// The ID or versionless ID of the key vault secret which stores the connection string of the database.
	DatabaseCredentialKeyVaultSecretID *string `json:"databaseCredentialKeyVaultSecretId,omitempty" tf:"database_credential_key_vault_secret_id,omitempty"`

	// The ID of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// The ID of the Backup Instance PostgreSQL.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Backup Vault within which the PostgreSQL Backup Instance should exist. Changing this forces a new Backup Instance PostgreSQL to be created.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type BackupInstancePostgreSQLParameters struct {

	// The ID of the Backup Policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dataprotection/v1beta1.BackupPolicyPostgreSQL
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// Reference to a BackupPolicyPostgreSQL in dataprotection to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDRef *v1.Reference `json:"backupPolicyIdRef,omitempty" tf:"-"`

	// Selector for a BackupPolicyPostgreSQL in dataprotection to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDSelector *v1.Selector `json:"backupPolicyIdSelector,omitempty" tf:"-"`

	// The ID or versionless ID of the key vault secret which stores the connection string of the database.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_id",true)
	// +kubebuilder:validation:Optional
	DatabaseCredentialKeyVaultSecretID *string `json:"databaseCredentialKeyVaultSecretId,omitempty" tf:"database_credential_key_vault_secret_id,omitempty"`

	// Reference to a Secret in keyvault to populate databaseCredentialKeyVaultSecretId.
	// +kubebuilder:validation:Optional
	DatabaseCredentialKeyVaultSecretIDRef *v1.Reference `json:"databaseCredentialKeyVaultSecretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in keyvault to populate databaseCredentialKeyVaultSecretId.
	// +kubebuilder:validation:Optional
	DatabaseCredentialKeyVaultSecretIDSelector *v1.Selector `json:"databaseCredentialKeyVaultSecretIdSelector,omitempty" tf:"-"`

	// The ID of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dbforpostgresql/v1beta1.Database
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// Reference to a Database in dbforpostgresql to populate databaseId.
	// +kubebuilder:validation:Optional
	DatabaseIDRef *v1.Reference `json:"databaseIdRef,omitempty" tf:"-"`

	// Selector for a Database in dbforpostgresql to populate databaseId.
	// +kubebuilder:validation:Optional
	DatabaseIDSelector *v1.Selector `json:"databaseIdSelector,omitempty" tf:"-"`

	// The location of the source database. Changing this forces a new Backup Instance PostgreSQL to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Backup Vault within which the PostgreSQL Backup Instance should exist. Changing this forces a new Backup Instance PostgreSQL to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/dataprotection/v1beta1.BackupVault
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`

	// Reference to a BackupVault in dataprotection to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDRef *v1.Reference `json:"vaultIdRef,omitempty" tf:"-"`

	// Selector for a BackupVault in dataprotection to populate vaultId.
	// +kubebuilder:validation:Optional
	VaultIDSelector *v1.Selector `json:"vaultIdSelector,omitempty" tf:"-"`
}

// BackupInstancePostgreSQLSpec defines the desired state of BackupInstancePostgreSQL
type BackupInstancePostgreSQLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupInstancePostgreSQLParameters `json:"forProvider"`
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
	InitProvider BackupInstancePostgreSQLInitParameters `json:"initProvider,omitempty"`
}

// BackupInstancePostgreSQLStatus defines the observed state of BackupInstancePostgreSQL.
type BackupInstancePostgreSQLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupInstancePostgreSQLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackupInstancePostgreSQL is the Schema for the BackupInstancePostgreSQLs API. Manages a Backup Instance to back up PostgreSQL.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupInstancePostgreSQL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   BackupInstancePostgreSQLSpec   `json:"spec"`
	Status BackupInstancePostgreSQLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupInstancePostgreSQLList contains a list of BackupInstancePostgreSQLs
type BackupInstancePostgreSQLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupInstancePostgreSQL `json:"items"`
}

// Repository type metadata.
var (
	BackupInstancePostgreSQL_Kind             = "BackupInstancePostgreSQL"
	BackupInstancePostgreSQL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupInstancePostgreSQL_Kind}.String()
	BackupInstancePostgreSQL_KindAPIVersion   = BackupInstancePostgreSQL_Kind + "." + CRDGroupVersion.String()
	BackupInstancePostgreSQL_GroupVersionKind = CRDGroupVersion.WithKind(BackupInstancePostgreSQL_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupInstancePostgreSQL{}, &BackupInstancePostgreSQLList{})
}
