/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HalTemplate struct {
	Properties  []Property `json:"properties,omitempty"`
	Title       string     `json:"title,omitempty"`
	Method      HttpMethod `json:"method"`
	ContentType string     `json:"contentType,omitempty"`
}
