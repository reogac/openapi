package models

type AccessTokenReq struct {
	NfType               NFType      `json:"nfType,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	Scope                string      `json:"scope"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
}
