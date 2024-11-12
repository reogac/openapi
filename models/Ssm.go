package models

type Ssm struct {
	SourceIpAddr IpAddr `json:"sourceIpAddr"`
	DestIpAddr   IpAddr `json:"destIpAddr"`
}
