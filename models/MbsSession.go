package models
type MbsSession struct {
	 MbsServInfo	*MbsServiceInfo	`json:"mbsServInfo,omitempty"`
	 ActivityStatus	MbsSessionActivityStatus	`json:"activityStatus,omitempty"`
	 AnyUeInd	*bool	`json:"anyUeInd,omitempty"`
	 TmgiAllocReq	*bool	`json:"tmgiAllocReq,omitempty"`
	 ExpirationTime	string	`json:"expirationTime,omitempty"`
	 ServiceType	MbsServiceType	`json:"serviceType"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 RedMbsServArea	*MbsServiceArea	`json:"redMbsServArea,omitempty"`
	 ActivationTime	string	`json:"activationTime,omitempty"`
	 Tmgi	*Tmgi	`json:"tmgi,omitempty"`
	 AreaSessionId	*int	`json:"areaSessionId,omitempty"`
	 IngressTunAddr	[]TunnelAddress	`json:"ingressTunAddr,omitempty"`
	 Ssm	*Ssm	`json:"ssm,omitempty"`
	 MbsFsaIdList	[]string	`json:"mbsFsaIdList,omitempty"`
	 LocationDependent	*bool	`json:"locationDependent,omitempty"`
	 ExtMbsServiceArea	*ExternalMbsServiceArea	`json:"extMbsServiceArea,omitempty"`
	 ExtRedMbsServArea	*ExternalMbsServiceArea	`json:"extRedMbsServArea,omitempty"`
	 TerminationTime	string	`json:"terminationTime,omitempty"`
	 StartTime	string	`json:"startTime,omitempty"`
	 MbsSessionSubsc	*MbsSessionSubscription	`json:"mbsSessionSubsc,omitempty"`
	 MbsSessionId	*MbsSessionId	`json:"mbsSessionId,omitempty"`
	 IngressTunAddrReq	*bool	`json:"ingressTunAddrReq,omitempty"`
	 MbsServiceArea	*MbsServiceArea	`json:"mbsServiceArea,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
}
