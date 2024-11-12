package models

type AccessTokenReq struct {
	TargetSnpn           *PlmnIdNid  `json:"targetSnpn,omitempty"`
	TargetNfSetId        string      `json:"targetNfSetId,omitempty"`
	NfType               string      `json:"nfType,omitempty"`
	TargetNfInstanceId   string      `json:"targetNfInstanceId,omitempty"`
	RequesterPlmn        *PlmnId     `json:"requesterPlmn,omitempty"`
	RequesterSnpnList    []PlmnIdNid `json:"requesterSnpnList,omitempty"`
	TargetPlmn           *PlmnId     `json:"targetPlmn,omitempty"`
	RequesterFqdn        string      `json:"requesterFqdn,omitempty"`
	TargetSnssaiList     []Snssai    `json:"targetSnssaiList,omitempty"`
	Grant_type           Grant_type  `json:"grant_type"`
	NfInstanceId         string      `json:"nfInstanceId"`
	TargetNfType         string      `json:"targetNfType,omitempty"`
	Scope                string      `json:"scope"`
	RequesterSnssaiList  []Snssai    `json:"requesterSnssaiList,omitempty"`
	RequesterPlmnList    []PlmnId    `json:"requesterPlmnList,omitempty"`
	TargetNfServiceSetId string      `json:"targetNfServiceSetId,omitempty"`
	SourceNfInstanceId   string      `json:"sourceNfInstanceId,omitempty"`
	TargetNsiList        []string    `json:"targetNsiList,omitempty"`
	HnrfAccessTokenUri   string      `json:"hnrfAccessTokenUri,omitempty"`
}
