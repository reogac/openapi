package models
type Trigger struct {
	 MaxNumberOfccc	*int	`json:"maxNumberOfccc,omitempty"`
	 TariffTimeChange	string	`json:"tariffTimeChange,omitempty"`
	 TriggerType	string	`json:"triggerType"`
	 TriggerCategory	string	`json:"triggerCategory"`
	 TimeLimit	*int	`json:"timeLimit,omitempty"`
	 VolumeLimit	*int	`json:"volumeLimit,omitempty"`
	 VolumeLimit64	*int	`json:"volumeLimit64,omitempty"`
	 EventLimit	*int	`json:"eventLimit,omitempty"`
}