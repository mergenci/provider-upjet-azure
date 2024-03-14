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

type SentinelAlertRuleMSSecurityIncidentInitParameters struct {

	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Only create incidents when the alert display name doesn't contain text from this list.
	// +listType=set
	DisplayNameExcludeFilter []*string `json:"displayNameExcludeFilter,omitempty" tf:"display_name_exclude_filter,omitempty"`

	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	// +listType=set
	DisplayNameFilter []*string `json:"displayNameFilter,omitempty" tf:"display_name_filter,omitempty"`

	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Microsoft Security Service from where the alert will be generated. Possible values are Azure Active Directory Identity Protection, Azure Advanced Threat Protection, Azure Security Center, Azure Security Center for IoT, Microsoft Cloud App Security, Microsoft Defender Advanced Threat Protection and Office 365 Advanced Threat Protection.
	ProductFilter *string `json:"productFilter,omitempty" tf:"product_filter,omitempty"`

	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are High, Medium, Low and Informational.
	// +listType=set
	SeverityFilter []*string `json:"severityFilter,omitempty" tf:"severity_filter,omitempty"`
}

type SentinelAlertRuleMSSecurityIncidentObservation struct {

	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// The description of this Sentinel MS Security Incident Alert Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Only create incidents when the alert display name doesn't contain text from this list.
	// +listType=set
	DisplayNameExcludeFilter []*string `json:"displayNameExcludeFilter,omitempty" tf:"display_name_exclude_filter,omitempty"`

	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	// +listType=set
	DisplayNameFilter []*string `json:"displayNameFilter,omitempty" tf:"display_name_filter,omitempty"`

	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Sentinel MS Security Incident Alert Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// The Microsoft Security Service from where the alert will be generated. Possible values are Azure Active Directory Identity Protection, Azure Advanced Threat Protection, Azure Security Center, Azure Security Center for IoT, Microsoft Cloud App Security, Microsoft Defender Advanced Threat Protection and Office 365 Advanced Threat Protection.
	ProductFilter *string `json:"productFilter,omitempty" tf:"product_filter,omitempty"`

	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are High, Medium, Low and Informational.
	// +listType=set
	SeverityFilter []*string `json:"severityFilter,omitempty" tf:"severity_filter,omitempty"`
}

type SentinelAlertRuleMSSecurityIncidentParameters struct {

	// The GUID of the alert rule template which is used to create this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	// +kubebuilder:validation:Optional
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// The description of this Sentinel MS Security Incident Alert Rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The friendly name of this Sentinel MS Security Incident Alert Rule.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Only create incidents when the alert display name doesn't contain text from this list.
	// +kubebuilder:validation:Optional
	// +listType=set
	DisplayNameExcludeFilter []*string `json:"displayNameExcludeFilter,omitempty" tf:"display_name_exclude_filter,omitempty"`

	// Only create incidents when the alert display name contain text from this list, leave empty to apply no filter.
	// +kubebuilder:validation:Optional
	// +listType=set
	DisplayNameFilter []*string `json:"displayNameFilter,omitempty" tf:"display_name_filter,omitempty"`

	// Should this Sentinel MS Security Incident Alert Rule be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Log Analytics Workspace this Sentinel MS Security Incident Alert Rule belongs to. Changing this forces a new Sentinel MS Security Incident Alert Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/securityinsights/v1beta1.SentinelLogAnalyticsWorkspaceOnboarding
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("workspace_id",false)
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// Reference to a SentinelLogAnalyticsWorkspaceOnboarding in securityinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDRef *v1.Reference `json:"logAnalyticsWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a SentinelLogAnalyticsWorkspaceOnboarding in securityinsights to populate logAnalyticsWorkspaceId.
	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceIDSelector *v1.Selector `json:"logAnalyticsWorkspaceIdSelector,omitempty" tf:"-"`

	// The Microsoft Security Service from where the alert will be generated. Possible values are Azure Active Directory Identity Protection, Azure Advanced Threat Protection, Azure Security Center, Azure Security Center for IoT, Microsoft Cloud App Security, Microsoft Defender Advanced Threat Protection and Office 365 Advanced Threat Protection.
	// +kubebuilder:validation:Optional
	ProductFilter *string `json:"productFilter,omitempty" tf:"product_filter,omitempty"`

	// Only create incidents from alerts when alert severity level is contained in this list. Possible values are High, Medium, Low and Informational.
	// +kubebuilder:validation:Optional
	// +listType=set
	SeverityFilter []*string `json:"severityFilter,omitempty" tf:"severity_filter,omitempty"`
}

// SentinelAlertRuleMSSecurityIncidentSpec defines the desired state of SentinelAlertRuleMSSecurityIncident
type SentinelAlertRuleMSSecurityIncidentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleMSSecurityIncidentParameters `json:"forProvider"`
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
	InitProvider SentinelAlertRuleMSSecurityIncidentInitParameters `json:"initProvider,omitempty"`
}

// SentinelAlertRuleMSSecurityIncidentStatus defines the observed state of SentinelAlertRuleMSSecurityIncident.
type SentinelAlertRuleMSSecurityIncidentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleMSSecurityIncidentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SentinelAlertRuleMSSecurityIncident is the Schema for the SentinelAlertRuleMSSecurityIncidents API. Manages a Sentinel MS Security Incident Alert Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAlertRuleMSSecurityIncident struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.productFilter) || (has(self.initProvider) && has(self.initProvider.productFilter))",message="spec.forProvider.productFilter is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.severityFilter) || (has(self.initProvider) && has(self.initProvider.severityFilter))",message="spec.forProvider.severityFilter is a required parameter"
	Spec   SentinelAlertRuleMSSecurityIncidentSpec   `json:"spec"`
	Status SentinelAlertRuleMSSecurityIncidentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMSSecurityIncidentList contains a list of SentinelAlertRuleMSSecurityIncidents
type SentinelAlertRuleMSSecurityIncidentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleMSSecurityIncident `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleMSSecurityIncident_Kind             = "SentinelAlertRuleMSSecurityIncident"
	SentinelAlertRuleMSSecurityIncident_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelAlertRuleMSSecurityIncident_Kind}.String()
	SentinelAlertRuleMSSecurityIncident_KindAPIVersion   = SentinelAlertRuleMSSecurityIncident_Kind + "." + CRDGroupVersion.String()
	SentinelAlertRuleMSSecurityIncident_GroupVersionKind = CRDGroupVersion.WithKind(SentinelAlertRuleMSSecurityIncident_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleMSSecurityIncident{}, &SentinelAlertRuleMSSecurityIncidentList{})
}
