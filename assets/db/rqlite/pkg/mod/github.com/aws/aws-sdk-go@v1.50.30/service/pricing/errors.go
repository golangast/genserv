// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// General authentication failure. The request wasn't signed correctly.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeExpiredNextTokenException for service response error code
	// "ExpiredNextTokenException".
	//
	// The pagination token expired. Try again without a pagination token.
	ErrCodeExpiredNextTokenException = "ExpiredNextTokenException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// An error on the server occurred during the processing of your request. Try
	// again later.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The pagination token is invalid. Try again without a pagination token.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// One or more parameters had an invalid value.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The requested resource can't be found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// You've made too many requests exceeding service quotas.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"ExpiredNextTokenException": newErrorExpiredNextTokenException,
	"InternalErrorException":    newErrorInternalErrorException,
	"InvalidNextTokenException": newErrorInvalidNextTokenException,
	"InvalidParameterException": newErrorInvalidParameterException,
	"NotFoundException":         newErrorNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
}
