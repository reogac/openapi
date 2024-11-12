package models

type UeContextCreateError struct {
	TargetToSourceFailureData *N2InfoContent `json:"targetToSourceFailureData,omitempty"`
	Error                     ProblemDetails `json:"error"`
	NgapCause                 *NgApCause     `json:"ngapCause,omitempty"`
}
