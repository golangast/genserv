// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package translateiface provides an interface to enable mocking the Amazon Translate service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package translateiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/translate"
)

// TranslateAPI provides an interface to enable mocking the
// translate.Translate service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Translate.
//	func myFunc(svc translateiface.TranslateAPI) bool {
//	    // Make svc.CreateParallelData request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := translate.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockTranslateClient struct {
//	    translateiface.TranslateAPI
//	}
//	func (m *mockTranslateClient) CreateParallelData(input *translate.CreateParallelDataInput) (*translate.CreateParallelDataOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockTranslateClient{}
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
type TranslateAPI interface {
	CreateParallelData(*translate.CreateParallelDataInput) (*translate.CreateParallelDataOutput, error)
	CreateParallelDataWithContext(aws.Context, *translate.CreateParallelDataInput, ...request.Option) (*translate.CreateParallelDataOutput, error)
	CreateParallelDataRequest(*translate.CreateParallelDataInput) (*request.Request, *translate.CreateParallelDataOutput)

	DeleteParallelData(*translate.DeleteParallelDataInput) (*translate.DeleteParallelDataOutput, error)
	DeleteParallelDataWithContext(aws.Context, *translate.DeleteParallelDataInput, ...request.Option) (*translate.DeleteParallelDataOutput, error)
	DeleteParallelDataRequest(*translate.DeleteParallelDataInput) (*request.Request, *translate.DeleteParallelDataOutput)

	DeleteTerminology(*translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error)
	DeleteTerminologyWithContext(aws.Context, *translate.DeleteTerminologyInput, ...request.Option) (*translate.DeleteTerminologyOutput, error)
	DeleteTerminologyRequest(*translate.DeleteTerminologyInput) (*request.Request, *translate.DeleteTerminologyOutput)

	DescribeTextTranslationJob(*translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error)
	DescribeTextTranslationJobWithContext(aws.Context, *translate.DescribeTextTranslationJobInput, ...request.Option) (*translate.DescribeTextTranslationJobOutput, error)
	DescribeTextTranslationJobRequest(*translate.DescribeTextTranslationJobInput) (*request.Request, *translate.DescribeTextTranslationJobOutput)

	GetParallelData(*translate.GetParallelDataInput) (*translate.GetParallelDataOutput, error)
	GetParallelDataWithContext(aws.Context, *translate.GetParallelDataInput, ...request.Option) (*translate.GetParallelDataOutput, error)
	GetParallelDataRequest(*translate.GetParallelDataInput) (*request.Request, *translate.GetParallelDataOutput)

	GetTerminology(*translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error)
	GetTerminologyWithContext(aws.Context, *translate.GetTerminologyInput, ...request.Option) (*translate.GetTerminologyOutput, error)
	GetTerminologyRequest(*translate.GetTerminologyInput) (*request.Request, *translate.GetTerminologyOutput)

	ImportTerminology(*translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error)
	ImportTerminologyWithContext(aws.Context, *translate.ImportTerminologyInput, ...request.Option) (*translate.ImportTerminologyOutput, error)
	ImportTerminologyRequest(*translate.ImportTerminologyInput) (*request.Request, *translate.ImportTerminologyOutput)

	ListLanguages(*translate.ListLanguagesInput) (*translate.ListLanguagesOutput, error)
	ListLanguagesWithContext(aws.Context, *translate.ListLanguagesInput, ...request.Option) (*translate.ListLanguagesOutput, error)
	ListLanguagesRequest(*translate.ListLanguagesInput) (*request.Request, *translate.ListLanguagesOutput)

	ListLanguagesPages(*translate.ListLanguagesInput, func(*translate.ListLanguagesOutput, bool) bool) error
	ListLanguagesPagesWithContext(aws.Context, *translate.ListLanguagesInput, func(*translate.ListLanguagesOutput, bool) bool, ...request.Option) error

	ListParallelData(*translate.ListParallelDataInput) (*translate.ListParallelDataOutput, error)
	ListParallelDataWithContext(aws.Context, *translate.ListParallelDataInput, ...request.Option) (*translate.ListParallelDataOutput, error)
	ListParallelDataRequest(*translate.ListParallelDataInput) (*request.Request, *translate.ListParallelDataOutput)

	ListParallelDataPages(*translate.ListParallelDataInput, func(*translate.ListParallelDataOutput, bool) bool) error
	ListParallelDataPagesWithContext(aws.Context, *translate.ListParallelDataInput, func(*translate.ListParallelDataOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*translate.ListTagsForResourceInput) (*translate.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *translate.ListTagsForResourceInput, ...request.Option) (*translate.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*translate.ListTagsForResourceInput) (*request.Request, *translate.ListTagsForResourceOutput)

	ListTerminologies(*translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error)
	ListTerminologiesWithContext(aws.Context, *translate.ListTerminologiesInput, ...request.Option) (*translate.ListTerminologiesOutput, error)
	ListTerminologiesRequest(*translate.ListTerminologiesInput) (*request.Request, *translate.ListTerminologiesOutput)

	ListTerminologiesPages(*translate.ListTerminologiesInput, func(*translate.ListTerminologiesOutput, bool) bool) error
	ListTerminologiesPagesWithContext(aws.Context, *translate.ListTerminologiesInput, func(*translate.ListTerminologiesOutput, bool) bool, ...request.Option) error

	ListTextTranslationJobs(*translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error)
	ListTextTranslationJobsWithContext(aws.Context, *translate.ListTextTranslationJobsInput, ...request.Option) (*translate.ListTextTranslationJobsOutput, error)
	ListTextTranslationJobsRequest(*translate.ListTextTranslationJobsInput) (*request.Request, *translate.ListTextTranslationJobsOutput)

	ListTextTranslationJobsPages(*translate.ListTextTranslationJobsInput, func(*translate.ListTextTranslationJobsOutput, bool) bool) error
	ListTextTranslationJobsPagesWithContext(aws.Context, *translate.ListTextTranslationJobsInput, func(*translate.ListTextTranslationJobsOutput, bool) bool, ...request.Option) error

	StartTextTranslationJob(*translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error)
	StartTextTranslationJobWithContext(aws.Context, *translate.StartTextTranslationJobInput, ...request.Option) (*translate.StartTextTranslationJobOutput, error)
	StartTextTranslationJobRequest(*translate.StartTextTranslationJobInput) (*request.Request, *translate.StartTextTranslationJobOutput)

	StopTextTranslationJob(*translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error)
	StopTextTranslationJobWithContext(aws.Context, *translate.StopTextTranslationJobInput, ...request.Option) (*translate.StopTextTranslationJobOutput, error)
	StopTextTranslationJobRequest(*translate.StopTextTranslationJobInput) (*request.Request, *translate.StopTextTranslationJobOutput)

	TagResource(*translate.TagResourceInput) (*translate.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *translate.TagResourceInput, ...request.Option) (*translate.TagResourceOutput, error)
	TagResourceRequest(*translate.TagResourceInput) (*request.Request, *translate.TagResourceOutput)

	Text(*translate.TextInput) (*translate.TextOutput, error)
	TextWithContext(aws.Context, *translate.TextInput, ...request.Option) (*translate.TextOutput, error)
	TextRequest(*translate.TextInput) (*request.Request, *translate.TextOutput)

	TranslateDocument(*translate.TranslateDocumentInput) (*translate.TranslateDocumentOutput, error)
	TranslateDocumentWithContext(aws.Context, *translate.TranslateDocumentInput, ...request.Option) (*translate.TranslateDocumentOutput, error)
	TranslateDocumentRequest(*translate.TranslateDocumentInput) (*request.Request, *translate.TranslateDocumentOutput)

	UntagResource(*translate.UntagResourceInput) (*translate.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *translate.UntagResourceInput, ...request.Option) (*translate.UntagResourceOutput, error)
	UntagResourceRequest(*translate.UntagResourceInput) (*request.Request, *translate.UntagResourceOutput)

	UpdateParallelData(*translate.UpdateParallelDataInput) (*translate.UpdateParallelDataOutput, error)
	UpdateParallelDataWithContext(aws.Context, *translate.UpdateParallelDataInput, ...request.Option) (*translate.UpdateParallelDataOutput, error)
	UpdateParallelDataRequest(*translate.UpdateParallelDataInput) (*request.Request, *translate.UpdateParallelDataOutput)
}

var _ TranslateAPI = (*translate.Translate)(nil)
