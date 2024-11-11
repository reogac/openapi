package models

type MbsSession struct {
	MbsSessionId      *MbsSessionId           `json:"mbsSessionId,omitempty"`
	ExpirationTime    string                  `json:"expirationTime,omitempty"`
	ServiceType       string                  `json:"serviceType"`
	AreaSessionId     *int                    `json:"areaSessionId,omitempty"`
	MbsServInfo       *MbsServiceInfo         `json:"mbsServInfo,omitempty"`
	AnyUeInd          *bool                   `json:"anyUeInd,omitempty"`
	IngressTunAddrReq *bool                   `json:"ingressTunAddrReq,omitempty"`
	Ssm               *Ssm                    `json:"ssm,omitempty"`
	MbsServiceArea    *MbsServiceArea         `json:"mbsServiceArea,omitempty"`
	Dnn               string                  `json:"dnn,omitempty"`
	Snssai            *Snssai                 `json:"snssai,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription `json:"mbsSessionSubsc,omitempty"`
	TmgiAllocReq      *bool                   `json:"tmgiAllocReq,omitempty"`
	Tmgi              *Tmgi                   `json:"tmgi,omitempty"`
	RedMbsServArea    *MbsServiceArea         `json:"redMbsServArea,omitempty"`
	ActivationTime    string                  `json:"activationTime,omitempty"`
	ActivityStatus    string                  `json:"activityStatus,omitempty"`
	LocationDependent *bool                   `json:"locationDependent,omitempty"`
	IngressTunAddr    []TunnelAddress         `json:"ingressTunAddr,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea `json:"extMbsServiceArea,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea `json:"extRedMbsServArea,omitempty"`
	StartTime         string                  `json:"startTime,omitempty"`
	TerminationTime   string                  `json:"terminationTime,omitempty"`
	MbsFsaIdList      []string                `json:"mbsFsaIdList,omitempty"`
}
