package models

type SmsSupport string

// Define constant values for SmsSupport
const (
	SMSSUPPORT_3GPP     SmsSupport = "3GPP"
	SMSSUPPORT_NON_3GPP SmsSupport = "NON_3GPP"
	SMSSUPPORT_BOTH     SmsSupport = "BOTH"
	SMSSUPPORT_NONE     SmsSupport = "NONE"
)
