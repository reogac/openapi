/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UEAuthenticationCtx struct {
	Links              map[string]Link `json:"_links"`
	ServingNetworkName string          `json:"servingNetworkName,omitempty"`
	AuthType           AuthType        `json:"authType"`
	FiveGAuthData      FiveGAuthData   `json:"5gAuthData"`
}
