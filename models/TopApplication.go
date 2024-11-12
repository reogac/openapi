package models

type TopApplication struct {
	Ratio           *int      `json:"ratio,omitempty"`
	AppId           string    `json:"appId,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
}
