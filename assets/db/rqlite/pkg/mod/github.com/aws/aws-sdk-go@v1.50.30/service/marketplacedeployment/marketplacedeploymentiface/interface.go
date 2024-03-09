// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package marketplacedeploymentiface provides an interface to enable mocking the AWS Marketplace Deployment Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package marketplacedeploymentiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/marketplacedeployment"
)

// MarketplaceDeploymentAPI provides an interface to enable mocking the
// marketplacedeployment.MarketplaceDeployment service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Marketplace Deployment Service.
//	func myFunc(svc marketplacedeploymentiface.MarketplaceDeploymentAPI) bool {
//	    // Make svc.ListTagsForResource request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := marketplacedeployment.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockMarketplaceDeploymentClient struct {
//	    marketplacedeploymentiface.MarketplaceDeploymentAPI
//	}
//	func (m *mockMarketplaceDeploymentClient) ListTagsForResource(input *marketplacedeployment.ListTagsForResourceInput) (*marketplacedeployment.ListTagsForResourceOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockMarketplaceDeploymentClient{}
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
type MarketplaceDeploymentAPI interface {
	ListTagsForResource(*marketplacedeployment.ListTagsForResourceInput) (*marketplacedeployment.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *marketplacedeployment.ListTagsForResourceInput, ...request.Option) (*marketplacedeployment.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*marketplacedeployment.ListTagsForResourceInput) (*request.Request, *marketplacedeployment.ListTagsForResourceOutput)

	PutDeploymentParameter(*marketplacedeployment.PutDeploymentParameterInput) (*marketplacedeployment.PutDeploymentParameterOutput, error)
	PutDeploymentParameterWithContext(aws.Context, *marketplacedeployment.PutDeploymentParameterInput, ...request.Option) (*marketplacedeployment.PutDeploymentParameterOutput, error)
	PutDeploymentParameterRequest(*marketplacedeployment.PutDeploymentParameterInput) (*request.Request, *marketplacedeployment.PutDeploymentParameterOutput)

	TagResource(*marketplacedeployment.TagResourceInput) (*marketplacedeployment.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *marketplacedeployment.TagResourceInput, ...request.Option) (*marketplacedeployment.TagResourceOutput, error)
	TagResourceRequest(*marketplacedeployment.TagResourceInput) (*request.Request, *marketplacedeployment.TagResourceOutput)

	UntagResource(*marketplacedeployment.UntagResourceInput) (*marketplacedeployment.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *marketplacedeployment.UntagResourceInput, ...request.Option) (*marketplacedeployment.UntagResourceOutput, error)
	UntagResourceRequest(*marketplacedeployment.UntagResourceInput) (*request.Request, *marketplacedeployment.UntagResourceOutput)
}

var _ MarketplaceDeploymentAPI = (*marketplacedeployment.MarketplaceDeployment)(nil)
