package models

type PatchItem struct {
	From string `json:"from,omitempty"`
	Op   string `json:"op"`
	Path string `json:"path"`
}
