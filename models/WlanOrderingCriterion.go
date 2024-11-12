package models

type WlanOrderingCriterion string

// Define constant values for WlanOrderingCriterion
const (
	WLANORDERINGCRITERION_TIME_SLOT_START WlanOrderingCriterion = "TIME_SLOT_START"
	WLANORDERINGCRITERION_NUMBER_OF_UES   WlanOrderingCriterion = "NUMBER_OF_UES"
	WLANORDERINGCRITERION_RSSI            WlanOrderingCriterion = "RSSI"
	WLANORDERINGCRITERION_RTT             WlanOrderingCriterion = "RTT"
	WLANORDERINGCRITERION_TRAFFIC_INFO    WlanOrderingCriterion = "TRAFFIC_INFO"
)
