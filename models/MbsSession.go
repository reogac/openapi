package models

type MbsSession struct {
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
}
