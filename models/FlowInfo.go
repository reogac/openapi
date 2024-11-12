package models

type FlowInfo struct {
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
	FlowId           int      `json:"flowId"`
}
