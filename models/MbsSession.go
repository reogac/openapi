package models
type MbsSession struct {
	 TmgiAllocReq	*bool	`json:"tmgiAllocReq,omitempty"`
	 ExpirationTime	string	`json:"expirationTime,omitempty"`
	 Ssm	*Ssm	`json:"ssm,omitempty"`
	 MbsServiceArea	*MbsServiceArea	`json:"mbsServiceArea,omitempty"`
	 ActivityStatus	string	`json:"activityStatus,omitempty"`
	 MbsSessionId	*MbsSessionId	`json:"mbsSessionId,omitempty"`
	 Tmgi	*Tmgi	`json:"tmgi,omitempty"`
	 AreaSessionId	*int	`json:"areaSessionId,omitempty"`
	 IngressTunAddr	[]TunnelAddress	`json:"ingressTunAddr,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 MbsServInfo	*MbsServiceInfo	`json:"mbsServInfo,omitempty"`
	 ServiceType	string	`json:"serviceType"`
	 ExtMbsServiceArea	*ExternalMbsServiceArea	`json:"extMbsServiceArea,omitempty"`
	 RedMbsServArea	*MbsServiceArea	`json:"redMbsServArea,omitempty"`
	 ExtRedMbsServArea	*ExternalMbsServiceArea	`json:"extRedMbsServArea,omitempty"`
	 MbsSessionSubsc	*MbsSessionSubscription	`json:"mbsSessionSubsc,omitempty"`
	 AnyUeInd	*bool	`json:"anyUeInd,omitempty"`
	 MbsFsaIdList	[]string	`json:"mbsFsaIdList,omitempty"`
	 LocationDependent	*bool	`json:"locationDependent,omitempty"`
	 IngressTunAddrReq	*bool	`json:"ingressTunAddrReq,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 ActivationTime	string	`json:"activationTime,omitempty"`
	 StartTime	string	`json:"startTime,omitempty"`
	 TerminationTime	string	`json:"terminationTime,omitempty"`
}
