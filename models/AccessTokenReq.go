package models
type AccessTokenReq struct {
	 NfType	string	`json:"nfType,omitempty"`
	 RequesterPlmn	*PlmnId	`json:"requesterPlmn,omitempty"`
	 RequesterSnpnList	[]PlmnIdNid	`json:"requesterSnpnList,omitempty"`
	 HnrfAccessTokenUri	string	`json:"hnrfAccessTokenUri,omitempty"`
	 SourceNfInstanceId	string	`json:"sourceNfInstanceId,omitempty"`
	 Grant_type	Grant_type	`json:"grant_type"`
	 RequesterPlmnList	[]PlmnId	`json:"requesterPlmnList,omitempty"`
	 RequesterSnssaiList	[]Snssai	`json:"requesterSnssaiList,omitempty"`
	 RequesterFqdn	string	`json:"requesterFqdn,omitempty"`
	 TargetPlmn	*PlmnId	`json:"targetPlmn,omitempty"`
	 TargetSnpn	*PlmnIdNid	`json:"targetSnpn,omitempty"`
	 TargetNfSetId	string	`json:"targetNfSetId,omitempty"`
	 Scope	string	`json:"scope"`
	 TargetNfInstanceId	string	`json:"targetNfInstanceId,omitempty"`
	 TargetSnssaiList	[]Snssai	`json:"targetSnssaiList,omitempty"`
	 TargetNsiList	[]string	`json:"targetNsiList,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 TargetNfType	string	`json:"targetNfType,omitempty"`
	 TargetNfServiceSetId	string	`json:"targetNfServiceSetId,omitempty"`
}
