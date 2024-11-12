package models

type ComplexQuery struct {
	Cnf *Cnf `json:"Cnf,omitempty"`
	Dnf *Dnf `json:"Dnf,omitempty"`
}
