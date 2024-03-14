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

type PermissionsInitParameters struct {

	// A list of data actions that are allowed for the Cosmos DB SQL Role Definition.
	// +listType=set
	DataActions []*string `json:"dataActions,omitempty" tf:"data_actions,omitempty"`
}

type PermissionsObservation struct {

	// A list of data actions that are allowed for the Cosmos DB SQL Role Definition.
	// +listType=set
	DataActions []*string `json:"dataActions,omitempty" tf:"data_actions,omitempty"`
}

type PermissionsParameters struct {

	// A list of data actions that are allowed for the Cosmos DB SQL Role Definition.
	// +kubebuilder:validation:Optional
	// +listType=set
	DataActions []*string `json:"dataActions" tf:"data_actions,omitempty"`
}

type SQLRoleDefinitionInitParameters struct {

	// The name of the Cosmos DB Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// A list of fully qualified scopes at or below which Role Assignments may be created using this Cosmos DB SQL Role Definition. It will allow application of this Cosmos DB SQL Role Definition on the entire Database Account or any underlying Database/Collection. Scopes higher than Database Account are not enforceable as assignable scopes.
	// +listType=set
	AssignableScopes []*string `json:"assignableScopes,omitempty" tf:"assignable_scopes,omitempty"`

	// An user-friendly name for the Cosmos DB SQL Role Definition which must be unique for the Database Account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A permissions block as defined below.
	Permissions []PermissionsInitParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The name of the Resource Group in which the Cosmos DB SQL Role Definition is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The GUID as the name of the Cosmos DB SQL Role Definition - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The type of the Cosmos DB SQL Role Definition. Possible values are BuiltInRole and CustomRole. Defaults to CustomRole. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLRoleDefinitionObservation struct {

	// The name of the Cosmos DB Account. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// A list of fully qualified scopes at or below which Role Assignments may be created using this Cosmos DB SQL Role Definition. It will allow application of this Cosmos DB SQL Role Definition on the entire Database Account or any underlying Database/Collection. Scopes higher than Database Account are not enforceable as assignable scopes.
	// +listType=set
	AssignableScopes []*string `json:"assignableScopes,omitempty" tf:"assignable_scopes,omitempty"`

	// The ID of the Cosmos DB SQL Role Definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An user-friendly name for the Cosmos DB SQL Role Definition which must be unique for the Database Account.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A permissions block as defined below.
	Permissions []PermissionsObservation `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The name of the Resource Group in which the Cosmos DB SQL Role Definition is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The GUID as the name of the Cosmos DB SQL Role Definition - one will be generated if not specified. Changing this forces a new resource to be created.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The type of the Cosmos DB SQL Role Definition. Possible values are BuiltInRole and CustomRole. Defaults to CustomRole. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SQLRoleDefinitionParameters struct {

	// The name of the Cosmos DB Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// Reference to a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate accountName.
	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// A list of fully qualified scopes at or below which Role Assignments may be created using this Cosmos DB SQL Role Definition. It will allow application of this Cosmos DB SQL Role Definition on the entire Database Account or any underlying Database/Collection. Scopes higher than Database Account are not enforceable as assignable scopes.
	// +kubebuilder:validation:Optional
	// +listType=set
	AssignableScopes []*string `json:"assignableScopes,omitempty" tf:"assignable_scopes,omitempty"`

	// An user-friendly name for the Cosmos DB SQL Role Definition which must be unique for the Database Account.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A permissions block as defined below.
	// +kubebuilder:validation:Optional
	Permissions []PermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The name of the Resource Group in which the Cosmos DB SQL Role Definition is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The GUID as the name of the Cosmos DB SQL Role Definition - one will be generated if not specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The type of the Cosmos DB SQL Role Definition. Possible values are BuiltInRole and CustomRole. Defaults to CustomRole. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SQLRoleDefinitionSpec defines the desired state of SQLRoleDefinition
type SQLRoleDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLRoleDefinitionParameters `json:"forProvider"`
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
	InitProvider SQLRoleDefinitionInitParameters `json:"initProvider,omitempty"`
}

// SQLRoleDefinitionStatus defines the observed state of SQLRoleDefinition.
type SQLRoleDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLRoleDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SQLRoleDefinition is the Schema for the SQLRoleDefinitions API. Manages a Cosmos DB SQL Role Definition.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLRoleDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.assignableScopes) || (has(self.initProvider) && has(self.initProvider.assignableScopes))",message="spec.forProvider.assignableScopes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permissions) || (has(self.initProvider) && has(self.initProvider.permissions))",message="spec.forProvider.permissions is a required parameter"
	Spec   SQLRoleDefinitionSpec   `json:"spec"`
	Status SQLRoleDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleDefinitionList contains a list of SQLRoleDefinitions
type SQLRoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLRoleDefinition `json:"items"`
}

// Repository type metadata.
var (
	SQLRoleDefinition_Kind             = "SQLRoleDefinition"
	SQLRoleDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLRoleDefinition_Kind}.String()
	SQLRoleDefinition_KindAPIVersion   = SQLRoleDefinition_Kind + "." + CRDGroupVersion.String()
	SQLRoleDefinition_GroupVersionKind = CRDGroupVersion.WithKind(SQLRoleDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLRoleDefinition{}, &SQLRoleDefinitionList{})
}
