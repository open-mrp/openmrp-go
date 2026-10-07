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

// List and manage parts.
//
// CatalogPartService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogPartService] method instead.
type CatalogPartService struct {
	options []option.RequestOption
	// List and manage parts.
	Actions CatalogPartActionService
}

// NewCatalogPartService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogPartService(opts ...option.RequestOption) (r CatalogPartService) {
	r = CatalogPartService{}
	r.options = opts
	r.Actions = NewCatalogPartActionService(opts...)
	return
}

// Creates a part with the specified SKU and category.
//
// Inventory tracking for the new part starts at a zero on-hand quantity in the
// category's base unit.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `parts:create`.
func (r *CatalogPartService) New(ctx context.Context, params CatalogPartNewParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/parts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a part by ID.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `parts:read`.
func (r *CatalogPartService) Get(ctx context.Context, id string, query CatalogPartGetParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/parts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a part.
//
// Fields not provided retain their current values. The SKU, description, notes,
// and category are editable here; the part's attributes are changed through the
// item endpoints.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `parts:update`.
func (r *CatalogPartService) Update(ctx context.Context, id string, params CatalogPartUpdateParams, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/parts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of parts for the current account, most recently created
// first.
//
// The `q` search term matches the part's SKU or description. When it is supplied,
// the parts whose SKU matches it most closely are returned first, ordered by
// creation time within each level of match.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `parts:read`.
func (r *CatalogPartService) List(ctx context.Context, query CatalogPartListParams, opts ...option.RequestOption) (res *ListPart, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/parts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a part.
//
// This is a soft delete: the part is marked deleted and no longer returned by
// other endpoints, but the record is retained. Deleting an already-deleted part
// returns an error.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `parts:delete`.
func (r *CatalogPartService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Part, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/parts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a part.
//
// The properties CategoryID, SKU are required.
type CreatePartRequestParam struct {
	// ID of the item category to place the part in.
	//
	// The category's unit group determines the base unit used for the part's rates
	// (`unit_value`, `unit_cost`, `burn_rate`).
	CategoryID string `json:"category_id" api:"required"`
	// Stock keeping unit code for the part.
	//
	// Must be unique within the account; creating a part with a SKU already used by
	// another item fails with a conflict error.
	SKU string `json:"sku" api:"required"`
	// Free-form description of the part.
	Description param.Opt[string] `json:"description,omitzero"`
	// Free-form notes about the part.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// IDs of existing attributes to link to the part at creation time.
	//
	// Each attribute's property must be one the part's category carries; an attribute
	// from any other property fails the whole request.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitCost RateInputParam `json:"unit_cost,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreatePartRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreatePartRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreatePartRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListPart struct {
	// Resources in this page.
	Data []Part `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListPartObject `json:"object" api:"required"`
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
func (r ListPart) RawJSON() string { return r.JSON.raw }
func (r *ListPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListPartObject string

const (
	ListPartObjectList ListPartObject = "list"
)

// A part in the account's catalog: a component used in production.
//
// Part-level data such as the SKU, description, category, pricing, and attributes
// lives on the underlying `item`.
type Part struct {
	// Part ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "part".
	Object PartObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Item        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Part) RawJSON() string { return r.JSON.raw }
func (r *Part) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type PartObject string

const (
	PartObjectPart PartObject = "part"
)

// Request to partially update a part.
type UpdatePartRequestParam struct {
	// New free-form description of the part.
	Description param.Opt[string] `json:"description,omitzero"`
	// New free-form notes about the part.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ID of the item category to move the part to.
	//
	// The move is the one Change Item Category makes: the category has to be a product
	// category and has to carry the properties of every attribute the part already
	// has, and the part's rate units switch to the category's base unit while their
	// numbers stay as they were. It is applied before the other fields in the request.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// New stock keeping unit code for the part.
	//
	// Must remain unique within the account; a conflict error is returned if another
	// item already uses it.
	SKU param.Opt[string] `json:"sku,omitzero"`
	paramObj
}

func (r UpdatePartRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdatePartRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdatePartRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPartNewParams struct {
	// Request to create a part.
	CreatePartRequest CreatePartRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogPartNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreatePartRequest)
}
func (r *CatalogPartNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPartNewParams]'s query parameters as `url.Values`.
func (r CatalogPartNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPartGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogPartGetParams]'s query parameters as `url.Values`.
func (r CatalogPartGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPartUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a part.
	UpdatePartRequest UpdatePartRequestParam
	paramObj
}

func (r CatalogPartUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdatePartRequest)
}
func (r *CatalogPartUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPartUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogPartUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogPartListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Only return parts created at or before this time.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Only return parts created at or after this time.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Only return parts carrying at least one of these attributes.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Only return parts belonging to any of these item categories.
	CategoryIDs []string `query:"category_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogPartListParams]'s query parameters as `url.Values`.
func (r CatalogPartListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
