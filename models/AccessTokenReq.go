package models

type AccessTokenReq struct {
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetNfType         string      `json:"targetNfType,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	Grant_type           string      `json:"grant_type"`
	NfType               string      `json:"nfType,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	Scope                string      `json:"scope"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
}
