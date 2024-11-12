package models

type AuthEvent struct {
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	Success                    bool     `json:"success"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthType                   AuthType `json:"authType"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
}
