package models

type WlanPerformanceInfo struct {
	WlanPerSsidInfos []WlanPerSsIdPerformanceInfo `json:"wlanPerSsidInfos"`
	NetworkArea      *NetworkAreaInfo             `json:"networkArea,omitempty"`
}
