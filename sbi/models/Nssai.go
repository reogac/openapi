/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Nssai struct {
	SingleNssais         []Snssai                        `json:"singleNssais,omitempty"`
	ProvisioningTime     string                          `json:"provisioningTime,omitempty"`
	AdditionalSnssaiData map[string]AdditionalSnssaiData `json:"additionalSnssaiData,omitempty"`
	SuppressNssrgInd     *bool                           `json:"suppressNssrgInd,omitempty"`
	SupportedFeatures    string                          `json:"supportedFeatures,omitempty"`
	DefaultSingleNssais  []Snssai                        `json:"defaultSingleNssais"`
}
