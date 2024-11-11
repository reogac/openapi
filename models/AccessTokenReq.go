package models
type AccessTokenReq struct {
	 TargetNfSetId	string	`json:"targetNfSetId,omitempty"`
	 Grant_type	string	`json:"grant_type"`
	 NfType	string	`json:"nfType,omitempty"`
	 RequesterPlmn	*PlmnId	`json:"requesterPlmn,omitempty"`
	 HnrfAccessTokenUri	string	`json:"hnrfAccessTokenUri,omitempty"`
	 SourceNfInstanceId	string	`json:"sourceNfInstanceId,omitempty"`
	 TargetNfType	string	`json:"targetNfType,omitempty"`
	 TargetNfInstanceId	string	`json:"targetNfInstanceId,omitempty"`
	 RequesterFqdn	string	`json:"requesterFqdn,omitempty"`
	 RequesterSnpnList	[]PlmnIdNid	`json:"requesterSnpnList,omitempty"`
	 TargetSnpn	*PlmnIdNid	`json:"targetSnpn,omitempty"`
	 TargetNsiList	[]string	`json:"targetNsiList,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 Scope	string	`json:"scope"`
	 RequesterPlmnList	[]PlmnId	`json:"requesterPlmnList,omitempty"`
	 RequesterSnssaiList	[]Snssai	`json:"requesterSnssaiList,omitempty"`
	 TargetPlmn	*PlmnId	`json:"targetPlmn,omitempty"`
	 TargetSnssaiList	[]Snssai	`json:"targetSnssaiList,omitempty"`
	 TargetNfServiceSetId	string	`json:"targetNfServiceSetId,omitempty"`
}
