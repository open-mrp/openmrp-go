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

// List and manage units.
//
// CatalogUnitService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitService] method instead.
type CatalogUnitService struct {
	options []option.RequestOption
	// List and manage units.
	Actions CatalogUnitActionService
}

// NewCatalogUnitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogUnitService(opts ...option.RequestOption) (r CatalogUnitService) {
	r = CatalogUnitService{}
	r.options = opts
	r.Actions = NewCatalogUnitActionService(opts...)
	return
}

// Creates a unit of measurement owned by your account, in addition to the system
// units the platform already provides.
//
// The name and abbreviation must each be unique within the account. A unit created
// here is never a base unit, so its conversion ratio is interpreted relative to
// the base unit of the chosen dimension.
//
// This endpoint requires the permission: `units:create`.
func (r *CatalogUnitService) New(ctx context.Context, params CatalogUnitNewParams, opts ...option.RequestOption) (res *Unit, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/units"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a unit by ID, including both account-owned and global system units.
//
// This endpoint requires the permission: `units:read`.
func (r *CatalogUnitService) Get(ctx context.Context, id string, query CatalogUnitGetParams, opts ...option.RequestOption) (res *Unit, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/units/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a unit owned by your account.
//
// System units cannot be modified, and a unit's dimension is fixed once it is
// created.
//
// This endpoint requires the permission: `units:update`.
func (r *CatalogUnitService) Update(ctx context.Context, id string, params CatalogUnitUpdateParams, opts ...option.RequestOption) (res *Unit, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/units/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of units for the current account, including both
// account-owned and global system units.
//
// This endpoint requires the permission: `units:read`.
func (r *CatalogUnitService) List(ctx context.Context, query CatalogUnitListParams, opts ...option.RequestOption) (res *ListUnit, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/units"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a unit owned by your account.
//
// The unit is also removed from every unit group it belongs to. System units,
// which are shared across all accounts, cannot be deleted, and neither can a unit
// that is a unit group's base unit — change the group's base unit or delete the
// group first — or one that any quantity or price is recorded in.
//
// This endpoint requires the permission: `units:delete`.
func (r *CatalogUnitService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogUnitDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/units/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a unit.
//
// The properties Abbreviation, Name, OffsetDenominator, OffsetNumerator,
// RatioDenominator, RatioNumerator, Type are required.
type CreateUnitRequestParam struct {
	// Short abbreviation for the unit (e.g. "g").
	//
	// Must be unique within the account.
	Abbreviation string `json:"abbreviation" api:"required"`
	// Display name of the unit (e.g. "Gram").
	//
	// Must be unique within the account.
	Name string `json:"name" api:"required"`
	// Denominator of the conversion offset.
	//
	// Must not be zero, so send `1` when the unit has no offset.
	OffsetDenominator string `json:"offset_denominator" api:"required" format:"decimal"`
	// Numerator of the conversion offset, applied after the ratio for scales that do
	// not share a zero point, such as temperature.
	//
	// Send `0` for units that convert by ratio alone.
	OffsetNumerator string `json:"offset_numerator" api:"required" format:"decimal"`
	// Denominator of the ratio that converts a quantity in this unit into the
	// dimension's base unit.
	//
	// Must not be zero.
	RatioDenominator string `json:"ratio_denominator" api:"required" format:"decimal"`
	// Numerator of the ratio that converts a quantity in this unit into the
	// dimension's base unit.
	//
	// A quantity is converted with
	// `value × (ratio_numerator / ratio_denominator) + (offset_numerator / offset_denominator)`,
	// so a kilogram in a gram-based dimension has a numerator of `1000` and a
	// denominator of `1`.
	RatioNumerator string `json:"ratio_numerator" api:"required" format:"decimal"`
	// The dimension this unit measures, such as mass, volume, or currency.
	//
	// Units can only be converted to other units of the same dimension, and the
	// dimension cannot be changed after the unit is created.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CreateUnitRequestType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r CreateUnitRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateUnitRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateUnitRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The dimension this unit measures, such as mass, volume, or currency.
//
// Units can only be converted to other units of the same dimension, and the
// dimension cannot be changed after the unit is created.
type CreateUnitRequestType string

const (
	CreateUnitRequestTypeCurrency    CreateUnitRequestType = "currency"
	CreateUnitRequestTypeQuantity    CreateUnitRequestType = "quantity"
	CreateUnitRequestTypeTime        CreateUnitRequestType = "time"
	CreateUnitRequestTypeMass        CreateUnitRequestType = "mass"
	CreateUnitRequestTypeVolume      CreateUnitRequestType = "volume"
	CreateUnitRequestTypeLength      CreateUnitRequestType = "length"
	CreateUnitRequestTypeTemperature CreateUnitRequestType = "temperature"
	CreateUnitRequestTypeArea        CreateUnitRequestType = "area"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListUnit struct {
	// Resources in this page.
	Data []Unit `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListUnitObject `json:"object" api:"required"`
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
func (r ListUnit) RawJSON() string { return r.JSON.raw }
func (r *ListUnit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListUnitObject string

const (
	ListUnitObjectList ListUnitObject = "list"
)

// Unit of measurement used for conversions and product quantities.
type Unit struct {
	// Unit ID.
	ID string `json:"id" api:"required"`
	// Short abbreviation for the unit (e.g. "g", "kg").
	Abbreviation string `json:"abbreviation" api:"required"`
	// When this unit was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this is the base unit for its dimension.
	//
	// Every other unit's conversion ratio is expressed relative to the base unit. Base
	// units are platform-defined; units created through the API are never base units.
	IsBaseUnit bool `json:"is_base_unit" api:"required"`
	// Display name of the unit (e.g. "Gram", "Kilogram").
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "unit".
	Object UnitObject `json:"object" api:"required"`
	// Denominator of the conversion offset applied after the ratio.
	//
	// Never zero; a unit with no offset carries a numerator of `0` over a denominator
	// of `1`.
	OffsetDenominator string `json:"offset_denominator" api:"required" format:"decimal"`
	// Numerator of the conversion offset, applied after the ratio for scales that do
	// not share a zero point, such as temperature.
	//
	// Zero for units that convert by ratio alone.
	OffsetNumerator string `json:"offset_numerator" api:"required" format:"decimal"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Denominator of the ratio that converts a quantity in this unit into the
	// dimension's base unit.
	//
	// Cannot be zero.
	RatioDenominator string `json:"ratio_denominator" api:"required" format:"decimal"`
	// Numerator of the ratio that converts a quantity in this unit into the
	// dimension's base unit.
	//
	// A quantity is converted with
	// `value × (ratio_numerator / ratio_denominator) + (offset_numerator / offset_denominator)`,
	// so a kilogram in a gram-based dimension has a numerator of `1000` and a
	// denominator of `1`.
	RatioNumerator string `json:"ratio_numerator" api:"required" format:"decimal"`
	// The dimension this unit measures, such as mass, volume, or currency.
	//
	// A unit can only be converted to another unit of the same dimension. The
	// `quantity` dimension is for discrete countable items rather than a physical
	// measure.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type UnitType `json:"type" api:"required"`
	// When this unit was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Abbreviation      respjson.Field
		CreatedAt         respjson.Field
		IsBaseUnit        respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		OffsetDenominator respjson.Field
		OffsetNumerator   respjson.Field
		Owner             respjson.Field
		RatioDenominator  respjson.Field
		RatioNumerator    respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Unit) RawJSON() string { return r.JSON.raw }
func (r *Unit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type UnitObject string

const (
	UnitObjectUnit UnitObject = "unit"
)

// The dimension this unit measures, such as mass, volume, or currency.
//
// A unit can only be converted to another unit of the same dimension. The
// `quantity` dimension is for discrete countable items rather than a physical
// measure.
type UnitType string

const (
	UnitTypeCurrency    UnitType = "currency"
	UnitTypeQuantity    UnitType = "quantity"
	UnitTypeTime        UnitType = "time"
	UnitTypeMass        UnitType = "mass"
	UnitTypeVolume      UnitType = "volume"
	UnitTypeLength      UnitType = "length"
	UnitTypeTemperature UnitType = "temperature"
	UnitTypeArea        UnitType = "area"
)

// Request to partially update a unit.
type UpdateUnitRequestParam struct {
	// Short abbreviation for the unit.
	//
	// Must be unique within the account.
	Abbreviation param.Opt[string] `json:"abbreviation,omitzero"`
	// Display name of the unit.
	//
	// Must be unique within the account.
	Name param.Opt[string] `json:"name,omitzero"`
	// Denominator of the conversion offset.
	//
	// Must not be zero.
	OffsetDenominator param.Opt[string] `json:"offset_denominator,omitzero" format:"decimal"`
	// Numerator of the conversion offset, applied after the ratio for scales that do
	// not share a zero point, such as temperature.
	OffsetNumerator param.Opt[string] `json:"offset_numerator,omitzero" format:"decimal"`
	// Denominator of the ratio that converts a quantity in this unit into the
	// dimension's base unit.
	//
	// Must not be zero.
	RatioDenominator param.Opt[string] `json:"ratio_denominator,omitzero" format:"decimal"`
	// Numerator of the ratio that converts a quantity in this unit into the
	// dimension's base unit.
	//
	// A quantity is converted with
	// `value × (ratio_numerator / ratio_denominator) + (offset_numerator / offset_denominator)`.
	RatioNumerator param.Opt[string] `json:"ratio_numerator,omitzero" format:"decimal"`
	paramObj
}

func (r UpdateUnitRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateUnitRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateUnitRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogUnitDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogUnitDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogUnitNewParams struct {
	// Request to create a unit.
	CreateUnitRequest CreateUnitRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogUnitNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateUnitRequest)
}
func (r *CatalogUnitNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitNewParams]'s query parameters as `url.Values`.
func (r CatalogUnitNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitGetParams]'s query parameters as `url.Values`.
func (r CatalogUnitGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a unit.
	UpdateUnitRequest UpdateUnitRequestParam
	paramObj
}

func (r CatalogUnitUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateUnitRequest)
}
func (r *CatalogUnitUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogUnitListParams struct {
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
	// Filter by unit dimension.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type CatalogUnitListParamsType `query:"type,omitzero" json:"-"`
	// Return only units that belong to at least one of the given unit groups.
	UnitGroupIDs []string `query:"unit_group_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogUnitListParams]'s query parameters as `url.Values`.
func (r CatalogUnitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by unit dimension.
type CatalogUnitListParamsType string

const (
	CatalogUnitListParamsTypeCurrency    CatalogUnitListParamsType = "currency"
	CatalogUnitListParamsTypeQuantity    CatalogUnitListParamsType = "quantity"
	CatalogUnitListParamsTypeTime        CatalogUnitListParamsType = "time"
	CatalogUnitListParamsTypeMass        CatalogUnitListParamsType = "mass"
	CatalogUnitListParamsTypeVolume      CatalogUnitListParamsType = "volume"
	CatalogUnitListParamsTypeLength      CatalogUnitListParamsType = "length"
	CatalogUnitListParamsTypeTemperature CatalogUnitListParamsType = "temperature"
	CatalogUnitListParamsTypeArea        CatalogUnitListParamsType = "area"
)
