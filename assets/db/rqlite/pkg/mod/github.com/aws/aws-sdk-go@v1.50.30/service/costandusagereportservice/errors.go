// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeDuplicateReportNameException for service response error code
	// "DuplicateReportNameException".
	//
	// A report with the specified name already exists in the account. Specify a
	// different report name.
	ErrCodeDuplicateReportNameException = "DuplicateReportNameException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// An error on the server occurred during the processing of your request. Try
	// again later.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeReportLimitReachedException for service response error code
	// "ReportLimitReachedException".
	//
	// This account already has five reports defined. To define a new report, you
	// must delete an existing report.
	ErrCodeReportLimitReachedException = "ReportLimitReachedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified report (ReportName) in the request doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by an Amazon Web Services
	// service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DuplicateReportNameException": newErrorDuplicateReportNameException,
	"InternalErrorException":       newErrorInternalErrorException,
	"ReportLimitReachedException":  newErrorReportLimitReachedException,
	"ResourceNotFoundException":    newErrorResourceNotFoundException,
	"ValidationException":          newErrorValidationException,
}
