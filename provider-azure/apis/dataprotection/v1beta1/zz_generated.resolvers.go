/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	rconfig "github.com/upbound/official-providers/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VaultIDRef,
		Selector:     mg.Spec.ForProvider.VaultIDSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultID")
	}
	mg.Spec.ForProvider.VaultID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VaultNameRef,
		Selector:     mg.Spec.ForProvider.VaultNameSelector,
		To: reference.To{
			List:    &BackupVaultList{},
			Managed: &BackupVault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultName")
	}
	mg.Spec.ForProvider.VaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this BackupVault.
func (mg *BackupVault) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}