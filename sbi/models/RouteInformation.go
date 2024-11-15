/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RouteInformation struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	PortNumber int    `json:"portNumber"`
}
