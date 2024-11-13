package models

type AuthEvent struct {
	NfInstanceId               string   `json:"nfInstanceId"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	Success                    bool     `json:"success"`
	AuthType                   AuthType `json:"authType"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
}
