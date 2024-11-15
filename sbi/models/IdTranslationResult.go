/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IdTranslationResult struct {
	Gpsi              string   `json:"gpsi,omitempty"`
	AdditionalSupis   []string `json:"additionalSupis,omitempty"`
	AdditionalGpsis   []string `json:"additionalGpsis,omitempty"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty"`
	Supi              string   `json:"supi"`
}
