// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *HealthcareDICOMService) ResolveReferences( // ResolveReferences of this HealthcareDICOMService.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("healthcareapis.azure.upbound.io", "v1beta1", "HealthcareWorkspace", "HealthcareWorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HealthcareFHIRService.
func (mg *HealthcareFHIRService) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("healthcareapis.azure.upbound.io", "v1beta1", "HealthcareWorkspace", "HealthcareWorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HealthcareMedtechService.
func (mg *HealthcareMedtechService) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "ConsumerGroup", "ConsumerGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubConsumerGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventHubConsumerGroupNameRef,
			Selector:     mg.Spec.ForProvider.EventHubConsumerGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubConsumerGroupName")
	}
	mg.Spec.ForProvider.EventHubConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubConsumerGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta2", "EventHub", "EventHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventHubNameRef,
			Selector:     mg.Spec.ForProvider.EventHubNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubName")
	}
	mg.Spec.ForProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta2", "EventHubNamespace", "EventHubNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubNamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventHubNamespaceNameRef,
			Selector:     mg.Spec.ForProvider.EventHubNamespaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubNamespaceName")
	}
	mg.Spec.ForProvider.EventHubNamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubNamespaceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("healthcareapis.azure.upbound.io", "v1beta1", "HealthcareWorkspace", "HealthcareWorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta1", "ConsumerGroup", "ConsumerGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventHubConsumerGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventHubConsumerGroupNameRef,
			Selector:     mg.Spec.InitProvider.EventHubConsumerGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventHubConsumerGroupName")
	}
	mg.Spec.InitProvider.EventHubConsumerGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventHubConsumerGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta2", "EventHub", "EventHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventHubNameRef,
			Selector:     mg.Spec.InitProvider.EventHubNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventHubName")
	}
	mg.Spec.InitProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventHubNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("eventhub.azure.upbound.io", "v1beta2", "EventHubNamespace", "EventHubNamespaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventHubNamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventHubNamespaceNameRef,
			Selector:     mg.Spec.InitProvider.EventHubNamespaceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.EventHubNamespaceName")
	}
	mg.Spec.InitProvider.EventHubNamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.EventHubNamespaceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this HealthcareService.
func (mg *HealthcareService) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
