package models

type PcscfRestorationNotification struct {
	Supi        string        `json:"supi"`
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
}
