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

type SpringCloudCertificateInitParameters struct {

	// The content of uploaded certificate. Changing this forces a new resource to be created.
	CertificateContent *string `json:"certificateContent,omitempty" tf:"certificate_content,omitempty"`

	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KeyVaultCertificateID *string `json:"keyVaultCertificateId,omitempty" tf:"key_vault_certificate_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultCertificateId.
	// +kubebuilder:validation:Optional
	KeyVaultCertificateIDRef *v1.Reference `json:"keyVaultCertificateIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultCertificateId.
	// +kubebuilder:validation:Optional
	KeyVaultCertificateIDSelector *v1.Selector `json:"keyVaultCertificateIdSelector,omitempty" tf:"-"`
}

type SpringCloudCertificateObservation struct {

	// The content of uploaded certificate. Changing this forces a new resource to be created.
	CertificateContent *string `json:"certificateContent,omitempty" tf:"certificate_content,omitempty"`

	// The ID of the Spring Cloud Certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	KeyVaultCertificateID *string `json:"keyVaultCertificateId,omitempty" tf:"key_vault_certificate_id,omitempty"`

	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The thumbprint of the Spring Cloud certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudCertificateParameters struct {

	// The content of uploaded certificate. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CertificateContent *string `json:"certificateContent,omitempty" tf:"certificate_content,omitempty"`

	// Specifies the ID of the Key Vault Certificate resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultCertificateID *string `json:"keyVaultCertificateId,omitempty" tf:"key_vault_certificate_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultCertificateId.
	// +kubebuilder:validation:Optional
	KeyVaultCertificateIDRef *v1.Reference `json:"keyVaultCertificateIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultCertificateId.
	// +kubebuilder:validation:Optional
	KeyVaultCertificateIDSelector *v1.Selector `json:"keyVaultCertificateIdSelector,omitempty" tf:"-"`

	// Specifies the name of the resource group in which to create the Spring Cloud Certificate. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the name of the Spring Cloud Service resource. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudService
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`
}

// SpringCloudCertificateSpec defines the desired state of SpringCloudCertificate
type SpringCloudCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudCertificateParameters `json:"forProvider"`
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
	InitProvider SpringCloudCertificateInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudCertificateStatus defines the observed state of SpringCloudCertificate.
type SpringCloudCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SpringCloudCertificate is the Schema for the SpringCloudCertificates API. Manages an Azure Spring Cloud Certificate.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpringCloudCertificateSpec   `json:"spec"`
	Status            SpringCloudCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudCertificateList contains a list of SpringCloudCertificates
type SpringCloudCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudCertificate `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudCertificate_Kind             = "SpringCloudCertificate"
	SpringCloudCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudCertificate_Kind}.String()
	SpringCloudCertificate_KindAPIVersion   = SpringCloudCertificate_Kind + "." + CRDGroupVersion.String()
	SpringCloudCertificate_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudCertificate{}, &SpringCloudCertificateList{})
}
