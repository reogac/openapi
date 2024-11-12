package models

type UeContextInSmsfData struct {
	SmsfInfo3GppAccess    *SmsfInfo `json:"smsfInfo3GppAccess,omitempty"`
	SmsfInfoNon3GppAccess *SmsfInfo `json:"smsfInfoNon3GppAccess,omitempty"`
}
