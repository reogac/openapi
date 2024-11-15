/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
}
