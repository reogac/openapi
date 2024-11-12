package models

type Ssm struct {
	DestIpAddr   IpAddr `json:"destIpAddr"`
	SourceIpAddr IpAddr `json:"sourceIpAddr"`
}
