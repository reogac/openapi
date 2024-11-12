package models

type RedirectResponse struct {
	Cause      string `json:"cause,omitempty"`
	TargetScp  string `json:"targetScp,omitempty"`
	TargetSepp string `json:"targetSepp,omitempty"`
}
