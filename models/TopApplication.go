package models

type TopApplication struct {
	IpTrafficFilter *FlowInfo `json:"ipTrafficFilter,omitempty"`
	Ratio           *int      `json:"ratio,omitempty"`
	AppId           string    `json:"appId,omitempty"`
}
