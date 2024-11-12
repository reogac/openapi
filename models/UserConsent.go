package models

type UserConsent string

// Define constant values for UserConsent
const (
	USERCONSENT_CONSENT_NOT_GIVEN UserConsent = "CONSENT_NOT_GIVEN"
	USERCONSENT_CONSENT_GIVEN     UserConsent = "CONSENT_GIVEN"
)
