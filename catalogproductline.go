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

// List and manage product lines.
//
// CatalogProductLineService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogProductLineService] method instead.
type CatalogProductLineService struct {
	options []option.RequestOption
	// List and manage product lines.
	Actions CatalogProductLineActionService
}

// NewCatalogProductLineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogProductLineService(opts ...option.RequestOption) (r CatalogProductLineService) {
	r = CatalogProductLineService{}
	r.options = opts
	r.Actions = NewCatalogProductLineActionService(opts...)
	return
}

// Creates a product line owned by your account.
//
// The new line starts with no products; assign products to it by setting their
// product line. Customers and account groups can only be granted access to lines
// your account owns, so this is the starting point for scoping a customer's
// catalog.
//
// This endpoint requires the permission: `product_lines:create`.
func (r *CatalogProductLineService) New(ctx context.Context, params CatalogProductLineNewParams, opts ...option.RequestOption) (res *ProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/product-lines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single product line by ID.
//
// Both the product lines your account owns and the shared system lines can be
// retrieved.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `product_lines:read`.
func (r *CatalogProductLineService) Get(ctx context.Context, id string, query CatalogProductLineGetParams, opts ...option.RequestOption) (res *ProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/product-lines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a product line your account owns.
//
// Only the provided fields are changed. The reserved `shipping`, `service`,
// `credit`, and `tax` lines cannot be updated, and neither can the shared system
// lines, which belong to no single account.
//
// This endpoint requires the permission: `product_lines:update`.
func (r *CatalogProductLineService) Update(ctx context.Context, id string, params CatalogProductLineUpdateParams, opts ...option.RequestOption) (res *ProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/product-lines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of product lines, newest first.
//
// Covers both the product lines your account owns and the shared system lines. The
// `q` search term is matched against the product line name.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `product_lines:read`.
func (r *CatalogProductLineService) List(ctx context.Context, query CatalogProductLineListParams, opts ...option.RequestOption) (res *ListProductLine, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/product-lines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a product line your account owns.
//
// The reserved `shipping`, `service`, `credit`, and `tax` lines cannot be deleted,
// and neither can the shared system lines, which belong to no single account.
// Deleting a line that was already deleted returns an already-deleted error rather
// than succeeding silently.
//
// This endpoint requires the permission: `product_lines:delete`.
func (r *CatalogProductLineService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CatalogProductLineDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/product-lines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a product line.
//
// The properties CommissionPolicy, FreightPolicy, Name, UnitGroupID are required.
type CreateProductLineRequestParam struct {
	// Default commission policy for products in this product line.
	//
	//   - `commission_exempt`: no commission applies to these products.
	//   - `commission_applied`: commission applies to these products, unless overridden
	//     elsewhere.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy CreateProductLineRequestCommissionPolicy `json:"commission_policy,omitzero" api:"required"`
	// Default freight policy for products in this product line.
	//
	//   - `free_freight`: these products do not incur a freight charge.
	//   - `billed_freight`: freight is billed for these products, unless overridden
	//     elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy CreateProductLineRequestFreightPolicy `json:"freight_policy,omitzero" api:"required"`
	// Display name of the product line.
	//
	// Must be unique among the product lines visible to your account, including the
	// shared system lines; a duplicate name returns a conflict error.
	Name string `json:"name" api:"required"`
	// ID of the unit group to associate with this product line.
	//
	// The unit group determines the set of units available to products in this product
	// line. It must be a unit group your account owns or one of the shared system unit
	// groups.
	UnitGroupID string `json:"unit_group_id" api:"required"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	DefaultLot QuantityInputParam `json:"default_lot,omitzero"`
	// How products in this line are produced when they do not say for themselves.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy CreateProductLineRequestFulfillmentPolicy `json:"fulfillment_policy,omitzero"`
	paramObj
}

func (r CreateProductLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProductLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProductLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default commission policy for products in this product line.
//
//   - `commission_exempt`: no commission applies to these products.
//   - `commission_applied`: commission applies to these products, unless overridden
//     elsewhere.
type CreateProductLineRequestCommissionPolicy string

const (
	CreateProductLineRequestCommissionPolicyCommissionApplied CreateProductLineRequestCommissionPolicy = "commission_applied"
	CreateProductLineRequestCommissionPolicyCommissionExempt  CreateProductLineRequestCommissionPolicy = "commission_exempt"
)

// Default freight policy for products in this product line.
//
//   - `free_freight`: these products do not incur a freight charge.
//   - `billed_freight`: freight is billed for these products, unless overridden
//     elsewhere.
type CreateProductLineRequestFreightPolicy string

const (
	CreateProductLineRequestFreightPolicyFreeFreight   CreateProductLineRequestFreightPolicy = "free_freight"
	CreateProductLineRequestFreightPolicyBilledFreight CreateProductLineRequestFreightPolicy = "billed_freight"
)

// How products in this line are produced when they do not say for themselves.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
type CreateProductLineRequestFulfillmentPolicy string

const (
	CreateProductLineRequestFulfillmentPolicyMakeToStock CreateProductLineRequestFulfillmentPolicy = "make_to_stock"
	CreateProductLineRequestFulfillmentPolicyMakeToOrder CreateProductLineRequestFulfillmentPolicy = "make_to_order"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductLine struct {
	// Resources in this page.
	Data []ProductLine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductLineObject `json:"object" api:"required"`
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
func (r ListProductLine) RawJSON() string { return r.JSON.raw }
func (r *ListProductLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductLineObject string

const (
	ListProductLineObjectList ListProductLineObject = "list"
)

// A named grouping of related products in your catalog.
//
// A product line carries the default commission and freight policies for the
// products assigned to it, along with the unit group that determines how those
// products are measured. Product lines are also the unit that catalog access is
// granted over, for both customers and account groups.
type ProductLine struct {
	// Product line ID.
	ID string `json:"id" api:"required"`
	// Default commission policy for products in this product line.
	//
	//   - `commission_exempt`: no commission applies to these products.
	//   - `commission_applied`: commission applies to these products, unless overridden
	//     elsewhere.
	//
	// Null to customer and supplier portal users, like the rest of your commission
	// settings.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy ProductLineCommissionPolicy `json:"commission_policy" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	DefaultLot Quantity `json:"default_lot" api:"required"`
	// Free-form description of the product line.
	Description string `json:"description" api:"required"`
	// Default freight policy for products in this product line.
	//
	//   - `free_freight`: these products do not incur a freight charge.
	//   - `billed_freight`: freight is billed for these products, unless overridden
	//     elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy ProductLineFreightPolicy `json:"freight_policy" api:"required"`
	// How products in this line are produced when they do not say for themselves.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Null falls through to the account default. Always null to customer and supplier
	// portal users, like the rest of your production planning.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy ProductLineFulfillmentPolicy `json:"fulfillment_policy" api:"required"`
	// Display name of the product line.
	//
	// Unique among the product lines visible to your account, which includes the
	// shared system lines.
	Name string `json:"name" api:"required"`
	// Free-form notes about the product line.
	//
	// Null to customer and supplier portal users: they are your own team's notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "product_line".
	Object ProductLineObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// A named collection of units that share one dimension, defining which units a
	// product can be ordered in.
	//
	// Each associated unit carries its own discount and customer portal visibility,
	// applied when an order line is priced in that unit. A product takes its unit
	// group from its product line, falling back to its item category.
	UnitGroup UnitGroup `json:"unit_group" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CommissionPolicy  respjson.Field
		CreatedAt         respjson.Field
		DefaultLot        respjson.Field
		Description       respjson.Field
		FreightPolicy     respjson.Field
		FulfillmentPolicy respjson.Field
		Name              respjson.Field
		Notes             respjson.Field
		Object            respjson.Field
		Owner             respjson.Field
		UnitGroup         respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductLine) RawJSON() string { return r.JSON.raw }
func (r *ProductLine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default commission policy for products in this product line.
//
//   - `commission_exempt`: no commission applies to these products.
//   - `commission_applied`: commission applies to these products, unless overridden
//     elsewhere.
//
// Null to customer and supplier portal users, like the rest of your commission
// settings.
type ProductLineCommissionPolicy string

const (
	ProductLineCommissionPolicyCommissionApplied ProductLineCommissionPolicy = "commission_applied"
	ProductLineCommissionPolicyCommissionExempt  ProductLineCommissionPolicy = "commission_exempt"
)

// Default freight policy for products in this product line.
//
//   - `free_freight`: these products do not incur a freight charge.
//   - `billed_freight`: freight is billed for these products, unless overridden
//     elsewhere.
type ProductLineFreightPolicy string

const (
	ProductLineFreightPolicyFreeFreight   ProductLineFreightPolicy = "free_freight"
	ProductLineFreightPolicyBilledFreight ProductLineFreightPolicy = "billed_freight"
)

// How products in this line are produced when they do not say for themselves.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
//
// Null falls through to the account default. Always null to customer and supplier
// portal users, like the rest of your production planning.
type ProductLineFulfillmentPolicy string

const (
	ProductLineFulfillmentPolicyMakeToStock ProductLineFulfillmentPolicy = "make_to_stock"
	ProductLineFulfillmentPolicyMakeToOrder ProductLineFulfillmentPolicy = "make_to_order"
)

// Resource type identifier.
type ProductLineObject string

const (
	ProductLineObjectProductLine ProductLineObject = "product_line"
)

// Request to partially update a product line.
type UpdateProductLineRequestParam struct {
	// Display name of the product line.
	//
	// Must be unique among the product lines visible to your account, including the
	// shared system lines; a duplicate name returns a conflict error.
	Name param.Opt[string] `json:"name,omitzero"`
	// ID of the unit group to associate with this product line.
	//
	// The unit group determines the set of units available to products in this product
	// line. It must be a unit group your account owns or one of the shared system unit
	// groups. A lot already stored on the line is not rechecked when the group
	// changes, so send `default_lot` alongside to keep the two consistent.
	UnitGroupID param.Opt[string] `json:"unit_group_id,omitzero"`
	// How products in this line are produced when they do not say for themselves.
	//
	//   - `make_to_stock`: built to the forecast, holding a safety stock against its
	//     variability.
	//   - `make_to_order`: built only against orders already on the book, holding no
	//     buffer.
	//
	// Clearing it returns the line's products to the account default.
	//
	// Any of "make_to_stock", "make_to_order".
	FulfillmentPolicy UpdateProductLineRequestFulfillmentPolicy `json:"fulfillment_policy,omitzero"`
	// Default commission policy for products in this product line.
	//
	//   - `commission_exempt`: no commission applies to these products.
	//   - `commission_applied`: commission applies to these products, unless overridden
	//     elsewhere.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpdateProductLineRequestCommissionPolicy `json:"commission_policy,omitzero"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	DefaultLot QuantityInputParam `json:"default_lot,omitzero"`
	// Default freight policy for products in this product line.
	//
	//   - `free_freight`: these products do not incur a freight charge.
	//   - `billed_freight`: freight is billed for these products, unless overridden
	//     elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpdateProductLineRequestFreightPolicy `json:"freight_policy,omitzero"`
	paramObj
}

func (r UpdateProductLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default commission policy for products in this product line.
//
//   - `commission_exempt`: no commission applies to these products.
//   - `commission_applied`: commission applies to these products, unless overridden
//     elsewhere.
type UpdateProductLineRequestCommissionPolicy string

const (
	UpdateProductLineRequestCommissionPolicyCommissionApplied UpdateProductLineRequestCommissionPolicy = "commission_applied"
	UpdateProductLineRequestCommissionPolicyCommissionExempt  UpdateProductLineRequestCommissionPolicy = "commission_exempt"
)

// Default freight policy for products in this product line.
//
//   - `free_freight`: these products do not incur a freight charge.
//   - `billed_freight`: freight is billed for these products, unless overridden
//     elsewhere.
type UpdateProductLineRequestFreightPolicy string

const (
	UpdateProductLineRequestFreightPolicyFreeFreight   UpdateProductLineRequestFreightPolicy = "free_freight"
	UpdateProductLineRequestFreightPolicyBilledFreight UpdateProductLineRequestFreightPolicy = "billed_freight"
)

// How products in this line are produced when they do not say for themselves.
//
//   - `make_to_stock`: built to the forecast, holding a safety stock against its
//     variability.
//   - `make_to_order`: built only against orders already on the book, holding no
//     buffer.
//
// Clearing it returns the line's products to the account default.
type UpdateProductLineRequestFulfillmentPolicy string

const (
	UpdateProductLineRequestFulfillmentPolicyMakeToStock UpdateProductLineRequestFulfillmentPolicy = "make_to_stock"
	UpdateProductLineRequestFulfillmentPolicyMakeToOrder UpdateProductLineRequestFulfillmentPolicy = "make_to_order"
)

type CatalogProductLineDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogProductLineDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogProductLineDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogProductLineNewParams struct {
	// Request to create a product line.
	CreateProductLineRequest CreateProductLineRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group", "default_lot",
	// "default_lot.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogProductLineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProductLineRequest)
}
func (r *CatalogProductLineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductLineNewParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductLineGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group", "default_lot",
	// "default_lot.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductLineGetParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductLineUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "owner", "owner.account", "unit_group", "default_lot",
	// "default_lot.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a product line.
	UpdateProductLineRequest UpdateProductLineRequestParam
	paramObj
}

func (r CatalogProductLineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductLineRequest)
}
func (r *CatalogProductLineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductLineUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductLineListParams struct {
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
	// Any of "owner", "owner.account", "unit_group", "default_lot",
	// "default_lot.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductLineListParams]'s query parameters as
// `url.Values`.
func (r CatalogProductLineListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
