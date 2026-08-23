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

// List and manage volume discounts.
//
// SaleVolumeDiscountService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleVolumeDiscountService] method instead.
type SaleVolumeDiscountService struct {
	options []option.RequestOption
}

// NewSaleVolumeDiscountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleVolumeDiscountService(opts ...option.RequestOption) (r SaleVolumeDiscountService) {
	r = SaleVolumeDiscountService{}
	r.options = opts
	return
}

// Creates a volume discount with its tiers and scoping associations.
//
// The discount name must be unique within the account; creating a discount with an
// existing name returns a conflict error.
//
// Each scoping list narrows the order lines the discount applies to, and an empty
// list places no restriction on that dimension. Because tier thresholds are
// compared against quantities converted into `unit_ids`, a discount created
// without any units never reaches a threshold above zero.
//
// This endpoint requires the permission: `discounts:create`.
func (r *SaleVolumeDiscountService) New(ctx context.Context, params SaleVolumeDiscountNewParams, opts ...option.RequestOption) (res *VolumeDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/volume-discounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a volume discount by ID.
//
// This endpoint requires the permissions: `discounts:read`, `customers:read`,
// `suppliers:read`.
func (r *SaleVolumeDiscountService) Get(ctx context.Context, id string, query SaleVolumeDiscountGetParams, opts ...option.RequestOption) (res *VolumeDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/volume-discounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a volume discount.
//
// The tier and association lists are only applied when their corresponding `has_*`
// flag is `true`, in which case they replace the existing set entirely. Tiers use
// upsert semantics: tiers with an `id` are updated, tiers without one are created,
// and existing tiers omitted from the list are deleted.
//
// The name must remain unique within the account; reusing another discount's name
// returns a conflict error. Order lines that have already been priced keep the
// unit price they were given; the revised discount applies to lines priced after
// the change.
//
// This endpoint requires the permission: `discounts:update`.
func (r *SaleVolumeDiscountService) Update(ctx context.Context, id string, params SaleVolumeDiscountUpdateParams, opts ...option.RequestOption) (res *VolumeDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/volume-discounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of volume discounts, newest first.
//
// The search term matches the discount name, the name of a customer group it is
// scoped to, or the name of a product line it is scoped to. Customer portal users
// see only discounts with no customer-group restriction plus those scoped to a
// group their own account belongs to.
//
// This endpoint requires the permissions: `discounts:read`, `customers:read`,
// `suppliers:read`.
func (r *SaleVolumeDiscountService) List(ctx context.Context, query SaleVolumeDiscountListParams, opts ...option.RequestOption) (res *ListVolumeDiscount, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/sales/volume-discounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a volume discount along with its tiers and scoping associations.
//
// Deletion is permanent; further requests against the deleted ID return an error.
//
// Order lines that have already been priced keep the unit price they were given;
// only lines priced after the deletion lose the discount.
//
// This endpoint requires the permission: `discounts:delete`.
func (r *SaleVolumeDiscountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SaleVolumeDiscountDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/volume-discounts/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a volume discount.
//
// The properties Name, Tiers are required.
type CreateVolumeDiscountRequestParam struct {
	// Display name of the volume discount.
	//
	// Must be unique within the account.
	Name string `json:"name" api:"required"`
	// Tiers for this volume discount.
	Tiers []CreateVolumeDiscountTierInputParam `json:"tiers,omitzero" api:"required"`
	// Attribute IDs to scope the discount to.
	//
	// When set, an item qualifies only if it has every listed attribute.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// Item category IDs to scope the discount to.
	//
	// When empty, all categories qualify.
	CategoryIDs []string `json:"category_ids,omitzero"`
	// Account group IDs to scope the discount to specific customer groups.
	//
	// When empty, all customers qualify. A discount scoped to a group the buyer
	// belongs to is preferred over an unscoped one when both could apply to the same
	// order line.
	CustomerGroupIDs []string `json:"customer_group_ids,omitzero"`
	// Product line IDs to scope the discount to.
	//
	// When empty, all product lines qualify.
	ProductLineIDs []string `json:"product_line_ids,omitzero"`
	// IDs of the units that ordered quantities are measured in when evaluating tier
	// thresholds.
	//
	// Quantities ordered in other units are converted into one of these before being
	// compared against a threshold. Leaving this empty makes the discount inert: the
	// quantity always evaluates to zero, so no threshold above zero is ever reached.
	UnitIDs []string `json:"unit_ids,omitzero"`
	paramObj
}

func (r CreateVolumeDiscountRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateVolumeDiscountRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateVolumeDiscountRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume discount tier to create.
//
// The properties DiscountPercentage, Name, Threshold are required.
type CreateVolumeDiscountTierInputParam struct {
	// Fraction of the price taken off once the threshold is met, as a decimal string.
	//
	// This is a multiplier, not a whole percent: `0.05` takes 5% off. When an order
	// meets several tiers of the same discount, their reductions compound.
	DiscountPercentage string `json:"discount_percentage" api:"required" format:"decimal"`
	// Display name of the tier.
	Name string `json:"name" api:"required"`
	// Minimum ordered quantity at which this tier's discount begins to apply, as a
	// decimal string.
	//
	// The quantity compared against the threshold is the total across every line on
	// the order that falls within the discount's scope, converted into one of the
	// discount's units.
	Threshold string `json:"threshold" api:"required" format:"decimal"`
	// ID of another tier that this tier follows.
	//
	// Tier IDs are assigned when the discount is created, so a tier created in this
	// same request cannot be referenced here. The link is stored with the tier but
	// does not affect pricing: every tier whose threshold is met applies, regardless
	// of any parent.
	ParentTierID param.Opt[string] `json:"parent_tier_id,omitzero"`
	paramObj
}

func (r CreateVolumeDiscountTierInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateVolumeDiscountTierInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateVolumeDiscountTierInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListVolumeDiscount struct {
	// Resources in this page.
	Data []VolumeDiscount `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListVolumeDiscountObject `json:"object" api:"required"`
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
func (r ListVolumeDiscount) RawJSON() string { return r.JSON.raw }
func (r *ListVolumeDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListVolumeDiscountObject string

const (
	ListVolumeDiscountObjectList ListVolumeDiscountObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListVolumeDiscountTier struct {
	// Resources in this page.
	Data []VolumeDiscountTier `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListVolumeDiscountTierObject `json:"object" api:"required"`
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
func (r ListVolumeDiscountTier) RawJSON() string { return r.JSON.raw }
func (r *ListVolumeDiscountTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListVolumeDiscountTierObject string

const (
	ListVolumeDiscountTierObjectList ListVolumeDiscountTierObject = "list"
)

// Request to partially update a volume discount.
//
// The properties HasAttributes, HasCategories, HasCustomerGroups, HasProductLines,
// HasTiers, HasUnits are required.
type UpdateVolumeDiscountRequestParam struct {
	// Whether to apply the `attribute_ids` field; when `false`, it is ignored.
	HasAttributes bool `json:"has_attributes" api:"required"`
	// Whether to apply the `category_ids` field; when `false`, it is ignored.
	HasCategories bool `json:"has_categories" api:"required"`
	// Whether to apply the `customer_group_ids` field; when `false`, it is ignored.
	HasCustomerGroups bool `json:"has_customer_groups" api:"required"`
	// Whether to apply the `product_line_ids` field; when `false`, it is ignored.
	HasProductLines bool `json:"has_product_lines" api:"required"`
	// Whether to apply the `tiers` field.
	//
	// When `true`, the discount's tiers are replaced with the contents of `tiers` (an
	// empty list deletes all tiers). When `false`, `tiers` is ignored.
	HasTiers bool `json:"has_tiers" api:"required"`
	// Whether to apply the `unit_ids` field; when `false`, it is ignored.
	HasUnits bool `json:"has_units" api:"required"`
	// Display name of the volume discount.
	//
	// Must be unique within the account.
	Name param.Opt[string] `json:"name,omitzero"`
	// Attribute IDs to set.
	//
	// Only applied when `has_attributes` is `true`, in which case they replace the
	// existing set entirely.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// Item category IDs to set.
	//
	// Only applied when `has_categories` is `true`, in which case they replace the
	// existing set entirely.
	CategoryIDs []string `json:"category_ids,omitzero"`
	// Account group IDs to set as customer groups.
	//
	// Only applied when `has_customer_groups` is `true`, in which case they replace
	// the existing set entirely.
	CustomerGroupIDs []string `json:"customer_group_ids,omitzero"`
	// Product line IDs to set.
	//
	// Only applied when `has_product_lines` is `true`, in which case they replace the
	// existing set entirely.
	ProductLineIDs []string `json:"product_line_ids,omitzero"`
	// The full set of tiers for this discount.
	//
	// Only applied when `has_tiers` is `true`. Tiers with an `id` are updated, tiers
	// without an `id` are created, and existing tiers not present in the list are
	// deleted.
	Tiers []UpdateVolumeDiscountTierInputParam `json:"tiers,omitzero"`
	// IDs of the units to set as acceptable units.
	//
	// Only applied when `has_units` is `true`, in which case they replace the existing
	// set entirely. Clearing every unit makes the discount inert, since ordered
	// quantity then always evaluates to zero.
	UnitIDs []string `json:"unit_ids,omitzero"`
	paramObj
}

func (r UpdateVolumeDiscountRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateVolumeDiscountRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateVolumeDiscountRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume discount tier to upsert.
//
// Each entry is written as a whole: send every value you want the tier to keep,
// since values left out are not carried over from the existing tier.
type UpdateVolumeDiscountTierInputParam struct {
	// ID of an existing tier to update.
	//
	// Omit to create a new tier.
	ID param.Opt[string] `json:"id,omitzero"`
	// Fraction of the price taken off once the threshold is met, as a decimal string.
	//
	// This is a multiplier, not a whole percent: `0.05` takes 5% off. When an order
	// meets several tiers of the same discount, their reductions compound.
	DiscountPercentage param.Opt[string] `json:"discount_percentage,omitzero" format:"decimal"`
	// Display name of the tier.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of another tier in this discount that this tier follows.
	//
	// The link is stored with the tier but does not affect pricing. Omitting it when
	// updating an existing tier clears the link.
	ParentTierID param.Opt[string] `json:"parent_tier_id,omitzero"`
	// Minimum ordered quantity at which this tier's discount begins to apply, as a
	// decimal string.
	//
	// The quantity compared against the threshold is the total across every line on
	// the order that falls within the discount's scope, converted into one of the
	// discount's units.
	Threshold param.Opt[string] `json:"threshold,omitzero" format:"decimal"`
	paramObj
}

func (r UpdateVolumeDiscountTierInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateVolumeDiscountTierInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateVolumeDiscountTierInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A quantity-based discount with tiered percentage rates.
//
// A volume discount reduces the price once the ordered quantity reaches a tier's
// threshold. The customer group associations scope which customers qualify, and
// the product line, category, and attribute associations scope which order lines
// qualify; an empty list on any of them means no restriction on that dimension.
// Acceptable units are not a scope: they are the units the ordered quantity is
// measured in, and a discount with none of them never reaches a threshold above
// zero.
//
// At most one volume discount is applied to a given order line: among the
// discounts whose scope the line matches and whose thresholds are met, those
// scoped to a customer group the buyer belongs to take precedence. An account
// price for the same line overrides the discounted price entirely.
type VolumeDiscount struct {
	// Volume discount ID.
	ID string `json:"id" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	AcceptableUnits ListUnit `json:"acceptable_units" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Attributes ListAttribute `json:"attributes" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Categories ListItemCategory `json:"categories" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	CustomerGroups ListAccountGroup `json:"customer_groups" api:"required"`
	// Display name of the volume discount.
	//
	// Must be unique within the account.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "volume_discount".
	Object VolumeDiscountObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ProductLines ListProductLine `json:"product_lines" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Tiers ListVolumeDiscountTier `json:"tiers" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AcceptableUnits respjson.Field
		Attributes      respjson.Field
		Categories      respjson.Field
		CreatedAt       respjson.Field
		CustomerGroups  respjson.Field
		Name            respjson.Field
		Object          respjson.Field
		ProductLines    respjson.Field
		Tiers           respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeDiscount) RawJSON() string { return r.JSON.raw }
func (r *VolumeDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type VolumeDiscountObject string

const (
	VolumeDiscountObjectVolumeDiscount VolumeDiscountObject = "volume_discount"
)

// A quantity threshold within a volume discount, and the reduction that applies at
// or above it.
type VolumeDiscountTier struct {
	// Volume discount tier ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Fraction of the price taken off once the threshold is met, as a decimal string.
	//
	// This is a multiplier, not a whole percent: `0.05` takes 5% off. When an order
	// meets several tiers of the same discount, their reductions compound: meeting a
	// `0.1` tier and a `0.2` tier multiplies the price by `0.9 × 0.8`, a 28% reduction
	// overall.
	DiscountPercentage string `json:"discount_percentage" api:"required" format:"decimal"`
	// Display name of the tier.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "volume_discount_tier".
	Object VolumeDiscountTierObject `json:"object" api:"required"`
	// Minimum ordered quantity at which this tier's discount begins to apply, as a
	// decimal string.
	//
	// The quantity compared against the threshold is the total across every line on
	// the order that falls within the discount's scope, converted into one of the
	// discount's acceptable units — not the quantity of a single line.
	Threshold string `json:"threshold" api:"required" format:"decimal"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		DiscountPercentage respjson.Field
		Name               respjson.Field
		Object             respjson.Field
		Threshold          respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeDiscountTier) RawJSON() string { return r.JSON.raw }
func (r *VolumeDiscountTier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type VolumeDiscountTierObject string

const (
	VolumeDiscountTierObjectVolumeDiscountTier VolumeDiscountTierObject = "volume_discount_tier"
)

type SaleVolumeDiscountDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleVolumeDiscountDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleVolumeDiscountDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleVolumeDiscountNewParams struct {
	// Request to create a volume discount.
	CreateVolumeDiscountRequest CreateVolumeDiscountRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer_groups", "product_lines", "categories",
	// "categories.properties", "attributes", "acceptable_units".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleVolumeDiscountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateVolumeDiscountRequest)
}
func (r *SaleVolumeDiscountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleVolumeDiscountNewParams]'s query parameters as
// `url.Values`.
func (r SaleVolumeDiscountNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleVolumeDiscountGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer_groups", "product_lines", "categories",
	// "categories.properties", "attributes", "acceptable_units".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleVolumeDiscountGetParams]'s query parameters as
// `url.Values`.
func (r SaleVolumeDiscountGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleVolumeDiscountUpdateParams struct {
	// Request to partially update a volume discount.
	UpdateVolumeDiscountRequest UpdateVolumeDiscountRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "customer_groups", "product_lines", "categories",
	// "categories.properties", "attributes", "acceptable_units".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleVolumeDiscountUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateVolumeDiscountRequest)
}
func (r *SaleVolumeDiscountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleVolumeDiscountUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleVolumeDiscountUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleVolumeDiscountListParams struct {
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
	// Any of "customer_groups", "product_lines", "categories",
	// "categories.properties", "attributes", "acceptable_units".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SaleVolumeDiscountListParams]'s query parameters as
// `url.Values`.
func (r SaleVolumeDiscountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
