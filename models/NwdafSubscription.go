package models

type NwdafSubscription struct {
	NwdafEvtSubsServiceUri  string                   `json:"nwdafEvtSubsServiceUri"`
	NwdafEventsSubscription NnwdafEventsSubscription `json:"nwdafEventsSubscription"`
}
