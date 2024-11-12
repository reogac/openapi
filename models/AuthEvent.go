package models

type AuthEvent struct {
	UdrRestartInd              *bool    `json:"udrRestartInd,omitempty"`
	TimeStamp                  string   `json:"timeStamp"`
	AuthType                   AuthType `json:"authType"`
	ServingNetworkName         string   `json:"servingNetworkName"`
	NfSetId                    string   `json:"nfSetId,omitempty"`
	ResetIds                   []string `json:"resetIds,omitempty"`
	DataRestorationCallbackUri string   `json:"dataRestorationCallbackUri,omitempty"`
	NfInstanceId               string   `json:"nfInstanceId"`
	Success                    bool     `json:"success"`
	AuthRemovalInd             *bool    `json:"authRemovalInd,omitempty"`
}
