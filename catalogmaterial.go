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

// List and manage materials.
//
// CatalogMaterialService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogMaterialService] method instead.
type CatalogMaterialService struct {
	options []option.RequestOption
	// List and manage materials.
	Actions CatalogMaterialActionService
}

// NewCatalogMaterialService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogMaterialService(opts ...option.RequestOption) (r CatalogMaterialService) {
	r = CatalogMaterialService{}
	r.options = opts
	r.Actions = NewCatalogMaterialActionService(opts...)
	return
}

// Creates a material together with the catalog item that carries its SKU,
// description, category, pricing, and attributes.
//
// Inventory tracking for the new material starts at a zero on-hand quantity in the
// category's base unit. The item's consumption rate (`burn_rate`) also starts at
// zero and cannot be supplied here — it is derived from recorded consumption as
// production happens.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `materials:create`.
func (r *CatalogMaterialService) New(ctx context.Context, params CatalogMaterialNewParams, opts ...option.RequestOption) (res *Material, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/materials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a material by ID.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `materials:read`.
func (r *CatalogMaterialService) Get(ctx context.Context, id string, query CatalogMaterialGetParams, opts ...option.RequestOption) (res *Material, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/materials/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a material.
//
// Fields not provided retain their current values. Only the cost side of pricing
// can be changed here; the selling price set at creation is not editable through
// this endpoint.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `materials:update`.
func (r *CatalogMaterialService) Update(ctx context.Context, id string, params CatalogMaterialUpdateParams, opts ...option.RequestOption) (res *Material, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/materials/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of materials, newest first.
//
// `q` matches against SKU and description, with closer SKU matches ranked first.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `materials:read`.
func (r *CatalogMaterialService) List(ctx context.Context, query CatalogMaterialListParams, opts ...option.RequestOption) (res *ListMaterial, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/materials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a material.
//
// This is a soft delete: the material and the catalog item behind it stop being
// returned by other endpoints, but the records are retained. The response is the
// material as it stood immediately before deletion, and deleting an
// already-deleted material returns an error.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `materials:delete`.
func (r *CatalogMaterialService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *Material, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/materials/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a material.
//
// The properties CategoryID, SKU are required.
type CreateMaterialRequestParam struct {
	// ID of the item category to place the material in.
	//
	// The category's unit group determines the base unit used for the material's rates
	// (`unit_value`, `unit_cost`, `burn_rate`).
	CategoryID string `json:"category_id" api:"required"`
	// Stock keeping unit code for the material.
	//
	// Must be unique within the account; creating a material with a SKU already used
	// by another item fails with a conflict error.
	SKU string `json:"sku" api:"required"`
	// Free-form description of the material.
	Description param.Opt[string] `json:"description,omitzero"`
	// Free-form notes about the material.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// IDs of existing attributes to link to the material at creation time.
	//
	// Each attribute's property must be one the material's category carries; an
	// attribute from any other property fails the whole request.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// A quantity, given as a decimal value and the unit it is measured in.
	LeadTime QuantityInputRequestParam `json:"lead_time,omitzero"`
	// A quantity, given as a decimal value and the unit it is measured in.
	OrderPoint QuantityInputRequestParam `json:"order_point,omitzero"`
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

func (r CreateMaterialRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMaterialRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMaterialRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMaterial struct {
	// Resources in this page.
	Data []Material `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMaterialObject `json:"object" api:"required"`
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
func (r ListMaterial) RawJSON() string { return r.JSON.raw }
func (r *ListMaterial) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMaterialObject string

const (
	ListMaterialObjectList ListMaterialObject = "list"
)

// A material in the account's catalog: a raw material or component consumed in
// production.
//
// Material-level data such as the SKU, description, category, pricing, and
// attributes lives on the underlying `item`; the material record adds the
// reordering fields `order_point` and `lead_time`.
type Material struct {
	// Material ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	LeadTime Quantity `json:"lead_time" api:"required"`
	// Resource type identifier.
	//
	// Any of "material".
	Object MaterialObject `json:"object" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	OrderPoint Quantity `json:"order_point" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Item        respjson.Field
		LeadTime    respjson.Field
		Object      respjson.Field
		OrderPoint  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Material) RawJSON() string { return r.JSON.raw }
func (r *Material) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MaterialObject string

const (
	MaterialObjectMaterial MaterialObject = "material"
)

// A measured amount: a numeric value together with the unit it is expressed in.
//
// Quantities are shared building blocks rather than standalone records — other
// resources point at them to report stock levels, ordered and packed amounts,
// money, weights, and durations.
type Quantity struct {
	// Quantity ID.
	ID string `json:"id" api:"required"`
	// Formatted value with unit abbreviation (e.g. "$1,234.56" or "100 kg").
	DisplayValue string `json:"display_value" api:"required"`
	// Resource type identifier.
	//
	// Any of "quantity".
	Object QuantityObject `json:"object" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// Raw decimal value of the quantity, as a string to preserve precision.
	//
	// This is the unformatted machine value; see `display_value` for the
	// human-readable rendering with unit and thousands separators.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayValue respjson.Field
		Object       respjson.Field
		Unit         respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quantity) RawJSON() string { return r.JSON.raw }
func (r *Quantity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type QuantityObject string

const (
	QuantityObjectQuantity QuantityObject = "quantity"
)

// A quantity, given as a decimal value and the unit it is measured in.
//
// The properties UnitID, Value are required.
type QuantityInputRequestParam struct {
	// ID of the unit the value is expressed in.
	UnitID string `json:"unit_id" api:"required"`
	// Decimal value of the quantity.
	Value string `json:"value" api:"required" format:"decimal"`
	paramObj
}

func (r QuantityInputRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow QuantityInputRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QuantityInputRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A value expressed as a ratio of two units, supplied on create and update
// requests.
//
// A unit price, for example, has a currency as its numerator unit and the unit the
// product is bought or sold by as its denominator.
//
// The properties DenominatorUnitID, NumeratorUnitID, Value are required.
type RateInputParam struct {
	// ID of the unit for the rate's denominator (the per-unit basis).
	DenominatorUnitID string `json:"denominator_unit_id" api:"required"`
	// ID of the unit for the rate's numerator (e.g. the currency of a price).
	NumeratorUnitID string `json:"numerator_unit_id" api:"required"`
	// Decimal value of the rate, expressed as the amount of the numerator unit per one
	// denominator unit.
	Value string `json:"value" api:"required" format:"decimal"`
	paramObj
}

func (r RateInputParam) MarshalJSON() (data []byte, err error) {
	type shadow RateInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RateInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to update a material.
type UpdateMaterialRequestParam struct {
	// ID of the item category to move the material to.
	//
	// The move is the one Change Item Category makes: the category has to be a
	// material category and has to carry the properties of every attribute the
	// material already has, and the material's rate and order-point units switch to
	// the category's base unit while their numbers stay as they were. It is applied
	// before the other fields in the request, so an `order_point` or `unit_cost` sent
	// alongside is written after it.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// New description for the material.
	Description param.Opt[string] `json:"description,omitzero"`
	// New notes for the material.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// New stock keeping unit code for the material.
	//
	// Must remain unique within the account; a conflict error is returned if another
	// item already uses it.
	SKU param.Opt[string] `json:"sku,omitzero"`
	// A quantity, given as a decimal value and the unit it is measured in.
	LeadTime QuantityInputRequestParam `json:"lead_time,omitzero"`
	// A quantity, given as a decimal value and the unit it is measured in.
	OrderPoint QuantityInputRequestParam `json:"order_point,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitCost RateInputParam `json:"unit_cost,omitzero"`
	paramObj
}

func (r UpdateMaterialRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMaterialRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMaterialRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogMaterialNewParams struct {
	// Request to create a material.
	CreateMaterialRequest CreateMaterialRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogMaterialNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMaterialRequest)
}
func (r *CatalogMaterialNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogMaterialNewParams]'s query parameters as
// `url.Values`.
func (r CatalogMaterialNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogMaterialGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogMaterialGetParams]'s query parameters as
// `url.Values`.
func (r CatalogMaterialGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogMaterialUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "item", "item.category", "item.category.properties",
	// "item.category.unit_group", "item.unit_value", "item.unit_cost",
	// "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a material.
	UpdateMaterialRequest UpdateMaterialRequestParam
	paramObj
}

func (r CatalogMaterialUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMaterialRequest)
}
func (r *CatalogMaterialUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogMaterialUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogMaterialUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogMaterialListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to materials created on or before this date.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter to materials created on or after this date.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Filter to materials carrying any of these attributes.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Filter to materials in any of these categories.
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

// URLQuery serializes [CatalogMaterialListParams]'s query parameters as
// `url.Values`.
func (r CatalogMaterialListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
