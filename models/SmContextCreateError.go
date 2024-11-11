type SmContextCreateError struct {
	N1SmMsg      *RefToBinaryData  `json:"n1SmMsg,omitempty"`
	N2SmInfo     *RefToBinaryData  `json:"n2SmInfo,omitempty"`
	N2SmInfoType string            `json:"n2SmInfoType,omitempty"`
	RecoveryTime string            `json:"recoveryTime,omitempty"`
	Error        ExtProblemDetails `json:"error"`
}
