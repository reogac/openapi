package models

type NsiLoadLevelInfo struct {
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	Snssai                   Snssai           `json:"snssai"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NsiId                    string           `json:"nsiId,omitempty"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
}
