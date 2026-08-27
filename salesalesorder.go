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

// SaleSalesOrderService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderService] method instead.
type SaleSalesOrderService struct {
	options []option.RequestOption
	// List, view, create, update, and delete sales orders.
	Actions SaleSalesOrderActionService
	// List, view, create, update, and delete sales orders.
	Lines SaleSalesOrderLineService
}

// NewSaleSalesOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleSalesOrderService(opts ...option.RequestOption) (r SaleSalesOrderService) {
	r = SaleSalesOrderService{}
	r.options = opts
	r.Actions = NewSaleSalesOrderActionService(opts...)
	r.Lines = NewSaleSalesOrderLineService(opts...)
	return
}

// Creates a sales order in `estimate` status.
//
// The order number is assigned automatically, and a sales rep is auto-assigned
// when none is provided. Line prices and costs are resolved server-side from each
// product. A shipping line carrying the estimated freight charge is added to the
// order, plus a negative-priced discount line when an order discount is supplied.
// The order is not committed for fulfillment until it is issued.
//
// This endpoint requires the permission: `sales_orders:create`.
func (r *SaleSalesOrderService) New(ctx context.Context, params SaleSalesOrderNewParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a sales order by ID.
//
// This endpoint requires the permissions: `customers:read`, `suppliers:read`,
// `sales_orders:read`.
func (r *SaleSalesOrderService) Get(ctx context.Context, id string, query SaleSalesOrderGetParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a sales order.
//
// Changing the carrier, service level, or ship-to address propagates to the
// order's existing shipments, but never re-prices the freight line: request a
// fresh estimate from the quote-freight endpoint and apply it to the shipping line
// yourself. Order status is changed through the issue, unissue, close, and reopen
// actions instead of this endpoint.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderService) Update(ctx context.Context, id string, params SaleSalesOrderUpdateParams, opts ...option.RequestOption) (res *SalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of sales orders for the current account, newest first.
//
// A free-text search term (`q`) is matched as an exact value against the order
// number and the customer purchase order number, and still respects the other
// filters. Customer accounts calling this endpoint only ever see their own orders.
//
// This endpoint requires the permissions: `sales_orders:read`, `customers:read`,
// `suppliers:read`.
func (r *SaleSalesOrderService) List(ctx context.Context, query SaleSalesOrderListParams, opts ...option.RequestOption) (res *ListSalesOrder, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a sales order and all its related records.
//
// Removes the order's lines, pick, shipment and invoice lines, and email contacts,
// and releases any inventory it had reserved. Fulfilled orders cannot be deleted.
//
// This endpoint requires the permission: `sales_orders:delete`.
func (r *SaleSalesOrderService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleSalesOrderDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Creates a hosted payment checkout session for a sales order.
//
// Requires an active Stripe integration on the account and a customer that already
// exists in Stripe. The customer is charged a single amount covering every line on
// the order, including its freight and discount lines, and the checkout link is
// emailed to the address provided. Fails with a conflict if the order already has
// a payment.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderService) Checkout(ctx context.Context, id string, body SaleSalesOrderCheckoutParams, opts ...option.RequestOption) (res *CheckoutSalesOrderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/checkout", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Calculates the unit price for each line without creating an order.
//
// Use this to display prices to users as they build an order. Prices are computed
// server-side from the product's list price, contracted account prices, and
// applicable discounts — the same logic used when an order is created. Internal
// price overrides are not accepted here; the calculated price is always returned.
//
// This endpoint requires the permission: `sales_orders:read`.
func (r *SaleSalesOrderService) PriceQuote(ctx context.Context, body SaleSalesOrderPriceQuoteParams, opts ...option.RequestOption) (res *QuoteSalesOrderPricesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/price-quote"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists the statuses a sales order can be in.
//
// The statuses are platform-provided and the same for every account, so the result
// is small and stable enough to cache. Use it to label orders in your own
// interface; an order moves between statuses through its issue, unissue, close,
// and reopen actions rather than by being assigned a status.
func (r *SaleSalesOrderService) GetStatuses(ctx context.Context, query SaleSalesOrderGetStatusesParams, opts ...option.RequestOption) (res *ListSalesOrderStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/sales-orders/statuses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to create a checkout session for a sales order.
//
// The property Email is required.
type CheckoutSalesOrderRequestParam struct {
	// Email address to send the checkout link to.
	Email string `json:"email" api:"required"`
	paramObj
}

func (r CheckoutSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckoutSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckoutSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Checkout session result.
type CheckoutSalesOrderResponse struct {
	// URL of the hosted payment page where the customer completes the checkout.
	CheckoutURL string `json:"checkout_url" api:"required"`
	// Resource type identifier.
	//
	// Any of "checkout_sales_order".
	Object CheckoutSalesOrderResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckoutURL respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckoutSalesOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *CheckoutSalesOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CheckoutSalesOrderResponseObject string

const (
	CheckoutSalesOrderResponseObjectCheckoutSalesOrder CheckoutSalesOrderResponseObject = "checkout_sales_order"
)

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
type Commitment struct {
	// Days the customer's receiving calendar and the plant's shipping calendar pulled
	// the ship-by date back, beyond what carrier transit accounted for.
	//
	// Zero means every date along the way already fell on an open day. This is what
	// explains a ship-by date that is earlier than transit alone would suggest.
	CalendarAdjustmentDays int64 `json:"calendar_adjustment_days" api:"required"`
	// When freight leaving on the ship-by date would reach the customer: transit
	// walked forward from it and landed on a day their dock receives.
	//
	// Reported by the commitment preview, which is asked what a set of inputs would
	// produce and so computes the arrival too. A record carries the commitment it was
	// stamped with, not a projection, and leaves this null.
	EstimatedDeliveryDate time.Time `json:"estimated_delivery_date" api:"required" format:"date-time"`
	// Calendar days between issue and the ship-by date.
	LeadTimeDays int64 `json:"lead_time_days" api:"required"`
	// Days between issue and the ship-by date, set on this record alone in place of
	// the customer's standing lead time.
	LeadTimeOverrideDays int64 `json:"lead_time_override_days" api:"required"`
	// Which rule produced the ship-by date.
	//
	// Any of "customer", "parent_customer", "account_group", "account", "manual",
	// "order_lead_time", "order_ship_by".
	LeadTimeSource CommitmentLeadTimeSource `json:"lead_time_source" api:"required"`
	// Resource type identifier.
	//
	// Any of "commitment".
	Object CommitmentObject `json:"object" api:"required"`
	// Date delivery was promised to the customer, if one was committed.
	PromisedAt time.Time `json:"promised_at" api:"required" format:"date-time"`
	// When the record is contractually due to ship.
	//
	// Stamped at issue. With a promised delivery date, this is that date less the
	// carrier's transit for the order's lane and less any day the customer cannot
	// receive on — when the order has to leave to arrive when promised. Otherwise it
	// comes from a lead time, whether the order's own or the one on the customer, its
	// parent account, its account group, or the account.
	//
	// Always a day the plant actually ships on, whichever rule produced it, and
	// carries the plant's pickup cutoff as its time of day when the shipping calendar
	// sets one — the moment freight has to be tendered by, not just the day. Midnight
	// UTC means no cutoff is configured rather than a deadline at midnight.
	//
	// Recomputed while the order is still open whenever something it was derived from
	// moves — the basis above, or the carrier, service level, or ship-to address the
	// transit was quoted on. Renegotiating a customer's standing lead time or adding a
	// holiday to a calendar does not reach back into commitments already made. Cleared
	// if the order is unissued.
	ShipByDate time.Time `json:"ship_by_date" api:"required" format:"date-time"`
	// The ship date pinned by hand, bypassing transit and the customer's receiving
	// days.
	ShipByOverrideDate time.Time `json:"ship_by_override_date" api:"required" format:"date-time"`
	// Business days the carrier needs to cover this lane, subtracted from the promised
	// delivery date to reach the ship-by date.
	//
	// Only set when a delivery date was promised and the lane could be priced. Without
	// it the ship-by date falls back to the promised date itself.
	TransitDays int64 `json:"transit_days" api:"required"`
	// Where the transit estimate came from.
	//
	// Any of "carrier_lane", "service_level".
	TransitSource CommitmentTransitSource `json:"transit_source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CalendarAdjustmentDays respjson.Field
		EstimatedDeliveryDate  respjson.Field
		LeadTimeDays           respjson.Field
		LeadTimeOverrideDays   respjson.Field
		LeadTimeSource         respjson.Field
		Object                 respjson.Field
		PromisedAt             respjson.Field
		ShipByDate             respjson.Field
		ShipByOverrideDate     respjson.Field
		TransitDays            respjson.Field
		TransitSource          respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Commitment) RawJSON() string { return r.JSON.raw }
func (r *Commitment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which rule produced the ship-by date.
type CommitmentLeadTimeSource string

const (
	CommitmentLeadTimeSourceCustomer       CommitmentLeadTimeSource = "customer"
	CommitmentLeadTimeSourceParentCustomer CommitmentLeadTimeSource = "parent_customer"
	CommitmentLeadTimeSourceAccountGroup   CommitmentLeadTimeSource = "account_group"
	CommitmentLeadTimeSourceAccount        CommitmentLeadTimeSource = "account"
	CommitmentLeadTimeSourceManual         CommitmentLeadTimeSource = "manual"
	CommitmentLeadTimeSourceOrderLeadTime  CommitmentLeadTimeSource = "order_lead_time"
	CommitmentLeadTimeSourceOrderShipBy    CommitmentLeadTimeSource = "order_ship_by"
)

// Resource type identifier.
type CommitmentObject string

const (
	CommitmentObjectCommitment CommitmentObject = "commitment"
)

// Where the transit estimate came from.
type CommitmentTransitSource string

const (
	CommitmentTransitSourceCarrierLane  CommitmentTransitSource = "carrier_lane"
	CommitmentTransitSourceServiceLevel CommitmentTransitSource = "service_level"
)

// A rate calculated on demand rather than stored.
//
// The same shape as a rate minus the fields only a persisted row can have: it
// carries no ID and no timestamps because nothing was written. Used where a figure
// is derived per request, such as an analysis comparing one customer's price
// against the median other customers pay.
type ComputedRate struct {
	// Unit of measurement used for conversions and product quantities.
	DenominatorUnit Unit `json:"denominator_unit" api:"required"`
	// Human-readable formatted value (e.g. "$25.50 / pr").
	DisplayValue string `json:"display_value" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	NumeratorUnit Unit `json:"numerator_unit" api:"required"`
	// Resource type identifier.
	//
	// Any of "computed_rate".
	Object ComputedRateObject `json:"object" api:"required"`
	// Decimal value of the rate, as a string to preserve precision.
	//
	// Expressed as the amount of the numerator unit per one denominator unit.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DenominatorUnit respjson.Field
		DisplayValue    respjson.Field
		NumeratorUnit   respjson.Field
		Object          respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputedRate) RawJSON() string { return r.JSON.raw }
func (r *ComputedRate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ComputedRateObject string

const (
	ComputedRateObjectComputedRate ComputedRateObject = "computed_rate"
)

// Line item input for a create sales order request.
//
// The item, unit cost, and (unless an internal user supplies a `unit_price`
// override) the unit price are resolved server-side from the product. The quantity
// unit must belong to the product's unit group.
//
// The properties ProductID, Quantity are required.
type CreateSalesOrderLineInputParam struct {
	// ID of the product being ordered.
	ProductID string `json:"product_id" api:"required"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	// Description recorded on the line.
	//
	// Defaults to the product's description when omitted.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// SKU recorded on the line.
	//
	// Defaults to the product's SKU when omitted.
	ProductSKU param.Opt[string] `json:"product_sku,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreateSalesOrderLineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesOrderLineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesOrderLineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to create a sales order.
//
// The properties BillToAddressID, BuyerAccountID, Lines, PriorityCode,
// ShipToAddressID are required.
type CreateSalesOrderRequestParam struct {
	// Bill-to address ID.
	//
	// Must reference an existing address on the order's owner or buyer account.
	BillToAddressID string `json:"bill_to_address_id" api:"required"`
	// ID of the customer account the order is for.
	BuyerAccountID string `json:"buyer_account_id" api:"required"`
	// The line items to put on the order.
	//
	// The freight line, and the discount line when `order_discount_id` is supplied,
	// are added on top of these automatically.
	Lines []CreateSalesOrderLineInputParam `json:"lines,omitzero" api:"required"`
	// Fulfillment priority used to rank the order on the shop floor.
	//
	// Any of "low", "normal", "high".
	PriorityCode CreateSalesOrderRequestPriorityCode `json:"priority_code,omitzero" api:"required"`
	// Ship-to address ID.
	//
	// Must reference an existing address on the order's owner or buyer account.
	ShipToAddressID string `json:"ship_to_address_id" api:"required"`
	// Carrier billing account number charged when `carrier_billing_type` is
	// `third_party`.
	CarrierBillingAccountNumber param.Opt[string] `json:"carrier_billing_account_number,omitzero"`
	// ID of the carrier that will ship the order.
	//
	// Falls back to the customer's default carrier; the order is rejected when neither
	// is available.
	CarrierID param.Opt[string] `json:"carrier_id,omitzero"`
	// The customer's own purchase order number, for cross-referencing.
	//
	// Must be unique among your orders for this customer.
	CustomerPurchaseOrderNumber param.Opt[string] `json:"customer_purchase_order_number,omitzero"`
	// Days between this order being issued and it being due to ship, replacing the
	// customer's standing lead time for this order alone.
	//
	// Already a ship lead time, so no carrier transit is subtracted from it. Mutually
	// exclusive with promised_at and ship_by_override_date.
	LeadTimeOverrideDays param.Opt[int64] `json:"lead_time_override_days,omitzero"`
	// Free-form note about the order.
	Note param.Opt[string] `json:"note,omitzero"`
	// The order-level discount to apply, given as either its ID or its unique code.
	//
	// The discount is realized as an extra negative-priced line on the order rather
	// than as a separate total.
	OrderDiscountID param.Opt[string] `json:"order_discount_id,omitzero"`
	// ID of the payment terms for the order.
	//
	// Falls back to the customer's default payment term; the order is rejected when
	// neither is available.
	PaymentTermID param.Opt[string] `json:"payment_term_id,omitzero"`
	// Date delivery is promised to the customer.
	//
	// The order's ship-by date is worked back from this: the goods have to reach the
	// customer on a day they receive, so transit and both operating calendars are
	// subtracted from it. Mutually exclusive with lead_time_override_days and
	// ship_by_override_date.
	PromisedAt param.Opt[time.Time] `json:"promised_at,omitzero" format:"date-time"`
	// ID of the account user to credit as the order's sales rep.
	//
	// When omitted, a rep is assigned automatically: the customer's default sales rep
	// first, then the sales territory matching the ship-to postal code, then the
	// ship-to state. No rep is assigned when the customer is commission-exempt or
	// every ordered product belongs to a commission-exempt product line.
	SalesRepID param.Opt[string] `json:"sales_rep_id,omitzero"`
	// ID of the carrier service level the order ships on.
	//
	// Falls back to the customer's default service level, but only when `carrier_id`
	// is also omitted — supplying a carrier without a service level leaves the service
	// level unset.
	ServiceLevelID param.Opt[string] `json:"service_level_id,omitzero"`
	// The exact date the order is due to ship, bypassing transit and the customer's
	// receiving days.
	//
	// Still moved back to the nearest earlier day the plant ships on, since a date
	// nobody can ship on is not a deadline. Mutually exclusive with promised_at and
	// lead_time_override_days.
	ShipByOverrideDate param.Opt[time.Time] `json:"ship_by_override_date,omitzero" format:"date-time"`
	// ID of the shipping terms for the order.
	//
	// Falls back to the customer's default shipping term; the order is rejected when
	// neither is available.
	ShippingTermID param.Opt[string] `json:"shipping_term_id,omitzero"`
	// Users who should receive order acknowledgement emails for this order.
	//
	// Each must be a user on the customer's account.
	AcknowledgementEmailContacts []SalesOrderEmailContactInputParam `json:"acknowledgement_email_contacts,omitzero"`
	// Who is billed for freight.
	//
	//   - `sender`: the sender pays for shipping.
	//   - `third_party`: a third party pays for shipping, using the carrier billing
	//     account number.
	//
	// Any of "sender", "third_party".
	CarrierBillingType CreateSalesOrderRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// Users who should receive invoice emails for this order.
	//
	// Each must be a user on the customer's account.
	InvoiceEmailContacts []SalesOrderEmailContactInputParam `json:"invoice_email_contacts,omitzero"`
	paramObj
}

func (r CreateSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Fulfillment priority used to rank the order on the shop floor.
type CreateSalesOrderRequestPriorityCode string

const (
	CreateSalesOrderRequestPriorityCodeLow    CreateSalesOrderRequestPriorityCode = "low"
	CreateSalesOrderRequestPriorityCodeNormal CreateSalesOrderRequestPriorityCode = "normal"
	CreateSalesOrderRequestPriorityCodeHigh   CreateSalesOrderRequestPriorityCode = "high"
)

// Who is billed for freight.
//
//   - `sender`: the sender pays for shipping.
//   - `third_party`: a third party pays for shipping, using the carrier billing
//     account number.
type CreateSalesOrderRequestCarrierBillingType string

const (
	CreateSalesOrderRequestCarrierBillingTypeSender     CreateSalesOrderRequestCarrierBillingType = "sender"
	CreateSalesOrderRequestCarrierBillingTypeThirdParty CreateSalesOrderRequestCarrierBillingType = "third_party"
)

// CreatedBy describes who created a resource and their relationship to the account
// that owns it.
//
// It is resolved from the resource's create audit event.
type CreatedBy struct {
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Resource type identifier.
	//
	// Any of "created_by".
	Object CreatedByObject `json:"object" api:"required"`
	// The creator's relationship to the account that owns the resource.
	//
	// - `internal`: created by a user of the owning account.
	// - `customer`: created by a customer of the owning account.
	// - `system`: created automatically with no human actor (e.g. an EDI import).
	//
	// Any of "internal", "customer", "system".
	Relation CreatedByRelation `json:"relation" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Actor       respjson.Field
		Object      respjson.Field
		Relation    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedBy) RawJSON() string { return r.JSON.raw }
func (r *CreatedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CreatedByObject string

const (
	CreatedByObjectCreatedBy CreatedByObject = "created_by"
)

// The creator's relationship to the account that owns the resource.
//
// - `internal`: created by a user of the owning account.
// - `customer`: created by a customer of the owning account.
// - `system`: created automatically with no human actor (e.g. an EDI import).
type CreatedByRelation string

const (
	CreatedByRelationInternal CreatedByRelation = "internal"
	CreatedByRelationCustomer CreatedByRelation = "customer"
	CreatedByRelationSystem   CreatedByRelation = "system"
)

// Freight describes the carrier selection and freight billing for a record.
//
// It is a generic, reusable sub-resource shared by anything that carries shipping
// configuration — a sales order, a purchase order, or a shipment.
type Freight struct {
	// Carrier account number to bill, used when `billing_type` is `third_party`.
	BillingAccountNumber string `json:"billing_account_number" api:"required"`
	// Which party the carrier bills for the shipment.
	//
	// - `sender`: the shipper (your account) is billed.
	// - `third_party`: a third party is billed via `billing_account_number`.
	//
	// Any of "sender", "third_party".
	BillingType FreightBillingType `json:"billing_type" api:"required"`
	// A shipping carrier configured for fulfilling orders.
	//
	// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
	// through Shippo for live rating and label purchase; other carriers represent
	// self-managed shipping methods such as will call or local delivery.
	Carrier Carrier `json:"carrier" api:"required"`
	// Resource type identifier.
	//
	// Any of "freight".
	Object FreightObject `json:"object" api:"required"`
	// How freight is arranged and billed for the record.
	//
	// - `free_freight`: no shipping cost to the buyer.
	// - `billed_freight`: freight is billed to the buyer.
	//
	// Sales orders, purchase orders, and shipments do not carry a policy of their own.
	// Freight on those records is waived when the customer's freight preferences, the
	// customer's type group, any of its pricing groups, the customer's shipping term,
	// or any product line on the order is `free_freight`.
	//
	// Any of "free_freight", "billed_freight".
	Policy FreightPolicy `json:"policy" api:"required"`
	// A shipping speed or method offered by a carrier, such as ground or overnight.
	//
	// Carriers connected through Shippo have their service levels synced from the
	// carrier itself; any carrier can also have service levels you create by hand.
	ServiceLevel ServiceLevel `json:"service_level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAccountNumber respjson.Field
		BillingType          respjson.Field
		Carrier              respjson.Field
		Object               respjson.Field
		Policy               respjson.Field
		ServiceLevel         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Freight) RawJSON() string { return r.JSON.raw }
func (r *Freight) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Which party the carrier bills for the shipment.
//
// - `sender`: the shipper (your account) is billed.
// - `third_party`: a third party is billed via `billing_account_number`.
type FreightBillingType string

const (
	FreightBillingTypeSender     FreightBillingType = "sender"
	FreightBillingTypeThirdParty FreightBillingType = "third_party"
)

// Resource type identifier.
type FreightObject string

const (
	FreightObjectFreight FreightObject = "freight"
)

// How freight is arranged and billed for the record.
//
// - `free_freight`: no shipping cost to the buyer.
// - `billed_freight`: freight is billed to the buyer.
//
// Sales orders, purchase orders, and shipments do not carry a policy of their own.
// Freight on those records is waived when the customer's freight preferences, the
// customer's type group, any of its pricing groups, the customer's shipping term,
// or any product line on the order is `free_freight`.
type FreightPolicy string

const (
	FreightPolicyFreeFreight   FreightPolicy = "free_freight"
	FreightPolicyBilledFreight FreightPolicy = "billed_freight"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListQuotedSalesOrderLine struct {
	// Resources in this page.
	Data []QuotedSalesOrderLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListQuotedSalesOrderLineObject `json:"object" api:"required"`
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
func (r ListQuotedSalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *ListQuotedSalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListQuotedSalesOrderLineObject string

const (
	ListQuotedSalesOrderLineObjectList ListQuotedSalesOrderLineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListRecord struct {
	// Resources in this page.
	Data []Record `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListRecordObject `json:"object" api:"required"`
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
func (r ListRecord) RawJSON() string { return r.JSON.raw }
func (r *ListRecord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListRecordObject string

const (
	ListRecordObjectList ListRecordObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListSalesOrder struct {
	// Resources in this page.
	Data []SalesOrder `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesOrderObject `json:"object" api:"required"`
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
func (r ListSalesOrder) RawJSON() string { return r.JSON.raw }
func (r *ListSalesOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesOrderObject string

const (
	ListSalesOrderObjectList ListSalesOrderObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListSalesOrderLine struct {
	// Resources in this page.
	Data []SalesOrderLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesOrderLineObject `json:"object" api:"required"`
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
func (r ListSalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *ListSalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesOrderLineObject string

const (
	ListSalesOrderLineObjectList ListSalesOrderLineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListSalesOrderStatus struct {
	// Resources in this page.
	Data []SalesOrderStatus `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSalesOrderStatusObject `json:"object" api:"required"`
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
func (r ListSalesOrderStatus) RawJSON() string { return r.JSON.raw }
func (r *ListSalesOrderStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSalesOrderStatusObject string

const (
	ListSalesOrderStatusObjectList ListSalesOrderStatusObject = "list"
)

// A sales order's email recipients, grouped by the notification they receive.
type OrderContact struct {
	// Email addresses that receive order acknowledgements for this order.
	Acknowledgement []string `json:"acknowledgement" api:"required"`
	// Email addresses that receive invoices for this order.
	Invoice []string `json:"invoice" api:"required"`
	// Resource type identifier.
	//
	// Any of "order_contact".
	Object OrderContactObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Acknowledgement respjson.Field
		Invoice         respjson.Field
		Object          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderContact) RawJSON() string { return r.JSON.raw }
func (r *OrderContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type OrderContactObject string

const (
	OrderContactObjectOrderContact OrderContactObject = "order_contact"
)

// A line to price in a quote request.
//
// The properties ProductID, Quantity are required.
type QuoteSalesOrderLineInputParam struct {
	// ID of the product to price.
	ProductID string `json:"product_id" api:"required"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	paramObj
}

func (r QuoteSalesOrderLineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow QuoteSalesOrderLineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteSalesOrderLineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to quote sales-order line prices without creating an order.
//
// The properties BuyerAccountID, Lines are required.
type QuoteSalesOrderPricesRequestParam struct {
	// ID of the customer account the prices are for.
	BuyerAccountID string `json:"buyer_account_id" api:"required"`
	// Lines to price.
	Lines []QuoteSalesOrderLineInputParam `json:"lines,omitzero" api:"required"`
	paramObj
}

func (r QuoteSalesOrderPricesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow QuoteSalesOrderPricesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuoteSalesOrderPricesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quoted unit prices for the requested lines, in request order.
type QuoteSalesOrderPricesResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lines ListQuotedSalesOrderLine `json:"lines" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_price_quote".
	Object QuoteSalesOrderPricesResponseObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lines       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteSalesOrderPricesResponse) RawJSON() string { return r.JSON.raw }
func (r *QuoteSalesOrderPricesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuoteSalesOrderPricesResponseObject string

const (
	QuoteSalesOrderPricesResponseObjectSalesOrderPriceQuote QuoteSalesOrderPricesResponseObject = "sales_order_price_quote"
)

// One priced line in a quote response.
type QuotedSalesOrderLine struct {
	// Resource type identifier.
	//
	// Any of "sales_order_price_quote_line".
	Object QuotedSalesOrderLineObject `json:"object" api:"required"`
	// A catalog entry as it is sold: an inventory item together with its product type,
	// product line, and customer portal visibility.
	//
	// Every product is backed by exactly one item, which carries the SKU, description,
	// pricing, attributes, and inventory position. Creating a product creates that
	// item; deleting the product deletes it.
	Product Product `json:"product" api:"required"`
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
		Product     respjson.Field
		UnitPrice   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotedSalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *QuotedSalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuotedSalesOrderLineObject string

const (
	QuotedSalesOrderLineObjectSalesOrderPriceQuoteLine QuotedSalesOrderLineObject = "sales_order_price_quote_line"
)

// Record is a lightweight reference to a business record — a sales order, purchase
// order, pick, shipment, production run, invoice, etc.
//
// Like the `actor` and `entity` references, it carries just enough to identify and
// label the referenced record without embedding its full resource. The `status`
// and `metadata` fields hold type-specific detail that varies by the kind of
// record referenced.
type Record struct {
	// Unique identifier for the record.
	ID string `json:"id" api:"required"`
	// Type-specific metadata.
	//
	// The set of keys varies by record type.
	Metadata map[string]string `json:"metadata" api:"required"`
	// Human-readable record number, when the record has one.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "record".
	Object RecordObject `json:"object" api:"required"`
	// Type-specific status code, when applicable.
	Status string `json:"status" api:"required"`
	// The kind of business record referenced.
	//
	// Determines how to resolve the record and which `status` and `metadata` keys may
	// appear.
	//
	// - `sales_order`: a customer order.
	// - `purchase_order`: an order placed with a supplier.
	// - `receiving_order`: an inbound order being received into inventory.
	// - `pick`: a warehouse pick task.
	// - `shipment`: an outbound shipment.
	// - `delivery`: a delivery of one or more shipments to a destination.
	// - `production_run`: a manufacturing production run.
	// - `invoice`: a customer invoice.
	// - `transaction`: a payment or financial transaction.
	// - `settlement`: a settlement reconciling transactions against invoices.
	//
	// Any of "sales_order", "purchase_order", "receiving_order", "pick", "shipment",
	// "delivery", "production_run", "invoice", "transaction", "settlement".
	Type RecordType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Metadata    respjson.Field
		Number      respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Record) RawJSON() string { return r.JSON.raw }
func (r *Record) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RecordObject string

const (
	RecordObjectRecord RecordObject = "record"
)

// The kind of business record referenced.
//
// Determines how to resolve the record and which `status` and `metadata` keys may
// appear.
//
// - `sales_order`: a customer order.
// - `purchase_order`: an order placed with a supplier.
// - `receiving_order`: an inbound order being received into inventory.
// - `pick`: a warehouse pick task.
// - `shipment`: an outbound shipment.
// - `delivery`: a delivery of one or more shipments to a destination.
// - `production_run`: a manufacturing production run.
// - `invoice`: a customer invoice.
// - `transaction`: a payment or financial transaction.
// - `settlement`: a settlement reconciling transactions against invoices.
type RecordType string

const (
	RecordTypeSalesOrder     RecordType = "sales_order"
	RecordTypePurchaseOrder  RecordType = "purchase_order"
	RecordTypeReceivingOrder RecordType = "receiving_order"
	RecordTypePick           RecordType = "pick"
	RecordTypeShipment       RecordType = "shipment"
	RecordTypeDelivery       RecordType = "delivery"
	RecordTypeProductionRun  RecordType = "production_run"
	RecordTypeInvoice        RecordType = "invoice"
	RecordTypeTransaction    RecordType = "transaction"
	RecordTypeSettlement     RecordType = "settlement"
)

// An order placed by a customer, tracked from estimate through fulfillment.
type SalesOrder struct {
	// Sales order ID.
	ID string `json:"id" api:"required"`
	// Whether an order acknowledgment has been sent to the customer.
	//
	// Becomes `sent` when the order is issued with customer notification requested and
	// the order has acknowledgement contacts to send to. It can also be set directly
	// when an acknowledgement was sent outside OpenMRP.
	//
	// Any of "not_sent", "sent".
	AcknowledgmentStatus SalesOrderAcknowledgmentStatus `json:"acknowledgment_status" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	BillToAddress Address `json:"bill_to_address" api:"required"`
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
	// When the order was fulfilled and closed.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// A sales order's email recipients, grouped by the notification they receive.
	Contacts OrderContact `json:"contacts" api:"required"`
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
	// The customer's own purchase order number, for cross-referencing.
	//
	// Unique among this customer's orders.
	CustomerPurchaseOrderNumber string `json:"customer_purchase_order_number" api:"required"`
	// When this estimate expires, if an expiration was set.
	ExpiredAt time.Time `json:"expired_at" api:"required" format:"date-time"`
	// When the first shipment against this order went out.
	FirstShipAt time.Time `json:"first_ship_at" api:"required" format:"date-time"`
	// Freight describes the carrier selection and freight billing for a record.
	//
	// It is a generic, reusable sub-resource shared by anything that carries shipping
	// configuration — a sales order, a purchase order, or a shipment.
	Freight Freight `json:"freight" api:"required"`
	// When the order was issued (moved out of `estimate`).
	IssuedAt time.Time `json:"issued_at" api:"required" format:"date-time"`
	// Number of lines on this order.
	LineCount int64 `json:"line_count" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Lines ListSalesOrderLine `json:"lines" api:"required"`
	// Free-form note about the order.
	Note string `json:"note" api:"required"`
	// Human-readable order number, e.g. `SO-001`.
	//
	// Assigned automatically when the order is created; unique within your account.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order".
	Object SalesOrderObject `json:"object" api:"required"`
	// A discount code that can be applied to a sales order.
	//
	// An order discount reduces the order total by either a percentage or a fixed
	// amount, depending on `discount_type`. The reduction is capped at the order total
	// and rounded to the nearest cent.
	OrderDiscount OrderDiscount `json:"order_discount" api:"required"`
	// Stripe payment intent IDs recorded against this order.
	PaymentIntentIDs []string `json:"payment_intent_ids" api:"required"`
	// Payment state of the order, derived from settlement allocations, invoices, and
	// Stripe payments.
	//
	// Any of "unpaid", "partially_paid", "paid".
	PaymentStatus SalesOrderPaymentStatus `json:"payment_status" api:"required"`
	// A payment term describing when payment is due (e.g. `Net 30`), assignable to
	// customers, sales orders, purchase orders, and invoices.
	PaymentTerm PaymentTerm `json:"payment_term" api:"required"`
	// Fulfillment priority, used to rank orders on the shop floor.
	//
	// Any of "low", "normal", "high".
	Priority SalesOrderPriority `json:"priority" api:"required"`
	// The fulfillment records produced from a sales order.
	//
	// The group itself is returned only when at least one of its members has been
	// expanded.
	Related SalesOrderRelated `json:"related" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	SalesRep Actor `json:"sales_rep" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	ShipToAddress Address `json:"ship_to_address" api:"required"`
	// A named freight pricing rule that decides what a buyer pays for shipping.
	//
	// A customer's default shipping term is evaluated whenever freight is quoted for
	// one of their orders. Freight exemptions on the customer, its type group, or any
	// of its price groups are checked first and zero the freight charge before the
	// shipping term is considered.
	ShippingTerm ShippingTerm `json:"shipping_term" api:"required"`
	// Order lifecycle status.
	//
	//   - `estimate`: a draft quote that has not yet been committed; not counted as a
	//     real order.
	//   - `issued`: the order has been issued and is being fulfilled.
	//   - `fulfilled`: the order has been completed and closed.
	//
	// Status changes are made through the issue, unissue, close, and reopen action
	// endpoints rather than by updating this field.
	//
	// Any of "estimate", "issued", "fulfilled".
	Status SalesOrderStatus `json:"status" api:"required"`
	// Derived monetary totals for a sales order or one of its lines.
	//
	// Fulfillment runs ordered -> picked -> packed -> invoiced, and each downstream
	// stage reports both the money that has reached it and its progress against the
	// ordered baseline.
	Totals SalesOrderTotals `json:"totals" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                          respjson.Field
		AcknowledgmentStatus        respjson.Field
		BillToAddress               respjson.Field
		Commitment                  respjson.Field
		CompletedAt                 respjson.Field
		Contacts                    respjson.Field
		CreatedAt                   respjson.Field
		CreatedBy                   respjson.Field
		Customer                    respjson.Field
		CustomerPurchaseOrderNumber respjson.Field
		ExpiredAt                   respjson.Field
		FirstShipAt                 respjson.Field
		Freight                     respjson.Field
		IssuedAt                    respjson.Field
		LineCount                   respjson.Field
		Lines                       respjson.Field
		Note                        respjson.Field
		Number                      respjson.Field
		Object                      respjson.Field
		OrderDiscount               respjson.Field
		PaymentIntentIDs            respjson.Field
		PaymentStatus               respjson.Field
		PaymentTerm                 respjson.Field
		Priority                    respjson.Field
		Related                     respjson.Field
		SalesRep                    respjson.Field
		ShipToAddress               respjson.Field
		ShippingTerm                respjson.Field
		Status                      respjson.Field
		Totals                      respjson.Field
		UpdatedAt                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrder) RawJSON() string { return r.JSON.raw }
func (r *SalesOrder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether an order acknowledgment has been sent to the customer.
//
// Becomes `sent` when the order is issued with customer notification requested and
// the order has acknowledgement contacts to send to. It can also be set directly
// when an acknowledgement was sent outside OpenMRP.
type SalesOrderAcknowledgmentStatus string

const (
	SalesOrderAcknowledgmentStatusNotSent SalesOrderAcknowledgmentStatus = "not_sent"
	SalesOrderAcknowledgmentStatusSent    SalesOrderAcknowledgmentStatus = "sent"
)

// Resource type identifier.
type SalesOrderObject string

const (
	SalesOrderObjectSalesOrder SalesOrderObject = "sales_order"
)

// Payment state of the order, derived from settlement allocations, invoices, and
// Stripe payments.
type SalesOrderPaymentStatus string

const (
	SalesOrderPaymentStatusUnpaid        SalesOrderPaymentStatus = "unpaid"
	SalesOrderPaymentStatusPartiallyPaid SalesOrderPaymentStatus = "partially_paid"
	SalesOrderPaymentStatusPaid          SalesOrderPaymentStatus = "paid"
)

// Fulfillment priority, used to rank orders on the shop floor.
type SalesOrderPriority string

const (
	SalesOrderPriorityLow    SalesOrderPriority = "low"
	SalesOrderPriorityNormal SalesOrderPriority = "normal"
	SalesOrderPriorityHigh   SalesOrderPriority = "high"
)

// Order lifecycle status.
//
//   - `estimate`: a draft quote that has not yet been committed; not counted as a
//     real order.
//   - `issued`: the order has been issued and is being fulfilled.
//   - `fulfilled`: the order has been completed and closed.
//
// Status changes are made through the issue, unissue, close, and reopen action
// endpoints rather than by updating this field.
type SalesOrderStatus string

const (
	SalesOrderStatusEstimate  SalesOrderStatus = "estimate"
	SalesOrderStatusIssued    SalesOrderStatus = "issued"
	SalesOrderStatusFulfilled SalesOrderStatus = "fulfilled"
)

// A user subscribed to one of a sales order's email notifications.
//
// The property AccountUserID is required.
type SalesOrderEmailContactInputParam struct {
	// ID of the account user who should receive the notification.
	AccountUserID string `json:"account_user_id" api:"required"`
	paramObj
}

func (r SalesOrderEmailContactInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SalesOrderEmailContactInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SalesOrderEmailContactInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single line item on a sales order.
type SalesOrderLine struct {
	// Sales order line ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// Position of the line on the order.
	//
	// Assigned automatically in sequence, starting at `1`. Product lines are numbered
	// first and the automatically generated freight and discount lines always sit at
	// the bottom; removing a line renumbers the rest so the sequence stays contiguous.
	LineItemNumber int64 `json:"line_item_number" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_line".
	Object SalesOrderLineObject `json:"object" api:"required"`
	// A catalog entry as it is sold: an inventory item together with its product type,
	// product line, and customer portal visibility.
	//
	// Every product is backed by exactly one item, which carries the SKU, description,
	// pricing, attributes, and inventory position. Creating a product creates that
	// item; deleting the product deletes it.
	Product Product `json:"product" api:"required"`
	// Description recorded on this line, taken from the product unless the line
	// supplies its own.
	ProductDescription string `json:"product_description" api:"required"`
	// SKU recorded on this line.
	//
	// Taken from the product unless the line supplies its own, and editable
	// afterwards, so it preserves what was sold even if the product's SKU later
	// changes.
	ProductSKU string `json:"product_sku" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	QuantityOrdered Quantity `json:"quantity_ordered" api:"required"`
	// Derived monetary totals for a sales order or one of its lines.
	//
	// Fulfillment runs ordered -> picked -> packed -> invoiced, and each downstream
	// stage reports both the money that has reached it and its progress against the
	// ordered baseline.
	Totals SalesOrderTotals `json:"totals" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitCost Rate `json:"unit_cost" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitPrice Rate `json:"unit_price" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Item               respjson.Field
		LineItemNumber     respjson.Field
		Object             respjson.Field
		Product            respjson.Field
		ProductDescription respjson.Field
		ProductSKU         respjson.Field
		QuantityOrdered    respjson.Field
		Totals             respjson.Field
		UnitCost           respjson.Field
		UnitPrice          respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderLine) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderLineObject string

const (
	SalesOrderLineObjectSalesOrderLine SalesOrderLineObject = "sales_order_line"
)

// The fulfillment records produced from a sales order.
//
// The group itself is returned only when at least one of its members has been
// expanded.
type SalesOrderRelated struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Invoices ListRecord `json:"invoices" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_related".
	Object SalesOrderRelatedObject `json:"object" api:"required"`
	// Record is a lightweight reference to a business record — a sales order, purchase
	// order, pick, shipment, production run, invoice, etc.
	//
	// Like the `actor` and `entity` references, it carries just enough to identify and
	// label the referenced record without embedding its full resource. The `status`
	// and `metadata` fields hold type-specific detail that varies by the kind of
	// record referenced.
	Pick Record `json:"pick" api:"required"`
	// Record is a lightweight reference to a business record — a sales order, purchase
	// order, pick, shipment, production run, invoice, etc.
	//
	// Like the `actor` and `entity` references, it carries just enough to identify and
	// label the referenced record without embedding its full resource. The `status`
	// and `metadata` fields hold type-specific detail that varies by the kind of
	// record referenced.
	ProductionRun Record `json:"production_run" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Shipments ListRecord `json:"shipments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoices      respjson.Field
		Object        respjson.Field
		Pick          respjson.Field
		ProductionRun respjson.Field
		Shipments     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderRelated) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderRelated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderRelatedObject string

const (
	SalesOrderRelatedObjectSalesOrderRelated SalesOrderRelatedObject = "sales_order_related"
)

// The monetary amount that has reached one fulfillment stage, together with how
// far that stage has progressed.
type SalesOrderStageTotal struct {
	// Amount that has reached this stage, as a decimal string (unit price times the
	// quantity at this stage).
	Amount string `json:"amount" api:"required" format:"decimal"`
	// Progress through this stage, as a fraction between 0 and 1.
	//
	// Calculated as the quantity that has reached this stage divided by the quantity
	// ordered, so `1` means the whole order has cleared the stage and `0` means
	// nothing has reached it yet.
	Completion float64 `json:"completion" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_stage_total".
	Object SalesOrderStageTotalObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Completion  respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderStageTotal) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderStageTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderStageTotalObject string

const (
	SalesOrderStageTotalObjectSalesOrderStageTotal SalesOrderStageTotalObject = "sales_order_stage_total"
)

// Derived monetary totals for a sales order or one of its lines.
//
// Fulfillment runs ordered -> picked -> packed -> invoiced, and each downstream
// stage reports both the money that has reached it and its progress against the
// ordered baseline.
type SalesOrderTotals struct {
	// The monetary amount that has reached one fulfillment stage, together with how
	// far that stage has progressed.
	Invoiced SalesOrderStageTotal `json:"invoiced" api:"required"`
	// Resource type identifier.
	//
	// Any of "sales_order_totals".
	Object SalesOrderTotalsObject `json:"object" api:"required"`
	// Total ordered amount as a decimal string (unit price x quantity ordered).
	//
	// This is the baseline the stage completions are measured against.
	Ordered string `json:"ordered" api:"required" format:"decimal"`
	// The monetary amount that has reached one fulfillment stage, together with how
	// far that stage has progressed.
	Packed SalesOrderStageTotal `json:"packed" api:"required"`
	// The monetary amount that has reached one fulfillment stage, together with how
	// far that stage has progressed.
	Picked SalesOrderStageTotal `json:"picked" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoiced    respjson.Field
		Object      respjson.Field
		Ordered     respjson.Field
		Packed      respjson.Field
		Picked      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SalesOrderTotals) RawJSON() string { return r.JSON.raw }
func (r *SalesOrderTotals) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type SalesOrderTotalsObject string

const (
	SalesOrderTotalsObjectSalesOrderTotals SalesOrderTotalsObject = "sales_order_totals"
)

// Request to update a sales order.
type UpdateSalesOrderRequestParam struct {
	// Carrier billing account number charged when `carrier_billing_type` is
	// `third_party`.
	CarrierBillingAccountNumber param.Opt[string] `json:"carrier_billing_account_number,omitzero"`
	// The customer's own purchase order number, for cross-referencing.
	CustomerPurchaseOrderNumber param.Opt[string] `json:"customer_purchase_order_number,omitzero"`
	// Days between this order being issued and it being due to ship, replacing the
	// customer's standing lead time for this order alone. Mutually exclusive with
	// promised_at and ship_by_override_date; clear one to switch to another.
	LeadTimeOverrideDays param.Opt[int64] `json:"lead_time_override_days,omitzero"`
	// Free-form note about the order.
	Note param.Opt[string] `json:"note,omitzero"`
	// ID of the order-level discount recorded on the order.
	//
	// Changing this does not add, reprice, or remove the order's discount line; adjust
	// that line directly.
	OrderDiscountID param.Opt[string] `json:"order_discount_id,omitzero"`
	// Date delivery is promised to the customer.
	PromisedAt param.Opt[time.Time] `json:"promised_at,omitzero" format:"date-time"`
	// ID of the account user to credit as the order's sales rep.
	SalesRepID param.Opt[string] `json:"sales_rep_id,omitzero"`
	// ID of the carrier service level the order ships on.
	ServiceLevelID param.Opt[string] `json:"service_level_id,omitzero"`
	// The exact date the order is due to ship, bypassing transit and the customer's
	// receiving days. Mutually exclusive with promised_at and lead_time_override_days.
	ShipByOverrideDate param.Opt[time.Time] `json:"ship_by_override_date,omitzero" format:"date-time"`
	// Billing address ID.
	//
	// Re-points the order to an existing address. To change an address's contents, use
	// the update-address endpoint.
	BillingAddressID param.Opt[string] `json:"billing_address_id,omitzero"`
	// ID of the carrier that will ship the order.
	CarrierID param.Opt[string] `json:"carrier_id,omitzero"`
	// Moves the order to a different customer account.
	//
	// Existing lines keep the prices they were created with; they are not re-priced
	// against the new customer.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// ID of the payment terms for the order.
	PaymentTermID param.Opt[string] `json:"payment_term_id,omitzero"`
	// Shipping address ID.
	//
	// Re-points the order to an existing address. To change an address's contents, use
	// the update-address endpoint.
	ShippingAddressID param.Opt[string] `json:"shipping_address_id,omitzero"`
	// ID of the shipping terms for the order.
	ShippingTermID param.Opt[string] `json:"shipping_term_id,omitzero"`
	// Who is billed for freight.
	//
	//   - `sender`: the sender pays for shipping.
	//   - `third_party`: a third party pays for shipping, using the carrier billing
	//     account number.
	//
	// Any of "sender", "third_party".
	CarrierBillingType UpdateSalesOrderRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// Replaces the acknowledgement email contacts on the order.
	//
	// An empty list clears all contacts; omitting the field leaves existing contacts
	// untouched.
	AcknowledgementEmailContacts []SalesOrderEmailContactInputParam `json:"acknowledgement_email_contacts,omitzero"`
	// Acknowledgment status of the order.
	//
	// Set to `sent` to mark the acknowledgement as sent without emailing the customer,
	// or `not_sent` to reset it.
	//
	// Any of "not_sent", "sent".
	AcknowledgmentStatus UpdateSalesOrderRequestAcknowledgmentStatus `json:"acknowledgment_status,omitzero"`
	// Replaces the invoice email contacts on the order.
	//
	// An empty list clears all contacts; omitting the field leaves existing contacts
	// untouched.
	InvoiceEmailContacts []SalesOrderEmailContactInputParam `json:"invoice_email_contacts,omitzero"`
	// New fulfillment priority for the order.
	//
	// Any of "low", "normal", "high".
	PriorityCode UpdateSalesOrderRequestPriorityCode `json:"priority_code,omitzero"`
	paramObj
}

func (r UpdateSalesOrderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateSalesOrderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateSalesOrderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Acknowledgment status of the order.
//
// Set to `sent` to mark the acknowledgement as sent without emailing the customer,
// or `not_sent` to reset it.
type UpdateSalesOrderRequestAcknowledgmentStatus string

const (
	UpdateSalesOrderRequestAcknowledgmentStatusNotSent UpdateSalesOrderRequestAcknowledgmentStatus = "not_sent"
	UpdateSalesOrderRequestAcknowledgmentStatusSent    UpdateSalesOrderRequestAcknowledgmentStatus = "sent"
)

// Who is billed for freight.
//
//   - `sender`: the sender pays for shipping.
//   - `third_party`: a third party pays for shipping, using the carrier billing
//     account number.
type UpdateSalesOrderRequestCarrierBillingType string

const (
	UpdateSalesOrderRequestCarrierBillingTypeSender     UpdateSalesOrderRequestCarrierBillingType = "sender"
	UpdateSalesOrderRequestCarrierBillingTypeThirdParty UpdateSalesOrderRequestCarrierBillingType = "third_party"
)

// New fulfillment priority for the order.
type UpdateSalesOrderRequestPriorityCode string

const (
	UpdateSalesOrderRequestPriorityCodeLow    UpdateSalesOrderRequestPriorityCode = "low"
	UpdateSalesOrderRequestPriorityCodeNormal UpdateSalesOrderRequestPriorityCode = "normal"
	UpdateSalesOrderRequestPriorityCodeHigh   UpdateSalesOrderRequestPriorityCode = "high"
)

type SaleSalesOrderDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderNewParams struct {
	// Request to create a sales order.
	CreateSalesOrderRequest CreateSalesOrderRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "bill_to_address", "ship_to_address", "freight",
	// "payment_term", "shipping_term", "order_discount", "totals", "contacts",
	// "related.pick", "related.production_run", "related.shipments",
	// "related.invoices", "lines", "lines.product", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleSalesOrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSalesOrderRequest)
}
func (r *SaleSalesOrderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderNewParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "created_by", "bill_to_address",
	// "ship_to_address", "freight", "payment_term", "shipping_term", "order_discount",
	// "totals", "contacts", "related.pick", "related.production_run",
	// "related.shipments", "related.invoices", "lines", "lines.product",
	// "lines.product.item", "lines.product.item.category",
	// "lines.product.item.category.properties",
	// "lines.product.item.category.unit_group",
	// "lines.product.item.category.unit_group.base_unit",
	// "lines.product.item.category.unit_group.associated_units",
	// "lines.product.item.category.unit_group.associated_units.unit",
	// "lines.product.product_line", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderGetParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "bill_to_address", "ship_to_address", "freight",
	// "payment_term", "shipping_term", "order_discount", "totals", "contacts",
	// "related.pick", "related.production_run", "related.shipments",
	// "related.invoices", "lines", "lines.product", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a sales order.
	UpdateSalesOrderRequest UpdateSalesOrderRequestParam
	paramObj
}

func (r SaleSalesOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateSalesOrderRequest)
}
func (r *SaleSalesOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Latest order creation date to include, in `YYYY-MM-DD` format.
	//
	// Compared against the creation timestamp at the start of that day, so orders
	// created later on the end date itself are excluded; pass the following day to
	// include them.
	EndsAt param.Opt[string] `query:"ends_at,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Restricts results to orders that are, or are not, past their ship-by date.
	//
	// An order is past due when it is still `issued` and its ship-by date has passed.
	// A fulfilled order that shipped late is not past due — it is delivered, and how
	// late it was is a delivery-performance question rather than a backlog one.
	PastDue param.Opt[bool] `query:"past_due,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Earliest ship-by date to include, in `YYYY-MM-DD` format. Inclusive of the date
	// itself.
	ShipByAfter param.Opt[string] `query:"ship_by_after,omitzero" json:"-"`
	// Latest ship-by date to include, in `YYYY-MM-DD` format. Inclusive of the date
	// itself.
	ShipByBefore param.Opt[string] `query:"ship_by_before,omitzero" json:"-"`
	// Earliest order creation date to include, in `YYYY-MM-DD` format.
	StartsAt param.Opt[string] `query:"starts_at,omitzero" json:"-"`
	// Restricts results to orders placed by customers belonging to any of these
	// account groups.
	CustomerGroupIDs []string `query:"customer_group_ids,omitzero" json:"-"`
	// Restricts results to orders placed by any of these customers.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer", "sales_rep", "created_by", "bill_to_address",
	// "ship_to_address", "freight", "payment_term", "shipping_term", "order_discount",
	// "totals", "contacts", "related.pick", "related.production_run",
	// "related.shipments", "related.invoices", "lines", "lines.product",
	// "lines.product.item", "lines.product.item.category",
	// "lines.product.item.category.properties",
	// "lines.product.item.category.unit_group",
	// "lines.product.item.category.unit_group.base_unit",
	// "lines.product.item.category.unit_group.associated_units",
	// "lines.product.item.category.unit_group.associated_units.unit",
	// "lines.product.product_line", "lines.quantity_ordered",
	// "lines.quantity_ordered.unit", "lines.unit_price",
	// "lines.unit_price.numerator_unit", "lines.unit_price.denominator_unit",
	// "lines.unit_cost", "lines.unit_cost.numerator_unit",
	// "lines.unit_cost.denominator_unit", "lines.totals".
	Include []string `query:"include,omitzero" json:"-"`
	// Restricts results to orders that have at least one line for any of these
	// inventory items.
	ItemIDs []string `query:"item_ids,omitzero" json:"-"`
	// Restricts results to orders that have at least one line whose product belongs to
	// any of these product lines.
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	// Restricts results to orders credited to any of these sales reps.
	//
	// These are account user IDs, matching the `sales_rep` on the order.
	SalesRepIDs []string `query:"sales_rep_ids,omitzero" json:"-"`
	// Restricts results to orders in any of these lifecycle statuses.
	//
	// Any of "estimate", "issued", "fulfilled".
	StatusCodes []string `query:"status_codes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderListParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderCheckoutParams struct {
	// Request to create a checkout session for a sales order.
	CheckoutSalesOrderRequest CheckoutSalesOrderRequestParam
	paramObj
}

func (r SaleSalesOrderCheckoutParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CheckoutSalesOrderRequest)
}
func (r *SaleSalesOrderCheckoutParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderPriceQuoteParams struct {
	// Request to quote sales-order line prices without creating an order.
	QuoteSalesOrderPricesRequest QuoteSalesOrderPricesRequestParam
	paramObj
}

func (r SaleSalesOrderPriceQuoteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.QuoteSalesOrderPricesRequest)
}
func (r *SaleSalesOrderPriceQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderGetStatusesParams struct {
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
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleSalesOrderGetStatusesParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderGetStatusesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
