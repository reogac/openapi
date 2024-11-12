package models

type PtwParameters struct {
	PtwValue         string        `json:"ptwValue"`
	ExtendedPtwValue string        `json:"extendedPtwValue,omitempty"`
	OperationMode    OperationMode `json:"operationMode"`
}
