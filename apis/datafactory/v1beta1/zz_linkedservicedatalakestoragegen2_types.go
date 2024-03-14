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

type LinkedServiceDataLakeStorageGen2InitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with storage_account_key and use_managed_identity.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with service_principal_id, service_principal_key, tenant and use_managed_identity.
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// The endpoint for the Azure Data Lake Storage Gen2 service.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with service_principal_id, service_principal_key, tenant and storage_account_key.
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type LinkedServiceDataLakeStorageGen2Observation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Data Lake Storage Gen2 Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with storage_account_key and use_managed_identity.
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with service_principal_id, service_principal_key, tenant and use_managed_identity.
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// The endpoint for the Azure Data Lake Storage Gen2 service.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with service_principal_id, service_principal_key, tenant and storage_account_key.
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

type LinkedServiceDataLakeStorageGen2Parameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The service principal id with which to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with storage_account_key and use_managed_identity.
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// The service principal key with which to authenticate against the Azure Data Lake Storage Gen2 account.
	// +kubebuilder:validation:Optional
	ServicePrincipalKey *string `json:"servicePrincipalKey,omitempty" tf:"service_principal_key,omitempty"`

	// The Storage Account Key with which to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with service_principal_id, service_principal_key, tenant and use_managed_identity.
	// +kubebuilder:validation:Optional
	StorageAccountKey *string `json:"storageAccountKey,omitempty" tf:"storage_account_key,omitempty"`

	// The tenant id or name in which the service principal exists to authenticate against the Azure Data Lake Storage Gen2 account.
	// +kubebuilder:validation:Optional
	Tenant *string `json:"tenant,omitempty" tf:"tenant,omitempty"`

	// The endpoint for the Azure Data Lake Storage Gen2 service.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Whether to use the Data Factory's managed identity to authenticate against the Azure Data Lake Storage Gen2 account. Incompatible with service_principal_id, service_principal_key, tenant and storage_account_key.
	// +kubebuilder:validation:Optional
	UseManagedIdentity *bool `json:"useManagedIdentity,omitempty" tf:"use_managed_identity,omitempty"`
}

// LinkedServiceDataLakeStorageGen2Spec defines the desired state of LinkedServiceDataLakeStorageGen2
type LinkedServiceDataLakeStorageGen2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceDataLakeStorageGen2Parameters `json:"forProvider"`
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
	InitProvider LinkedServiceDataLakeStorageGen2InitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceDataLakeStorageGen2Status defines the observed state of LinkedServiceDataLakeStorageGen2.
type LinkedServiceDataLakeStorageGen2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceDataLakeStorageGen2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkedServiceDataLakeStorageGen2 is the Schema for the LinkedServiceDataLakeStorageGen2s API. Manages a Linked Service (connection) between Data Lake Storage Gen2 and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceDataLakeStorageGen2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   LinkedServiceDataLakeStorageGen2Spec   `json:"spec"`
	Status LinkedServiceDataLakeStorageGen2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceDataLakeStorageGen2List contains a list of LinkedServiceDataLakeStorageGen2s
type LinkedServiceDataLakeStorageGen2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceDataLakeStorageGen2 `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceDataLakeStorageGen2_Kind             = "LinkedServiceDataLakeStorageGen2"
	LinkedServiceDataLakeStorageGen2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceDataLakeStorageGen2_Kind}.String()
	LinkedServiceDataLakeStorageGen2_KindAPIVersion   = LinkedServiceDataLakeStorageGen2_Kind + "." + CRDGroupVersion.String()
	LinkedServiceDataLakeStorageGen2_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceDataLakeStorageGen2_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceDataLakeStorageGen2{}, &LinkedServiceDataLakeStorageGen2List{})
}
