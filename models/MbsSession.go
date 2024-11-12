package models

type MbsSession struct {
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
}
