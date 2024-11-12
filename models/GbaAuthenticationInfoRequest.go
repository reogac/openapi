package models

type GbaAuthenticationInfoRequest struct {
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	AuthType              GbaAuthType            `json:"authType"`
}
