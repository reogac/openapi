package models

type RedirectAddressType string

// Define constant values for RedirectAddressType
const (
	REDIRECTADDRESSTYPE_IPV4_ADDR RedirectAddressType = "IPV4_ADDR"
	REDIRECTADDRESSTYPE_IPV6_ADDR RedirectAddressType = "IPV6_ADDR"
	REDIRECTADDRESSTYPE_URL       RedirectAddressType = "URL"
	REDIRECTADDRESSTYPE_SIP_URI   RedirectAddressType = "SIP_URI"
)