/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Dnai             string                `json:"dnai,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
}
