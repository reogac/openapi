package models
type UeContextInSmsfData struct {
	 SmsfInfoNon3GppAccess	*SmsfInfo	`json:"smsfInfoNon3GppAccess,omitempty"`
	 SmsfInfo3GppAccess	*SmsfInfo	`json:"smsfInfo3GppAccess,omitempty"`
}
