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

type BudgetResourceGroupFilterInitParameters struct {

	// One or more dimension blocks as defined below to filter the budget on.
	Dimension []FilterDimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	Not []FilterNotInitParameters `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	Tag []BudgetResourceGroupFilterTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetResourceGroupFilterObservation struct {

	// One or more dimension blocks as defined below to filter the budget on.
	Dimension []FilterDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	Not []FilterNotObservation `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	Tag []BudgetResourceGroupFilterTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetResourceGroupFilterParameters struct {

	// One or more dimension blocks as defined below to filter the budget on.
	// +kubebuilder:validation:Optional
	Dimension []FilterDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	// +kubebuilder:validation:Optional
	Not []FilterNotParameters `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	// +kubebuilder:validation:Optional
	Tag []BudgetResourceGroupFilterTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetResourceGroupFilterTagInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetResourceGroupFilterTagObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetResourceGroupFilterTagParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type BudgetResourceGroupInitParameters struct {

	// The total amount of cost to track with the budget.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Resource Group Consumption Budget
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	Filter []BudgetResourceGroupFilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more notification blocks as defined below.
	Notification []BudgetResourceGroupNotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDRef *v1.Reference `json:"resourceGroupIdRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDSelector *v1.Selector `json:"resourceGroupIdSelector,omitempty" tf:"-"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	TimePeriod []BudgetResourceGroupTimePeriodInitParameters `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetResourceGroupNotificationInitParameters struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Specifies a list of Action Group IDs to send the budget notification to when the threshold is exceeded.
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// Specifies a list of contact roles to send the budget notification to when the threshold is exceeded.
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// Should the notification be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type BudgetResourceGroupNotificationObservation struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Specifies a list of Action Group IDs to send the budget notification to when the threshold is exceeded.
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// Specifies a list of contact roles to send the budget notification to when the threshold is exceeded.
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// Should the notification be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type BudgetResourceGroupNotificationParameters struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	// +kubebuilder:validation:Optional
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Specifies a list of Action Group IDs to send the budget notification to when the threshold is exceeded.
	// +kubebuilder:validation:Optional
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// Specifies a list of contact roles to send the budget notification to when the threshold is exceeded.
	// +kubebuilder:validation:Optional
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// Should the notification be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type BudgetResourceGroupObservation struct {

	// The total amount of cost to track with the budget.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Resource Group Consumption Budget
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	Filter []BudgetResourceGroupFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the Resource Group Consumption Budget.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more notification blocks as defined below.
	Notification []BudgetResourceGroupNotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	TimePeriod []BudgetResourceGroupTimePeriodObservation `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetResourceGroupParameters struct {

	// The total amount of cost to track with the budget.
	// +kubebuilder:validation:Optional
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Resource Group Consumption Budget
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	// +kubebuilder:validation:Optional
	Filter []BudgetResourceGroupFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// The name which should be used for this Resource Group Consumption Budget. Changing this forces a new Resource Group Consumption Budget to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// One or more notification blocks as defined below.
	// +kubebuilder:validation:Optional
	Notification []BudgetResourceGroupNotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the Resource Group to create the consumption budget for in the form of /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1. Changing this forces a new Resource Group Consumption Budget to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDRef *v1.Reference `json:"resourceGroupIdRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDSelector *v1.Selector `json:"resourceGroupIdSelector,omitempty" tf:"-"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	// +kubebuilder:validation:Optional
	TimePeriod []BudgetResourceGroupTimePeriodParameters `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetResourceGroupTimePeriodInitParameters struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Resource Group Consumption Budget to be created.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type BudgetResourceGroupTimePeriodObservation struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Resource Group Consumption Budget to be created.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type BudgetResourceGroupTimePeriodParameters struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Resource Group Consumption Budget to be created.
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate" tf:"start_date,omitempty"`
}

type FilterDimensionInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterDimensionObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterDimensionParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type FilterNotDimensionInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterNotDimensionObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterNotDimensionParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type FilterNotInitParameters struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	Dimension []FilterNotDimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	Tag []NotTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FilterNotObservation struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	Dimension []FilterNotDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	Tag []NotTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type FilterNotParameters struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	// +kubebuilder:validation:Optional
	Dimension []FilterNotDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	// +kubebuilder:validation:Optional
	Tag []NotTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type NotTagInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NotTagObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type NotTagParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

// BudgetResourceGroupSpec defines the desired state of BudgetResourceGroup
type BudgetResourceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetResourceGroupParameters `json:"forProvider"`
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
	InitProvider BudgetResourceGroupInitParameters `json:"initProvider,omitempty"`
}

// BudgetResourceGroupStatus defines the observed state of BudgetResourceGroup.
type BudgetResourceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetResourceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BudgetResourceGroup is the Schema for the BudgetResourceGroups API. Manages a Resource Group Consumption Budget.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BudgetResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.amount) || (has(self.initProvider) && has(self.initProvider.amount))",message="spec.forProvider.amount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notification) || (has(self.initProvider) && has(self.initProvider.notification))",message="spec.forProvider.notification is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timePeriod) || (has(self.initProvider) && has(self.initProvider.timePeriod))",message="spec.forProvider.timePeriod is a required parameter"
	Spec   BudgetResourceGroupSpec   `json:"spec"`
	Status BudgetResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetResourceGroupList contains a list of BudgetResourceGroups
type BudgetResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetResourceGroup `json:"items"`
}

// Repository type metadata.
var (
	BudgetResourceGroup_Kind             = "BudgetResourceGroup"
	BudgetResourceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BudgetResourceGroup_Kind}.String()
	BudgetResourceGroup_KindAPIVersion   = BudgetResourceGroup_Kind + "." + CRDGroupVersion.String()
	BudgetResourceGroup_GroupVersionKind = CRDGroupVersion.WithKind(BudgetResourceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&BudgetResourceGroup{}, &BudgetResourceGroupList{})
}
