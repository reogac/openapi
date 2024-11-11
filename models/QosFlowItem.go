package models

type QosFlowItem struct {
	NullQoSProfileIndex    *bool      `json:"nullQoSProfileIndex,omitempty"`
	NgApCause              *NgApCause `json:"ngApCause,omitempty"`
	Qfi                    int        `json:"qfi"`
	Cause                  string     `json:"cause,omitempty"`
	CurrentQosProfileIndex *int       `json:"currentQosProfileIndex,omitempty"`
}
