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

type KeyPropertyInitParameters struct {

	// The name of the key property. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The data type of the key property. Valid values include Bool, DateTime, Double, String. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type KeyPropertyObservation struct {

	// The name of the key property. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The data type of the key property. Valid values include Bool, DateTime, Double, String. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type KeyPropertyParameters struct {

	// The name of the key property. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The data type of the key property. Valid values include Bool, DateTime, Double, String. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ReferenceDataSetInitParameters struct {

	// The comparison behavior that will be used to compare keys. Valid values include Ordinal and OrdinalIgnoreCase. Defaults to Ordinal. Changing this forces a new resource to be created.
	DataStringComparisonBehavior *string `json:"dataStringComparisonBehavior,omitempty" tf:"data_string_comparison_behavior,omitempty"`

	// A key_property block as defined below. Changing this forces a new resource to be created.
	KeyProperty []KeyPropertyInitParameters `json:"keyProperty,omitempty" tf:"key_property,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ReferenceDataSetObservation struct {

	// The comparison behavior that will be used to compare keys. Valid values include Ordinal and OrdinalIgnoreCase. Defaults to Ordinal. Changing this forces a new resource to be created.
	DataStringComparisonBehavior *string `json:"dataStringComparisonBehavior,omitempty" tf:"data_string_comparison_behavior,omitempty"`

	// The ID of the IoT Time Series Insights Reference Data Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A key_property block as defined below. Changing this forces a new resource to be created.
	KeyProperty []KeyPropertyObservation `json:"keyProperty,omitempty" tf:"key_property,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	TimeSeriesInsightsEnvironmentID *string `json:"timeSeriesInsightsEnvironmentId,omitempty" tf:"time_series_insights_environment_id,omitempty"`
}

type ReferenceDataSetParameters struct {

	// The comparison behavior that will be used to compare keys. Valid values include Ordinal and OrdinalIgnoreCase. Defaults to Ordinal. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataStringComparisonBehavior *string `json:"dataStringComparisonBehavior,omitempty" tf:"data_string_comparison_behavior,omitempty"`

	// A key_property block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	KeyProperty []KeyPropertyParameters `json:"keyProperty,omitempty" tf:"key_property,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/timeseriesinsights/v1beta1.StandardEnvironment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TimeSeriesInsightsEnvironmentID *string `json:"timeSeriesInsightsEnvironmentId,omitempty" tf:"time_series_insights_environment_id,omitempty"`

	// Reference to a StandardEnvironment in timeseriesinsights to populate timeSeriesInsightsEnvironmentId.
	// +kubebuilder:validation:Optional
	TimeSeriesInsightsEnvironmentIDRef *v1.Reference `json:"timeSeriesInsightsEnvironmentIdRef,omitempty" tf:"-"`

	// Selector for a StandardEnvironment in timeseriesinsights to populate timeSeriesInsightsEnvironmentId.
	// +kubebuilder:validation:Optional
	TimeSeriesInsightsEnvironmentIDSelector *v1.Selector `json:"timeSeriesInsightsEnvironmentIdSelector,omitempty" tf:"-"`
}

// ReferenceDataSetSpec defines the desired state of ReferenceDataSet
type ReferenceDataSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReferenceDataSetParameters `json:"forProvider"`
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
	InitProvider ReferenceDataSetInitParameters `json:"initProvider,omitempty"`
}

// ReferenceDataSetStatus defines the observed state of ReferenceDataSet.
type ReferenceDataSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReferenceDataSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ReferenceDataSet is the Schema for the ReferenceDataSets API. Manages an Azure IoT Time Series Insights Reference Data Set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ReferenceDataSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keyProperty) || (has(self.initProvider) && has(self.initProvider.keyProperty))",message="spec.forProvider.keyProperty is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   ReferenceDataSetSpec   `json:"spec"`
	Status ReferenceDataSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceDataSetList contains a list of ReferenceDataSets
type ReferenceDataSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceDataSet `json:"items"`
}

// Repository type metadata.
var (
	ReferenceDataSet_Kind             = "ReferenceDataSet"
	ReferenceDataSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReferenceDataSet_Kind}.String()
	ReferenceDataSet_KindAPIVersion   = ReferenceDataSet_Kind + "." + CRDGroupVersion.String()
	ReferenceDataSet_GroupVersionKind = CRDGroupVersion.WithKind(ReferenceDataSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ReferenceDataSet{}, &ReferenceDataSetList{})
}
