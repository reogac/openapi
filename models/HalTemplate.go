package models

type HalTemplate struct {
	Method      HttpMethod `json:"method"`
	ContentType string     `json:"contentType,omitempty"`
	Properties  []Property `json:"properties,omitempty"`
	Title       string     `json:"title,omitempty"`
}
