package models
type TriggerRequest struct {
	 Supi	string	`json:"supi"`
	 FailedPcscf	*PcscfAddress	`json:"failedPcscf,omitempty"`
}
