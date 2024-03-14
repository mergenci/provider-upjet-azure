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

type ConditionAlertRuleNameInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionAlertRuleNameObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionAlertRuleNameParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorConditionInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionMonitorConditionObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionMonitorConditionParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionSignalTypeInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionSignalTypeObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionSignalTypeParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionTargetResourceGroupInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionTargetResourceGroupObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionTargetResourceGroupParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionTargetResourceInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionTargetResourceObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConditionTargetResourceParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionAlertContextInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionAlertContextObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionAlertContextParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionAlertRuleIDInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionAlertRuleIDObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionAlertRuleIDParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionDescriptionInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionDescriptionObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionDescriptionParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionInitParameters struct {

	// A alert_context block as defined above.
	AlertContext []MonitorAlertProcessingRuleSuppressionConditionAlertContextInitParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined above.
	AlertRuleID []MonitorAlertProcessingRuleSuppressionConditionAlertRuleIDInitParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A alert_rule_name block as defined above.
	AlertRuleName []ConditionAlertRuleNameInitParameters `json:"alertRuleName,omitempty" tf:"alert_rule_name,omitempty"`

	// A description block as defined below.
	Description []MonitorAlertProcessingRuleSuppressionConditionDescriptionInitParameters `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor_condition block as defined below.
	MonitorCondition []ConditionMonitorConditionInitParameters `json:"monitorCondition,omitempty" tf:"monitor_condition,omitempty"`

	// A monitor_service block as defined below.
	MonitorService []MonitorAlertProcessingRuleSuppressionConditionMonitorServiceInitParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	Severity []MonitorAlertProcessingRuleSuppressionConditionSeverityInitParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// A signal_type block as defined below.
	SignalType []ConditionSignalTypeInitParameters `json:"signalType,omitempty" tf:"signal_type,omitempty"`

	// A target_resource block as defined below.
	TargetResource []ConditionTargetResourceInitParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`

	// A target_resource_group block as defined below.
	TargetResourceGroup []ConditionTargetResourceGroupInitParameters `json:"targetResourceGroup,omitempty" tf:"target_resource_group,omitempty"`

	// A target_resource_type block as defined below.
	TargetResourceType []MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeInitParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionMonitorServiceInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionMonitorServiceObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionMonitorServiceParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionObservation struct {

	// A alert_context block as defined above.
	AlertContext []MonitorAlertProcessingRuleSuppressionConditionAlertContextObservation `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined above.
	AlertRuleID []MonitorAlertProcessingRuleSuppressionConditionAlertRuleIDObservation `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A alert_rule_name block as defined above.
	AlertRuleName []ConditionAlertRuleNameObservation `json:"alertRuleName,omitempty" tf:"alert_rule_name,omitempty"`

	// A description block as defined below.
	Description []MonitorAlertProcessingRuleSuppressionConditionDescriptionObservation `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor_condition block as defined below.
	MonitorCondition []ConditionMonitorConditionObservation `json:"monitorCondition,omitempty" tf:"monitor_condition,omitempty"`

	// A monitor_service block as defined below.
	MonitorService []MonitorAlertProcessingRuleSuppressionConditionMonitorServiceObservation `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	Severity []MonitorAlertProcessingRuleSuppressionConditionSeverityObservation `json:"severity,omitempty" tf:"severity,omitempty"`

	// A signal_type block as defined below.
	SignalType []ConditionSignalTypeObservation `json:"signalType,omitempty" tf:"signal_type,omitempty"`

	// A target_resource block as defined below.
	TargetResource []ConditionTargetResourceObservation `json:"targetResource,omitempty" tf:"target_resource,omitempty"`

	// A target_resource_group block as defined below.
	TargetResourceGroup []ConditionTargetResourceGroupObservation `json:"targetResourceGroup,omitempty" tf:"target_resource_group,omitempty"`

	// A target_resource_type block as defined below.
	TargetResourceType []MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeObservation `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionParameters struct {

	// A alert_context block as defined above.
	// +kubebuilder:validation:Optional
	AlertContext []MonitorAlertProcessingRuleSuppressionConditionAlertContextParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined above.
	// +kubebuilder:validation:Optional
	AlertRuleID []MonitorAlertProcessingRuleSuppressionConditionAlertRuleIDParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A alert_rule_name block as defined above.
	// +kubebuilder:validation:Optional
	AlertRuleName []ConditionAlertRuleNameParameters `json:"alertRuleName,omitempty" tf:"alert_rule_name,omitempty"`

	// A description block as defined below.
	// +kubebuilder:validation:Optional
	Description []MonitorAlertProcessingRuleSuppressionConditionDescriptionParameters `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor_condition block as defined below.
	// +kubebuilder:validation:Optional
	MonitorCondition []ConditionMonitorConditionParameters `json:"monitorCondition,omitempty" tf:"monitor_condition,omitempty"`

	// A monitor_service block as defined below.
	// +kubebuilder:validation:Optional
	MonitorService []MonitorAlertProcessingRuleSuppressionConditionMonitorServiceParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	// +kubebuilder:validation:Optional
	Severity []MonitorAlertProcessingRuleSuppressionConditionSeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// A signal_type block as defined below.
	// +kubebuilder:validation:Optional
	SignalType []ConditionSignalTypeParameters `json:"signalType,omitempty" tf:"signal_type,omitempty"`

	// A target_resource block as defined below.
	// +kubebuilder:validation:Optional
	TargetResource []ConditionTargetResourceParameters `json:"targetResource,omitempty" tf:"target_resource,omitempty"`

	// A target_resource_group block as defined below.
	// +kubebuilder:validation:Optional
	TargetResourceGroup []ConditionTargetResourceGroupParameters `json:"targetResourceGroup,omitempty" tf:"target_resource_group,omitempty"`

	// A target_resource_type block as defined below.
	// +kubebuilder:validation:Optional
	TargetResourceType []MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionSeverityInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionSeverityObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionSeverityParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeInitParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeObservation struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionConditionTargetResourceTypeParameters struct {

	// The operator for a given condition. Possible values are Equals, NotEquals, Contains, and DoesNotContain.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types. (e.g. Microsoft.Compute/VirtualMachines)
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionInitParameters struct {

	// A condition block as defined below.
	Condition []MonitorAlertProcessingRuleSuppressionConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Alert Processing Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the Alert Processing Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A schedule block as defined below.
	Schedule []MonitorAlertProcessingRuleSuppressionScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of resource IDs which will be the target of Alert Processing Rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// References to ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesRefs []v1.Reference `json:"scopesRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesSelector *v1.Selector `json:"scopesSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Alert Processing Rule.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionObservation struct {

	// A condition block as defined below.
	Condition []MonitorAlertProcessingRuleSuppressionConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Alert Processing Rule.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the Alert Processing Rule be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the Alert Processing Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Alert Processing Rule should exist. Changing this forces a new Alert Processing Rule to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A schedule block as defined below.
	Schedule []MonitorAlertProcessingRuleSuppressionScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of resource IDs which will be the target of Alert Processing Rule.
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// A mapping of tags which should be assigned to the Alert Processing Rule.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionParameters struct {

	// A condition block as defined below.
	// +kubebuilder:validation:Optional
	Condition []MonitorAlertProcessingRuleSuppressionConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Alert Processing Rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the Alert Processing Rule be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Resource Group where the Alert Processing Rule should exist. Changing this forces a new Alert Processing Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A schedule block as defined below.
	// +kubebuilder:validation:Optional
	Schedule []MonitorAlertProcessingRuleSuppressionScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// A list of resource IDs which will be the target of Alert Processing Rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// References to ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesRefs []v1.Reference `json:"scopesRefs,omitempty" tf:"-"`

	// Selector for a list of ResourceGroup in azure to populate scopes.
	// +kubebuilder:validation:Optional
	ScopesSelector *v1.Selector `json:"scopesSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Alert Processing Rule.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionScheduleInitParameters struct {

	// Specifies the Alert Processing Rule effective start time (Y-m-d'T'H:M:S).
	EffectiveFrom *string `json:"effectiveFrom,omitempty" tf:"effective_from,omitempty"`

	// Specifies the Alert Processing Rule effective end time (Y-m-d'T'H:M:S).
	EffectiveUntil *string `json:"effectiveUntil,omitempty" tf:"effective_until,omitempty"`

	// A recurrence block as defined above.
	Recurrence []ScheduleRecurrenceInitParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The time zone (e.g. Pacific Standard time, Eastern Standard Time). Defaults to UTC. possible values are defined here.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionScheduleObservation struct {

	// Specifies the Alert Processing Rule effective start time (Y-m-d'T'H:M:S).
	EffectiveFrom *string `json:"effectiveFrom,omitempty" tf:"effective_from,omitempty"`

	// Specifies the Alert Processing Rule effective end time (Y-m-d'T'H:M:S).
	EffectiveUntil *string `json:"effectiveUntil,omitempty" tf:"effective_until,omitempty"`

	// A recurrence block as defined above.
	Recurrence []ScheduleRecurrenceObservation `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The time zone (e.g. Pacific Standard time, Eastern Standard Time). Defaults to UTC. possible values are defined here.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type MonitorAlertProcessingRuleSuppressionScheduleParameters struct {

	// Specifies the Alert Processing Rule effective start time (Y-m-d'T'H:M:S).
	// +kubebuilder:validation:Optional
	EffectiveFrom *string `json:"effectiveFrom,omitempty" tf:"effective_from,omitempty"`

	// Specifies the Alert Processing Rule effective end time (Y-m-d'T'H:M:S).
	// +kubebuilder:validation:Optional
	EffectiveUntil *string `json:"effectiveUntil,omitempty" tf:"effective_until,omitempty"`

	// A recurrence block as defined above.
	// +kubebuilder:validation:Optional
	Recurrence []ScheduleRecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`

	// The time zone (e.g. Pacific Standard time, Eastern Standard Time). Defaults to UTC. possible values are defined here.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type RecurrenceDailyInitParameters struct {

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceDailyObservation struct {

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceDailyParameters struct {

	// Specifies the recurrence end time (H:M:S).
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type RecurrenceMonthlyInitParameters struct {

	// Specifies a list of dayOfMonth to recurrence. Possible values are integers between 1 - 31.
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceMonthlyObservation struct {

	// Specifies a list of dayOfMonth to recurrence. Possible values are integers between 1 - 31.
	DaysOfMonth []*float64 `json:"daysOfMonth,omitempty" tf:"days_of_month,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceMonthlyParameters struct {

	// Specifies a list of dayOfMonth to recurrence. Possible values are integers between 1 - 31.
	// +kubebuilder:validation:Optional
	DaysOfMonth []*float64 `json:"daysOfMonth" tf:"days_of_month,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceWeeklyInitParameters struct {

	// Specifies a list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday.
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceWeeklyObservation struct {

	// Specifies a list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday.
	DaysOfWeek []*string `json:"daysOfWeek,omitempty" tf:"days_of_week,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type RecurrenceWeeklyParameters struct {

	// Specifies a list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday.
	// +kubebuilder:validation:Optional
	DaysOfWeek []*string `json:"daysOfWeek" tf:"days_of_week,omitempty"`

	// Specifies the recurrence end time (H:M:S).
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the recurrence start time (H:M:S).
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ScheduleRecurrenceInitParameters struct {

	// One or more daily blocks as defined above.
	Daily []RecurrenceDailyInitParameters `json:"daily,omitempty" tf:"daily,omitempty"`

	// One or more monthly blocks as defined above.
	Monthly []RecurrenceMonthlyInitParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// One or more weekly blocks as defined below.
	Weekly []RecurrenceWeeklyInitParameters `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type ScheduleRecurrenceObservation struct {

	// One or more daily blocks as defined above.
	Daily []RecurrenceDailyObservation `json:"daily,omitempty" tf:"daily,omitempty"`

	// One or more monthly blocks as defined above.
	Monthly []RecurrenceMonthlyObservation `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// One or more weekly blocks as defined below.
	Weekly []RecurrenceWeeklyObservation `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

type ScheduleRecurrenceParameters struct {

	// One or more daily blocks as defined above.
	// +kubebuilder:validation:Optional
	Daily []RecurrenceDailyParameters `json:"daily,omitempty" tf:"daily,omitempty"`

	// One or more monthly blocks as defined above.
	// +kubebuilder:validation:Optional
	Monthly []RecurrenceMonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// One or more weekly blocks as defined below.
	// +kubebuilder:validation:Optional
	Weekly []RecurrenceWeeklyParameters `json:"weekly,omitempty" tf:"weekly,omitempty"`
}

// MonitorAlertProcessingRuleSuppressionSpec defines the desired state of MonitorAlertProcessingRuleSuppression
type MonitorAlertProcessingRuleSuppressionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorAlertProcessingRuleSuppressionParameters `json:"forProvider"`
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
	InitProvider MonitorAlertProcessingRuleSuppressionInitParameters `json:"initProvider,omitempty"`
}

// MonitorAlertProcessingRuleSuppressionStatus defines the observed state of MonitorAlertProcessingRuleSuppression.
type MonitorAlertProcessingRuleSuppressionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorAlertProcessingRuleSuppressionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MonitorAlertProcessingRuleSuppression is the Schema for the MonitorAlertProcessingRuleSuppressions API. Manages an Alert Processing Rule which suppress notifications.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorAlertProcessingRuleSuppression struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorAlertProcessingRuleSuppressionSpec   `json:"spec"`
	Status            MonitorAlertProcessingRuleSuppressionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAlertProcessingRuleSuppressionList contains a list of MonitorAlertProcessingRuleSuppressions
type MonitorAlertProcessingRuleSuppressionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorAlertProcessingRuleSuppression `json:"items"`
}

// Repository type metadata.
var (
	MonitorAlertProcessingRuleSuppression_Kind             = "MonitorAlertProcessingRuleSuppression"
	MonitorAlertProcessingRuleSuppression_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorAlertProcessingRuleSuppression_Kind}.String()
	MonitorAlertProcessingRuleSuppression_KindAPIVersion   = MonitorAlertProcessingRuleSuppression_Kind + "." + CRDGroupVersion.String()
	MonitorAlertProcessingRuleSuppression_GroupVersionKind = CRDGroupVersion.WithKind(MonitorAlertProcessingRuleSuppression_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorAlertProcessingRuleSuppression{}, &MonitorAlertProcessingRuleSuppressionList{})
}
