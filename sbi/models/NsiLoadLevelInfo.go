/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiLoadLevelInfo struct {
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	Snssai                   Snssai           `json:"snssai"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NsiId                    string           `json:"nsiId,omitempty"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
}
