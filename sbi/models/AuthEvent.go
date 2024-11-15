/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:38 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthEvent struct {
	NfSetId                    string   `json:"nfSetId,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	Success                    bool     `json:"success"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthType                   AuthType `json:"authType"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
}
