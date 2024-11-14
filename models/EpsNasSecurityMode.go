package models
type EpsNasSecurityMode struct {
	 IntegrityAlgorithm	EpsNasIntegrityAlgorithm	`json:"integrityAlgorithm"`
	 CipheringAlgorithm	EpsNasCipheringAlgorithm	`json:"cipheringAlgorithm"`
}
