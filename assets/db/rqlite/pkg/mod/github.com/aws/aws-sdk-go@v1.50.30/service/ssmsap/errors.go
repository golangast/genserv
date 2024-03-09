// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssmsap

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// A conflict has occurred.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An internal error has occurred.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource is not available.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// The request is not authorized.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by an AWS service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":         newErrorConflictException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"UnauthorizedException":     newErrorUnauthorizedException,
	"ValidationException":       newErrorValidationException,
}
