package models

type ServiceExperienceInfo struct {
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Dnai             string                `json:"dnai,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
}
