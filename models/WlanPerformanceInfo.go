package models

type WlanPerformanceInfo struct {
	NetworkArea      *NetworkAreaInfo             `json:"networkArea,omitempty"`
	WlanPerSsidInfos []WlanPerSsIdPerformanceInfo `json:"wlanPerSsidInfos"`
}
