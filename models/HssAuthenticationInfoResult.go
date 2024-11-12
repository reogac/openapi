package models

type HssAuthenticationInfoResult struct {
	SupportedFeatures        string                   `json:"supportedFeatures,omitempty"`
	HssAuthenticationVectors HssAuthenticationVectors `json:"hssAuthenticationVectors"`
}
