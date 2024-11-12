package models
type AccessTokenErr struct {
	 ErrorDescription	string	`json:"error_description,omitempty"`
	 ErrorUri	string	`json:"error_uri,omitempty"`
	 Error	Error	`json:"error"`
}
