package models

type QosFlowTunnel struct {
	TunnelInfo TunnelInfo `json:"tunnelInfo"`
	QfiList    []int      `json:"qfiList"`
}
