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

type MSSQLManagedInstanceFailoverGroupInitParameters struct {

	// The ID of the Azure SQL Managed Instance which will be replicated using a Managed Instance Failover Group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDRef *v1.Reference `json:"managedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDSelector *v1.Selector `json:"managedInstanceIdSelector,omitempty" tf:"-"`

	// The ID of the Azure SQL Managed Instance which will be replicated to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	PartnerManagedInstanceID *string `json:"partnerManagedInstanceId,omitempty" tf:"partner_managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate partnerManagedInstanceId.
	// +kubebuilder:validation:Optional
	PartnerManagedInstanceIDRef *v1.Reference `json:"partnerManagedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate partnerManagedInstanceId.
	// +kubebuilder:validation:Optional
	PartnerManagedInstanceIDSelector *v1.Selector `json:"partnerManagedInstanceIdSelector,omitempty" tf:"-"`

	// A read_write_endpoint_failover_policy block as defined below.
	ReadWriteEndpointFailoverPolicy []MSSQLManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyInitParameters `json:"readWriteEndpointFailoverPolicy,omitempty" tf:"read_write_endpoint_failover_policy,omitempty"`

	// Failover policy for the read-only endpoint. Defaults to true.
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`
}

type MSSQLManagedInstanceFailoverGroupObservation struct {

	// The ID of the Managed Instance Failover Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Managed Instance Failover Group should exist. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Azure SQL Managed Instance which will be replicated using a Managed Instance Failover Group. Changing this forces a new resource to be created.
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// The ID of the Azure SQL Managed Instance which will be replicated to. Changing this forces a new resource to be created.
	PartnerManagedInstanceID *string `json:"partnerManagedInstanceId,omitempty" tf:"partner_managed_instance_id,omitempty"`

	// A partner_region block as defined below.
	PartnerRegion []PartnerRegionObservation `json:"partnerRegion,omitempty" tf:"partner_region,omitempty"`

	// A read_write_endpoint_failover_policy block as defined below.
	ReadWriteEndpointFailoverPolicy []MSSQLManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyObservation `json:"readWriteEndpointFailoverPolicy,omitempty" tf:"read_write_endpoint_failover_policy,omitempty"`

	// Failover policy for the read-only endpoint. Defaults to true.
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`

	// The local replication role of the Managed Instance Failover Group.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type MSSQLManagedInstanceFailoverGroupParameters struct {

	// The Azure Region where the Managed Instance Failover Group should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The ID of the Azure SQL Managed Instance which will be replicated using a Managed Instance Failover Group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ManagedInstanceID *string `json:"managedInstanceId,omitempty" tf:"managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDRef *v1.Reference `json:"managedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate managedInstanceId.
	// +kubebuilder:validation:Optional
	ManagedInstanceIDSelector *v1.Selector `json:"managedInstanceIdSelector,omitempty" tf:"-"`

	// The ID of the Azure SQL Managed Instance which will be replicated to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLManagedInstance
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PartnerManagedInstanceID *string `json:"partnerManagedInstanceId,omitempty" tf:"partner_managed_instance_id,omitempty"`

	// Reference to a MSSQLManagedInstance to populate partnerManagedInstanceId.
	// +kubebuilder:validation:Optional
	PartnerManagedInstanceIDRef *v1.Reference `json:"partnerManagedInstanceIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLManagedInstance to populate partnerManagedInstanceId.
	// +kubebuilder:validation:Optional
	PartnerManagedInstanceIDSelector *v1.Selector `json:"partnerManagedInstanceIdSelector,omitempty" tf:"-"`

	// A read_write_endpoint_failover_policy block as defined below.
	// +kubebuilder:validation:Optional
	ReadWriteEndpointFailoverPolicy []MSSQLManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyParameters `json:"readWriteEndpointFailoverPolicy,omitempty" tf:"read_write_endpoint_failover_policy,omitempty"`

	// Failover policy for the read-only endpoint. Defaults to true.
	// +kubebuilder:validation:Optional
	ReadonlyEndpointFailoverPolicyEnabled *bool `json:"readonlyEndpointFailoverPolicyEnabled,omitempty" tf:"readonly_endpoint_failover_policy_enabled,omitempty"`
}

type MSSQLManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyInitParameters struct {

	// Applies only if mode is Automatic. The grace period in minutes before failover with data loss is attempted.
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// The failover mode. Possible values are Automatic or Manual.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type MSSQLManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyObservation struct {

	// Applies only if mode is Automatic. The grace period in minutes before failover with data loss is attempted.
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// The failover mode. Possible values are Automatic or Manual.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type MSSQLManagedInstanceFailoverGroupReadWriteEndpointFailoverPolicyParameters struct {

	// Applies only if mode is Automatic. The grace period in minutes before failover with data loss is attempted.
	// +kubebuilder:validation:Optional
	GraceMinutes *float64 `json:"graceMinutes,omitempty" tf:"grace_minutes,omitempty"`

	// The failover mode. Possible values are Automatic or Manual.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`
}

type PartnerRegionInitParameters struct {
}

type PartnerRegionObservation struct {

	// The Azure Region where the Managed Instance Failover Group partner exists.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The partner replication role of the Managed Instance Failover Group.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type PartnerRegionParameters struct {
}

// MSSQLManagedInstanceFailoverGroupSpec defines the desired state of MSSQLManagedInstanceFailoverGroup
type MSSQLManagedInstanceFailoverGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLManagedInstanceFailoverGroupParameters `json:"forProvider"`
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
	InitProvider MSSQLManagedInstanceFailoverGroupInitParameters `json:"initProvider,omitempty"`
}

// MSSQLManagedInstanceFailoverGroupStatus defines the observed state of MSSQLManagedInstanceFailoverGroup.
type MSSQLManagedInstanceFailoverGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLManagedInstanceFailoverGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MSSQLManagedInstanceFailoverGroup is the Schema for the MSSQLManagedInstanceFailoverGroups API. Manages an Azure SQL Managed Instance Failover Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLManagedInstanceFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.readWriteEndpointFailoverPolicy) || (has(self.initProvider) && has(self.initProvider.readWriteEndpointFailoverPolicy))",message="spec.forProvider.readWriteEndpointFailoverPolicy is a required parameter"
	Spec   MSSQLManagedInstanceFailoverGroupSpec   `json:"spec"`
	Status MSSQLManagedInstanceFailoverGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLManagedInstanceFailoverGroupList contains a list of MSSQLManagedInstanceFailoverGroups
type MSSQLManagedInstanceFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLManagedInstanceFailoverGroup `json:"items"`
}

// Repository type metadata.
var (
	MSSQLManagedInstanceFailoverGroup_Kind             = "MSSQLManagedInstanceFailoverGroup"
	MSSQLManagedInstanceFailoverGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLManagedInstanceFailoverGroup_Kind}.String()
	MSSQLManagedInstanceFailoverGroup_KindAPIVersion   = MSSQLManagedInstanceFailoverGroup_Kind + "." + CRDGroupVersion.String()
	MSSQLManagedInstanceFailoverGroup_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLManagedInstanceFailoverGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLManagedInstanceFailoverGroup{}, &MSSQLManagedInstanceFailoverGroupList{})
}
