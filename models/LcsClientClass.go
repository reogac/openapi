package models

type LcsClientClass string

// Define constant values for LcsClientClass
const (
	LCSCLIENTCLASS_BROADCAST_SERVICE          LcsClientClass = "BROADCAST_SERVICE"
	LCSCLIENTCLASS_OM_IN_HPLMN                LcsClientClass = "OM_IN_HPLMN"
	LCSCLIENTCLASS_OM_IN_VPLMN                LcsClientClass = "OM_IN_VPLMN"
	LCSCLIENTCLASS_ANONYMOUS_LOCATION_SERVICE LcsClientClass = "ANONYMOUS_LOCATION_SERVICE"
	LCSCLIENTCLASS_SPECIFIC_SERVICE           LcsClientClass = "SPECIFIC_SERVICE"
)