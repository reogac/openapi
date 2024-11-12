package models

type AuthEvent struct {
	ServingNetworkName         string   `json:"servingNetworkName"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	Success                    bool     `json:"success"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	AuthType                   AuthType `json:"authType"`
}
