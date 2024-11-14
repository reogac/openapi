package models
type ContextInfo struct {
	 RequestHeaders	[]string	`json:"requestHeaders,omitempty"`
	 OrigHeaders	[]string	`json:"origHeaders,omitempty"`
}
