package models

type ServiceExperienceInfo struct {
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
}
