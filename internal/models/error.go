package models

type ErrorResponse struct {
	ErrorCode          string   `json:"errorCode"`
	ErrorMessage       string   `json:"errorMessage"`
	NumericErrorCode   int      `json:"numericErrorCode"`
	MessageVars        []string `json:"messageVars"`
	OriginatingService string   `json:"originatingService"`
	Intent             string   `json:"intent"`
}

type GenericHTTPErrorResponse struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}
