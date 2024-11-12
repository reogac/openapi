package models

type RedirectResponse struct {
	TargetScp  string `json:"targetScp,omitempty"`
	TargetSepp string `json:"targetSepp,omitempty"`
	Cause      string `json:"cause,omitempty"`
}
