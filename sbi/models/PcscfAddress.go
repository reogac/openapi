/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PcscfAddress struct {
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`
	Fqdn      string   `json:"fqdn,omitempty"`
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`
}
