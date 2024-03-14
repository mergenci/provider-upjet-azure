// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceDeploymentScriptAzureCli.
func (mg *ResourceDeploymentScriptAzureCli) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceDeploymentScriptAzurePowerShell.
func (mg *ResourceDeploymentScriptAzurePowerShell) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ResourceGroupTemplateDeployment.
func (mg *ResourceGroupTemplateDeployment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this SubscriptionTemplateDeployment.
func (mg *SubscriptionTemplateDeployment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
