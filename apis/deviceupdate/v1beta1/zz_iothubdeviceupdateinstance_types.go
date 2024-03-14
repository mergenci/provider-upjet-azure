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

type DiagnosticStorageAccountInitParameters struct {

	// Resource ID of the Diagnostic Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Account in storage to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type DiagnosticStorageAccountObservation struct {

	// Resource ID of the Diagnostic Storage Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DiagnosticStorageAccountParameters struct {

	// Connection String of the Diagnostic Storage Account.
	// +kubebuilder:validation:Required
	ConnectionStringSecretRef v1.SecretKeySelector `json:"connectionStringSecretRef" tf:"-"`

	// Resource ID of the Diagnostic Storage Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Account in storage to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type IOTHubDeviceUpdateInstanceInitParameters struct {

	// Whether the diagnostic log collection is enabled. Possible values are true and false. Defaults to false.
	DiagnosticEnabled *bool `json:"diagnosticEnabled,omitempty" tf:"diagnostic_enabled,omitempty"`

	// A diagnostic_storage_account block as defined below.
	DiagnosticStorageAccount []DiagnosticStorageAccountInitParameters `json:"diagnosticStorageAccount,omitempty" tf:"diagnostic_storage_account,omitempty"`

	// Specifies the ID of the IoT Hub associated with the IoT Hub Device Update Instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the IoT Hub Device Update Instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDeviceUpdateInstanceObservation struct {

	// Specifies the ID of the IoT Hub Device Update Account where the IoT Hub Device Update Instance exists. Changing this forces a new resource to be created.
	DeviceUpdateAccountID *string `json:"deviceUpdateAccountId,omitempty" tf:"device_update_account_id,omitempty"`

	// Whether the diagnostic log collection is enabled. Possible values are true and false. Defaults to false.
	DiagnosticEnabled *bool `json:"diagnosticEnabled,omitempty" tf:"diagnostic_enabled,omitempty"`

	// A diagnostic_storage_account block as defined below.
	DiagnosticStorageAccount []DiagnosticStorageAccountObservation `json:"diagnosticStorageAccount,omitempty" tf:"diagnostic_storage_account,omitempty"`

	// The ID of the IoT Hub Device Update Instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the IoT Hub associated with the IoT Hub Device Update Instance. Changing this forces a new resource to be created.
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// A mapping of tags which should be assigned to the IoT Hub Device Update Instance.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IOTHubDeviceUpdateInstanceParameters struct {

	// Specifies the ID of the IoT Hub Device Update Account where the IoT Hub Device Update Instance exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/deviceupdate/v1beta1.IOTHubDeviceUpdateAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeviceUpdateAccountID *string `json:"deviceUpdateAccountId,omitempty" tf:"device_update_account_id,omitempty"`

	// Reference to a IOTHubDeviceUpdateAccount in deviceupdate to populate deviceUpdateAccountId.
	// +kubebuilder:validation:Optional
	DeviceUpdateAccountIDRef *v1.Reference `json:"deviceUpdateAccountIdRef,omitempty" tf:"-"`

	// Selector for a IOTHubDeviceUpdateAccount in deviceupdate to populate deviceUpdateAccountId.
	// +kubebuilder:validation:Optional
	DeviceUpdateAccountIDSelector *v1.Selector `json:"deviceUpdateAccountIdSelector,omitempty" tf:"-"`

	// Whether the diagnostic log collection is enabled. Possible values are true and false. Defaults to false.
	// +kubebuilder:validation:Optional
	DiagnosticEnabled *bool `json:"diagnosticEnabled,omitempty" tf:"diagnostic_enabled,omitempty"`

	// A diagnostic_storage_account block as defined below.
	// +kubebuilder:validation:Optional
	DiagnosticStorageAccount []DiagnosticStorageAccountParameters `json:"diagnosticStorageAccount,omitempty" tf:"diagnostic_storage_account,omitempty"`

	// Specifies the ID of the IoT Hub associated with the IoT Hub Device Update Instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the IoT Hub Device Update Instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// IOTHubDeviceUpdateInstanceSpec defines the desired state of IOTHubDeviceUpdateInstance
type IOTHubDeviceUpdateInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubDeviceUpdateInstanceParameters `json:"forProvider"`
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
	InitProvider IOTHubDeviceUpdateInstanceInitParameters `json:"initProvider,omitempty"`
}

// IOTHubDeviceUpdateInstanceStatus defines the observed state of IOTHubDeviceUpdateInstance.
type IOTHubDeviceUpdateInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubDeviceUpdateInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IOTHubDeviceUpdateInstance is the Schema for the IOTHubDeviceUpdateInstances API. Manages an IoT Hub Device Update Instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubDeviceUpdateInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubDeviceUpdateInstanceSpec   `json:"spec"`
	Status            IOTHubDeviceUpdateInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDeviceUpdateInstanceList contains a list of IOTHubDeviceUpdateInstances
type IOTHubDeviceUpdateInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubDeviceUpdateInstance `json:"items"`
}

// Repository type metadata.
var (
	IOTHubDeviceUpdateInstance_Kind             = "IOTHubDeviceUpdateInstance"
	IOTHubDeviceUpdateInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubDeviceUpdateInstance_Kind}.String()
	IOTHubDeviceUpdateInstance_KindAPIVersion   = IOTHubDeviceUpdateInstance_Kind + "." + CRDGroupVersion.String()
	IOTHubDeviceUpdateInstance_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubDeviceUpdateInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubDeviceUpdateInstance{}, &IOTHubDeviceUpdateInstanceList{})
}
