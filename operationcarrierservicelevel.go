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

// List and manage service levels (shipping service levels).
//
// OperationCarrierServiceLevelService contains methods and other services that
// help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationCarrierServiceLevelService] method instead.
type OperationCarrierServiceLevelService struct {
	options []option.RequestOption
}

// NewOperationCarrierServiceLevelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationCarrierServiceLevelService(opts ...option.RequestOption) (r OperationCarrierServiceLevelService) {
	r = OperationCarrierServiceLevelService{}
	r.options = opts
	return
}

// Adds a shipping service level to a carrier.
//
// Use this for self-managed carriers, or to add a service a connected carrier does
// not publish. Service levels created here are never removed by a later sync of
// the carrier's services.
//
// This endpoint requires the permission: `carriers:create`.
func (r *OperationCarrierServiceLevelService) New(ctx context.Context, carrierID string, params OperationCarrierServiceLevelNewParams, opts ...option.RequestOption) (res *ServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if carrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels", url.PathEscape(carrierID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a service level by ID.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `carriers:read`.
func (r *OperationCarrierServiceLevelService) Get(ctx context.Context, id string, params OperationCarrierServiceLevelGetParams, opts ...option.RequestOption) (res *ServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CarrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels/%s", url.PathEscape(params.CarrierID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Updates a service level's name, code, customer portal visibility, or default
// status.
//
// Only the fields you send are changed. System-owned service levels cannot be
// updated.
//
// This endpoint requires the permission: `carriers:update`.
func (r *OperationCarrierServiceLevelService) Update(ctx context.Context, id string, params OperationCarrierServiceLevelUpdateParams, opts ...option.RequestOption) (res *ServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CarrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels/%s", url.PathEscape(params.CarrierID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of the service levels a carrier offers.
//
// Use this rather than the `service_levels` field on the carrier itself when a
// carrier has more than a handful of services, since that inline list is capped.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `carriers:read`.
func (r *OperationCarrierServiceLevelService) List(ctx context.Context, carrierID string, query OperationCarrierServiceLevelListParams, opts ...option.RequestOption) (res *ListServiceLevel, err error) {
	opts = slices.Concat(r.options, opts)
	if carrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels", url.PathEscape(carrierID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a service level so it can no longer be selected on
// shipments.
//
// System-owned service levels and the carrier's default service level cannot be
// deleted; to remove a default, first clear its `is_default` flag or promote
// another service level in its place.
//
// This endpoint requires the permission: `carriers:delete`.
func (r *OperationCarrierServiceLevelService) Delete(ctx context.Context, id string, body OperationCarrierServiceLevelDeleteParams, opts ...option.RequestOption) (res *OperationCarrierServiceLevelDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CarrierID == "" {
		err = errors.New("missing required carrier_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/carriers/%s/service-levels/%s", url.PathEscape(body.CarrierID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a service level.
//
// The properties Code, IsDefault, Name are required.
type CreateServiceLevelRequestParam struct {
	// Carrier-specific code identifying this service level (e.g. `fedex_ground`).
	//
	// Must be unique among the carrier's service levels, and is returned as the
	// service level's `service_level_token`.
	Code string `json:"code" api:"required"`
	// Whether this becomes the carrier's default service level, pre-selected when the
	// carrier is chosen.
	//
	// Each carrier has at most one default; setting this to `true` clears the
	// carrier's existing default.
	IsDefault bool `json:"is_default" api:"required"`
	// Human-readable name for the service level, shown to customers at checkout when
	// the service level is visible.
	Name string `json:"name" api:"required"`
	// Business days this service typically takes in transit, used to work an order's
	// ship-by date back from a promised delivery date.
	//
	// A fallback: when a carrier can rate the lane, the transit it quotes is used
	// instead. Leave unset for carriers that can be rated, and set it for those that
	// cannot (freight, will-call), where it is the only transit the system will have.
	DefaultTransitDays param.Opt[int64] `json:"default_transit_days,omitzero"`
	// Whether customers can see and select this service level at checkout in the
	// customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility CreateServiceLevelRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r CreateServiceLevelRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateServiceLevelRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateServiceLevelRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether customers can see and select this service level at checkout in the
// customer portal.
type CreateServiceLevelRequestCustomerPortalVisibility string

const (
	CreateServiceLevelRequestCustomerPortalVisibilityVisible CreateServiceLevelRequestCustomerPortalVisibility = "visible"
	CreateServiceLevelRequestCustomerPortalVisibilityHidden  CreateServiceLevelRequestCustomerPortalVisibility = "hidden"
)

// Request to update a service level.
type UpdateServiceLevelRequestParam struct {
	// Business days this service typically takes in transit, used to work an order's
	// ship-by date back from a promised delivery date.
	//
	// A fallback: when a carrier can rate the lane, the transit it quotes is used
	// instead. Set to null to remove it, which leaves transit unknown for lanes the
	// carrier cannot rate.
	DefaultTransitDays param.Opt[int64] `json:"default_transit_days,omitzero"`
	// Carrier-specific code identifying this service level (e.g. `fedex_ground`).
	//
	// Must be unique among the carrier's service levels. For a service level synced
	// from a connected carrier the `service_level_token` used for rating is fixed by
	// the carrier and a code change does not affect it; for one you created yourself,
	// the token follows the code.
	Code param.Opt[string] `json:"code,omitzero"`
	// Whether this is the carrier's default service level, pre-selected when the
	// carrier is chosen.
	//
	// Each carrier has at most one default; setting this to `true` clears the
	// carrier's existing default.
	IsDefault param.Opt[bool] `json:"is_default,omitzero"`
	// Human-readable name for the service level, shown to customers at checkout when
	// the service level is visible.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether customers can see and select this service level at checkout in the
	// customer portal.
	//
	// Any of "visible", "hidden".
	CustomerPortalVisibility UpdateServiceLevelRequestCustomerPortalVisibility `json:"customer_portal_visibility,omitzero"`
	paramObj
}

func (r UpdateServiceLevelRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateServiceLevelRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateServiceLevelRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether customers can see and select this service level at checkout in the
// customer portal.
type UpdateServiceLevelRequestCustomerPortalVisibility string

const (
	UpdateServiceLevelRequestCustomerPortalVisibilityVisible UpdateServiceLevelRequestCustomerPortalVisibility = "visible"
	UpdateServiceLevelRequestCustomerPortalVisibilityHidden  UpdateServiceLevelRequestCustomerPortalVisibility = "hidden"
)

type OperationCarrierServiceLevelDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationCarrierServiceLevelDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationCarrierServiceLevelDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationCarrierServiceLevelNewParams struct {
	// Request to create a service level.
	CreateServiceLevelRequest CreateServiceLevelRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationCarrierServiceLevelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateServiceLevelRequest)
}
func (r *OperationCarrierServiceLevelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierServiceLevelNewParams]'s query parameters
// as `url.Values`.
func (r OperationCarrierServiceLevelNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelGetParams struct {
	CarrierID string `path:"carrier_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierServiceLevelGetParams]'s query parameters
// as `url.Values`.
func (r OperationCarrierServiceLevelGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelUpdateParams struct {
	CarrierID string `path:"carrier_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a service level.
	UpdateServiceLevelRequest UpdateServiceLevelRequestParam
	paramObj
}

func (r OperationCarrierServiceLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateServiceLevelRequest)
}
func (r *OperationCarrierServiceLevelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationCarrierServiceLevelUpdateParams]'s query
// parameters as `url.Values`.
func (r OperationCarrierServiceLevelUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelListParams struct {
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
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationCarrierServiceLevelListParams]'s query parameters
// as `url.Values`.
func (r OperationCarrierServiceLevelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationCarrierServiceLevelDeleteParams struct {
	CarrierID string `path:"carrier_id" api:"required" json:"-"`
	paramObj
}
