package models
type MbsSession struct {
	 TmgiAllocReq	*bool	`json:"tmgiAllocReq,omitempty"`
	 ExpirationTime	string	`json:"expirationTime,omitempty"`
	 AreaSessionId	*int	`json:"areaSessionId,omitempty"`
	 Ssm	*Ssm	`json:"ssm,omitempty"`
	 ActivationTime	string	`json:"activationTime,omitempty"`
	 MbsSessionSubsc	*MbsSessionSubscription	`json:"mbsSessionSubsc,omitempty"`
	 MbsSessionId	*MbsSessionId	`json:"mbsSessionId,omitempty"`
	 ServiceType	MbsServiceType	`json:"serviceType"`
	 IngressTunAddrReq	*bool	`json:"ingressTunAddrReq,omitempty"`
	 ExtRedMbsServArea	*ExternalMbsServiceArea	`json:"extRedMbsServArea,omitempty"`
	 TerminationTime	string	`json:"terminationTime,omitempty"`
	 StartTime	string	`json:"startTime,omitempty"`
	 MbsServInfo	*MbsServiceInfo	`json:"mbsServInfo,omitempty"`
	 LocationDependent	*bool	`json:"locationDependent,omitempty"`
	 MbsServiceArea	*MbsServiceArea	`json:"mbsServiceArea,omitempty"`
	 ExtMbsServiceArea	*ExternalMbsServiceArea	`json:"extMbsServiceArea,omitempty"`
	 RedMbsServArea	*MbsServiceArea	`json:"redMbsServArea,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 ActivityStatus	MbsSessionActivityStatus	`json:"activityStatus,omitempty"`
	 AnyUeInd	*bool	`json:"anyUeInd,omitempty"`
	 MbsFsaIdList	[]string	`json:"mbsFsaIdList,omitempty"`
	 Tmgi	*Tmgi	`json:"tmgi,omitempty"`
	 IngressTunAddr	[]TunnelAddress	`json:"ingressTunAddr,omitempty"`
}
