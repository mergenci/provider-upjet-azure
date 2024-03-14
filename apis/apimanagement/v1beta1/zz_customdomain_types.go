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

type CustomDomainDeveloperPortalInitParameters struct {

	// The Hostname to use for the corresponding endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_secret_id",true)
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainDeveloperPortalObservation struct {
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`

	CertificateStatus *string `json:"certificateStatus,omitempty" tf:"certificate_status,omitempty"`

	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The Hostname to use for the corresponding endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CustomDomainDeveloperPortalParameters struct {

	// The password associated with the certificate provided above.
	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// The Base64 Encoded Certificate. (Mutually exclusive with key_vault_id.)
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// The Hostname to use for the corresponding endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_secret_id",true)
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	// +kubebuilder:validation:Optional
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainInitParameters struct {

	// One or more developer_portal blocks as defined below.
	DeveloperPortal []CustomDomainDeveloperPortalInitParameters `json:"developerPortal,omitempty" tf:"developer_portal,omitempty"`

	// One or more gateway blocks as defined below.
	Gateway []GatewayInitParameters `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// One or more management blocks as defined below.
	Management []CustomDomainManagementInitParameters `json:"management,omitempty" tf:"management,omitempty"`

	// One or more portal blocks as defined below.
	Portal []CustomDomainPortalInitParameters `json:"portal,omitempty" tf:"portal,omitempty"`

	// One or more scm blocks as defined below.
	Scm []CustomDomainScmInitParameters `json:"scm,omitempty" tf:"scm,omitempty"`
}

type CustomDomainManagementInitParameters struct {

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainManagementObservation struct {
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`

	CertificateStatus *string `json:"certificateStatus,omitempty" tf:"certificate_status,omitempty"`

	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CustomDomainManagementParameters struct {

	// The password associated with the certificate provided above.
	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// The Base64 Encoded Certificate. (Mutually exclusive with key_vault_id.)
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// The Hostname to use for the API Proxy Endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	// +kubebuilder:validation:Optional
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainObservation struct {

	// The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// One or more developer_portal blocks as defined below.
	DeveloperPortal []CustomDomainDeveloperPortalObservation `json:"developerPortal,omitempty" tf:"developer_portal,omitempty"`

	// One or more gateway blocks as defined below.
	Gateway []GatewayObservation `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The ID of the API Management Custom Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more management blocks as defined below.
	Management []CustomDomainManagementObservation `json:"management,omitempty" tf:"management,omitempty"`

	// One or more portal blocks as defined below.
	Portal []CustomDomainPortalObservation `json:"portal,omitempty" tf:"portal,omitempty"`

	// One or more scm blocks as defined below.
	Scm []CustomDomainScmObservation `json:"scm,omitempty" tf:"scm,omitempty"`
}

type CustomDomainParameters struct {

	// The ID of the API Management service for which to configure Custom Domains. Changing this forces a new API Management Custom Domain resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/apimanagement/v1beta1.Management
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// Reference to a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDRef *v1.Reference `json:"apiManagementIdRef,omitempty" tf:"-"`

	// Selector for a Management in apimanagement to populate apiManagementId.
	// +kubebuilder:validation:Optional
	APIManagementIDSelector *v1.Selector `json:"apiManagementIdSelector,omitempty" tf:"-"`

	// One or more developer_portal blocks as defined below.
	// +kubebuilder:validation:Optional
	DeveloperPortal []CustomDomainDeveloperPortalParameters `json:"developerPortal,omitempty" tf:"developer_portal,omitempty"`

	// One or more gateway blocks as defined below.
	// +kubebuilder:validation:Optional
	Gateway []GatewayParameters `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// One or more management blocks as defined below.
	// +kubebuilder:validation:Optional
	Management []CustomDomainManagementParameters `json:"management,omitempty" tf:"management,omitempty"`

	// One or more portal blocks as defined below.
	// +kubebuilder:validation:Optional
	Portal []CustomDomainPortalParameters `json:"portal,omitempty" tf:"portal,omitempty"`

	// One or more scm blocks as defined below.
	// +kubebuilder:validation:Optional
	Scm []CustomDomainScmParameters `json:"scm,omitempty" tf:"scm,omitempty"`
}

type CustomDomainPortalInitParameters struct {

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainPortalObservation struct {
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`

	CertificateStatus *string `json:"certificateStatus,omitempty" tf:"certificate_status,omitempty"`

	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CustomDomainPortalParameters struct {

	// The password associated with the certificate provided above.
	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// The Base64 Encoded Certificate. (Mutually exclusive with key_vault_id.)
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// The Hostname to use for the API Proxy Endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	// +kubebuilder:validation:Optional
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainScmInitParameters struct {

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type CustomDomainScmObservation struct {
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`

	CertificateStatus *string `json:"certificateStatus,omitempty" tf:"certificate_status,omitempty"`

	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type CustomDomainScmParameters struct {

	// The password associated with the certificate provided above.
	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// The Base64 Encoded Certificate. (Mutually exclusive with key_vault_id.)
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// The Hostname to use for the API Proxy Endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	// +kubebuilder:validation:Optional
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type GatewayInitParameters struct {

	// Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to false.
	DefaultSSLBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding,omitempty"`

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_secret_id",true)
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

type GatewayObservation struct {
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`

	CertificateStatus *string `json:"certificateStatus,omitempty" tf:"certificate_status,omitempty"`

	// Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to false.
	DefaultSSLBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding,omitempty"`

	Expiry *string `json:"expiry,omitempty" tf:"expiry,omitempty"`

	// The Hostname to use for the API Proxy Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type GatewayParameters struct {

	// The password associated with the certificate provided above.
	// +kubebuilder:validation:Optional
	CertificatePasswordSecretRef *v1.SecretKeySelector `json:"certificatePasswordSecretRef,omitempty" tf:"-"`

	// The Base64 Encoded Certificate. (Mutually exclusive with key_vault_id.)
	// +kubebuilder:validation:Optional
	CertificateSecretRef *v1.SecretKeySelector `json:"certificateSecretRef,omitempty" tf:"-"`

	// Is the certificate associated with this Hostname the Default SSL Certificate? This is used when an SNI header isn't specified by a client. Defaults to false.
	// +kubebuilder:validation:Optional
	DefaultSSLBinding *bool `json:"defaultSslBinding,omitempty" tf:"default_ssl_binding,omitempty"`

	// The Hostname to use for the API Proxy Endpoint.
	// +kubebuilder:validation:Optional
	HostName *string `json:"hostName" tf:"host_name,omitempty"`

	// The ID of the Key Vault Secret containing the SSL Certificate, which must be should be of the type application/x-pkcs12.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/keyvault/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("versionless_secret_id",true)
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in keyvault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// Should Client Certificate Negotiation be enabled for this Hostname? Defaults to false.
	// +kubebuilder:validation:Optional
	NegotiateClientCertificate *bool `json:"negotiateClientCertificate,omitempty" tf:"negotiate_client_certificate,omitempty"`

	// System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to the keyVault containing the SSL certificate.
	// +kubebuilder:validation:Optional
	SSLKeyvaultIdentityClientID *string `json:"sslKeyvaultIdentityClientId,omitempty" tf:"ssl_keyvault_identity_client_id,omitempty"`
}

// CustomDomainSpec defines the desired state of CustomDomain
type CustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDomainParameters `json:"forProvider"`
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
	InitProvider CustomDomainInitParameters `json:"initProvider,omitempty"`
}

// CustomDomainStatus defines the observed state of CustomDomain.
type CustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomDomain is the Schema for the CustomDomains API. Manages a API Management Custom Domain.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDomainSpec   `json:"spec"`
	Status            CustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDomainList contains a list of CustomDomains
type CustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDomain `json:"items"`
}

// Repository type metadata.
var (
	CustomDomain_Kind             = "CustomDomain"
	CustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDomain_Kind}.String()
	CustomDomain_KindAPIVersion   = CustomDomain_Kind + "." + CRDGroupVersion.String()
	CustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(CustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDomain{}, &CustomDomainList{})
}
