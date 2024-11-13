package models

type SubscriptionData struct {
	GuamiList    []Guami `json:"guamiList,omitempty"`
	AmfStatusUri string  `json:"amfStatusUri"`
}
