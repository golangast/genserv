// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package licensemanageriface provides an interface to enable mocking the AWS License Manager service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package licensemanageriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/licensemanager"
)

// LicenseManagerAPI provides an interface to enable mocking the
// licensemanager.LicenseManager service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS License Manager.
//	func myFunc(svc licensemanageriface.LicenseManagerAPI) bool {
//	    // Make svc.AcceptGrant request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := licensemanager.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockLicenseManagerClient struct {
//	    licensemanageriface.LicenseManagerAPI
//	}
//	func (m *mockLicenseManagerClient) AcceptGrant(input *licensemanager.AcceptGrantInput) (*licensemanager.AcceptGrantOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockLicenseManagerClient{}
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
type LicenseManagerAPI interface {
	AcceptGrant(*licensemanager.AcceptGrantInput) (*licensemanager.AcceptGrantOutput, error)
	AcceptGrantWithContext(aws.Context, *licensemanager.AcceptGrantInput, ...request.Option) (*licensemanager.AcceptGrantOutput, error)
	AcceptGrantRequest(*licensemanager.AcceptGrantInput) (*request.Request, *licensemanager.AcceptGrantOutput)

	CheckInLicense(*licensemanager.CheckInLicenseInput) (*licensemanager.CheckInLicenseOutput, error)
	CheckInLicenseWithContext(aws.Context, *licensemanager.CheckInLicenseInput, ...request.Option) (*licensemanager.CheckInLicenseOutput, error)
	CheckInLicenseRequest(*licensemanager.CheckInLicenseInput) (*request.Request, *licensemanager.CheckInLicenseOutput)

	CheckoutBorrowLicense(*licensemanager.CheckoutBorrowLicenseInput) (*licensemanager.CheckoutBorrowLicenseOutput, error)
	CheckoutBorrowLicenseWithContext(aws.Context, *licensemanager.CheckoutBorrowLicenseInput, ...request.Option) (*licensemanager.CheckoutBorrowLicenseOutput, error)
	CheckoutBorrowLicenseRequest(*licensemanager.CheckoutBorrowLicenseInput) (*request.Request, *licensemanager.CheckoutBorrowLicenseOutput)

	CheckoutLicense(*licensemanager.CheckoutLicenseInput) (*licensemanager.CheckoutLicenseOutput, error)
	CheckoutLicenseWithContext(aws.Context, *licensemanager.CheckoutLicenseInput, ...request.Option) (*licensemanager.CheckoutLicenseOutput, error)
	CheckoutLicenseRequest(*licensemanager.CheckoutLicenseInput) (*request.Request, *licensemanager.CheckoutLicenseOutput)

	CreateGrant(*licensemanager.CreateGrantInput) (*licensemanager.CreateGrantOutput, error)
	CreateGrantWithContext(aws.Context, *licensemanager.CreateGrantInput, ...request.Option) (*licensemanager.CreateGrantOutput, error)
	CreateGrantRequest(*licensemanager.CreateGrantInput) (*request.Request, *licensemanager.CreateGrantOutput)

	CreateGrantVersion(*licensemanager.CreateGrantVersionInput) (*licensemanager.CreateGrantVersionOutput, error)
	CreateGrantVersionWithContext(aws.Context, *licensemanager.CreateGrantVersionInput, ...request.Option) (*licensemanager.CreateGrantVersionOutput, error)
	CreateGrantVersionRequest(*licensemanager.CreateGrantVersionInput) (*request.Request, *licensemanager.CreateGrantVersionOutput)

	CreateLicense(*licensemanager.CreateLicenseInput) (*licensemanager.CreateLicenseOutput, error)
	CreateLicenseWithContext(aws.Context, *licensemanager.CreateLicenseInput, ...request.Option) (*licensemanager.CreateLicenseOutput, error)
	CreateLicenseRequest(*licensemanager.CreateLicenseInput) (*request.Request, *licensemanager.CreateLicenseOutput)

	CreateLicenseConfiguration(*licensemanager.CreateLicenseConfigurationInput) (*licensemanager.CreateLicenseConfigurationOutput, error)
	CreateLicenseConfigurationWithContext(aws.Context, *licensemanager.CreateLicenseConfigurationInput, ...request.Option) (*licensemanager.CreateLicenseConfigurationOutput, error)
	CreateLicenseConfigurationRequest(*licensemanager.CreateLicenseConfigurationInput) (*request.Request, *licensemanager.CreateLicenseConfigurationOutput)

	CreateLicenseConversionTaskForResource(*licensemanager.CreateLicenseConversionTaskForResourceInput) (*licensemanager.CreateLicenseConversionTaskForResourceOutput, error)
	CreateLicenseConversionTaskForResourceWithContext(aws.Context, *licensemanager.CreateLicenseConversionTaskForResourceInput, ...request.Option) (*licensemanager.CreateLicenseConversionTaskForResourceOutput, error)
	CreateLicenseConversionTaskForResourceRequest(*licensemanager.CreateLicenseConversionTaskForResourceInput) (*request.Request, *licensemanager.CreateLicenseConversionTaskForResourceOutput)

	CreateLicenseManagerReportGenerator(*licensemanager.CreateLicenseManagerReportGeneratorInput) (*licensemanager.CreateLicenseManagerReportGeneratorOutput, error)
	CreateLicenseManagerReportGeneratorWithContext(aws.Context, *licensemanager.CreateLicenseManagerReportGeneratorInput, ...request.Option) (*licensemanager.CreateLicenseManagerReportGeneratorOutput, error)
	CreateLicenseManagerReportGeneratorRequest(*licensemanager.CreateLicenseManagerReportGeneratorInput) (*request.Request, *licensemanager.CreateLicenseManagerReportGeneratorOutput)

	CreateLicenseVersion(*licensemanager.CreateLicenseVersionInput) (*licensemanager.CreateLicenseVersionOutput, error)
	CreateLicenseVersionWithContext(aws.Context, *licensemanager.CreateLicenseVersionInput, ...request.Option) (*licensemanager.CreateLicenseVersionOutput, error)
	CreateLicenseVersionRequest(*licensemanager.CreateLicenseVersionInput) (*request.Request, *licensemanager.CreateLicenseVersionOutput)

	CreateToken(*licensemanager.CreateTokenInput) (*licensemanager.CreateTokenOutput, error)
	CreateTokenWithContext(aws.Context, *licensemanager.CreateTokenInput, ...request.Option) (*licensemanager.CreateTokenOutput, error)
	CreateTokenRequest(*licensemanager.CreateTokenInput) (*request.Request, *licensemanager.CreateTokenOutput)

	DeleteGrant(*licensemanager.DeleteGrantInput) (*licensemanager.DeleteGrantOutput, error)
	DeleteGrantWithContext(aws.Context, *licensemanager.DeleteGrantInput, ...request.Option) (*licensemanager.DeleteGrantOutput, error)
	DeleteGrantRequest(*licensemanager.DeleteGrantInput) (*request.Request, *licensemanager.DeleteGrantOutput)

	DeleteLicense(*licensemanager.DeleteLicenseInput) (*licensemanager.DeleteLicenseOutput, error)
	DeleteLicenseWithContext(aws.Context, *licensemanager.DeleteLicenseInput, ...request.Option) (*licensemanager.DeleteLicenseOutput, error)
	DeleteLicenseRequest(*licensemanager.DeleteLicenseInput) (*request.Request, *licensemanager.DeleteLicenseOutput)

	DeleteLicenseConfiguration(*licensemanager.DeleteLicenseConfigurationInput) (*licensemanager.DeleteLicenseConfigurationOutput, error)
	DeleteLicenseConfigurationWithContext(aws.Context, *licensemanager.DeleteLicenseConfigurationInput, ...request.Option) (*licensemanager.DeleteLicenseConfigurationOutput, error)
	DeleteLicenseConfigurationRequest(*licensemanager.DeleteLicenseConfigurationInput) (*request.Request, *licensemanager.DeleteLicenseConfigurationOutput)

	DeleteLicenseManagerReportGenerator(*licensemanager.DeleteLicenseManagerReportGeneratorInput) (*licensemanager.DeleteLicenseManagerReportGeneratorOutput, error)
	DeleteLicenseManagerReportGeneratorWithContext(aws.Context, *licensemanager.DeleteLicenseManagerReportGeneratorInput, ...request.Option) (*licensemanager.DeleteLicenseManagerReportGeneratorOutput, error)
	DeleteLicenseManagerReportGeneratorRequest(*licensemanager.DeleteLicenseManagerReportGeneratorInput) (*request.Request, *licensemanager.DeleteLicenseManagerReportGeneratorOutput)

	DeleteToken(*licensemanager.DeleteTokenInput) (*licensemanager.DeleteTokenOutput, error)
	DeleteTokenWithContext(aws.Context, *licensemanager.DeleteTokenInput, ...request.Option) (*licensemanager.DeleteTokenOutput, error)
	DeleteTokenRequest(*licensemanager.DeleteTokenInput) (*request.Request, *licensemanager.DeleteTokenOutput)

	ExtendLicenseConsumption(*licensemanager.ExtendLicenseConsumptionInput) (*licensemanager.ExtendLicenseConsumptionOutput, error)
	ExtendLicenseConsumptionWithContext(aws.Context, *licensemanager.ExtendLicenseConsumptionInput, ...request.Option) (*licensemanager.ExtendLicenseConsumptionOutput, error)
	ExtendLicenseConsumptionRequest(*licensemanager.ExtendLicenseConsumptionInput) (*request.Request, *licensemanager.ExtendLicenseConsumptionOutput)

	GetAccessToken(*licensemanager.GetAccessTokenInput) (*licensemanager.GetAccessTokenOutput, error)
	GetAccessTokenWithContext(aws.Context, *licensemanager.GetAccessTokenInput, ...request.Option) (*licensemanager.GetAccessTokenOutput, error)
	GetAccessTokenRequest(*licensemanager.GetAccessTokenInput) (*request.Request, *licensemanager.GetAccessTokenOutput)

	GetGrant(*licensemanager.GetGrantInput) (*licensemanager.GetGrantOutput, error)
	GetGrantWithContext(aws.Context, *licensemanager.GetGrantInput, ...request.Option) (*licensemanager.GetGrantOutput, error)
	GetGrantRequest(*licensemanager.GetGrantInput) (*request.Request, *licensemanager.GetGrantOutput)

	GetLicense(*licensemanager.GetLicenseInput) (*licensemanager.GetLicenseOutput, error)
	GetLicenseWithContext(aws.Context, *licensemanager.GetLicenseInput, ...request.Option) (*licensemanager.GetLicenseOutput, error)
	GetLicenseRequest(*licensemanager.GetLicenseInput) (*request.Request, *licensemanager.GetLicenseOutput)

	GetLicenseConfiguration(*licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error)
	GetLicenseConfigurationWithContext(aws.Context, *licensemanager.GetLicenseConfigurationInput, ...request.Option) (*licensemanager.GetLicenseConfigurationOutput, error)
	GetLicenseConfigurationRequest(*licensemanager.GetLicenseConfigurationInput) (*request.Request, *licensemanager.GetLicenseConfigurationOutput)

	GetLicenseConversionTask(*licensemanager.GetLicenseConversionTaskInput) (*licensemanager.GetLicenseConversionTaskOutput, error)
	GetLicenseConversionTaskWithContext(aws.Context, *licensemanager.GetLicenseConversionTaskInput, ...request.Option) (*licensemanager.GetLicenseConversionTaskOutput, error)
	GetLicenseConversionTaskRequest(*licensemanager.GetLicenseConversionTaskInput) (*request.Request, *licensemanager.GetLicenseConversionTaskOutput)

	GetLicenseManagerReportGenerator(*licensemanager.GetLicenseManagerReportGeneratorInput) (*licensemanager.GetLicenseManagerReportGeneratorOutput, error)
	GetLicenseManagerReportGeneratorWithContext(aws.Context, *licensemanager.GetLicenseManagerReportGeneratorInput, ...request.Option) (*licensemanager.GetLicenseManagerReportGeneratorOutput, error)
	GetLicenseManagerReportGeneratorRequest(*licensemanager.GetLicenseManagerReportGeneratorInput) (*request.Request, *licensemanager.GetLicenseManagerReportGeneratorOutput)

	GetLicenseUsage(*licensemanager.GetLicenseUsageInput) (*licensemanager.GetLicenseUsageOutput, error)
	GetLicenseUsageWithContext(aws.Context, *licensemanager.GetLicenseUsageInput, ...request.Option) (*licensemanager.GetLicenseUsageOutput, error)
	GetLicenseUsageRequest(*licensemanager.GetLicenseUsageInput) (*request.Request, *licensemanager.GetLicenseUsageOutput)

	GetServiceSettings(*licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error)
	GetServiceSettingsWithContext(aws.Context, *licensemanager.GetServiceSettingsInput, ...request.Option) (*licensemanager.GetServiceSettingsOutput, error)
	GetServiceSettingsRequest(*licensemanager.GetServiceSettingsInput) (*request.Request, *licensemanager.GetServiceSettingsOutput)

	ListAssociationsForLicenseConfiguration(*licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
	ListAssociationsForLicenseConfigurationWithContext(aws.Context, *licensemanager.ListAssociationsForLicenseConfigurationInput, ...request.Option) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error)
	ListAssociationsForLicenseConfigurationRequest(*licensemanager.ListAssociationsForLicenseConfigurationInput) (*request.Request, *licensemanager.ListAssociationsForLicenseConfigurationOutput)

	ListDistributedGrants(*licensemanager.ListDistributedGrantsInput) (*licensemanager.ListDistributedGrantsOutput, error)
	ListDistributedGrantsWithContext(aws.Context, *licensemanager.ListDistributedGrantsInput, ...request.Option) (*licensemanager.ListDistributedGrantsOutput, error)
	ListDistributedGrantsRequest(*licensemanager.ListDistributedGrantsInput) (*request.Request, *licensemanager.ListDistributedGrantsOutput)

	ListFailuresForLicenseConfigurationOperations(*licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
	ListFailuresForLicenseConfigurationOperationsWithContext(aws.Context, *licensemanager.ListFailuresForLicenseConfigurationOperationsInput, ...request.Option) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error)
	ListFailuresForLicenseConfigurationOperationsRequest(*licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*request.Request, *licensemanager.ListFailuresForLicenseConfigurationOperationsOutput)

	ListLicenseConfigurations(*licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error)
	ListLicenseConfigurationsWithContext(aws.Context, *licensemanager.ListLicenseConfigurationsInput, ...request.Option) (*licensemanager.ListLicenseConfigurationsOutput, error)
	ListLicenseConfigurationsRequest(*licensemanager.ListLicenseConfigurationsInput) (*request.Request, *licensemanager.ListLicenseConfigurationsOutput)

	ListLicenseConversionTasks(*licensemanager.ListLicenseConversionTasksInput) (*licensemanager.ListLicenseConversionTasksOutput, error)
	ListLicenseConversionTasksWithContext(aws.Context, *licensemanager.ListLicenseConversionTasksInput, ...request.Option) (*licensemanager.ListLicenseConversionTasksOutput, error)
	ListLicenseConversionTasksRequest(*licensemanager.ListLicenseConversionTasksInput) (*request.Request, *licensemanager.ListLicenseConversionTasksOutput)

	ListLicenseManagerReportGenerators(*licensemanager.ListLicenseManagerReportGeneratorsInput) (*licensemanager.ListLicenseManagerReportGeneratorsOutput, error)
	ListLicenseManagerReportGeneratorsWithContext(aws.Context, *licensemanager.ListLicenseManagerReportGeneratorsInput, ...request.Option) (*licensemanager.ListLicenseManagerReportGeneratorsOutput, error)
	ListLicenseManagerReportGeneratorsRequest(*licensemanager.ListLicenseManagerReportGeneratorsInput) (*request.Request, *licensemanager.ListLicenseManagerReportGeneratorsOutput)

	ListLicenseSpecificationsForResource(*licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
	ListLicenseSpecificationsForResourceWithContext(aws.Context, *licensemanager.ListLicenseSpecificationsForResourceInput, ...request.Option) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error)
	ListLicenseSpecificationsForResourceRequest(*licensemanager.ListLicenseSpecificationsForResourceInput) (*request.Request, *licensemanager.ListLicenseSpecificationsForResourceOutput)

	ListLicenseVersions(*licensemanager.ListLicenseVersionsInput) (*licensemanager.ListLicenseVersionsOutput, error)
	ListLicenseVersionsWithContext(aws.Context, *licensemanager.ListLicenseVersionsInput, ...request.Option) (*licensemanager.ListLicenseVersionsOutput, error)
	ListLicenseVersionsRequest(*licensemanager.ListLicenseVersionsInput) (*request.Request, *licensemanager.ListLicenseVersionsOutput)

	ListLicenses(*licensemanager.ListLicensesInput) (*licensemanager.ListLicensesOutput, error)
	ListLicensesWithContext(aws.Context, *licensemanager.ListLicensesInput, ...request.Option) (*licensemanager.ListLicensesOutput, error)
	ListLicensesRequest(*licensemanager.ListLicensesInput) (*request.Request, *licensemanager.ListLicensesOutput)

	ListReceivedGrants(*licensemanager.ListReceivedGrantsInput) (*licensemanager.ListReceivedGrantsOutput, error)
	ListReceivedGrantsWithContext(aws.Context, *licensemanager.ListReceivedGrantsInput, ...request.Option) (*licensemanager.ListReceivedGrantsOutput, error)
	ListReceivedGrantsRequest(*licensemanager.ListReceivedGrantsInput) (*request.Request, *licensemanager.ListReceivedGrantsOutput)

	ListReceivedGrantsForOrganization(*licensemanager.ListReceivedGrantsForOrganizationInput) (*licensemanager.ListReceivedGrantsForOrganizationOutput, error)
	ListReceivedGrantsForOrganizationWithContext(aws.Context, *licensemanager.ListReceivedGrantsForOrganizationInput, ...request.Option) (*licensemanager.ListReceivedGrantsForOrganizationOutput, error)
	ListReceivedGrantsForOrganizationRequest(*licensemanager.ListReceivedGrantsForOrganizationInput) (*request.Request, *licensemanager.ListReceivedGrantsForOrganizationOutput)

	ListReceivedLicenses(*licensemanager.ListReceivedLicensesInput) (*licensemanager.ListReceivedLicensesOutput, error)
	ListReceivedLicensesWithContext(aws.Context, *licensemanager.ListReceivedLicensesInput, ...request.Option) (*licensemanager.ListReceivedLicensesOutput, error)
	ListReceivedLicensesRequest(*licensemanager.ListReceivedLicensesInput) (*request.Request, *licensemanager.ListReceivedLicensesOutput)

	ListReceivedLicensesForOrganization(*licensemanager.ListReceivedLicensesForOrganizationInput) (*licensemanager.ListReceivedLicensesForOrganizationOutput, error)
	ListReceivedLicensesForOrganizationWithContext(aws.Context, *licensemanager.ListReceivedLicensesForOrganizationInput, ...request.Option) (*licensemanager.ListReceivedLicensesForOrganizationOutput, error)
	ListReceivedLicensesForOrganizationRequest(*licensemanager.ListReceivedLicensesForOrganizationInput) (*request.Request, *licensemanager.ListReceivedLicensesForOrganizationOutput)

	ListResourceInventory(*licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error)
	ListResourceInventoryWithContext(aws.Context, *licensemanager.ListResourceInventoryInput, ...request.Option) (*licensemanager.ListResourceInventoryOutput, error)
	ListResourceInventoryRequest(*licensemanager.ListResourceInventoryInput) (*request.Request, *licensemanager.ListResourceInventoryOutput)

	ListTagsForResource(*licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *licensemanager.ListTagsForResourceInput, ...request.Option) (*licensemanager.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*licensemanager.ListTagsForResourceInput) (*request.Request, *licensemanager.ListTagsForResourceOutput)

	ListTokens(*licensemanager.ListTokensInput) (*licensemanager.ListTokensOutput, error)
	ListTokensWithContext(aws.Context, *licensemanager.ListTokensInput, ...request.Option) (*licensemanager.ListTokensOutput, error)
	ListTokensRequest(*licensemanager.ListTokensInput) (*request.Request, *licensemanager.ListTokensOutput)

	ListUsageForLicenseConfiguration(*licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
	ListUsageForLicenseConfigurationWithContext(aws.Context, *licensemanager.ListUsageForLicenseConfigurationInput, ...request.Option) (*licensemanager.ListUsageForLicenseConfigurationOutput, error)
	ListUsageForLicenseConfigurationRequest(*licensemanager.ListUsageForLicenseConfigurationInput) (*request.Request, *licensemanager.ListUsageForLicenseConfigurationOutput)

	RejectGrant(*licensemanager.RejectGrantInput) (*licensemanager.RejectGrantOutput, error)
	RejectGrantWithContext(aws.Context, *licensemanager.RejectGrantInput, ...request.Option) (*licensemanager.RejectGrantOutput, error)
	RejectGrantRequest(*licensemanager.RejectGrantInput) (*request.Request, *licensemanager.RejectGrantOutput)

	TagResource(*licensemanager.TagResourceInput) (*licensemanager.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *licensemanager.TagResourceInput, ...request.Option) (*licensemanager.TagResourceOutput, error)
	TagResourceRequest(*licensemanager.TagResourceInput) (*request.Request, *licensemanager.TagResourceOutput)

	UntagResource(*licensemanager.UntagResourceInput) (*licensemanager.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *licensemanager.UntagResourceInput, ...request.Option) (*licensemanager.UntagResourceOutput, error)
	UntagResourceRequest(*licensemanager.UntagResourceInput) (*request.Request, *licensemanager.UntagResourceOutput)

	UpdateLicenseConfiguration(*licensemanager.UpdateLicenseConfigurationInput) (*licensemanager.UpdateLicenseConfigurationOutput, error)
	UpdateLicenseConfigurationWithContext(aws.Context, *licensemanager.UpdateLicenseConfigurationInput, ...request.Option) (*licensemanager.UpdateLicenseConfigurationOutput, error)
	UpdateLicenseConfigurationRequest(*licensemanager.UpdateLicenseConfigurationInput) (*request.Request, *licensemanager.UpdateLicenseConfigurationOutput)

	UpdateLicenseManagerReportGenerator(*licensemanager.UpdateLicenseManagerReportGeneratorInput) (*licensemanager.UpdateLicenseManagerReportGeneratorOutput, error)
	UpdateLicenseManagerReportGeneratorWithContext(aws.Context, *licensemanager.UpdateLicenseManagerReportGeneratorInput, ...request.Option) (*licensemanager.UpdateLicenseManagerReportGeneratorOutput, error)
	UpdateLicenseManagerReportGeneratorRequest(*licensemanager.UpdateLicenseManagerReportGeneratorInput) (*request.Request, *licensemanager.UpdateLicenseManagerReportGeneratorOutput)

	UpdateLicenseSpecificationsForResource(*licensemanager.UpdateLicenseSpecificationsForResourceInput) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
	UpdateLicenseSpecificationsForResourceWithContext(aws.Context, *licensemanager.UpdateLicenseSpecificationsForResourceInput, ...request.Option) (*licensemanager.UpdateLicenseSpecificationsForResourceOutput, error)
	UpdateLicenseSpecificationsForResourceRequest(*licensemanager.UpdateLicenseSpecificationsForResourceInput) (*request.Request, *licensemanager.UpdateLicenseSpecificationsForResourceOutput)

	UpdateServiceSettings(*licensemanager.UpdateServiceSettingsInput) (*licensemanager.UpdateServiceSettingsOutput, error)
	UpdateServiceSettingsWithContext(aws.Context, *licensemanager.UpdateServiceSettingsInput, ...request.Option) (*licensemanager.UpdateServiceSettingsOutput, error)
	UpdateServiceSettingsRequest(*licensemanager.UpdateServiceSettingsInput) (*request.Request, *licensemanager.UpdateServiceSettingsOutput)
}

var _ LicenseManagerAPI = (*licensemanager.LicenseManager)(nil)
