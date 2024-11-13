package models

type Arp struct {
	PreemptCap    PreemptionCapability    `json:"preemptCap"`
	PreemptVuln   PreemptionVulnerability `json:"preemptVuln"`
	PriorityLevel int                     `json:"priorityLevel"`
}
