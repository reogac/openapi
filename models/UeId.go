package models

type UeId struct {
	GpsiList []string `json:"gpsiList,omitempty"`
	Supi     string   `json:"supi"`
}
