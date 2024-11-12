package models

type MbsSession struct {
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
}
