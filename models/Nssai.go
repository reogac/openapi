package models

type Nssai struct {
	AdditionalSnssaiData map[string]AdditionalSnssaiData `json:"additionalSnssaiData,omitempty"`
	SuppressNssrgInd     *bool                           `json:"suppressNssrgInd,omitempty"`
	SupportedFeatures    string                          `json:"supportedFeatures,omitempty"`
	DefaultSingleNssais  []Snssai                        `json:"defaultSingleNssais"`
	SingleNssais         []Snssai                        `json:"singleNssais,omitempty"`
	ProvisioningTime     string                          `json:"provisioningTime,omitempty"`
}