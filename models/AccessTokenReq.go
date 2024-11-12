package models

type AccessTokenReq struct {
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	NfType               NFType      `json:"nfType,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetNfType         NFType      `json:"targetNfType,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	Grant_type           Grant_type  `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	Scope                string      `json:"scope"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
}
