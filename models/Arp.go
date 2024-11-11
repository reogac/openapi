package models

type Arp struct {
	PreemptVuln   string `json:"preemptVuln"`
	PriorityLevel int    `json:"priorityLevel"`
	PreemptCap    string `json:"preemptCap"`
}
