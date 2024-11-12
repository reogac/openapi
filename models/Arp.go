package models

type Arp struct {
	PreemptVuln   PreemptionVulnerability `json:"preemptVuln"`
	PriorityLevel int                     `json:"priorityLevel"`
	PreemptCap    PreemptionCapability    `json:"preemptCap"`
}
