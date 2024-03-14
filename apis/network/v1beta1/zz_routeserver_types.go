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

type RouteServerInitParameters struct {

	// Whether to enable route exchange between Azure Route Server and the gateway(s)
	BranchToBranchTrafficEnabled *bool `json:"branchToBranchTrafficEnabled,omitempty" tf:"branch_to_branch_traffic_enabled,omitempty"`

	// Specifies the supported Azure location where the Route Server should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Public IP Address. This option is required since September 1st 2021. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// The SKU of the Route Server. The only possible value is Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The ID of the Subnet that the Route Server will reside. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RouteServerObservation struct {

	// Whether to enable route exchange between Azure Route Server and the gateway(s)
	BranchToBranchTrafficEnabled *bool `json:"branchToBranchTrafficEnabled,omitempty" tf:"branch_to_branch_traffic_enabled,omitempty"`

	// The ID of the Route Server .
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the supported Azure location where the Route Server should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Public IP Address. This option is required since September 1st 2021. Changing this forces a new resource to be created.
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Specifies the name of the Resource Group where the Route Server should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	RoutingState *string `json:"routingState,omitempty" tf:"routing_state,omitempty"`

	// The SKU of the Route Server. The only possible value is Standard. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The ID of the Subnet that the Route Server will reside. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	VirtualRouterAsn *float64 `json:"virtualRouterAsn,omitempty" tf:"virtual_router_asn,omitempty"`

	// +listType=set
	VirtualRouterIps []*string `json:"virtualRouterIps,omitempty" tf:"virtual_router_ips,omitempty"`
}

type RouteServerParameters struct {

	// Whether to enable route exchange between Azure Route Server and the gateway(s)
	// +kubebuilder:validation:Optional
	BranchToBranchTrafficEnabled *bool `json:"branchToBranchTrafficEnabled,omitempty" tf:"branch_to_branch_traffic_enabled,omitempty"`

	// Specifies the supported Azure location where the Route Server should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Public IP Address. This option is required since September 1st 2021. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PublicIP
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id,omitempty"`

	// Reference to a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDRef *v1.Reference `json:"publicIpAddressIdRef,omitempty" tf:"-"`

	// Selector for a PublicIP in network to populate publicIpAddressId.
	// +kubebuilder:validation:Optional
	PublicIPAddressIDSelector *v1.Selector `json:"publicIpAddressIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Resource Group where the Route Server should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Route Server. The only possible value is Standard. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// The ID of the Subnet that the Route Server will reside. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RouteServerSpec defines the desired state of RouteServer
type RouteServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteServerParameters `json:"forProvider"`
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
	InitProvider RouteServerInitParameters `json:"initProvider,omitempty"`
}

// RouteServerStatus defines the observed state of RouteServer.
type RouteServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouteServer is the Schema for the RouteServers API. Manages an Azure Route Server
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RouteServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   RouteServerSpec   `json:"spec"`
	Status RouteServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteServerList contains a list of RouteServers
type RouteServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteServer `json:"items"`
}

// Repository type metadata.
var (
	RouteServer_Kind             = "RouteServer"
	RouteServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteServer_Kind}.String()
	RouteServer_KindAPIVersion   = RouteServer_Kind + "." + CRDGroupVersion.String()
	RouteServer_GroupVersionKind = CRDGroupVersion.WithKind(RouteServer_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteServer{}, &RouteServerList{})
}
