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

// Analyze sales, orders, manufacturing, materials, and other business metrics.
//
// CoreAnalyticsService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAnalyticsService] method instead.
type CoreAnalyticsService struct {
	options []option.RequestOption
}

// NewCoreAnalyticsService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreAnalyticsService(opts ...option.RequestOption) (r CoreAnalyticsService) {
	r = CoreAnalyticsService{}
	r.options = opts
	return
}

// Returns how reliably promised delivery dates were met.
//
// Orders are counted in the period their promise came due, not the period they
// shipped — an order promised in March and shipped in May is March's miss. On time
// means the first shipment left on or before the promised date, because the
// promise is that the order starts moving by then; judging on the last shipment
// would fail an order the customer received on time in two boxes. On time in full
// adds that the whole ordered quantity was packed.
//
// The denominator is orders that were due, not orders that shipped, so an order
// past its date and still unshipped counts against the rate rather than being held
// back until it moves. Excluding open orders would let a plant with a growing late
// backlog report perfect delivery.
//
// Only orders carrying a ship-by commitment participate. An order with no
// commitment cannot be late, and counting it as on time would inflate the rate
// with orders nobody promised anything about — `uncommitted_order_count` says how
// many were excluded, so the gap is visible rather than silent.
//
// Every rate is null rather than zero when nothing was due, and average lateness
// is measured over late orders only.
//
// The same window is also returned sliced by customer, customer group, product
// line, and the rule each ship-by date came from — each ordered worst-first, and
// each derived from the same set of orders as the headline so a drilldown always
// adds up to it. `by_product_line` is the one exception to that: an order spanning
// two lines is counted under both, because a late order is late for every line on
// it.
//
// Every filter is empty-means-all and they combine with AND. They narrow
// `uncommitted_order_count` too, so the excluded count always describes the same
// slice of the order book the rates do.
//
// This endpoint requires the permission: `sales_orders:read`.
func (r *CoreAnalyticsService) UpdateDeliveryPerformance(ctx context.Context, body CoreAnalyticsUpdateDeliveryPerformanceParams, opts ...option.RequestOption) (res *AnalyzeDeliveryPerformanceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/analytics/delivery-performance"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns Overall Equipment Effectiveness (OEE) metrics by department.
//
// Availability is the scheduled machine time the plant actually planned, net of
// logged machine downtime — the planned time comes from the published production
// schedule (or `planned_time` when supplied), and a department the schedule never
// covered has no availability rather than a fabricated one. Departments with
// `has_downtime_data` false have no downtime measured, and their ratios are
// returned as null rather than as 100%.
//
// This endpoint requires the permission: `machine_downtime:read`.
func (r *CoreAnalyticsService) UpdateOee(ctx context.Context, body CoreAnalyticsUpdateOeeParams, opts ...option.RequestOption) (res *AnalyzeOeeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/analytics/oee"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns Overall Equipment Effectiveness (OEE) by production week.
//
// Each period carries the same four terms `/v1/core/analytics/oee` reports for a
// single window, rolled up across departments and weighted by seconds rather than
// averaged, so a department that ran for an hour does not weigh as heavily as one
// that ran all week. Weeks start on Monday, and the first and last period of a
// window are clipped to the window itself.
//
// Only departments with scheduled time take part: a department with no machines
// has no availability, so counting its output in quality would leave the three
// terms describing different plants. Compare two windows by calling this twice.
//
// This endpoint requires the permission: `machine_downtime:read`.
func (r *CoreAnalyticsService) UpdateOeeTrend(ctx context.Context, body CoreAnalyticsUpdateOeeTrendParams, opts ...option.RequestOption) (res *AnalyzeOeeTrendResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/analytics/oee-trend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns actual production measured against the plan that was live at the time.
//
// The baseline for each week is the version that froze it — the plan committed for
// that week — so a version published after the week ended cannot rewrite a week
// the floor has already worked, while a plan published on the week's own start day
// still counts as the plan it froze. `baseline_schedules` names the versions used.
//
// Two ratios are returned because either alone misleads: `attainment_pct` caps
// each campaign at what was asked for, so over-building one SKU cannot hide a miss
// on another, while `output_ratio_pct` is uncapped and is what reveals
// over-production. Production with no matching planned campaign is reported as
// `unplanned_quantity` rather than discarded — that number is the clearest signal
// a schedule is being worked around.
//
// Every ratio is null rather than zero when nothing was planned, and
// `has_baseline` is false when nothing was ever published over the period.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *CoreAnalyticsService) UpdateScheduleAttainment(ctx context.Context, body CoreAnalyticsUpdateScheduleAttainmentParams, opts ...option.RequestOption) (res *AnalyzeScheduleAttainmentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/analytics/schedule-attainment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// AnalyzeDeliveryPerformanceRequest is the request to measure promises against
// shipments.
//
// The properties EndsAt, StartsAt are required.
type AnalyzeDeliveryPerformanceRequestParam struct {
	// The end date for the analysis period.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// The start date for the analysis period.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// Only measure orders whose customer sits in these groups.
	CustomerGroupIDs []string `json:"customer_group_ids,omitzero"`
	// Only measure orders bought by these customers. Their child accounts are
	// included, matching how the sales analytics resolve a customer.
	CustomerIDs []string `json:"customer_ids,omitzero"`
	// The period to break the results down by. Defaults to `week`.
	//
	// Any of "day", "week", "month".
	Granularity AnalyzeDeliveryPerformanceRequestGranularity `json:"granularity,omitzero"`
	// Only measure orders containing at least one line in these product lines.
	ProductLineIDs []string `json:"product_line_ids,omitzero"`
	// Only measure orders owned by these sales reps.
	SalesRepIDs []string `json:"sales_rep_ids,omitzero"`
	paramObj
}

func (r AnalyzeDeliveryPerformanceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AnalyzeDeliveryPerformanceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyzeDeliveryPerformanceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The period to break the results down by. Defaults to `week`.
type AnalyzeDeliveryPerformanceRequestGranularity string

const (
	AnalyzeDeliveryPerformanceRequestGranularityDay   AnalyzeDeliveryPerformanceRequestGranularity = "day"
	AnalyzeDeliveryPerformanceRequestGranularityWeek  AnalyzeDeliveryPerformanceRequestGranularity = "week"
	AnalyzeDeliveryPerformanceRequestGranularityMonth AnalyzeDeliveryPerformanceRequestGranularity = "month"
)

// How reliably promised delivery dates were met.
type AnalyzeDeliveryPerformanceResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Backlog ListDeliveryBacklogBucket `json:"backlog" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ByCommitmentSource ListDeliveryBreakdown `json:"by_commitment_source" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ByCustomer ListDeliveryBreakdown `json:"by_customer" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ByCustomerGroup ListDeliveryBreakdown `json:"by_customer_group" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ByProductLine ListDeliveryBreakdown `json:"by_product_line" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lateness ListDeliveryLatenessBucket `json:"lateness" api:"required"`
	// Resource type identifier.
	//
	// Any of "analyze_delivery_performance_response".
	Object AnalyzeDeliveryPerformanceResponseObject `json:"object" api:"required"`
	// Delivery reliability for one period, or for a whole window.
	Overall DeliveryPerformance `json:"overall" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Periods ListDeliveryPerformance `json:"periods" api:"required"`
	// Issued orders in the window carrying no ship-by date, excluded from every rate
	// above.
	//
	// Reported so the exclusion is visible: a delivery score computed over half the
	// order book, silently, is worse than one that says which half. A non-zero count
	// here means orders placed before commitments were tracked still need a ship-by
	// date.
	UncommittedOrderCount int64 `json:"uncommitted_order_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Backlog               respjson.Field
		ByCommitmentSource    respjson.Field
		ByCustomer            respjson.Field
		ByCustomerGroup       respjson.Field
		ByProductLine         respjson.Field
		Lateness              respjson.Field
		Object                respjson.Field
		Overall               respjson.Field
		Periods               respjson.Field
		UncommittedOrderCount respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyzeDeliveryPerformanceResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyzeDeliveryPerformanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AnalyzeDeliveryPerformanceResponseObject string

const (
	AnalyzeDeliveryPerformanceResponseObjectAnalyzeDeliveryPerformanceResponse AnalyzeDeliveryPerformanceResponseObject = "analyze_delivery_performance_response"
)

// AnalyzeOeeRequest is the request to analyze Overall Equipment Effectiveness
// (OEE).
//
// The properties EndsAt, StartsAt are required.
type AnalyzeOeeRequestParam struct {
	// The end date for the analysis period.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// The start date for the analysis period.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// Optional department IDs to filter by.
	DepartmentIDs []string `json:"department_ids,omitzero"`
	// Overrides the scheduled production time per department for the period. When
	// omitted it is taken from the published production schedule, so this is only
	// needed to measure a period the schedule does not cover. Availability,
	// performance and OEE are only returned for departments the scheduled time covers.
	PlannedTime []OeeDepartmentPlannedTimeParam `json:"planned_time,omitzero"`
	paramObj
}

func (r AnalyzeOeeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AnalyzeOeeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyzeOeeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AnalyzeOeeResponse represents the response from the analyze OEE endpoint.
type AnalyzeOeeResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Departments ListOeeDepartment `json:"departments" api:"required"`
	// Resource type identifier.
	//
	// Any of "analyze_oee_response".
	Object AnalyzeOeeResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Departments respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyzeOeeResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyzeOeeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AnalyzeOeeResponseObject string

const (
	AnalyzeOeeResponseObjectAnalyzeOeeResponse AnalyzeOeeResponseObject = "analyze_oee_response"
)

// AnalyzeOeeTrendRequest is the request to analyze Overall Equipment Effectiveness
// (OEE) over time.
//
// The properties EndsAt, StartsAt are required.
type AnalyzeOeeTrendRequestParam struct {
	// The end date for the analysis period.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// The start date for the analysis period.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// Restrict the analysis to these departments.
	DepartmentIDs []string `json:"department_ids,omitzero"`
	paramObj
}

func (r AnalyzeOeeTrendRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AnalyzeOeeTrendRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyzeOeeTrendRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AnalyzeOeeTrendResponse represents the response from the OEE trend endpoint.
type AnalyzeOeeTrendResponse struct {
	// Resource type identifier.
	//
	// Any of "analyze_oee_trend_response".
	Object AnalyzeOeeTrendResponseObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Periods ListOeeTrendPeriod `json:"periods" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		Periods     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyzeOeeTrendResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyzeOeeTrendResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AnalyzeOeeTrendResponseObject string

const (
	AnalyzeOeeTrendResponseObjectAnalyzeOeeTrendResponse AnalyzeOeeTrendResponseObject = "analyze_oee_trend_response"
)

// AnalyzeScheduleAttainmentRequest is the request to measure production against
// plan.
//
// The properties EndsAt, StartsAt are required.
type AnalyzeScheduleAttainmentRequestParam struct {
	// The end date for the analysis period.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// The start date for the analysis period.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// Only measure production in these departments.
	DepartmentIDs []string `json:"department_ids,omitzero"`
	// The dimension to break the results down by. Defaults to `week`.
	//
	// Any of "week", "machine", "department", "item".
	GroupBy AnalyzeScheduleAttainmentRequestGroupBy `json:"group_by,omitzero"`
	// Only measure production on these machines.
	MachineIDs []string `json:"machine_ids,omitzero"`
	paramObj
}

func (r AnalyzeScheduleAttainmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AnalyzeScheduleAttainmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AnalyzeScheduleAttainmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The dimension to break the results down by. Defaults to `week`.
type AnalyzeScheduleAttainmentRequestGroupBy string

const (
	AnalyzeScheduleAttainmentRequestGroupByWeek       AnalyzeScheduleAttainmentRequestGroupBy = "week"
	AnalyzeScheduleAttainmentRequestGroupByMachine    AnalyzeScheduleAttainmentRequestGroupBy = "machine"
	AnalyzeScheduleAttainmentRequestGroupByDepartment AnalyzeScheduleAttainmentRequestGroupBy = "department"
	AnalyzeScheduleAttainmentRequestGroupByItem       AnalyzeScheduleAttainmentRequestGroupBy = "item"
)

// Actual production measured against the plan that was live at the time.
//
// The baseline for each week is the version that was published on or before that
// week began, so republishing mid-horizon cannot rewrite a week the floor has
// already worked. `baseline_schedules` names the versions used, so any number here
// can be traced back to the plan that produced it.
type AnalyzeScheduleAttainmentResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	BaselineSchedules ListEntity `json:"baseline_schedules" api:"required"`
	// Whether the period had a plan to measure against. When `no_baseline`, every
	// ratio is null and the period has no plan rather than a missed one.
	//
	// Any of "measured", "no_baseline".
	BaselineStatus AnalyzeScheduleAttainmentResponseBaselineStatus `json:"baseline_status" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Buckets ListAttainmentBucket `json:"buckets" api:"required"`
	// End of the measured period.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	FrozenAdherence ListFrozenAdherence `json:"frozen_adherence" api:"required"`
	// The dimension the breakdown is grouped by.
	//
	// Any of "week", "machine", "department", "item".
	GroupBy AnalyzeScheduleAttainmentResponseGroupBy `json:"group_by" api:"required"`
	// Resource type identifier.
	//
	// Any of "analyze_schedule_attainment_response".
	Object AnalyzeScheduleAttainmentResponseObject `json:"object" api:"required"`
	// Machines the plan asked for over this window.
	//
	// Every figure in this response covers those machines only. Production scanned
	// onto a machine no published version scheduled is excluded outright, so the score
	// measures the plan that was made rather than the whole plant against it.
	ScheduledMachineCount int64 `json:"scheduled_machine_count" api:"required"`
	// Start of the measured period.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// One row of a schedule-attainment breakdown.
	//
	// Both ratios are reported because either alone misleads. `attainment_pct` caps
	// each SKU at what was asked for, so over-building one easy item cannot paper over
	// a total miss on another; `output_ratio_pct` does not cap, so it is the only one
	// that reveals over-production.
	Totals AttainmentBucket `json:"totals" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaselineSchedules     respjson.Field
		BaselineStatus        respjson.Field
		Buckets               respjson.Field
		EndsAt                respjson.Field
		FrozenAdherence       respjson.Field
		GroupBy               respjson.Field
		Object                respjson.Field
		ScheduledMachineCount respjson.Field
		StartsAt              respjson.Field
		Totals                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AnalyzeScheduleAttainmentResponse) RawJSON() string { return r.JSON.raw }
func (r *AnalyzeScheduleAttainmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the period had a plan to measure against. When `no_baseline`, every
// ratio is null and the period has no plan rather than a missed one.
type AnalyzeScheduleAttainmentResponseBaselineStatus string

const (
	AnalyzeScheduleAttainmentResponseBaselineStatusMeasured   AnalyzeScheduleAttainmentResponseBaselineStatus = "measured"
	AnalyzeScheduleAttainmentResponseBaselineStatusNoBaseline AnalyzeScheduleAttainmentResponseBaselineStatus = "no_baseline"
)

// The dimension the breakdown is grouped by.
type AnalyzeScheduleAttainmentResponseGroupBy string

const (
	AnalyzeScheduleAttainmentResponseGroupByWeek       AnalyzeScheduleAttainmentResponseGroupBy = "week"
	AnalyzeScheduleAttainmentResponseGroupByMachine    AnalyzeScheduleAttainmentResponseGroupBy = "machine"
	AnalyzeScheduleAttainmentResponseGroupByDepartment AnalyzeScheduleAttainmentResponseGroupBy = "department"
	AnalyzeScheduleAttainmentResponseGroupByItem       AnalyzeScheduleAttainmentResponseGroupBy = "item"
)

// Resource type identifier.
type AnalyzeScheduleAttainmentResponseObject string

const (
	AnalyzeScheduleAttainmentResponseObjectAnalyzeScheduleAttainmentResponse AnalyzeScheduleAttainmentResponseObject = "analyze_schedule_attainment_response"
)

// One row of a schedule-attainment breakdown.
//
// Both ratios are reported because either alone misleads. `attainment_pct` caps
// each SKU at what was asked for, so over-building one easy item cannot paper over
// a total miss on another; `output_ratio_pct` does not cap, so it is the only one
// that reveals over-production.
type AttainmentBucket struct {
	// Units actually produced.
	ActualQuantity float64 `json:"actual_quantity" api:"required"`
	// Share of the plan that was met. Null when nothing was planned.
	AttainmentPct float64 `json:"attainment_pct" api:"required"`
	// Batches scanned in this bucket.
	BatchCount int64 `json:"batch_count" api:"required"`
	// Identifies the bucket within the chosen grouping — a week start, machine ID,
	// department ID or item ID.
	Key string `json:"key" api:"required"`
	// Display label for the bucket.
	Label string `json:"label" api:"required"`
	// Units produced that were planned for, capped per campaign at what was asked.
	MatchedQuantity float64 `json:"matched_quantity" api:"required"`
	// Output as a share of plan, uncapped. Null when nothing was planned.
	OutputRatioPct float64 `json:"output_ratio_pct" api:"required"`
	// Planned campaigns in this bucket.
	PlannedLines int64 `json:"planned_lines" api:"required"`
	// Units the live plan called for.
	PlannedQuantity float64 `json:"planned_quantity" api:"required"`
	// Machine hours the plan called for.
	PlannedRunHours float64 `json:"planned_run_hours" api:"required"`
	// Units produced with no matching planned campaign.
	UnplannedQuantity float64 `json:"unplanned_quantity" api:"required"`
	// Units scrapped.
	WasteQuantity float64 `json:"waste_quantity" api:"required"`
	// First day of the week, when grouping by week.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActualQuantity    respjson.Field
		AttainmentPct     respjson.Field
		BatchCount        respjson.Field
		Key               respjson.Field
		Label             respjson.Field
		MatchedQuantity   respjson.Field
		OutputRatioPct    respjson.Field
		PlannedLines      respjson.Field
		PlannedQuantity   respjson.Field
		PlannedRunHours   respjson.Field
		UnplannedQuantity respjson.Field
		WasteQuantity     respjson.Field
		WeekStartsAt      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttainmentBucket) RawJSON() string { return r.JSON.raw }
func (r *AttainmentBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One age band of orders past their promise and still unshipped.
type DeliveryBacklogBucket struct {
	// Name of the band.
	Label string `json:"label" api:"required"`
	// Upper bound in days late; `0` means unbounded.
	MaxDaysLate int64 `json:"max_days_late" api:"required"`
	// Lower bound of the band in days late.
	MinDaysLate int64 `json:"min_days_late" api:"required"`
	// Resource type identifier.
	//
	// Any of "delivery_backlog_bucket".
	Object DeliveryBacklogBucketObject `json:"object" api:"required"`
	// Orders in the band.
	OrderCount int64 `json:"order_count" api:"required"`
	// Quantity still owed across them, which is what remains unpacked rather than what
	// was ordered.
	Units float64 `json:"units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label       respjson.Field
		MaxDaysLate respjson.Field
		MinDaysLate respjson.Field
		Object      respjson.Field
		OrderCount  respjson.Field
		Units       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryBacklogBucket) RawJSON() string { return r.JSON.raw }
func (r *DeliveryBacklogBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DeliveryBacklogBucketObject string

const (
	DeliveryBacklogBucketObjectDeliveryBacklogBucket DeliveryBacklogBucketObject = "delivery_backlog_bucket"
)

// Delivery performance for one slice of the order book.
type DeliveryBreakdown struct {
	// Identifier of the slice — a customer, customer group, product line, or
	// commitment source. Empty when the dimension is unset on the orders in it.
	Key string `json:"key" api:"required"`
	// Display name for the slice.
	Label string `json:"label" api:"required"`
	// Resource type identifier.
	//
	// Any of "delivery_breakdown".
	Object DeliveryBreakdownObject `json:"object" api:"required"`
	// Delivery reliability for one period, or for a whole window.
	Performance DeliveryPerformance `json:"performance" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Label       respjson.Field
		Object      respjson.Field
		Performance respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryBreakdown) RawJSON() string { return r.JSON.raw }
func (r *DeliveryBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DeliveryBreakdownObject string

const (
	DeliveryBreakdownObjectDeliveryBreakdown DeliveryBreakdownObject = "delivery_breakdown"
)

// One band of how far the window's misses missed by.
type DeliveryLatenessBucket struct {
	// Name of the band.
	Label string `json:"label" api:"required"`
	// Upper bound in days late; `0` means unbounded.
	MaxDaysLate int64 `json:"max_days_late" api:"required"`
	// Lower bound of the band in days late.
	MinDaysLate int64 `json:"min_days_late" api:"required"`
	// Resource type identifier.
	//
	// Any of "delivery_lateness_bucket".
	Object DeliveryLatenessBucketObject `json:"object" api:"required"`
	// Orders in the band, shipped and unshipped.
	OrderCount int64 `json:"order_count" api:"required"`
	// How many of them have since shipped. The remainder are still owed, and are the
	// same orders `backlog` counts.
	ShippedCount int64 `json:"shipped_count" api:"required"`
	// Quantity still unpacked across the band's orders.
	Units float64 `json:"units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label        respjson.Field
		MaxDaysLate  respjson.Field
		MinDaysLate  respjson.Field
		Object       respjson.Field
		OrderCount   respjson.Field
		ShippedCount respjson.Field
		Units        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryLatenessBucket) RawJSON() string { return r.JSON.raw }
func (r *DeliveryLatenessBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DeliveryLatenessBucketObject string

const (
	DeliveryLatenessBucketObjectDeliveryLatenessBucket DeliveryLatenessBucketObject = "delivery_lateness_bucket"
)

// Delivery reliability for one period, or for a whole window.
type DeliveryPerformance struct {
	// Average lead time these orders were promised.
	//
	// The gap between this and `average_lead_time_days` is what a lead time is
	// renegotiated on.
	AverageCommittedLeadTimeDays float64 `json:"average_committed_lead_time_days" api:"required"`
	// Average days late, over late orders only.
	//
	// Averaging over every order would dilute a real problem into a number that looks
	// fine.
	AverageDaysLate float64 `json:"average_days_late" api:"required"`
	// Average days from issue to first shipment, over orders that have shipped.
	AverageLeadTimeDays float64 `json:"average_lead_time_days" api:"required"`
	// Orders whose promised ship date fell in this period.
	//
	// This is the denominator for both rates below — orders that were due, not orders
	// that shipped. Measuring against shipments only would let unshipped late orders
	// disappear from the score.
	CommittedOrderCount int64 `json:"committed_order_count" api:"required"`
	// How many shipped late, plus those already past their date and still unshipped.
	LateOrderCount int64 `json:"late_order_count" api:"required"`
	// How many due in this period have not shipped at all.
	//
	// These count against on-time: a promise not yet met is not a promise kept.
	NotYetShippedCount int64 `json:"not_yet_shipped_count" api:"required"`
	// Resource type identifier.
	//
	// Any of "delivery_performance".
	Object DeliveryPerformanceObject `json:"object" api:"required"`
	// How many shipped on time and complete.
	OnTimeInFullCount int64 `json:"on_time_in_full_count" api:"required"`
	// Share of due orders that shipped on time and complete, as a percentage.
	OnTimeInFullPct float64 `json:"on_time_in_full_pct" api:"required"`
	// How many shipped on or before the promised date.
	OnTimeOrderCount int64 `json:"on_time_order_count" api:"required"`
	// Share of due orders that shipped on time, as a percentage.
	//
	// Null rather than zero when nothing was due, so a quiet week does not render as
	// total failure.
	OnTimePct float64 `json:"on_time_pct" api:"required"`
	// First day of the period; absent on the overall figure.
	PeriodStart time.Time `json:"period_start" api:"required" format:"date-time"`
	// How many of them have shipped at all.
	ShippedOrderCount int64 `json:"shipped_order_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AverageCommittedLeadTimeDays respjson.Field
		AverageDaysLate              respjson.Field
		AverageLeadTimeDays          respjson.Field
		CommittedOrderCount          respjson.Field
		LateOrderCount               respjson.Field
		NotYetShippedCount           respjson.Field
		Object                       respjson.Field
		OnTimeInFullCount            respjson.Field
		OnTimeInFullPct              respjson.Field
		OnTimeOrderCount             respjson.Field
		OnTimePct                    respjson.Field
		PeriodStart                  respjson.Field
		ShippedOrderCount            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeliveryPerformance) RawJSON() string { return r.JSON.raw }
func (r *DeliveryPerformance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DeliveryPerformanceObject string

const (
	DeliveryPerformanceObjectDeliveryPerformance DeliveryPerformanceObject = "delivery_performance"
)

// How well a published commitment survived the week it covered.
type FrozenAdherence struct {
	// Total absolute unit change across frozen-week deviations.
	AbsDeltaUnits float64 `json:"abs_delta_units" api:"required"`
	// Campaigns added into the frozen window after publish.
	AddedLines int64 `json:"added_lines" api:"required"`
	// Frozen campaigns that were changed after publish.
	DeviatedLines int64 `json:"deviated_lines" api:"required"`
	// Campaigns frozen at publish.
	FrozenLineCount int64 `json:"frozen_line_count" api:"required"`
	// Units frozen at publish.
	FrozenPlannedQuantity float64 `json:"frozen_planned_quantity" api:"required"`
	// Last day of the frozen window.
	FrozenThroughAt time.Time `json:"frozen_through_at" api:"required" format:"date-time"`
	// Share of frozen campaigns that survived untouched. Null when nothing was frozen.
	LineAdherencePct float64 `json:"line_adherence_pct" api:"required"`
	// Campaigns the floor ran inside the frozen window that the frozen plan never
	// called for, counted per machine-week-SKU.
	//
	// Working around a commitment breaks it as surely as editing it does, so this
	// scores alongside the hand edits rather than beside them.
	OffPlanLines int64 `json:"off_plan_lines" api:"required"`
	// Units behind those off-plan campaigns.
	OffPlanQuantity float64 `json:"off_plan_quantity" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Schedule Entity `json:"schedule" api:"required"`
	// Share of frozen units that survived untouched. Null when nothing was frozen.
	UnitsAdherencePct float64 `json:"units_adherence_pct" api:"required"`
	// Version number of that schedule.
	Version int64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbsDeltaUnits         respjson.Field
		AddedLines            respjson.Field
		DeviatedLines         respjson.Field
		FrozenLineCount       respjson.Field
		FrozenPlannedQuantity respjson.Field
		FrozenThroughAt       respjson.Field
		LineAdherencePct      respjson.Field
		OffPlanLines          respjson.Field
		OffPlanQuantity       respjson.Field
		Schedule              respjson.Field
		UnitsAdherencePct     respjson.Field
		Version               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FrozenAdherence) RawJSON() string { return r.JSON.raw }
func (r *FrozenAdherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAttainmentBucket struct {
	// Resources in this page.
	Data []AttainmentBucket `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAttainmentBucketObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListAttainmentBucket) RawJSON() string { return r.JSON.raw }
func (r *ListAttainmentBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAttainmentBucketObject string

const (
	ListAttainmentBucketObjectList ListAttainmentBucketObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDeliveryBacklogBucket struct {
	// Resources in this page.
	Data []DeliveryBacklogBucket `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDeliveryBacklogBucketObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDeliveryBacklogBucket) RawJSON() string { return r.JSON.raw }
func (r *ListDeliveryBacklogBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDeliveryBacklogBucketObject string

const (
	ListDeliveryBacklogBucketObjectList ListDeliveryBacklogBucketObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDeliveryBreakdown struct {
	// Resources in this page.
	Data []DeliveryBreakdown `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDeliveryBreakdownObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDeliveryBreakdown) RawJSON() string { return r.JSON.raw }
func (r *ListDeliveryBreakdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDeliveryBreakdownObject string

const (
	ListDeliveryBreakdownObjectList ListDeliveryBreakdownObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDeliveryLatenessBucket struct {
	// Resources in this page.
	Data []DeliveryLatenessBucket `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDeliveryLatenessBucketObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDeliveryLatenessBucket) RawJSON() string { return r.JSON.raw }
func (r *ListDeliveryLatenessBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDeliveryLatenessBucketObject string

const (
	ListDeliveryLatenessBucketObjectList ListDeliveryLatenessBucketObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDeliveryPerformance struct {
	// Resources in this page.
	Data []DeliveryPerformance `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDeliveryPerformanceObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListDeliveryPerformance) RawJSON() string { return r.JSON.raw }
func (r *ListDeliveryPerformance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDeliveryPerformanceObject string

const (
	ListDeliveryPerformanceObjectList ListDeliveryPerformanceObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListFrozenAdherence struct {
	// Resources in this page.
	Data []FrozenAdherence `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListFrozenAdherenceObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListFrozenAdherence) RawJSON() string { return r.JSON.raw }
func (r *ListFrozenAdherence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListFrozenAdherenceObject string

const (
	ListFrozenAdherenceObjectList ListFrozenAdherenceObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOeeDepartment struct {
	// Resources in this page.
	Data []OeeDepartment `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOeeDepartmentObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListOeeDepartment) RawJSON() string { return r.JSON.raw }
func (r *ListOeeDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOeeDepartmentObject string

const (
	ListOeeDepartmentObjectList ListOeeDepartmentObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOeeDowntimeReason struct {
	// Resources in this page.
	Data []OeeDowntimeReason `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOeeDowntimeReasonObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListOeeDowntimeReason) RawJSON() string { return r.JSON.raw }
func (r *ListOeeDowntimeReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOeeDowntimeReasonObject string

const (
	ListOeeDowntimeReasonObjectList ListOeeDowntimeReasonObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOeeTrendPeriod struct {
	// Resources in this page.
	Data []OeeTrendPeriod `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOeeTrendPeriodObject `json:"object" api:"required"`
	// PageInfo describes where the current page sits within a paginated result set and
	// how to move to the adjacent pages.
	//
	// Page a list by following the URLs below rather than assembling cursors yourself.
	// For a top-level list endpoint the URL repeats the original request's query
	// string with only the cursor swapped, so following it preserves the same filters,
	// search term, and page size.
	PageInfo PageInfo `json:"page_info" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Object      respjson.Field
		PageInfo    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListOeeTrendPeriod) RawJSON() string { return r.JSON.raw }
func (r *ListOeeTrendPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOeeTrendPeriodObject string

const (
	ListOeeTrendPeriodObjectList ListOeeTrendPeriodObject = "list"
)

// OeeDepartment represents OEE metrics for a single department.
type OeeDepartment struct {
	// Data-quality warnings for this grouping. Empty when the numbers can be taken at
	// face value.
	//
	// Any of "performance_above_capacity".
	Anomalies []string `json:"anomalies" api:"required"`
	// Logged downtime charged against availability, in seconds.
	AvailabilityLossSeconds float64 `json:"availability_loss_seconds" api:"required"`
	// Run time divided by scheduled time.
	AvailabilityPct float64 `json:"availability_pct" api:"required"`
	// Time spent changing over between products, in seconds.
	ChangeoverSeconds float64 `json:"changeover_seconds" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Department Entity `json:"department" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	DowntimeBreakdown ListOeeDowntimeReason `json:"downtime_breakdown" api:"required"`
	// Number of downtime events logged in the period.
	DowntimeEventCount int64 `json:"downtime_event_count" api:"required"`
	// The estimated runtime in hours.
	EstimatedRuntimeHours float64 `json:"estimated_runtime_hours" api:"required"`
	// The number of good units produced.
	GoodUnits float64 `json:"good_units" api:"required"`
	// Whether availability was measured from logged downtime or estimated from
	// runtime. A department with no logged downtime computes as perfectly available,
	// so an estimate is labelled rather than presented as a measurement.
	//
	// Any of "measured", "estimated".
	MeasurementStatus OeeDepartmentMeasurementStatus `json:"measurement_status" api:"required"`
	// Time nobody planned to run, removed from the OEE denominator rather than counted
	// as a loss.
	NotScheduledSeconds float64 `json:"not_scheduled_seconds" api:"required"`
	// Availability multiplied by performance multiplied by quality.
	OeePct float64 `json:"oee_pct" api:"required"`
	// Logged downtime charged against performance, in seconds.
	PerformanceLossSeconds float64 `json:"performance_loss_seconds" api:"required"`
	// Standard seconds earned divided by run time: how fast the department ran against
	// the designed speed of its production steps.
	PerformancePct float64 `json:"performance_pct" api:"required"`
	// Logged downtime charged against quality, in seconds.
	QualityLossSeconds float64 `json:"quality_loss_seconds" api:"required"`
	// Good units divided by total units produced.
	QualityPct float64 `json:"quality_pct" api:"required"`
	// Scheduled time net of availability losses, in seconds.
	RunTimeSeconds float64 `json:"run_time_seconds" api:"required"`
	// Planned time net of not-scheduled downtime, in seconds.
	ScheduledSeconds float64 `json:"scheduled_seconds" api:"required"`
	// The number of seconds units.
	SecondsUnits float64 `json:"seconds_units" api:"required"`
	// The time this output should have taken at each production step's own labor rate:
	// ideal cycle time multiplied by the units produced. This is the numerator of
	// Performance.
	StandardSecondsEarned float64 `json:"standard_seconds_earned" api:"required"`
	// The number of waste units.
	WasteUnits float64 `json:"waste_units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Anomalies               respjson.Field
		AvailabilityLossSeconds respjson.Field
		AvailabilityPct         respjson.Field
		ChangeoverSeconds       respjson.Field
		Department              respjson.Field
		DowntimeBreakdown       respjson.Field
		DowntimeEventCount      respjson.Field
		EstimatedRuntimeHours   respjson.Field
		GoodUnits               respjson.Field
		MeasurementStatus       respjson.Field
		NotScheduledSeconds     respjson.Field
		OeePct                  respjson.Field
		PerformanceLossSeconds  respjson.Field
		PerformancePct          respjson.Field
		QualityLossSeconds      respjson.Field
		QualityPct              respjson.Field
		RunTimeSeconds          respjson.Field
		ScheduledSeconds        respjson.Field
		SecondsUnits            respjson.Field
		StandardSecondsEarned   respjson.Field
		WasteUnits              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OeeDepartment) RawJSON() string { return r.JSON.raw }
func (r *OeeDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether availability was measured from logged downtime or estimated from
// runtime. A department with no logged downtime computes as perfectly available,
// so an estimate is labelled rather than presented as a measurement.
type OeeDepartmentMeasurementStatus string

const (
	OeeDepartmentMeasurementStatusMeasured  OeeDepartmentMeasurementStatus = "measured"
	OeeDepartmentMeasurementStatusEstimated OeeDepartmentMeasurementStatus = "estimated"
)

// OeeDepartmentPlannedTime supplies the scheduled production time for one
// department.
//
// The properties DepartmentID, PlannedHours are required.
type OeeDepartmentPlannedTimeParam struct {
	// The department ID.
	DepartmentID string `json:"department_id" api:"required"`
	// Scheduled production hours for the period.
	PlannedHours float64 `json:"planned_hours" api:"required"`
	paramObj
}

func (r OeeDepartmentPlannedTimeParam) MarshalJSON() (data []byte, err error) {
	type shadow OeeDepartmentPlannedTimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OeeDepartmentPlannedTimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OeeDowntimeReason represents one reason's contribution to a department's
// downtime.
type OeeDowntimeReason struct {
	// Downtime attributed to this reason, in seconds.
	DowntimeSeconds float64 `json:"downtime_seconds" api:"required"`
	// Number of events logged against this reason.
	EventCount int64 `json:"event_count" api:"required"`
	// Which OEE term this reason charges.
	//
	// Any of "availability", "performance", "quality", "not_scheduled".
	OeeBucket OeeDowntimeReasonOeeBucket `json:"oee_bucket" api:"required"`
	// Why the machine stopped.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Reason OeeDowntimeReasonReason `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DowntimeSeconds respjson.Field
		EventCount      respjson.Field
		OeeBucket       respjson.Field
		Reason          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OeeDowntimeReason) RawJSON() string { return r.JSON.raw }
func (r *OeeDowntimeReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which OEE term this reason charges.
type OeeDowntimeReasonOeeBucket string

const (
	OeeDowntimeReasonOeeBucketAvailability OeeDowntimeReasonOeeBucket = "availability"
	OeeDowntimeReasonOeeBucketPerformance  OeeDowntimeReasonOeeBucket = "performance"
	OeeDowntimeReasonOeeBucketQuality      OeeDowntimeReasonOeeBucket = "quality"
	OeeDowntimeReasonOeeBucketNotScheduled OeeDowntimeReasonOeeBucket = "not_scheduled"
)

// Why the machine stopped.
type OeeDowntimeReasonReason string

const (
	OeeDowntimeReasonReasonBreakdown          OeeDowntimeReasonReason = "breakdown"
	OeeDowntimeReasonReasonChangeover         OeeDowntimeReasonReason = "changeover"
	OeeDowntimeReasonReasonMaterialShortage   OeeDowntimeReasonReason = "material_shortage"
	OeeDowntimeReasonReasonNoOperator         OeeDowntimeReasonReason = "no_operator"
	OeeDowntimeReasonReasonPlannedMaintenance OeeDowntimeReasonReason = "planned_maintenance"
	OeeDowntimeReasonReasonMinorStop          OeeDowntimeReasonReason = "minor_stop"
	OeeDowntimeReasonReasonQualityHold        OeeDowntimeReasonReason = "quality_hold"
	OeeDowntimeReasonReasonNoSchedule         OeeDowntimeReasonReason = "no_schedule"
)

// OeeTrendPeriod represents one production week of OEE, rolled up across the
// departments that had scheduled time in it. Departments with no scheduled time
// have no OEE and take no part in the roll-up, so their output is not counted here
// either.
type OeeTrendPeriod struct {
	// Logged downtime charged against availability, in seconds.
	AvailabilityLossSeconds float64 `json:"availability_loss_seconds" api:"required"`
	// Run time divided by scheduled time.
	AvailabilityPct float64 `json:"availability_pct" api:"required"`
	// Number of downtime events overlapping this period.
	DowntimeEventCount int64 `json:"downtime_event_count" api:"required"`
	// The instant this period ends, exclusive.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// The number of good units produced.
	GoodUnits float64 `json:"good_units" api:"required"`
	// Whether availability was measured from logged downtime or estimated from
	// runtime.
	//
	// Any of "measured", "estimated".
	MeasurementStatus OeeTrendPeriodMeasurementStatus `json:"measurement_status" api:"required"`
	// Time nobody planned to run, removed from the denominator rather than counted as
	// a loss.
	NotScheduledSeconds float64 `json:"not_scheduled_seconds" api:"required"`
	// Availability multiplied by performance multiplied by quality.
	OeePct float64 `json:"oee_pct" api:"required"`
	// Standard seconds earned divided by run time.
	PerformancePct float64 `json:"performance_pct" api:"required"`
	// Good units divided by total units produced.
	QualityPct float64 `json:"quality_pct" api:"required"`
	// Scheduled time net of availability losses, in seconds.
	RunTimeSeconds float64 `json:"run_time_seconds" api:"required"`
	// Planned time net of not-scheduled downtime, in seconds.
	ScheduledSeconds float64 `json:"scheduled_seconds" api:"required"`
	// The number of seconds units.
	SecondsUnits float64 `json:"seconds_units" api:"required"`
	// The time this output should have taken at each production step's own labor rate:
	// ideal cycle time multiplied by the units produced.
	StandardSecondsEarned float64 `json:"standard_seconds_earned" api:"required"`
	// The first instant this period covers. Weeks start on Monday; the first and last
	// periods of a window are clipped to the window itself.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	// The number of waste units.
	WasteUnits float64 `json:"waste_units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailabilityLossSeconds respjson.Field
		AvailabilityPct         respjson.Field
		DowntimeEventCount      respjson.Field
		EndsAt                  respjson.Field
		GoodUnits               respjson.Field
		MeasurementStatus       respjson.Field
		NotScheduledSeconds     respjson.Field
		OeePct                  respjson.Field
		PerformancePct          respjson.Field
		QualityPct              respjson.Field
		RunTimeSeconds          respjson.Field
		ScheduledSeconds        respjson.Field
		SecondsUnits            respjson.Field
		StandardSecondsEarned   respjson.Field
		StartsAt                respjson.Field
		WasteUnits              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OeeTrendPeriod) RawJSON() string { return r.JSON.raw }
func (r *OeeTrendPeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether availability was measured from logged downtime or estimated from
// runtime.
type OeeTrendPeriodMeasurementStatus string

const (
	OeeTrendPeriodMeasurementStatusMeasured  OeeTrendPeriodMeasurementStatus = "measured"
	OeeTrendPeriodMeasurementStatusEstimated OeeTrendPeriodMeasurementStatus = "estimated"
)

type CoreAnalyticsUpdateDeliveryPerformanceParams struct {
	// AnalyzeDeliveryPerformanceRequest is the request to measure promises against
	// shipments.
	AnalyzeDeliveryPerformanceRequest AnalyzeDeliveryPerformanceRequestParam
	paramObj
}

func (r CoreAnalyticsUpdateDeliveryPerformanceParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AnalyzeDeliveryPerformanceRequest)
}
func (r *CoreAnalyticsUpdateDeliveryPerformanceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoreAnalyticsUpdateOeeParams struct {
	// AnalyzeOeeRequest is the request to analyze Overall Equipment Effectiveness
	// (OEE).
	AnalyzeOeeRequest AnalyzeOeeRequestParam
	paramObj
}

func (r CoreAnalyticsUpdateOeeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AnalyzeOeeRequest)
}
func (r *CoreAnalyticsUpdateOeeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoreAnalyticsUpdateOeeTrendParams struct {
	// AnalyzeOeeTrendRequest is the request to analyze Overall Equipment Effectiveness
	// (OEE) over time.
	AnalyzeOeeTrendRequest AnalyzeOeeTrendRequestParam
	paramObj
}

func (r CoreAnalyticsUpdateOeeTrendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AnalyzeOeeTrendRequest)
}
func (r *CoreAnalyticsUpdateOeeTrendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoreAnalyticsUpdateScheduleAttainmentParams struct {
	// AnalyzeScheduleAttainmentRequest is the request to measure production against
	// plan.
	AnalyzeScheduleAttainmentRequest AnalyzeScheduleAttainmentRequestParam
	paramObj
}

func (r CoreAnalyticsUpdateScheduleAttainmentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AnalyzeScheduleAttainmentRequest)
}
func (r *CoreAnalyticsUpdateScheduleAttainmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
