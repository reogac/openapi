package models

type ServiceExperienceInfo struct {
	NetworkArea      *NetworkAreaInfo    `json:"networkArea,omitempty"`
	RatFreq          *RatFreqInformation `json:"ratFreq,omitempty"`
	SvcExprc         SvcExperience       `json:"svcExprc"`
	SvcExprcVariance *float64            `json:"svcExprcVariance,omitempty"`
	Dnai             string              `json:"dnai,omitempty"`
	Dnn              string              `json:"dnn,omitempty"`
	NsiId            string              `json:"nsiId,omitempty"`
	Supis            []string            `json:"supis,omitempty"`
	UpfInfo          *UpfInformation     `json:"upfInfo,omitempty"`
	AppServerInst    *AddrFqdn           `json:"appServerInst,omitempty"`
	Snssai           *Snssai             `json:"snssai,omitempty"`
	Confidence       *int                `json:"confidence,omitempty"`
	Ratio            *int                `json:"ratio,omitempty"`
	AppId            string              `json:"appId,omitempty"`
	SrvExpcType      string              `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo      `json:"ueLocs,omitempty"`
}
