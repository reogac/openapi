package models

type AuthEvent struct {
	TimeStamp                  string   `json:"timeStamp"`
	AuthType                   AuthType `json:"authType"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	Success                    bool     `json:"success"`
}
