package models

type AccessTokenReq struct {
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	Scope                string      `json:"scope"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	NfType               string      `json:"nfType,omitempty"`
	TargetNfType         string      `json:"targetNfType,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	Grant_type           string      `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
}
