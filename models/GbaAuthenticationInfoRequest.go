package models

type GbaAuthenticationInfoRequest struct {
	AuthType              GbaAuthType            `json:"authType"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
}
