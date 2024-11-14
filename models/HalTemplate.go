package models
type HalTemplate struct {
	 ContentType	string	`json:"contentType,omitempty"`
	 Properties	[]Property	`json:"properties,omitempty"`
	 Title	string	`json:"title,omitempty"`
	 Method	HttpMethod	`json:"method"`
}
