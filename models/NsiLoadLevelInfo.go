package models

type NsiLoadLevelInfo struct {
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	NsiId                    string           `json:"nsiId,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
}
