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

type PublicIPInitParameters struct {

	// Defines the allocation method for this IP address. Possible values are Static or Dynamic.
	AllocationMethod *string `json:"allocationMethod,omitempty" tf:"allocation_method,omitempty"`

	// The DDoS protection mode of the public IP. Possible values are Disabled, Enabled, and VirtualNetworkInherited. Defaults to VirtualNetworkInherited.
	DDOSProtectionMode *string `json:"ddosProtectionMode,omitempty" tf:"ddos_protection_mode,omitempty"`

	// The ID of DDoS protection plan associated with the public IP.
	DDOSProtectionPlanID *string `json:"ddosProtectionPlanId,omitempty" tf:"ddos_protection_plan_id,omitempty"`

	// Label for the Domain Name. Will be used to make up the FQDN. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Public IP should exist. Changing this forces a new Public IP to be created.
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// A mapping of IP tags to assign to the public IP. Changing this forces a new resource to be created.
	// +mapType=granular
	IPTags map[string]*string `json:"ipTags,omitempty" tf:"ip_tags,omitempty"`

	// The IP Version to use, IPv6 or IPv4. Changing this forces a new resource to be created.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the supported Azure location where the Public IP should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// If specified then public IP address allocated will be provided from the public IP prefix resource. Changing this forces a new resource to be created.
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty" tf:"reverse_fqdn,omitempty"`

	// The SKU of the Public IP. Accepted values are Basic and Standard. Defaults to Basic. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The SKU Tier that should be used for the Public IP. Possible values are Regional and Global. Defaults to Regional. Changing this forces a new resource to be created.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A collection containing the availability zone to allocate the Public IP in. Changing this forces a new resource to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type PublicIPObservation struct {

	// Defines the allocation method for this IP address. Possible values are Static or Dynamic.
	AllocationMethod *string `json:"allocationMethod,omitempty" tf:"allocation_method,omitempty"`

	// The DDoS protection mode of the public IP. Possible values are Disabled, Enabled, and VirtualNetworkInherited. Defaults to VirtualNetworkInherited.
	DDOSProtectionMode *string `json:"ddosProtectionMode,omitempty" tf:"ddos_protection_mode,omitempty"`

	// The ID of DDoS protection plan associated with the public IP.
	DDOSProtectionPlanID *string `json:"ddosProtectionPlanId,omitempty" tf:"ddos_protection_plan_id,omitempty"`

	// Label for the Domain Name. Will be used to make up the FQDN. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Public IP should exist. Changing this forces a new Public IP to be created.
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// Fully qualified domain name of the A DNS record associated with the public IP. domain_name_label must be specified to get the fqdn. This is the concatenation of the domain_name_label and the regionalized DNS zone
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The ID of this Public IP.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address value that was allocated.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A mapping of IP tags to assign to the public IP. Changing this forces a new resource to be created.
	// +mapType=granular
	IPTags map[string]*string `json:"ipTags,omitempty" tf:"ip_tags,omitempty"`

	// The IP Version to use, IPv6 or IPv4. Changing this forces a new resource to be created.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the supported Azure location where the Public IP should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// If specified then public IP address allocated will be provided from the public IP prefix resource. Changing this forces a new resource to be created.
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// The name of the Resource Group where this Public IP should exist. Changing this forces a new Public IP to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty" tf:"reverse_fqdn,omitempty"`

	// The SKU of the Public IP. Accepted values are Basic and Standard. Defaults to Basic. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The SKU Tier that should be used for the Public IP. Possible values are Regional and Global. Defaults to Regional. Changing this forces a new resource to be created.
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A collection containing the availability zone to allocate the Public IP in. Changing this forces a new resource to be created.
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type PublicIPParameters struct {

	// Defines the allocation method for this IP address. Possible values are Static or Dynamic.
	// +kubebuilder:validation:Optional
	AllocationMethod *string `json:"allocationMethod,omitempty" tf:"allocation_method,omitempty"`

	// The DDoS protection mode of the public IP. Possible values are Disabled, Enabled, and VirtualNetworkInherited. Defaults to VirtualNetworkInherited.
	// +kubebuilder:validation:Optional
	DDOSProtectionMode *string `json:"ddosProtectionMode,omitempty" tf:"ddos_protection_mode,omitempty"`

	// The ID of DDoS protection plan associated with the public IP.
	// +kubebuilder:validation:Optional
	DDOSProtectionPlanID *string `json:"ddosProtectionPlanId,omitempty" tf:"ddos_protection_plan_id,omitempty"`

	// Label for the Domain Name. Will be used to make up the FQDN. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	// +kubebuilder:validation:Optional
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Public IP should exist. Changing this forces a new Public IP to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// A mapping of IP tags to assign to the public IP. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	IPTags map[string]*string `json:"ipTags,omitempty" tf:"ip_tags,omitempty"`

	// The IP Version to use, IPv6 or IPv4. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// Specifies the supported Azure location where the Public IP should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// If specified then public IP address allocated will be provided from the public IP prefix resource. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`

	// The name of the Resource Group where this Public IP should exist. Changing this forces a new Public IP to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	// +kubebuilder:validation:Optional
	ReverseFqdn *string `json:"reverseFqdn,omitempty" tf:"reverse_fqdn,omitempty"`

	// The SKU of the Public IP. Accepted values are Basic and Standard. Defaults to Basic. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The SKU Tier that should be used for the Public IP. Possible values are Regional and Global. Defaults to Regional. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SkuTier *string `json:"skuTier,omitempty" tf:"sku_tier,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A collection containing the availability zone to allocate the Public IP in. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	// +listType=set
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

// PublicIPSpec defines the desired state of PublicIP
type PublicIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicIPParameters `json:"forProvider"`
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
	InitProvider PublicIPInitParameters `json:"initProvider,omitempty"`
}

// PublicIPStatus defines the observed state of PublicIP.
type PublicIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PublicIP is the Schema for the PublicIPs API. Manages a Public IP Address.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PublicIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.allocationMethod) || (has(self.initProvider) && has(self.initProvider.allocationMethod))",message="spec.forProvider.allocationMethod is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   PublicIPSpec   `json:"spec"`
	Status PublicIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPList contains a list of PublicIPs
type PublicIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIP `json:"items"`
}

// Repository type metadata.
var (
	PublicIP_Kind             = "PublicIP"
	PublicIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicIP_Kind}.String()
	PublicIP_KindAPIVersion   = PublicIP_Kind + "." + CRDGroupVersion.String()
	PublicIP_GroupVersionKind = CRDGroupVersion.WithKind(PublicIP_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicIP{}, &PublicIPList{})
}
