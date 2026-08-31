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
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and export inventory change logs.
//
// OperationInventoryChangeLogService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationInventoryChangeLogService] method instead.
type OperationInventoryChangeLogService struct {
	options []option.RequestOption
	// List and export inventory change logs.
	Actions OperationInventoryChangeLogActionService
}

// NewOperationInventoryChangeLogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationInventoryChangeLogService(opts ...option.RequestOption) (r OperationInventoryChangeLogService) {
	r = OperationInventoryChangeLogService{}
	r.options = opts
	r.Actions = NewOperationInventoryChangeLogActionService(opts...)
	return
}

// Returns an inventory change log by ID.
//
// This endpoint requires the permission: `inventory_logs:read`.
func (r *OperationInventoryChangeLogService) Get(ctx context.Context, id string, query OperationInventoryChangeLogGetParams, opts ...option.RequestOption) (res *InventoryChangeLog, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/inventory-change-logs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of inventory change logs, newest first.
//
// Filters combine with AND, while the values within a single filter combine with
// OR. The `q` search term matches changes affecting items whose SKU contains it,
// as a case-insensitive substring.
//
// This endpoint requires the permission: `inventory_logs:read`.
func (r *OperationInventoryChangeLogService) List(ctx context.Context, query OperationInventoryChangeLogListParams, opts ...option.RequestOption) (res *ListInventoryChangeLog, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/inventory-change-logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A record of a single change to an item's on-hand inventory.
//
// Every inventory movement — production scans, manual user adjustments, and
// automatic system actions — produces one entry, forming an audit trail of how
// on-hand quantities changed over time.
type InventoryChangeLog struct {
	// Inventory change log ID.
	ID string `json:"id" api:"required"`
	// Action that produced this inventory change.
	//
	//   - `scan`: change driven by a scan, typically a production step.
	//   - `user_action`: change made manually by a user.
	//   - `system_action`: change made automatically by the system.
	//   - `user_correction`: manual adjustment a user made to correct an inventory
	//     discrepancy.
	//
	// Any of "scan", "user_action", "system_action", "user_correction".
	ActionType InventoryChangeLogActionType `json:"action_type" api:"required"`
	// Timestamp when this change was recorded.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "inventory_change_log".
	Object InventoryChangeLogObject `json:"object" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	Quantity Quantity `json:"quantity" api:"required"`
	// A station on the production floor where operators scan batches to perform a
	// batch operation, such as initializing or moving a batch.
	ResponsibleScanningStation ScanningStation `json:"responsible_scanning_station" api:"required"`
	// A user's global profile, shared across every account they belong to.
	//
	// Account-specific settings (status, role, department) live on the account user
	// resource that links the user to each account.
	ResponsibleUser User `json:"responsible_user" api:"required"`
	// Timestamp when this record was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		ActionType                 respjson.Field
		CreatedAt                  respjson.Field
		Item                       respjson.Field
		Object                     respjson.Field
		Quantity                   respjson.Field
		ResponsibleScanningStation respjson.Field
		ResponsibleUser            respjson.Field
		UpdatedAt                  respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InventoryChangeLog) RawJSON() string { return r.JSON.raw }
func (r *InventoryChangeLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Action that produced this inventory change.
//
//   - `scan`: change driven by a scan, typically a production step.
//   - `user_action`: change made manually by a user.
//   - `system_action`: change made automatically by the system.
//   - `user_correction`: manual adjustment a user made to correct an inventory
//     discrepancy.
type InventoryChangeLogActionType string

const (
	InventoryChangeLogActionTypeScan           InventoryChangeLogActionType = "scan"
	InventoryChangeLogActionTypeUserAction     InventoryChangeLogActionType = "user_action"
	InventoryChangeLogActionTypeSystemAction   InventoryChangeLogActionType = "system_action"
	InventoryChangeLogActionTypeUserCorrection InventoryChangeLogActionType = "user_correction"
)

// Resource type identifier.
type InventoryChangeLogObject string

const (
	InventoryChangeLogObjectInventoryChangeLog InventoryChangeLogObject = "inventory_change_log"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListInventoryChangeLog struct {
	// Resources in this page.
	Data []InventoryChangeLog `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListInventoryChangeLogObject `json:"object" api:"required"`
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
func (r ListInventoryChangeLog) RawJSON() string { return r.JSON.raw }
func (r *ListInventoryChangeLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListInventoryChangeLogObject string

const (
	ListInventoryChangeLogObjectList ListInventoryChangeLogObject = "list"
)

type OperationInventoryChangeLogGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "responsible_user", "responsible_scanning_station".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationInventoryChangeLogGetParams]'s query parameters as
// `url.Values`.
func (r OperationInventoryChangeLogGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationInventoryChangeLogListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Restricts results to change logs created on or before this timestamp.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Restricts results to change logs created on or after this timestamp.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Restricts results to these action types.
	//
	// Any of "scan", "user_action", "system_action", "user_correction".
	ActionTypes []string `query:"action_types,omitzero" json:"-"`
	// Restricts results to changes made by these users.
	//
	// Changes that were recorded without a responsible user are excluded whenever this
	// filter is set.
	ChangedByUserIDs []string `query:"changed_by_user_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "responsible_user", "responsible_scanning_station".
	Include []string `query:"include,omitzero" json:"-"`
	// Restricts results to changes affecting these items.
	ItemIDs []string `query:"item_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationInventoryChangeLogListParams]'s query parameters
// as `url.Values`.
func (r OperationInventoryChangeLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
