package response

import (
	"github.com/mlctrez/lexstream/amsapi"
	"github.com/mlctrez/lexstream/amsapi/header"
	"github.com/mlctrez/lexstream/amsapi/types"
)

// AlexaError returns a general error response body.
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-error-response.html#general-errors
func AlexaError(type_ types.ErrorType, headerMessageId string, message string, subtypes ...types.ErrorType) *amsapi.Response {
	return &amsapi.Response{
		Header: &header.Header{
			Namespace: "Alexa", Name: "ErrorResponse",
			MessageId: headerMessageId, PayloadVersion: "3.0",
		},
		Payload: ErrorPayload{Type: type_, SubTypes: subtypes, Message: message},
	}
}

// MediaError returns a media specific response body.
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-error-response.html#media-specific-errors
func MediaError(type_ types.ErrorType, headerMessageId string, message string, subtypes ...types.ErrorType) *amsapi.Response {
	return &amsapi.Response{
		Header: &header.Header{
			Namespace: "Alexa.Media", Name: "ErrorResponse",
			MessageId: headerMessageId, PayloadVersion: "1.0",
		},
		Payload: ErrorPayload{Type: type_, SubTypes: subtypes, Message: message},
	}
}

// AudioError returns audio specific response body.
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-error-response.html#audio-specific-errors
func AudioError(type_ types.ErrorType, headerMessageId string, message string, retryPeriod types.AudioSkipRetryPeriod) *amsapi.Response {
	return &amsapi.Response{
		Header: &header.Header{
			Namespace: "Alexa.Audio", Name: "ErrorResponse",
			MessageId: headerMessageId, PayloadVersion: "1.0",
		},
		Payload: ErrorPayload{Type: type_, RetryPeriod: retryPeriod, Message: message},
	}
}

// ErrorPayload is the response payload used by AlexaError, MediaError, and AudioError.
// SubTypes and RetryPeriod are only valid for certain error types.
type ErrorPayload struct {
	Type        types.ErrorType            `json:"type"`
	SubTypes    []types.ErrorType          `json:"subtypes,omitempty"`
	Message     string                     `json:"message"`
	RetryPeriod types.AudioSkipRetryPeriod `json:"retryPeriod,omitempty"`
}
