package models
type AccessTokenErr struct {
	 Error	Error	`json:"error"`
	 ErrorDescription	string	`json:"error_description,omitempty"`
	 ErrorUri	string	`json:"error_uri,omitempty"`
}
