package models

type ComplexQuery struct {
	Dnf *Dnf `json:"Dnf,omitempty"`
	Cnf *Cnf `json:"Cnf,omitempty"`
}
