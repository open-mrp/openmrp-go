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

// List and manage account prices.
//
// SaleAccountPriceService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleAccountPriceService] method instead.
type SaleAccountPriceService struct {
	options []option.RequestOption
	// List and manage account prices.
	Actions SaleAccountPriceActionService
}

// NewSaleAccountPriceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleAccountPriceService(opts ...option.RequestOption) (r SaleAccountPriceService) {
	r = SaleAccountPriceService{}
	r.options = opts
	r.Actions = NewSaleAccountPriceActionService(opts...)
	return
}

// Creates a customer-specific price for a product line.
//
// When a sales order line for the recipient matches the price's product line and
// attributes, this price replaces the unit price the line would otherwise be
// given, including the effect of any volume discount. If more than one account
// price matches a line, the most recently created one wins.
//
// This endpoint requires the permission: `discounts:create`.
func (r *SaleAccountPriceService) New(ctx context.Context, params SaleAccountPriceNewParams, opts ...option.RequestOption) (res *AccountPrice, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/account-prices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns an account price by ID.
//
// A customer portal user can only retrieve a price whose recipient is their own
// account or its parent; any other price is reported as not found.
//
// This endpoint requires the permissions: `discounts:read`, `customers:read`,
// `suppliers:read`.
func (r *SaleAccountPriceService) Get(ctx context.Context, id string, query SaleAccountPriceGetParams, opts ...option.RequestOption) (res *AccountPrice, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-prices/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account price.
//
// Only the provided fields are changed. If `category_ids` or `attribute_ids` are
// provided, they replace the existing set entirely.
//
// Order lines that have already been priced keep the unit price they were given;
// the new price applies to lines priced after the change.
//
// This endpoint requires the permission: `discounts:update`.
func (r *SaleAccountPriceService) Update(ctx context.Context, id string, params SaleAccountPriceUpdateParams, opts ...option.RequestOption) (res *AccountPrice, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-prices/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of account prices, newest first.
//
// The search term matches the recipient customer's name or their customer number.
// Customer portal users always see only the prices that apply to their own
// account, whatever `recipient_account_id` is set to.
//
// This endpoint requires the permissions: `discounts:read`, `customers:read`,
// `suppliers:read`.
func (r *SaleAccountPriceService) List(ctx context.Context, query SaleAccountPriceListParams, opts ...option.RequestOption) (res *ListAccountPrice, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/account-prices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an account price.
//
// The price's category and attribute associations and its rate record are removed
// with it. Deletion is permanent; further requests against the deleted ID return
// an error.
//
// Order lines that have already been priced keep the unit price they were given;
// only lines priced after the deletion revert to standard pricing.
//
// This endpoint requires the permission: `discounts:delete`.
func (r *SaleAccountPriceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleAccountPriceDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/account-prices/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A customer-specific price for a product line.
//
// When a sales order line matches an account price, that price replaces the unit
// price the line would otherwise be given — including the effect of any volume
// discount — rather than discounting it. If more than one account price matches a
// line, the most recently created one wins.
type AccountPrice struct {
	// Account price ID.
	ID string `json:"id" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Attributes ListAttribute `json:"attributes" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Categories ListItemCategory `json:"categories" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "account_price".
	Object AccountPriceObject `json:"object" api:"required"`
	// A named grouping of related products in your catalog.
	//
	// A product line carries the default commission and freight policies for the
	// products assigned to it, along with the unit group that determines how those
	// products are measured. Product lines are also the unit that catalog access is
	// granted over, for both customers and account groups.
	ProductLine ProductLine `json:"product_line" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	Rate Rate `json:"rate" api:"required"`
	// A business you sell to, with its contact details, default fulfillment settings,
	// and order policies.
	RecipientAccount Customer `json:"recipient_account" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Attributes       respjson.Field
		Categories       respjson.Field
		CreatedAt        respjson.Field
		Object           respjson.Field
		ProductLine      respjson.Field
		Rate             respjson.Field
		RecipientAccount respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPrice) RawJSON() string { return r.JSON.raw }
func (r *AccountPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountPriceObject string

const (
	AccountPriceObjectAccountPrice AccountPriceObject = "account_price"
)

// A shipping carrier configured for fulfilling orders.
//
// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
// through Shippo for live rating and label purchase; other carriers represent
// self-managed shipping methods such as will call or local delivery.
type Carrier struct {
	// Carrier ID.
	ID string `json:"id" api:"required"`
	// Your account number with this carrier.
	//
	// UPS and USPS carrier accounts are connected to Shippo using this number; FedEx
	// carriers authorize through OAuth instead, so their account number is not used to
	// connect them.
	AccountNumber string `json:"account_number" api:"required"`
	// Well-known carrier identifier, set only for recognized carriers and absent for
	// custom ones.
	//
	//   - `fedex`, `ups`, `usps`: integrated carriers managed through Shippo (live
	//     rating and labels).
	//   - `will_call`: customer picks the order up; no carrier shipment.
	//   - `delivery`: delivered by your own vehicles/drivers.
	//   - `ltl`, `ltl1`: less-than-truckload freight carriers.
	//   - `freight_collect`: freight billed to and arranged by the receiver.
	//
	// Any of "fedex", "ups", "usps", "will_call", "delivery", "ltl", "ltl1",
	// "freight_collect".
	Code CarrierCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether customers can see and select this carrier at checkout in the customer
	// portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CarrierCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Soft-delete timestamp.
	DeletedAt time.Time `json:"deleted_at" api:"required" format:"date-time"`
	// Human-readable name for the carrier, unique among the carriers visible to your
	// account.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "carrier".
	Object CarrierObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ServiceLevels ListServiceLevel `json:"service_levels" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AccountNumber            respjson.Field
		Code                     respjson.Field
		CreatedAt                respjson.Field
		CustomerPortalVisibility respjson.Field
		DeletedAt                respjson.Field
		Name                     respjson.Field
		Object                   respjson.Field
		Owner                    respjson.Field
		ServiceLevels            respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Carrier) RawJSON() string { return r.JSON.raw }
func (r *Carrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Well-known carrier identifier, set only for recognized carriers and absent for
// custom ones.
//
//   - `fedex`, `ups`, `usps`: integrated carriers managed through Shippo (live
//     rating and labels).
//   - `will_call`: customer picks the order up; no carrier shipment.
//   - `delivery`: delivered by your own vehicles/drivers.
//   - `ltl`, `ltl1`: less-than-truckload freight carriers.
//   - `freight_collect`: freight billed to and arranged by the receiver.
type CarrierCode string

const (
	CarrierCodeFedex          CarrierCode = "fedex"
	CarrierCodeUps            CarrierCode = "ups"
	CarrierCodeUsps           CarrierCode = "usps"
	CarrierCodeWillCall       CarrierCode = "will_call"
	CarrierCodeDelivery       CarrierCode = "delivery"
	CarrierCodeLtl            CarrierCode = "ltl"
	CarrierCodeLtl1           CarrierCode = "ltl1"
	CarrierCodeFreightCollect CarrierCode = "freight_collect"
)

// Whether customers can see and select this carrier at checkout in the customer
// portal.
type CarrierCustomerPortalVisibility string

const (
	CarrierCustomerPortalVisibilityVisible CarrierCustomerPortalVisibility = "visible"
	CarrierCustomerPortalVisibilityHidden  CarrierCustomerPortalVisibility = "hidden"
)

// Resource type identifier.
type CarrierObject string

const (
	CarrierObjectCarrier CarrierObject = "carrier"
)

// Request to create an account price.
//
// The properties ProductLineID, Rate, RecipientAccountID are required.
type CreateAccountPriceRequestParam struct {
	// ID of the product line whose products this price applies to.
	ProductLineID string `json:"product_line_id" api:"required"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	Rate RateInputParam `json:"rate,omitzero" api:"required"`
	// ID of the customer this price is offered to.
	//
	// A price recorded against a parent customer account also applies to orders placed
	// by its child accounts.
	RecipientAccountID string `json:"recipient_account_id" api:"required"`
	// Attribute IDs to constrain this price to.
	//
	// When set, the price applies only to items that have every listed attribute.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// Item category IDs to record on this price.
	//
	// Order pricing matches an account price on its product line and attributes only,
	// so categories recorded here do not narrow which products the price applies to.
	CategoryIDs []string `json:"category_ids,omitzero"`
	paramObj
}

func (r CreateAccountPriceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAccountPriceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAccountPriceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A business you sell to, with its contact details, default fulfillment settings,
// and order policies.
type Customer struct {
	// Customer ID.
	ID string `json:"id" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	BillToAddress Address `json:"bill_to_address" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ChildAccounts *ListCustomer `json:"child_accounts" api:"required"`
	// How sales commission applies to this customer's orders.
	//
	//   - `commission_exempt`: this customer's orders are exempt from sales commission.
	//   - `commission_applied`: sales commission is calculated on this customer's
	//     orders.
	//
	// The customer counts as exempt if this field, its `type` group, or any of its
	// `price_groups` is `commission_exempt`. Exempt customers never have a sales rep
	// assigned automatically when an order is created without one.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CustomerCommissionPolicy `json:"commission_policy" api:"required"`
	// Customer contact information.
	ContactInfo CustomerContactInfo `json:"contact_info" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	CreditLimit Quantity `json:"credit_limit" api:"required"`
	// Values used to fill in a new sales order for this customer when the order does
	// not supply its own.
	Defaults CustomerDefaults `json:"defaults" api:"required"`
	// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
	// documents with this customer.
	//
	// Any of "enabled", "disabled".
	EdiStatus CustomerEdiStatus `json:"edi_status" api:"required"`
	// Customer freight and carrier settings.
	FreightPreferences CustomerFreightPreferences `json:"freight_preferences" api:"required"`
	// The customer's business name, as shown throughout the app and on documents.
	Name string `json:"name" api:"required"`
	// Free-form note about the customer.
	Note string `json:"note" api:"required"`
	// Customer notification settings.
	NotificationPreferences CustomerNotificationPreferences `json:"notification_preferences" api:"required"`
	// Human-readable customer number used to identify the account, distinct from the
	// `id`.
	//
	// Unique within your account.
	Number string `json:"number" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer".
	Object CustomerObject `json:"object" api:"required"`
	// A business you sell to, with its contact details, default fulfillment settings,
	// and order policies.
	ParentAccount *Customer `json:"parent_account" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	PriceGroups ListAccountGroup `json:"price_groups" api:"required"`
	// The customer's position in the account hierarchy.
	//
	// - `standalone`: no parent or child accounts.
	// - `parent`: has one or more child accounts (see `child_accounts`).
	// - `child`: belongs to a parent account (see `parent_account`).
	//
	// Any of "standalone", "parent", "child".
	RelationshipType CustomerRelationshipType `json:"relationship_type" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	ShipToAddress Address `json:"ship_to_address" api:"required"`
	// The customer's account standing.
	//
	//   - `normal`: standard account with no restrictions.
	//   - `preferred`: account flagged for prioritized handling.
	//   - `hold_shipment`: the customer's shipments should be held, typically over a
	//     credit problem, while orders can still be placed.
	//   - `hold_all`: all activity for the customer should be held.
	//
	// The hold statuses are advisory: OpenMRP flags the customer's orders as being on
	// credit hold, but requests to create orders or shipments for the customer are not
	// rejected.
	//
	// Any of "normal", "preferred", "hold_shipment", "hold_all".
	Status CustomerStatus `json:"status" api:"required"`
	// A named grouping of customer accounts, used for pricing rules or to categorize
	// accounts.
	//
	// A customer carries at most one group of type `type_group` as its customer type,
	// plus any number of groups of type `pricing_group`. Membership of either kind can
	// scope a volume discount to the customer and open up product lines for it to
	// order from.
	Type AccountGroup `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		BillToAddress           respjson.Field
		ChildAccounts           respjson.Field
		CommissionPolicy        respjson.Field
		ContactInfo             respjson.Field
		CreatedAt               respjson.Field
		CreditLimit             respjson.Field
		Defaults                respjson.Field
		EdiStatus               respjson.Field
		FreightPreferences      respjson.Field
		Name                    respjson.Field
		Note                    respjson.Field
		NotificationPreferences respjson.Field
		Number                  respjson.Field
		Object                  respjson.Field
		ParentAccount           respjson.Field
		PriceGroups             respjson.Field
		RelationshipType        respjson.Field
		ShipToAddress           respjson.Field
		Status                  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How sales commission applies to this customer's orders.
//
//   - `commission_exempt`: this customer's orders are exempt from sales commission.
//   - `commission_applied`: sales commission is calculated on this customer's
//     orders.
//
// The customer counts as exempt if this field, its `type` group, or any of its
// `price_groups` is `commission_exempt`. Exempt customers never have a sales rep
// assigned automatically when an order is created without one.
type CustomerCommissionPolicy string

const (
	CustomerCommissionPolicyCommissionApplied CustomerCommissionPolicy = "commission_applied"
	CustomerCommissionPolicyCommissionExempt  CustomerCommissionPolicy = "commission_exempt"
)

// Whether EDI (Electronic Data Interchange) is enabled for exchanging orders and
// documents with this customer.
type CustomerEdiStatus string

const (
	CustomerEdiStatusEnabled  CustomerEdiStatus = "enabled"
	CustomerEdiStatusDisabled CustomerEdiStatus = "disabled"
)

// Resource type identifier.
type CustomerObject string

const (
	CustomerObjectCustomer CustomerObject = "customer"
)

// The customer's position in the account hierarchy.
//
// - `standalone`: no parent or child accounts.
// - `parent`: has one or more child accounts (see `child_accounts`).
// - `child`: belongs to a parent account (see `parent_account`).
type CustomerRelationshipType string

const (
	CustomerRelationshipTypeStandalone CustomerRelationshipType = "standalone"
	CustomerRelationshipTypeParent     CustomerRelationshipType = "parent"
	CustomerRelationshipTypeChild      CustomerRelationshipType = "child"
)

// The customer's account standing.
//
//   - `normal`: standard account with no restrictions.
//   - `preferred`: account flagged for prioritized handling.
//   - `hold_shipment`: the customer's shipments should be held, typically over a
//     credit problem, while orders can still be placed.
//   - `hold_all`: all activity for the customer should be held.
//
// The hold statuses are advisory: OpenMRP flags the customer's orders as being on
// credit hold, but requests to create orders or shipments for the customer are not
// rejected.
type CustomerStatus string

const (
	CustomerStatusNormal       CustomerStatus = "normal"
	CustomerStatusPreferred    CustomerStatus = "preferred"
	CustomerStatusHoldShipment CustomerStatus = "hold_shipment"
	CustomerStatusHoldAll      CustomerStatus = "hold_all"
)

// Customer contact information.
type CustomerContactInfo struct {
	// Email address.
	Email string `json:"email" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_contact_info".
	Object CustomerContactInfoObject `json:"object" api:"required"`
	// Phone number.
	Phone string `json:"phone" api:"required"`
	// Website URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Object      respjson.Field
		Phone       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerContactInfo) RawJSON() string { return r.JSON.raw }
func (r *CustomerContactInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CustomerContactInfoObject string

const (
	CustomerContactInfoObjectCustomerContactInfo CustomerContactInfoObject = "customer_contact_info"
)

// Values used to fill in a new sales order for this customer when the order does
// not supply its own.
type CustomerDefaults struct {
	// How this customer's orders are produced.
	//
	//   - `make_to_stock`: their order history feeds the production-schedule forecast,
	//     so stock is built ahead of their demand.
	//   - `make_to_order`: their history is left out of the forecast; their orders are
	//     produced only once placed, and fit into the schedule on their own ship-by
	//     dates.
	//
	// With none set here the customer inherits its account group's policy, then falls
	// back to make-to-stock.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy CustomerDefaultsFulfillmentPolicy `json:"fulfillment_policy" api:"required"`
	// Calendar days between an order being issued and it being due to ship.
	//
	// Sets each order's `ship_by_date` when it is issued. With none set here the
	// customer inherits its parent account's lead time, then its account group's, then
	// the account default.
	LeadTimeDays int64 `json:"lead_time_days" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_defaults".
	Object CustomerDefaultsObject `json:"object" api:"required"`
	// A payment term describing when payment is due (e.g. `Net 30`), assignable to
	// customers, sales orders, purchase orders, and invoices.
	PaymentTerm PaymentTerm `json:"payment_term" api:"required"`
	// Priority level used to order work on sales orders, purchase orders, and picks.
	//
	// The levels are platform-provided and the same for every account, so they cannot
	// be created, renamed, or removed. A customer can carry a default priority that
	// pre-fills new orders for them.
	Priority Priority `json:"priority" api:"required"`
	// The operating calendar naming the days this customer's dock accepts freight.
	//
	// A promised delivery date is worked back from a day the customer can actually
	// receive on. With none set here the customer inherits its account group's
	// calendar, then the account default, then Monday to Friday.
	ReceiveCalendarID string `json:"receive_calendar_id" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the `user`
	// sub-resource, which is shared across every account the user belongs to.
	SalesRep AccountUser `json:"sales_rep" api:"required"`
	// A named freight pricing rule that decides what a buyer pays for shipping.
	//
	// A customer's default shipping term is evaluated whenever freight is quoted for
	// one of their orders. Freight exemptions on the customer, its type group, or any
	// of its price groups are checked first and zero the freight charge before the
	// shipping term is considered.
	ShippingTerm ShippingTerm `json:"shipping_term" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FulfillmentPolicy respjson.Field
		LeadTimeDays      respjson.Field
		Object            respjson.Field
		PaymentTerm       respjson.Field
		Priority          respjson.Field
		ReceiveCalendarID respjson.Field
		SalesRep          respjson.Field
		ShippingTerm      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerDefaults) RawJSON() string { return r.JSON.raw }
func (r *CustomerDefaults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How this customer's orders are produced.
//
//   - `make_to_stock`: their order history feeds the production-schedule forecast,
//     so stock is built ahead of their demand.
//   - `make_to_order`: their history is left out of the forecast; their orders are
//     produced only once placed, and fit into the schedule on their own ship-by
//     dates.
//
// With none set here the customer inherits its account group's policy, then falls
// back to make-to-stock.
type CustomerDefaultsFulfillmentPolicy string

const (
	CustomerDefaultsFulfillmentPolicyMakeToStock CustomerDefaultsFulfillmentPolicy = "make_to_stock"
	CustomerDefaultsFulfillmentPolicyMakeToOrder CustomerDefaultsFulfillmentPolicy = "make_to_order"
)

// Resource type identifier.
type CustomerDefaultsObject string

const (
	CustomerDefaultsObjectCustomerDefaults CustomerDefaultsObject = "customer_defaults"
)

// Customer freight and carrier settings.
type CustomerFreightPreferences struct {
	// Carrier billing account number charged when `billing_type` is `third_party`.
	BillingAccount string `json:"billing_account" api:"required"`
	// Who pays the carrier for shipments.
	//
	// - `sender`: the shipper (you) pays the carrier.
	// - `third_party`: a third party is billed, using `billing_account`.
	//
	// Any of "sender", "third_party".
	BillingType CustomerFreightPreferencesBillingType `json:"billing_type" api:"required"`
	// A shipping carrier configured for fulfilling orders.
	//
	// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
	// through Shippo for live rating and label purchase; other carriers represent
	// self-managed shipping methods such as will call or local delivery.
	Carrier Carrier `json:"carrier" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_freight_preferences".
	Object CustomerFreightPreferencesObject `json:"object" api:"required"`
	// A shipping speed or method offered by a carrier, such as ground or overnight.
	//
	// Carriers connected through Shippo have their service levels synced from the
	// carrier itself; any carrier can also have service levels you create by hand.
	ServiceLevel ServiceLevel `json:"service_level" api:"required"`
	// Freight policy applied to this customer's orders.
	//
	// - `free_freight`: the customer is not billed for freight.
	// - `billed_freight`: freight is billed to the customer.
	//
	// Freight is waived when this field, the customer's `type` group, any of its
	// `price_groups`, or any product line the ordered products belong to is
	// `free_freight`, so a shipment can come back freight-exempt even while this field
	// is `billed_freight`.
	//
	// Any of "free_freight", "billed_freight".
	Status CustomerFreightPreferencesStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BillingAccount respjson.Field
		BillingType    respjson.Field
		Carrier        respjson.Field
		Object         respjson.Field
		ServiceLevel   respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerFreightPreferences) RawJSON() string { return r.JSON.raw }
func (r *CustomerFreightPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays the carrier for shipments.
//
// - `sender`: the shipper (you) pays the carrier.
// - `third_party`: a third party is billed, using `billing_account`.
type CustomerFreightPreferencesBillingType string

const (
	CustomerFreightPreferencesBillingTypeSender     CustomerFreightPreferencesBillingType = "sender"
	CustomerFreightPreferencesBillingTypeThirdParty CustomerFreightPreferencesBillingType = "third_party"
)

// Resource type identifier.
type CustomerFreightPreferencesObject string

const (
	CustomerFreightPreferencesObjectCustomerFreightPreferences CustomerFreightPreferencesObject = "customer_freight_preferences"
)

// Freight policy applied to this customer's orders.
//
// - `free_freight`: the customer is not billed for freight.
// - `billed_freight`: freight is billed to the customer.
//
// Freight is waived when this field, the customer's `type` group, any of its
// `price_groups`, or any product line the ordered products belong to is
// `free_freight`, so a shipment can come back freight-exempt even while this field
// is `billed_freight`.
type CustomerFreightPreferencesStatus string

const (
	CustomerFreightPreferencesStatusFreeFreight   CustomerFreightPreferencesStatus = "free_freight"
	CustomerFreightPreferencesStatusBilledFreight CustomerFreightPreferencesStatus = "billed_freight"
)

// Customer notification settings.
type CustomerNotificationPreferences struct {
	// Whether anyone is set up to receive invoice emails for this customer.
	//
	// Derived from the customer's notification recipients: true when at least one of
	// them is configured for invoice notifications.
	AcceptsInvoiceEmails bool `json:"accepts_invoice_emails" api:"required"`
	// Resource type identifier.
	//
	// Any of "customer_notification_preferences".
	Object CustomerNotificationPreferencesObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptsInvoiceEmails respjson.Field
		Object               respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerNotificationPreferences) RawJSON() string { return r.JSON.raw }
func (r *CustomerNotificationPreferences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CustomerNotificationPreferencesObject string

const (
	CustomerNotificationPreferencesObjectCustomerNotificationPreferences CustomerNotificationPreferencesObject = "customer_notification_preferences"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAccountPrice struct {
	// Resources in this page.
	Data []AccountPrice `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountPriceObject `json:"object" api:"required"`
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
func (r ListAccountPrice) RawJSON() string { return r.JSON.raw }
func (r *ListAccountPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountPriceObject string

const (
	ListAccountPriceObjectList ListAccountPriceObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListCustomer struct {
	// Resources in this page.
	Data []Customer `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListCustomerObject `json:"object" api:"required"`
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
func (r ListCustomer) RawJSON() string { return r.JSON.raw }
func (r *ListCustomer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListCustomerObject string

const (
	ListCustomerObjectList ListCustomerObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListServiceLevel struct {
	// Resources in this page.
	Data []ServiceLevel `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListServiceLevelObject `json:"object" api:"required"`
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
func (r ListServiceLevel) RawJSON() string { return r.JSON.raw }
func (r *ListServiceLevel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListServiceLevelObject string

const (
	ListServiceLevelObjectList ListServiceLevelObject = "list"
)

// A payment term describing when payment is due (e.g. `Net 30`), assignable to
// customers, sales orders, purchase orders, and invoices.
type PaymentTerm struct {
	// Payment term ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name (e.g. `Net 30`), unique among the payment terms visible to your
	// account.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "payment_term".
	Object PaymentTermObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Whether this payment term is still in active use.
	//
	// Payment terms created through the API are always `active`, and no endpoint
	// changes a term's status. List Payment Terms returns inactive terms alongside
	// active ones, so filter them out yourself if you only want the ones still on
	// offer.
	//
	// Any of "active", "inactive".
	Status PaymentTermStatus `json:"status" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentTerm) RawJSON() string { return r.JSON.raw }
func (r *PaymentTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PaymentTermObject string

const (
	PaymentTermObjectPaymentTerm PaymentTermObject = "payment_term"
)

// Whether this payment term is still in active use.
//
// Payment terms created through the API are always `active`, and no endpoint
// changes a term's status. List Payment Terms returns inactive terms alongside
// active ones, so filter them out yourself if you only want the ones still on
// offer.
type PaymentTermStatus string

const (
	PaymentTermStatusActive   PaymentTermStatus = "active"
	PaymentTermStatusInactive PaymentTermStatus = "inactive"
)

// Priority level used to order work on sales orders, purchase orders, and picks.
//
// The levels are platform-provided and the same for every account, so they cannot
// be created, renamed, or removed. A customer can carry a default priority that
// pre-fills new orders for them.
type Priority struct {
	// Priority ID.
	ID string `json:"id" api:"required"`
	// Machine-readable code identifying the priority level.
	//
	// Other resources refer to a priority by this code rather than by its ID, such as
	// a sales order's `priority`, and it can be used in place of the ID when
	// retrieving a priority.
	//
	// Any of "low", "normal", "high".
	Code PriorityCode `json:"code" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the priority level.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "priority".
	Object PriorityObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Priority) RawJSON() string { return r.JSON.raw }
func (r *Priority) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable code identifying the priority level.
//
// Other resources refer to a priority by this code rather than by its ID, such as
// a sales order's `priority`, and it can be used in place of the ID when
// retrieving a priority.
type PriorityCode string

const (
	PriorityCodeLow    PriorityCode = "low"
	PriorityCodeNormal PriorityCode = "normal"
	PriorityCodeHigh   PriorityCode = "high"
)

// Resource type identifier.
type PriorityObject string

const (
	PriorityObjectPriority PriorityObject = "priority"
)

// A shipping speed or method offered by a carrier, such as ground or overnight.
//
// Carriers connected through Shippo have their service levels synced from the
// carrier itself; any carrier can also have service levels you create by hand.
type ServiceLevel struct {
	// Service level ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether customers can see and select this service level at checkout in the
	// customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility ServiceLevelCustomerPortalVisibility `json:"customer_portal_visibility" api:"required"`
	// Business days this service typically takes in transit, used to work an order's
	// ship-by date back from a promised delivery date.
	//
	// A fallback for lanes the carrier has not quoted. Null means transit is unknown
	// for this service rather than instant, so a ship-by date falls back to the
	// promised delivery date itself.
	DefaultTransitDays int64 `json:"default_transit_days" api:"required"`
	// Whether this is the carrier's default service level, pre-selected when the
	// carrier is chosen.
	//
	// Each carrier has at most one default; setting a new default clears the previous
	// one. A default service level cannot be deleted until another service level takes
	// its place or the flag is cleared.
	IsDefault bool `json:"is_default" api:"required"`
	// Human-readable name for the service level, shown to customers at checkout when
	// the service level is visible.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "service_level".
	Object ServiceLevelObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Carrier-specific code identifying this service level (e.g. `fedex_ground`,
	// `ups_next_day_air`).
	//
	// For service levels synced from a connected carrier this is the carrier's own
	// token, which is what rate shopping and label purchase are keyed on; for service
	// levels you create yourself it is the `code` you supplied.
	ServiceLevelToken string `json:"service_level_token" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		CreatedAt                respjson.Field
		CustomerPortalVisibility respjson.Field
		DefaultTransitDays       respjson.Field
		IsDefault                respjson.Field
		Name                     respjson.Field
		Object                   respjson.Field
		Owner                    respjson.Field
		ServiceLevelToken        respjson.Field
		UpdatedAt                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceLevel) RawJSON() string { return r.JSON.raw }
func (r *ServiceLevel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether customers can see and select this service level at checkout in the
// customer portal.
type ServiceLevelCustomerPortalVisibility string

const (
	ServiceLevelCustomerPortalVisibilityVisible ServiceLevelCustomerPortalVisibility = "visible"
	ServiceLevelCustomerPortalVisibilityHidden  ServiceLevelCustomerPortalVisibility = "hidden"
)

// Resource type identifier.
type ServiceLevelObject string

const (
	ServiceLevelObjectServiceLevel ServiceLevelObject = "service_level"
)

// A named freight pricing rule that decides what a buyer pays for shipping.
//
// A customer's default shipping term is evaluated whenever freight is quoted for
// one of their orders. Freight exemptions on the customer, its type group, or any
// of its price groups are checked first and zero the freight charge before the
// shipping term is considered.
type ShippingTerm struct {
	// Shipping term ID.
	ID string `json:"id" api:"required"`
	// When this shipping term was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	FlatRate Quantity `json:"flat_rate" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	FreeShippingServiceLevels ListServiceLevel `json:"free_shipping_service_levels" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	MinimumOrderValue Quantity `json:"minimum_order_value" api:"required"`
	// Human-readable name for the shipping term, used to identify it when assigning
	// shipping terms to customers and orders.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "shipping_term".
	Object ShippingTermObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Freight pricing model applied by this shipping term.
	//
	//   - `free_freight`: the buyer is never charged for shipping.
	//   - `flat_rate_freight`: the buyer is charged the fixed amount in `flat_rate`,
	//     regardless of what the carrier would have charged.
	//   - `carrier_rate_freight`: the buyer is charged the rate the carrier quotes for
	//     the order's carrier and service level.
	//
	// Any of "free_freight", "flat_rate_freight", "carrier_rate_freight".
	Type ShippingTermType `json:"type" api:"required"`
	// When this shipping term was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		CreatedAt                 respjson.Field
		FlatRate                  respjson.Field
		FreeShippingServiceLevels respjson.Field
		MinimumOrderValue         respjson.Field
		Name                      respjson.Field
		Object                    respjson.Field
		Owner                     respjson.Field
		Type                      respjson.Field
		UpdatedAt                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShippingTerm) RawJSON() string { return r.JSON.raw }
func (r *ShippingTerm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ShippingTermObject string

const (
	ShippingTermObjectShippingTerm ShippingTermObject = "shipping_term"
)

// Freight pricing model applied by this shipping term.
//
//   - `free_freight`: the buyer is never charged for shipping.
//   - `flat_rate_freight`: the buyer is charged the fixed amount in `flat_rate`,
//     regardless of what the carrier would have charged.
//   - `carrier_rate_freight`: the buyer is charged the rate the carrier quotes for
//     the order's carrier and service level.
type ShippingTermType string

const (
	ShippingTermTypeFreeFreight        ShippingTermType = "free_freight"
	ShippingTermTypeFlatRateFreight    ShippingTermType = "flat_rate_freight"
	ShippingTermTypeCarrierRateFreight ShippingTermType = "carrier_rate_freight"
)

// Request to partially update an account price.
type UpdateAccountPriceRequestParam struct {
	// ID of the product line whose products this price applies to.
	ProductLineID param.Opt[string] `json:"product_line_id,omitzero"`
	// ID of the customer this price is offered to.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero"`
	// Attribute IDs to constrain this price to.
	//
	// When provided, replaces the existing set of attributes entirely; an empty list
	// removes all attribute constraints.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// Item category IDs to record on this price.
	//
	// When provided, replaces the existing set of categories entirely; an empty list
	// removes them all. Categories are recorded only — they do not narrow which
	// products the price applies to.
	CategoryIDs []string `json:"category_ids,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	Rate RateInputParam `json:"rate,omitzero"`
	paramObj
}

func (r UpdateAccountPriceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAccountPriceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAccountPriceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountPriceDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleAccountPriceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleAccountPriceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleAccountPriceNewParams struct {
	// Request to create an account price.
	CreateAccountPriceRequest CreateAccountPriceRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "recipient_account", "product_line", "categories", "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleAccountPriceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAccountPriceRequest)
}
func (r *SaleAccountPriceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleAccountPriceNewParams]'s query parameters as
// `url.Values`.
func (r SaleAccountPriceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleAccountPriceGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "recipient_account", "product_line", "categories", "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleAccountPriceGetParams]'s query parameters as
// `url.Values`.
func (r SaleAccountPriceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleAccountPriceUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "recipient_account", "product_line", "categories", "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update an account price.
	UpdateAccountPriceRequest UpdateAccountPriceRequestParam
	paramObj
}

func (r SaleAccountPriceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAccountPriceRequest)
}
func (r *SaleAccountPriceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleAccountPriceUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleAccountPriceUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleAccountPriceListParams struct {
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
	// Filters results to prices whose recipient is this customer account.
	//
	// A child account also matches the prices recorded against its parent, since those
	// price its orders too.
	RecipientAccountID param.Opt[string] `query:"recipient_account_id,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "recipient_account", "product_line", "categories", "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleAccountPriceListParams]'s query parameters as
// `url.Values`.
func (r SaleAccountPriceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
