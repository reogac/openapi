package models

type ChangeItem struct {
	Op   ChangeType `json:"op"`
	Path string     `json:"path"`
	From string     `json:"from,omitempty"`
}
