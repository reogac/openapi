package models

type SmContextCreateError struct {
	Error        ExtProblemDetails `json:"error"`
	N1SmMsg      *RefToBinaryData  `json:"n1SmMsg,omitempty"`
	N2SmInfo     *RefToBinaryData  `json:"n2SmInfo,omitempty"`
	N2SmInfoType N2SmInfoType      `json:"n2SmInfoType,omitempty"`
	RecoveryTime string            `json:"recoveryTime,omitempty"`
}
