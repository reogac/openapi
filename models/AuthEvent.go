/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AuthEvent struct {
	Success                    bool     `json:"success"`
	AuthType                   AuthType `json:"authType"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	TimeStamp                  string   `json:"timeStamp"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
}
