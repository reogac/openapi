package models

type MbsSession struct {
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
}
