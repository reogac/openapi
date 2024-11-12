package models

type UpConfidentiality string

// Define constant values for UpConfidentiality
const (
	UPCONFIDENTIALITY_REQUIRED   UpConfidentiality = "REQUIRED"
	UPCONFIDENTIALITY_PREFERRED  UpConfidentiality = "PREFERRED"
	UPCONFIDENTIALITY_NOT_NEEDED UpConfidentiality = "NOT_NEEDED"
)
