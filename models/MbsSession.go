package models

type MbsSession struct {
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
}
