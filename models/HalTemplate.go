package models
type HalTemplate struct {
	 Title	string	`json:"title,omitempty"`
	 Method	HttpMethod	`json:"method"`
	 ContentType	string	`json:"contentType,omitempty"`
	 Properties	[]Property	`json:"properties,omitempty"`
}
