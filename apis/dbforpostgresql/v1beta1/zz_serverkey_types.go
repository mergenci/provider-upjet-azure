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

type ServerKeyInitParameters struct {

	// The URL to a Key Vault Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Server
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

type ServerKeyObservation struct {

	// The ID of the PostgreSQL Server Key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URL to a Key Vault Key.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type ServerKeyParameters struct {

	// The URL to a Key Vault Key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// Reference to a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDRef *v1.Reference `json:"keyVaultKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in keyvault to populate keyVaultKeyId.
	// +kubebuilder:validation:Optional
	KeyVaultKeyIDSelector *v1.Selector `json:"keyVaultKeyIdSelector,omitempty" tf:"-"`

	// The ID of the PostgreSQL Server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Server
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`
}

// ServerKeySpec defines the desired state of ServerKey
type ServerKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerKeyParameters `json:"forProvider"`
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
	InitProvider ServerKeyInitParameters `json:"initProvider,omitempty"`
}

// ServerKeyStatus defines the observed state of ServerKey.
type ServerKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerKey is the Schema for the ServerKeys API. Manages a PostgreSQL Server Key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ServerKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerKeySpec   `json:"spec"`
	Status            ServerKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerKeyList contains a list of ServerKeys
type ServerKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerKey `json:"items"`
}

// Repository type metadata.
var (
	ServerKey_Kind             = "ServerKey"
	ServerKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerKey_Kind}.String()
	ServerKey_KindAPIVersion   = ServerKey_Kind + "." + CRDGroupVersion.String()
	ServerKey_GroupVersionKind = CRDGroupVersion.WithKind(ServerKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerKey{}, &ServerKeyList{})
}
