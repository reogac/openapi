package models
type AccessTokenErr struct {
	 ErrorUri	string	`json:"error_uri,omitempty"`
	 Error	Error	`json:"error"`
	 ErrorDescription	string	`json:"error_description,omitempty"`
}
