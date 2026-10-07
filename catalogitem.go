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
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage inventory items.
//
// CatalogItemService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemService] method instead.
type CatalogItemService struct {
	options []option.RequestOption
	// List and manage inventory items.
	Inventory CatalogItemInventoryService
	// List and manage inventory items.
	Attributes CatalogItemAttributeService
	// List and manage inventory items.
	Actions CatalogItemActionService
}

// NewCatalogItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCatalogItemService(opts ...option.RequestOption) (r CatalogItemService) {
	r = CatalogItemService{}
	r.options = opts
	r.Inventory = NewCatalogItemInventoryService(opts...)
	r.Attributes = NewCatalogItemAttributeService(opts...)
	r.Actions = NewCatalogItemActionService(opts...)
	return
}

// Returns a single item by ID.
//
// This endpoint requires the permission: `items:read`.
func (r *CatalogItemService) Get(ctx context.Context, id string, query CatalogItemGetParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of items, newest first.
//
// Items backed by a non-sale product — the service, shipping, tax, credit, and
// return products that carry charges on orders — are left out, so this reflects
// the catalog you sell and stock rather than every item row. `q` matches against
// SKU and description, with closer SKU matches ranked first.
//
// This endpoint requires the permission: `items:read`.
func (r *CatalogItemService) List(ctx context.Context, query CatalogItemListParams, opts ...option.RequestOption) (res *ListItem, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/items"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Moves an item to a different category and returns the updated item.
//
// The item's rate units (unit value, unit cost, burn rate) and any related
// order-point, consumption, and production quantity units are switched to the new
// category's base unit. Only the units change — the numbers attached to them are
// carried over as they were, so review any figure whose meaning depends on the
// unit after moving between categories that count differently.
//
// Re-assigning the item's current category succeeds and changes nothing.
//
// This endpoint requires the permission: `items:update`.
func (r *CatalogItemService) ChangeCategory(ctx context.Context, categoryID string, params CatalogItemChangeCategoryParams, opts ...option.RequestOption) (res *Item, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/category/%s", url.PathEscape(params.ID), url.PathEscape(categoryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the lot this item is made in — how many, counted in what.
//
// A lot is a doff, a pallet, a batch: the quantity production is issued in. The
// unit is what makes it meaningful, since 60 pairs and 60 eaches are different
// lots, so `quantity` should never be read without `unit`.
//
// Resolved through the same chain the production schedule uses, most specific
// first: a per-item override, then the item's own product line, then the product
// lines of the finished goods it becomes, then the account-wide default. `source`
// names which rule applied. Intermediate items like greige are not sold and have
// no product line of their own, which is why they inherit from what they become.
//
// `quantity` is `0` when nothing in the chain supplies a lot. That means the item
// has no lot convention, not that its lot is zero.
//
// This endpoint requires the permission: `items:read`.
func (r *CatalogItemService) GetLotDefault(ctx context.Context, id string, query CatalogItemGetLotDefaultParams, opts ...option.RequestOption) (res *ItemLotDefault, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/items/%s/lot-default", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// An entry in your catalog: something you sell, consume, or build with.
type Item struct {
	// Item ID.
	ID string `json:"id" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Attributes ListAttribute `json:"attributes" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	BurnRate Rate `json:"burn_rate" api:"required"`
	// A grouping of related catalog items that defines the unit group and properties
	// available to the items within it.
	Category ItemCategory `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Item description.
	Description string `json:"description" api:"required"`
	// Free-form notes about the item.
	//
	// Null to customer and supplier portal users: they are your own team's notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "item".
	Object ItemObject `json:"object" api:"required"`
	// Stock keeping unit code, unique within the account.
	SKU string `json:"sku" api:"required"`
	// What kind of item this is.
	//
	// - `product`: a finished product.
	// - `material`: a raw material or component consumed in production.
	// - `part`: a part used in production.
	//
	// Any of "product", "material", "part".
	Type ItemType `json:"type" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitCost Rate `json:"unit_cost" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	UnitValue Rate `json:"unit_value" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attributes  respjson.Field
		BurnRate    respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Notes       respjson.Field
		Object      respjson.Field
		SKU         respjson.Field
		Type        respjson.Field
		UnitCost    respjson.Field
		UnitValue   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Item) RawJSON() string { return r.JSON.raw }
func (r *Item) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemObject string

const (
	ItemObjectItem ItemObject = "item"
)

// What kind of item this is.
//
// - `product`: a finished product.
// - `material`: a raw material or component consumed in production.
// - `part`: a part used in production.
type ItemType string

const (
	ItemTypeProduct  ItemType = "product"
	ItemTypeMaterial ItemType = "material"
	ItemTypePart     ItemType = "part"
)

// A grouping of related catalog items that defines the unit group and properties
// available to the items within it.
type ItemCategory struct {
	// Item category ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the item category.
	Name string `json:"name" api:"required"`
	// Free-form notes about the item category.
	//
	// Null to customer and supplier portal users: they are your own team's notes.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "item_category".
	Object ItemCategoryObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Properties ListProperty `json:"properties" api:"required"`
	// What kind of items this category groups.
	//
	//   - `material_category`: groups raw materials and components (items of type
	//     `material`).
	//   - `product_category`: groups finished products and parts (items of type
	//     `product` or `part`).
	//
	// An item can only be assigned to a category whose type matches the item's `type`,
	// and the category's type is fixed at creation.
	//
	// Any of "material_category", "product_category".
	Type ItemCategoryType `json:"type" api:"required"`
	// A named collection of units that share one dimension, defining which units a
	// product can be ordered in.
	//
	// Each associated unit carries its own discount and customer portal visibility,
	// applied when an order line is priced in that unit. A product takes its unit
	// group from its product line, falling back to its item category.
	UnitGroup UnitGroup `json:"unit_group" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		UnitGroup   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemCategory) RawJSON() string { return r.JSON.raw }
func (r *ItemCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemCategoryObject string

const (
	ItemCategoryObjectItemCategory ItemCategoryObject = "item_category"
)

// What kind of items this category groups.
//
//   - `material_category`: groups raw materials and components (items of type
//     `material`).
//   - `product_category`: groups finished products and parts (items of type
//     `product` or `part`).
//
// An item can only be assigned to a category whose type matches the item's `type`,
// and the category's type is fixed at creation.
type ItemCategoryType string

const (
	ItemCategoryTypeMaterialCategory ItemCategoryType = "material_category"
	ItemCategoryTypeProductCategory  ItemCategoryType = "product_category"
)

// The lot an item is made in — how many, counted in what.
//
// A lot is the quantity production is issued in: a doff, a pallet, a batch. The
// unit is what makes it meaningful, since 60 pairs and 60 eaches are different
// lots.
type ItemLotDefault struct {
	// Entity is a polymorphic reference to any resource in the system.
	Item Entity `json:"item" api:"required"`
	// Resource type identifier.
	//
	// Any of "item_lot_default".
	Object ItemLotDefaultObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	ProductLine Entity `json:"product_line" api:"required"`
	// Units in one lot.
	//
	// `0` means the item has no lot convention, not that its lot is zero.
	Quantity float64 `json:"quantity" api:"required"`
	// Which rule in the chain produced this lot.
	//
	//   - `item_override`: a lot size set on the item itself.
	//   - `product_line`: the convention of the line the item sells under.
	//   - `downstream_product_line`: inherited from the finished goods this item
	//     becomes, for intermediates that are not themselves sold.
	//   - `account_default`: the account-wide fallback.
	//
	// Empty when no rule in the chain supplies a lot, which is the same case
	// `quantity` reports as `0`.
	//
	// Any of "item_override", "product_line", "downstream_product_line",
	// "account_default", "".
	Source ItemLotDefaultSource `json:"source" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	Unit Unit `json:"unit" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Item        respjson.Field
		Object      respjson.Field
		ProductLine respjson.Field
		Quantity    respjson.Field
		Source      respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ItemLotDefault) RawJSON() string { return r.JSON.raw }
func (r *ItemLotDefault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ItemLotDefaultObject string

const (
	ItemLotDefaultObjectItemLotDefault ItemLotDefaultObject = "item_lot_default"
)

// Which rule in the chain produced this lot.
//
//   - `item_override`: a lot size set on the item itself.
//   - `product_line`: the convention of the line the item sells under.
//   - `downstream_product_line`: inherited from the finished goods this item
//     becomes, for intermediates that are not themselves sold.
//   - `account_default`: the account-wide fallback.
//
// Empty when no rule in the chain supplies a lot, which is the same case
// `quantity` reports as `0`.
type ItemLotDefaultSource string

const (
	ItemLotDefaultSourceItemOverride          ItemLotDefaultSource = "item_override"
	ItemLotDefaultSourceProductLine           ItemLotDefaultSource = "product_line"
	ItemLotDefaultSourceDownstreamProductLine ItemLotDefaultSource = "downstream_product_line"
	ItemLotDefaultSourceAccountDefault        ItemLotDefaultSource = "account_default"
	ItemLotDefaultSourceEmpty                 ItemLotDefaultSource = ""
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListItem struct {
	// Resources in this page.
	Data []Item `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListItemObject `json:"object" api:"required"`
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
func (r ListItem) RawJSON() string { return r.JSON.raw }
func (r *ListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListItemObject string

const (
	ListItemObjectList ListItemObject = "list"
)

// Value expressed as a ratio of two units, such as a price per kilogram or a
// throughput per hour.
type Rate struct {
	// Rate ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unit of measurement used for conversions and product quantities.
	DenominatorUnit Unit `json:"denominator_unit" api:"required"`
	// Human-readable formatted value (e.g. "$25.50 / kg" or "100 kg / hr").
	DisplayValue string `json:"display_value" api:"required"`
	// Unit of measurement used for conversions and product quantities.
	NumeratorUnit Unit `json:"numerator_unit" api:"required"`
	// Resource type identifier.
	//
	// Any of "rate".
	Object RateObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Decimal value of the rate, as a string to preserve precision.
	//
	// Expressed as the amount of the numerator unit per one denominator unit.
	Value string `json:"value" api:"required" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		DenominatorUnit respjson.Field
		DisplayValue    respjson.Field
		NumeratorUnit   respjson.Field
		Object          respjson.Field
		UpdatedAt       respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rate) RawJSON() string { return r.JSON.raw }
func (r *Rate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RateObject string

const (
	RateObjectRate RateObject = "rate"
)

type CatalogItemGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemGetParams]'s query parameters as `url.Values`.
func (r CatalogItemGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to items created on or before this date.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Filter to items created on or after this date.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Filter to materials this supplier account supplies to you.
	//
	// Only materials can have suppliers, so combining this with a `types` filter that
	// excludes `material` returns nothing.
	SupplierID param.Opt[string] `query:"supplier_id,omitzero" json:"-"`
	// Filter to items carrying any of these attributes.
	AttributeIDs []string `query:"attribute_ids,omitzero" json:"-"`
	// Filter to items in any of these categories.
	CategoryIDs []string `query:"category_ids,omitzero" json:"-"`
	// Filter to items any of these customers are allowed to order.
	//
	// A customer qualifies when its relationship, its account group, or its price
	// group grants access to the product line the item's product sits in. Items with
	// no product line, including materials and parts, never match.
	CustomerIDs []string `query:"customer_ids,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter to items whose product belongs to any of these product lines.
	ProductLineIDs []string `query:"product_line_ids,omitzero" json:"-"`
	// Restricts results based on where the item is produced in its production flow.
	//
	//   - `all`: no restriction.
	//   - `initial_only`: only items produced by an initial production step, i.e. a step
	//     with no upstream steps feeding into it.
	//
	// Any of "all", "initial_only".
	SubassemblyFilter CatalogItemListParamsSubassemblyFilter `query:"subassembly_filter,omitzero" json:"-"`
	// Filter to items of these types (`product`, `material`, `part`).
	//
	// Any of "product", "material", "part".
	Types []string `query:"types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemListParams]'s query parameters as `url.Values`.
func (r CatalogItemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restricts results based on where the item is produced in its production flow.
//
//   - `all`: no restriction.
//   - `initial_only`: only items produced by an initial production step, i.e. a step
//     with no upstream steps feeding into it.
type CatalogItemListParamsSubassemblyFilter string

const (
	CatalogItemListParamsSubassemblyFilterAll         CatalogItemListParamsSubassemblyFilter = "all"
	CatalogItemListParamsSubassemblyFilterInitialOnly CatalogItemListParamsSubassemblyFilter = "initial_only"
)

type CatalogItemChangeCategoryParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "category", "unit_value", "unit_cost", "burn_rate", "attributes",
	// "category.unit_group", "category.properties", "category.unit_group.base_unit",
	// "category.unit_group.associated_units",
	// "category.unit_group.associated_units.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemChangeCategoryParams]'s query parameters as
// `url.Values`.
func (r CatalogItemChangeCategoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemGetLotDefaultParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CatalogItemGetLotDefaultParams]'s query parameters as
// `url.Values`.
func (r CatalogItemGetLotDefaultParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
