// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta14 "github.com/upbound/provider-azure/apis/eventhub/v1beta1"
	v1beta12 "github.com/upbound/provider-azure/apis/keyvault/v1beta1"
	v1beta13 "github.com/upbound/provider-azure/apis/managedidentity/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/network/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this NetworkACL.
func (mg *NetworkACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrivateEndpoint); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateEndpoint[i3].ID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrivateEndpoint[i3].IDRef,
			Selector:     mg.Spec.ForProvider.PrivateEndpoint[i3].IDSelector,
			To: reference.To{
				List:    &v1beta1.PrivateEndpointList{},
				Managed: &v1beta1.PrivateEndpoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateEndpoint[i3].ID")
		}
		mg.Spec.ForProvider.PrivateEndpoint[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateEndpoint[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SignalrServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SignalrServiceIDRef,
		Selector:     mg.Spec.ForProvider.SignalrServiceIDSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SignalrServiceID")
	}
	mg.Spec.ForProvider.SignalrServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SignalrServiceIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.PrivateEndpoint); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateEndpoint[i3].ID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PrivateEndpoint[i3].IDRef,
			Selector:     mg.Spec.InitProvider.PrivateEndpoint[i3].IDSelector,
			To: reference.To{
				List:    &v1beta1.PrivateEndpointList{},
				Managed: &v1beta1.PrivateEndpoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PrivateEndpoint[i3].ID")
		}
		mg.Spec.InitProvider.PrivateEndpoint[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PrivateEndpoint[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SignalrServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SignalrServiceIDRef,
		Selector:     mg.Spec.InitProvider.SignalrServiceIDSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SignalrServiceID")
	}
	mg.Spec.InitProvider.SignalrServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SignalrServiceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SignalrSharedPrivateLinkResource.
func (mg *SignalrSharedPrivateLinkResource) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SignalrServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.SignalrServiceIDRef,
		Selector:     mg.Spec.ForProvider.SignalrServiceIDSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SignalrServiceID")
	}
	mg.Spec.ForProvider.SignalrServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SignalrServiceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.TargetResourceIDRef,
		Selector:     mg.Spec.ForProvider.TargetResourceIDSelector,
		To: reference.To{
			List:    &v1beta12.VaultList{},
			Managed: &v1beta12.Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetResourceID")
	}
	mg.Spec.ForProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SignalrServiceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.SignalrServiceIDRef,
		Selector:     mg.Spec.InitProvider.SignalrServiceIDSelector,
		To: reference.To{
			List:    &ServiceList{},
			Managed: &Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SignalrServiceID")
	}
	mg.Spec.InitProvider.SignalrServiceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SignalrServiceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TargetResourceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.TargetResourceIDRef,
		Selector:     mg.Spec.InitProvider.TargetResourceIDSelector,
		To: reference.To{
			List:    &v1beta12.VaultList{},
			Managed: &v1beta12.Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TargetResourceID")
	}
	mg.Spec.InitProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WebPubsub.
func (mg *WebPubsub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.InitProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ResourceGroupName")
	}
	mg.Spec.InitProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WebPubsubHub.
func (mg *WebPubsubHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventHandler); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.EventHandler[i3].Auth); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHandler[i3].Auth[i4].ManagedIdentityID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.EventHandler[i3].Auth[i4].ManagedIdentityIDRef,
				Selector:     mg.Spec.ForProvider.EventHandler[i3].Auth[i4].ManagedIdentityIDSelector,
				To: reference.To{
					List:    &v1beta13.UserAssignedIdentityList{},
					Managed: &v1beta13.UserAssignedIdentity{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.EventHandler[i3].Auth[i4].ManagedIdentityID")
			}
			mg.Spec.ForProvider.EventHandler[i3].Auth[i4].ManagedIdentityID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.EventHandler[i3].Auth[i4].ManagedIdentityIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventListener); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventListener[i3].EventHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventListener[i3].EventHubNameRef,
			Selector:     mg.Spec.ForProvider.EventListener[i3].EventHubNameSelector,
			To: reference.To{
				List:    &v1beta14.EventHubList{},
				Managed: &v1beta14.EventHub{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventListener[i3].EventHubName")
		}
		mg.Spec.ForProvider.EventListener[i3].EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventListener[i3].EventHubNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.EventListener); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventListener[i3].EventHubNamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EventListener[i3].EventHubNamespaceNameRef,
			Selector:     mg.Spec.ForProvider.EventListener[i3].EventHubNamespaceNameSelector,
			To: reference.To{
				List:    &v1beta14.EventHubNamespaceList{},
				Managed: &v1beta14.EventHubNamespace{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EventListener[i3].EventHubNamespaceName")
		}
		mg.Spec.ForProvider.EventListener[i3].EventHubNamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EventListener[i3].EventHubNamespaceNameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WebPubsubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WebPubsubIDRef,
		Selector:     mg.Spec.ForProvider.WebPubsubIDSelector,
		To: reference.To{
			List:    &WebPubsubList{},
			Managed: &WebPubsub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WebPubsubID")
	}
	mg.Spec.ForProvider.WebPubsubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WebPubsubIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventHandler); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.EventHandler[i3].Auth); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventHandler[i3].Auth[i4].ManagedIdentityID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.EventHandler[i3].Auth[i4].ManagedIdentityIDRef,
				Selector:     mg.Spec.InitProvider.EventHandler[i3].Auth[i4].ManagedIdentityIDSelector,
				To: reference.To{
					List:    &v1beta13.UserAssignedIdentityList{},
					Managed: &v1beta13.UserAssignedIdentity{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.EventHandler[i3].Auth[i4].ManagedIdentityID")
			}
			mg.Spec.InitProvider.EventHandler[i3].Auth[i4].ManagedIdentityID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.EventHandler[i3].Auth[i4].ManagedIdentityIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventListener); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventListener[i3].EventHubName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventListener[i3].EventHubNameRef,
			Selector:     mg.Spec.InitProvider.EventListener[i3].EventHubNameSelector,
			To: reference.To{
				List:    &v1beta14.EventHubList{},
				Managed: &v1beta14.EventHub{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventListener[i3].EventHubName")
		}
		mg.Spec.InitProvider.EventListener[i3].EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventListener[i3].EventHubNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.EventListener); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.EventListener[i3].EventHubNamespaceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.EventListener[i3].EventHubNamespaceNameRef,
			Selector:     mg.Spec.InitProvider.EventListener[i3].EventHubNamespaceNameSelector,
			To: reference.To{
				List:    &v1beta14.EventHubNamespaceList{},
				Managed: &v1beta14.EventHubNamespace{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.EventListener[i3].EventHubNamespaceName")
		}
		mg.Spec.InitProvider.EventListener[i3].EventHubNamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.EventListener[i3].EventHubNamespaceNameRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WebPubsubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.WebPubsubIDRef,
		Selector:     mg.Spec.InitProvider.WebPubsubIDSelector,
		To: reference.To{
			List:    &WebPubsubList{},
			Managed: &WebPubsub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WebPubsubID")
	}
	mg.Spec.InitProvider.WebPubsubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WebPubsubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this WebPubsubNetworkACL.
func (mg *WebPubsubNetworkACL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrivateEndpoint); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateEndpoint[i3].ID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.PrivateEndpoint[i3].IDRef,
			Selector:     mg.Spec.ForProvider.PrivateEndpoint[i3].IDSelector,
			To: reference.To{
				List:    &v1beta1.PrivateEndpointList{},
				Managed: &v1beta1.PrivateEndpoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.PrivateEndpoint[i3].ID")
		}
		mg.Spec.ForProvider.PrivateEndpoint[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.PrivateEndpoint[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WebPubsubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.WebPubsubIDRef,
		Selector:     mg.Spec.ForProvider.WebPubsubIDSelector,
		To: reference.To{
			List:    &WebPubsubList{},
			Managed: &WebPubsub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WebPubsubID")
	}
	mg.Spec.ForProvider.WebPubsubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WebPubsubIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.PrivateEndpoint); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateEndpoint[i3].ID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.PrivateEndpoint[i3].IDRef,
			Selector:     mg.Spec.InitProvider.PrivateEndpoint[i3].IDSelector,
			To: reference.To{
				List:    &v1beta1.PrivateEndpointList{},
				Managed: &v1beta1.PrivateEndpoint{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.PrivateEndpoint[i3].ID")
		}
		mg.Spec.InitProvider.PrivateEndpoint[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.PrivateEndpoint[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WebPubsubID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.WebPubsubIDRef,
		Selector:     mg.Spec.InitProvider.WebPubsubIDSelector,
		To: reference.To{
			List:    &WebPubsubList{},
			Managed: &WebPubsub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WebPubsubID")
	}
	mg.Spec.InitProvider.WebPubsubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WebPubsubIDRef = rsp.ResolvedReference

	return nil
}
