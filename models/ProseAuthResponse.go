package models

type ProseAuthResponse struct {
	EapPayload string          `json:"eapPayload"`
	Links      map[string]Link `json:"_links"`
}
