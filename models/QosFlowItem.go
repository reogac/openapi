package models
type QosFlowItem struct {
	 Qfi	int	`json:"qfi"`
	 Cause	Cause	`json:"cause,omitempty"`
	 CurrentQosProfileIndex	*int	`json:"currentQosProfileIndex,omitempty"`
	 NullQoSProfileIndex	*bool	`json:"nullQoSProfileIndex,omitempty"`
	 NgApCause	*NgApCause	`json:"ngApCause,omitempty"`
}
