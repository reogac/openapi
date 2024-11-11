package models

type AccessTokenReq struct {
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	Scope                string      `json:"scope"`
	TargetNfType         string      `json:"targetNfType,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	NfType               string      `json:"nfType,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	Grant_type           string      `json:"grant_type"`
}
