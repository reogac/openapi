package models

type PcscfRestorationNotification struct {
	FailedPcscf *PcscfAddress `json:"failedPcscf,omitempty"`
	Supi        string        `json:"supi"`
}
