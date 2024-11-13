package models

type NwdafSubscription struct {
	NwdafEventsSubscription NnwdafEventsSubscription `json:"nwdafEventsSubscription"`
	NwdafEvtSubsServiceUri  string                   `json:"nwdafEvtSubsServiceUri"`
}
