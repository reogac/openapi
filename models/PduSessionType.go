package models

type PduSessionType string

// Define constant values for PduSessionType
const (
	PDUSESSIONTYPE_IPV4         PduSessionType = "IPV4"
	PDUSESSIONTYPE_IPV6         PduSessionType = "IPV6"
	PDUSESSIONTYPE_IPV4V6       PduSessionType = "IPV4V6"
	PDUSESSIONTYPE_UNSTRUCTURED PduSessionType = "UNSTRUCTURED"
	PDUSESSIONTYPE_ETHERNET     PduSessionType = "ETHERNET"
)