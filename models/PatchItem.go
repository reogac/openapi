package models

type PatchItem struct {
	From string         `json:"from,omitempty"`
	Op   PatchOperation `json:"op"`
	Path string         `json:"path"`
}
