package models

type ExtendedSmSubsData struct {
	IndividualSmSubsData []SessionManagementSubscriptionData `json:"individualSmSubsData,omitempty"`
	SharedSmSubsDataIds  []string                            `json:"sharedSmSubsDataIds"`
}
