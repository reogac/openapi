package models
type AccessTokenReq struct {
	 TargetNfType	NFType	`json:"targetNfType,omitempty"`
	 Scope	string	`json:"scope"`
	 RequesterPlmn	*PlmnId	`json:"requesterPlmn,omitempty"`
	 RequesterSnpnList	[]PlmnIdNid	`json:"requesterSnpnList,omitempty"`
	 TargetNsiList	[]string	`json:"targetNsiList,omitempty"`
	 TargetNfServiceSetId	string	`json:"targetNfServiceSetId,omitempty"`
	 GrantType	GrantType	`json:"grant_type"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 TargetSnssaiList	[]Snssai	`json:"targetSnssaiList,omitempty"`
	 TargetNfSetId	string	`json:"targetNfSetId,omitempty"`
	 RequesterPlmnList	[]PlmnId	`json:"requesterPlmnList,omitempty"`
	 TargetSnpn	*PlmnIdNid	`json:"targetSnpn,omitempty"`
	 SourceNfInstanceId	string	`json:"sourceNfInstanceId,omitempty"`
	 NfType	NFType	`json:"nfType,omitempty"`
	 TargetPlmn	*PlmnId	`json:"targetPlmn,omitempty"`
	 RequesterFqdn	string	`json:"requesterFqdn,omitempty"`
	 HnrfAccessTokenUri	string	`json:"hnrfAccessTokenUri,omitempty"`
	 TargetNfInstanceId	string	`json:"targetNfInstanceId,omitempty"`
	 RequesterSnssaiList	[]Snssai	`json:"requesterSnssaiList,omitempty"`
}
