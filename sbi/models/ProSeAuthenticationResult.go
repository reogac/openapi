/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationResult struct {
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	KnrProSe          string `json:"knrProSe,omitempty"`
	Nonce2            string `json:"nonce2,omitempty"`
}
