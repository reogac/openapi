package models
type HoState string
// Define constant values for HoState
const (
	 HOSTATE_NONE HoState = "NONE"
	 HOSTATE_PREPARING HoState = "PREPARING"
	 HOSTATE_PREPARED HoState = "PREPARED"
	 HOSTATE_COMPLETED HoState = "COMPLETED"
	 HOSTATE_CANCELLED HoState = "CANCELLED"
) 
