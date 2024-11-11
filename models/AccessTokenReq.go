package models

type AccessTokenReq struct {
	Scope                string      `json:"scope"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetNfType         string      `json:"targetNfType,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	Grant_type           string      `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	NfType               string      `json:"nfType,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
}
