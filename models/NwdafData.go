package models

type NwdafData struct {
	NwdafInstanceId string   `json:"nwdafInstanceId"`
	NwdafEvents     []string `json:"nwdafEvents,omitempty"`
}
