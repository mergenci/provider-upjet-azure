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

type NotificationRecipientEmailInitParameters struct {
}

type NotificationRecipientEmailObservation struct {

	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
	APIManagementID *string `json:"apiManagementId,omitempty" tf:"api_management_id,omitempty"`

	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The ID of the API Management Notification Recipient Email.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are AccountClosedPublisher, BCC, NewApplicationNotificationMessage, NewIssuePublisherNotificationMessage, PurchasePublisherNotificationMessage, QuotaLimitApproachingPublisherNotificationMessage, and RequestPublisherNotificationMessage.
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`
}

type NotificationRecipientEmailParameters struct {

	// The ID of the API Management Service from which to create this Notification Recipient Email. Changing this forces a new API Management Notification Recipient Email to be created.
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

	// The recipient email address. Changing this forces a new API Management Notification Recipient Email to be created.
	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// The Notification Name to be received. Changing this forces a new API Management Notification Recipient Email to be created. Possible values are AccountClosedPublisher, BCC, NewApplicationNotificationMessage, NewIssuePublisherNotificationMessage, PurchasePublisherNotificationMessage, QuotaLimitApproachingPublisherNotificationMessage, and RequestPublisherNotificationMessage.
	// +kubebuilder:validation:Required
	NotificationType *string `json:"notificationType" tf:"notification_type,omitempty"`
}

// NotificationRecipientEmailSpec defines the desired state of NotificationRecipientEmail
type NotificationRecipientEmailSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationRecipientEmailParameters `json:"forProvider"`
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
	InitProvider NotificationRecipientEmailInitParameters `json:"initProvider,omitempty"`
}

// NotificationRecipientEmailStatus defines the observed state of NotificationRecipientEmail.
type NotificationRecipientEmailStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationRecipientEmailObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NotificationRecipientEmail is the Schema for the NotificationRecipientEmails API. Manages a API Management Notification Recipient Email.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NotificationRecipientEmail struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationRecipientEmailSpec   `json:"spec"`
	Status            NotificationRecipientEmailStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationRecipientEmailList contains a list of NotificationRecipientEmails
type NotificationRecipientEmailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotificationRecipientEmail `json:"items"`
}

// Repository type metadata.
var (
	NotificationRecipientEmail_Kind             = "NotificationRecipientEmail"
	NotificationRecipientEmail_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NotificationRecipientEmail_Kind}.String()
	NotificationRecipientEmail_KindAPIVersion   = NotificationRecipientEmail_Kind + "." + CRDGroupVersion.String()
	NotificationRecipientEmail_GroupVersionKind = CRDGroupVersion.WithKind(NotificationRecipientEmail_Kind)
)

func init() {
	SchemeBuilder.Register(&NotificationRecipientEmail{}, &NotificationRecipientEmailList{})
}
