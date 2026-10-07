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

// List and manage order discounts.
//
// SaleOrderDiscountService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleOrderDiscountService] method instead.
type SaleOrderDiscountService struct {
	options []option.RequestOption
	// List and manage order discounts.
	Actions SaleOrderDiscountActionService
}

// NewSaleOrderDiscountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleOrderDiscountService(opts ...option.RequestOption) (r SaleOrderDiscountService) {
	r = SaleOrderDiscountService{}
	r.options = opts
	r.Actions = NewSaleOrderDiscountActionService(opts...)
	return
}

// Creates an order discount that buyers can then redeem on a sales order by its
// code.
//
// The code must be unique within your account; reusing a code that another
// discount already holds returns a conflict error. Creating the discount does not
// apply it to anything — a discount only affects an order once that order
// references it.
//
// This endpoint requires the permission: `discounts:create`.
func (r *SaleOrderDiscountService) New(ctx context.Context, body SaleOrderDiscountNewParams, opts ...option.RequestOption) (res *OrderDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/order-discounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns an order discount by ID.
//
// This endpoint requires the permission: `discounts:read`.
func (r *SaleOrderDiscountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OrderDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/order-discounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Partially updates an order discount.
//
// Only the fields you send are changed; the rest keep their current values.
// Changing `code` to one another discount already holds returns a conflict error.
// Edits apply to future orders only — orders that already used this discount keep
// the reduction they were given.
//
// This endpoint requires the permission: `discounts:update`.
func (r *SaleOrderDiscountService) Update(ctx context.Context, id string, body SaleOrderDiscountUpdateParams, opts ...option.RequestOption) (res *OrderDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/order-discounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of the order discounts defined for the current account,
// newest first.
//
// Pass `q` to narrow the list to discounts whose name or code contains the search
// text.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `discounts:read`.
func (r *SaleOrderDiscountService) List(ctx context.Context, query SaleOrderDiscountListParams, opts ...option.RequestOption) (res *ListOrderDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/order-discounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an order discount and returns it as it was just before deletion.
//
// Deletion is permanent; further requests against the deleted ID return an error.
//
// The code can no longer be redeemed, but sales orders that already used the
// discount keep the reduction that was applied to them; their totals are not
// recalculated.
//
// This endpoint requires the permission: `discounts:delete`.
func (r *SaleOrderDiscountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OrderDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/order-discounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create an order discount.
//
// The properties Code, DiscountType, Name are required.
type CreateOrderDiscountRequestParam struct {
	// The code a buyer enters to apply this discount to an order.
	//
	// Codes are unique within your account and are compared without regard to letter
	// case, so `SAVE10` collides with `save10`.
	Code string `json:"code" api:"required"`
	// How the discount is calculated.
	//
	// - `percentage`: the order total is reduced by the fraction in `percentage`.
	// - `amount`: the order total is reduced by the flat amount in `amount`.
	//
	// Any of "percentage", "amount".
	DiscountType CreateOrderDiscountRequestDiscountType `json:"discount_type,omitzero" api:"required"`
	// Display name of the discount.
	Name string `json:"name" api:"required"`
	// The flat amount to take off the order total, as a decimal string.
	//
	// Only read when `discount_type` is `amount`. Leaving it out stores `0`, which
	// produces a discount that takes nothing off.
	Amount param.Opt[string] `json:"amount,omitzero" format:"decimal"`
	// The fraction of the order total to take off, as a decimal string.
	//
	// This is a multiplier, not a whole percent: send `0.1` to take 10% off. Only read
	// when `discount_type` is `percentage`. Leaving it out stores `0`, which produces
	// a discount that takes nothing off.
	Percentage param.Opt[string] `json:"percentage,omitzero" format:"decimal"`
	paramObj
}

func (r CreateOrderDiscountRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateOrderDiscountRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateOrderDiscountRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the discount is calculated.
//
// - `percentage`: the order total is reduced by the fraction in `percentage`.
// - `amount`: the order total is reduced by the flat amount in `amount`.
type CreateOrderDiscountRequestDiscountType string

const (
	CreateOrderDiscountRequestDiscountTypePercentage CreateOrderDiscountRequestDiscountType = "percentage"
	CreateOrderDiscountRequestDiscountTypeAmount     CreateOrderDiscountRequestDiscountType = "amount"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListOrderDiscount struct {
	// Resources in this page.
	Data []OrderDiscount `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListOrderDiscountObject `json:"object" api:"required"`
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
func (r ListOrderDiscount) RawJSON() string { return r.JSON.raw }
func (r *ListOrderDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListOrderDiscountObject string

const (
	ListOrderDiscountObjectList ListOrderDiscountObject = "list"
)

// A discount code that can be applied to a sales order.
//
// An order discount reduces the order total by either a percentage or a fixed
// amount, depending on `discount_type`. The reduction is capped at the order total
// and rounded to the nearest cent.
type OrderDiscount struct {
	// Order discount ID.
	ID string `json:"id" api:"required"`
	// The flat amount taken off the order total, as a decimal string.
	//
	// Only read when `discount_type` is `amount`.
	Amount string `json:"amount" api:"required" format:"decimal"`
	// The code a buyer enters to apply this discount to an order.
	//
	// Codes are unique within your account and are matched without regard to letter
	// case.
	Code string `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// How the discount is calculated.
	//
	// - `percentage`: the order total is reduced by the fraction in `percentage`.
	// - `amount`: the order total is reduced by the flat amount in `amount`.
	//
	// Any of "percentage", "amount".
	DiscountType OrderDiscountDiscountType `json:"discount_type" api:"required"`
	// Display name of the discount.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "order_discount".
	Object OrderDiscountObject `json:"object" api:"required"`
	// How many sales orders this discount has been applied to, across all buyers.
	OrderCount int64 `json:"order_count" api:"required"`
	// The fraction of the order total taken off, as a decimal string.
	//
	// This is a multiplier, not a whole percent: `0.1` takes 10% off. Only read when
	// `discount_type` is `percentage`.
	Percentage string `json:"percentage" api:"required" format:"decimal"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		Code         respjson.Field
		CreatedAt    respjson.Field
		DiscountType respjson.Field
		Name         respjson.Field
		Object       respjson.Field
		OrderCount   respjson.Field
		Percentage   respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderDiscount) RawJSON() string { return r.JSON.raw }
func (r *OrderDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the discount is calculated.
//
// - `percentage`: the order total is reduced by the fraction in `percentage`.
// - `amount`: the order total is reduced by the flat amount in `amount`.
type OrderDiscountDiscountType string

const (
	OrderDiscountDiscountTypePercentage OrderDiscountDiscountType = "percentage"
	OrderDiscountDiscountTypeAmount     OrderDiscountDiscountType = "amount"
)

// Resource type identifier.
type OrderDiscountObject string

const (
	OrderDiscountObjectOrderDiscount OrderDiscountObject = "order_discount"
)

// Request to partially update an order discount.
type UpdateOrderDiscountRequestParam struct {
	// The flat amount to take off the order total, as a decimal string.
	//
	// Only read when `discount_type` is `amount`.
	Amount param.Opt[string] `json:"amount,omitzero" format:"decimal"`
	// The code a buyer enters to apply this discount to an order.
	//
	// Codes are unique within your account and are compared without regard to letter
	// case.
	Code param.Opt[string] `json:"code,omitzero"`
	// Display name of the discount.
	Name param.Opt[string] `json:"name,omitzero"`
	// The fraction of the order total to take off, as a decimal string.
	//
	// This is a multiplier, not a whole percent: send `0.1` to take 10% off. Only read
	// when `discount_type` is `percentage`.
	Percentage param.Opt[string] `json:"percentage,omitzero" format:"decimal"`
	// How the discount is calculated.
	//
	// - `percentage`: the order total is reduced by the fraction in `percentage`.
	// - `amount`: the order total is reduced by the flat amount in `amount`.
	//
	// Switching the type does not move the stored figure across, so send the matching
	// `percentage` or `amount` in the same request or the discount will take nothing
	// off.
	//
	// Any of "percentage", "amount".
	DiscountType UpdateOrderDiscountRequestDiscountType `json:"discount_type,omitzero"`
	paramObj
}

func (r UpdateOrderDiscountRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateOrderDiscountRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateOrderDiscountRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the discount is calculated.
//
// - `percentage`: the order total is reduced by the fraction in `percentage`.
// - `amount`: the order total is reduced by the flat amount in `amount`.
//
// Switching the type does not move the stored figure across, so send the matching
// `percentage` or `amount` in the same request or the discount will take nothing
// off.
type UpdateOrderDiscountRequestDiscountType string

const (
	UpdateOrderDiscountRequestDiscountTypePercentage UpdateOrderDiscountRequestDiscountType = "percentage"
	UpdateOrderDiscountRequestDiscountTypeAmount     UpdateOrderDiscountRequestDiscountType = "amount"
)

type SaleOrderDiscountNewParams struct {
	// Request to create an order discount.
	CreateOrderDiscountRequest CreateOrderDiscountRequestParam
	paramObj
}

func (r SaleOrderDiscountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateOrderDiscountRequest)
}
func (r *SaleOrderDiscountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleOrderDiscountUpdateParams struct {
	// Request to partially update an order discount.
	UpdateOrderDiscountRequest UpdateOrderDiscountRequestParam
	paramObj
}

func (r SaleOrderDiscountUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateOrderDiscountRequest)
}
func (r *SaleOrderDiscountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleOrderDiscountListParams struct {
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

// URLQuery serializes [SaleOrderDiscountListParams]'s query parameters as
// `url.Values`.
func (r SaleOrderDiscountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
