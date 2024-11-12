package models
type NsiLoadLevelInfo struct {
	 NumOfExceedLoadLevelThr	*int	`json:"numOfExceedLoadLevelThr,omitempty"`
	 ExceedLoadLevelThrInd	*bool	`json:"exceedLoadLevelThrInd,omitempty"`
	 NetworkArea	*NetworkAreaInfo	`json:"networkArea,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 LoadLevelInformation	int	`json:"loadLevelInformation"`
	 Snssai	Snssai	`json:"snssai"`
	 ResUsage	*ResourceUsage	`json:"resUsage,omitempty"`
	 NumOfUes	*NumberAverage	`json:"numOfUes,omitempty"`
	 NumOfPduSess	*NumberAverage	`json:"numOfPduSess,omitempty"`
	 NsiId	string	`json:"nsiId,omitempty"`
	 TimePeriod	*TimeWindow	`json:"timePeriod,omitempty"`
	 ResUsgThrCrossTimePeriod	[]TimeWindow	`json:"resUsgThrCrossTimePeriod,omitempty"`
}
