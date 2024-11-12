package models

type NotifyItem struct {
	Changes    []ChangeItem `json:"changes"`
	ResourceId string       `json:"resourceId"`
}
