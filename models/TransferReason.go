package models

type TransferReason string

// Define constant values for TransferReason
const (
	TRANSFERREASON_INIT_REG              TransferReason = "INIT_REG"
	TRANSFERREASON_MOBI_REG              TransferReason = "MOBI_REG"
	TRANSFERREASON_MOBI_REG_UE_VALIDATED TransferReason = "MOBI_REG_UE_VALIDATED"
)