package models

type RequestTrigger string

// Define constant values for RequestTrigger
const (
	REQUESTTRIGGER_LOC_CH            RequestTrigger = "LOC_CH"
	REQUESTTRIGGER_PRA_CH            RequestTrigger = "PRA_CH"
	REQUESTTRIGGER_UE_POLICY         RequestTrigger = "UE_POLICY"
	REQUESTTRIGGER_PLMN_CH           RequestTrigger = "PLMN_CH"
	REQUESTTRIGGER_CON_STATE_CH      RequestTrigger = "CON_STATE_CH"
	REQUESTTRIGGER_GROUP_ID_LIST_CHG RequestTrigger = "GROUP_ID_LIST_CHG"
	REQUESTTRIGGER_UE_CAP_CH         RequestTrigger = "UE_CAP_CH"
)