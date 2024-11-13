package models

type QosFlowItem struct {
	NgApCause              *NgApCause `json:"ngApCause,omitempty"`
	Qfi                    int        `json:"qfi"`
	Cause                  Cause      `json:"cause,omitempty"`
	CurrentQosProfileIndex *int       `json:"currentQosProfileIndex,omitempty"`
	NullQoSProfileIndex    *bool      `json:"nullQoSProfileIndex,omitempty"`
}
