/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Dnai             string                `json:"dnai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
}
