package models

type MbsSession struct {
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
}
