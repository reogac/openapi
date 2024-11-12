package models

type ServiceExperienceInfo struct {
	Supis            []string              `json:"supis,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
}
