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

// Manage customer accounts.
//
// SaleCustomerService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleCustomerService] method instead.
type SaleCustomerService struct {
	options []option.RequestOption
	// Manage customer accounts.
	Actions SaleCustomerActionService
}

// NewSaleCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSaleCustomerService(opts ...option.RequestOption) (r SaleCustomerService) {
	r = SaleCustomerService{}
	r.options = opts
	r.Actions = NewSaleCustomerActionService(opts...)
	return
}

// Creates a customer account with its default addresses, fulfillment settings, and
// order policies.
//
// If `number` is omitted, the next sequential customer number is assigned
// automatically.
//
// This endpoint requires the permission: `customers:create`.
func (r *SaleCustomerService) New(ctx context.Context, params SaleCustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a customer by ID.
//
// This endpoint requires the permissions: `customers:read`, `suppliers:read`.
func (r *SaleCustomerService) Get(ctx context.Context, id string, query SaleCustomerGetParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a customer account.
//
// Only the fields provided in the request are changed. Nullable fields can be set
// to `null` to clear their current value.
//
// This endpoint requires the permission: `customers:update`.
func (r *SaleCustomerService) Update(ctx context.Context, id string, params SaleCustomerUpdateParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of customers for the current account.
//
// This endpoint requires the permission: `customers:read`.
func (r *SaleCustomerService) List(ctx context.Context, query SaleCustomerListParams, opts ...option.RequestOption) (res *ListCustomer, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a customer.
//
// Fails with a conflict error if any sales orders still reference the customer;
// delete or reassign those orders, or merge the customer into another first.
//
// This endpoint requires the permission: `customers:delete`.
func (r *SaleCustomerService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleCustomerDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the ship-by lead time a new order for this customer would be committed
// to.
//
// Resolved through the same chain the issue path stamps onto an order, most
// specific first: a lead time set on the customer, then on its parent account,
// then on the customer's account group, then the account-wide default. `source`
// names which rule applied, so a form can show where the number came from rather
// than leaving a rep to guess.
//
// A lead time set on a parent account therefore governs every child account under
// it that has not set its own, which is how a head office's terms are given to its
// locations without repeating them on each one.
//
// This is a preview of a commitment, not the commitment itself. An order takes its
// own `ship_by_date` when it is issued and keeps it afterwards, so changing a lead
// time here moves what future orders will promise and leaves promises already made
// alone.
//
// This endpoint requires the permission: `customers:read`.
func (r *SaleCustomerService) GetLeadTime(ctx context.Context, id string, query SaleCustomerGetLeadTimeParams, opts ...option.RequestOption) (res *CustomerLeadTime, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s/lead-time", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to create a customer.
//
// The properties BillToAddress, CustomerTypeGroupID, DefaultCarrierID,
// DefaultPaymentTermID, DefaultShippingTermID, Name, ShipToAddress are required.
type CreateCustomerRequestParam struct {
	// Address details supplied when creating an address, either on its own or inline
	// on another resource.
	//
	// A few requests, such as shipping rate estimates, take these same fields for a
	// one-off address that is never saved to the account.
	BillToAddress AddressInputParam `json:"bill_to_address,omitzero" api:"required"`
	// ID of the account group of type `type_group` that categorizes this customer (for
	// example "Distributors").
	CustomerTypeGroupID string `json:"customer_type_group_id" api:"required"`
	// ID of the carrier used on this customer's orders when the order does not specify
	// one.
	DefaultCarrierID string `json:"default_carrier_id" api:"required"`
	// ID of the payment term used on this customer's orders when the order does not
	// specify one.
	DefaultPaymentTermID string `json:"default_payment_term_id" api:"required"`
	// ID of the shipping term used on this customer's orders when the order does not
	// specify one.
	DefaultShippingTermID string `json:"default_shipping_term_id" api:"required"`
	// The customer's business name, as shown throughout the app and on documents.
	Name string `json:"name" api:"required"`
	// Address details supplied when creating an address, either on its own or inline
	// on another resource.
	//
	// A few requests, such as shipping rate estimates, take these same fields for a
	// one-off address that is never saved to the account.
	ShipToAddress AddressInputParam `json:"ship_to_address,omitzero" api:"required"`
	// Carrier billing account number charged when `carrier_billing_type` is
	// `third_party`.
	CarrierBillingAccount param.Opt[string] `json:"carrier_billing_account,omitzero"`
	// The ID of the account user to credit as the sales rep on this customer's orders.
	//
	// Must be an account user on your own account.
	DefaultSalesRepID param.Opt[string] `json:"default_sales_rep_id,omitzero"`
	// ID of the carrier service level used when an order takes its carrier from this
	// customer's default.
	DefaultServiceLevelID param.Opt[string] `json:"default_service_level_id,omitzero"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// Calendar days between an order being issued and it being due to ship.
	//
	// Sets each order's `ship_by_date` when it is issued. Leave unset to inherit the
	// parent account's lead time, then the customer's account group lead time, then
	// the account default.
	LeadTimeDays param.Opt[int64] `json:"lead_time_days,omitzero"`
	// Free-form note about the customer.
	Note param.Opt[string] `json:"note,omitzero"`
	// Human-readable customer number used to identify the account, distinct from the
	// `id`.
	//
	// Must be unique within your account. If omitted, the next sequential number is
	// assigned automatically.
	Number param.Opt[string] `json:"number,omitzero"`
	// Phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The operating calendar naming the days this customer's dock accepts freight.
	//
	// Sits in the same chain as lead_time_days: leaving it unset falls through to the
	// customer's group, then the account default, then Monday to Friday. A promised
	// delivery date is never worked back from a day nobody is there to receive on.
	ReceiveCalendarID param.Opt[string] `json:"receive_calendar_id,omitzero"`
	// Website URL.
	URL param.Opt[string] `json:"url,omitzero"`
	// Who pays the carrier for shipments.
	//
	// - `sender`: the shipper (you) pays the carrier.
	// - `third_party`: a third party is billed, using `carrier_billing_account`.
	//
	// Any of "sender", "third_party".
	CarrierBillingType CreateCustomerRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// How sales commission applies to this customer's orders.
	//
	//   - `commission_exempt`: this customer's orders are exempt from sales commission.
	//   - `commission_applied`: sales commission is calculated on this customer's
	//     orders.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CreateCustomerRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	CreditLimit QuantityInputParam `json:"credit_limit,omitzero"`
	// IDs of the account groups of type `pricing_group` to assign to this customer,
	// used to apply pricing rules.
	CustomerPriceGroupIDs []string `json:"customer_price_group_ids,omitzero"`
	// Priority used to pre-fill new orders for this customer.
	//
	// Any of "low", "normal", "high".
	DefaultPriority CreateCustomerRequestDefaultPriority `json:"default_priority,omitzero"`
	// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
	// documents with this customer.
	//
	// Any of "enabled", "disabled".
	EdiStatus CreateCustomerRequestEdiStatus `json:"edi_status,omitzero"`
	// Whether this customer is billed for freight on their orders.
	//
	// - `free_freight`: the customer is not billed for freight.
	// - `billed_freight`: freight is billed to the customer.
	//
	// Freight is also waived when the customer's type group, one of its price groups,
	// or a product line the ordered products belong to is `free_freight`.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy CreateCustomerRequestFreightPolicy `json:"freight_policy,omitzero"`
	// How this customer's orders are produced.
	//
	//   - `make_to_stock`: their order history feeds the production-schedule forecast,
	//     so stock is built ahead of their demand.
	//   - `make_to_order`: their history is left out of the forecast; their orders are
	//     produced only once placed, and fit into the schedule on their own ship-by
	//     dates.
	//
	// Leave unset to inherit the customer's account group policy, then the
	// make-to-stock default.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy CreateCustomerRequestFulfillmentPolicy `json:"fulfillment_policy,omitzero"`
	// The customer's account standing.
	//
	//   - `normal`: standard account with no restrictions.
	//   - `preferred`: account flagged for prioritized handling.
	//   - `hold_shipment`: the customer's shipments should be held, typically over a
	//     credit problem, while orders can still be placed.
	//   - `hold_all`: all activity for the customer should be held.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status CreateCustomerRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r CreateCustomerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCustomerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCustomerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays the carrier for shipments.
//
// - `sender`: the shipper (you) pays the carrier.
// - `third_party`: a third party is billed, using `carrier_billing_account`.
type CreateCustomerRequestCarrierBillingType string

const (
	CreateCustomerRequestCarrierBillingTypeSender     CreateCustomerRequestCarrierBillingType = "sender"
	CreateCustomerRequestCarrierBillingTypeThirdParty CreateCustomerRequestCarrierBillingType = "third_party"
)

// How sales commission applies to this customer's orders.
//
//   - `commission_exempt`: this customer's orders are exempt from sales commission.
//   - `commission_applied`: sales commission is calculated on this customer's
//     orders.
type CreateCustomerRequestCommissionPolicy string

const (
	CreateCustomerRequestCommissionPolicyCommissionApplied CreateCustomerRequestCommissionPolicy = "commission_applied"
	CreateCustomerRequestCommissionPolicyCommissionExempt  CreateCustomerRequestCommissionPolicy = "commission_exempt"
)

// Priority used to pre-fill new orders for this customer.
type CreateCustomerRequestDefaultPriority string

const (
	CreateCustomerRequestDefaultPriorityLow    CreateCustomerRequestDefaultPriority = "low"
	CreateCustomerRequestDefaultPriorityNormal CreateCustomerRequestDefaultPriority = "normal"
	CreateCustomerRequestDefaultPriorityHigh   CreateCustomerRequestDefaultPriority = "high"
)

// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
// documents with this customer.
type CreateCustomerRequestEdiStatus string

const (
	CreateCustomerRequestEdiStatusEnabled  CreateCustomerRequestEdiStatus = "enabled"
	CreateCustomerRequestEdiStatusDisabled CreateCustomerRequestEdiStatus = "disabled"
)

// Whether this customer is billed for freight on their orders.
//
// - `free_freight`: the customer is not billed for freight.
// - `billed_freight`: freight is billed to the customer.
//
// Freight is also waived when the customer's type group, one of its price groups,
// or a product line the ordered products belong to is `free_freight`.
type CreateCustomerRequestFreightPolicy string

const (
	CreateCustomerRequestFreightPolicyFreeFreight   CreateCustomerRequestFreightPolicy = "free_freight"
	CreateCustomerRequestFreightPolicyBilledFreight CreateCustomerRequestFreightPolicy = "billed_freight"
)

// How this customer's orders are produced.
//
//   - `make_to_stock`: their order history feeds the production-schedule forecast,
//     so stock is built ahead of their demand.
//   - `make_to_order`: their history is left out of the forecast; their orders are
//     produced only once placed, and fit into the schedule on their own ship-by
//     dates.
//
// Leave unset to inherit the customer's account group policy, then the
// make-to-stock default.
type CreateCustomerRequestFulfillmentPolicy string

const (
	CreateCustomerRequestFulfillmentPolicyMakeToStock CreateCustomerRequestFulfillmentPolicy = "make_to_stock"
	CreateCustomerRequestFulfillmentPolicyMakeToOrder CreateCustomerRequestFulfillmentPolicy = "make_to_order"
)

// The customer's account standing.
//
//   - `normal`: standard account with no restrictions.
//   - `preferred`: account flagged for prioritized handling.
//   - `hold_shipment`: the customer's shipments should be held, typically over a
//     credit problem, while orders can still be placed.
//   - `hold_all`: all activity for the customer should be held.
type CreateCustomerRequestStatus string

const (
	CreateCustomerRequestStatusNormal       CreateCustomerRequestStatus = "normal"
	CreateCustomerRequestStatusPreferred    CreateCustomerRequestStatus = "preferred"
	CreateCustomerRequestStatusHoldShipment CreateCustomerRequestStatus = "hold_shipment"
	CreateCustomerRequestStatusHoldAll      CreateCustomerRequestStatus = "hold_all"
)

// The ship-by lead time a new order for this customer would be committed to.
type CustomerLeadTime struct {
	// A named grouping of customer accounts, used for pricing rules or to categorize
	// accounts.
	//
	// A customer carries at most one group of type `type_group` as its customer type,
	// plus any number of groups of type `pricing_group`. Membership of either kind can
	// scope a volume discount to the customer and open up product lines for it to
	// order from.
	AccountGroup AccountGroup `json:"account_group" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Customer Entity `json:"customer" api:"required"`
	// Calendar days between an order being issued and it being due to ship.
	//
	// `0` means same-day: an order issued today would be due to ship today.
	Days int64 `json:"days" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_lead_time".
	Object CustomerLeadTimeObject `json:"object" api:"required"`
	// A business you sell to, with its contact details, default fulfillment settings,
	// and order policies.
	ParentCustomer Customer `json:"parent_customer" api:"required"`
	// Which rule in the chain produced this lead time.
	//
	// - `customer`: a lead time set on the customer itself.
	// - `parent_customer`: inherited from the customer's parent account.
	// - `account_group`: inherited from the customer's account group.
	// - `account`: the account-wide fallback.
	//
	// The shared `manual` value cannot appear here: it means a promised date was set
	// on one specific order, which is a fact about that order rather than about the
	// customer.
	//
	// Any of "customer", "parent_customer", "account_group", "account", "manual",
	// "order_lead_time", "order_ship_by".
	Source CustomerLeadTimeSource `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountGroup   respjson.Field
		Customer       respjson.Field
		Days           respjson.Field
		Object         respjson.Field
		ParentCustomer respjson.Field
		Source         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerLeadTime) RawJSON() string { return r.JSON.raw }
func (r *CustomerLeadTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CustomerLeadTimeObject string

const (
	CustomerLeadTimeObjectCustomerLeadTime CustomerLeadTimeObject = "customer_lead_time"
)

// Which rule in the chain produced this lead time.
//
// - `customer`: a lead time set on the customer itself.
// - `parent_customer`: inherited from the customer's parent account.
// - `account_group`: inherited from the customer's account group.
// - `account`: the account-wide fallback.
//
// The shared `manual` value cannot appear here: it means a promised date was set
// on one specific order, which is a fact about that order rather than about the
// customer.
type CustomerLeadTimeSource string

const (
	CustomerLeadTimeSourceCustomer       CustomerLeadTimeSource = "customer"
	CustomerLeadTimeSourceParentCustomer CustomerLeadTimeSource = "parent_customer"
	CustomerLeadTimeSourceAccountGroup   CustomerLeadTimeSource = "account_group"
	CustomerLeadTimeSourceAccount        CustomerLeadTimeSource = "account"
	CustomerLeadTimeSourceManual         CustomerLeadTimeSource = "manual"
	CustomerLeadTimeSourceOrderLeadTime  CustomerLeadTimeSource = "order_lead_time"
	CustomerLeadTimeSourceOrderShipBy    CustomerLeadTimeSource = "order_ship_by"
)

// Request to partially update a customer.
type UpdateCustomerRequestParam struct {
	// ID of an existing address to use as the default billing address.
	//
	// The address is linked to the customer's account if it is not already.
	BillToAddressID param.Opt[string] `json:"bill_to_address_id,omitzero"`
	// Carrier billing account number charged when `carrier_billing_type` is
	// `third_party`.
	CarrierBillingAccount param.Opt[string] `json:"carrier_billing_account,omitzero"`
	// The ID of the account user to credit as the sales rep on this customer's orders.
	//
	// Must be an account user on your own account.
	DefaultSalesRepID param.Opt[string] `json:"default_sales_rep_id,omitzero"`
	// ID of the carrier service level used when an order takes its carrier from this
	// customer's default.
	DefaultServiceLevelID param.Opt[string] `json:"default_service_level_id,omitzero"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// Calendar days between an order being issued and it being due to ship.
	//
	// Sets each order's `ship_by_date` when it is issued. Clear it to inherit the
	// parent account's lead time, then the customer's account group lead time, then
	// the account default.
	LeadTimeDays param.Opt[int64] `json:"lead_time_days,omitzero"`
	// Free-form note about the customer.
	Note param.Opt[string] `json:"note,omitzero"`
	// Phone number.
	Phone param.Opt[string] `json:"phone,omitzero"`
	// The operating calendar naming the days this customer's dock accepts freight.
	// Clearing it returns the customer to their group's calendar, then the account
	// default.
	ReceiveCalendarID param.Opt[string] `json:"receive_calendar_id,omitzero"`
	// ID of an existing address to use as the default shipping address.
	//
	// The address is linked to the customer's account if it is not already.
	ShipToAddressID param.Opt[string] `json:"ship_to_address_id,omitzero"`
	// Website URL.
	URL param.Opt[string] `json:"url,omitzero"`
	// ID of the account group of type `type_group` that categorizes this customer (for
	// example "Distributors").
	CustomerTypeGroupID param.Opt[string] `json:"customer_type_group_id,omitzero"`
	// ID of the carrier used on this customer's orders when the order does not specify
	// one.
	DefaultCarrierID param.Opt[string] `json:"default_carrier_id,omitzero"`
	// ID of the payment term used on this customer's orders when the order does not
	// specify one.
	DefaultPaymentTermID param.Opt[string] `json:"default_payment_term_id,omitzero"`
	// ID of the shipping term used on this customer's orders when the order does not
	// specify one.
	DefaultShippingTermID param.Opt[string] `json:"default_shipping_term_id,omitzero"`
	// The customer's business name, as shown throughout the app and on documents.
	Name param.Opt[string] `json:"name,omitzero"`
	// Human-readable customer number used to identify the account, distinct from the
	// `id`.
	//
	// Must be unique within your account.
	Number param.Opt[string] `json:"number,omitzero"`
	// How this customer's orders are produced.
	//
	//   - `make_to_stock`: their order history feeds the production-schedule forecast,
	//     so stock is built ahead of their demand.
	//   - `make_to_order`: their history is left out of the forecast; their orders are
	//     produced only once placed, and fit into the schedule on their own ship-by
	//     dates.
	//
	// Clearing it returns the customer to their account group policy, then the
	// make-to-stock default.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy UpdateCustomerRequestFulfillmentPolicy `json:"fulfillment_policy,omitzero"`
	// Who pays the carrier for shipments.
	//
	// - `sender`: the shipper (you) pays the carrier.
	// - `third_party`: a third party is billed, using `carrier_billing_account`.
	//
	// Any of "sender", "third_party".
	CarrierBillingType UpdateCustomerRequestCarrierBillingType `json:"carrier_billing_type,omitzero"`
	// How sales commission applies to this customer's orders.
	//
	//   - `commission_exempt`: this customer's orders are exempt from sales commission.
	//   - `commission_applied`: sales commission is calculated on this customer's
	//     orders.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpdateCustomerRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	CreditLimit QuantityInputParam `json:"credit_limit,omitzero"`
	// IDs of the account groups of type `pricing_group` to assign to this customer,
	// used to apply pricing rules.
	//
	// When provided, replaces the customer's full set of existing price groups.
	CustomerPriceGroupIDs []string `json:"customer_price_group_ids,omitzero"`
	// Priority used to pre-fill new orders for this customer.
	//
	// Any of "low", "normal", "high".
	DefaultPriority UpdateCustomerRequestDefaultPriority `json:"default_priority,omitzero"`
	// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
	// documents with this customer.
	//
	// Any of "enabled", "disabled".
	EdiStatus UpdateCustomerRequestEdiStatus `json:"edi_status,omitzero"`
	// Whether this customer is billed for freight on their orders.
	//
	// - `free_freight`: the customer is not billed for freight.
	// - `billed_freight`: freight is billed to the customer.
	//
	// Freight is also waived when the customer's type group, one of its price groups,
	// or a product line the ordered products belong to is `free_freight`.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpdateCustomerRequestFreightPolicy `json:"freight_policy,omitzero"`
	// The customer's account standing.
	//
	//   - `normal`: standard account with no restrictions.
	//   - `preferred`: account flagged for prioritized handling.
	//   - `hold_shipment`: the customer's shipments should be held, typically over a
	//     credit problem, while orders can still be placed.
	//   - `hold_all`: all activity for the customer should be held.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status UpdateCustomerRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r UpdateCustomerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCustomerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCustomerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays the carrier for shipments.
//
// - `sender`: the shipper (you) pays the carrier.
// - `third_party`: a third party is billed, using `carrier_billing_account`.
type UpdateCustomerRequestCarrierBillingType string

const (
	UpdateCustomerRequestCarrierBillingTypeSender     UpdateCustomerRequestCarrierBillingType = "sender"
	UpdateCustomerRequestCarrierBillingTypeThirdParty UpdateCustomerRequestCarrierBillingType = "third_party"
)

// How sales commission applies to this customer's orders.
//
//   - `commission_exempt`: this customer's orders are exempt from sales commission.
//   - `commission_applied`: sales commission is calculated on this customer's
//     orders.
type UpdateCustomerRequestCommissionPolicy string

const (
	UpdateCustomerRequestCommissionPolicyCommissionApplied UpdateCustomerRequestCommissionPolicy = "commission_applied"
	UpdateCustomerRequestCommissionPolicyCommissionExempt  UpdateCustomerRequestCommissionPolicy = "commission_exempt"
)

// Priority used to pre-fill new orders for this customer.
type UpdateCustomerRequestDefaultPriority string

const (
	UpdateCustomerRequestDefaultPriorityLow    UpdateCustomerRequestDefaultPriority = "low"
	UpdateCustomerRequestDefaultPriorityNormal UpdateCustomerRequestDefaultPriority = "normal"
	UpdateCustomerRequestDefaultPriorityHigh   UpdateCustomerRequestDefaultPriority = "high"
)

// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
// documents with this customer.
type UpdateCustomerRequestEdiStatus string

const (
	UpdateCustomerRequestEdiStatusEnabled  UpdateCustomerRequestEdiStatus = "enabled"
	UpdateCustomerRequestEdiStatusDisabled UpdateCustomerRequestEdiStatus = "disabled"
)

// Whether this customer is billed for freight on their orders.
//
// - `free_freight`: the customer is not billed for freight.
// - `billed_freight`: freight is billed to the customer.
//
// Freight is also waived when the customer's type group, one of its price groups,
// or a product line the ordered products belong to is `free_freight`.
type UpdateCustomerRequestFreightPolicy string

const (
	UpdateCustomerRequestFreightPolicyFreeFreight   UpdateCustomerRequestFreightPolicy = "free_freight"
	UpdateCustomerRequestFreightPolicyBilledFreight UpdateCustomerRequestFreightPolicy = "billed_freight"
)

// How this customer's orders are produced.
//
//   - `make_to_stock`: their order history feeds the production-schedule forecast,
//     so stock is built ahead of their demand.
//   - `make_to_order`: their history is left out of the forecast; their orders are
//     produced only once placed, and fit into the schedule on their own ship-by
//     dates.
//
// Clearing it returns the customer to their account group policy, then the
// make-to-stock default.
type UpdateCustomerRequestFulfillmentPolicy string

const (
	UpdateCustomerRequestFulfillmentPolicyMakeToStock UpdateCustomerRequestFulfillmentPolicy = "make_to_stock"
	UpdateCustomerRequestFulfillmentPolicyMakeToOrder UpdateCustomerRequestFulfillmentPolicy = "make_to_order"
)

// The customer's account standing.
//
//   - `normal`: standard account with no restrictions.
//   - `preferred`: account flagged for prioritized handling.
//   - `hold_shipment`: the customer's shipments should be held, typically over a
//     credit problem, while orders can still be placed.
//   - `hold_all`: all activity for the customer should be held.
type UpdateCustomerRequestStatus string

const (
	UpdateCustomerRequestStatusNormal       UpdateCustomerRequestStatus = "normal"
	UpdateCustomerRequestStatusPreferred    UpdateCustomerRequestStatus = "preferred"
	UpdateCustomerRequestStatusHoldShipment UpdateCustomerRequestStatus = "hold_shipment"
	UpdateCustomerRequestStatusHoldAll      UpdateCustomerRequestStatus = "hold_all"
)

type SaleCustomerDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleCustomerDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleCustomerDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleCustomerNewParams struct {
	// Request to create a customer.
	CreateCustomerRequest CreateCustomerRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.carrier.service_levels",
	// "freight_preferences.service_level", "defaults.payment_term",
	// "defaults.shipping_term", "defaults.sales_rep", "defaults.sales_rep.user",
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit",
	// "credit_limit.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleCustomerNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCustomerRequest)
}
func (r *SaleCustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleCustomerNewParams]'s query parameters as `url.Values`.
func (r SaleCustomerNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleCustomerGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.carrier.service_levels",
	// "freight_preferences.service_level", "defaults.payment_term",
	// "defaults.shipping_term", "defaults.sales_rep", "defaults.sales_rep.user",
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit",
	// "credit_limit.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleCustomerGetParams]'s query parameters as `url.Values`.
func (r SaleCustomerGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleCustomerUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.carrier.service_levels",
	// "freight_preferences.service_level", "defaults.payment_term",
	// "defaults.shipping_term", "defaults.sales_rep", "defaults.sales_rep.user",
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit",
	// "credit_limit.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a customer.
	UpdateCustomerRequest UpdateCustomerRequestParam
	paramObj
}

func (r SaleCustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCustomerRequest)
}
func (r *SaleCustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleCustomerUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleCustomerUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleCustomerListParams struct {
	// Filter to customers with any address in this city (exact match).
	//
	// When combined with `state` or `postal_code`, a single address must match all
	// provided values.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to customers created at or before this timestamp (inclusive).
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter to customers with any address in this postal code (exact match).
	PostalCode param.Opt[string] `query:"postal_code,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter to customers created at or after this timestamp (inclusive).
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Filter to customers with any address in this state (exact match).
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// Filter by default carrier IDs.
	CarrierIDs []string `query:"carrier_ids,omitzero" json:"-"`
	// Filter by the commission policy set on the customer itself.
	//
	// Policies inherited from the customer's type group or price groups are not
	// considered here.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionStatusCodes []string `query:"commission_status_codes,omitzero" json:"-"`
	// Filter by customer type group IDs (the account group of type `type_group`
	// returned in the customer's `type` field).
	CustomerGroupIDs []string `query:"customer_group_ids,omitzero" json:"-"`
	// Filter by the freight policy set on the customer itself.
	//
	// Policies inherited from the customer's type group or price groups are not
	// considered here.
	//
	// Any of "free_freight", "billed_freight".
	FreightStatusCodes []string `query:"freight_status_codes,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.carrier.service_levels",
	// "freight_preferences.service_level", "defaults.payment_term",
	// "defaults.shipping_term", "defaults.sales_rep", "defaults.sales_rep.user",
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit",
	// "credit_limit.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by whether the customer has child accounts.
	//
	// Any of "parent", "non_parent".
	ParentAccountStatus SaleCustomerListParamsParentAccountStatus `query:"parent_account_status,omitzero" json:"-"`
	// Filter by default payment term IDs.
	PaymentTermIDs []string `query:"payment_term_ids,omitzero" json:"-"`
	// Filter to customers that belong to any of these pricing groups.
	PricingGroupIDs []string `query:"pricing_group_ids,omitzero" json:"-"`
	// Filter to customers whose default sales rep is one of these account users.
	SalesRepIDs []string `query:"sales_rep_ids,omitzero" json:"-"`
	// Filter by default service level IDs.
	ServiceLevelIDs []string `query:"service_level_ids,omitzero" json:"-"`
	// Filter by default shipping term IDs.
	ShippingTermIDs []string `query:"shipping_term_ids,omitzero" json:"-"`
	// Filter by the customer's account standing.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	StatusCodes []string `query:"status_codes,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleCustomerListParams]'s query parameters as `url.Values`.
func (r SaleCustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by whether the customer has child accounts.
type SaleCustomerListParamsParentAccountStatus string

const (
	SaleCustomerListParamsParentAccountStatusParent    SaleCustomerListParamsParentAccountStatus = "parent"
	SaleCustomerListParamsParentAccountStatusNonParent SaleCustomerListParamsParentAccountStatus = "non_parent"
)

type SaleCustomerGetLeadTimeParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account_group", "parent_customer".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleCustomerGetLeadTimeParams]'s query parameters as
// `url.Values`.
func (r SaleCustomerGetLeadTimeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
