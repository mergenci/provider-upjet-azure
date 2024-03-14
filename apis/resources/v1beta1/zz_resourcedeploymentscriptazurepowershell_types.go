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

type ResourceDeploymentScriptAzurePowerShellContainerInitParameters struct {

	// Container group name, if not specified then the name will get auto-generated. For more information, please refer to the Container Configuration documentation.
	ContainerGroupName *string `json:"containerGroupName,omitempty" tf:"container_group_name,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellContainerObservation struct {

	// Container group name, if not specified then the name will get auto-generated. For more information, please refer to the Container Configuration documentation.
	ContainerGroupName *string `json:"containerGroupName,omitempty" tf:"container_group_name,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellContainerParameters struct {

	// Container group name, if not specified then the name will get auto-generated. For more information, please refer to the Container Configuration documentation.
	// +kubebuilder:validation:Optional
	ContainerGroupName *string `json:"containerGroupName,omitempty" tf:"container_group_name,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellEnvironmentVariableInitParameters struct {

	// Specifies the name of the environment variable.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the environment variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellEnvironmentVariableObservation struct {

	// Specifies the name of the environment variable.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the value of the environment variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellEnvironmentVariableParameters struct {

	// Specifies the name of the environment variable.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the value of the secure environment variable.
	// +kubebuilder:validation:Optional
	SecureValueSecretRef *v1.SecretKeySelector `json:"secureValueSecretRef,omitempty" tf:"-"`

	// Specifies the value of the environment variable.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellIdentityInitParameters struct {

	// Specifies the list of user-assigned managed identity IDs associated with the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// References to UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsRefs []v1.Reference `json:"identityIdsRefs,omitempty" tf:"-"`

	// Selector for a list of UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsSelector *v1.Selector `json:"identityIdsSelector,omitempty" tf:"-"`

	// Type of the managed identity. The only possible value is UserAssigned. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellIdentityObservation struct {

	// Specifies the list of user-assigned managed identity IDs associated with the resource. Changing this forces a new resource to be created.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Type of the managed identity. The only possible value is UserAssigned. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellIdentityParameters struct {

	// Specifies the list of user-assigned managed identity IDs associated with the resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/managedidentity/v1beta1.UserAssignedIdentity
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// References to UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsRefs []v1.Reference `json:"identityIdsRefs,omitempty" tf:"-"`

	// Selector for a list of UserAssignedIdentity in managedidentity to populate identityIds.
	// +kubebuilder:validation:Optional
	IdentityIdsSelector *v1.Selector `json:"identityIdsSelector,omitempty" tf:"-"`

	// Type of the managed identity. The only possible value is UserAssigned. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellInitParameters struct {

	// Specifies the cleanup preference when the script execution gets in a terminal state. Possible values are Always, OnExpiration, OnSuccess. Defaults to Always. Changing this forces a new Resource Deployment Script to be created.
	CleanupPreference *string `json:"cleanupPreference,omitempty" tf:"cleanup_preference,omitempty"`

	// Command line arguments to pass to the script. Changing this forces a new Resource Deployment Script to be created.
	CommandLine *string `json:"commandLine,omitempty" tf:"command_line,omitempty"`

	// A container block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Container []ResourceDeploymentScriptAzurePowerShellContainerInitParameters `json:"container,omitempty" tf:"container,omitempty"`

	// An environment_variable block as defined below. Changing this forces a new Resource Deployment Script to be created.
	EnvironmentVariable []ResourceDeploymentScriptAzurePowerShellEnvironmentVariableInitParameters `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID. Changing this forces a new Resource Deployment Script to be created.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// An identity block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Identity []ResourceDeploymentScriptAzurePowerShellIdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Resource Deployment Script. The name length must be from 1 to 260 characters. The name can only contain alphanumeric, underscore, parentheses, hyphen and period, and it cannot end with a period. Changing this forces a new Resource Deployment Script to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Uri for the script. This is the entry point for the external script. Changing this forces a new Resource Deployment Script to be created.
	PrimaryScriptURI *string `json:"primaryScriptUri,omitempty" tf:"primary_script_uri,omitempty"`

	// Specifies the name of the Resource Group where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. The time duration should be between 1 hour and 26 hours (inclusive) and should be specified in ISO 8601 format. Changing this forces a new Resource Deployment Script to be created.
	RetentionInterval *string `json:"retentionInterval,omitempty" tf:"retention_interval,omitempty"`

	// Script body. Changing this forces a new Resource Deployment Script to be created.
	ScriptContent *string `json:"scriptContent,omitempty" tf:"script_content,omitempty"`

	// A storage_account block as defined below. Changing this forces a new Resource Deployment Script to be created.
	StorageAccount []ResourceDeploymentScriptAzurePowerShellStorageAccountInitParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Supporting files for the external script. Changing this forces a new Resource Deployment Script to be created.
	SupportingScriptUris []*string `json:"supportingScriptUris,omitempty" tf:"supporting_script_uris,omitempty"`

	// A mapping of tags which should be assigned to the Resource Deployment Script.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum allowed script execution time specified in ISO 8601 format. Needs to be greater than 0 and smaller than 1 day. Defaults to P1D. Changing this forces a new Resource Deployment Script to be created.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Azure PowerShell module version to be used. The supported versions are 2.7, 2.8, 3.0, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 5.0, 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0, 6.1, 6.2, 6.3, 6.4, 6.5, 6.6, 7.0, 7.1, 7.2, 7.3, 7.4, 7.5, 8.0, 8.1, 8.2, 8.3, 9.0. Changing this forces a new Resource Deployment Script to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellObservation struct {

	// Specifies the cleanup preference when the script execution gets in a terminal state. Possible values are Always, OnExpiration, OnSuccess. Defaults to Always. Changing this forces a new Resource Deployment Script to be created.
	CleanupPreference *string `json:"cleanupPreference,omitempty" tf:"cleanup_preference,omitempty"`

	// Command line arguments to pass to the script. Changing this forces a new Resource Deployment Script to be created.
	CommandLine *string `json:"commandLine,omitempty" tf:"command_line,omitempty"`

	// A container block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Container []ResourceDeploymentScriptAzurePowerShellContainerObservation `json:"container,omitempty" tf:"container,omitempty"`

	// An environment_variable block as defined below. Changing this forces a new Resource Deployment Script to be created.
	EnvironmentVariable []ResourceDeploymentScriptAzurePowerShellEnvironmentVariableObservation `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID. Changing this forces a new Resource Deployment Script to be created.
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// The ID of the Resource Deployment Script.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below. Changing this forces a new Resource Deployment Script to be created.
	Identity []ResourceDeploymentScriptAzurePowerShellIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Resource Deployment Script. The name length must be from 1 to 260 characters. The name can only contain alphanumeric, underscore, parentheses, hyphen and period, and it cannot end with a period. Changing this forces a new Resource Deployment Script to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// List of script outputs.
	Outputs *string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// Uri for the script. This is the entry point for the external script. Changing this forces a new Resource Deployment Script to be created.
	PrimaryScriptURI *string `json:"primaryScriptUri,omitempty" tf:"primary_script_uri,omitempty"`

	// Specifies the name of the Resource Group where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. The time duration should be between 1 hour and 26 hours (inclusive) and should be specified in ISO 8601 format. Changing this forces a new Resource Deployment Script to be created.
	RetentionInterval *string `json:"retentionInterval,omitempty" tf:"retention_interval,omitempty"`

	// Script body. Changing this forces a new Resource Deployment Script to be created.
	ScriptContent *string `json:"scriptContent,omitempty" tf:"script_content,omitempty"`

	// A storage_account block as defined below. Changing this forces a new Resource Deployment Script to be created.
	StorageAccount []ResourceDeploymentScriptAzurePowerShellStorageAccountObservation `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Supporting files for the external script. Changing this forces a new Resource Deployment Script to be created.
	SupportingScriptUris []*string `json:"supportingScriptUris,omitempty" tf:"supporting_script_uris,omitempty"`

	// A mapping of tags which should be assigned to the Resource Deployment Script.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum allowed script execution time specified in ISO 8601 format. Needs to be greater than 0 and smaller than 1 day. Defaults to P1D. Changing this forces a new Resource Deployment Script to be created.
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Azure PowerShell module version to be used. The supported versions are 2.7, 2.8, 3.0, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 5.0, 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0, 6.1, 6.2, 6.3, 6.4, 6.5, 6.6, 7.0, 7.1, 7.2, 7.3, 7.4, 7.5, 8.0, 8.1, 8.2, 8.3, 9.0. Changing this forces a new Resource Deployment Script to be created.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellParameters struct {

	// Specifies the cleanup preference when the script execution gets in a terminal state. Possible values are Always, OnExpiration, OnSuccess. Defaults to Always. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	CleanupPreference *string `json:"cleanupPreference,omitempty" tf:"cleanup_preference,omitempty"`

	// Command line arguments to pass to the script. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	CommandLine *string `json:"commandLine,omitempty" tf:"command_line,omitempty"`

	// A container block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Container []ResourceDeploymentScriptAzurePowerShellContainerParameters `json:"container,omitempty" tf:"container,omitempty"`

	// An environment_variable block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	EnvironmentVariable []ResourceDeploymentScriptAzurePowerShellEnvironmentVariableParameters `json:"environmentVariable,omitempty" tf:"environment_variable,omitempty"`

	// Gets or sets how the deployment script should be forced to execute even if the script resource has not changed. Can be current time stamp or a GUID. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// An identity block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Identity []ResourceDeploymentScriptAzurePowerShellIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the Azure Region where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the name which should be used for this Resource Deployment Script. The name length must be from 1 to 260 characters. The name can only contain alphanumeric, underscore, parentheses, hyphen and period, and it cannot end with a period. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Uri for the script. This is the entry point for the external script. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	PrimaryScriptURI *string `json:"primaryScriptUri,omitempty" tf:"primary_script_uri,omitempty"`

	// Specifies the name of the Resource Group where the Resource Deployment Script should exist. Changing this forces a new Resource Deployment Script to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Interval for which the service retains the script resource after it reaches a terminal state. Resource will be deleted when this duration expires. The time duration should be between 1 hour and 26 hours (inclusive) and should be specified in ISO 8601 format. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	RetentionInterval *string `json:"retentionInterval,omitempty" tf:"retention_interval,omitempty"`

	// Script body. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	ScriptContent *string `json:"scriptContent,omitempty" tf:"script_content,omitempty"`

	// A storage_account block as defined below. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	StorageAccount []ResourceDeploymentScriptAzurePowerShellStorageAccountParameters `json:"storageAccount,omitempty" tf:"storage_account,omitempty"`

	// Supporting files for the external script. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	SupportingScriptUris []*string `json:"supportingScriptUris,omitempty" tf:"supporting_script_uris,omitempty"`

	// A mapping of tags which should be assigned to the Resource Deployment Script.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Maximum allowed script execution time specified in ISO 8601 format. Needs to be greater than 0 and smaller than 1 day. Defaults to P1D. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Azure PowerShell module version to be used. The supported versions are 2.7, 2.8, 3.0, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 5.0, 5.1, 5.2, 5.3, 5.4, 5.5, 5.6, 5.7, 5.8, 5.9, 6.0, 6.1, 6.2, 6.3, 6.4, 6.5, 6.6, 7.0, 7.1, 7.2, 7.3, 7.4, 7.5, 8.0, 8.1, 8.2, 8.3, 9.0. Changing this forces a new Resource Deployment Script to be created.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellStorageAccountInitParameters struct {

	// Specifies the storage account name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellStorageAccountObservation struct {

	// Specifies the storage account name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResourceDeploymentScriptAzurePowerShellStorageAccountParameters struct {

	// Specifies the storage account access key.
	// +kubebuilder:validation:Required
	KeySecretRef v1.SecretKeySelector `json:"keySecretRef" tf:"-"`

	// Specifies the storage account name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// ResourceDeploymentScriptAzurePowerShellSpec defines the desired state of ResourceDeploymentScriptAzurePowerShell
type ResourceDeploymentScriptAzurePowerShellSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceDeploymentScriptAzurePowerShellParameters `json:"forProvider"`
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
	InitProvider ResourceDeploymentScriptAzurePowerShellInitParameters `json:"initProvider,omitempty"`
}

// ResourceDeploymentScriptAzurePowerShellStatus defines the observed state of ResourceDeploymentScriptAzurePowerShell.
type ResourceDeploymentScriptAzurePowerShellStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceDeploymentScriptAzurePowerShellObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResourceDeploymentScriptAzurePowerShell is the Schema for the ResourceDeploymentScriptAzurePowerShells API. Manages a Resource Deployment Script of Azure PowerShell.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceDeploymentScriptAzurePowerShell struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.retentionInterval) || (has(self.initProvider) && has(self.initProvider.retentionInterval))",message="spec.forProvider.retentionInterval is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   ResourceDeploymentScriptAzurePowerShellSpec   `json:"spec"`
	Status ResourceDeploymentScriptAzurePowerShellStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceDeploymentScriptAzurePowerShellList contains a list of ResourceDeploymentScriptAzurePowerShells
type ResourceDeploymentScriptAzurePowerShellList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceDeploymentScriptAzurePowerShell `json:"items"`
}

// Repository type metadata.
var (
	ResourceDeploymentScriptAzurePowerShell_Kind             = "ResourceDeploymentScriptAzurePowerShell"
	ResourceDeploymentScriptAzurePowerShell_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceDeploymentScriptAzurePowerShell_Kind}.String()
	ResourceDeploymentScriptAzurePowerShell_KindAPIVersion   = ResourceDeploymentScriptAzurePowerShell_Kind + "." + CRDGroupVersion.String()
	ResourceDeploymentScriptAzurePowerShell_GroupVersionKind = CRDGroupVersion.WithKind(ResourceDeploymentScriptAzurePowerShell_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceDeploymentScriptAzurePowerShell{}, &ResourceDeploymentScriptAzurePowerShellList{})
}
