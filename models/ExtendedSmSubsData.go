package models
type ExtendedSmSubsData struct {
	 SharedSmSubsDataIds	[]string	`json:"sharedSmSubsDataIds"`
	 IndividualSmSubsData	[]SessionManagementSubscriptionData	`json:"individualSmSubsData,omitempty"`
}
