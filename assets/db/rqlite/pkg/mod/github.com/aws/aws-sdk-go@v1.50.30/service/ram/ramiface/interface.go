// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ramiface provides an interface to enable mocking the AWS Resource Access Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ramiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ram"
)

// RAMAPI provides an interface to enable mocking the
// ram.RAM service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Resource Access Manager.
//	func myFunc(svc ramiface.RAMAPI) bool {
//	    // Make svc.AcceptResourceShareInvitation request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := ram.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockRAMClient struct {
//	    ramiface.RAMAPI
//	}
//	func (m *mockRAMClient) AcceptResourceShareInvitation(input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockRAMClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RAMAPI interface {
	AcceptResourceShareInvitation(*ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationWithContext(aws.Context, *ram.AcceptResourceShareInvitationInput, ...request.Option) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationRequest(*ram.AcceptResourceShareInvitationInput) (*request.Request, *ram.AcceptResourceShareInvitationOutput)

	AssociateResourceShare(*ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareWithContext(aws.Context, *ram.AssociateResourceShareInput, ...request.Option) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareRequest(*ram.AssociateResourceShareInput) (*request.Request, *ram.AssociateResourceShareOutput)

	AssociateResourceSharePermission(*ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error)
	AssociateResourceSharePermissionWithContext(aws.Context, *ram.AssociateResourceSharePermissionInput, ...request.Option) (*ram.AssociateResourceSharePermissionOutput, error)
	AssociateResourceSharePermissionRequest(*ram.AssociateResourceSharePermissionInput) (*request.Request, *ram.AssociateResourceSharePermissionOutput)

	CreatePermission(*ram.CreatePermissionInput) (*ram.CreatePermissionOutput, error)
	CreatePermissionWithContext(aws.Context, *ram.CreatePermissionInput, ...request.Option) (*ram.CreatePermissionOutput, error)
	CreatePermissionRequest(*ram.CreatePermissionInput) (*request.Request, *ram.CreatePermissionOutput)

	CreatePermissionVersion(*ram.CreatePermissionVersionInput) (*ram.CreatePermissionVersionOutput, error)
	CreatePermissionVersionWithContext(aws.Context, *ram.CreatePermissionVersionInput, ...request.Option) (*ram.CreatePermissionVersionOutput, error)
	CreatePermissionVersionRequest(*ram.CreatePermissionVersionInput) (*request.Request, *ram.CreatePermissionVersionOutput)

	CreateResourceShare(*ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareWithContext(aws.Context, *ram.CreateResourceShareInput, ...request.Option) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareRequest(*ram.CreateResourceShareInput) (*request.Request, *ram.CreateResourceShareOutput)

	DeletePermission(*ram.DeletePermissionInput) (*ram.DeletePermissionOutput, error)
	DeletePermissionWithContext(aws.Context, *ram.DeletePermissionInput, ...request.Option) (*ram.DeletePermissionOutput, error)
	DeletePermissionRequest(*ram.DeletePermissionInput) (*request.Request, *ram.DeletePermissionOutput)

	DeletePermissionVersion(*ram.DeletePermissionVersionInput) (*ram.DeletePermissionVersionOutput, error)
	DeletePermissionVersionWithContext(aws.Context, *ram.DeletePermissionVersionInput, ...request.Option) (*ram.DeletePermissionVersionOutput, error)
	DeletePermissionVersionRequest(*ram.DeletePermissionVersionInput) (*request.Request, *ram.DeletePermissionVersionOutput)

	DeleteResourceShare(*ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareWithContext(aws.Context, *ram.DeleteResourceShareInput, ...request.Option) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareRequest(*ram.DeleteResourceShareInput) (*request.Request, *ram.DeleteResourceShareOutput)

	DisassociateResourceShare(*ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareWithContext(aws.Context, *ram.DisassociateResourceShareInput, ...request.Option) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareRequest(*ram.DisassociateResourceShareInput) (*request.Request, *ram.DisassociateResourceShareOutput)

	DisassociateResourceSharePermission(*ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error)
	DisassociateResourceSharePermissionWithContext(aws.Context, *ram.DisassociateResourceSharePermissionInput, ...request.Option) (*ram.DisassociateResourceSharePermissionOutput, error)
	DisassociateResourceSharePermissionRequest(*ram.DisassociateResourceSharePermissionInput) (*request.Request, *ram.DisassociateResourceSharePermissionOutput)

	EnableSharingWithAwsOrganization(*ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationWithContext(aws.Context, *ram.EnableSharingWithAwsOrganizationInput, ...request.Option) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationRequest(*ram.EnableSharingWithAwsOrganizationInput) (*request.Request, *ram.EnableSharingWithAwsOrganizationOutput)

	GetPermission(*ram.GetPermissionInput) (*ram.GetPermissionOutput, error)
	GetPermissionWithContext(aws.Context, *ram.GetPermissionInput, ...request.Option) (*ram.GetPermissionOutput, error)
	GetPermissionRequest(*ram.GetPermissionInput) (*request.Request, *ram.GetPermissionOutput)

	GetResourcePolicies(*ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesWithContext(aws.Context, *ram.GetResourcePoliciesInput, ...request.Option) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesRequest(*ram.GetResourcePoliciesInput) (*request.Request, *ram.GetResourcePoliciesOutput)

	GetResourcePoliciesPages(*ram.GetResourcePoliciesInput, func(*ram.GetResourcePoliciesOutput, bool) bool) error
	GetResourcePoliciesPagesWithContext(aws.Context, *ram.GetResourcePoliciesInput, func(*ram.GetResourcePoliciesOutput, bool) bool, ...request.Option) error

	GetResourceShareAssociations(*ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsWithContext(aws.Context, *ram.GetResourceShareAssociationsInput, ...request.Option) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsRequest(*ram.GetResourceShareAssociationsInput) (*request.Request, *ram.GetResourceShareAssociationsOutput)

	GetResourceShareAssociationsPages(*ram.GetResourceShareAssociationsInput, func(*ram.GetResourceShareAssociationsOutput, bool) bool) error
	GetResourceShareAssociationsPagesWithContext(aws.Context, *ram.GetResourceShareAssociationsInput, func(*ram.GetResourceShareAssociationsOutput, bool) bool, ...request.Option) error

	GetResourceShareInvitations(*ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsWithContext(aws.Context, *ram.GetResourceShareInvitationsInput, ...request.Option) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsRequest(*ram.GetResourceShareInvitationsInput) (*request.Request, *ram.GetResourceShareInvitationsOutput)

	GetResourceShareInvitationsPages(*ram.GetResourceShareInvitationsInput, func(*ram.GetResourceShareInvitationsOutput, bool) bool) error
	GetResourceShareInvitationsPagesWithContext(aws.Context, *ram.GetResourceShareInvitationsInput, func(*ram.GetResourceShareInvitationsOutput, bool) bool, ...request.Option) error

	GetResourceShares(*ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesWithContext(aws.Context, *ram.GetResourceSharesInput, ...request.Option) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesRequest(*ram.GetResourceSharesInput) (*request.Request, *ram.GetResourceSharesOutput)

	GetResourceSharesPages(*ram.GetResourceSharesInput, func(*ram.GetResourceSharesOutput, bool) bool) error
	GetResourceSharesPagesWithContext(aws.Context, *ram.GetResourceSharesInput, func(*ram.GetResourceSharesOutput, bool) bool, ...request.Option) error

	ListPendingInvitationResources(*ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesWithContext(aws.Context, *ram.ListPendingInvitationResourcesInput, ...request.Option) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesRequest(*ram.ListPendingInvitationResourcesInput) (*request.Request, *ram.ListPendingInvitationResourcesOutput)

	ListPendingInvitationResourcesPages(*ram.ListPendingInvitationResourcesInput, func(*ram.ListPendingInvitationResourcesOutput, bool) bool) error
	ListPendingInvitationResourcesPagesWithContext(aws.Context, *ram.ListPendingInvitationResourcesInput, func(*ram.ListPendingInvitationResourcesOutput, bool) bool, ...request.Option) error

	ListPermissionAssociations(*ram.ListPermissionAssociationsInput) (*ram.ListPermissionAssociationsOutput, error)
	ListPermissionAssociationsWithContext(aws.Context, *ram.ListPermissionAssociationsInput, ...request.Option) (*ram.ListPermissionAssociationsOutput, error)
	ListPermissionAssociationsRequest(*ram.ListPermissionAssociationsInput) (*request.Request, *ram.ListPermissionAssociationsOutput)

	ListPermissionAssociationsPages(*ram.ListPermissionAssociationsInput, func(*ram.ListPermissionAssociationsOutput, bool) bool) error
	ListPermissionAssociationsPagesWithContext(aws.Context, *ram.ListPermissionAssociationsInput, func(*ram.ListPermissionAssociationsOutput, bool) bool, ...request.Option) error

	ListPermissionVersions(*ram.ListPermissionVersionsInput) (*ram.ListPermissionVersionsOutput, error)
	ListPermissionVersionsWithContext(aws.Context, *ram.ListPermissionVersionsInput, ...request.Option) (*ram.ListPermissionVersionsOutput, error)
	ListPermissionVersionsRequest(*ram.ListPermissionVersionsInput) (*request.Request, *ram.ListPermissionVersionsOutput)

	ListPermissionVersionsPages(*ram.ListPermissionVersionsInput, func(*ram.ListPermissionVersionsOutput, bool) bool) error
	ListPermissionVersionsPagesWithContext(aws.Context, *ram.ListPermissionVersionsInput, func(*ram.ListPermissionVersionsOutput, bool) bool, ...request.Option) error

	ListPermissions(*ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error)
	ListPermissionsWithContext(aws.Context, *ram.ListPermissionsInput, ...request.Option) (*ram.ListPermissionsOutput, error)
	ListPermissionsRequest(*ram.ListPermissionsInput) (*request.Request, *ram.ListPermissionsOutput)

	ListPermissionsPages(*ram.ListPermissionsInput, func(*ram.ListPermissionsOutput, bool) bool) error
	ListPermissionsPagesWithContext(aws.Context, *ram.ListPermissionsInput, func(*ram.ListPermissionsOutput, bool) bool, ...request.Option) error

	ListPrincipals(*ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsWithContext(aws.Context, *ram.ListPrincipalsInput, ...request.Option) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsRequest(*ram.ListPrincipalsInput) (*request.Request, *ram.ListPrincipalsOutput)

	ListPrincipalsPages(*ram.ListPrincipalsInput, func(*ram.ListPrincipalsOutput, bool) bool) error
	ListPrincipalsPagesWithContext(aws.Context, *ram.ListPrincipalsInput, func(*ram.ListPrincipalsOutput, bool) bool, ...request.Option) error

	ListReplacePermissionAssociationsWork(*ram.ListReplacePermissionAssociationsWorkInput) (*ram.ListReplacePermissionAssociationsWorkOutput, error)
	ListReplacePermissionAssociationsWorkWithContext(aws.Context, *ram.ListReplacePermissionAssociationsWorkInput, ...request.Option) (*ram.ListReplacePermissionAssociationsWorkOutput, error)
	ListReplacePermissionAssociationsWorkRequest(*ram.ListReplacePermissionAssociationsWorkInput) (*request.Request, *ram.ListReplacePermissionAssociationsWorkOutput)

	ListReplacePermissionAssociationsWorkPages(*ram.ListReplacePermissionAssociationsWorkInput, func(*ram.ListReplacePermissionAssociationsWorkOutput, bool) bool) error
	ListReplacePermissionAssociationsWorkPagesWithContext(aws.Context, *ram.ListReplacePermissionAssociationsWorkInput, func(*ram.ListReplacePermissionAssociationsWorkOutput, bool) bool, ...request.Option) error

	ListResourceSharePermissions(*ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceSharePermissionsWithContext(aws.Context, *ram.ListResourceSharePermissionsInput, ...request.Option) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceSharePermissionsRequest(*ram.ListResourceSharePermissionsInput) (*request.Request, *ram.ListResourceSharePermissionsOutput)

	ListResourceSharePermissionsPages(*ram.ListResourceSharePermissionsInput, func(*ram.ListResourceSharePermissionsOutput, bool) bool) error
	ListResourceSharePermissionsPagesWithContext(aws.Context, *ram.ListResourceSharePermissionsInput, func(*ram.ListResourceSharePermissionsOutput, bool) bool, ...request.Option) error

	ListResourceTypes(*ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error)
	ListResourceTypesWithContext(aws.Context, *ram.ListResourceTypesInput, ...request.Option) (*ram.ListResourceTypesOutput, error)
	ListResourceTypesRequest(*ram.ListResourceTypesInput) (*request.Request, *ram.ListResourceTypesOutput)

	ListResourceTypesPages(*ram.ListResourceTypesInput, func(*ram.ListResourceTypesOutput, bool) bool) error
	ListResourceTypesPagesWithContext(aws.Context, *ram.ListResourceTypesInput, func(*ram.ListResourceTypesOutput, bool) bool, ...request.Option) error

	ListResources(*ram.ListResourcesInput) (*ram.ListResourcesOutput, error)
	ListResourcesWithContext(aws.Context, *ram.ListResourcesInput, ...request.Option) (*ram.ListResourcesOutput, error)
	ListResourcesRequest(*ram.ListResourcesInput) (*request.Request, *ram.ListResourcesOutput)

	ListResourcesPages(*ram.ListResourcesInput, func(*ram.ListResourcesOutput, bool) bool) error
	ListResourcesPagesWithContext(aws.Context, *ram.ListResourcesInput, func(*ram.ListResourcesOutput, bool) bool, ...request.Option) error

	PromotePermissionCreatedFromPolicy(*ram.PromotePermissionCreatedFromPolicyInput) (*ram.PromotePermissionCreatedFromPolicyOutput, error)
	PromotePermissionCreatedFromPolicyWithContext(aws.Context, *ram.PromotePermissionCreatedFromPolicyInput, ...request.Option) (*ram.PromotePermissionCreatedFromPolicyOutput, error)
	PromotePermissionCreatedFromPolicyRequest(*ram.PromotePermissionCreatedFromPolicyInput) (*request.Request, *ram.PromotePermissionCreatedFromPolicyOutput)

	PromoteResourceShareCreatedFromPolicy(*ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
	PromoteResourceShareCreatedFromPolicyWithContext(aws.Context, *ram.PromoteResourceShareCreatedFromPolicyInput, ...request.Option) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
	PromoteResourceShareCreatedFromPolicyRequest(*ram.PromoteResourceShareCreatedFromPolicyInput) (*request.Request, *ram.PromoteResourceShareCreatedFromPolicyOutput)

	RejectResourceShareInvitation(*ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationWithContext(aws.Context, *ram.RejectResourceShareInvitationInput, ...request.Option) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationRequest(*ram.RejectResourceShareInvitationInput) (*request.Request, *ram.RejectResourceShareInvitationOutput)

	ReplacePermissionAssociations(*ram.ReplacePermissionAssociationsInput) (*ram.ReplacePermissionAssociationsOutput, error)
	ReplacePermissionAssociationsWithContext(aws.Context, *ram.ReplacePermissionAssociationsInput, ...request.Option) (*ram.ReplacePermissionAssociationsOutput, error)
	ReplacePermissionAssociationsRequest(*ram.ReplacePermissionAssociationsInput) (*request.Request, *ram.ReplacePermissionAssociationsOutput)

	SetDefaultPermissionVersion(*ram.SetDefaultPermissionVersionInput) (*ram.SetDefaultPermissionVersionOutput, error)
	SetDefaultPermissionVersionWithContext(aws.Context, *ram.SetDefaultPermissionVersionInput, ...request.Option) (*ram.SetDefaultPermissionVersionOutput, error)
	SetDefaultPermissionVersionRequest(*ram.SetDefaultPermissionVersionInput) (*request.Request, *ram.SetDefaultPermissionVersionOutput)

	TagResource(*ram.TagResourceInput) (*ram.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *ram.TagResourceInput, ...request.Option) (*ram.TagResourceOutput, error)
	TagResourceRequest(*ram.TagResourceInput) (*request.Request, *ram.TagResourceOutput)

	UntagResource(*ram.UntagResourceInput) (*ram.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *ram.UntagResourceInput, ...request.Option) (*ram.UntagResourceOutput, error)
	UntagResourceRequest(*ram.UntagResourceInput) (*request.Request, *ram.UntagResourceOutput)

	UpdateResourceShare(*ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareWithContext(aws.Context, *ram.UpdateResourceShareInput, ...request.Option) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareRequest(*ram.UpdateResourceShareInput) (*request.Request, *ram.UpdateResourceShareOutput)
}

var _ RAMAPI = (*ram.RAM)(nil)
