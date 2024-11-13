package models

type ServiceExperienceInfo struct {
	Dnn              string                `json:"dnn,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
}
