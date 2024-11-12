package models
type NotifyItem struct {
	 ResourceId	string	`json:"resourceId"`
	 Changes	[]ChangeItem	`json:"changes"`
}
