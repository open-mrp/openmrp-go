// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Generate and review machine-level production schedules.
//
// OperationProductionScheduleService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleService] method instead.
type OperationProductionScheduleService struct {
	options []option.RequestOption
	// Generate and review machine-level production schedules.
	Lines OperationProductionScheduleLineService
	// Generate and review machine-level production schedules.
	Actions OperationProductionScheduleActionService
}

// NewOperationProductionScheduleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationProductionScheduleService(opts ...option.RequestOption) (r OperationProductionScheduleService) {
	r = OperationProductionScheduleService{}
	r.options = opts
	r.Lines = NewOperationProductionScheduleLineService(opts...)
	r.Actions = NewOperationProductionScheduleActionService(opts...)
	return
}

// Generates and saves a new production schedule.
//
// The plan is saved as a draft: nothing is frozen yet, so campaigns can be added,
// changed and removed without having to give a reason. Generating again creates a
// new version rather than replacing this one, because attainment is measured
// against whichever version was live at the time.
//
// The solver plans the constraint department — the room that sets the pace of the
// factory — so production schedule settings must name one and it must have
// machines that are included in planning. Without that there is nothing to
// schedule and the request is rejected rather than returning an empty plan.
//
// Alongside the campaigns, the version stores the assumptions it was solved with,
// the per-item policies behind each campaign, and the downstream department work
// implied by the plan.
//
// This endpoint requires the permission: `production_schedules:create`.
func (r *OperationProductionScheduleService) New(ctx context.Context, body OperationProductionScheduleNewParams, opts ...option.RequestOption) (res *ProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single production schedule version.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a paginated list of production schedule versions, newest first.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) List(ctx context.Context, query OperationProductionScheduleListParams, opts ...option.RequestOption) (res *ListProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a draft schedule along with its planned campaigns and its item policy
// snapshot.
//
// Only drafts can be deleted. A published version is the baseline attainment is
// measured against, so removing it would erase the record of what was promised —
// archive those instead.
//
// This endpoint requires the permission: `production_schedules:delete`.
func (r *OperationProductionScheduleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationProductionScheduleDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the customer commitments this schedule version does not meet, soonest
// first.
//
// Three ways an order lands here. `past_due` means the constraint stage needed to
// start before this plan begins. `undated` means the order carries no ship-by
// commitment at all, so it is treated as owed now. `short` means the plan simply
// does not build enough of it in time — the campaigns it does allocate are listed
// alongside, because building three hundred of five hundred is a different
// conversation from building none.
//
// Read from the version's own record rather than re-solved, so what comes back is
// what was decided when the plan was made. A version generated before commitments
// were tracked reports nothing, which is correct: it made no promises it could
// break.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetAtRiskOrders(ctx context.Context, id string, opts ...option.RequestOption) (res *ListScheduleOrderCoverage, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/at-risk-orders", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the published schedule covering today.
//
// Responds 404 when no published version covers today, which is the normal state
// before the first schedule is published. Drafts are never returned here — a plan
// nobody has committed to is not the current plan.
//
// At most one version is ever current: publishing a new one supersedes every
// published version its horizon overlaps, so republishing mid-horizon takes over
// immediately.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetCurrent(ctx context.Context, opts ...option.RequestOption) (res *ProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedules/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the department work implied by a schedule's constraint plan.
//
// The solver schedules only the constraint; every other department's work is
// derived from it by walking the production-step graph, applying each step's
// lead-time offset and yield. That makes this the work list a supervisor reads,
// rather than a second plan someone has to maintain.
//
// `explosion_depth` is how many steps downstream the work sits, which is what a
// readiness indicator keys off. Depth 0 is the constraint's own campaigns, so a
// plant with nothing configured downstream of its constraint still gets the work
// it actually scheduled. Work whose derived week falls past the schedule's horizon
// is still returned — a department needs to see it coming.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetDerivedLines(ctx context.Context, id string, query OperationProductionScheduleGetDerivedLinesParams, opts ...option.RequestOption) (res *ListProductionScheduleDerivedLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/derived-lines", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the append-only log of hand changes made to a schedule, most recent
// first.
//
// This is what frozen-week adherence is measured from. A change recorded as frozen
// was inside the freeze window at the moment it was made, and stays that way
// regardless of what is published later.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetDeviations(ctx context.Context, id string, query OperationProductionScheduleGetDeviationsParams, opts ...option.RequestOption) (res *ListProductionScheduleDeviation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/deviations", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the per-finished-SKU inventory targets behind a schedule version,
// grouped under the constraint item each one is made from.
//
// The item policies pool every finished good a constraint item feeds into one
// echelon figure, which is what the build decision is made against. These rows are
// what that pooling hides: each finished SKU's own demand, its own variability,
// its own stock, and a buffer sized against the finishing lead time rather than
// the constraint's.
//
// The two stages do not overlap, so together they describe the whole network's
// stock without counting any of it twice: the constraint stage holds its pooled
// buffer, and the finished stage holds these.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetFinishedPolicies(ctx context.Context, id string, opts ...option.RequestOption) (res *ListProductionScheduleFinishedPolicy, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/finished-policies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the second stage of a schedule: how many of which finished good to make
// from the knitted parts, week by week.
//
// The constraint plan says how much greige to knit and deliberately does not say
// what to turn it into — a family's demand is pooled onto the greige precisely so
// the buffer can sit at the undifferentiated stage, where it is cheapest. These
// lines are where that pooling is undone, against each finished SKU's own stock
// position, its own orders, and the hours the rest of the factory has that week.
//
// Levelled, not merely allocated. Work that does not fit a week moves to the next
// one rather than being dropped, so the plan never asks the second stage for more
// hours than it has. Two things bound it, and they are reported separately in the
// schedule's diagnostics because they call for opposite responses: a SKU held back
// for want of greige is a knitting problem, and a SKU held back for want of hours
// is a finishing one.
//
// Everything is counted in the constraint item's unit, so `greige_consumed` here
// and `planned_quantity` on the constraint plan are directly comparable — which is
// what lets the two stages be reconciled rather than only read side by side.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetFinishingLines(ctx context.Context, id string, query OperationProductionScheduleGetFinishingLinesParams, opts ...option.RequestOption) (res *ListProductionScheduleFinishingLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/finishing-lines", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the per-item policy behind a schedule version, ordered by constraint run
// hours descending.
//
// This is the "why" behind every campaign: lot size, reorder point, safety stock
// and lead times as they stood when the plan was generated.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetItemPolicies(ctx context.Context, id string, opts ...option.RequestOption) (res *ListProductionScheduleItemPolicy, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/item-policies", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns what releasing a week would create, without creating it.
//
// The lots are resolved exactly as the release itself resolves them, so what a
// planner is shown and what the floor receives cannot drift apart.
//
// `is_releasable` is false when the week is empty or already released, with
// `blocked_reason` saying which; `existing_production_run_id` names the run a
// released week is already tied to.
//
// Cancelled campaigns and campaigns planned at zero are excluded here exactly as
// the release excludes them, so a week holding nothing but those previews as
// empty.
//
// Lots the floor is already holding are named as such. A batch with
// `carried_forward_from` set is a ticket an earlier week issued and nobody worked,
// which the release moves into the new run rather than reissuing, so nothing has
// to be reprinted.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleService) GetWeekReleasePreview(ctx context.Context, id string, query OperationProductionScheduleGetWeekReleasePreviewParams, opts ...option.RequestOption) (res *ReleaseScheduleWeekPreview, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/week-release-preview", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to generate a production schedule.
type GenerateProductionScheduleRequestParam struct {
	// Number of weeks the plan should cover, overriding the account's configured
	// horizon for this version only.
	HorizonWeeks param.Opt[int64] `json:"horizon_weeks,omitzero"`
	// Human-readable label for the version, such as the week it was cut for.
	//
	// Purely for recognising the version in a list; versions are numbered
	// automatically and the number is what identifies them.
	Name param.Opt[string] `json:"name,omitzero"`
	// The instant to plan against, which is what stock, demand history and active
	// demand overrides are read as of.
	//
	// Left unset, the plan is solved against the moment the request arrives. The
	// horizon starts on the account's configured week-start day on or before this
	// instant, so backdating this shifts the whole week grid.
	PlanningAsOf param.Opt[time.Time] `json:"planning_as_of,omitzero" format:"date-time"`
	// How future demand is derived, overriding the account's configured basis for this
	// version only.
	//
	//   - `trailing_12`: demand is the trailing twelve months of orders.
	//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
	//     a season arriving early or late rather than flattening it.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis GenerateProductionScheduleRequestDemandBasis `json:"demand_basis,omitzero"`
	paramObj
}

func (r GenerateProductionScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow GenerateProductionScheduleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerateProductionScheduleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How future demand is derived, overriding the account's configured basis for this
// version only.
//
//   - `trailing_12`: demand is the trailing twelve months of orders.
//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
//     a season arriving early or late rather than flattening it.
type GenerateProductionScheduleRequestDemandBasis string

const (
	GenerateProductionScheduleRequestDemandBasisTrailing12  GenerateProductionScheduleRequestDemandBasis = "trailing_12"
	GenerateProductionScheduleRequestDemandBasisSeasonalEma GenerateProductionScheduleRequestDemandBasis = "seasonal_ema"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionSchedule struct {
	// Resources in this page.
	Data []ProductionSchedule `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleObject `json:"object" api:"required"`
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
func (r ListProductionSchedule) RawJSON() string { return r.JSON.raw }
func (r *ListProductionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleObject string

const (
	ListProductionScheduleObjectList ListProductionScheduleObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionScheduleDerivedLine struct {
	// Resources in this page.
	Data []ProductionScheduleDerivedLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleDerivedLineObject `json:"object" api:"required"`
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
func (r ListProductionScheduleDerivedLine) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleDerivedLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleDerivedLineObject string

const (
	ListProductionScheduleDerivedLineObjectList ListProductionScheduleDerivedLineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionScheduleDeviation struct {
	// Resources in this page.
	Data []ProductionScheduleDeviation `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleDeviationObject `json:"object" api:"required"`
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
func (r ListProductionScheduleDeviation) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleDeviationObject string

const (
	ListProductionScheduleDeviationObjectList ListProductionScheduleDeviationObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionScheduleFinishedPolicy struct {
	// Resources in this page.
	Data []ProductionScheduleFinishedPolicy `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleFinishedPolicyObject `json:"object" api:"required"`
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
func (r ListProductionScheduleFinishedPolicy) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleFinishedPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleFinishedPolicyObject string

const (
	ListProductionScheduleFinishedPolicyObjectList ListProductionScheduleFinishedPolicyObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionScheduleFinishingLine struct {
	// Resources in this page.
	Data []ProductionScheduleFinishingLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleFinishingLineObject `json:"object" api:"required"`
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
func (r ListProductionScheduleFinishingLine) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleFinishingLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleFinishingLineObject string

const (
	ListProductionScheduleFinishingLineObjectList ListProductionScheduleFinishingLineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionScheduleItemPolicy struct {
	// Resources in this page.
	Data []ProductionScheduleItemPolicy `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionScheduleItemPolicyObject `json:"object" api:"required"`
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
func (r ListProductionScheduleItemPolicy) RawJSON() string { return r.JSON.raw }
func (r *ListProductionScheduleItemPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionScheduleItemPolicyObject string

const (
	ListProductionScheduleItemPolicyObjectList ListProductionScheduleItemPolicyObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListReleaseScheduleBatch struct {
	// Resources in this page.
	Data []ReleaseScheduleBatch `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListReleaseScheduleBatchObject `json:"object" api:"required"`
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
func (r ListReleaseScheduleBatch) RawJSON() string { return r.JSON.raw }
func (r *ListReleaseScheduleBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListReleaseScheduleBatchObject string

const (
	ListReleaseScheduleBatchObjectList ListReleaseScheduleBatchObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListReleasedScheduleLine struct {
	// Resources in this page.
	Data []ReleasedScheduleLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListReleasedScheduleLineObject `json:"object" api:"required"`
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
func (r ListReleasedScheduleLine) RawJSON() string { return r.JSON.raw }
func (r *ListReleasedScheduleLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListReleasedScheduleLineObject string

const (
	ListReleasedScheduleLineObjectList ListReleasedScheduleLineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleAppliedOverride struct {
	// Resources in this page.
	Data []ScheduleAppliedOverride `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleAppliedOverrideObject `json:"object" api:"required"`
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
func (r ListScheduleAppliedOverride) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleAppliedOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleAppliedOverrideObject string

const (
	ListScheduleAppliedOverrideObjectList ListScheduleAppliedOverrideObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleAtRiskOrder struct {
	// Resources in this page.
	Data []ScheduleAtRiskOrder `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleAtRiskOrderObject `json:"object" api:"required"`
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
func (r ListScheduleAtRiskOrder) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleAtRiskOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleAtRiskOrderObject string

const (
	ListScheduleAtRiskOrderObjectList ListScheduleAtRiskOrderObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleOrderCoverage struct {
	// Resources in this page.
	Data []ScheduleOrderCoverage `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleOrderCoverageObject `json:"object" api:"required"`
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
func (r ListScheduleOrderCoverage) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleOrderCoverage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleOrderCoverageObject string

const (
	ListScheduleOrderCoverageObjectList ListScheduleOrderCoverageObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleOrderCoverageLine struct {
	// Resources in this page.
	Data []ScheduleOrderCoverageLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleOrderCoverageLineObject `json:"object" api:"required"`
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
func (r ListScheduleOrderCoverageLine) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleOrderCoverageLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleOrderCoverageLineObject string

const (
	ListScheduleOrderCoverageLineObjectList ListScheduleOrderCoverageLineObject = "list"
)

// A saved production schedule.
//
// A published version is a record rather than a document that keeps being edited:
// generating again creates a new version, and publishing supersedes the previous
// one rather than changing it, because attainment is measured against whichever
// version was live at the time.
type ProductionSchedule struct {
	// Schedule ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Which demand basis produced the plan.
	//
	//   - `trailing_12`: demand is taken from the trailing twelve months of orders.
	//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
	//     a season arriving earlier or later than usual.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis ProductionScheduleDemandBasis `json:"demand_basis" api:"required"`
	// What the solver could not do, and why the plan differs from raw history.
	Diagnostics ScheduleDiagnostics `json:"diagnostics" api:"required"`
	// Why generation failed, when it did.
	ErrorMessage string `json:"error_message" api:"required"`
	// Number of lines that were frozen at publish.
	//
	// Captured once and never recomputed, because frozen-week adherence measures
	// against what was committed to.
	FrozenLineCount int64 `json:"frozen_line_count" api:"required"`
	// Total quantity frozen at publish.
	FrozenPlannedQuantity float64 `json:"frozen_planned_quantity" api:"required"`
	// The last day the frozen window covers, set when the version is published.
	FrozenThroughAt time.Time `json:"frozen_through_at" api:"required" format:"date-time"`
	// How many leading weeks freeze on publish.
	//
	// Publishing freezes every campaign that starts inside the window; changing one
	// afterwards requires a reason and is recorded in the deviation log.
	FrozenWeeks int64 `json:"frozen_weeks" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	GeneratedBy Actor `json:"generated_by" api:"required"`
	// What triggered the generation.
	//
	// - `manual`: someone asked for this version.
	// - `scheduled`: the account's generation cadence produced it on its own.
	//
	// Any of "manual", "scheduled".
	GenerationSource ProductionScheduleGenerationSource `json:"generation_source" api:"required"`
	// First instant of the last day of the horizon.
	HorizonEndsAt time.Time `json:"horizon_ends_at" api:"required" format:"date-time"`
	// First instant of the horizon.
	HorizonStartsAt time.Time `json:"horizon_starts_at" api:"required" format:"date-time"`
	// Length of the horizon in weeks.
	HorizonWeeks int64 `json:"horizon_weeks" api:"required"`
	// Label for the version, such as the planning cycle it was generated for.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule".
	Object ProductionScheduleObject `json:"object" api:"required"`
	// The instant the plan was calculated against.
	PlanningAsOfAt time.Time `json:"planning_as_of_at" api:"required" format:"date-time"`
	// When this version was published.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	PublishedBy Actor `json:"published_by" api:"required"`
	// The planning assumptions used, frozen at generation so the plan stays
	// explainable after settings change.
	SettingsSnapshot map[string]any `json:"settings_snapshot" api:"required"`
	// Version of the solver that produced the plan.
	SolverVersion string `json:"solver_version" api:"required"`
	// Where this version is in its lifecycle.
	//
	// - `draft`: still editable and commits to nothing.
	// - `generating`: a scheduled solve is still building this version.
	// - `published`: live, with its leading weeks frozen as a commitment to the floor.
	// - `superseded`: a later version was published over an overlapping horizon.
	// - `archived`: retired without being replaced.
	// - `failed`: the solver could not produce a plan; `error_message` says why.
	//
	// Any of "draft", "generating", "published", "superseded", "archived", "failed".
	Status ProductionScheduleStatus `json:"status" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	SupersededBy Entity `json:"superseded_by" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Sequential version number within the account.
	//
	// Regenerating a draft re-solves it in place and keeps its number; only generating
	// a new plan takes the next one.
	Version int64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		DemandBasis           respjson.Field
		Diagnostics           respjson.Field
		ErrorMessage          respjson.Field
		FrozenLineCount       respjson.Field
		FrozenPlannedQuantity respjson.Field
		FrozenThroughAt       respjson.Field
		FrozenWeeks           respjson.Field
		GeneratedBy           respjson.Field
		GenerationSource      respjson.Field
		HorizonEndsAt         respjson.Field
		HorizonStartsAt       respjson.Field
		HorizonWeeks          respjson.Field
		Name                  respjson.Field
		Object                respjson.Field
		PlanningAsOfAt        respjson.Field
		PublishedAt           respjson.Field
		PublishedBy           respjson.Field
		SettingsSnapshot      respjson.Field
		SolverVersion         respjson.Field
		Status                respjson.Field
		SupersededBy          respjson.Field
		UpdatedAt             respjson.Field
		Version               respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionSchedule) RawJSON() string { return r.JSON.raw }
func (r *ProductionSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which demand basis produced the plan.
//
//   - `trailing_12`: demand is taken from the trailing twelve months of orders.
//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
//     a season arriving earlier or later than usual.
type ProductionScheduleDemandBasis string

const (
	ProductionScheduleDemandBasisTrailing12  ProductionScheduleDemandBasis = "trailing_12"
	ProductionScheduleDemandBasisSeasonalEma ProductionScheduleDemandBasis = "seasonal_ema"
)

// What triggered the generation.
//
// - `manual`: someone asked for this version.
// - `scheduled`: the account's generation cadence produced it on its own.
type ProductionScheduleGenerationSource string

const (
	ProductionScheduleGenerationSourceManual    ProductionScheduleGenerationSource = "manual"
	ProductionScheduleGenerationSourceScheduled ProductionScheduleGenerationSource = "scheduled"
)

// Resource type identifier.
type ProductionScheduleObject string

const (
	ProductionScheduleObjectProductionSchedule ProductionScheduleObject = "production_schedule"
)

// Where this version is in its lifecycle.
//
// - `draft`: still editable and commits to nothing.
// - `generating`: a scheduled solve is still building this version.
// - `published`: live, with its leading weeks frozen as a commitment to the floor.
// - `superseded`: a later version was published over an overlapping horizon.
// - `archived`: retired without being replaced.
// - `failed`: the solver could not produce a plan; `error_message` says why.
type ProductionScheduleStatus string

const (
	ProductionScheduleStatusDraft      ProductionScheduleStatus = "draft"
	ProductionScheduleStatusGenerating ProductionScheduleStatus = "generating"
	ProductionScheduleStatusPublished  ProductionScheduleStatus = "published"
	ProductionScheduleStatusSuperseded ProductionScheduleStatus = "superseded"
	ProductionScheduleStatusArchived   ProductionScheduleStatus = "archived"
	ProductionScheduleStatusFailed     ProductionScheduleStatus = "failed"
)

// Downstream department work implied by a constraint campaign.
//
// The solver only schedules the constraint; every other department's work follows
// from it by walking the production-step graph. `explosion_depth` is how many
// steps downstream this sits — depth 1 waits only on the constraint, depth 3 waits
// on two intermediate steps — which is what a readiness indicator keys off.
//
// The derived week can fall past the schedule's horizon when a long chain follows
// a late campaign. That work is still returned rather than dropped, because a
// department needs to see it coming.
type ProductionScheduleDerivedLine struct {
	// Derived line ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Department Entity `json:"department" api:"required"`
	// How many steps downstream of the constraint this work sits.
	ExplosionDepth int64 `json:"explosion_depth" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_derived_line".
	Object ProductionScheduleDerivedLineObject `json:"object" api:"required"`
	// Weeks after the constraint campaign this work starts.
	OffsetWeeks int64 `json:"offset_weeks" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	PlannedUnit Entity `json:"planned_unit" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionStep Entity `json:"production_step" api:"required"`
	// Units implied for this step.
	Quantity float64 `json:"quantity" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	SourceLine Entity `json:"source_line" api:"required"`
	// State of the derived work.
	//
	// Derived rows are discarded and rebuilt from the constraint plan every time the
	// version is solved, and are only ever written as `planned`, so they report what
	// the plan implies rather than what the floor has done.
	//
	// Any of "planned", "released", "in_progress", "complete", "cancelled".
	Status ProductionScheduleDerivedLineStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Horizon week the work falls in, zero-based.
	WeekIndex int64 `json:"week_index" api:"required"`
	// First instant of that week.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Department         respjson.Field
		ExplosionDepth     respjson.Field
		Item               respjson.Field
		Object             respjson.Field
		OffsetWeeks        respjson.Field
		PlannedUnit        respjson.Field
		ProductionSchedule respjson.Field
		ProductionStep     respjson.Field
		Quantity           respjson.Field
		SourceLine         respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		WeekIndex          respjson.Field
		WeekStartsAt       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleDerivedLine) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleDerivedLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionScheduleDerivedLineObject string

const (
	ProductionScheduleDerivedLineObjectProductionScheduleDerivedLine ProductionScheduleDerivedLineObject = "production_schedule_derived_line"
)

// State of the derived work.
//
// Derived rows are discarded and rebuilt from the constraint plan every time the
// version is solved, and are only ever written as `planned`, so they report what
// the plan implies rather than what the floor has done.
type ProductionScheduleDerivedLineStatus string

const (
	ProductionScheduleDerivedLineStatusPlanned    ProductionScheduleDerivedLineStatus = "planned"
	ProductionScheduleDerivedLineStatusReleased   ProductionScheduleDerivedLineStatus = "released"
	ProductionScheduleDerivedLineStatusInProgress ProductionScheduleDerivedLineStatus = "in_progress"
	ProductionScheduleDerivedLineStatusComplete   ProductionScheduleDerivedLineStatus = "complete"
	ProductionScheduleDerivedLineStatusCancelled  ProductionScheduleDerivedLineStatus = "cancelled"
)

// One hand change to a production schedule.
//
// The log is append-only: it is what frozen-week adherence is measured from, and a
// plan edited back into shape has to stay distinguishable from one that was right
// the first time. `before` and `after` are full snapshots of the line, so a
// deviation stays readable after the line it describes is deleted.
//
// `freeze_status` is recorded when the change is made, from the freeze window as
// it stood at that moment. It is never re-derived, so a later publish cannot
// retroactively reclassify a past edit.
type ProductionScheduleDeviation struct {
	// Deviation ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Snapshot of the line after the change, null when the change removed it. Encoded
	// as a JSON value (object, array, string, number, boolean, or null), not a
	// JSON-encoded string.
	After any `json:"after" api:"required"`
	// Snapshot of the line before the change, null when the change created it. Encoded
	// as a JSON value (object, array, string, number, boolean, or null), not a
	// JSON-encoded string.
	Before any `json:"before" api:"required"`
	// When the change was made.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Signed change in planned units.
	DeltaQuantity float64 `json:"delta_quantity" api:"required"`
	// Signed change in planned run hours.
	DeltaRunHours float64 `json:"delta_run_hours" api:"required"`
	// What kind of change this was.
	//
	// Derived from the change itself rather than supplied by the person making it. An
	// edit that both moves a campaign to another machine and changes its quantity is
	// recorded as the machine change, because that is what a planner has to react to
	// first.
	//
	// Any of "line_added", "line_removed", "quantity_changed", "machine_changed",
	// "resequenced", "week_moved".
	DeviationType ProductionScheduleDeviationDeviationType `json:"deviation_type" api:"required"`
	// Whether the change fell inside the frozen window when it was made.
	//
	// Any of "frozen", "flexible".
	FreezeStatus ProductionScheduleDeviationFreezeStatus `json:"freeze_status" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Line Entity `json:"line" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_deviation".
	Object ProductionScheduleDeviationObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Why the change was made.
	//
	// A change inside a frozen week has to supply one; outside it a reason is left to
	// the planner.
	//
	//   - `machine_down`: the machine the campaign was on stopped running.
	//   - `material_shortage`: the material the campaign needs did not arrive.
	//   - `rush_order`: demand that could not wait for the next plan.
	//   - `quality_hold`: the work was stopped over a quality problem.
	//   - `over_run`: the floor produced more than the plan asked for.
	//   - `under_run`: the floor produced less than the plan asked for.
	//   - `capacity_change`: the available machine time changed, such as a shutdown or
	//     an added shift.
	//   - `other`: something outside the list, which should be spelled out in
	//     `reason_note`.
	//
	// Any of "machine_down", "material_shortage", "rush_order", "quality_hold",
	// "over_run", "under_run", "capacity_change", "other".
	Reason ProductionScheduleDeviationReason `json:"reason" api:"required"`
	// Free-form explanation of the change.
	ReasonNote string `json:"reason_note" api:"required"`
	// The horizon week the change affected, zero-based.
	WeekIndex int64 `json:"week_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Actor              respjson.Field
		After              respjson.Field
		Before             respjson.Field
		CreatedAt          respjson.Field
		DeltaQuantity      respjson.Field
		DeltaRunHours      respjson.Field
		DeviationType      respjson.Field
		FreezeStatus       respjson.Field
		Item               respjson.Field
		Line               respjson.Field
		Machine            respjson.Field
		Object             respjson.Field
		ProductionSchedule respjson.Field
		Reason             respjson.Field
		ReasonNote         respjson.Field
		WeekIndex          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleDeviation) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleDeviation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What kind of change this was.
//
// Derived from the change itself rather than supplied by the person making it. An
// edit that both moves a campaign to another machine and changes its quantity is
// recorded as the machine change, because that is what a planner has to react to
// first.
type ProductionScheduleDeviationDeviationType string

const (
	ProductionScheduleDeviationDeviationTypeLineAdded       ProductionScheduleDeviationDeviationType = "line_added"
	ProductionScheduleDeviationDeviationTypeLineRemoved     ProductionScheduleDeviationDeviationType = "line_removed"
	ProductionScheduleDeviationDeviationTypeQuantityChanged ProductionScheduleDeviationDeviationType = "quantity_changed"
	ProductionScheduleDeviationDeviationTypeMachineChanged  ProductionScheduleDeviationDeviationType = "machine_changed"
	ProductionScheduleDeviationDeviationTypeResequenced     ProductionScheduleDeviationDeviationType = "resequenced"
	ProductionScheduleDeviationDeviationTypeWeekMoved       ProductionScheduleDeviationDeviationType = "week_moved"
)

// Whether the change fell inside the frozen window when it was made.
type ProductionScheduleDeviationFreezeStatus string

const (
	ProductionScheduleDeviationFreezeStatusFrozen   ProductionScheduleDeviationFreezeStatus = "frozen"
	ProductionScheduleDeviationFreezeStatusFlexible ProductionScheduleDeviationFreezeStatus = "flexible"
)

// Resource type identifier.
type ProductionScheduleDeviationObject string

const (
	ProductionScheduleDeviationObjectProductionScheduleDeviation ProductionScheduleDeviationObject = "production_schedule_deviation"
)

// Why the change was made.
//
// A change inside a frozen week has to supply one; outside it a reason is left to
// the planner.
//
//   - `machine_down`: the machine the campaign was on stopped running.
//   - `material_shortage`: the material the campaign needs did not arrive.
//   - `rush_order`: demand that could not wait for the next plan.
//   - `quality_hold`: the work was stopped over a quality problem.
//   - `over_run`: the floor produced more than the plan asked for.
//   - `under_run`: the floor produced less than the plan asked for.
//   - `capacity_change`: the available machine time changed, such as a shutdown or
//     an added shift.
//   - `other`: something outside the list, which should be spelled out in
//     `reason_note`.
type ProductionScheduleDeviationReason string

const (
	ProductionScheduleDeviationReasonMachineDown      ProductionScheduleDeviationReason = "machine_down"
	ProductionScheduleDeviationReasonMaterialShortage ProductionScheduleDeviationReason = "material_shortage"
	ProductionScheduleDeviationReasonRushOrder        ProductionScheduleDeviationReason = "rush_order"
	ProductionScheduleDeviationReasonQualityHold      ProductionScheduleDeviationReason = "quality_hold"
	ProductionScheduleDeviationReasonOverRun          ProductionScheduleDeviationReason = "over_run"
	ProductionScheduleDeviationReasonUnderRun         ProductionScheduleDeviationReason = "under_run"
	ProductionScheduleDeviationReasonCapacityChange   ProductionScheduleDeviationReason = "capacity_change"
	ProductionScheduleDeviationReasonOther            ProductionScheduleDeviationReason = "other"
)

// One finished SKU's own inventory target, snapshotted onto a schedule version.
//
// The item policy pools every finished good a constraint item feeds into one
// echelon figure, which is the right basis for deciding whether to build. These
// rows are what that pooling hides: this SKU's own demand, its own variability,
// and a buffer sized against the finishing lead time rather than the constraint's
// — because finishing, not the constraint, is what replenishes this stock.
//
// The two stages do not overlap. The constraint stage holds its pooled buffer and
// the finished stage holds these, so together they describe the whole network's
// stock without counting any of it twice.
type ProductionScheduleFinishedPolicy struct {
	// Finished policy ID.
	ID string `json:"id" api:"required"`
	// This SKU's own annual demand.
	AnnualDemand float64 `json:"annual_demand" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	GreigeItem Entity `json:"greige_item" api:"required"`
	// SKU of that constraint item.
	GreigeSKU string `json:"greige_sku" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_finished_policy".
	Object ProductionScheduleFinishedPolicyObject `json:"object" api:"required"`
	// This SKU's own stock, not the echelon it contributes to.
	OnHand float64 `json:"on_hand" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductLine Entity `json:"product_line" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Stock position at which this finished good needs replenishing.
	ReorderPoint float64 `json:"reorder_point" api:"required"`
	// Buffer held as this finished good, covering the finishing lead time.
	SafetyStock float64 `json:"safety_stock" api:"required"`
	// This SKU's own weekly demand variability.
	//
	// The constraint buffer pools these as the root of the sum of squares; these
	// targets use them one at a time.
	SigmaWeekly float64 `json:"sigma_weekly" api:"required"`
	// SKU of the finished good.
	SKU string `json:"sku" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// This SKU's own weekly demand.
	WeeklyDemand float64 `json:"weekly_demand" api:"required"`
	// Weeks of demand this SKU's own stock covers.
	WeeksOfCover float64 `json:"weeks_of_cover" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AnnualDemand       respjson.Field
		CreatedAt          respjson.Field
		GreigeItem         respjson.Field
		GreigeSKU          respjson.Field
		Item               respjson.Field
		Object             respjson.Field
		OnHand             respjson.Field
		ProductLine        respjson.Field
		ProductionSchedule respjson.Field
		ReorderPoint       respjson.Field
		SafetyStock        respjson.Field
		SigmaWeekly        respjson.Field
		SKU                respjson.Field
		UpdatedAt          respjson.Field
		WeeklyDemand       respjson.Field
		WeeksOfCover       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleFinishedPolicy) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleFinishedPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionScheduleFinishedPolicyObject string

const (
	ProductionScheduleFinishedPolicyObjectProductionScheduleFinishedPolicy ProductionScheduleFinishedPolicyObject = "production_schedule_finished_policy"
)

// One finished good's build in one week: the second stage of the plan.
//
// The constraint plan says how much greige to knit and deliberately does not say
// what to turn it into — a family's demand is pooled onto the greige precisely so
// the buffer can sit at the undifferentiated stage. These lines are where that
// pooling is undone: how many of which finished good to make from the knitted
// parts, decided against each SKU's own stock position, its own orders, and the
// hours the rest of the factory has that week.
//
// Quantities are counted in the constraint item's unit, so `greige_consumed` and
// the knit plan's `planned_quantity` are directly comparable. That is what lets
// the two stages be reconciled rather than merely read side by side.
type ProductionScheduleFinishingLine struct {
	// Finishing line ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Department Entity `json:"department" api:"required"`
	// How much of the week's draw on this SKU is an order rather than a forecast.
	FirmUnits float64 `json:"firm_units" api:"required"`
	// Units of the constraint item this takes out of the greige buffer.
	//
	// Equal to `planned_quantity` unless a finishing yield loss means a finished unit
	// costs more than one knitted one.
	GreigeConsumed float64 `json:"greige_consumed" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	GreigeItem Entity `json:"greige_item" api:"required"`
	// SKU of that constraint item.
	GreigeSKU string `json:"greige_sku" api:"required"`
	// Whether the line sits inside the published frozen window.
	IsFrozen bool `json:"is_frozen" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_finishing_line".
	Object ProductionScheduleFinishingLineObject `json:"object" api:"required"`
	// Units in one lot.
	PlannedLotUnits float64 `json:"planned_lot_units" api:"required"`
	// How many lots the quantity breaks into.
	PlannedLots int64 `json:"planned_lots" api:"required"`
	// Units of the finished good to make.
	PlannedQuantity float64 `json:"planned_quantity" api:"required"`
	// Hours of the second stage's capacity this line consumes.
	PlannedRunHours float64 `json:"planned_run_hours" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionStep Entity `json:"production_step" api:"required"`
	// And after it lands.
	ProjectedOnHandAfter float64 `json:"projected_on_hand_after" api:"required"`
	// This SKU's own projected stock before the line lands.
	ProjectedOnHandBefore float64 `json:"projected_on_hand_before" api:"required"`
	// SKU of the finished good, as it stood when the plan was generated.
	SKU string `json:"sku" api:"required"`
	// Whether the solver produced this line or a person did.
	//
	// Any of "solver", "manual".
	Source ProductionScheduleFinishingLineSource `json:"source" api:"required"`
	// Where the line stands.
	//
	// Any of "planned", "released", "in_progress", "complete", "cancelled".
	Status ProductionScheduleFinishingLineStatus `json:"status" api:"required"`
	// Abbreviation of the unit everything on this line is counted in.
	Unit string `json:"unit" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// First day of the week this is planned in.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Department            respjson.Field
		FirmUnits             respjson.Field
		GreigeConsumed        respjson.Field
		GreigeItem            respjson.Field
		GreigeSKU             respjson.Field
		IsFrozen              respjson.Field
		Item                  respjson.Field
		Object                respjson.Field
		PlannedLotUnits       respjson.Field
		PlannedLots           respjson.Field
		PlannedQuantity       respjson.Field
		PlannedRunHours       respjson.Field
		ProductionSchedule    respjson.Field
		ProductionStep        respjson.Field
		ProjectedOnHandAfter  respjson.Field
		ProjectedOnHandBefore respjson.Field
		SKU                   respjson.Field
		Source                respjson.Field
		Status                respjson.Field
		Unit                  respjson.Field
		UpdatedAt             respjson.Field
		WeekIndex             respjson.Field
		WeekStartsAt          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleFinishingLine) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleFinishingLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionScheduleFinishingLineObject string

const (
	ProductionScheduleFinishingLineObjectProductionScheduleFinishingLine ProductionScheduleFinishingLineObject = "production_schedule_finishing_line"
)

// Whether the solver produced this line or a person did.
type ProductionScheduleFinishingLineSource string

const (
	ProductionScheduleFinishingLineSourceSolver ProductionScheduleFinishingLineSource = "solver"
	ProductionScheduleFinishingLineSourceManual ProductionScheduleFinishingLineSource = "manual"
)

// Where the line stands.
type ProductionScheduleFinishingLineStatus string

const (
	ProductionScheduleFinishingLineStatusPlanned    ProductionScheduleFinishingLineStatus = "planned"
	ProductionScheduleFinishingLineStatusReleased   ProductionScheduleFinishingLineStatus = "released"
	ProductionScheduleFinishingLineStatusInProgress ProductionScheduleFinishingLineStatus = "in_progress"
	ProductionScheduleFinishingLineStatusComplete   ProductionScheduleFinishingLineStatus = "complete"
	ProductionScheduleFinishingLineStatusCancelled  ProductionScheduleFinishingLineStatus = "cancelled"
)

// The per-item policy behind a schedule version.
//
// Snapshotted at generation rather than recomputed, so a historical plan can still
// explain itself after costs, demand or settings move.
type ProductionScheduleItemPolicy struct {
	// Policy ID.
	ID string `json:"id" api:"required"`
	// ABC class by share of constraint run hours.
	//
	// - `a`: consumes the largest share of constraint capacity.
	// - `b`: moderate constraint consumption.
	// - `c`: consumes little constraint capacity.
	//
	// Any of "a", "b", "c".
	AbcClass ProductionScheduleItemPolicyAbcClass `json:"abc_class" api:"required"`
	// Demand used for planning, annualized.
	AnnualDemand float64 `json:"annual_demand" api:"required"`
	// Constraint hours this item's annual demand consumes.
	AnnualRunHours float64 `json:"annual_run_hours" api:"required"`
	// What the constraint stage holds on average: its buffer, plus half a campaign as
	// one lands and drains.
	AverageGreigeInventory float64 `json:"average_greige_inventory" api:"required"`
	// Observed or default lead time at the constraint.
	ConstraintLeadTimeWeeks float64 `json:"constraint_lead_time_weeks" api:"required"`
	// Limits the solver hit while sizing this item's campaigns, empty when the policy
	// was applied as calculated.
	//
	//   - `eoq_capped`: the economic lot size did not fit one machine-week and was cut
	//     back to what does, so campaigns run shorter and more often than the cost
	//     calculation alone would ask for.
	//   - `capacity_starved`: the item was already below its trigger point and never won
	//     a slot in the horizon, so the plan does not replenish it.
	//
	// Any of "eoq_capped", "capacity_starved".
	Constraints []string `json:"constraints" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Economic order quantity: the campaign size that balances the cost of a
	// changeover against the cost of holding what it produces.
	EoqUnits float64 `json:"eoq_units" api:"required"`
	// Lead time from the constraint to sellable stock.
	FinishLeadTimeWeeks float64 `json:"finish_lead_time_weeks" api:"required"`
	// Outstanding quantity the order book already owed for this item over the horizon.
	FirmDemandUnits float64 `json:"firm_demand_units" api:"required"`
	// Quantity the forecast projected for the same window.
	ForecastDemandUnits float64 `json:"forecast_demand_units" api:"required"`
	// How this item was planned.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer, so its safety stocks and reorder point are all zero.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy ProductionScheduleItemPolicyFulfillmentPolicy `json:"fulfillment_policy" api:"required"`
	// Annual cost of holding one unit.
	HoldingCost float64 `json:"holding_cost" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// What the constraint stage holds at its peak: its buffer plus a whole campaign.
	MaxGreigeInventory float64 `json:"max_greige_inventory" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_item_policy".
	Object ProductionScheduleItemPolicyObject `json:"object" api:"required"`
	// Stock at the constraint plus everything downstream of it.
	//
	// This is what the build decision is made against — stock already finished still
	// counts against building more.
	OnHandEchelon float64 `json:"on_hand_echelon" api:"required"`
	// Stock sitting at the constraint stage on its own.
	//
	// Kept alongside the echelon total because that total cannot be decomposed back
	// into its stages once summed.
	OnHandGreige float64 `json:"on_hand_greige" api:"required"`
	// Ceiling on how far ahead this item is built.
	OrderUpTo float64 `json:"order_up_to" api:"required"`
	// Which rule decided that policy: the item itself, its product line, or the
	// account default.
	//
	// Any of "item", "product_line", "account_default".
	PolicySource ProductionScheduleItemPolicyPolicySource `json:"policy_source" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	PrimaryMachine Entity `json:"primary_machine" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionStep Entity `json:"production_step" api:"required"`
	// The physical greige store at the end of each horizon week — the constraint stage
	// on its own, which `projected_on_hand` cannot be decomposed back into.
	//
	// A week where this dips to `safety_stock_primary` is the week knitting is meant
	// to replenish, even where `projected_on_hand` still reads full because the stock
	// is held downstream as finished goods. Empty for a schedule generated before the
	// greige buffer existed.
	ProjectedGreigeOnHand []float64 `json:"projected_greige_on_hand" api:"required"`
	// The echelon position at the end of each horizon week, after that week's
	// campaigns land and its demand is drawn down.
	//
	// A run of weeks with no campaign is stock draining toward `reorder_point`; this
	// is what makes that visible rather than looking like the solver did nothing.
	ProjectedOnHand []float64 `json:"projected_on_hand" api:"required"`
	// Stock position at which a campaign is triggered.
	ReorderPoint float64 `json:"reorder_point" api:"required"`
	// Buffer held as finished goods.
	SafetyStockDownstream float64 `json:"safety_stock_downstream" api:"required"`
	// Buffer held at the constraint.
	SafetyStockPrimary float64 `json:"safety_stock_primary" api:"required"`
	// How long one unit occupies the constraint.
	SecondsPerUnit float64 `json:"seconds_per_unit" api:"required"`
	// Cost of one changeover.
	SetupCost float64 `json:"setup_cost" api:"required"`
	// Summed weekly variability of the finished goods this item becomes.
	SigmaDownstreamSum float64 `json:"sigma_downstream_sum" api:"required"`
	// Pooled weekly demand variability at the constraint.
	SigmaWeeklyPooled float64 `json:"sigma_weekly_pooled" api:"required"`
	// SKU of the item.
	SKU string `json:"sku" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Unit Entity `json:"unit" api:"required"`
	// Abbreviation of the unit every quantity in this policy is counted in, for
	// display.
	//
	// A reorder point of 2,508 is uninterpretable without it, so the two are never
	// meaningful apart.
	UnitAbbreviation string `json:"unit_abbreviation" api:"required"`
	// Standard cost per unit.
	UnitCost float64 `json:"unit_cost" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Demand used for planning, per week.
	WeeklyDemand float64 `json:"weekly_demand" api:"required"`
	// Weeks of demand the current stock covers.
	WeeksOfCover float64 `json:"weeks_of_cover" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		AbcClass                respjson.Field
		AnnualDemand            respjson.Field
		AnnualRunHours          respjson.Field
		AverageGreigeInventory  respjson.Field
		ConstraintLeadTimeWeeks respjson.Field
		Constraints             respjson.Field
		CreatedAt               respjson.Field
		EoqUnits                respjson.Field
		FinishLeadTimeWeeks     respjson.Field
		FirmDemandUnits         respjson.Field
		ForecastDemandUnits     respjson.Field
		FulfillmentPolicy       respjson.Field
		HoldingCost             respjson.Field
		Item                    respjson.Field
		MaxGreigeInventory      respjson.Field
		Object                  respjson.Field
		OnHandEchelon           respjson.Field
		OnHandGreige            respjson.Field
		OrderUpTo               respjson.Field
		PolicySource            respjson.Field
		PrimaryMachine          respjson.Field
		ProductionSchedule      respjson.Field
		ProductionStep          respjson.Field
		ProjectedGreigeOnHand   respjson.Field
		ProjectedOnHand         respjson.Field
		ReorderPoint            respjson.Field
		SafetyStockDownstream   respjson.Field
		SafetyStockPrimary      respjson.Field
		SecondsPerUnit          respjson.Field
		SetupCost               respjson.Field
		SigmaDownstreamSum      respjson.Field
		SigmaWeeklyPooled       respjson.Field
		SKU                     respjson.Field
		Unit                    respjson.Field
		UnitAbbreviation        respjson.Field
		UnitCost                respjson.Field
		UpdatedAt               respjson.Field
		WeeklyDemand            respjson.Field
		WeeksOfCover            respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleItemPolicy) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleItemPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ABC class by share of constraint run hours.
//
// - `a`: consumes the largest share of constraint capacity.
// - `b`: moderate constraint consumption.
// - `c`: consumes little constraint capacity.
type ProductionScheduleItemPolicyAbcClass string

const (
	ProductionScheduleItemPolicyAbcClassA ProductionScheduleItemPolicyAbcClass = "a"
	ProductionScheduleItemPolicyAbcClassB ProductionScheduleItemPolicyAbcClass = "b"
	ProductionScheduleItemPolicyAbcClassC ProductionScheduleItemPolicyAbcClass = "c"
)

// How this item was planned.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer, so its safety stocks and reorder point are all zero.
type ProductionScheduleItemPolicyFulfillmentPolicy string

const (
	ProductionScheduleItemPolicyFulfillmentPolicyMakeToStock ProductionScheduleItemPolicyFulfillmentPolicy = "make_to_stock"
	ProductionScheduleItemPolicyFulfillmentPolicyMakeToOrder ProductionScheduleItemPolicyFulfillmentPolicy = "make_to_order"
)

// Resource type identifier.
type ProductionScheduleItemPolicyObject string

const (
	ProductionScheduleItemPolicyObjectProductionScheduleItemPolicy ProductionScheduleItemPolicyObject = "production_schedule_item_policy"
)

// Which rule decided that policy: the item itself, its product line, or the
// account default.
type ProductionScheduleItemPolicyPolicySource string

const (
	ProductionScheduleItemPolicyPolicySourceItem           ProductionScheduleItemPolicyPolicySource = "item"
	ProductionScheduleItemPolicyPolicySourceProductLine    ProductionScheduleItemPolicyPolicySource = "product_line"
	ProductionScheduleItemPolicyPolicySourceAccountDefault ProductionScheduleItemPolicyPolicySource = "account_default"
)

// One batch a release created, or would create: a single lot off one planned
// campaign.
type ReleaseScheduleBatch struct {
	// Entity is a polymorphic reference to any resource in the system.
	Batch Entity `json:"batch" api:"required"`
	// The number of the run this ticket came off, when the batch already existed.
	//
	// Present on a lot carried forward from an earlier week that the floor never
	// worked. The ticket is already printed and on the floor, so the release moves it
	// into the new run rather than issuing a replacement.
	CarriedForwardFrom string `json:"carried_forward_from" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Units in this lot.
	//
	// The last lot of a campaign is short when the planned quantity is not a whole
	// number of lots.
	Quantity float64 `json:"quantity" api:"required"`
	// The item's SKU, as it stood when the plan was generated.
	SKU string `json:"sku" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Batch              respjson.Field
		CarriedForwardFrom respjson.Field
		Item               respjson.Field
		Quantity           respjson.Field
		SKU                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReleaseScheduleBatch) RawJSON() string { return r.JSON.raw }
func (r *ReleaseScheduleBatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What releasing a week would create, with nothing written.
//
// A release makes a numbered production run and every batch under it, which is
// real work to undo by hand, so the confirmation is driven by this rather than by
// a count computed in the browser.
type ReleaseScheduleWeekPreview struct {
	// How many batches the run would hold, created and carried forward together.
	BatchCount int64 `json:"batch_count" api:"required"`
	// Why the week cannot be released, phrased for display.
	//
	// A week is blocked when it has already been released to the floor, or when it
	// holds nothing to release.
	BlockedReason string `json:"blocked_reason" api:"required"`
	// How many of `batch_count` would be moved off an earlier run rather than created.
	CarriedForwardBatchCount int64 `json:"carried_forward_batch_count" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ExistingProductionRun Entity `json:"existing_production_run" api:"required"`
	// Whether the week can be released.
	IsReleasable bool `json:"is_releasable" api:"required"`
	// How many campaigns would be released.
	LineCount int64 `json:"line_count" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lines ListReleasedScheduleLine `json:"lines" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_week_release_preview".
	Object ReleaseScheduleWeekPreviewObject `json:"object" api:"required"`
	// Total units that would be released.
	TotalQuantity float64 `json:"total_quantity" api:"required"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// First instant of the week.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchCount               respjson.Field
		BlockedReason            respjson.Field
		CarriedForwardBatchCount respjson.Field
		ExistingProductionRun    respjson.Field
		IsReleasable             respjson.Field
		LineCount                respjson.Field
		Lines                    respjson.Field
		Object                   respjson.Field
		TotalQuantity            respjson.Field
		WeekIndex                respjson.Field
		WeekStartsAt             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReleaseScheduleWeekPreview) RawJSON() string { return r.JSON.raw }
func (r *ReleaseScheduleWeekPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ReleaseScheduleWeekPreviewObject string

const (
	ReleaseScheduleWeekPreviewObjectProductionScheduleWeekReleasePreview ReleaseScheduleWeekPreviewObject = "production_schedule_week_release_preview"
)

// One planned campaign and the lots it broke into.
type ReleasedScheduleLine struct {
	// How many batches the campaign broke into.
	BatchCount int64 `json:"batch_count" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Batches ListReleaseScheduleBatch `json:"batches" api:"required"`
	// How much of `planned_quantity` is covered by tickets an earlier week already
	// issued.
	CarriedForwardQuantity float64 `json:"carried_forward_quantity" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Line Entity `json:"line" api:"required"`
	// Units in one lot.
	LotUnits float64 `json:"lot_units" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// Total units planned for the campaign.
	PlannedQuantity float64 `json:"planned_quantity" api:"required"`
	// The item's SKU, as it stood when the plan was generated.
	SKU string `json:"sku" api:"required"`
	// Abbreviation of the unit the quantity and the lot are counted in.
	//
	// `6 × 60` is not an instruction until it says 6 × 60 of what.
	Unit string `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchCount             respjson.Field
		Batches                respjson.Field
		CarriedForwardQuantity respjson.Field
		Item                   respjson.Field
		Line                   respjson.Field
		LotUnits               respjson.Field
		Machine                respjson.Field
		PlannedQuantity        respjson.Field
		SKU                    respjson.Field
		Unit                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReleasedScheduleLine) RawJSON() string { return r.JSON.raw }
func (r *ReleasedScheduleLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A demand override that changed a number, recorded so the plan can explain
// itself.
type ScheduleAppliedOverride struct {
	// How the override was expressed.
	//
	// - `absolute`: the override replaced the forecast for the month outright.
	// - `delta_units`: the override was added to the forecast.
	// - `delta_percent`: the override scaled the forecast.
	//
	// Any of "absolute", "delta_units", "delta_percent".
	Adjustment ScheduleAppliedOverrideAdjustment `json:"adjustment" api:"required"`
	// Demand after the override.
	After float64 `json:"after" api:"required"`
	// Demand before the override.
	Before float64 `json:"before" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// The first instant of the month the override applied to.
	MonthStartsAt time.Time `json:"month_starts_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Override Entity `json:"override" api:"required"`
	// Why the override exists.
	//
	// Any of "new_customer", "lost_account", "promotion", "seasonal_shift",
	// "new_product", "discontinued", "market_intelligence", "other".
	Reason ScheduleAppliedOverrideReason `json:"reason" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adjustment    respjson.Field
		After         respjson.Field
		Before        respjson.Field
		Item          respjson.Field
		MonthStartsAt respjson.Field
		Override      respjson.Field
		Reason        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleAppliedOverride) RawJSON() string { return r.JSON.raw }
func (r *ScheduleAppliedOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the override was expressed.
//
// - `absolute`: the override replaced the forecast for the month outright.
// - `delta_units`: the override was added to the forecast.
// - `delta_percent`: the override scaled the forecast.
type ScheduleAppliedOverrideAdjustment string

const (
	ScheduleAppliedOverrideAdjustmentAbsolute     ScheduleAppliedOverrideAdjustment = "absolute"
	ScheduleAppliedOverrideAdjustmentDeltaUnits   ScheduleAppliedOverrideAdjustment = "delta_units"
	ScheduleAppliedOverrideAdjustmentDeltaPercent ScheduleAppliedOverrideAdjustment = "delta_percent"
)

// Why the override exists.
type ScheduleAppliedOverrideReason string

const (
	ScheduleAppliedOverrideReasonNewCustomer        ScheduleAppliedOverrideReason = "new_customer"
	ScheduleAppliedOverrideReasonLostAccount        ScheduleAppliedOverrideReason = "lost_account"
	ScheduleAppliedOverrideReasonPromotion          ScheduleAppliedOverrideReason = "promotion"
	ScheduleAppliedOverrideReasonSeasonalShift      ScheduleAppliedOverrideReason = "seasonal_shift"
	ScheduleAppliedOverrideReasonNewProduct         ScheduleAppliedOverrideReason = "new_product"
	ScheduleAppliedOverrideReasonDiscontinued       ScheduleAppliedOverrideReason = "discontinued"
	ScheduleAppliedOverrideReasonMarketIntelligence ScheduleAppliedOverrideReason = "market_intelligence"
	ScheduleAppliedOverrideReasonOther              ScheduleAppliedOverrideReason = "other"
)

// An order commitment the plan does not meet.
type ScheduleAtRiskOrder struct {
	// Horizon week the constraint stage has to finish in for the order to ship on
	// time.
	DueWeek int64 `json:"due_week" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "schedule_at_risk_order".
	Object ScheduleAtRiskOrderObject `json:"object" api:"required"`
	// Why the commitment is at risk.
	//
	//   - `past_due`: production needed to start before this plan begins.
	//   - `undated`: the order carries no ship-by commitment, so it is treated as owed
	//     now.
	//   - `short`: the plan projects less stock than the order needs in the week it is
	//     needed.
	//
	// Any of "past_due", "undated", "short".
	Reason ScheduleAtRiskOrderReason `json:"reason" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	SalesOrder Entity `json:"sales_order" api:"required"`
	// SKU of that item.
	SKU string `json:"sku" api:"required"`
	// Outstanding quantity still owed.
	Units float64 `json:"units" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DueWeek     respjson.Field
		Item        respjson.Field
		Object      respjson.Field
		Reason      respjson.Field
		SalesOrder  respjson.Field
		SKU         respjson.Field
		Units       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleAtRiskOrder) RawJSON() string { return r.JSON.raw }
func (r *ScheduleAtRiskOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ScheduleAtRiskOrderObject string

const (
	ScheduleAtRiskOrderObjectScheduleAtRiskOrder ScheduleAtRiskOrderObject = "schedule_at_risk_order"
)

// Why the commitment is at risk.
//
//   - `past_due`: production needed to start before this plan begins.
//   - `undated`: the order carries no ship-by commitment, so it is treated as owed
//     now.
//   - `short`: the plan projects less stock than the order needs in the week it is
//     needed.
type ScheduleAtRiskOrderReason string

const (
	ScheduleAtRiskOrderReasonPastDue ScheduleAtRiskOrderReason = "past_due"
	ScheduleAtRiskOrderReasonUndated ScheduleAtRiskOrderReason = "undated"
	ScheduleAtRiskOrderReasonShort   ScheduleAtRiskOrderReason = "short"
)

// What the solver could not do, and why the plan differs from raw history.
type ScheduleDiagnostics struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	AppliedOverrides ListScheduleAppliedOverride `json:"applied_overrides" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	AtRiskOrders ListScheduleAtRiskOrder `json:"at_risk_orders" api:"required"`
	// Average inputs a product transition introduces, measured from history.
	AverageInputsAdded float64 `json:"average_inputs_added" api:"required"`
	// Items below their reorder point that never won a slot in the horizon.
	//
	// This is the signal that the plant is short of capacity.
	CapacityStarvedSKUs []string `json:"capacity_starved_skus" api:"required"`
	// Minutes of changeover the model adds for each new input a product transition
	// introduces.
	//
	// Calibrated from measured production against `average_inputs_added`, so the
	// modelled changeover lands on the time the floor actually reports rather than on
	// a fixed allowance.
	ChangeoverSlopeMinutes float64 `json:"changeover_slope_minutes" api:"required"`
	// Machines the constraint department contributed to this solve.
	ConstraintMachineCount int64 `json:"constraint_machine_count" api:"required"`
	// Items whose economic lot size was reduced to fit one machine-week, meaning
	// shorter and more frequent campaigns.
	EoqCappedSKUs []string `json:"eoq_capped_skus" api:"required"`
	// Number of items the merchant has excluded from planning.
	ExcludedItemCount int64 `json:"excluded_item_count" api:"required"`
	// How the second stage fared: what it could not make, and which of the two things
	// it ran out of.
	//
	// The two starvation lists are the point of planning in two stages at all. A
	// finished good held back for want of greige is a knitting problem — knit more of
	// it, or knit it sooner — and one held back for want of hours is a finishing
	// problem: another shift, or a different mix. A single "short" list would throw
	// that distinction away, and it is the only thing this model knows that a
	// one-stage plan does not.
	Finishing ScheduleFinishingDiagnostics `json:"finishing" api:"required"`
	// Whether the second stage's capacity was estimated rather than counted from
	// machines.
	FinishingCapacityIsEstimated bool `json:"finishing_capacity_is_estimated" api:"required"`
	// Machines outside the constraint department that the second stage was sized from.
	//
	// Zero means its capacity was estimated from the shift pattern alone rather than
	// counted.
	FinishingMachineCount int64 `json:"finishing_machine_count" api:"required"`
	// Outstanding order quantity this plan owes, expressed in the constraint item's
	// own unit.
	//
	// Zero means nothing is on order and the plan is driven purely by the forecast.
	FirmDemandUnits float64 `json:"firm_demand_units" api:"required"`
	// Items with no measured run rate, which cannot be scheduled because their machine
	// time is unknown.
	ItemsWithoutRunRate []string `json:"items_without_run_rate" api:"required"`
	// Machines in the constraint department with no production step.
	//
	// Their campaigns derive no downstream department work.
	MachinesWithoutStep int64 `json:"machines_without_step" api:"required"`
	// Planned items built only against the order book rather than to a forecast.
	MakeToOrderItemCount int64 `json:"make_to_order_item_count" api:"required"`
	// Batches found on those machines in the demand window.
	//
	// Zero means nothing has been scanned there, which is why a plan can be empty even
	// with machines configured.
	MeasuredBatchCount int64 `json:"measured_batch_count" api:"required"`
	// Open orders carrying no ship-by commitment, dated at the front of the horizon
	// because they are issued and unshipped.
	//
	// A non-zero count means orders placed before commitments were tracked still need
	// a ship-by date.
	UndatedFirmOrderCount int64 `json:"undated_firm_order_count" api:"required"`
	// Items that cannot fit even a single lot into a machine-week and are therefore
	// never scheduled.
	UnschedulableSKUs []string `json:"unschedulable_skus" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppliedOverrides             respjson.Field
		AtRiskOrders                 respjson.Field
		AverageInputsAdded           respjson.Field
		CapacityStarvedSKUs          respjson.Field
		ChangeoverSlopeMinutes       respjson.Field
		ConstraintMachineCount       respjson.Field
		EoqCappedSKUs                respjson.Field
		ExcludedItemCount            respjson.Field
		Finishing                    respjson.Field
		FinishingCapacityIsEstimated respjson.Field
		FinishingMachineCount        respjson.Field
		FirmDemandUnits              respjson.Field
		ItemsWithoutRunRate          respjson.Field
		MachinesWithoutStep          respjson.Field
		MakeToOrderItemCount         respjson.Field
		MeasuredBatchCount           respjson.Field
		UndatedFirmOrderCount        respjson.Field
		UnschedulableSKUs            respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleDiagnostics) RawJSON() string { return r.JSON.raw }
func (r *ScheduleDiagnostics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the second stage fared: what it could not make, and which of the two things
// it ran out of.
//
// The two starvation lists are the point of planning in two stages at all. A
// finished good held back for want of greige is a knitting problem — knit more of
// it, or knit it sooner — and one held back for want of hours is a finishing
// problem: another shift, or a different mix. A single "short" list would throw
// that distinction away, and it is the only thing this model knows that a
// one-stage plan does not.
type ScheduleFinishingDiagnostics struct {
	// Finished goods that had greige and never had hours.
	CapacityStarvedSKUs []string `json:"capacity_starved_skus" api:"required"`
	// Finished goods that wanted building across the whole horizon and never had
	// greige to build from.
	GreigeStarvedSKUs []string `json:"greige_starved_skus" api:"required"`
	// Finished goods with no measured finishing rate, which cannot be levelled because
	// the hours they cost are unknown.
	ItemsWithoutRunRate []string `json:"items_without_run_rate" api:"required"`
	// How many finishing lines the plan holds.
	LineCount int64 `json:"line_count" api:"required"`
	// Hours the plan asks of it, week by week.
	PlannedHoursByWeek []float64 `json:"planned_hours_by_week" api:"required"`
	// Total finished units the stage plans across the horizon.
	TotalPlannedUnits float64 `json:"total_planned_units" api:"required"`
	// Constraint output the horizon never converts into anything.
	//
	// A large figure means the two stages are planned against different demand, which
	// is worth looking at rather than leaving as an unexplained pile of greige.
	UnusedGreigeUnits float64 `json:"unused_greige_units" api:"required"`
	// Those hours as a fraction of capacity, week by week.
	UtilisationByWeek []float64 `json:"utilisation_by_week" api:"required"`
	// Hours the second stage can work in one week.
	WeeklyCapacityHours float64 `json:"weekly_capacity_hours" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CapacityStarvedSKUs respjson.Field
		GreigeStarvedSKUs   respjson.Field
		ItemsWithoutRunRate respjson.Field
		LineCount           respjson.Field
		PlannedHoursByWeek  respjson.Field
		TotalPlannedUnits   respjson.Field
		UnusedGreigeUnits   respjson.Field
		UtilisationByWeek   respjson.Field
		WeeklyCapacityHours respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleFinishingDiagnostics) RawJSON() string { return r.JSON.raw }
func (r *ScheduleFinishingDiagnostics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An order this schedule does not build in time, with the campaigns covering the
// part it does.
type ScheduleOrderCoverage struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	CoveringLines ListScheduleOrderCoverageLine `json:"covering_lines" api:"required"`
	// Horizon week the constraint stage has to finish in for the order to ship on
	// time.
	DueWeek int64 `json:"due_week" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "schedule_order_coverage".
	Object ScheduleOrderCoverageObject `json:"object" api:"required"`
	// Why the commitment is at risk.
	//
	// Any of "past_due", "undated", "short".
	Reason ScheduleOrderCoverageReason `json:"reason" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	SalesOrder Entity `json:"sales_order" api:"required"`
	// The date this order is contractually due to ship.
	ShipByDate time.Time `json:"ship_by_date" api:"required" format:"date-time"`
	// SKU of that item.
	SKU string `json:"sku" api:"required"`
	// Quantity the plan does not build in time.
	//
	// Less than the whole order when the plan builds part of it — a mostly-built order
	// is mostly built, and reporting the full quantity would read as a total miss.
	UnitsAtRisk float64 `json:"units_at_risk" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CoveringLines respjson.Field
		DueWeek       respjson.Field
		Item          respjson.Field
		Object        respjson.Field
		Reason        respjson.Field
		SalesOrder    respjson.Field
		ShipByDate    respjson.Field
		SKU           respjson.Field
		UnitsAtRisk   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleOrderCoverage) RawJSON() string { return r.JSON.raw }
func (r *ScheduleOrderCoverage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ScheduleOrderCoverageObject string

const (
	ScheduleOrderCoverageObjectScheduleOrderCoverage ScheduleOrderCoverageObject = "schedule_order_coverage"
)

// Why the commitment is at risk.
type ScheduleOrderCoverageReason string

const (
	ScheduleOrderCoverageReasonPastDue ScheduleOrderCoverageReason = "past_due"
	ScheduleOrderCoverageReasonUndated ScheduleOrderCoverageReason = "undated"
	ScheduleOrderCoverageReasonShort   ScheduleOrderCoverageReason = "short"
)

// One campaign earmarked for an order.
type ScheduleOrderCoverageLine struct {
	// Quantity of that campaign earmarked for this order.
	AllocatedQuantity float64 `json:"allocated_quantity" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// Resource type identifier.
	//
	// Any of "schedule_order_coverage_line".
	Object ScheduleOrderCoverageLineObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionScheduleLine Entity `json:"production_schedule_line" api:"required"`
	// Horizon week it runs in.
	WeekIndex int64 `json:"week_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllocatedQuantity      respjson.Field
		Machine                respjson.Field
		Object                 respjson.Field
		ProductionScheduleLine respjson.Field
		WeekIndex              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleOrderCoverageLine) RawJSON() string { return r.JSON.raw }
func (r *ScheduleOrderCoverageLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ScheduleOrderCoverageLineObject string

const (
	ScheduleOrderCoverageLineObjectScheduleOrderCoverageLine ScheduleOrderCoverageLineObject = "schedule_order_coverage_line"
)

type OperationProductionScheduleDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationProductionScheduleDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationProductionScheduleDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleNewParams struct {
	// Request to generate a production schedule.
	GenerateProductionScheduleRequest GenerateProductionScheduleRequestParam
	paramObj
}

func (r OperationProductionScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GenerateProductionScheduleRequest)
}
func (r *OperationProductionScheduleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Only return versions in these lifecycle states.
	//
	//   - `draft`: still editable and committed to nothing.
	//   - `generating`: the solver is still building the version.
	//   - `published`: live, with its first weeks frozen as a commitment to the floor.
	//   - `superseded`: a later version was published over the same horizon and replaced
	//     this one.
	//   - `archived`: retired without being replaced.
	//   - `failed`: the solver could not produce a plan.
	//
	// Any of "draft", "generating", "published", "superseded", "archived", "failed".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleListParams]'s query parameters
// as `url.Values`.
func (r OperationProductionScheduleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationProductionScheduleGetDerivedLinesParams struct {
	// Only return work in this horizon week, zero-based.
	WeekIndex param.Opt[int64] `query:"week_index,omitzero" json:"-"`
	// Only return work owned by these departments.
	DepartmentIDs []string `query:"department_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleGetDerivedLinesParams]'s query
// parameters as `url.Values`.
func (r OperationProductionScheduleGetDerivedLinesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationProductionScheduleGetDeviationsParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Whether the change fell inside the frozen window.
	//
	// Judged against the freeze as it stood when the change was made, not as it stands
	// now, so a later publish cannot reclassify history.
	Frozen param.Opt[bool] `query:"frozen,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleGetDeviationsParams]'s query
// parameters as `url.Values`.
func (r OperationProductionScheduleGetDeviationsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationProductionScheduleGetFinishingLinesParams struct {
	// Only the finishing planned for this finished good.
	ItemID param.Opt[string] `query:"item_id,omitzero" json:"-"`
	// Only the finishing planned for this week, zero-based from the start of the
	// horizon.
	WeekIndex param.Opt[int64] `query:"week_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleGetFinishingLinesParams]'s query
// parameters as `url.Values`.
func (r OperationProductionScheduleGetFinishingLinesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationProductionScheduleGetWeekReleasePreviewParams struct {
	// Preview the week as if every batch were newly issued.
	//
	// By default the preview counts tickets an earlier week issued and the floor never
	// worked against this week's campaigns, because that is what releasing would do.
	SkipCarryForward param.Opt[bool] `query:"skip_carry_forward,omitzero" json:"-"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex param.Opt[int64] `query:"week_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationProductionScheduleGetWeekReleasePreviewParams]'s
// query parameters as `url.Values`.
func (r OperationProductionScheduleGetWeekReleasePreviewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
