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

type SSHPublicKeyInitParameters struct {

	// The Azure Region where the SSH Public Key should exist. Changing this forces a new SSH Public Key to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// SSH public key used to authenticate to a virtual machine through ssh. the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// A mapping of tags which should be assigned to the SSH Public Key.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SSHPublicKeyObservation struct {

	// The ID of the SSH Public Key.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the SSH Public Key should exist. Changing this forces a new SSH Public Key to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// SSH public key used to authenticate to a virtual machine through ssh. the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The name of the Resource Group where the SSH Public Key should exist. Changing this forces a new SSH Public Key to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags which should be assigned to the SSH Public Key.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SSHPublicKeyParameters struct {

	// The Azure Region where the SSH Public Key should exist. Changing this forces a new SSH Public Key to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// SSH public key used to authenticate to a virtual machine through ssh. the provided public key needs to be at least 2048-bit and in ssh-rsa format.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The name of the Resource Group where the SSH Public Key should exist. Changing this forces a new SSH Public Key to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the SSH Public Key.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SSHPublicKeySpec defines the desired state of SSHPublicKey
type SSHPublicKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHPublicKeyParameters `json:"forProvider"`
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
	InitProvider SSHPublicKeyInitParameters `json:"initProvider,omitempty"`
}

// SSHPublicKeyStatus defines the observed state of SSHPublicKey.
type SSHPublicKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHPublicKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SSHPublicKey is the Schema for the SSHPublicKeys API. Manages a SSH Public Key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SSHPublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   SSHPublicKeySpec   `json:"spec"`
	Status SSHPublicKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHPublicKeyList contains a list of SSHPublicKeys
type SSHPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHPublicKey `json:"items"`
}

// Repository type metadata.
var (
	SSHPublicKey_Kind             = "SSHPublicKey"
	SSHPublicKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHPublicKey_Kind}.String()
	SSHPublicKey_KindAPIVersion   = SSHPublicKey_Kind + "." + CRDGroupVersion.String()
	SSHPublicKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHPublicKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHPublicKey{}, &SSHPublicKeyList{})
}
