package models

type TopApplication struct {
	AppId           string    `json:"appId,omitempty"`
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
	Ratio           *int      `json:"ratio,omitempty"`
}
