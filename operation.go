// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// OperationService contains methods and other services that help with interacting
// with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationService] method instead.
type OperationService struct {
	options []option.RequestOption
	// List and manage shipping terms.
	ShippingTerms OperationShippingTermService
	// List and manage carriers and their Shippo integrations.
	Carriers OperationCarrierService
	// List and manage departments.
	Departments OperationDepartmentService
	// List and export inventory change logs.
	InventoryChangeLogs OperationInventoryChangeLogService
	// List and manage machines.
	Machines OperationMachineService
	// Log and review machine stoppages. Downtime is the source of OEE availability and
	// changeover time.
	MachineDowntimeEvents OperationMachineDowntimeEventService
	// Adjust the demand a production schedule plans against. Overrides are how
	// management accounts for demand that sales history cannot see.
	DemandOverrides OperationDemandOverrideService
	// Generate and review machine-level production schedules.
	ProductionSchedules OperationProductionScheduleService
	// The planning assumptions production schedules are solved against, and the
	// per-resource overrides that mark which machines constrain the plan.
	ProductionScheduleSettings OperationProductionScheduleSettingService
	// The planning assumptions production schedules are solved against, and the
	// per-resource overrides that mark which machines constrain the plan.
	FulfillmentRecommendations OperationFulfillmentRecommendationService
	// The days a plant tenders freight and a customer's dock accepts it, less the
	// holidays and shutdowns either side is closed for. Every ship-by date is resolved
	// against them, so an order is never committed to a day nobody can act on.
	OperatingCalendars OperationOperatingCalendarService
	// List, view, pick, void, and pack picks and pick lines.
	Picks OperationPickService
	// List and manage locations.
	Locations OperationLocationService
	// List and manage locations.
	LocationTypes OperationLocationTypeService
	Shipments     OperationShipmentService
	// List and manage scanning stations.
	ScanningStations OperationScanningStationService
}

// NewOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationService(opts ...option.RequestOption) (r OperationService) {
	r = OperationService{}
	r.options = opts
	r.ShippingTerms = NewOperationShippingTermService(opts...)
	r.Carriers = NewOperationCarrierService(opts...)
	r.Departments = NewOperationDepartmentService(opts...)
	r.InventoryChangeLogs = NewOperationInventoryChangeLogService(opts...)
	r.Machines = NewOperationMachineService(opts...)
	r.MachineDowntimeEvents = NewOperationMachineDowntimeEventService(opts...)
	r.DemandOverrides = NewOperationDemandOverrideService(opts...)
	r.ProductionSchedules = NewOperationProductionScheduleService(opts...)
	r.ProductionScheduleSettings = NewOperationProductionScheduleSettingService(opts...)
	r.FulfillmentRecommendations = NewOperationFulfillmentRecommendationService(opts...)
	r.OperatingCalendars = NewOperationOperatingCalendarService(opts...)
	r.Picks = NewOperationPickService(opts...)
	r.Locations = NewOperationLocationService(opts...)
	r.LocationTypes = NewOperationLocationTypeService(opts...)
	r.Shipments = NewOperationShipmentService(opts...)
	r.ScanningStations = NewOperationScanningStationService(opts...)
	return
}

// Returns the demand override types, which describe how an override's value
// adjusts the forecast.
//
// The taxonomy is platform-provided and identical for every account; each type's
// `code` is a value accepted as an override's `adjustment`.
//
// This endpoint requires the permission: `demand_overrides:read`.
func (r *OperationService) GetDemandOverrideTypes(ctx context.Context, opts ...option.RequestOption) (res *ListDemandOverrideType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/demand-override-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the downtime reasons available when logging a stoppage.
//
// The list is the same for every account and is ordered for display, so it can be
// rendered straight into a reason picker. Each reason carries the OEE term its
// stoppages charge, which is what makes the choice of reason matter beyond
// labeling.
//
// This endpoint requires the permission: `machine_downtime:read`.
func (r *OperationService) GetMachineDowntimeReasons(ctx context.Context, opts ...option.RequestOption) (res *ListMachineDowntimeReason, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/machine-downtime-reasons"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns what every machine is running right now, how much is left on it, and
// what is queued behind that.
//
// The whole floor comes back in one response rather than a page at a time, so a
// wall display can render it in a single call.
//
// Assembled from the published schedule, the batches the floor has scanned against
// each campaign, and any open downtime. A campaign is `current` once its week is
// released and while it still has batches to scan; when the last one is scanned it
// hands over to the next, so this advances on its own as a shift progresses.
//
// A machine with an open stoppage reads `down` even when it has a released
// campaign, because a broken machine is not producing whatever the plan says. A
// machine with nothing released reads `idle`, which is a state worth seeing rather
// than an absence from the list.
//
// Reads the published version rather than the newest draft: the floor works to
// what was committed, and a draft regenerating underneath a wall display would
// make machines appear to change job on their own. With nothing published every
// machine reads idle rather than the request failing.
//
// This endpoint requires the permission: `machines:read`.
func (r *OperationService) GetMachineStatus(ctx context.Context, query OperationGetMachineStatusParams, opts ...option.RequestOption) (res *ListMachineStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/machine-status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the kinds of hand change a schedule deviation can record.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationService) GetScheduleDeviationTypes(ctx context.Context, opts ...option.RequestOption) (res *ListScheduleDeviationType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/schedule-deviation-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A way of adjusting planned demand.
//
// `absolute` replaces the forecast for each month an override covers,
// `delta_units` adds to it, and `delta_percent` scales it. When several overrides
// land on the same month they are applied in that order.
type DemandOverrideType struct {
	// Override type ID.
	ID string `json:"id" api:"required"`
	// The value to send as an override's `adjustment`.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Code DemandOverrideTypeCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the type.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "demand_override_type".
	Object DemandOverrideTypeObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DemandOverrideType) RawJSON() string { return r.JSON.raw }
func (r *DemandOverrideType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The value to send as an override's `adjustment`.
type DemandOverrideTypeCode string

const (
	DemandOverrideTypeCodeAbsolute     DemandOverrideTypeCode = "absolute"
	DemandOverrideTypeCodeDeltaUnits   DemandOverrideTypeCode = "delta_units"
	DemandOverrideTypeCodeDeltaPercent DemandOverrideTypeCode = "delta_percent"
)

// Resource type identifier.
type DemandOverrideTypeObject string

const (
	DemandOverrideTypeObjectDemandOverrideType DemandOverrideTypeObject = "demand_override_type"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDemandOverrideType struct {
	// Resources in this page.
	Data []DemandOverrideType `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDemandOverrideTypeObject `json:"object" api:"required"`
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
func (r ListDemandOverrideType) RawJSON() string { return r.JSON.raw }
func (r *ListDemandOverrideType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDemandOverrideTypeObject string

const (
	ListDemandOverrideTypeObjectList ListDemandOverrideTypeObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMachineDowntimeReason struct {
	// Resources in this page.
	Data []MachineDowntimeReason `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMachineDowntimeReasonObject `json:"object" api:"required"`
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
func (r ListMachineDowntimeReason) RawJSON() string { return r.JSON.raw }
func (r *ListMachineDowntimeReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMachineDowntimeReasonObject string

const (
	ListMachineDowntimeReasonObjectList ListMachineDowntimeReasonObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMachineStatus struct {
	// Resources in this page.
	Data []MachineStatus `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMachineStatusObject `json:"object" api:"required"`
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
func (r ListMachineStatus) RawJSON() string { return r.JSON.raw }
func (r *ListMachineStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMachineStatusObject string

const (
	ListMachineStatusObjectList ListMachineStatusObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleDeviationType struct {
	// Resources in this page.
	Data []ScheduleDeviationType `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleDeviationTypeObject `json:"object" api:"required"`
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
func (r ListScheduleDeviationType) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleDeviationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleDeviationTypeObject string

const (
	ListScheduleDeviationTypeObjectList ListScheduleDeviationTypeObject = "list"
)

// One campaign on a machine, with how far through it the floor is.
//
// A campaign is one item scheduled to run on one machine for one week. Progress is
// taken from the batches the floor has scanned against it rather than reported by
// hand, so it advances on its own as a shift runs.
type MachineCampaign struct {
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Quantity the plan asked for.
	PlannedQuantity float64 `json:"planned_quantity" api:"required"`
	// Machine hours the plan allocates to the campaign.
	PlannedRunHours float64 `json:"planned_run_hours" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionRun Entity `json:"production_run" api:"required"`
	// Batches issued to the floor for this campaign.
	ReleasedBatchCount int64 `json:"released_batch_count" api:"required"`
	// Quantity still to make.
	//
	// Never negative: an over-run shows up in `scanned_quantity` rather than as
	// negative remaining work.
	RemainingQuantity float64 `json:"remaining_quantity" api:"required"`
	// Batches of this campaign the floor has scanned.
	ScannedBatchCount int64 `json:"scanned_batch_count" api:"required"`
	// Quantity the floor has scanned so far.
	ScannedQuantity float64 `json:"scanned_quantity" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ScheduleLine Entity `json:"schedule_line" api:"required"`
	// SKU of the item.
	SKU string `json:"sku" api:"required"`
	// Where the campaign is in its lifecycle.
	//
	//   - `planned`: scheduled, but not yet released to the floor.
	//   - `released`: issued to the floor as a production run, so batches can be scanned
	//     against it.
	//   - `in_progress`: being run.
	//   - `complete`: finished.
	//   - `cancelled`: will not be run.
	//
	// Any of "planned", "released", "in_progress", "complete", "cancelled".
	Status MachineCampaignStatus `json:"status" api:"required"`
	// Unit the quantities are counted in.
	Unit string `json:"unit" api:"required"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// First day of the week the campaign belongs to.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item               respjson.Field
		PlannedQuantity    respjson.Field
		PlannedRunHours    respjson.Field
		ProductionRun      respjson.Field
		ReleasedBatchCount respjson.Field
		RemainingQuantity  respjson.Field
		ScannedBatchCount  respjson.Field
		ScannedQuantity    respjson.Field
		ScheduleLine       respjson.Field
		SKU                respjson.Field
		Status             respjson.Field
		Unit               respjson.Field
		WeekIndex          respjson.Field
		WeekStartsAt       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineCampaign) RawJSON() string { return r.JSON.raw }
func (r *MachineCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the campaign is in its lifecycle.
//
//   - `planned`: scheduled, but not yet released to the floor.
//   - `released`: issued to the floor as a production run, so batches can be scanned
//     against it.
//   - `in_progress`: being run.
//   - `complete`: finished.
//   - `cancelled`: will not be run.
type MachineCampaignStatus string

const (
	MachineCampaignStatusPlanned    MachineCampaignStatus = "planned"
	MachineCampaignStatusReleased   MachineCampaignStatus = "released"
	MachineCampaignStatusInProgress MachineCampaignStatus = "in_progress"
	MachineCampaignStatusComplete   MachineCampaignStatus = "complete"
	MachineCampaignStatusCancelled  MachineCampaignStatus = "cancelled"
)

// A reason a machine stopped running.
//
// The `oee_bucket` decides which OEE term the stoppage charges: `availability`
// losses reduce run time, `performance` losses are minor stops and speed loss,
// `quality` losses cover rework and holds, and `not_scheduled` time is removed
// from the OEE calculation entirely rather than counted against it.
type MachineDowntimeReason struct {
	// Downtime reason ID.
	ID string `json:"id" api:"required"`
	// Stable code used when logging downtime.
	//
	// This is the value to send as `reason` when creating or updating a downtime
	// event.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Code MachineDowntimeReasonCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the reason.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine_downtime_reason".
	Object MachineDowntimeReasonObject `json:"object" api:"required"`
	// Which OEE term this reason charges.
	//
	// Any of "availability", "performance", "quality", "not_scheduled".
	OeeBucket MachineDowntimeReasonOeeBucket `json:"oee_bucket" api:"required"`
	// Whether the stoppage was scheduled in advance, such as preventive maintenance.
	//
	// Any of "planned", "unplanned".
	PlanningStatus MachineDowntimeReasonPlanningStatus `json:"planning_status" api:"required"`
	// Display order, ascending.
	SortOrder int64 `json:"sort_order" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Code           respjson.Field
		CreatedAt      respjson.Field
		Name           respjson.Field
		Object         respjson.Field
		OeeBucket      respjson.Field
		PlanningStatus respjson.Field
		SortOrder      respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineDowntimeReason) RawJSON() string { return r.JSON.raw }
func (r *MachineDowntimeReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable code used when logging downtime.
//
// This is the value to send as `reason` when creating or updating a downtime
// event.
type MachineDowntimeReasonCode string

const (
	MachineDowntimeReasonCodeBreakdown          MachineDowntimeReasonCode = "breakdown"
	MachineDowntimeReasonCodeChangeover         MachineDowntimeReasonCode = "changeover"
	MachineDowntimeReasonCodeMaterialShortage   MachineDowntimeReasonCode = "material_shortage"
	MachineDowntimeReasonCodeNoOperator         MachineDowntimeReasonCode = "no_operator"
	MachineDowntimeReasonCodePlannedMaintenance MachineDowntimeReasonCode = "planned_maintenance"
	MachineDowntimeReasonCodeMinorStop          MachineDowntimeReasonCode = "minor_stop"
	MachineDowntimeReasonCodeQualityHold        MachineDowntimeReasonCode = "quality_hold"
	MachineDowntimeReasonCodeNoSchedule         MachineDowntimeReasonCode = "no_schedule"
)

// Resource type identifier.
type MachineDowntimeReasonObject string

const (
	MachineDowntimeReasonObjectMachineDowntimeReason MachineDowntimeReasonObject = "machine_downtime_reason"
)

// Which OEE term this reason charges.
type MachineDowntimeReasonOeeBucket string

const (
	MachineDowntimeReasonOeeBucketAvailability MachineDowntimeReasonOeeBucket = "availability"
	MachineDowntimeReasonOeeBucketPerformance  MachineDowntimeReasonOeeBucket = "performance"
	MachineDowntimeReasonOeeBucketQuality      MachineDowntimeReasonOeeBucket = "quality"
	MachineDowntimeReasonOeeBucketNotScheduled MachineDowntimeReasonOeeBucket = "not_scheduled"
)

// Whether the stoppage was scheduled in advance, such as preventive maintenance.
type MachineDowntimeReasonPlanningStatus string

const (
	MachineDowntimeReasonPlanningStatusPlanned   MachineDowntimeReasonPlanningStatus = "planned"
	MachineDowntimeReasonPlanningStatusUnplanned MachineDowntimeReasonPlanningStatus = "unplanned"
)

// The reason for a stoppage, as carried on a downtime event.
//
// A denormalized view of the reason taxonomy: the stable code plus the display
// name and OEE bucket resolved from it at read time.
type MachineDowntimeReasonSummary struct {
	// Stable code identifying the reason.
	//
	// Any of "breakdown", "changeover", "material_shortage", "no_operator",
	// "planned_maintenance", "minor_stop", "quality_hold", "no_schedule".
	Code MachineDowntimeReasonSummaryCode `json:"code" api:"required"`
	// Display name of the reason.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine_downtime_reason".
	Object MachineDowntimeReasonSummaryObject `json:"object" api:"required"`
	// Which OEE term this reason charges.
	//
	// Any of "availability", "performance", "quality", "not_scheduled".
	OeeBucket MachineDowntimeReasonSummaryOeeBucket `json:"oee_bucket" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		OeeBucket   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineDowntimeReasonSummary) RawJSON() string { return r.JSON.raw }
func (r *MachineDowntimeReasonSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable code identifying the reason.
type MachineDowntimeReasonSummaryCode string

const (
	MachineDowntimeReasonSummaryCodeBreakdown          MachineDowntimeReasonSummaryCode = "breakdown"
	MachineDowntimeReasonSummaryCodeChangeover         MachineDowntimeReasonSummaryCode = "changeover"
	MachineDowntimeReasonSummaryCodeMaterialShortage   MachineDowntimeReasonSummaryCode = "material_shortage"
	MachineDowntimeReasonSummaryCodeNoOperator         MachineDowntimeReasonSummaryCode = "no_operator"
	MachineDowntimeReasonSummaryCodePlannedMaintenance MachineDowntimeReasonSummaryCode = "planned_maintenance"
	MachineDowntimeReasonSummaryCodeMinorStop          MachineDowntimeReasonSummaryCode = "minor_stop"
	MachineDowntimeReasonSummaryCodeQualityHold        MachineDowntimeReasonSummaryCode = "quality_hold"
	MachineDowntimeReasonSummaryCodeNoSchedule         MachineDowntimeReasonSummaryCode = "no_schedule"
)

// Resource type identifier.
type MachineDowntimeReasonSummaryObject string

const (
	MachineDowntimeReasonSummaryObjectMachineDowntimeReason MachineDowntimeReasonSummaryObject = "machine_downtime_reason"
)

// Which OEE term this reason charges.
type MachineDowntimeReasonSummaryOeeBucket string

const (
	MachineDowntimeReasonSummaryOeeBucketAvailability MachineDowntimeReasonSummaryOeeBucket = "availability"
	MachineDowntimeReasonSummaryOeeBucketPerformance  MachineDowntimeReasonSummaryOeeBucket = "performance"
	MachineDowntimeReasonSummaryOeeBucketQuality      MachineDowntimeReasonSummaryOeeBucket = "quality"
	MachineDowntimeReasonSummaryOeeBucketNotScheduled MachineDowntimeReasonSummaryOeeBucket = "not_scheduled"
)

// An open stoppage on a machine.
type MachineDowntimeSummary struct {
	// Entity is a polymorphic reference to any resource in the system.
	Event Entity `json:"event" api:"required"`
	// Free-text note left by whoever logged it.
	Note string `json:"note" api:"required"`
	// The reason for a stoppage, as carried on a downtime event.
	//
	// A denormalized view of the reason taxonomy: the stable code plus the display
	// name and OEE bucket resolved from it at read time.
	Reason MachineDowntimeReasonSummary `json:"reason" api:"required"`
	// When the machine went down.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Note        respjson.Field
		Reason      respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineDowntimeSummary) RawJSON() string { return r.JSON.raw }
func (r *MachineDowntimeSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What one machine is doing right now.
//
// Assembled from the published schedule, the batches the floor has scanned against
// it, and any open downtime. A machine with an open stoppage reads `down` even
// when it has a released campaign, because a broken machine is not producing
// whatever the plan says.
type MachineStatus struct {
	// One campaign on a machine, with how far through it the floor is.
	//
	// A campaign is one item scheduled to run on one machine for one week. Progress is
	// taken from the batches the floor has scanned against it rather than reported by
	// hand, so it advances on its own as a shift runs.
	Current MachineCampaign `json:"current" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Department Entity `json:"department" api:"required"`
	// An open stoppage on a machine.
	Downtime MachineDowntimeSummary `json:"downtime" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// One campaign on a machine, with how far through it the floor is.
	//
	// A campaign is one item scheduled to run on one machine for one week. Progress is
	// taken from the batches the floor has scanned against it rather than reported by
	// hand, so it advances on its own as a shift runs.
	Next MachineCampaign `json:"next" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine_status".
	Object MachineStatusObject `json:"object" api:"required"`
	// What the machine is doing.
	//
	// - `running`: a released campaign with work still to scan.
	// - `idle`: nothing released to it.
	// - `down`: an open downtime event, which outranks running.
	//
	// Any of "running", "idle", "down".
	Status MachineStatusStatus `json:"status" api:"required"`
	// Unit the week's quantities are counted in.
	Unit string `json:"unit" api:"required"`
	// Quantity planned on this machine for the current week.
	//
	// Summed across every campaign scheduled on the machine that week, not just the
	// current one.
	WeekPlannedQuantity float64 `json:"week_planned_quantity" api:"required"`
	// Machine hours the plan allocates on this machine for the current week.
	WeekPlannedRunHours float64 `json:"week_planned_run_hours" api:"required"`
	// Quantity scanned on this machine so far in the current week.
	WeekScannedQuantity float64 `json:"week_scanned_quantity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Current             respjson.Field
		Department          respjson.Field
		Downtime            respjson.Field
		Machine             respjson.Field
		Next                respjson.Field
		Object              respjson.Field
		Status              respjson.Field
		Unit                respjson.Field
		WeekPlannedQuantity respjson.Field
		WeekPlannedRunHours respjson.Field
		WeekScannedQuantity respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MachineStatus) RawJSON() string { return r.JSON.raw }
func (r *MachineStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MachineStatusObject string

const (
	MachineStatusObjectMachineStatus MachineStatusObject = "machine_status"
)

// What the machine is doing.
//
// - `running`: a released campaign with work still to scan.
// - `idle`: nothing released to it.
// - `down`: an open downtime event, which outranks running.
type MachineStatusStatus string

const (
	MachineStatusStatusRunning MachineStatusStatus = "running"
	MachineStatusStatusIdle    MachineStatusStatus = "idle"
	MachineStatusStatusDown    MachineStatusStatus = "down"
)

// A kind of hand change to a plan.
type ScheduleDeviationType struct {
	// Deviation type ID.
	ID string `json:"id" api:"required"`
	// Stable code recorded on a deviation.
	//
	// Any of "line_added", "line_removed", "quantity_changed", "machine_changed",
	// "resequenced", "week_moved".
	Code ScheduleDeviationTypeCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the type.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "schedule_deviation_type".
	Object ScheduleDeviationTypeObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleDeviationType) RawJSON() string { return r.JSON.raw }
func (r *ScheduleDeviationType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stable code recorded on a deviation.
type ScheduleDeviationTypeCode string

const (
	ScheduleDeviationTypeCodeLineAdded       ScheduleDeviationTypeCode = "line_added"
	ScheduleDeviationTypeCodeLineRemoved     ScheduleDeviationTypeCode = "line_removed"
	ScheduleDeviationTypeCodeQuantityChanged ScheduleDeviationTypeCode = "quantity_changed"
	ScheduleDeviationTypeCodeMachineChanged  ScheduleDeviationTypeCode = "machine_changed"
	ScheduleDeviationTypeCodeResequenced     ScheduleDeviationTypeCode = "resequenced"
	ScheduleDeviationTypeCodeWeekMoved       ScheduleDeviationTypeCode = "week_moved"
)

// Resource type identifier.
type ScheduleDeviationTypeObject string

const (
	ScheduleDeviationTypeObjectScheduleDeviationType ScheduleDeviationTypeObject = "schedule_deviation_type"
)

type OperationGetMachineStatusParams struct {
	// The moment to read the floor at.
	//
	// Chooses the week the campaigns are read for, and the published schedule whose
	// horizon covers that moment; open downtime and scan progress are always read as
	// they stand now. Omit it to read the floor as it is at this instant.
	AsOf param.Opt[time.Time] `query:"as_of,omitzero" format:"date-time" json:"-"`
	// Only include machines in these departments.
	DepartmentIDs []string `query:"department_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationGetMachineStatusParams]'s query parameters as
// `url.Values`.
func (r OperationGetMachineStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
