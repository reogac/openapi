package models
type EapAuthMethodResponse struct {
	 Links	map[string]Link	`json:"_links"`
	 EapPayload	string	`json:"eapPayload"`
}
