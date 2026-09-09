package errs

import (
	"fmt"

	"github.com/vaultnite/transcendent/internal/buildinfo"
	"github.com/vaultnite/transcendent/internal/models"
)

func NewGenericHTTPError(statusCode int, message string) models.GenericHTTPErrorResponse {
	return models.GenericHTTPErrorResponse{
		Error:   statusCode,
		Message: message,
	}
}

func NewError(errorCode string, errorMessageFormat string, numericErrorCode int, originatingService string, intent string, a ...any) models.ErrorResponse {
	messageVars := make([]string, len(a))
	for i, v := range a {
		messageVars[i] = fmt.Sprintf("%v", v)
	}

	return models.ErrorResponse{
		ErrorCode:          errorCode,
		ErrorMessage:       fmt.Sprintf(errorMessageFormat, a...),
		MessageVars:        messageVars,
		NumericErrorCode:   numericErrorCode,
		OriginatingService: originatingService,
		Intent:             intent,
	}
}

func NewInternalServerError(originatingService string) models.ErrorResponse {
	return NewError(
		"errors.com.transcendent.common.server_error",
		"Sorry an error occurred and we were unable to resolve it (tracking id: 0)",
		1000,
		originatingService,
		buildinfo.Branch,
	)
}

func NewUnsupportedGrantTypeError(originatingService string, grantType string) models.ErrorResponse {
	return NewError(
		"errors.com.transcendent.common.server_error",
		"Malformed grant type in token request: %s",
		1016,
		originatingService,
		buildinfo.Branch,
		grantType,
	)
}
