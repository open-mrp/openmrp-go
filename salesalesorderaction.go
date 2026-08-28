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

// List, view, create, update, and delete sales orders.
//
// SaleSalesOrderActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderActionService] method instead.
type SaleSalesOrderActionService struct {
	options []option.RequestOption
}

// NewSaleSalesOrderActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleSalesOrderActionService(opts ...option.RequestOption) (r SaleSalesOrderActionService) {
	r = SaleSalesOrderActionService{}
	r.options = opts
	return
}

// Deletes multiple sales orders in a single atomic operation.
//
// Each order is torn down exactly as it would be by deleting it on its own.
// Fulfilled orders cannot be deleted; if any requested order fails this check, no
// orders are deleted.
//
// This endpoint requires the permission: `sales_orders:delete`.
func (r *SaleSalesOrderActionService) BulkDelete(ctx context.Context, body SaleSalesOrderActionBulkDeleteParams, opts ...option.RequestOption) (res *SaleSalesOrderActionBulkDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/actions/bulk-delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Closes a sales order, transitioning it from `issued` to `fulfilled`.
//
// Stamps the order's completion timestamp and closes its pick, packing every pick
// line that is still open so the pick reads as complete alongside the order. Only
// an order in `issued` can be closed, and once it is fulfilled it can no longer be
// deleted, nor can its lines be removed, until it is reopened.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Close(ctx context.Context, id string, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/close", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Creates a production run from a sales order.
//
// Walks the production flow behind each item-backed line to work out what actually
// has to be made, then creates one batch for each item that is produced directly
// from raw materials, sized to cover every line that needs it. Reserves the
// material inventory those batches consume and links the run to the order. The
// caller becomes the run's responsible user. An order can have at most one
// production run, and a line whose item has no production flow contributes no
// batches.
//
// This endpoint requires the permission: `production_runs:create`.
func (r *SaleSalesOrderActionService) NewProductionRun(ctx context.Context, id string, body SaleSalesOrderActionNewProductionRunParams, opts ...option.RequestOption) (res *ProductionRun, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/create-production-run", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Issues a sales order, transitioning it from `estimate` to `issued`.
//
// Issuing commits the order for fulfillment: a pick is created for the order's
// sale lines and inventory is reserved for each line tied to an inventory item.
// Only an order still in `estimate` can be issued.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Issue(ctx context.Context, id string, body SaleSalesOrderActionIssueParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/issue", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Reopens a sales order, transitioning it from `fulfilled` back to `issued`.
//
// Clears the order's completion timestamp and reopens its pick, unpacking every
// pick line that is not yet fully picked so the outstanding work can be resumed;
// lines already picked in full stay packed. Only an order in `fulfilled` can be
// reopened.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Open(ctx context.Context, id string, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/open", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Previews the ship-by date a set of commitment inputs would produce, without
// creating or changing anything.
//
// Runs the same resolution an order runs when it is issued: a promised delivery
// date has the customer's receiving days, the carrier's transit, and the plant's
// shipping days worked back through it, while a lead time or a pinned ship date is
// snapped onto the next earlier day the plant ships. The returned steps are that
// derivation in order, so a caller can show why a date is what it is rather than
// restating the rules.
//
// At most one of `promised_at`, `lead_time_override_days`, and
// `ship_by_override_date` may be set; they are alternative answers to the same
// question.
//
// Advisory rather than binding. Carrier transit comes from a lane cache warmed in
// the background, so a lane nobody has shipped yet quotes against the service
// level's default or against no transit at all, and the date stamped at issue may
// differ once the lane has been rated.
//
// This endpoint requires the permission: `sales_orders:read`.
func (r *SaleSalesOrderActionService) QuoteCommitment(ctx context.Context, body SaleSalesOrderActionQuoteCommitmentParams, opts ...option.RequestOption) (res *QuoteSalesOrderCommitmentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/actions/quote-commitment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Re-estimates the freight (shipping) charge for an order using the latest carrier
// rates.
//
// Computes what the order's freight charge would be from its current ship-to
// address, carrier, service level, and line items — applying the same
// freight-exemption, flat-rate, and live carrier-rate logic used when the order is
// created. The order is not modified: the returned amount is a quote to review,
// and callers apply it by updating the order's shipping line. Use this to refresh
// freight after changing the address or line items, or at any time to re-price
// against current rates.
//
// This endpoint requires the permission: `sales_orders:read`.
func (r *SaleSalesOrderActionService) QuoteFreight(ctx context.Context, id string, opts ...option.RequestOption) (res *QuoteSalesOrderFreightResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/quote-freight", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Unissues a sales order, transitioning it from `issued` back to `estimate`.
//
// Deletes the order's pick, discarding any picking progress recorded against it,
// and releases the inventory reserved when the order was issued. Only an order in
// `issued` can be unissued.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderActionService) Unissue(ctx context.Context, id string, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/actions/unissue", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Request to bulk delete sales orders.
//
// The property SalesOrderIDs is required.
type BulkDeleteSalesOrdersRequestParam struct {
	// IDs of the sales orders to delete.
	SalesOrderIDs []string `json:"sales_order_ids,omitzero" api:"required"`
	paramObj
}

func (r BulkDeleteSalesOrdersRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkDeleteSalesOrdersRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkDeleteSalesOrdersRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CommitmentQuoteStep is one rule's contribution to a previewed ship-by date.
//
// Returned as an ordered list so a caller can show why a date is what it is
// without reimplementing the arithmetic, and so the explanation cannot drift from
// the calculation that produced it.
type CommitmentQuoteStep struct {
	// Which rule applied.
	//
	// Any of "basis", "receive_calendar", "carrier_transit", "ship_calendar",
	// "pickup_cutoff".
	Code CommitmentQuoteStepCode `json:"code" api:"required"`
	// Where the running date stood after this rule.
	Date time.Time `json:"date" api:"required" format:"date-time"`
	// How far this rule pulled the date back. Zero means the rule applied and changed
	// nothing, which is worth showing: it says the date was already on an open day.
	DaysMoved int64 `json:"days_moved" api:"required"`
	// The rule's own parameter — where a transit estimate came from, or the cutoff
	// time applied. Null for a rule that takes none, rather than an empty string:
	// snapping onto an open day has no parameter to report.
	Detail string `json:"detail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Date        respjson.Field
		DaysMoved   respjson.Field
		Detail      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommitmentQuoteStep) RawJSON() string { return r.JSON.raw }
func (r *CommitmentQuoteStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which rule applied.
type CommitmentQuoteStepCode string

const (
	CommitmentQuoteStepCodeBasis           CommitmentQuoteStepCode = "basis"
	CommitmentQuoteStepCodeReceiveCalendar CommitmentQuoteStepCode = "receive_calendar"
	CommitmentQuoteStepCodeCarrierTransit  CommitmentQuoteStepCode = "carrier_transit"
	CommitmentQuoteStepCodeShipCalendar    CommitmentQuoteStepCode = "ship_calendar"
	CommitmentQuoteStepCodePickupCutoff    CommitmentQuoteStepCode = "pickup_cutoff"
)

// Request to issue a sales order.
//
// The property NotifyCustomer is required.
type IssueSalesOrderRequestParam struct {
	// Whether to notify the customer.
	//
	// When `true`, an order acknowledgement email with a PDF of the order is sent to
	// the acknowledgement contacts on the order and the order's
	// `acknowledgment_status` becomes `sent`. An order with no acknowledgement
	// contacts sends nothing and leaves its `acknowledgment_status` unchanged.
	NotifyCustomer bool `json:"notify_customer" api:"required"`
	paramObj
}

func (r IssueSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow IssueSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IssueSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A production run: the group of shop-floor batches that are executed together,
// tracked from the first batch scan through to completion.
type ProductionRun struct {
	// Production run ID.
	ID string `json:"id" api:"required"`
	// Number of batches currently recorded against this run.
	BatchCount int64 `json:"batch_count" api:"required"`
	// Time the run finished production.
	//
	// Set automatically once every batch in the run has been scanned or deleted. From
	// that point the run can no longer be updated and no further batches can be added
	// to it.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Production run number, unique per account.
	//
	// Assigned automatically at creation as the next sequential number for the
	// account; can be changed via update.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_run".
	Object ProductionRunObject `json:"object" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the `user`
	// sub-resource, which is shared across every account the user belongs to.
	ResponsibleUser AccountUser `json:"responsible_user" api:"required"`
	// Time the run started production.
	//
	// Set automatically the first time a batch in the run is scanned at a station.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		BatchCount      respjson.Field
		CompletedAt     respjson.Field
		CreatedAt       respjson.Field
		Number          respjson.Field
		Object          respjson.Field
		ResponsibleUser respjson.Field
		StartedAt       respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionRun) RawJSON() string { return r.JSON.raw }
func (r *ProductionRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionRunObject string

const (
	ProductionRunObjectProductionRun ProductionRunObject = "production_run"
)

// Request to preview the ship-by date a set of commitment inputs would produce.
type QuoteSalesOrderCommitmentRequestParam struct {
	// The buying account, used to resolve its lead time and receiving days.
	BuyerAccountID param.Opt[string] `json:"buyer_account_id,omitzero"`
	// Carrier for the shipment.
	CarrierID param.Opt[string] `json:"carrier_id,omitzero"`
	// When the order would be issued. Defaults to the date sales_order_id was issued
	// on, or to now for an order that has not been issued — a lead time is measured
	// from issue, so an order built today but issued next week commits to next week's
	// date, and re-committing one issued last week still counts from last week.
	IssuedAt param.Opt[time.Time] `json:"issued_at,omitzero" format:"date-time"`
	// Days between issue and the order being due to ship, in place of the customer's
	// standing lead time.
	LeadTimeOverrideDays param.Opt[int64] `json:"lead_time_override_days,omitzero"`
	// Date delivery would be promised to the customer.
	PromisedAt param.Opt[time.Time] `json:"promised_at,omitzero" format:"date-time"`
	// An existing order to preview against. Its customer, ship-to address, carrier,
	// and service level are used, and the commitment fields below replace whatever it
	// currently carries.
	//
	// Omit it to preview an order that has not been created yet, supplying the parts
	// directly.
	SalesOrderID param.Opt[string] `json:"sales_order_id,omitzero"`
	// Service level for the shipment, which the lane's transit estimate is keyed on.
	ServiceLevelID param.Opt[string] `json:"service_level_id,omitzero"`
	// The exact date the order would be due to ship.
	ShipByOverrideDate param.Opt[time.Time] `json:"ship_by_override_date,omitzero" format:"date-time"`
	// The ship-to address, which decides the destination timezone and the lane transit
	// is quoted on.
	ShipToAddressID param.Opt[string] `json:"ship_to_address_id,omitzero"`
	paramObj
}

func (r QuoteSalesOrderCommitmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow QuoteSalesOrderCommitmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteSalesOrderCommitmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The ship-by date a set of commitment inputs would produce, and how it was
// reached.
type QuoteSalesOrderCommitmentResponse struct {
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
	// Resource type identifier.
	//
	// Any of "sales_order_commitment_quote".
	Object QuoteSalesOrderCommitmentResponseObject `json:"object" api:"required"`
	// The derivation in order, one entry per rule that moved the date.
	Steps []CommitmentQuoteStep `json:"steps" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Commitment  respjson.Field
		Object      respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteSalesOrderCommitmentResponse) RawJSON() string { return r.JSON.raw }
func (r *QuoteSalesOrderCommitmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuoteSalesOrderCommitmentResponseObject string

const (
	QuoteSalesOrderCommitmentResponseObjectSalesOrderCommitmentQuote QuoteSalesOrderCommitmentResponseObject = "sales_order_commitment_quote"
)

// The freshly estimated freight charge for a sales order.
type QuoteSalesOrderFreightResponse struct {
	// Resource type identifier.
	//
	// Any of "sales_order_freight_quote".
	Object QuoteSalesOrderFreightResponseObject `json:"object" api:"required"`
	// A rate calculated on demand rather than stored.
	//
	// The same shape as a rate minus the fields only a persisted row can have: it
	// carries no ID and no timestamps because nothing was written. Used where a figure
	// is derived per request, such as an analysis comparing one customer's price
	// against the median other customers pay.
	UnitPrice ComputedRate `json:"unit_price" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteSalesOrderFreightResponse) RawJSON() string { return r.JSON.raw }
func (r *QuoteSalesOrderFreightResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuoteSalesOrderFreightResponseObject string

const (
	QuoteSalesOrderFreightResponseObjectSalesOrderFreightQuote QuoteSalesOrderFreightResponseObject = "sales_order_freight_quote"
)

type SaleSalesOrderActionBulkDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderActionBulkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderActionBulkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderActionBulkDeleteParams struct {
	// Request to bulk delete sales orders.
	BulkDeleteSalesOrdersRequest BulkDeleteSalesOrdersRequestParam
	paramObj
}

func (r SaleSalesOrderActionBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkDeleteSalesOrdersRequest)
}
func (r *SaleSalesOrderActionBulkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderActionNewProductionRunParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "responsible_user", "responsible_user.user".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderActionNewProductionRunParams]'s query
// parameters as `url.Values`.
func (r SaleSalesOrderActionNewProductionRunParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderActionIssueParams struct {
	// Request to issue a sales order.
	IssueSalesOrderRequest IssueSalesOrderRequestParam
	paramObj
}

func (r SaleSalesOrderActionIssueParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IssueSalesOrderRequest)
}
func (r *SaleSalesOrderActionIssueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderActionQuoteCommitmentParams struct {
	// Request to preview the ship-by date a set of commitment inputs would produce.
	QuoteSalesOrderCommitmentRequest QuoteSalesOrderCommitmentRequestParam
	paramObj
}

func (r SaleSalesOrderActionQuoteCommitmentParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.QuoteSalesOrderCommitmentRequest)
}
func (r *SaleSalesOrderActionQuoteCommitmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
