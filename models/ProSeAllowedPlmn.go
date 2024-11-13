package models

type ProSeAllowedPlmn struct {
	ProseDirectAllowed []string `json:"proseDirectAllowed,omitempty"`
	VisitedPlmn        PlmnId   `json:"visitedPlmn"`
}
