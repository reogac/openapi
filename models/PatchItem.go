package models

type PatchItem struct {
	Op   PatchOperation `json:"op"`
	Path string         `json:"path"`
	From string         `json:"from,omitempty"`
}
