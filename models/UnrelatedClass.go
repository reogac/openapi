package models

type UnrelatedClass struct {
	ServiceTypeUnrelatedClasses []ServiceTypeUnrelatedClass `json:"serviceTypeUnrelatedClasses,omitempty"`
	DefaultUnrelatedClass       DefaultUnrelatedClass       `json:"defaultUnrelatedClass"`
}
