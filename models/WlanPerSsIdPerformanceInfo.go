package models
type WlanPerSsIdPerformanceInfo struct {
	 WlanPerTsInfos	[]WlanPerTsPerformanceInfo	`json:"wlanPerTsInfos"`
	 SsId	string	`json:"ssId"`
}
