package models

type MbsSession struct {
	ActivationTime    string                   `json:"activationTime,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
}
