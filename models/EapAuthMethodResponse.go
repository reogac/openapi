package models

type EapAuthMethodResponse struct {
	EapPayload string          `json:"eapPayload"`
	Links      map[string]Link `json:"_links"`
}
