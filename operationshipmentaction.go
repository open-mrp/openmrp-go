// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage shipments, shipment lines, and shipping operations.
//
// OperationShipmentActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationShipmentActionService] method instead.
type OperationShipmentActionService struct {
	options []option.RequestOption
}

// NewOperationShipmentActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationShipmentActionService(opts ...option.RequestOption) (r OperationShipmentActionService) {
	r = OperationShipmentActionService{}
	r.options = opts
	return
}

// Compares shipping rates across all of the account's carriers and service levels
// for the given addresses and parcels.
//
// Returns options sorted by rate ascending, after applying the account's freight
// rules: freight-exempt product lines or customers and free-freight shipping terms
// return no options, a flat-rate shipping term replaces carrier rates with the
// flat rate, and a met free-shipping minimum order value zeroes the rate on
// eligible options.
//
// Live carrier rates require the Shippo integration. Carriers that are not linked
// to a live-rating account are returned at a rate of `0`, while carriers that are
// linked but whose rates cannot be fetched are left out of the results entirely.
// Customer portal callers only see carriers and service levels that have been
// enabled for the portal.
//
// This endpoint requires the permissions: `shipments:read`, `customers:read`,
// `suppliers:read`.
func (r *OperationShipmentActionService) RateShop(ctx context.Context, body OperationShipmentActionRateShopParams, opts ...option.RequestOption) (res *RateShopResult, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/shipments/actions/rate-shop"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListRateShopOption struct {
	// Resources in this page.
	Data []RateShopOption `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListRateShopOptionObject `json:"object" api:"required"`
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
func (r ListRateShopOption) RawJSON() string { return r.JSON.raw }
func (r *ListRateShopOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListRateShopOptionObject string

const (
	ListRateShopOptionObjectList ListRateShopOptionObject = "list"
)

// A parcel's weight and dimensions for shipping rate calculations.
//
// The properties Height, Length, Weight, Width are required.
type ParcelInputParam struct {
	// Parcel height in inches.
	Height float64 `json:"height" api:"required"`
	// Parcel length in inches.
	Length float64 `json:"length" api:"required"`
	// Parcel weight in pounds.
	Weight float64 `json:"weight" api:"required"`
	// Parcel width in inches.
	Width float64 `json:"width" api:"required"`
	paramObj
}

func (r ParcelInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ParcelInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParcelInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single carrier and service level option returned by rate shopping.
type RateShopOption struct {
	// A shipping carrier configured for fulfilling orders.
	//
	// Carriers with a Shippo-supported `code` (`fedex`, `ups`, `usps`) are connected
	// through Shippo for live rating and label purchase; other carriers represent
	// self-managed shipping methods such as will call or local delivery.
	Carrier Carrier `json:"carrier" api:"required"`
	// Estimated number of days until delivery, when the carrier provides an estimate.
	EstimatedDays int64 `json:"estimated_days" api:"required"`
	// Resource type identifier.
	//
	// Any of "rate_shop_option".
	Object RateShopOptionObject `json:"object" api:"required"`
	// Quoted shipping rate for this carrier and service level.
	//
	// `0` when the carrier is not linked to a live-rating account, or when the
	// shipping term's free-shipping minimum order value has been met and this option
	// qualifies for free shipping. When the customer's shipping term applies a flat
	// rate, that amount replaces the rate on every option that is not already free.
	Rate float64 `json:"rate" api:"required"`
	// A shipping speed or method offered by a carrier, such as ground or overnight.
	//
	// Carriers connected through Shippo have their service levels synced from the
	// carrier itself; any carrier can also have service levels you create by hand.
	ServiceLevel ServiceLevel `json:"service_level" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Carrier       respjson.Field
		EstimatedDays respjson.Field
		Object        respjson.Field
		Rate          respjson.Field
		ServiceLevel  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RateShopOption) RawJSON() string { return r.JSON.raw }
func (r *RateShopOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RateShopOptionObject string

const (
	RateShopOptionObjectRateShopOption RateShopOptionObject = "rate_shop_option"
)

// Request to rate shop across carriers.
//
// The properties Parcels, ToAddress are required.
type RateShopRequestParam struct {
	// Parcels to rate shop.
	Parcels []ParcelInputParam `json:"parcels,omitzero" api:"required"`
	// Address details supplied when creating an address, either on its own or inline
	// on another resource.
	//
	// A few requests, such as shipping rate estimates, take these same fields for a
	// one-off address that is never saved to the account.
	ToAddress AddressInputParam `json:"to_address,omitzero" api:"required"`
	// ID of the customer the shipment is for, used to apply the customer's freight
	// policy and default shipping term.
	//
	// A customer that is freight exempt through its own policy or through one of its
	// groups, or whose shipping term is free freight, returns no options with
	// `exemption_type` set to `freight_exempt`; a flat-rate shipping term replaces
	// carrier rates with the flat rate. Omitting the customer skips all of these rules
	// and returns plain carrier rates.
	CustomerID param.Opt[string] `json:"customer_id,omitzero"`
	// Total value of the order, used to evaluate the free-shipping minimum order value
	// on the customer's shipping term.
	//
	// Free shipping applies only when the total is strictly above the threshold, and
	// only for the service levels the shipping term allows.
	OrderTotal param.Opt[float64] `json:"order_total,omitzero"`
	// Address details supplied when creating an address, either on its own or inline
	// on another resource.
	//
	// A few requests, such as shipping rate estimates, take these same fields for a
	// one-off address that is never saved to the account.
	FromAddress AddressInputParam `json:"from_address,omitzero"`
	// Product lines of the items being shipped, used to apply freight exemptions.
	//
	// If any listed product line is freight exempt, no options are returned and
	// `exemption_type` is `freight_exempt`.
	ProductLineIDs []string `json:"product_line_ids,omitzero"`
	paramObj
}

func (r RateShopRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RateShopRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RateShopRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The carrier and service level options returned by rate shopping, along with the
// freight rule that shaped their rates.
type RateShopResult struct {
	// Why a special freight outcome was applied to these options, if any.
	//
	//   - `freight_exempt`: the order is exempt from freight; no options are returned.
	//   - `minimum_order_met`: the customer's shipping term sets a free-shipping minimum
	//     order value and the order total exceeded it, so options are rated at zero. If
	//     the shipping term restricts free shipping to specific service levels, only
	//     those options are zeroed and the rest keep their carrier or flat rate.
	//   - `flat_rate`: the customer's shipping term applies a flat shipping rate, which
	//     replaced every option's carrier rate.
	//   - `none`: standard carrier rates apply with no exemption.
	//
	// Any of "freight_exempt", "minimum_order_met", "flat_rate", "none".
	ExemptionType RateShopResultExemptionType `json:"exemption_type" api:"required"`
	// Flat shipping amount applied to the options.
	//
	// Set when the customer's shipping term applies a flat rate, including when a met
	// free-shipping minimum has already rated some options at zero.
	FlatRate float64 `json:"flat_rate" api:"required"`
	// Resource type identifier.
	//
	// Any of "rate_shop_result".
	Object RateShopResultObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Options ListRateShopOption `json:"options" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExemptionType respjson.Field
		FlatRate      respjson.Field
		Object        respjson.Field
		Options       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RateShopResult) RawJSON() string { return r.JSON.raw }
func (r *RateShopResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why a special freight outcome was applied to these options, if any.
//
//   - `freight_exempt`: the order is exempt from freight; no options are returned.
//   - `minimum_order_met`: the customer's shipping term sets a free-shipping minimum
//     order value and the order total exceeded it, so options are rated at zero. If
//     the shipping term restricts free shipping to specific service levels, only
//     those options are zeroed and the rest keep their carrier or flat rate.
//   - `flat_rate`: the customer's shipping term applies a flat shipping rate, which
//     replaced every option's carrier rate.
//   - `none`: standard carrier rates apply with no exemption.
type RateShopResultExemptionType string

const (
	RateShopResultExemptionTypeFreightExempt   RateShopResultExemptionType = "freight_exempt"
	RateShopResultExemptionTypeMinimumOrderMet RateShopResultExemptionType = "minimum_order_met"
	RateShopResultExemptionTypeFlatRate        RateShopResultExemptionType = "flat_rate"
	RateShopResultExemptionTypeNone            RateShopResultExemptionType = "none"
)

// Resource type identifier.
type RateShopResultObject string

const (
	RateShopResultObjectRateShopResult RateShopResultObject = "rate_shop_result"
)

type OperationShipmentActionRateShopParams struct {
	// Request to rate shop across carriers.
	RateShopRequest RateShopRequestParam
	paramObj
}

func (r OperationShipmentActionRateShopParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RateShopRequest)
}
func (r *OperationShipmentActionRateShopParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
