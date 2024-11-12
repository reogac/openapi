package models

type AuthEvent struct {
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthType                   AuthType `json:"authType"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	Success                    bool     `json:"success"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
}
