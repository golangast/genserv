// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package freetieriface provides an interface to enable mocking the AWS Free Tier service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package freetieriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/freetier"
)

// FreeTierAPI provides an interface to enable mocking the
// freetier.FreeTier service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Free Tier.
//	func myFunc(svc freetieriface.FreeTierAPI) bool {
//	    // Make svc.GetFreeTierUsage request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := freetier.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockFreeTierClient struct {
//	    freetieriface.FreeTierAPI
//	}
//	func (m *mockFreeTierClient) GetFreeTierUsage(input *freetier.GetFreeTierUsageInput) (*freetier.GetFreeTierUsageOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockFreeTierClient{}
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
type FreeTierAPI interface {
	GetFreeTierUsage(*freetier.GetFreeTierUsageInput) (*freetier.GetFreeTierUsageOutput, error)
	GetFreeTierUsageWithContext(aws.Context, *freetier.GetFreeTierUsageInput, ...request.Option) (*freetier.GetFreeTierUsageOutput, error)
	GetFreeTierUsageRequest(*freetier.GetFreeTierUsageInput) (*request.Request, *freetier.GetFreeTierUsageOutput)

	GetFreeTierUsagePages(*freetier.GetFreeTierUsageInput, func(*freetier.GetFreeTierUsageOutput, bool) bool) error
	GetFreeTierUsagePagesWithContext(aws.Context, *freetier.GetFreeTierUsageInput, func(*freetier.GetFreeTierUsageOutput, bool) bool, ...request.Option) error
}

var _ FreeTierAPI = (*freetier.FreeTier)(nil)
