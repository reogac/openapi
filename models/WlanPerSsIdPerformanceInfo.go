package models
type WlanPerSsIdPerformanceInfo struct {
	 SsId	string	`json:"ssId"`
	 WlanPerTsInfos	[]WlanPerTsPerformanceInfo	`json:"wlanPerTsInfos"`
}
