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

// List and manage sales targets for account users.
//
// SaleAccountUserSalesTargetService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAccountUserSalesTargetService] method instead.
type SaleAccountUserSalesTargetService struct {
	options []option.RequestOption
}

// NewSaleAccountUserSalesTargetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSaleAccountUserSalesTargetService(opts ...option.RequestOption) (r SaleAccountUserSalesTargetService) {
	r = SaleAccountUserSalesTargetService{}
	r.options = opts
	return
}

// Creates a revenue goal for a sales rep covering a given period.
//
// The sales rep must be an active account user in your account, otherwise the
// request returns a not-found error. Periods are not checked for overlap, so a rep
// can hold several targets covering the same dates; use the upsert endpoint to
// change an existing target rather than adding another.
//
// This endpoint requires the permission: `sales_targets:create`.
func (r *SaleAccountUserSalesTargetService) New(ctx context.Context, id string, body SaleAccountUserSalesTargetNewParams, opts ...option.RequestOption) (res *SalesTarget, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-users/%s/sales-targets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Creates or updates a sales rep's revenue goal at an ID you choose.
//
// If no target with the given ID exists, one is created with the supplied dates,
// amount, and unit. If it already exists, only the amount value is updated — the
// dates and unit are left unchanged, so raising or lowering a goal mid-period is
// the intended use. The sales rep must be an active account user in your account,
// and an existing target must be filed under that rep, otherwise the request
// returns a not-found error.
//
// This endpoint requires the permission: `sales_targets:update`.
func (r *SaleAccountUserSalesTargetService) Update(ctx context.Context, targetID string, params SaleAccountUserSalesTargetUpdateParams, opts ...option.RequestOption) (res *SalesTarget, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-users/%s/sales-targets/%s", url.PathEscape(params.ID), url.PathEscape(targetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the revenue goals set for one sales rep, most recent period first.
//
// This endpoint does not support cursor pagination; passing a `cursor` returns a
// validation error, and the response carries no page cursors. Requesting targets
// for someone who is not an active account user in your account returns a
// not-found error.
//
// Pass `q` to narrow the list to targets whose ID or goal amount contains the
// search text.
//
// This endpoint requires the permission: `sales_targets:read`.
func (r *SaleAccountUserSalesTargetService) List(ctx context.Context, id string, query SaleAccountUserSalesTargetListParams, opts ...option.RequestOption) (res *ListSalesTarget, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-users/%s/sales-targets", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to create a sales target.
//
// The properties AmountUnitID, AmountValue, EndsAt, StartsAt are required.
type CreateSalesTargetRequestParam struct {
	// The unit the goal is denominated in, typically a currency unit. It must be one
	// of your account's units or a unit every account shares.
	AmountUnitID string `json:"amount_unit_id" api:"required"`
	// The revenue goal for the period, as a decimal string (e.g. `50000.00`).
	AmountValue string `json:"amount_value" api:"required" format:"decimal"`
	// End of the period the target applies to. It may not be before `starts_at`.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// Start of the period the target applies to (inclusive).
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	paramObj
}

func (r CreateSalesTargetRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesTargetRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesTargetRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListSalesTarget struct {
	// Resources in this page.
	Data []SalesTarget `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesTargetObject `json:"object" api:"required"`
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
func (r ListSalesTarget) RawJSON() string { return r.JSON.raw }
func (r *ListSalesTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesTargetObject string

const (
	ListSalesTargetObjectList ListSalesTargetObject = "list"
)

// A revenue goal assigned to a sales rep for a specific time period.
type SalesTarget struct {
	// Sales target ID.
	ID string `json:"id" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	Amount Quantity `json:"amount" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// End of the period this target applies to (e.g. the close of a quarter).
	EndAt time.Time `json:"end_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "sales_target".
	Object SalesTargetObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	SalesRep Entity `json:"sales_rep" api:"required"`
	// Start of the period this target applies to (inclusive).
	StartAt time.Time `json:"start_at" api:"required" format:"date-time"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		CreatedAt   respjson.Field
		EndAt       respjson.Field
		Object      respjson.Field
		SalesRep    respjson.Field
		StartAt     respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesTarget) RawJSON() string { return r.JSON.raw }
func (r *SalesTarget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesTargetObject string

const (
	SalesTargetObjectSalesTarget SalesTargetObject = "sales_target"
)

// Request to create or update a sales target.
//
// The properties AmountUnitID, AmountValue, EndsAt, StartsAt are required.
type UpsertSalesTargetRequestParam struct {
	// The unit the goal is denominated in, typically a currency unit. It must be one
	// of your account's units or a unit every account shares.
	//
	// Only applied when creating a new target; the unit on an existing target is not
	// changed.
	AmountUnitID string `json:"amount_unit_id" api:"required"`
	// The revenue goal for the period, as a decimal string (e.g. `75000.00`).
	//
	// This is the only value an existing target accepts; everything else on it stays
	// as it was.
	AmountValue string `json:"amount_value" api:"required" format:"decimal"`
	// End of the period the target applies to. It may not be before `starts_at`.
	//
	// Only applied when creating a new target; the dates on an existing target are not
	// changed.
	EndsAt time.Time `json:"ends_at" api:"required" format:"date-time"`
	// Start of the period the target applies to (inclusive).
	//
	// Only applied when creating a new target; the dates on an existing target are not
	// changed.
	StartsAt time.Time `json:"starts_at" api:"required" format:"date-time"`
	paramObj
}

func (r UpsertSalesTargetRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertSalesTargetRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertSalesTargetRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountUserSalesTargetNewParams struct {
	// Request to create a sales target.
	CreateSalesTargetRequest CreateSalesTargetRequestParam
	paramObj
}

func (r SaleAccountUserSalesTargetNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSalesTargetRequest)
}
func (r *SaleAccountUserSalesTargetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountUserSalesTargetUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Request to create or update a sales target.
	UpsertSalesTargetRequest UpsertSalesTargetRequestParam
	paramObj
}

func (r SaleAccountUserSalesTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpsertSalesTargetRequest)
}
func (r *SaleAccountUserSalesTargetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountUserSalesTargetListParams struct {
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
	paramObj
}

// URLQuery serializes [SaleAccountUserSalesTargetListParams]'s query parameters as
// `url.Values`.
func (r SaleAccountUserSalesTargetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
