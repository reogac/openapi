package models

type AccessTokenErr struct {
	Error_description string `json:"error_description,omitempty"`
	Error_uri         string `json:"error_uri,omitempty"`
	Error             Error  `json:"error"`
}
