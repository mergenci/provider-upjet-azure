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

type LinkedServiceAzureSearchInitParameters struct {

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

	// The key of the Azure Search Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/search/v1beta1.Service
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_key",true)
	SearchServiceKey *string `json:"searchServiceKey,omitempty" tf:"search_service_key,omitempty"`

	// Reference to a Service in search to populate searchServiceKey.
	// +kubebuilder:validation:Optional
	SearchServiceKeyRef *v1.Reference `json:"searchServiceKeyRef,omitempty" tf:"-"`

	// Selector for a Service in search to populate searchServiceKey.
	// +kubebuilder:validation:Optional
	SearchServiceKeySelector *v1.Selector `json:"searchServiceKeySelector,omitempty" tf:"-"`

	// The URL of the Search Service endpoint (e.g. https://{searchServiceName}.search.windows.net).
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkedServiceAzureSearchObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encrypted credential to connect to Azure Search Service.
	EncryptedCredential *string `json:"encryptedCredential,omitempty" tf:"encrypted_credential,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The key of the Azure Search Service.
	SearchServiceKey *string `json:"searchServiceKey,omitempty" tf:"search_service_key,omitempty"`

	// The URL of the Search Service endpoint (e.g. https://{searchServiceName}.search.windows.net).
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkedServiceAzureSearchParameters struct {

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

	// The key of the Azure Search Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/search/v1beta1.Service
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("primary_key",true)
	// +kubebuilder:validation:Optional
	SearchServiceKey *string `json:"searchServiceKey,omitempty" tf:"search_service_key,omitempty"`

	// Reference to a Service in search to populate searchServiceKey.
	// +kubebuilder:validation:Optional
	SearchServiceKeyRef *v1.Reference `json:"searchServiceKeyRef,omitempty" tf:"-"`

	// Selector for a Service in search to populate searchServiceKey.
	// +kubebuilder:validation:Optional
	SearchServiceKeySelector *v1.Selector `json:"searchServiceKeySelector,omitempty" tf:"-"`

	// The URL of the Search Service endpoint (e.g. https://{searchServiceName}.search.windows.net).
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// LinkedServiceAzureSearchSpec defines the desired state of LinkedServiceAzureSearch
type LinkedServiceAzureSearchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceAzureSearchParameters `json:"forProvider"`
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
	InitProvider LinkedServiceAzureSearchInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceAzureSearchStatus defines the observed state of LinkedServiceAzureSearch.
type LinkedServiceAzureSearchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceAzureSearchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkedServiceAzureSearch is the Schema for the LinkedServiceAzureSearchs API. Manages a Linked Service (connection) between Azure Search Service and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceAzureSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   LinkedServiceAzureSearchSpec   `json:"spec"`
	Status LinkedServiceAzureSearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureSearchList contains a list of LinkedServiceAzureSearchs
type LinkedServiceAzureSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceAzureSearch `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceAzureSearch_Kind             = "LinkedServiceAzureSearch"
	LinkedServiceAzureSearch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceAzureSearch_Kind}.String()
	LinkedServiceAzureSearch_KindAPIVersion   = LinkedServiceAzureSearch_Kind + "." + CRDGroupVersion.String()
	LinkedServiceAzureSearch_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceAzureSearch_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceAzureSearch{}, &LinkedServiceAzureSearchList{})
}
