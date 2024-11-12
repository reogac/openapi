package models

type AuthEvent struct {
	TimeStamp                  string   `json:"timeStamp"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	Success                    bool     `json:"success"`
	AuthType                   AuthType `json:"authType"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
}
