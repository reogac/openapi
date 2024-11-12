package models

type HalTemplate struct {
	Properties  []Property `json:"properties,omitempty"`
	Title       string     `json:"title,omitempty"`
	Method      HttpMethod `json:"method"`
	ContentType string     `json:"contentType,omitempty"`
}
