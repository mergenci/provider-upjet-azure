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

type CustomRuleInitParameters struct {

	// The action to perform when the rule is matched. Possible values are Allow, Block, Log, or Redirect.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Is the rule is enabled or disabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more match_condition block defined below. Can support up to 10 match_condition blocks.
	MatchCondition []MatchConditionInitParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the rule. Rules with a lower value will be evaluated before rules with a higher value. Defaults to 1.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The rate limit duration in minutes. Defaults to 1.
	RateLimitDurationInMinutes *float64 `json:"rateLimitDurationInMinutes,omitempty" tf:"rate_limit_duration_in_minutes,omitempty"`

	// The rate limit threshold. Defaults to 10.
	RateLimitThreshold *float64 `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`

	// The type of rule. Possible values are MatchRule or RateLimitRule.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CustomRuleObservation struct {

	// The action to perform when the rule is matched. Possible values are Allow, Block, Log, or Redirect.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Is the rule is enabled or disabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more match_condition block defined below. Can support up to 10 match_condition blocks.
	MatchCondition []MatchConditionObservation `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the rule. Rules with a lower value will be evaluated before rules with a higher value. Defaults to 1.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The rate limit duration in minutes. Defaults to 1.
	RateLimitDurationInMinutes *float64 `json:"rateLimitDurationInMinutes,omitempty" tf:"rate_limit_duration_in_minutes,omitempty"`

	// The rate limit threshold. Defaults to 10.
	RateLimitThreshold *float64 `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`

	// The type of rule. Possible values are MatchRule or RateLimitRule.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type CustomRuleParameters struct {

	// The action to perform when the rule is matched. Possible values are Allow, Block, Log, or Redirect.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// Is the rule is enabled or disabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more match_condition block defined below. Can support up to 10 match_condition blocks.
	// +kubebuilder:validation:Optional
	MatchCondition []MatchConditionParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// Gets name of the resource that is unique within a policy. This name can be used to access the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The priority of the rule. Rules with a lower value will be evaluated before rules with a higher value. Defaults to 1.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The rate limit duration in minutes. Defaults to 1.
	// +kubebuilder:validation:Optional
	RateLimitDurationInMinutes *float64 `json:"rateLimitDurationInMinutes,omitempty" tf:"rate_limit_duration_in_minutes,omitempty"`

	// The rate limit threshold. Defaults to 10.
	// +kubebuilder:validation:Optional
	RateLimitThreshold *float64 `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`

	// The type of rule. Possible values are MatchRule or RateLimitRule.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type FrontdoorFirewallPolicyInitParameters struct {

	// If a custom_rule block's action type is block, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty" tf:"custom_block_response_body,omitempty"`

	// If a custom_rule block's action type is block, this is the response status code. Possible values are 200, 403, 405, 406, or 429.
	CustomBlockResponseStatusCode *float64 `json:"customBlockResponseStatusCode,omitempty" tf:"custom_block_response_status_code,omitempty"`

	// One or more custom_rule blocks as defined below.
	CustomRule []CustomRuleInitParameters `json:"customRule,omitempty" tf:"custom_rule,omitempty"`

	// Is the policy a enabled state or disabled state. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more managed_rule blocks as defined below.
	ManagedRule []ManagedRuleInitParameters `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`

	// The firewall policy mode. Possible values are Detection, Prevention and defaults to Prevention.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// If action type is redirect, this field represents redirect URL for the client.
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontdoorFirewallPolicyObservation struct {

	// If a custom_rule block's action type is block, this is the response body. The body must be specified in base64 encoding.
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty" tf:"custom_block_response_body,omitempty"`

	// If a custom_rule block's action type is block, this is the response status code. Possible values are 200, 403, 405, 406, or 429.
	CustomBlockResponseStatusCode *float64 `json:"customBlockResponseStatusCode,omitempty" tf:"custom_block_response_status_code,omitempty"`

	// One or more custom_rule blocks as defined below.
	CustomRule []CustomRuleObservation `json:"customRule,omitempty" tf:"custom_rule,omitempty"`

	// Is the policy a enabled state or disabled state. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The Frontend Endpoints associated with this Front Door Web Application Firewall policy.
	FrontendEndpointIds []*string `json:"frontendEndpointIds,omitempty" tf:"frontend_endpoint_ids,omitempty"`

	// The ID of the Front Door Firewall Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where this Front Door Firewall Policy exists.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more managed_rule blocks as defined below.
	ManagedRule []ManagedRuleObservation `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`

	// The firewall policy mode. Possible values are Detection, Prevention and defaults to Prevention.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// If action type is redirect, this field represents redirect URL for the client.
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontdoorFirewallPolicyParameters struct {

	// If a custom_rule block's action type is block, this is the response body. The body must be specified in base64 encoding.
	// +kubebuilder:validation:Optional
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty" tf:"custom_block_response_body,omitempty"`

	// If a custom_rule block's action type is block, this is the response status code. Possible values are 200, 403, 405, 406, or 429.
	// +kubebuilder:validation:Optional
	CustomBlockResponseStatusCode *float64 `json:"customBlockResponseStatusCode,omitempty" tf:"custom_block_response_status_code,omitempty"`

	// One or more custom_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomRule []CustomRuleParameters `json:"customRule,omitempty" tf:"custom_rule,omitempty"`

	// Is the policy a enabled state or disabled state. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more managed_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	ManagedRule []ManagedRuleParameters `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`

	// The firewall policy mode. Possible values are Detection, Prevention and defaults to Prevention.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// If action type is redirect, this field represents redirect URL for the client.
	// +kubebuilder:validation:Optional
	RedirectURL *string `json:"redirectUrl,omitempty" tf:"redirect_url,omitempty"`

	// The name of the resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the Web Application Firewall Policy.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagedRuleExclusionInitParameters struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type ManagedRuleExclusionObservation struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type ManagedRuleExclusionParameters struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	// +kubebuilder:validation:Optional
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

type ManagedRuleInitParameters struct {

	// One or more exclusion blocks as defined below.
	Exclusion []ManagedRuleExclusionInitParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more override blocks as defined below.
	Override []OverrideInitParameters `json:"override,omitempty" tf:"override,omitempty"`

	// The name of the managed rule to use with this resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The version on the managed rule to use with this resource.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedRuleObservation struct {

	// One or more exclusion blocks as defined below.
	Exclusion []ManagedRuleExclusionObservation `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more override blocks as defined below.
	Override []OverrideObservation `json:"override,omitempty" tf:"override,omitempty"`

	// The name of the managed rule to use with this resource.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The version on the managed rule to use with this resource.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedRuleParameters struct {

	// One or more exclusion blocks as defined below.
	// +kubebuilder:validation:Optional
	Exclusion []ManagedRuleExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more override blocks as defined below.
	// +kubebuilder:validation:Optional
	Override []OverrideParameters `json:"override,omitempty" tf:"override,omitempty"`

	// The name of the managed rule to use with this resource.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The version on the managed rule to use with this resource.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type MatchConditionInitParameters struct {

	// Up to 600 possible values to match. Limit is in total across all match_condition blocks and match_values arguments. String value itself can be up to 256 characters long.
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Should the result of the condition be negated.
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// Up to 5 transforms to apply. Possible values are Lowercase, RemoveNulls, Trim, Uppercase, URLDecode orURLEncode.
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchConditionObservation struct {

	// Up to 600 possible values to match. Limit is in total across all match_condition blocks and match_values arguments. String value itself can be up to 256 characters long.
	MatchValues []*string `json:"matchValues,omitempty" tf:"match_values,omitempty"`

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Should the result of the condition be negated.
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// Up to 5 transforms to apply. Possible values are Lowercase, RemoveNulls, Trim, Uppercase, URLDecode orURLEncode.
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchConditionParameters struct {

	// Up to 600 possible values to match. Limit is in total across all match_condition blocks and match_values arguments. String value itself can be up to 256 characters long.
	// +kubebuilder:validation:Optional
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	// +kubebuilder:validation:Optional
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Should the result of the condition be negated.
	// +kubebuilder:validation:Optional
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// Up to 5 transforms to apply. Possible values are Lowercase, RemoveNulls, Trim, Uppercase, URLDecode orURLEncode.
	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type OverrideExclusionInitParameters struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type OverrideExclusionObservation struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type OverrideExclusionParameters struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	// +kubebuilder:validation:Optional
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

type OverrideInitParameters struct {

	// One or more exclusion blocks as defined below.
	Exclusion []OverrideExclusionInitParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more rule blocks as defined below. If none are specified, all of the rules in the group will be disabled.
	Rule []OverrideRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The managed rule group to override.
	RuleGroupName *string `json:"ruleGroupName,omitempty" tf:"rule_group_name,omitempty"`
}

type OverrideObservation struct {

	// One or more exclusion blocks as defined below.
	Exclusion []OverrideExclusionObservation `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more rule blocks as defined below. If none are specified, all of the rules in the group will be disabled.
	Rule []OverrideRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`

	// The managed rule group to override.
	RuleGroupName *string `json:"ruleGroupName,omitempty" tf:"rule_group_name,omitempty"`
}

type OverrideParameters struct {

	// One or more exclusion blocks as defined below.
	// +kubebuilder:validation:Optional
	Exclusion []OverrideExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// One or more rule blocks as defined below. If none are specified, all of the rules in the group will be disabled.
	// +kubebuilder:validation:Optional
	Rule []OverrideRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`

	// The managed rule group to override.
	// +kubebuilder:validation:Optional
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

type OverrideRuleInitParameters struct {

	// The action to be applied when the rule matches. Possible values are Allow, Block, Log, or Redirect.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Is the managed rule override enabled or disabled. Defaults to false
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more exclusion blocks as defined below.
	Exclusion []RuleExclusionInitParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// Identifier for the managed rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type OverrideRuleObservation struct {

	// The action to be applied when the rule matches. Possible values are Allow, Block, Log, or Redirect.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Is the managed rule override enabled or disabled. Defaults to false
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more exclusion blocks as defined below.
	Exclusion []RuleExclusionObservation `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// Identifier for the managed rule.
	RuleID *string `json:"ruleId,omitempty" tf:"rule_id,omitempty"`
}

type OverrideRuleParameters struct {

	// The action to be applied when the rule matches. Possible values are Allow, Block, Log, or Redirect.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// Is the managed rule override enabled or disabled. Defaults to false
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// One or more exclusion blocks as defined below.
	// +kubebuilder:validation:Optional
	Exclusion []RuleExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// Identifier for the managed rule.
	// +kubebuilder:validation:Optional
	RuleID *string `json:"ruleId" tf:"rule_id,omitempty"`
}

type RuleExclusionInitParameters struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type RuleExclusionObservation struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	MatchVariable *string `json:"matchVariable,omitempty" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`
}

type RuleExclusionParameters struct {

	// The request variable to compare with. Possible values are Cookies, PostArgs, QueryString, RemoteAddr, RequestBody, RequestHeader, RequestMethod, RequestUri, or SocketAddr.
	// +kubebuilder:validation:Optional
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// Comparison type to use for matching with the variable value. Possible values are Any, BeginsWith, Contains, EndsWith, Equal, GeoMatch, GreaterThan, GreaterThanOrEqual, IPMatch, LessThan, LessThanOrEqual or RegEx.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Match against a specific key if the match_variable is QueryString, PostArgs, RequestHeader or Cookies.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector" tf:"selector,omitempty"`
}

// FrontdoorFirewallPolicySpec defines the desired state of FrontdoorFirewallPolicy
type FrontdoorFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorFirewallPolicyParameters `json:"forProvider"`
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
	InitProvider FrontdoorFirewallPolicyInitParameters `json:"initProvider,omitempty"`
}

// FrontdoorFirewallPolicyStatus defines the observed state of FrontdoorFirewallPolicy.
type FrontdoorFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FrontdoorFirewallPolicy is the Schema for the FrontdoorFirewallPolicys API. Manages an Azure Front Door (classic) Web Application Firewall Policy instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorFirewallPolicySpec   `json:"spec"`
	Status            FrontdoorFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorFirewallPolicyList contains a list of FrontdoorFirewallPolicys
type FrontdoorFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorFirewallPolicy_Kind             = "FrontdoorFirewallPolicy"
	FrontdoorFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorFirewallPolicy_Kind}.String()
	FrontdoorFirewallPolicy_KindAPIVersion   = FrontdoorFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FrontdoorFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorFirewallPolicy{}, &FrontdoorFirewallPolicyList{})
}
