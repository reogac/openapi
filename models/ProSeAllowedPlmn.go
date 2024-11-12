package models

type ProSeAllowedPlmn struct {
	VisitedPlmn        PlmnId   `json:"visitedPlmn"`
	ProseDirectAllowed []string `json:"proseDirectAllowed,omitempty"`
}
