/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProseSubscriptionData struct {
	ProseAllowedPlmn []ProSeAllowedPlmn `json:"proseAllowedPlmn,omitempty"`
	ProseServiceAuth *ProseServiceAuth  `json:"proseServiceAuth,omitempty"`
	NrUePc5Ambr      string             `json:"nrUePc5Ambr,omitempty"`
}
