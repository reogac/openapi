package models

type PtwParameters struct {
	OperationMode    OperationMode `json:"operationMode"`
	PtwValue         string        `json:"ptwValue"`
	ExtendedPtwValue string        `json:"extendedPtwValue,omitempty"`
}
