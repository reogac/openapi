package models

type NsiLoadLevelInfo struct {
	NsiId                    string           `json:"nsiId,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	Snssai                   Snssai           `json:"snssai"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
}
