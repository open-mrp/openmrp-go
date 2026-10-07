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
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Generate and review machine-level production schedules.
//
// OperationProductionScheduleActionService contains methods and other services
// that help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationProductionScheduleActionService] method instead.
type OperationProductionScheduleActionService struct {
	options []option.RequestOption
}

// NewOperationProductionScheduleActionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationProductionScheduleActionService(opts ...option.RequestOption) (r OperationProductionScheduleActionService) {
	r = OperationProductionScheduleActionService{}
	r.options = opts
	return
}

// Archives a schedule version, retiring it without discarding its history.
//
// Any version that is not already archived can be archived, including a draft that
// was never published. The version stays readable — its campaigns, policy snapshot
// and deviation log are kept — and it still backs any attainment already measured
// against it.
//
// Archiving does not supersede anything or promote another version in its place.
// To take a published version out of use by replacing it, generate and publish a
// newer one instead.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleActionService) Archive(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/actions/archive", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Runs the production scheduling solver and returns the plan without saving it.
//
// This is the inspection surface for the scheduler: it takes the same path a
// generated schedule will take, minus the write, so a plan can be reviewed and
// compared before anything depends on it. No version is created and nothing is
// numbered, so this can be called as often as needed.
//
// The solver plans the constraint department — the room that sets the pace of the
// factory — so production schedule settings must name one and it must have
// machines that are included in planning. Without that there is nothing to
// schedule and the request is rejected rather than returning an empty plan.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleActionService) Preview(ctx context.Context, body OperationProductionScheduleActionPreviewParams, opts ...option.RequestOption) (res *ProductionSchedulePreview, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/production-schedules/actions/preview"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns what regenerating this draft would change, without changing it.
//
// Every campaign either plan holds is listed, including the ones both agree on, so
// the caller can render a full side-by-side rather than a list of surprises. Only
// a draft can be previewed, for the same reason only a draft can be regenerated.
//
// The comparison is run the way a regenerate runs by default — hand-edited
// campaigns are kept, and the fresh solve plans around them — so they read as
// unchanged rather than as work the solver wants to take away. `manual_line_count`
// is how many campaigns on the draft were placed or edited by hand, which is the
// work a `replace_all` regenerate is putting at risk.
//
// The horizon and demand basis default to the ones this version already has, so a
// plain call changes only how current the plan is, not what question is being
// asked.
//
// This endpoint requires the permission: `production_schedules:read`.
func (r *OperationProductionScheduleActionService) PreviewRegenerate(ctx context.Context, id string, body OperationProductionScheduleActionPreviewRegenerateParams, opts ...option.RequestOption) (res *ProductionScheduleRegeneratePreview, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/actions/preview-regenerate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Publishes a draft schedule, freezing its first weeks.
//
// Publishing is what makes a plan a commitment: the frozen weeks' lines are marked
// frozen, the frozen line count and quantity are captured onto the version, and
// any published version whose horizon overlaps this one's is superseded rather
// than rewritten. After this, a change inside the frozen window has to state a
// reason.
//
// Only a draft can be published. How many weeks freeze comes from the account's
// frozen-weeks setting as it stood when the version was generated, and a version
// generated with zero frozen weeks publishes without committing to anything.
//
// The frozen counts are snapshotted here and never recomputed, so adherence keeps
// the denominator it was committed to.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleActionService) Publish(ctx context.Context, id string, opts ...option.RequestOption) (res *ProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/actions/publish", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Re-solves a draft in place, keeping its version number.
//
// Only a draft can be regenerated. A published version is a commitment the floor
// is already working to, and a superseded or archived one is history; re-solving
// either in place would change what a week was measured against after the fact. To
// replan against a published version, generate a new one — publishing it
// supersedes the current one.
//
// The version number is kept deliberately: minting a new version for every
// re-solve would fill the list with drafts nobody asked for and make the version
// number meaningless as a count of the plans actually considered.
//
// Every hand edit a `replace_all` destroys is written to the deviation log before
// it goes, so "where did my change go" stays answerable. Call `preview-regenerate`
// first to see what a re-solve would change.
//
// Aside from the hand edits a `preserve_manual` run keeps, the version's
// campaigns, policy snapshot, derived department work, solver diagnostics and
// settings snapshot are all replaced with the fresh solve's, so the plan can still
// explain itself afterwards.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleActionService) Regenerate(ctx context.Context, id string, body OperationProductionScheduleActionRegenerateParams, opts ...option.RequestOption) (res *ProductionSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/actions/regenerate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Turns one planned week into a production run.
//
// Each campaign in the week becomes one batch per planned lot, using the lot size
// the campaign was planned at. A 360-unit campaign at a 60-unit lot arrives on the
// floor as six batches, not one instruction to make 360; a quantity that is not a
// whole number of lots trails a single short lot at the end of the run.
//
// The release is atomic. A run holding half a week's batches is worse than no run,
// because the missing half looks like work nobody was asked to do and attainment
// would count it as unplanned production.
//
// Releasing the same week twice fails rather than creating a second run. Each
// released line records the run now carrying it, and a line that is already
// released is never re-pointed.
//
// Lots an earlier week already issued are carried forward rather than reissued.
// When a week fell short, the next plan asks for the shortfall — and the batches
// covering it are usually already printed and sitting on the floor. Those tickets
// are moved into this run and counted against the campaign, so only the genuinely
// new work is created. `carried_forward_batch_count` says how many arrived that
// way, and each one names the run it came off. Send `skip_carry_forward` to issue
// the whole week new instead.
//
// Cancelled campaigns and campaigns planned at zero are left behind rather than
// released. A week that would produce an implausible number of batches is rejected
// outright, since that is far more likely to be a misconfigured lot size than a
// real week's work.
//
// This endpoint requires the permission: `production_schedules:update`.
func (r *OperationProductionScheduleActionService) ReleaseWeek(ctx context.Context, id string, body OperationProductionScheduleActionReleaseWeekParams, opts ...option.RequestOption) (res *ReleaseScheduleWeekResult, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/production-schedules/%s/actions/release-week", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleCampaign struct {
	// Resources in this page.
	Data []ScheduleCampaign `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleCampaignObject `json:"object" api:"required"`
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
func (r ListScheduleCampaign) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleCampaignObject string

const (
	ListScheduleCampaignObjectList ListScheduleCampaignObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleDiffLine struct {
	// Resources in this page.
	Data []ScheduleDiffLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleDiffLineObject `json:"object" api:"required"`
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
func (r ListScheduleDiffLine) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleDiffLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleDiffLineObject string

const (
	ListScheduleDiffLineObjectList ListScheduleDiffLineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListSchedulePolicy struct {
	// Resources in this page.
	Data []SchedulePolicy `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSchedulePolicyObject `json:"object" api:"required"`
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
func (r ListSchedulePolicy) RawJSON() string { return r.JSON.raw }
func (r *ListSchedulePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSchedulePolicyObject string

const (
	ListSchedulePolicyObjectList ListSchedulePolicyObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScheduleProjection struct {
	// Resources in this page.
	Data []ScheduleProjection `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScheduleProjectionObject `json:"object" api:"required"`
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
func (r ListScheduleProjection) RawJSON() string { return r.JSON.raw }
func (r *ListScheduleProjection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScheduleProjectionObject string

const (
	ListScheduleProjectionObjectList ListScheduleProjectionObject = "list"
)

// Request to preview a production schedule.
type PreviewProductionScheduleRequestParam struct {
	// Number of weeks the plan should cover, overriding the account's configured
	// horizon for this preview only.
	HorizonWeeks param.Opt[int64] `json:"horizon_weeks,omitzero"`
	// The instant to plan against, which is what stock, demand history and active
	// demand overrides are read as of.
	//
	// Left unset, the preview is solved against the moment the request arrives. The
	// horizon starts on the account's configured week-start day on or before this
	// instant, so backdating this shifts the whole week grid.
	PlanningAsOf param.Opt[time.Time] `json:"planning_as_of,omitzero" format:"date-time"`
	// How future demand is derived, overriding the account's configured basis for this
	// preview only.
	//
	//   - `trailing_12`: demand is the trailing twelve months of orders.
	//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
	//     a season arriving early or late rather than flattening it.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis PreviewProductionScheduleRequestDemandBasis `json:"demand_basis,omitzero"`
	paramObj
}

func (r PreviewProductionScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewProductionScheduleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewProductionScheduleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How future demand is derived, overriding the account's configured basis for this
// preview only.
//
//   - `trailing_12`: demand is the trailing twelve months of orders.
//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
//     a season arriving early or late rather than flattening it.
type PreviewProductionScheduleRequestDemandBasis string

const (
	PreviewProductionScheduleRequestDemandBasisTrailing12  PreviewProductionScheduleRequestDemandBasis = "trailing_12"
	PreviewProductionScheduleRequestDemandBasisSeasonalEma PreviewProductionScheduleRequestDemandBasis = "seasonal_ema"
)

// Request to see what a re-solve would change.
type PreviewRegenerateProductionScheduleRequestParam struct {
	// Number of weeks the re-solve should cover, defaulting to the horizon this
	// version already has.
	HorizonWeeks param.Opt[int64] `json:"horizon_weeks,omitzero"`
	// The instant to plan against, which is what stock, demand history and active
	// demand overrides are read as of.
	//
	// Defaults to now rather than to the instant the version was first generated, so a
	// plain call answers "what would the solver say today". Because the horizon
	// re-anchors to the week containing this instant, a campaign can appear under a
	// different `week_index` than the one stored on the draft.
	PlanningAsOf param.Opt[time.Time] `json:"planning_as_of,omitzero" format:"date-time"`
	// How future demand is derived, defaulting to the basis this version was solved
	// with.
	//
	//   - `trailing_12`: demand is the trailing twelve months of orders.
	//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
	//     a season arriving early or late rather than flattening it.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis PreviewRegenerateProductionScheduleRequestDemandBasis `json:"demand_basis,omitzero"`
	paramObj
}

func (r PreviewRegenerateProductionScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewRegenerateProductionScheduleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewRegenerateProductionScheduleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How future demand is derived, defaulting to the basis this version was solved
// with.
//
//   - `trailing_12`: demand is the trailing twelve months of orders.
//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
//     a season arriving early or late rather than flattening it.
type PreviewRegenerateProductionScheduleRequestDemandBasis string

const (
	PreviewRegenerateProductionScheduleRequestDemandBasisTrailing12  PreviewRegenerateProductionScheduleRequestDemandBasis = "trailing_12"
	PreviewRegenerateProductionScheduleRequestDemandBasisSeasonalEma PreviewRegenerateProductionScheduleRequestDemandBasis = "seasonal_ema"
)

// A production plan produced by the scheduling solver.
type ProductionSchedulePreview struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Campaigns ListScheduleCampaign `json:"campaigns" api:"required"`
	// What the solver could not do, and why the plan differs from raw history.
	Diagnostics ScheduleDiagnostics `json:"diagnostics" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_preview".
	Object ProductionSchedulePreviewObject `json:"object" api:"required"`
	// The instant the plan was calculated against.
	PlanningAsOfAt time.Time `json:"planning_as_of_at" api:"required" format:"date-time"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Policies ListSchedulePolicy `json:"policies" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Projections ListScheduleProjection `json:"projections" api:"required"`
	// Version of the solver that produced this plan.
	SolverVersion string `json:"solver_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Campaigns      respjson.Field
		Diagnostics    respjson.Field
		Object         respjson.Field
		PlanningAsOfAt respjson.Field
		Policies       respjson.Field
		Projections    respjson.Field
		SolverVersion  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionSchedulePreview) RawJSON() string { return r.JSON.raw }
func (r *ProductionSchedulePreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionSchedulePreviewObject string

const (
	ProductionSchedulePreviewObjectProductionSchedulePreview ProductionSchedulePreviewObject = "production_schedule_preview"
)

// What a regenerate would change about a draft, without changing it.
//
// A regenerate that silently discards hand-work is abandoned within two cycles, so
// the destructive mode states its cost as a number before it runs:
// `discarded_manual_count` is exactly how many hand-edited campaigns `replace_all`
// would destroy.
type ProductionScheduleRegeneratePreview struct {
	// Campaigns the fresh solve wants that the current plan does not have.
	AddedCount int64 `json:"added_count" api:"required"`
	// Campaigns both hold, in different quantities.
	ChangedCount int64 `json:"changed_count" api:"required"`
	// Hand-edited campaigns `replace_all` would destroy.
	DiscardedManualCount int64 `json:"discarded_manual_count" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lines ListScheduleDiffLine `json:"lines" api:"required"`
	// Hand-edited campaigns currently on the draft.
	ManualLineCount int64 `json:"manual_line_count" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_regenerate_preview".
	Object ProductionScheduleRegeneratePreviewObject `json:"object" api:"required"`
	// The instant the fresh solve planned from.
	//
	// Unless the caller names an instant, a regenerate plans from now rather than
	// replaying the one the draft was first generated against, so demand overrides
	// added since then are taken into account and the horizon re-anchors to today.
	PlanningAsOfAt time.Time `json:"planning_as_of_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductionSchedule Entity `json:"production_schedule" api:"required"`
	// Campaigns the current plan has that the fresh solve does not want.
	RemovedCount int64 `json:"removed_count" api:"required"`
	// Which solver produced the proposal.
	SolverVersion string `json:"solver_version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddedCount           respjson.Field
		ChangedCount         respjson.Field
		DiscardedManualCount respjson.Field
		Lines                respjson.Field
		ManualLineCount      respjson.Field
		Object               respjson.Field
		PlanningAsOfAt       respjson.Field
		ProductionSchedule   respjson.Field
		RemovedCount         respjson.Field
		SolverVersion        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionScheduleRegeneratePreview) RawJSON() string { return r.JSON.raw }
func (r *ProductionScheduleRegeneratePreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionScheduleRegeneratePreviewObject string

const (
	ProductionScheduleRegeneratePreviewObjectProductionScheduleRegeneratePreview ProductionScheduleRegeneratePreviewObject = "production_schedule_regenerate_preview"
)

// Request to re-solve a draft in place.
type RegenerateProductionScheduleRequestParam struct {
	// Number of weeks the re-solve should cover, defaulting to the horizon this
	// version already has.
	HorizonWeeks param.Opt[int64] `json:"horizon_weeks,omitzero"`
	// The instant to plan against, which is what stock, demand history and active
	// demand overrides are read as of.
	//
	// Defaults to now rather than to the instant the version was first generated, so a
	// plain call answers "what would the solver say today". Because the horizon
	// re-anchors to the week containing this instant, a kept campaign keeps the
	// calendar week it was planned in but can end up under a different `week_index`.
	PlanningAsOf param.Opt[time.Time] `json:"planning_as_of,omitzero" format:"date-time"`
	// How future demand is derived, defaulting to the basis this version was solved
	// with.
	//
	//   - `trailing_12`: demand is the trailing twelve months of orders.
	//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
	//     a season arriving early or late rather than flattening it.
	//
	// Any of "trailing_12", "seasonal_ema".
	DemandBasis RegenerateProductionScheduleRequestDemandBasis `json:"demand_basis,omitzero"`
	// What happens to the campaigns someone placed or edited by hand.
	//
	//   - `preserve_manual`: hand-edited campaigns are kept, and the fresh solve plans
	//     around them — their stock and machine time are facts the rest of the plan
	//     responds to.
	//   - `replace_all`: hand edits are discarded and the fresh solve is taken whole.
	//
	// Omitting this keeps hand edits, because the alternative destroys work silently.
	//
	// Any of "preserve_manual", "replace_all".
	MergeMode RegenerateProductionScheduleRequestMergeMode `json:"merge_mode,omitzero"`
	paramObj
}

func (r RegenerateProductionScheduleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RegenerateProductionScheduleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegenerateProductionScheduleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How future demand is derived, defaulting to the basis this version was solved
// with.
//
//   - `trailing_12`: demand is the trailing twelve months of orders.
//   - `seasonal_ema`: demand is a seasonal exponential moving average, which follows
//     a season arriving early or late rather than flattening it.
type RegenerateProductionScheduleRequestDemandBasis string

const (
	RegenerateProductionScheduleRequestDemandBasisTrailing12  RegenerateProductionScheduleRequestDemandBasis = "trailing_12"
	RegenerateProductionScheduleRequestDemandBasisSeasonalEma RegenerateProductionScheduleRequestDemandBasis = "seasonal_ema"
)

// What happens to the campaigns someone placed or edited by hand.
//
//   - `preserve_manual`: hand-edited campaigns are kept, and the fresh solve plans
//     around them — their stock and machine time are facts the rest of the plan
//     responds to.
//   - `replace_all`: hand edits are discarded and the fresh solve is taken whole.
//
// Omitting this keeps hand edits, because the alternative destroys work silently.
type RegenerateProductionScheduleRequestMergeMode string

const (
	RegenerateProductionScheduleRequestMergeModePreserveManual RegenerateProductionScheduleRequestMergeMode = "preserve_manual"
	RegenerateProductionScheduleRequestMergeModeReplaceAll     RegenerateProductionScheduleRequestMergeMode = "replace_all"
)

// Request to release one week of a production schedule to the floor.
//
// The properties ResponsibleUserID, WeekIndex are required.
type ReleaseProductionScheduleWeekRequestParam struct {
	// ID of the account user accountable for executing the run.
	//
	// Accepts either an account user ID or a user ID; it is resolved and stored as the
	// account user.
	ResponsibleUserID string `json:"responsible_user_id" api:"required"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// ID of the scanning station the batches will be scanned at.
	//
	// Applied to every batch this release creates, across all machines in the week.
	ScanningStationID param.Opt[string] `json:"scanning_station_id,omitzero"`
	// Issue the whole week as new batches, leaving an earlier week's unworked lots
	// where they are.
	//
	// Off unless you ask for it: reprinting a ticket the floor is already holding is
	// exactly what carrying work forward exists to prevent, so it takes a deliberate
	// choice to do it.
	SkipCarryForward param.Opt[bool] `json:"skip_carry_forward,omitzero"`
	paramObj
}

func (r ReleaseProductionScheduleWeekRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ReleaseProductionScheduleWeekRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReleaseProductionScheduleWeekRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The production run created from one week of a schedule.
//
// Each planned campaign becomes one batch per lot, so a 360-unit week at a 60-unit
// lot arrives on the floor as six batches rather than one instruction to make 360.
type ReleaseScheduleWeekResult struct {
	// How many batches the run holds across all campaigns, created and carried forward
	// together.
	BatchCount int64 `json:"batch_count" api:"required"`
	// How many of `batch_count` were moved off an earlier run rather than created.
	//
	// Tickets for these are already printed and on the floor.
	CarriedForwardBatchCount int64 `json:"carried_forward_batch_count" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lines ListReleasedScheduleLine `json:"lines" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_schedule_week_release".
	Object ReleaseScheduleWeekResultObject `json:"object" api:"required"`
	// A production run: the group of shop-floor batches that are executed together,
	// tracked from the first batch scan through to completion.
	ProductionRun ProductionRun `json:"production_run" api:"required"`
	// How many campaigns were released.
	ReleasedLineCount int64 `json:"released_line_count" api:"required"`
	// Total units released.
	TotalQuantity float64 `json:"total_quantity" api:"required"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// First instant of the released week.
	WeekStartsAt time.Time `json:"week_starts_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchCount               respjson.Field
		CarriedForwardBatchCount respjson.Field
		Lines                    respjson.Field
		Object                   respjson.Field
		ProductionRun            respjson.Field
		ReleasedLineCount        respjson.Field
		TotalQuantity            respjson.Field
		WeekIndex                respjson.Field
		WeekStartsAt             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReleaseScheduleWeekResult) RawJSON() string { return r.JSON.raw }
func (r *ReleaseScheduleWeekResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ReleaseScheduleWeekResultObject string

const (
	ReleaseScheduleWeekResultObjectProductionScheduleWeekRelease ReleaseScheduleWeekResultObject = "production_schedule_week_release"
)

// One planned production block: make this item, on this machine, in this week.
type ScheduleCampaign struct {
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Whole lots the quantity rounds to.
	Lots int64 `json:"lots" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// Constraint hours the campaign consumes.
	RunHours float64 `json:"run_hours" api:"required"`
	// SKU of the item.
	SKU string `json:"sku" api:"required"`
	// Quantity to produce.
	Units float64 `json:"units" api:"required"`
	// Zero-based week offset from the start of the horizon.
	WeekIndex int64 `json:"week_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item        respjson.Field
		Lots        respjson.Field
		Machine     respjson.Field
		RunHours    respjson.Field
		SKU         respjson.Field
		Units       respjson.Field
		WeekIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleCampaign) RawJSON() string { return r.JSON.raw }
func (r *ScheduleCampaign) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One campaign as the current plan and a fresh solve each see it.
type ScheduleDiffLine struct {
	// What the regenerate would do to this campaign.
	//
	// - `added`: the fresh solve wants a campaign the current plan does not have.
	// - `removed`: the current plan holds a campaign the fresh solve does not want.
	// - `changed`: both hold the campaign, in different quantities.
	// - `unchanged`: both agree on it.
	//
	// Any of "added", "removed", "changed", "unchanged".
	Change ScheduleDiffLineChange `json:"change" api:"required"`
	// Whether the current campaign was created or edited by a person.
	CurrentIsManual bool `json:"current_is_manual" api:"required"`
	// Units the current plan asks for.
	//
	// Zero when the campaign is being added.
	CurrentQuantity float64 `json:"current_quantity" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Machine Entity `json:"machine" api:"required"`
	// Units the fresh solve asks for.
	//
	// Zero when the campaign is being removed.
	ProposedQuantity float64 `json:"proposed_quantity" api:"required"`
	// SKU of that item.
	SKU string `json:"sku" api:"required"`
	// Zero-based horizon week.
	WeekIndex int64 `json:"week_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Change           respjson.Field
		CurrentIsManual  respjson.Field
		CurrentQuantity  respjson.Field
		Item             respjson.Field
		Machine          respjson.Field
		ProposedQuantity respjson.Field
		SKU              respjson.Field
		WeekIndex        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleDiffLine) RawJSON() string { return r.JSON.raw }
func (r *ScheduleDiffLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What the regenerate would do to this campaign.
//
// - `added`: the fresh solve wants a campaign the current plan does not have.
// - `removed`: the current plan holds a campaign the fresh solve does not want.
// - `changed`: both hold the campaign, in different quantities.
// - `unchanged`: both agree on it.
type ScheduleDiffLineChange string

const (
	ScheduleDiffLineChangeAdded     ScheduleDiffLineChange = "added"
	ScheduleDiffLineChangeRemoved   ScheduleDiffLineChange = "removed"
	ScheduleDiffLineChangeChanged   ScheduleDiffLineChange = "changed"
	ScheduleDiffLineChangeUnchanged ScheduleDiffLineChange = "unchanged"
)

// The inventory policy computed for one item.
type SchedulePolicy struct {
	// ABC class by share of constraint run hours.
	//
	// - `a`: consumes the largest share of constraint capacity.
	// - `b`: moderate constraint consumption.
	// - `c`: consumes little constraint capacity.
	//
	// Any of "a", "b", "c".
	AbcClass SchedulePolicyAbcClass `json:"abc_class" api:"required"`
	// Demand used for planning, annualized.
	AnnualDemand float64 `json:"annual_demand" api:"required"`
	// Constraint hours this item's annual demand consumes.
	AnnualRunHours float64 `json:"annual_run_hours" api:"required"`
	// What the constraint stage holds on average: its buffer plus half a campaign.
	AverageGreigeInventory float64 `json:"average_greige_inventory" api:"required"`
	// Observed or default lead time at the constraint.
	ConstraintLeadTimeWeeks float64 `json:"constraint_lead_time_weeks" api:"required"`
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
	FulfillmentPolicy SchedulePolicyFulfillmentPolicy `json:"fulfillment_policy" api:"required"`
	// Annual cost of holding one unit.
	//
	// Null unless the caller holds `costs:read`; customer and supplier portal users
	// never see it.
	HoldingCost float64 `json:"holding_cost" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// What the constraint stage holds at its peak: its buffer plus a whole campaign.
	MaxGreigeInventory float64 `json:"max_greige_inventory" api:"required"`
	// Stock on hand at the constraint plus everything downstream of it.
	OnHandEchelon float64 `json:"on_hand_echelon" api:"required"`
	// Stock sitting at the constraint stage on its own.
	OnHandGreige float64 `json:"on_hand_greige" api:"required"`
	// Ceiling on how far ahead this item is built.
	OrderUpTo float64 `json:"order_up_to" api:"required"`
	// Which rule decided that policy: the item itself, its product line, or the
	// account default.
	//
	// Any of "item", "product_line", "account_default".
	PolicySource SchedulePolicyPolicySource `json:"policy_source" api:"required"`
	// Stock position at which a campaign is triggered.
	ReorderPoint float64 `json:"reorder_point" api:"required"`
	// Buffer held as finished goods.
	SafetyStockDownstream float64 `json:"safety_stock_downstream" api:"required"`
	// Buffer held at the constraint, pooled across the finished goods it feeds.
	SafetyStockPrimary float64 `json:"safety_stock_primary" api:"required"`
	// How long one unit occupies the constraint.
	SecondsPerUnit float64 `json:"seconds_per_unit" api:"required"`
	// Cost of one changeover, used as the setup cost in the lot-size calculation.
	//
	// Null unless the caller holds `costs:read`; customer and supplier portal users
	// never see it.
	SetupCost float64 `json:"setup_cost" api:"required"`
	// SKU of the item.
	SKU string `json:"sku" api:"required"`
	// Standard cost per unit.
	//
	// Null unless the caller holds `costs:read`; customer and supplier portal users
	// never see it.
	UnitCost float64 `json:"unit_cost" api:"required"`
	// Demand used for planning, per week.
	WeeklyDemand float64 `json:"weekly_demand" api:"required"`
	// Weeks of demand the current stock covers.
	WeeksOfCover float64 `json:"weeks_of_cover" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AbcClass                respjson.Field
		AnnualDemand            respjson.Field
		AnnualRunHours          respjson.Field
		AverageGreigeInventory  respjson.Field
		ConstraintLeadTimeWeeks respjson.Field
		EoqUnits                respjson.Field
		FinishLeadTimeWeeks     respjson.Field
		FirmDemandUnits         respjson.Field
		ForecastDemandUnits     respjson.Field
		FulfillmentPolicy       respjson.Field
		HoldingCost             respjson.Field
		Item                    respjson.Field
		MaxGreigeInventory      respjson.Field
		OnHandEchelon           respjson.Field
		OnHandGreige            respjson.Field
		OrderUpTo               respjson.Field
		PolicySource            respjson.Field
		ReorderPoint            respjson.Field
		SafetyStockDownstream   respjson.Field
		SafetyStockPrimary      respjson.Field
		SecondsPerUnit          respjson.Field
		SetupCost               respjson.Field
		SKU                     respjson.Field
		UnitCost                respjson.Field
		WeeklyDemand            respjson.Field
		WeeksOfCover            respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SchedulePolicy) RawJSON() string { return r.JSON.raw }
func (r *SchedulePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ABC class by share of constraint run hours.
//
// - `a`: consumes the largest share of constraint capacity.
// - `b`: moderate constraint consumption.
// - `c`: consumes little constraint capacity.
type SchedulePolicyAbcClass string

const (
	SchedulePolicyAbcClassA SchedulePolicyAbcClass = "a"
	SchedulePolicyAbcClassB SchedulePolicyAbcClass = "b"
	SchedulePolicyAbcClassC SchedulePolicyAbcClass = "c"
)

// How this item was planned.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer, so its safety stocks and reorder point are all zero.
type SchedulePolicyFulfillmentPolicy string

const (
	SchedulePolicyFulfillmentPolicyMakeToStock SchedulePolicyFulfillmentPolicy = "make_to_stock"
	SchedulePolicyFulfillmentPolicyMakeToOrder SchedulePolicyFulfillmentPolicy = "make_to_order"
)

// Which rule decided that policy: the item itself, its product line, or the
// account default.
type SchedulePolicyPolicySource string

const (
	SchedulePolicyPolicySourceItem           SchedulePolicyPolicySource = "item"
	SchedulePolicyPolicySourceProductLine    SchedulePolicyPolicySource = "product_line"
	SchedulePolicyPolicySourceAccountDefault SchedulePolicyPolicySource = "account_default"
)

// An item's projected stock position across the horizon.
type ScheduleProjection struct {
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Projected stock at the end of each week of the horizon.
	OnHandByWeek []float64 `json:"on_hand_by_week" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item         respjson.Field
		OnHandByWeek respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleProjection) RawJSON() string { return r.JSON.raw }
func (r *ScheduleProjection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleActionPreviewParams struct {
	// Request to preview a production schedule.
	PreviewProductionScheduleRequest PreviewProductionScheduleRequestParam
	paramObj
}

func (r OperationProductionScheduleActionPreviewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreviewProductionScheduleRequest)
}
func (r *OperationProductionScheduleActionPreviewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleActionPreviewRegenerateParams struct {
	// Request to see what a re-solve would change.
	PreviewRegenerateProductionScheduleRequest PreviewRegenerateProductionScheduleRequestParam
	paramObj
}

func (r OperationProductionScheduleActionPreviewRegenerateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreviewRegenerateProductionScheduleRequest)
}
func (r *OperationProductionScheduleActionPreviewRegenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleActionRegenerateParams struct {
	// Request to re-solve a draft in place.
	RegenerateProductionScheduleRequest RegenerateProductionScheduleRequestParam
	paramObj
}

func (r OperationProductionScheduleActionRegenerateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RegenerateProductionScheduleRequest)
}
func (r *OperationProductionScheduleActionRegenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationProductionScheduleActionReleaseWeekParams struct {
	// Request to release one week of a production schedule to the floor.
	ReleaseProductionScheduleWeekRequest ReleaseProductionScheduleWeekRequestParam
	paramObj
}

func (r OperationProductionScheduleActionReleaseWeekParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ReleaseProductionScheduleWeekRequest)
}
func (r *OperationProductionScheduleActionReleaseWeekParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
