package models

type RedirectResponse struct {
	TargetSepp string `json:"targetSepp,omitempty"`
	Cause      string `json:"cause,omitempty"`
	TargetScp  string `json:"targetScp,omitempty"`
}
