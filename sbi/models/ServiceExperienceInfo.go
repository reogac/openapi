/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Supis            []string              `json:"supis,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
}
