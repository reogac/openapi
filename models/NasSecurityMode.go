package models

type NasSecurityMode struct {
	CipheringAlgorithm string `json:"cipheringAlgorithm"`
	IntegrityAlgorithm string `json:"integrityAlgorithm"`
}
