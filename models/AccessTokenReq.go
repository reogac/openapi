package models

type AccessTokenReq struct {
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	Scope                string      `json:"scope"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	NfType               NFType      `json:"nfType,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	GrantType            GrantType   `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
}
