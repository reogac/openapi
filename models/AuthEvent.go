package models

type AuthEvent struct {
	Success                    bool     `json:"success"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthType                   AuthType `json:"authType"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
}
