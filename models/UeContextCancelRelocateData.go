package models

type UeContextCancelRelocateData struct {
	RelocationCancelRequest RefToBinaryData `json:"relocationCancelRequest"`
	Supi                    string          `json:"supi,omitempty"`
}
