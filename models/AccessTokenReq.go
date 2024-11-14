package models
type AccessTokenReq struct {
	 Scope	string	`json:"scope"`
	 RequesterSnssaiList	[]Snssai	`json:"requesterSnssaiList,omitempty"`
	 TargetSnssaiList	[]Snssai	`json:"targetSnssaiList,omitempty"`
	 TargetNfSetId	string	`json:"targetNfSetId,omitempty"`
	 NfType	NFType	`json:"nfType,omitempty"`
	 TargetNfInstanceId	string	`json:"targetNfInstanceId,omitempty"`
	 RequesterSnpnList	[]PlmnIdNid	`json:"requesterSnpnList,omitempty"`
	 HnrfAccessTokenUri	string	`json:"hnrfAccessTokenUri,omitempty"`
	 SourceNfInstanceId	string	`json:"sourceNfInstanceId,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 RequesterFqdn	string	`json:"requesterFqdn,omitempty"`
	 TargetSnpn	*PlmnIdNid	`json:"targetSnpn,omitempty"`
	 TargetNfServiceSetId	string	`json:"targetNfServiceSetId,omitempty"`
	 GrantType	GrantType	`json:"grant_type"`
	 TargetNfType	NFType	`json:"targetNfType,omitempty"`
	 RequesterPlmn	*PlmnId	`json:"requesterPlmn,omitempty"`
	 RequesterPlmnList	[]PlmnId	`json:"requesterPlmnList,omitempty"`
	 TargetPlmn	*PlmnId	`json:"targetPlmn,omitempty"`
	 TargetNsiList	[]string	`json:"targetNsiList,omitempty"`
}
