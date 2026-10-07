// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage carriers and their Shippo integrations.
//
// OperationCarrierService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationCarrierService] method instead.
type OperationCarrierService struct {
	options []option.RequestOption
	// List and manage service levels (shipping service levels).
	ServiceLevels OperationCarrierServiceLevelService
}

// NewOperationCarrierService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationCarrierService(opts ...option.RequestOption) (r OperationCarrierService) {
	r = OperationCarrierService{}
	r.options = opts
	r.ServiceLevels = NewOperationCarrierServiceLevelService(opts...)
	return
}

// Creates a shipping carrier your account can ship orders with.
//
// Supplying a Shippo-supported code (`fedex`, `ups`, `usps`) connects a Shippo
// carrier account and creates a service level for every service that carrier
// offers, each hidden from the customer portal until you make it visible. This
// requires an active Shippo integration on the account and is skipped entirely for
// sandbox accounts, which get a carrier record with no service levels and no live
// rating.
//
// This endpoint requires the permission: `carriers:create`.
func (r *OperationCarrierService) New(ctx context.Context, params OperationCarrierNewParams, opts ...option.RequestOption) (res *Carrier, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/carriers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a carrier by ID.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `carriers:read`.
func (r *OperationCarrierService) Get(ctx context.Context, id string, query OperationCarrierGetParams, opts ...option.RequestOption) (res *Carrier, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a carrier's name and customer portal visibility.
//
// Only these two attributes can change: a carrier's code and account number are
// fixed at creation, and system-owned carriers cannot be updated at all.
//
// This endpoint requires the permission: `carriers:update`.
func (r *OperationCarrierService) Update(ctx context.Context, id string, params OperationCarrierUpdateParams, opts ...option.RequestOption) (res *Carrier, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of the carriers available to the current account.
//
// This covers the carriers you have created plus the platform-provided system
// carriers that every account shares.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `carriers:read`.
func (r *OperationCarrierService) List(ctx context.Context, query OperationCarrierListParams, opts ...option.RequestOption) (res *ListCarrier, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/carriers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a carrier and all of its service levels.
//
// If the carrier is connected through Shippo, its Shippo carrier account is
// deactivated. System-owned carriers cannot be deleted.
//
// This endpoint requires the permission: `carriers:delete`.
func (r *OperationCarrierService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationCarrierDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a carrier.
//
// The property Name is required.
type CreateCarrierRequestParam struct {
	// Human-readable name for the carrier.
	//
	// Must not match another carrier already visible to your account, including the
	// system-provided ones.
	Name string `json:"name" api:"required"`
	// Your account number with this carrier.
	//
	// Required when `code` is `ups` or `usps`, whose carrier accounts are connected to
	// Shippo using this number; FedEx authorizes through OAuth instead, so no account
	// number is needed.
	AccountNumber param.Opt[string] `json:"account_number,omitzero"`
	// Well-known carrier code.
	//
	// Providing a Shippo-supported code (`fedex`, `ups`, `usps`) connects the carrier
	// through Shippo and syncs its service levels; the other codes, such as
	// `will_call` and `delivery`, simply describe a self-managed shipping method. Omit
	// the code entirely when none of them fit. The code cannot be changed after the
	// carrier is created.
	//
	// Any of "fedex", "ups", "usps", "will_call", "delivery", "ltl", "ltl1",
	// "freight_collect".
	Code CreateCarrierRequestCode `json:"code,omitzero"`
	// Whether customers can see and select this carrier at checkout in the customer
	// portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CreateCarrierRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r CreateCarrierRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCarrierRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCarrierRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Well-known carrier code.
//
// Providing a Shippo-supported code (`fedex`, `ups`, `usps`) connects the carrier
// through Shippo and syncs its service levels; the other codes, such as
// `will_call` and `delivery`, simply describe a self-managed shipping method. Omit
// the code entirely when none of them fit. The code cannot be changed after the
// carrier is created.
type CreateCarrierRequestCode string

const (
	CreateCarrierRequestCodeFedex          CreateCarrierRequestCode = "fedex"
	CreateCarrierRequestCodeUps            CreateCarrierRequestCode = "ups"
	CreateCarrierRequestCodeUsps           CreateCarrierRequestCode = "usps"
	CreateCarrierRequestCodeWillCall       CreateCarrierRequestCode = "will_call"
	CreateCarrierRequestCodeDelivery       CreateCarrierRequestCode = "delivery"
	CreateCarrierRequestCodeLtl            CreateCarrierRequestCode = "ltl"
	CreateCarrierRequestCodeLtl1           CreateCarrierRequestCode = "ltl1"
	CreateCarrierRequestCodeFreightCollect CreateCarrierRequestCode = "freight_collect"
)

// Whether customers can see and select this carrier at checkout in the customer
// portal.
type CreateCarrierRequestCustomerPortalVisibility string

const (
	CreateCarrierRequestCustomerPortalVisibilityVisible CreateCarrierRequestCustomerPortalVisibility = "visible"
	CreateCarrierRequestCustomerPortalVisibilityHidden  CreateCarrierRequestCustomerPortalVisibility = "hidden"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListCarrier struct {
	// Resources in this page.
	Data []Carrier `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListCarrierObject `json:"object" api:"required"`
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
func (r ListCarrier) RawJSON() string { return r.JSON.raw }
func (r *ListCarrier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListCarrierObject string

const (
	ListCarrierObjectList ListCarrierObject = "list"
)

// Request to update a carrier.
type UpdateCarrierRequestParam struct {
	// Human-readable name for the carrier.
	//
	// Must not match another carrier already visible to your account, including the
	// system-provided ones.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether customers can see and select this carrier at checkout in the customer
	// portal.
	//
	// Each of the carrier's service levels carries its own customer portal visibility,
	// which this does not change.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UpdateCarrierRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r UpdateCarrierRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCarrierRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCarrierRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether customers can see and select this carrier at checkout in the customer
// portal.
//
// Each of the carrier's service levels carries its own customer portal visibility,
// which this does not change.
type UpdateCarrierRequestCustomerPortalVisibility string

const (
	UpdateCarrierRequestCustomerPortalVisibilityVisible UpdateCarrierRequestCustomerPortalVisibility = "visible"
	UpdateCarrierRequestCustomerPortalVisibilityHidden  UpdateCarrierRequestCustomerPortalVisibility = "hidden"
)

type OperationCarrierDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationCarrierDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationCarrierDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationCarrierNewParams struct {
	// Request to create a carrier.
	CreateCarrierRequest CreateCarrierRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationCarrierNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCarrierRequest)
}
func (r *OperationCarrierNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierNewParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierGetParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a carrier.
	UpdateCarrierRequest UpdateCarrierRequestParam
	paramObj
}

func (r OperationCarrierUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCarrierRequest)
}
func (r *OperationCarrierUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierListParams struct {
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
	// Any of "owner", "owner.account", "service_levels".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierListParams]'s query parameters as
// `url.Values`.
func (r OperationCarrierListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
