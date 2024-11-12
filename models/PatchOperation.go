package models

type PatchOperation string

// Define constant values for PatchOperation
const (
	PATCHOPERATION_ADD     PatchOperation = "add"
	PATCHOPERATION_COPY    PatchOperation = "copy"
	PATCHOPERATION_MOVE    PatchOperation = "move"
	PATCHOPERATION_REMOVE  PatchOperation = "remove"
	PATCHOPERATION_REPLACE PatchOperation = "replace"
	PATCHOPERATION_TEST    PatchOperation = "test"
)
