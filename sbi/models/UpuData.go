/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:08 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpuData struct {
	RoutingId        string   `json:"routingId,omitempty"`
	Drei             *bool    `json:"drei,omitempty"`
	Aol              *bool    `json:"aol,omitempty"`
	SecPacket        string   `json:"secPacket,omitempty"`
	DefaultConfNssai []Snssai `json:"defaultConfNssai,omitempty"`
}