package models
type AccessTokenErr struct {
	 Error_uri	string	`json:"error_uri,omitempty"`
	 Error	string	`json:"error"`
	 Error_description	string	`json:"error_description,omitempty"`
}
