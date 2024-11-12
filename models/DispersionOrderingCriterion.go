package models

type DispersionOrderingCriterion string

// Define constant values for DispersionOrderingCriterion
const (
	DISPERSIONORDERINGCRITERION_TIME_SLOT_START    DispersionOrderingCriterion = "TIME_SLOT_START"
	DISPERSIONORDERINGCRITERION_DISPERSION         DispersionOrderingCriterion = "DISPERSION"
	DISPERSIONORDERINGCRITERION_CLASSIFICATION     DispersionOrderingCriterion = "CLASSIFICATION"
	DISPERSIONORDERINGCRITERION_RANKING            DispersionOrderingCriterion = "RANKING"
	DISPERSIONORDERINGCRITERION_PERCENTILE_RANKING DispersionOrderingCriterion = "PERCENTILE_RANKING"
)
