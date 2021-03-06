package consumption

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// BillingFrequency enumerates the values for billing frequency.
type BillingFrequency string

const (
	// BillingFrequencyMonth ...
	BillingFrequencyMonth BillingFrequency = "Month"
	// BillingFrequencyQuarter ...
	BillingFrequencyQuarter BillingFrequency = "Quarter"
	// BillingFrequencyYear ...
	BillingFrequencyYear BillingFrequency = "Year"
)

// PossibleBillingFrequencyValues returns an array of possible values for the BillingFrequency const type.
func PossibleBillingFrequencyValues() []BillingFrequency {
	return []BillingFrequency{BillingFrequencyMonth, BillingFrequencyQuarter, BillingFrequencyYear}
}

// Bound enumerates the values for bound.
type Bound string

const (
	// BoundLower ...
	BoundLower Bound = "Lower"
	// BoundUpper ...
	BoundUpper Bound = "Upper"
)

// PossibleBoundValues returns an array of possible values for the Bound const type.
func PossibleBoundValues() []Bound {
	return []Bound{BoundLower, BoundUpper}
}

// ChargeType enumerates the values for charge type.
type ChargeType string

const (
	// ChargeTypeActual ...
	ChargeTypeActual ChargeType = "Actual"
	// ChargeTypeForecast ...
	ChargeTypeForecast ChargeType = "Forecast"
)

// PossibleChargeTypeValues returns an array of possible values for the ChargeType const type.
func PossibleChargeTypeValues() []ChargeType {
	return []ChargeType{ChargeTypeActual, ChargeTypeForecast}
}

// CultureCode enumerates the values for culture code.
type CultureCode string

const (
	// CultureCodeCsCz ...
	CultureCodeCsCz CultureCode = "cs-cz"
	// CultureCodeDaDk ...
	CultureCodeDaDk CultureCode = "da-dk"
	// CultureCodeDeDe ...
	CultureCodeDeDe CultureCode = "de-de"
	// CultureCodeEnGb ...
	CultureCodeEnGb CultureCode = "en-gb"
	// CultureCodeEnUs ...
	CultureCodeEnUs CultureCode = "en-us"
	// CultureCodeEsEs ...
	CultureCodeEsEs CultureCode = "es-es"
	// CultureCodeFrFr ...
	CultureCodeFrFr CultureCode = "fr-fr"
	// CultureCodeHuHu ...
	CultureCodeHuHu CultureCode = "hu-hu"
	// CultureCodeItIt ...
	CultureCodeItIt CultureCode = "it-it"
	// CultureCodeJaJp ...
	CultureCodeJaJp CultureCode = "ja-jp"
	// CultureCodeKoKr ...
	CultureCodeKoKr CultureCode = "ko-kr"
	// CultureCodeNbNo ...
	CultureCodeNbNo CultureCode = "nb-no"
	// CultureCodeNlNl ...
	CultureCodeNlNl CultureCode = "nl-nl"
	// CultureCodePlPl ...
	CultureCodePlPl CultureCode = "pl-pl"
	// CultureCodePtBr ...
	CultureCodePtBr CultureCode = "pt-br"
	// CultureCodePtPt ...
	CultureCodePtPt CultureCode = "pt-pt"
	// CultureCodeRuRu ...
	CultureCodeRuRu CultureCode = "ru-ru"
	// CultureCodeSvSe ...
	CultureCodeSvSe CultureCode = "sv-se"
	// CultureCodeTrTr ...
	CultureCodeTrTr CultureCode = "tr-tr"
	// CultureCodeZhCn ...
	CultureCodeZhCn CultureCode = "zh-cn"
	// CultureCodeZhTw ...
	CultureCodeZhTw CultureCode = "zh-tw"
)

// PossibleCultureCodeValues returns an array of possible values for the CultureCode const type.
func PossibleCultureCodeValues() []CultureCode {
	return []CultureCode{CultureCodeCsCz, CultureCodeDaDk, CultureCodeDeDe, CultureCodeEnGb, CultureCodeEnUs, CultureCodeEsEs, CultureCodeFrFr, CultureCodeHuHu, CultureCodeItIt, CultureCodeJaJp, CultureCodeKoKr, CultureCodeNbNo, CultureCodeNlNl, CultureCodePlPl, CultureCodePtBr, CultureCodePtPt, CultureCodeRuRu, CultureCodeSvSe, CultureCodeTrTr, CultureCodeZhCn, CultureCodeZhTw}
}

// Datagrain enumerates the values for datagrain.
type Datagrain string

const (
	// DatagrainDailyGrain Daily grain of data
	DatagrainDailyGrain Datagrain = "daily"
	// DatagrainMonthlyGrain Monthly grain of data
	DatagrainMonthlyGrain Datagrain = "monthly"
)

// PossibleDatagrainValues returns an array of possible values for the Datagrain const type.
func PossibleDatagrainValues() []Datagrain {
	return []Datagrain{DatagrainDailyGrain, DatagrainMonthlyGrain}
}

// EventType enumerates the values for event type.
type EventType string

const (
	// EventTypeNewCredit ...
	EventTypeNewCredit EventType = "NewCredit"
	// EventTypePendingAdjustments ...
	EventTypePendingAdjustments EventType = "PendingAdjustments"
	// EventTypePendingCharges ...
	EventTypePendingCharges EventType = "PendingCharges"
	// EventTypePendingExpiredCredit ...
	EventTypePendingExpiredCredit EventType = "PendingExpiredCredit"
	// EventTypePendingNewCredit ...
	EventTypePendingNewCredit EventType = "PendingNewCredit"
	// EventTypeSettledCharges ...
	EventTypeSettledCharges EventType = "SettledCharges"
	// EventTypeUnKnown ...
	EventTypeUnKnown EventType = "UnKnown"
)

// PossibleEventTypeValues returns an array of possible values for the EventType const type.
func PossibleEventTypeValues() []EventType {
	return []EventType{EventTypeNewCredit, EventTypePendingAdjustments, EventTypePendingCharges, EventTypePendingExpiredCredit, EventTypePendingNewCredit, EventTypeSettledCharges, EventTypeUnKnown}
}

// Grain enumerates the values for grain.
type Grain string

const (
	// GrainDaily ...
	GrainDaily Grain = "Daily"
	// GrainMonthly ...
	GrainMonthly Grain = "Monthly"
	// GrainYearly ...
	GrainYearly Grain = "Yearly"
)

// PossibleGrainValues returns an array of possible values for the Grain const type.
func PossibleGrainValues() []Grain {
	return []Grain{GrainDaily, GrainMonthly, GrainYearly}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindLegacy ...
	KindLegacy Kind = "legacy"
	// KindModern ...
	KindModern Kind = "modern"
	// KindUsageDetail ...
	KindUsageDetail Kind = "UsageDetail"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindLegacy, KindModern, KindUsageDetail}
}

// KindBasicChargeSummary enumerates the values for kind basic charge summary.
type KindBasicChargeSummary string

const (
	// KindBasicChargeSummaryKindChargeSummary ...
	KindBasicChargeSummaryKindChargeSummary KindBasicChargeSummary = "ChargeSummary"
	// KindBasicChargeSummaryKindLegacy ...
	KindBasicChargeSummaryKindLegacy KindBasicChargeSummary = "legacy"
	// KindBasicChargeSummaryKindModern ...
	KindBasicChargeSummaryKindModern KindBasicChargeSummary = "modern"
)

// PossibleKindBasicChargeSummaryValues returns an array of possible values for the KindBasicChargeSummary const type.
func PossibleKindBasicChargeSummaryValues() []KindBasicChargeSummary {
	return []KindBasicChargeSummary{KindBasicChargeSummaryKindChargeSummary, KindBasicChargeSummaryKindLegacy, KindBasicChargeSummaryKindModern}
}

// KindBasicReservationRecommendation enumerates the values for kind basic reservation recommendation.
type KindBasicReservationRecommendation string

const (
	// KindBasicReservationRecommendationKindLegacy ...
	KindBasicReservationRecommendationKindLegacy KindBasicReservationRecommendation = "legacy"
	// KindBasicReservationRecommendationKindModern ...
	KindBasicReservationRecommendationKindModern KindBasicReservationRecommendation = "modern"
	// KindBasicReservationRecommendationKindReservationRecommendation ...
	KindBasicReservationRecommendationKindReservationRecommendation KindBasicReservationRecommendation = "ReservationRecommendation"
)

// PossibleKindBasicReservationRecommendationValues returns an array of possible values for the KindBasicReservationRecommendation const type.
func PossibleKindBasicReservationRecommendationValues() []KindBasicReservationRecommendation {
	return []KindBasicReservationRecommendation{KindBasicReservationRecommendationKindLegacy, KindBasicReservationRecommendationKindModern, KindBasicReservationRecommendationKindReservationRecommendation}
}

// LookBackPeriod enumerates the values for look back period.
type LookBackPeriod string

const (
	// LookBackPeriodLast07Days Use 7 days of data for recommendations
	LookBackPeriodLast07Days LookBackPeriod = "Last7Days"
	// LookBackPeriodLast30Days Use 30 days of data for recommendations
	LookBackPeriodLast30Days LookBackPeriod = "Last30Days"
	// LookBackPeriodLast60Days Use 60 days of data for recommendations
	LookBackPeriodLast60Days LookBackPeriod = "Last60Days"
)

// PossibleLookBackPeriodValues returns an array of possible values for the LookBackPeriod const type.
func PossibleLookBackPeriodValues() []LookBackPeriod {
	return []LookBackPeriod{LookBackPeriodLast07Days, LookBackPeriodLast30Days, LookBackPeriodLast60Days}
}

// LotSource enumerates the values for lot source.
type LotSource string

const (
	// LotSourcePromotionalCredit ...
	LotSourcePromotionalCredit LotSource = "PromotionalCredit"
	// LotSourcePurchasedCredit ...
	LotSourcePurchasedCredit LotSource = "PurchasedCredit"
)

// PossibleLotSourceValues returns an array of possible values for the LotSource const type.
func PossibleLotSourceValues() []LotSource {
	return []LotSource{LotSourcePromotionalCredit, LotSourcePurchasedCredit}
}

// Metrictype enumerates the values for metrictype.
type Metrictype string

const (
	// MetrictypeActualCostMetricType Actual cost data.
	MetrictypeActualCostMetricType Metrictype = "actualcost"
	// MetrictypeAmortizedCostMetricType Amortized cost data.
	MetrictypeAmortizedCostMetricType Metrictype = "amortizedcost"
	// MetrictypeUsageMetricType Usage data.
	MetrictypeUsageMetricType Metrictype = "usage"
)

// PossibleMetrictypeValues returns an array of possible values for the Metrictype const type.
func PossibleMetrictypeValues() []Metrictype {
	return []Metrictype{MetrictypeActualCostMetricType, MetrictypeAmortizedCostMetricType, MetrictypeUsageMetricType}
}

// OperatorType enumerates the values for operator type.
type OperatorType string

const (
	// OperatorTypeEqualTo ...
	OperatorTypeEqualTo OperatorType = "EqualTo"
	// OperatorTypeGreaterThan ...
	OperatorTypeGreaterThan OperatorType = "GreaterThan"
	// OperatorTypeGreaterThanOrEqualTo ...
	OperatorTypeGreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

// PossibleOperatorTypeValues returns an array of possible values for the OperatorType const type.
func PossibleOperatorTypeValues() []OperatorType {
	return []OperatorType{OperatorTypeEqualTo, OperatorTypeGreaterThan, OperatorTypeGreaterThanOrEqualTo}
}

// Scope12 enumerates the values for scope 12.
type Scope12 string

const (
	// Scope12Shared ...
	Scope12Shared Scope12 = "Shared"
	// Scope12Single ...
	Scope12Single Scope12 = "Single"
)

// PossibleScope12Values returns an array of possible values for the Scope12 const type.
func PossibleScope12Values() []Scope12 {
	return []Scope12{Scope12Shared, Scope12Single}
}

// Scope14 enumerates the values for scope 14.
type Scope14 string

const (
	// Scope14Shared ...
	Scope14Shared Scope14 = "Shared"
	// Scope14Single ...
	Scope14Single Scope14 = "Single"
)

// PossibleScope14Values returns an array of possible values for the Scope14 const type.
func PossibleScope14Values() []Scope14 {
	return []Scope14{Scope14Shared, Scope14Single}
}

// Term enumerates the values for term.
type Term string

const (
	// TermP1Y 1 year reservation term
	TermP1Y Term = "P1Y"
	// TermP3Y 3 year reservation term
	TermP3Y Term = "P3Y"
)

// PossibleTermValues returns an array of possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{TermP1Y, TermP3Y}
}

// ThresholdType enumerates the values for threshold type.
type ThresholdType string

const (
	// ThresholdTypeActual ...
	ThresholdTypeActual ThresholdType = "Actual"
)

// PossibleThresholdTypeValues returns an array of possible values for the ThresholdType const type.
func PossibleThresholdTypeValues() []ThresholdType {
	return []ThresholdType{ThresholdTypeActual}
}

// TimeGrainType enumerates the values for time grain type.
type TimeGrainType string

const (
	// TimeGrainTypeAnnually ...
	TimeGrainTypeAnnually TimeGrainType = "Annually"
	// TimeGrainTypeBillingAnnual ...
	TimeGrainTypeBillingAnnual TimeGrainType = "BillingAnnual"
	// TimeGrainTypeBillingMonth ...
	TimeGrainTypeBillingMonth TimeGrainType = "BillingMonth"
	// TimeGrainTypeBillingQuarter ...
	TimeGrainTypeBillingQuarter TimeGrainType = "BillingQuarter"
	// TimeGrainTypeMonthly ...
	TimeGrainTypeMonthly TimeGrainType = "Monthly"
	// TimeGrainTypeQuarterly ...
	TimeGrainTypeQuarterly TimeGrainType = "Quarterly"
)

// PossibleTimeGrainTypeValues returns an array of possible values for the TimeGrainType const type.
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return []TimeGrainType{TimeGrainTypeAnnually, TimeGrainTypeBillingAnnual, TimeGrainTypeBillingMonth, TimeGrainTypeBillingQuarter, TimeGrainTypeMonthly, TimeGrainTypeQuarterly}
}
