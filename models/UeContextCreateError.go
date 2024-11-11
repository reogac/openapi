package models

type UeContextCreateError struct {
	NgapCause                 *NgApCause     `json:"ngapCause,omitempty"`
	TargetToSourceFailureData *N2InfoContent `json:"targetToSourceFailureData,omitempty"`
	Error                     ProblemDetails `json:"error"`
}
