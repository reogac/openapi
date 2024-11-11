package models

type UeContextCreateError struct {
	Error                     ProblemDetails `json:"error"`
	NgapCause                 *NgApCause     `json:"ngapCause,omitempty"`
	TargetToSourceFailureData *N2InfoContent `json:"targetToSourceFailureData,omitempty"`
}
