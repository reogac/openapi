package models

type SbiBindingLevel string

// Define constant values for SbiBindingLevel
const (
	SBIBINDINGLEVEL_NF_INSTANCE_BINDING         SbiBindingLevel = "NF_INSTANCE_BINDING"
	SBIBINDINGLEVEL_NF_SET_BINDING              SbiBindingLevel = "NF_SET_BINDING"
	SBIBINDINGLEVEL_NF_SERVICE_SET_BINDING      SbiBindingLevel = "NF_SERVICE_SET_BINDING"
	SBIBINDINGLEVEL_NF_SERVICE_INSTANCE_BINDING SbiBindingLevel = "NF_SERVICE_INSTANCE_BINDING"
)
