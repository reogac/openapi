package models

type EpsNasSecurityMode struct {
	IntegrityAlgorithm string `json:"integrityAlgorithm"`
	CipheringAlgorithm string `json:"cipheringAlgorithm"`
}
