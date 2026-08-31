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

// List, view, pick, void, and pack picks and pick lines.
//
// OperationPickService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationPickService] method instead.
type OperationPickService struct {
	options []option.RequestOption
	// List, view, pick, void, and pack picks and pick lines.
	Actions OperationPickActionService
	// List, view, pick, void, and pack picks and pick lines.
	Lines OperationPickLineService
}

// NewOperationPickService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationPickService(opts ...option.RequestOption) (r OperationPickService) {
	r = OperationPickService{}
	r.options = opts
	r.Actions = NewOperationPickActionService(opts...)
	r.Lines = NewOperationPickLineService(opts...)
	return
}

// Returns a pick by ID.
//
// This endpoint requires the permission: `picks:read`.
func (r *OperationPickService) Get(ctx context.Context, id string, query OperationPickGetParams, opts ...option.RequestOption) (res *Pick, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of picks, soonest ship-by date first.
//
// The `q` search term matches the pick number (which is the order number) and the
// customer PO number. To narrow by customer, use `customer_ids` rather than
// searching for a customer name.
//
// This endpoint requires the permissions: `picks:read`, `customers:read`,
// `suppliers:read`.
func (r *OperationPickService) List(ctx context.Context, query OperationPickListParams, opts ...option.RequestOption) (res *ListPick, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/picks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPick struct {
	// Resources in this page.
	Data []Pick `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPickObject `json:"object" api:"required"`
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
func (r ListPick) RawJSON() string { return r.JSON.raw }
func (r *ListPick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPickObject string

const (
	ListPickObjectList ListPickObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPickLine struct {
	// Resources in this page.
	Data []PickLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPickLineObject `json:"object" api:"required"`
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
func (r ListPickLine) RawJSON() string { return r.JSON.raw }
func (r *ListPickLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPickLineObject string

const (
	ListPickLineObjectList ListPickLineObject = "list"
)

// A warehouse picking task for a sales order, tracking the quantities to pull from
// inventory and pack for shipment.
//
// A pick is created automatically when a sales order is issued, with one line for
// each order line whose product is of type `sale` service, shipping, tax, credit
// and return lines are skipped — and nothing picked yet.
type Pick struct {
	// Pick ID.
	ID string `json:"id" api:"required"`
	// Commitment describes when a record is due to ship: what was asked for, what that
	// resolved to, and which rule decided.
	//
	// It is a generic, reusable sub-resource shared by anything carrying a ship-by
	// commitment — a sales order, the pick that fulfills it, or a preview of an order
	// that does not exist yet.
	//
	// The three inputs are alternative answers to the same question and at most one is
	// ever set; `lead_time_source` reports which of them, or which level of the
	// customer chain, produced the date. They are written flat on the create and
	// update bodies, the way a carrier is written as `carrier_id` and read back under
	// `freight`.
	Commitment Commitment `json:"commitment" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// CreatedBy describes who created a resource and their relationship to the account
	// that owns it.
	//
	// It is resolved from the resource's create audit event.
	CreatedBy CreatedBy `json:"created_by" api:"required"`
	// A business you sell to, with its contact details, default fulfillment settings,
	// and order policies.
	Customer Customer `json:"customer" api:"required"`
	// The customer's own purchase order number for the sales order this pick fulfills.
	CustomerPurchaseOrderNumber string `json:"customer_purchase_order_number" api:"required"`
	// Timestamp when the pick was finished.
	FinishedAt time.Time `json:"finished_at" api:"required" format:"date-time"`
	// Freight describes the carrier selection and freight billing for a record.
	//
	// It is a generic, reusable sub-resource shared by anything that carries shipping
	// configuration — a sales order, a purchase order, or a shipment.
	Freight Freight `json:"freight" api:"required"`
	// Timestamp of the most recent shipment sent (null until shipped).
	LastShippedAt time.Time `json:"last_shipped_at" api:"required" format:"date-time"`
	// Number of lines on this pick.
	LineCount int64 `json:"line_count" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lines ListPickLine `json:"lines" api:"required"`
	// Free-form note carried from the sales order this pick fulfills.
	Note string `json:"note" api:"required"`
	// Human-readable number that identifies the pick, distinct from the `id`.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "pick".
	Object PickObject `json:"object" api:"required"`
	// How urgently the pick should be worked.
	//
	// Any of "low", "normal", "high".
	Priority PickPriority `json:"priority" api:"required"`
	// Groups the records a pick sits between — the order it fulfills and the shipments
	// packed from it — and is returned only once at least one member has been
	// expanded.
	Related PickRelated `json:"related" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	ShipTo Address `json:"ship_to" api:"required"`
	// Progress through each fulfillment stage of a pick.
	Totals PickTotals `json:"totals" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		Commitment                  respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		Customer                    respjson.Field
		CustomerPurchaseOrderNumber respjson.Field
		FinishedAt                  respjson.Field
		Freight                     respjson.Field
		LastShippedAt               respjson.Field
		LineCount                   respjson.Field
		Lines                       respjson.Field
		Note                        respjson.Field
		Number                      respjson.Field
		Object                      respjson.Field
		Priority                    respjson.Field
		Related                     respjson.Field
		ShipTo                      respjson.Field
		Totals                      respjson.Field
		UpdatedAt                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pick) RawJSON() string { return r.JSON.raw }
func (r *Pick) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PickObject string

const (
	PickObjectPick PickObject = "pick"
)

// How urgently the pick should be worked.
type PickPriority string

const (
	PickPriorityLow    PickPriority = "low"
	PickPriorityNormal PickPriority = "normal"
	PickPriorityHigh   PickPriority = "high"
)

// A single line on a pick, tracking the quantity picked against one sales order
// line.
type PickLine struct {
	// Pick line ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "pick_line".
	Object PickLineObject `json:"object" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	OrderedQuantity Quantity `json:"ordered_quantity" api:"required"`
	// Timestamp when the line was packed.
	PackedAt time.Time `json:"packed_at" api:"required" format:"date-time"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	Quantity Quantity `json:"quantity" api:"required"`
	// A single line item on a sales order.
	SalesOrderLine SalesOrderLine `json:"sales_order_line" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Item            respjson.Field
		Object          respjson.Field
		OrderedQuantity respjson.Field
		PackedAt        respjson.Field
		Quantity        respjson.Field
		SalesOrderLine  respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PickLine) RawJSON() string { return r.JSON.raw }
func (r *PickLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PickLineObject string

const (
	PickLineObjectPickLine PickLineObject = "pick_line"
)

// Groups the records a pick sits between — the order it fulfills and the shipments
// packed from it — and is returned only once at least one member has been
// expanded.
type PickRelated struct {
	// Resource type identifier.
	//
	// Any of "pick_related".
	Object PickRelatedObject `json:"object" api:"required"`
	// Record is a lightweight reference to a business record — a sales order, purchase
	// order, pick, shipment, production run, invoice, etc.
	//
	// Like the `actor` and `entity` references, it carries just enough to identify and
	// label the referenced record without embedding its full resource. The `status`
	// and `metadata` fields hold type-specific detail that varies by the kind of
	// record referenced.
	SalesOrder Record `json:"sales_order" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Shipments ListRecord `json:"shipments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		SalesOrder  respjson.Field
		Shipments   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PickRelated) RawJSON() string { return r.JSON.raw }
func (r *PickRelated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PickRelatedObject string

const (
	PickRelatedObjectPickRelated PickRelatedObject = "pick_related"
)

// How far one fulfillment stage of a pick has progressed.
type PickStageTotal struct {
	// Progress as a fraction between 0 and 1.
	Completion float64 `json:"completion" api:"required"`
	// Resource type identifier.
	//
	// Any of "pick_stage_total".
	Object PickStageTotalObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completion  respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PickStageTotal) RawJSON() string { return r.JSON.raw }
func (r *PickStageTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PickStageTotalObject string

const (
	PickStageTotalObjectPickStageTotal PickStageTotalObject = "pick_stage_total"
)

// Progress through each fulfillment stage of a pick.
type PickTotals struct {
	// Resource type identifier.
	//
	// Any of "pick_totals".
	Object PickTotalsObject `json:"object" api:"required"`
	// How far one fulfillment stage of a pick has progressed.
	Packed PickStageTotal `json:"packed" api:"required"`
	// How far one fulfillment stage of a pick has progressed.
	Picked PickStageTotal `json:"picked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		Packed      respjson.Field
		Picked      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PickTotals) RawJSON() string { return r.JSON.raw }
func (r *PickTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PickTotalsObject string

const (
	PickTotalsObjectPickTotals PickTotalsObject = "pick_totals"
)

type OperationPickGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "created_by", "freight", "related.sales_order",
	// "related.shipments", "lines", "lines.item", "lines.sales_order_line",
	// "lines.sales_order_line.product", "lines.quantity", "lines.quantity.unit",
	// "lines.ordered_quantity", "lines.ordered_quantity.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationPickGetParams]'s query parameters as `url.Values`.
func (r OperationPickGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationPickListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Latest pick creation date to include, in `YYYY-MM-DD` format. Inclusive of the
	// date itself.
	EndsAt param.Opt[string] `query:"ends_at,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Earliest pick creation date to include, in `YYYY-MM-DD` format.
	StartsAt param.Opt[string] `query:"starts_at,omitzero" json:"-"`
	// Restricts results to picks whose customer belongs to any of these account
	// groups, matching the `type` on the customer.
	CustomerGroupIDs []string `query:"customer_group_ids,omitzero" json:"-"`
	// Restricts results to picks raised for any of these customers.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "created_by", "freight", "related.sales_order",
	// "related.shipments", "lines", "lines.item", "lines.sales_order_line",
	// "lines.sales_order_line.product", "lines.quantity", "lines.quantity.unit",
	// "lines.ordered_quantity", "lines.ordered_quantity.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Restricts results to picks with at least one line whose product belongs to any
	// of these product lines.
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	// Orders the results: `ship_by_date` puts the soonest delivery commitment first,
	// with picks whose order has no ship-by date last; `created_at` puts the newest
	// pick first.
	//
	// Any of "ship_by_date", "created_at".
	Sort OperationPickListParamsSort `query:"sort,omitzero" json:"-"`
	// Restricts results to picks in this state.
	//
	// - `open`: picks that have not been finished.
	// - `closed`: picks that have been finished.
	//
	// Any of "open", "closed".
	Status OperationPickListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationPickListParams]'s query parameters as
// `url.Values`.
func (r OperationPickListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Orders the results: `ship_by_date` puts the soonest delivery commitment first,
// with picks whose order has no ship-by date last; `created_at` puts the newest
// pick first.
type OperationPickListParamsSort string

const (
	OperationPickListParamsSortShipByDate OperationPickListParamsSort = "ship_by_date"
	OperationPickListParamsSortCreatedAt  OperationPickListParamsSort = "created_at"
)

// Restricts results to picks in this state.
//
// - `open`: picks that have not been finished.
// - `closed`: picks that have been finished.
type OperationPickListParamsStatus string

const (
	OperationPickListParamsStatusOpen   OperationPickListParamsStatus = "open"
	OperationPickListParamsStatusClosed OperationPickListParamsStatus = "closed"
)
