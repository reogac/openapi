package models
type NsiLoadLevelInfo struct {
	 NsiId	string	`json:"nsiId,omitempty"`
	 NumOfExceedLoadLevelThr	*int	`json:"numOfExceedLoadLevelThr,omitempty"`
	 NumOfPduSess	*NumberAverage	`json:"numOfPduSess,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 LoadLevelInformation	int	`json:"loadLevelInformation"`
	 ResUsage	*ResourceUsage	`json:"resUsage,omitempty"`
	 ExceedLoadLevelThrInd	*bool	`json:"exceedLoadLevelThrInd,omitempty"`
	 NetworkArea	*NetworkAreaInfo	`json:"networkArea,omitempty"`
	 TimePeriod	*TimeWindow	`json:"timePeriod,omitempty"`
	 ResUsgThrCrossTimePeriod	[]TimeWindow	`json:"resUsgThrCrossTimePeriod,omitempty"`
	 NumOfUes	*NumberAverage	`json:"numOfUes,omitempty"`
	 Snssai	Snssai	`json:"snssai"`
}
