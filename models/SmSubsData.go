package models

type SmSubsData struct {
	SessionManagementSubscriptionData []SessionManagementSubscriptionData `json:"SessionManagementSubscriptionData,omitempty"`
	ExtendedSmSubsData                *ExtendedSmSubsData                 `json:"ExtendedSmSubsData,omitempty"`
}
