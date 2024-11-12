package models

type SubscriptionData struct {
	AmfStatusUri string  `json:"amfStatusUri"`
	GuamiList    []Guami `json:"guamiList,omitempty"`
}
