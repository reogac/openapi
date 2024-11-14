package models
type SscModes struct {
	 DefaultSscMode	SscMode	`json:"defaultSscMode"`
	 AllowedSscModes	[]string	`json:"allowedSscModes,omitempty"`
}
