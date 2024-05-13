// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OutputServiceBusTopicInitParameters struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of property columns to add to the Service Bus Topic output.
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	Serialization *OutputServiceBusTopicSerializationInitParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta2.ServiceBusNamespace
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// Reference to a ServiceBusNamespace in servicebus to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceRef *v1.Reference `json:"servicebusNamespaceRef,omitempty" tf:"-"`

	// Selector for a ServiceBusNamespace in servicebus to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceSelector *v1.Selector `json:"servicebusNamespaceSelector,omitempty" tf:"-"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required if authentication_mode is ConnectionString.
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/streamanalytics/v1beta2.Job
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// A key-value pair of system property columns that will be attached to the outgoing messages for the Service Bus Topic Output.
	// +mapType=granular
	SystemPropertyColumns map[string]*string `json:"systemPropertyColumns,omitempty" tf:"system_property_columns,omitempty"`

	// The name of the Service Bus Topic.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.Topic
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Reference to a Topic in servicebus to populate topicName.
	// +kubebuilder:validation:Optional
	TopicNameRef *v1.Reference `json:"topicNameRef,omitempty" tf:"-"`

	// Selector for a Topic in servicebus to populate topicName.
	// +kubebuilder:validation:Optional
	TopicNameSelector *v1.Selector `json:"topicNameSelector,omitempty" tf:"-"`
}

type OutputServiceBusTopicObservation struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The ID of the Stream Analytics Output ServiceBus Topic.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Stream Output. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of property columns to add to the Service Bus Topic output.
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A serialization block as defined below.
	Serialization *OutputServiceBusTopicSerializationObservation `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required if authentication_mode is ConnectionString.
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// A key-value pair of system property columns that will be attached to the outgoing messages for the Service Bus Topic Output.
	// +mapType=granular
	SystemPropertyColumns map[string]*string `json:"systemPropertyColumns,omitempty" tf:"system_property_columns,omitempty"`

	// The name of the Service Bus Topic.
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

type OutputServiceBusTopicParameters struct {

	// The authentication mode for the Stream Output. Possible values are Msi and ConnectionString. Defaults to ConnectionString.
	// +kubebuilder:validation:Optional
	AuthenticationMode *string `json:"authenticationMode,omitempty" tf:"authentication_mode,omitempty"`

	// The name of the Stream Output. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of property columns to add to the Service Bus Topic output.
	// +kubebuilder:validation:Optional
	PropertyColumns []*string `json:"propertyColumns,omitempty" tf:"property_columns,omitempty"`

	// The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A serialization block as defined below.
	// +kubebuilder:validation:Optional
	Serialization *OutputServiceBusTopicSerializationParameters `json:"serialization,omitempty" tf:"serialization,omitempty"`

	// The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta2.ServiceBusNamespace
	// +kubebuilder:validation:Optional
	ServiceBusNamespace *string `json:"servicebusNamespace,omitempty" tf:"servicebus_namespace,omitempty"`

	// Reference to a ServiceBusNamespace in servicebus to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceRef *v1.Reference `json:"servicebusNamespaceRef,omitempty" tf:"-"`

	// Selector for a ServiceBusNamespace in servicebus to populate servicebusNamespace.
	// +kubebuilder:validation:Optional
	ServiceBusNamespaceSelector *v1.Selector `json:"servicebusNamespaceSelector,omitempty" tf:"-"`

	// The shared access policy key for the specified shared access policy. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyKeySecretRef *v1.SecretKeySelector `json:"sharedAccessPolicyKeySecretRef,omitempty" tf:"-"`

	// The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc. Required if authentication_mode is ConnectionString.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// The name of the Stream Analytics Job. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/streamanalytics/v1beta2.Job
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName,omitempty" tf:"stream_analytics_job_name,omitempty"`

	// Reference to a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameRef *v1.Reference `json:"streamAnalyticsJobNameRef,omitempty" tf:"-"`

	// Selector for a Job in streamanalytics to populate streamAnalyticsJobName.
	// +kubebuilder:validation:Optional
	StreamAnalyticsJobNameSelector *v1.Selector `json:"streamAnalyticsJobNameSelector,omitempty" tf:"-"`

	// A key-value pair of system property columns that will be attached to the outgoing messages for the Service Bus Topic Output.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	SystemPropertyColumns map[string]*string `json:"systemPropertyColumns,omitempty" tf:"system_property_columns,omitempty"`

	// The name of the Service Bus Topic.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.Topic
	// +kubebuilder:validation:Optional
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// Reference to a Topic in servicebus to populate topicName.
	// +kubebuilder:validation:Optional
	TopicNameRef *v1.Reference `json:"topicNameRef,omitempty" tf:"-"`

	// Selector for a Topic in servicebus to populate topicName.
	// +kubebuilder:validation:Optional
	TopicNameSelector *v1.Selector `json:"topicNameSelector,omitempty" tf:"-"`
}

type OutputServiceBusTopicSerializationInitParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OutputServiceBusTopicSerializationObservation struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OutputServiceBusTopicSerializationParameters struct {

	// The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to UTF8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are   (space), , (comma), 	 (tab), | (pipe) and ;.
	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// Specifies the format of the JSON the output will be written in. Possible values are Array and LineSeparated.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The serialization format used for outgoing data streams. Possible values are Avro, Csv, Json and Parquet.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// OutputServiceBusTopicSpec defines the desired state of OutputServiceBusTopic
type OutputServiceBusTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputServiceBusTopicParameters `json:"forProvider"`
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
	InitProvider OutputServiceBusTopicInitParameters `json:"initProvider,omitempty"`
}

// OutputServiceBusTopicStatus defines the observed state of OutputServiceBusTopic.
type OutputServiceBusTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputServiceBusTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OutputServiceBusTopic is the Schema for the OutputServiceBusTopics API. Manages a Stream Analytics Output to a ServiceBus Topic.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type OutputServiceBusTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serialization) || (has(self.initProvider) && has(self.initProvider.serialization))",message="spec.forProvider.serialization is a required parameter"
	Spec   OutputServiceBusTopicSpec   `json:"spec"`
	Status OutputServiceBusTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputServiceBusTopicList contains a list of OutputServiceBusTopics
type OutputServiceBusTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputServiceBusTopic `json:"items"`
}

// Repository type metadata.
var (
	OutputServiceBusTopic_Kind             = "OutputServiceBusTopic"
	OutputServiceBusTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputServiceBusTopic_Kind}.String()
	OutputServiceBusTopic_KindAPIVersion   = OutputServiceBusTopic_Kind + "." + CRDGroupVersion.String()
	OutputServiceBusTopic_GroupVersionKind = CRDGroupVersion.WithKind(OutputServiceBusTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputServiceBusTopic{}, &OutputServiceBusTopicList{})
}
