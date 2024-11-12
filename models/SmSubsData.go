package models

type SmSubsData struct {
	ExtendedSmSubsData                *ExtendedSmSubsData                 `json:"ExtendedSmSubsData,omitempty"`
	SessionManagementSubscriptionData []SessionManagementSubscriptionData `json:"SessionManagementSubscriptionData,omitempty"`
}
