package models

type UeId struct {
	Supi     string   `json:"supi"`
	GpsiList []string `json:"gpsiList,omitempty"`
}
