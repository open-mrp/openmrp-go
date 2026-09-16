// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// The planning assumptions production schedules are solved against, and the
// per-resource overrides that mark which machines constrain the plan.
//
// OperationProductionScheduleSettingService contains methods and other services
// that help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleSettingService] method instead.
type OperationProductionScheduleSettingService struct {
	options []option.RequestOption
	// The planning assumptions production schedules are solved against, and the
	// per-resource overrides that mark which machines constrain the plan.
	Resources OperationProductionScheduleSettingResourceService
	// The planning assumptions production schedules are solved against, and the
	// per-resource overrides that mark which machines constrain the plan.
	Items OperationProductionScheduleSettingItemService
}

// NewOperationProductionScheduleSettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewOperationProductionScheduleSettingService(opts ...option.RequestOption) (r OperationProductionScheduleSettingService) {
	r = OperationProductionScheduleSettingService{}
	r.options = opts
	r.Resources = NewOperationProductionScheduleSettingResourceService(opts...)
	r.Items = NewOperationProductionScheduleSettingItemService(opts...)
	return
}

// Replaces the planning assumptions production schedules are solved against.
//
// Settings are replaced wholesale rather than patched, because they are read as
// one coherent set: a horizon that no longer matches the frozen window, or a
// capacity headroom that no longer matches the shift pattern, would produce a plan
// nobody intended. Send the full set on every call — a value the request leaves
// out is never carried over from what was stored.
//
// The set is validated together, so a frozen window longer than the horizon, a
// minimum changeover above the maximum, or an active cadence with no valid
// schedule expression is rejected as a whole.
//
// Existing schedule versions are unaffected — each one records the assumptions it
// was solved under, so changing settings changes future plans only.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleSettingService) Update(ctx context.Context, body OperationProductionScheduleSettingUpdateParams, opts ...option.RequestOption) (res *ProductionScheduleSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the planning assumptions production schedules are solved against.
//
// The whole set is always returned. An account that has never saved settings reads
// back the values the solver would apply anyway, so a caller never has to know
// which assumptions are in play; `settings_status` says whether the values were
// saved on the account or are those defaults.
//
// Per-machine, per-department and per-step overrides of these assumptions are read
// separately.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleSettingService) List(ctx context.Context, opts ...option.RequestOption) (res *ProductionScheduleSettings, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedule-settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The planning assumptions a production schedule is solved against.
//
// The whole set is always returned. An account that has never saved settings reads
// back the values the solver would apply anyway, so a caller never has to know
// which assumptions are in play; `settings_status` says whether the values were
// saved on the account or are those defaults.
type ProductionScheduleSettings struct {
	// Whether a version produced by the cadence is published automatically.
	//
	// While active, a cadence run publishes as soon as it solves, committing its
	// frozen weeks without anyone reviewing the plan. Otherwise the run leaves a draft
	// for a planner to publish by hand. Versions generated on request are never
	// published automatically.
	//
	// Any of "active", "inactive".
	AutoPublishStatus ProductionScheduleSettingsAutoPublishStatus `json:"auto_publish_status" api:"required"`
	// Whether schedules are generated automatically on a recurring cadence.
	//
	// While active, each due tick queues a new schedule version; a generation cron
	// expression is required for the cadence to be saved.
	//
	// Any of "active", "inactive".
	CadenceStatus ProductionScheduleSettingsCadenceStatus `json:"cadence_status" api:"required"`
	// Share of machine time a plan may fill.
	//
	// Shifts, hours and work days give a machine's raw weekly hours; this trims them
	// to what may actually be planned. The remainder absorbs changeovers, which are
	// not scheduled as explicit blocks, so a value of 1 produces a plan that leaves no
	// time to set anything up.
	CapacityHeadroomPct float64 `json:"capacity_headroom_pct" api:"required"`
	// Typical changeover duration.
	//
	// Changeover time is modeled as rising with the number of new inputs a product
	// introduces, between the minimum and maximum below. The slope is calibrated from
	// production history so the model reproduces this average across the transitions
	// actually observed, which is why the value belongs at the changeover time the
	// floor typically reports rather than at a worst case.
	ChangeoverAvgMinutes float64 `json:"changeover_avg_minutes" api:"required"`
	// Hourly labor rate charged to a changeover.
	//
	// This is a dedicated technician rate rather than an allocated production rate,
	// because one person works a single machine through a changeover. Together with
	// the typical changeover duration it prices the setup cost that decides economic
	// campaign sizes. The constraint department's own labor rate takes precedence when
	// it has one, leaving this as the fallback.
	ChangeoverLaborRate float64 `json:"changeover_labor_rate" api:"required"`
	// Longest plausible changeover, and the ceiling of the changeover model.
	ChangeoverMaxMinutes float64 `json:"changeover_max_minutes" api:"required"`
	// Shortest plausible changeover, and the floor of the changeover model.
	ChangeoverMinMinutes float64 `json:"changeover_min_minutes" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ConstraintDepartment Entity `json:"constraint_department" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Weeks of lead time to assume at the constraint for an item with no measured
	// history.
	//
	// An item's own lead time, measured from production history, is used instead
	// whenever one can be observed.
	DefaultConstraintLeadTimeWeeks float64 `json:"default_constraint_lead_time_weeks" api:"required"`
	// Calendar days between an order being issued and it being due to ship.
	//
	// The last resort in the ship-by chain: a lead time set on the customer, on its
	// parent account, or on the customer's account group takes precedence. Zero means
	// same-day shipping.
	DefaultCustomerLeadTimeDays int64 `json:"default_customer_lead_time_days" api:"required"`
	// How a SKU is produced when neither it nor its product line says.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Any of "make_to_stock", "make_to_order".
	DefaultFulfillmentPolicy ProductionScheduleSettingsDefaultFulfillmentPolicy `json:"default_fulfillment_policy" api:"required"`
	// Units in a default production lot.
	//
	// The last resort in the lot-size chain: a lot set on the item, on its product
	// line, or on the finished goods an intermediate item becomes all take precedence.
	DefaultLotUnits float64 `json:"default_lot_units" api:"required"`
	// How the demand a plan is solved against is derived from history.
	//
	//   - `trailing_12`: the last twelve complete months of orders, spread evenly across
	//     the coming year.
	//   - `seasonal_ema`: a seasonally adjusted, exponentially smoothed projection that
	//     weights recent months more heavily. Falls back to the trailing baseline for an
	//     item with no history.
	//
	// Demand overrides are applied on top of whichever baseline is chosen.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis ProductionScheduleSettingsDemandBasis `json:"demand_basis" api:"required"`
	// Months of production history the solver measures run rates, changeover behavior
	// and lead times from.
	DemandWindowMonths int64 `json:"demand_window_months" api:"required"`
	// Weeks between coming off the constraint and being sellable.
	//
	// Added to the constraint's own lead time when reorder points are set, so a plan
	// replenishes early enough for a decision made today to become sellable stock.
	FinishLeadTimeWeeks float64 `json:"finish_lead_time_weeks" api:"required"`
	// Months of order history the demand baseline is drawn from.
	ForecastHistoryMonths int64 `json:"forecast_history_months" api:"required"`
	// Months the forecast projects forward.
	//
	// Only applies to the `seasonal_ema` basis. A projection of anything other than
	// twelve months is scaled to an annual rate, so the plan always reasons about a
	// year of demand.
	ForecastMonths int64 `json:"forecast_months" api:"required"`
	// Z-score used for the confidence interval around the seasonal demand forecast.
	//
	// The plan is solved against the central forecast, so this widens or narrows that
	// interval without changing what gets scheduled.
	ForecastZ float64 `json:"forecast_z" api:"required"`
	// How many leading weeks of the horizon become a commitment when a version is
	// published.
	//
	// Nothing is frozen while a version is still a draft. Once published, changing a
	// campaign inside the frozen window requires a reason and is recorded against the
	// plan. Cannot be longer than the planning horizon.
	FrozenWeeks int64 `json:"frozen_weeks" api:"required"`
	// Standard cron expression driving the generation cadence.
	GenerationCron string `json:"generation_cron" api:"required"`
	// Timezone the cadence is interpreted in.
	//
	// Decides when "every Wednesday at 6am" actually happens. A timezone the platform
	// does not recognize falls back to UTC.
	GenerationTimezone string `json:"generation_timezone" api:"required"`
	// Annual cost of holding stock, as a share of item value.
	//
	// Weighed against the cost of a changeover when campaigns are sized: a higher rate
	// favors shorter, more frequent runs.
	HoldingRatePct float64 `json:"holding_rate_pct" api:"required"`
	// Hours in a shift.
	HoursPerShift float64 `json:"hours_per_shift" api:"required"`
	// When the cadence last fired.
	//
	// Stamped when a run is queued rather than when the plan finishes solving, and the
	// next due time is measured from it.
	LastGeneratedAt time.Time `json:"last_generated_at" api:"required" format:"date-time"`
	// How many steps down the production flow a constraint item is traced to the
	// finished goods it becomes.
	//
	// Demand, stock and lot conventions are pooled onto the constraint item from every
	// finished good the trace reaches, so anything further down the flow than this
	// contributes nothing to the plan. The limit is also what stops a routing that
	// loops back on itself from being traced forever.
	MaxFlowDepth int64 `json:"max_flow_depth" api:"required"`
	// Ceiling on how far ahead any item is built.
	//
	// An item is only rebuilt once its projected stock falls below the lower of its
	// reorder point and this many weeks of demand, so a slow mover whose statistical
	// reorder point covers months of demand is not topped up ahead of items that are
	// actually short.
	MaxWeeksSupply float64 `json:"max_weeks_supply" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_settings".
	Object ProductionScheduleSettingsObject `json:"object" api:"required"`
	// How many weeks a generated plan covers.
	PlanningHorizonWeeks int64  `json:"planning_horizon_weeks" api:"required"`
	ReceiveCalendarID    string `json:"receive_calendar_id" api:"required"`
	// Z-score behind the safety stock targets.
	//
	// A higher value buys more cover against demand variability at both the constraint
	// and the finished goods stage, at the cost of carrying more stock.
	ServiceLevelZ float64 `json:"service_level_z" api:"required"`
	// Whether the values returned were saved on the account or are the defaults
	// applied when nothing has been saved.
	//
	// Any of "stored", "default".
	SettingsStatus ProductionScheduleSettingsSettingsStatus `json:"settings_status" api:"required"`
	// Shifts worked per day.
	ShiftsPerDay int64 `json:"shifts_per_day" api:"required"`
	// The account-wide operating calendars: the days the plant tenders freight, and
	// the days a customer's dock accepts it.
	//
	// Behind the per-address and per-customer links and ahead of a plain
	// Monday-to-Friday week. Null on both means every ship-by date is resolved against
	// weekdays alone.
	ShipCalendarID string `json:"ship_calendar_id" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Day a planning week starts, where 0 is Sunday.
	WeekStartDay int64 `json:"week_start_day" api:"required"`
	// Weeks worked per year.
	WeeksPerYear int64 `json:"weeks_per_year" api:"required"`
	// Days worked per week.
	WorkDaysPerWeek int64 `json:"work_days_per_week" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoPublishStatus              respjson.Field
		CadenceStatus                  respjson.Field
		CapacityHeadroomPct            respjson.Field
		ChangeoverAvgMinutes           respjson.Field
		ChangeoverLaborRate            respjson.Field
		ChangeoverMaxMinutes           respjson.Field
		ChangeoverMinMinutes           respjson.Field
		ConstraintDepartment           respjson.Field
		CreatedAt                      respjson.Field
		DefaultConstraintLeadTimeWeeks respjson.Field
		DefaultCustomerLeadTimeDays    respjson.Field
		DefaultFulfillmentPolicy       respjson.Field
		DefaultLotUnits                respjson.Field
		DemandBasis                    respjson.Field
		DemandWindowMonths             respjson.Field
		FinishLeadTimeWeeks            respjson.Field
		ForecastHistoryMonths          respjson.Field
		ForecastMonths                 respjson.Field
		ForecastZ                      respjson.Field
		FrozenWeeks                    respjson.Field
		GenerationCron                 respjson.Field
		GenerationTimezone             respjson.Field
		HoldingRatePct                 respjson.Field
		HoursPerShift                  respjson.Field
		LastGeneratedAt                respjson.Field
		MaxFlowDepth                   respjson.Field
		MaxWeeksSupply                 respjson.Field
		Object                         respjson.Field
		PlanningHorizonWeeks           respjson.Field
		ReceiveCalendarID              respjson.Field
		ServiceLevelZ                  respjson.Field
		SettingsStatus                 respjson.Field
		ShiftsPerDay                   respjson.Field
		ShipCalendarID                 respjson.Field
		UpdatedAt                      respjson.Field
		WeekStartDay                   respjson.Field
		WeeksPerYear                   respjson.Field
		WorkDaysPerWeek                respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleSettings) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a version produced by the cadence is published automatically.
//
// While active, a cadence run publishes as soon as it solves, committing its
// frozen weeks without anyone reviewing the plan. Otherwise the run leaves a draft
// for a planner to publish by hand. Versions generated on request are never
// published automatically.
type ProductionScheduleSettingsAutoPublishStatus string

const (
	ProductionScheduleSettingsAutoPublishStatusActive   ProductionScheduleSettingsAutoPublishStatus = "active"
	ProductionScheduleSettingsAutoPublishStatusInactive ProductionScheduleSettingsAutoPublishStatus = "inactive"
)

// Whether schedules are generated automatically on a recurring cadence.
//
// While active, each due tick queues a new schedule version; a generation cron
// expression is required for the cadence to be saved.
type ProductionScheduleSettingsCadenceStatus string

const (
	ProductionScheduleSettingsCadenceStatusActive   ProductionScheduleSettingsCadenceStatus = "active"
	ProductionScheduleSettingsCadenceStatusInactive ProductionScheduleSettingsCadenceStatus = "inactive"
)

// How a SKU is produced when neither it nor its product line says.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
type ProductionScheduleSettingsDefaultFulfillmentPolicy string

const (
	ProductionScheduleSettingsDefaultFulfillmentPolicyMakeToStock ProductionScheduleSettingsDefaultFulfillmentPolicy = "make_to_stock"
	ProductionScheduleSettingsDefaultFulfillmentPolicyMakeToOrder ProductionScheduleSettingsDefaultFulfillmentPolicy = "make_to_order"
)

// How the demand a plan is solved against is derived from history.
//
//   - `trailing_12`: the last twelve complete months of orders, spread evenly across
//     the coming year.
//   - `seasonal_ema`: a seasonally adjusted, exponentially smoothed projection that
//     weights recent months more heavily. Falls back to the trailing baseline for an
//     item with no history.
//
// Demand overrides are applied on top of whichever baseline is chosen.
type ProductionScheduleSettingsDemandBasis string

const (
	ProductionScheduleSettingsDemandBasisTrailing12  ProductionScheduleSettingsDemandBasis = "trailing_12"
	ProductionScheduleSettingsDemandBasisSeasonalEma ProductionScheduleSettingsDemandBasis = "seasonal_ema"
)

// Resource type identifier.
type ProductionScheduleSettingsObject string

const (
	ProductionScheduleSettingsObjectProductionScheduleSettings ProductionScheduleSettingsObject = "production_schedule_settings"
)

// Whether the values returned were saved on the account or are the defaults
// applied when nothing has been saved.
type ProductionScheduleSettingsSettingsStatus string

const (
	ProductionScheduleSettingsSettingsStatusStored  ProductionScheduleSettingsSettingsStatus = "stored"
	ProductionScheduleSettingsSettingsStatusDefault ProductionScheduleSettingsSettingsStatus = "default"
)

// Request to replace the account's planning assumptions.
//
// The properties AutoPublishStatus, CadenceStatus, CapacityHeadroomPct,
// ChangeoverAvgMinutes, ChangeoverLaborRate, ChangeoverMaxMinutes,
// ChangeoverMinMinutes, DefaultConstraintLeadTimeWeeks,
// DefaultCustomerLeadTimeDays, DefaultFulfillmentPolicy, DefaultLotUnits,
// DemandBasis, DemandWindowMonths, FinishLeadTimeWeeks, ForecastHistoryMonths,
// ForecastMonths, ForecastZ, FrozenWeeks, GenerationTimezone, HoldingRatePct,
// HoursPerShift, MaxFlowDepth, MaxWeeksSupply, PlanningHorizonWeeks,
// ServiceLevelZ, ShiftsPerDay, WeekStartDay, WeeksPerYear, WorkDaysPerWeek are
// required.
type UpdateProductionScheduleSettingsRequestParam struct {
	// Whether a version produced by the cadence is published automatically.
	//
	// While active, a cadence run publishes as soon as it solves, committing its
	// frozen weeks without anyone reviewing the plan. Otherwise the run leaves a draft
	// for a planner to publish by hand. Versions generated on request are never
	// published automatically.
	//
	// Any of "active", "inactive".
	AutoPublishStatus UpdateProductionScheduleSettingsRequestAutoPublishStatus `json:"auto_publish_status,omitzero" api:"required"`
	// Whether schedules are generated automatically on a recurring cadence.
	//
	// While active, each due tick queues a new schedule version.
	//
	// Any of "active", "inactive".
	CadenceStatus UpdateProductionScheduleSettingsRequestCadenceStatus `json:"cadence_status,omitzero" api:"required"`
	// Share of machine time a plan may fill.
	//
	// Shifts, hours and work days give a machine's raw weekly hours; this trims them
	// to what may actually be planned. The remainder absorbs changeovers, which are
	// not scheduled as explicit blocks, so a value of 1 produces a plan that leaves no
	// time to set anything up.
	CapacityHeadroomPct float64 `json:"capacity_headroom_pct" api:"required"`
	// Typical changeover duration.
	//
	// Changeover time is modeled as rising with the number of new inputs a product
	// introduces, between the minimum and maximum below. The slope is calibrated from
	// production history so the model reproduces this average across the transitions
	// actually observed; set it to the changeover time the floor typically reports
	// rather than to a worst case.
	ChangeoverAvgMinutes float64 `json:"changeover_avg_minutes" api:"required"`
	// Hourly labor rate charged to a changeover.
	//
	// This should be a dedicated technician rate rather than an allocated production
	// rate, because one person works a single machine through a changeover. The
	// constraint department's own labor rate takes precedence when it has one, leaving
	// this as the fallback.
	ChangeoverLaborRate float64 `json:"changeover_labor_rate" api:"required"`
	// Longest plausible changeover, and the ceiling of the changeover model.
	ChangeoverMaxMinutes float64 `json:"changeover_max_minutes" api:"required"`
	// Shortest plausible changeover, and the floor of the changeover model.
	//
	// Cannot exceed the maximum.
	ChangeoverMinMinutes float64 `json:"changeover_min_minutes" api:"required"`
	// Weeks of lead time to assume at the constraint for an item with no measured
	// history.
	//
	// An item's own lead time, measured from production history, is used instead
	// whenever one can be observed.
	DefaultConstraintLeadTimeWeeks float64 `json:"default_constraint_lead_time_weeks" api:"required"`
	// Calendar days between an order being issued and it being due to ship.
	//
	// The last resort in the ship-by chain: a lead time set on the customer, on its
	// parent account, or on the customer's account group takes precedence. Zero
	// commits the account to same-day shipping on every order that falls through to
	// it, so this update replaces the whole settings object and omitting the field is
	// not the same as leaving it alone.
	DefaultCustomerLeadTimeDays int64 `json:"default_customer_lead_time_days" api:"required"`
	// How a SKU is produced when neither it nor its product line says.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Any of "make_to_stock", "make_to_order".
	DefaultFulfillmentPolicy UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicy `json:"default_fulfillment_policy,omitzero" api:"required"`
	// Units in a default production lot.
	//
	// The last resort in the lot-size chain: a lot set on the item, on its product
	// line, or on the finished goods an intermediate item becomes all take precedence.
	DefaultLotUnits float64 `json:"default_lot_units" api:"required"`
	// How the demand a plan is solved against is derived from history.
	//
	//   - `trailing_12`: the last twelve complete months of orders, spread evenly across
	//     the coming year.
	//   - `seasonal_ema`: a seasonally adjusted, exponentially smoothed projection that
	//     weights recent months more heavily. Falls back to the trailing baseline for an
	//     item with no history.
	//
	// Demand overrides are applied on top of whichever baseline is chosen.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis UpdateProductionScheduleSettingsRequestDemandBasis `json:"demand_basis,omitzero" api:"required"`
	// Months of production history the solver measures run rates, changeover behavior
	// and lead times from.
	DemandWindowMonths int64 `json:"demand_window_months" api:"required"`
	// Weeks between coming off the constraint and being sellable.
	FinishLeadTimeWeeks float64 `json:"finish_lead_time_weeks" api:"required"`
	// Months of order history the demand baseline is drawn from.
	ForecastHistoryMonths int64 `json:"forecast_history_months" api:"required"`
	// Months the forecast projects forward.
	//
	// Only applies to the `seasonal_ema` basis. A projection of anything other than
	// twelve months is scaled to an annual rate, so the plan always reasons about a
	// year of demand.
	ForecastMonths int64 `json:"forecast_months" api:"required"`
	// Z-score used for the confidence interval around the seasonal demand forecast.
	//
	// The plan is solved against the central forecast, so this widens or narrows that
	// interval without changing what gets scheduled.
	ForecastZ float64 `json:"forecast_z" api:"required"`
	// How many leading weeks of the horizon become a commitment when a version is
	// published.
	//
	// Cannot be longer than the planning horizon. Once a version is published,
	// changing a campaign inside the frozen window requires a reason and is recorded
	// against the plan.
	FrozenWeeks int64 `json:"frozen_weeks" api:"required"`
	// Timezone the cadence is interpreted in.
	//
	// Decides when "every Wednesday at 6am" actually happens. A timezone the platform
	// does not recognize falls back to UTC.
	GenerationTimezone string `json:"generation_timezone" api:"required"`
	// Annual cost of holding stock, as a share of item value.
	//
	// Weighed against the cost of a changeover when campaigns are sized: a higher rate
	// favors shorter, more frequent runs.
	HoldingRatePct float64 `json:"holding_rate_pct" api:"required"`
	// Hours in a shift.
	HoursPerShift float64 `json:"hours_per_shift" api:"required"`
	// How many steps down the production flow a constraint item is traced to the
	// finished goods it becomes.
	//
	// Demand, stock and lot conventions are pooled onto the constraint item from every
	// finished good the trace reaches, so anything further down the flow than this
	// contributes nothing to the plan. The limit is also what stops a routing that
	// loops back on itself from being traced forever.
	MaxFlowDepth int64 `json:"max_flow_depth" api:"required"`
	// Ceiling on how far ahead any item is built.
	//
	// An item is only rebuilt once its projected stock falls below the lower of its
	// reorder point and this many weeks of demand, so a slow mover whose statistical
	// reorder point covers months of demand is not topped up ahead of items that are
	// actually short.
	MaxWeeksSupply float64 `json:"max_weeks_supply" api:"required"`
	// How many weeks a generated plan covers.
	PlanningHorizonWeeks int64 `json:"planning_horizon_weeks" api:"required"`
	// Z-score for service level safety stock targets.
	ServiceLevelZ float64 `json:"service_level_z" api:"required"`
	// Shifts worked per day.
	ShiftsPerDay int64 `json:"shifts_per_day" api:"required"`
	// Day a planning week starts, where 0 is Sunday.
	WeekStartDay int64 `json:"week_start_day" api:"required"`
	// Weeks worked per year.
	WeeksPerYear int64 `json:"weeks_per_year" api:"required"`
	// Days worked per week.
	WorkDaysPerWeek int64 `json:"work_days_per_week" api:"required"`
	// ID of the department that sets the pace of the factory, and the one campaigns
	// are planned onto.
	//
	// Every machine in the department is planned, and the work of downstream
	// departments is derived from what those machines are scheduled to run. Sending
	// null, or leaving the field out of a request that otherwise replaces the
	// settings, both leave the account with no constraint department — and generation
	// is refused until one is chosen again.
	ConstraintDepartmentID param.Opt[string] `json:"constraint_department_id,omitzero"`
	// Standard cron expression driving the generation cadence.
	//
	// Must be present and parse as a standard cron expression whenever the cadence is
	// active, otherwise the whole update is rejected.
	GenerationCron    param.Opt[string] `json:"generation_cron,omitzero"`
	ReceiveCalendarID param.Opt[string] `json:"receive_calendar_id,omitzero"`
	// The operating calendar naming the days this account's plant tenders freight, and
	// the one naming the days a customer's dock accepts it.
	//
	// These are the account-wide fallbacks: an address or a customer with its own
	// calendar overrides them, and an account with neither set falls back to a
	// Monday-to-Friday week with no closures.
	ShipCalendarID param.Opt[string] `json:"ship_calendar_id,omitzero"`
	paramObj
}

func (r UpdateProductionScheduleSettingsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductionScheduleSettingsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductionScheduleSettingsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether a version produced by the cadence is published automatically.
//
// While active, a cadence run publishes as soon as it solves, committing its
// frozen weeks without anyone reviewing the plan. Otherwise the run leaves a draft
// for a planner to publish by hand. Versions generated on request are never
// published automatically.
type UpdateProductionScheduleSettingsRequestAutoPublishStatus string

const (
	UpdateProductionScheduleSettingsRequestAutoPublishStatusActive   UpdateProductionScheduleSettingsRequestAutoPublishStatus = "active"
	UpdateProductionScheduleSettingsRequestAutoPublishStatusInactive UpdateProductionScheduleSettingsRequestAutoPublishStatus = "inactive"
)

// Whether schedules are generated automatically on a recurring cadence.
//
// While active, each due tick queues a new schedule version.
type UpdateProductionScheduleSettingsRequestCadenceStatus string

const (
	UpdateProductionScheduleSettingsRequestCadenceStatusActive   UpdateProductionScheduleSettingsRequestCadenceStatus = "active"
	UpdateProductionScheduleSettingsRequestCadenceStatusInactive UpdateProductionScheduleSettingsRequestCadenceStatus = "inactive"
)

// How a SKU is produced when neither it nor its product line says.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
type UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicy string

const (
	UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicyMakeToStock UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicy = "make_to_stock"
	UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicyMakeToOrder UpdateProductionScheduleSettingsRequestDefaultFulfillmentPolicy = "make_to_order"
)

// How the demand a plan is solved against is derived from history.
//
//   - `trailing_12`: the last twelve complete months of orders, spread evenly across
//     the coming year.
//   - `seasonal_ema`: a seasonally adjusted, exponentially smoothed projection that
//     weights recent months more heavily. Falls back to the trailing baseline for an
//     item with no history.
//
// Demand overrides are applied on top of whichever baseline is chosen.
type UpdateProductionScheduleSettingsRequestDemandBasis string

const (
	UpdateProductionScheduleSettingsRequestDemandBasisTrailing12  UpdateProductionScheduleSettingsRequestDemandBasis = "trailing_12"
	UpdateProductionScheduleSettingsRequestDemandBasisSeasonalEma UpdateProductionScheduleSettingsRequestDemandBasis = "seasonal_ema"
)

type OperationProductionScheduleSettingUpdateParams struct {
	// Request to replace the account's planning assumptions.
	UpdateProductionScheduleSettingsRequest UpdateProductionScheduleSettingsRequestParam
	paramObj
}

func (r OperationProductionScheduleSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductionScheduleSettingsRequest)
}
func (r *OperationProductionScheduleSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
