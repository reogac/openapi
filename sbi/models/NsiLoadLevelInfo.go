/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiLoadLevelInfo struct {
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	NsiId                    string           `json:"nsiId,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
}
