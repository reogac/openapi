package models

type AuthType string

// Define constant values for AuthType
const (
	AUTHTYPE_5G_AKA        AuthType = "5G_AKA"
	AUTHTYPE_EAP_AKA_PRIME AuthType = "EAP_AKA_PRIME"
	AUTHTYPE_EAP_TLS       AuthType = "EAP_TLS"
	AUTHTYPE_EAP_TTLS      AuthType = "EAP_TTLS"
)
