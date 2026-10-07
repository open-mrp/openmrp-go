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

// List and manage products.
//
// CatalogProductService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogProductService] method instead.
type CatalogProductService struct {
	options []option.RequestOption
	// List and manage products.
	Actions CatalogProductActionService
}

// NewCatalogProductService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogProductService(opts ...option.RequestOption) (r CatalogProductService) {
	r = CatalogProductService{}
	r.options = opts
	r.Actions = NewCatalogProductActionService(opts...)
	return
}

// Creates a product and its backing inventory item.
//
// The new item starts with zero on-hand inventory, and its pricing defaults to
// zero rates in the category's base unit unless `unit_price` or `unit_cost` is
// provided.
//
// Only products of type `sale` appear in the product list and export; products
// created with any other type are still usable on orders and invoices but must be
// retrieved by ID.
//
// This endpoint requires the permission: `items:create`.
func (r *CatalogProductService) New(ctx context.Context, params CatalogProductNewParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a product by ID.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `items:read`.
func (r *CatalogProductService) Get(ctx context.Context, id string, query CatalogProductGetParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a product.
//
// `sku`, `description`, `notes`, and `unit_price` all live on the product's
// backing item and are written there, so the change is visible on the item as
// well. The product line is reassigned through its own endpoint, and the product
// type cannot be changed after creation.
//
// This endpoint requires the permission: `items:update`.
func (r *CatalogProductService) Update(ctx context.Context, id string, params CatalogProductUpdateParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of products for the target account, newest first.
//
// Only products of type `sale` are listed — service, shipping, credit, return, and
// tax products are excluded and must be retrieved by ID. A request made by a
// customer-portal buyer always returns portal-visible products only, and its
// `customer_ids` filter is replaced with the buyer's own account, so the results
// reflect what that account is entitled to buy.
//
// The `q` search term is matched against the SKU and description of each product's
// item; when it is supplied, products whose SKU matches are returned ahead of the
// rest.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `items:read`.
func (r *CatalogProductService) List(ctx context.Context, query CatalogProductListParams, opts ...option.RequestOption) (res *ListProduct, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Soft-deletes a product and returns it as it stood at deletion.
//
// Deletion marks the product's backing item as deleted, so the item and its
// inventory drop out of catalog and inventory listings too. Deleting the same
// product again returns an error saying it has already been deleted.
//
// This endpoint requires the permission: `items:delete`.
func (r *CatalogProductService) Delete(ctx context.Context, id string, body CatalogProductDeleteParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Moves a product to a different product line.
//
// The target product line must be one your account owns or a shared system line;
// anything else fails as not found. Because customer accounts are granted access
// to whole product lines, moving a product changes which buyers can see and order
// it in the customer portal, and which default commission and freight policies
// apply to it.
//
// This endpoint requires the permission: `items:update`.
func (r *CatalogProductService) ChangeProductLine(ctx context.Context, productLineID string, params CatalogProductChangeProductLineParams, opts ...option.RequestOption) (res *Product, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if productLineID == "" {
		err = errors.New("missing required product_line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/products/%s/product-line/%s", url.PathEscape(params.ID), url.PathEscape(productLineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Request to create a product.
//
// The properties CategoryID, SKU, Type are required.
type CreateProductRequestParam struct {
	// ID of the item category for the product's item.
	//
	// The category's unit group determines the default units used for the product's
	// pricing rates and inventory tracking.
	CategoryID string `json:"category_id" api:"required"`
	// Stock keeping unit code for the product's item.
	//
	// Must be unique within the account; creation fails with a conflict error if
	// another item already uses it.
	SKU string `json:"sku" api:"required"`
	// Product type code, which determines how the product behaves on orders and
	// invoices.
	//
	// - `sale`: a standard sellable product.
	// - `service`: a non-physical service line, such as labor or installation.
	// - `shipping`: a shipping charge applied to an order.
	// - `credit`: a credit applied against an order or invoice.
	// - `return`: a returned product (RMA).
	// - `tax`: a tax line.
	//
	// Any of "sale", "service", "shipping", "credit", "return", "tax".
	Type CreateProductRequestType `json:"type,omitzero" api:"required"`
	// Free-form description of the product.
	Description param.Opt[string] `json:"description,omitzero"`
	// Free-form notes about the product.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ID of the product line to assign the product to.
	//
	// The product line must be one your account owns or a shared system line; anything
	// else fails as not found. Buyers are granted access to whole product lines, so a
	// product created without one never appears in the customer portal, whatever its
	// `portal_visibility`.
	ProductLineID param.Opt[string] `json:"product_line_id,omitzero"`
	// Attribute IDs to link to the product's item at creation time.
	//
	// Every ID must already exist in your account, and each attribute's property must
	// be one the item's category carries; an ID that fails either check fails the
	// whole request rather than being skipped.
	AttributeIDs []string `json:"attribute_ids,omitzero"`
	// Whether the product is shown to buyers in the customer portal.
	//
	//   - `visible`: buyers can see and order the product in the portal.
	//   - `hidden`: the product is concealed from the portal but remains usable
	//     internally.
	//
	// When omitted, the product is created hidden, so it must be set to `visible`
	// before buyers can see it.
	//
	// Any of "visible", "hidden".
	PortalVisibility CreateProductRequestPortalVisibility `json:"portal_visibility,omitzero"`
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

func (r CreateProductRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateProductRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateProductRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product type code, which determines how the product behaves on orders and
// invoices.
//
// - `sale`: a standard sellable product.
// - `service`: a non-physical service line, such as labor or installation.
// - `shipping`: a shipping charge applied to an order.
// - `credit`: a credit applied against an order or invoice.
// - `return`: a returned product (RMA).
// - `tax`: a tax line.
type CreateProductRequestType string

const (
	CreateProductRequestTypeSale     CreateProductRequestType = "sale"
	CreateProductRequestTypeService  CreateProductRequestType = "service"
	CreateProductRequestTypeShipping CreateProductRequestType = "shipping"
	CreateProductRequestTypeCredit   CreateProductRequestType = "credit"
	CreateProductRequestTypeReturn   CreateProductRequestType = "return"
	CreateProductRequestTypeTax      CreateProductRequestType = "tax"
)

// Whether the product is shown to buyers in the customer portal.
//
//   - `visible`: buyers can see and order the product in the portal.
//   - `hidden`: the product is concealed from the portal but remains usable
//     internally.
//
// When omitted, the product is created hidden, so it must be set to `visible`
// before buyers can see it.
type CreateProductRequestPortalVisibility string

const (
	CreateProductRequestPortalVisibilityVisible CreateProductRequestPortalVisibility = "visible"
	CreateProductRequestPortalVisibilityHidden  CreateProductRequestPortalVisibility = "hidden"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProduct struct {
	// Resources in this page.
	Data []Product `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductObject `json:"object" api:"required"`
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
func (r ListProduct) RawJSON() string { return r.JSON.raw }
func (r *ListProduct) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductObject string

const (
	ListProductObjectList ListProductObject = "list"
)

// A catalog entry as it is sold: an inventory item together with its product type,
// product line, and customer portal visibility.
//
// Every product is backed by exactly one item, which carries the SKU, description,
// pricing, attributes, and inventory position. Creating a product creates that
// item; deleting the product deletes it.
type Product struct {
	// Product ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An entry in your catalog: something you sell, consume, or build with.
	Item Item `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "product".
	Object ProductObject `json:"object" api:"required"`
	// Whether the product is shown to buyers in the customer portal.
	//
	//   - `visible`: buyers can see and order the product in the portal.
	//   - `hidden`: the product is concealed from the portal but remains usable
	//     internally.
	//
	// Visibility alone is not enough to expose a product: a buyer only sees it if
	// their account has also been granted access to the product's product line.
	//
	// Any of "visible", "hidden".
	PortalVisibility ProductPortalVisibility `json:"portal_visibility" api:"required"`
	// A named grouping of related products in your catalog.
	//
	// A product line carries the default commission and freight policies for the
	// products assigned to it, along with the unit group that determines how those
	// products are measured. Product lines are also the unit that catalog access is
	// granted over, for both customers and account groups.
	ProductLine ProductLine `json:"product_line" api:"required"`
	// Product type code, which determines how the product behaves on orders and
	// invoices.
	//
	// - `sale`: a standard sellable product.
	// - `service`: a non-physical service line, such as labor or installation.
	// - `shipping`: a shipping charge applied to an order.
	// - `credit`: a credit applied against an order or invoice.
	// - `return`: a returned product (RMA).
	// - `tax`: a tax line.
	//
	// Any of "sale", "service", "shipping", "credit", "return", "tax".
	Type ProductType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Item             respjson.Field
		Object           respjson.Field
		PortalVisibility respjson.Field
		ProductLine      respjson.Field
		Type             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Product) RawJSON() string { return r.JSON.raw }
func (r *Product) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductObject string

const (
	ProductObjectProduct ProductObject = "product"
)

// Whether the product is shown to buyers in the customer portal.
//
//   - `visible`: buyers can see and order the product in the portal.
//   - `hidden`: the product is concealed from the portal but remains usable
//     internally.
//
// Visibility alone is not enough to expose a product: a buyer only sees it if
// their account has also been granted access to the product's product line.
type ProductPortalVisibility string

const (
	ProductPortalVisibilityVisible ProductPortalVisibility = "visible"
	ProductPortalVisibilityHidden  ProductPortalVisibility = "hidden"
)

// Product type code, which determines how the product behaves on orders and
// invoices.
//
// - `sale`: a standard sellable product.
// - `service`: a non-physical service line, such as labor or installation.
// - `shipping`: a shipping charge applied to an order.
// - `credit`: a credit applied against an order or invoice.
// - `return`: a returned product (RMA).
// - `tax`: a tax line.
type ProductType string

const (
	ProductTypeSale     ProductType = "sale"
	ProductTypeService  ProductType = "service"
	ProductTypeShipping ProductType = "shipping"
	ProductTypeCredit   ProductType = "credit"
	ProductTypeReturn   ProductType = "return"
	ProductTypeTax      ProductType = "tax"
)

// Request to partially update a product.
type UpdateProductRequestParam struct {
	// Free-form description of the product.
	//
	// Send `null` to clear.
	Description param.Opt[string] `json:"description,omitzero"`
	// Free-form notes about the product.
	//
	// Send `null` to clear.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ID of the item category to move the product to.
	//
	// The move is the one Change Item Category makes: the category has to be a product
	// category and has to carry the properties of every attribute the product already
	// has, and the product's rate units switch to the category's base unit while their
	// numbers stay as they were. It is applied before the other fields in the request,
	// so a `unit_price` sent alongside is written after it.
	CategoryID param.Opt[string] `json:"category_id,omitzero"`
	// New stock keeping unit code for the product's item.
	//
	// Must be unique within the account; the update fails with a conflict error if
	// another item already uses it.
	SKU param.Opt[string] `json:"sku,omitzero"`
	// Whether the product is shown to buyers in the customer portal.
	//
	//   - `visible`: buyers can see and order the product in the portal.
	//   - `hidden`: the product is concealed from the portal but remains usable
	//     internally.
	//
	// Any of "visible", "hidden".
	PortalVisibility UpdateProductRequestPortalVisibility `json:"portal_visibility,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r UpdateProductRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateProductRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateProductRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the product is shown to buyers in the customer portal.
//
//   - `visible`: buyers can see and order the product in the portal.
//   - `hidden`: the product is concealed from the portal but remains usable
//     internally.
type UpdateProductRequestPortalVisibility string

const (
	UpdateProductRequestPortalVisibilityVisible UpdateProductRequestPortalVisibility = "visible"
	UpdateProductRequestPortalVisibilityHidden  UpdateProductRequestPortalVisibility = "hidden"
)

type CatalogProductNewParams struct {
	// Request to create a product.
	CreateProductRequest CreateProductRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogProductNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateProductRequest)
}
func (r *CatalogProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductNewParams]'s query parameters as
// `url.Values`.
func (r CatalogProductNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductGetParams]'s query parameters as
// `url.Values`.
func (r CatalogProductGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a product.
	UpdateProductRequest UpdateProductRequestParam
	paramObj
}

func (r CatalogProductUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateProductRequest)
}
func (r *CatalogProductUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductUpdateParams]'s query parameters as
// `url.Values`.
func (r CatalogProductUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// End of creation date range.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Start of creation date range.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Filter to products whose item carries at least one of these attributes.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Filter by the item category the product's item belongs to.
	CategoryIDs []string `query:"category_ids,omitzero" json:"-"`
	// Restrict results to products these customer accounts are entitled to buy.
	//
	// A product matches when its product line has been granted to the customer
	// directly, through the customer's account group, or through the account group
	// used for the customer's pricing. Combined with `product_line_ids` this widens
	// the results rather than narrowing them: products matching either filter are
	// returned.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by customer portal visibility.
	//
	// Any of "visible", "hidden".
	PortalVisibility CatalogProductListParamsPortalVisibility `query:"portal_visibility,omitzero" json:"-"`
	// Filter by product line IDs.
	//
	// Combined with `customer_ids`, products matching either filter are returned.
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductListParams]'s query parameters as
// `url.Values`.
func (r CatalogProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by customer portal visibility.
type CatalogProductListParamsPortalVisibility string

const (
	CatalogProductListParamsPortalVisibilityVisible CatalogProductListParamsPortalVisibility = "visible"
	CatalogProductListParamsPortalVisibilityHidden  CatalogProductListParamsPortalVisibility = "hidden"
)

type CatalogProductDeleteParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductDeleteParams]'s query parameters as
// `url.Values`.
func (r CatalogProductDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogProductChangeProductLineParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product_line", "product_line.unit_group",
	// "product_line.unit_group.base_unit", "product_line.unit_group.associated_units",
	// "product_line.unit_group.associated_units.unit", "item", "item.category",
	// "item.category.properties", "item.category.unit_group",
	// "item.category.unit_group.base_unit",
	// "item.category.unit_group.associated_units",
	// "item.category.unit_group.associated_units.unit", "item.unit_value",
	// "item.unit_cost", "item.burn_rate", "item.attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogProductChangeProductLineParams]'s query parameters
// as `url.Values`.
func (r CatalogProductChangeProductLineParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
