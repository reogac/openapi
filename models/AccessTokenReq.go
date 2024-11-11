type AccessTokenReq struct {
	 TargetPlmn	*PlmnId	`json:"targetPlmn,omitempty"`
	 TargetSnssaiList	[]Snssai	`json:"targetSnssaiList,omitempty"`
	 NfType	string	`json:"nfType,omitempty"`
	 TargetNfInstanceId	string	`json:"targetNfInstanceId,omitempty"`
	 RequesterPlmnList	[]PlmnId	`json:"requesterPlmnList,omitempty"`
	 Grant_type	string	`json:"grant_type"`
	 TargetNsiList	[]string	`json:"targetNsiList,omitempty"`
	 TargetNfSetId	string	`json:"targetNfSetId,omitempty"`
	 TargetNfServiceSetId	string	`json:"targetNfServiceSetId,omitempty"`
	 NfInstanceId	string	`json:"nfInstanceId"`
	 RequesterPlmn	*PlmnId	`json:"requesterPlmn,omitempty"`
	 RequesterSnssaiList	[]Snssai	`json:"requesterSnssaiList,omitempty"`
	 RequesterSnpnList	[]PlmnIdNid	`json:"requesterSnpnList,omitempty"`
	 TargetNfType	string	`json:"targetNfType,omitempty"`
	 Scope	string	`json:"scope"`
	 RequesterFqdn	string	`json:"requesterFqdn,omitempty"`
}
