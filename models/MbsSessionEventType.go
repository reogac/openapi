package models

type MbsSessionEventType string

// Define constant values for MbsSessionEventType
const (
	MBSSESSIONEVENTTYPE_MBS_REL_TMGI_EXPIRY       MbsSessionEventType = "MBS_REL_TMGI_EXPIRY"
	MBSSESSIONEVENTTYPE_BROADCAST_DELIVERY_STATUS MbsSessionEventType = "BROADCAST_DELIVERY_STATUS"
	MBSSESSIONEVENTTYPE_INGRESS_TUNNEL_ADD_CHANGE MbsSessionEventType = "INGRESS_TUNNEL_ADD_CHANGE"
)