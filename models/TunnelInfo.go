/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TunnelInfo struct {
	Ipv4Addr string     `json:"ipv4Addr,omitempty"`
	Ipv6Addr string     `json:"ipv6Addr,omitempty"`
	GtpTeid  string     `json:"gtpTeid"`
	AnType   AccessType `json:"anType,omitempty"`
}
