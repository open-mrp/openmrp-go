// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
)

// List and manage products.
//
// CatalogProductActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogProductActionService] method instead.
type CatalogProductActionService struct {
	options []option.RequestOption
}

// NewCatalogProductActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogProductActionService(opts ...option.RequestOption) (r CatalogProductActionService) {
	r = CatalogProductActionService{}
	r.options = opts
	return
}

// Creates or updates multiple products for the account, matched by SKU. Validates
// and resolves synchronously, then writes asynchronously — 202 with a job to poll.
//
// At most 1,000 products and an 8 MB request body per call.
//
// This endpoint requires the permissions: `items:create` and `items:update`.
func (r *CatalogProductActionService) BulkUpsert(ctx context.Context, params CatalogProductActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/products/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to bulk upsert products.
//
// The property Products is required.
type BulkUpsertProductsRequestParam struct {
	// Products to create or update, matched by SKU within the account.
	Products []UpsertProductInputParam `json:"products,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertProductsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertProductsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertProductsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for a single product in a bulk upsert operation.
//
// The properties Category, Properties, SKU are required.
type UpsertProductInputParam struct {
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	Category ObjectIdentifierParam `json:"category,omitzero" api:"required"`
	// Properties to attach to the product, matched/created by name + value. Additive —
	// existing attributes are not removed.
	Properties []UpsertProductPropertyParam `json:"properties,omitzero" api:"required"`
	// SKU for the product, used to match an existing product within the account. If it
	// exists the product is updated in place; otherwise a new product is created. A
	// SKU already used by a non-product item fails that row.
	SKU string `json:"sku" api:"required"`
	// Product description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Product notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Whether the product is shown to buyers in the customer portal. Defaults to
	// `hidden` on create; preserved when omitted on update.
	//
	// Any of "visible", "hidden".
	PortalVisibility UpsertProductInputPortalVisibility `json:"portal_visibility,omitzero"`
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	ProductLine ObjectIdentifierParam `json:"product_line,omitzero"`
	// Product type. Create-only; defaults to `sale` when omitted.
	//
	// Any of "sale", "service", "shipping", "credit", "return", "tax".
	Type UpsertProductInputType `json:"type,omitzero"`
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

func (r UpsertProductInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertProductInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertProductInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the product is shown to buyers in the customer portal. Defaults to
// `hidden` on create; preserved when omitted on update.
type UpsertProductInputPortalVisibility string

const (
	UpsertProductInputPortalVisibilityVisible UpsertProductInputPortalVisibility = "visible"
	UpsertProductInputPortalVisibilityHidden  UpsertProductInputPortalVisibility = "hidden"
)

// Product type. Create-only; defaults to `sale` when omitted.
type UpsertProductInputType string

const (
	UpsertProductInputTypeSale     UpsertProductInputType = "sale"
	UpsertProductInputTypeService  UpsertProductInputType = "service"
	UpsertProductInputTypeShipping UpsertProductInputType = "shipping"
	UpsertProductInputTypeCredit   UpsertProductInputType = "credit"
	UpsertProductInputTypeReturn   UpsertProductInputType = "return"
	UpsertProductInputTypeTax      UpsertProductInputType = "tax"
)

// Property name + value pair attached to a product. The property and its value (an
// attribute) are created if they do not yet exist.
//
// The properties Name, Value are required.
type UpsertProductPropertyParam struct {
	// Property name (e.g. "Color"). Matched case-insensitively; created if missing.
	Name string `json:"name" api:"required"`
	// Property value (e.g. "Red"). Matched case-insensitively; created under the
	// property if missing. A value already in use under a different property fails the
	// whole job.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r UpsertProductPropertyParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertProductPropertyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertProductPropertyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogProductActionBulkUpsertParams struct {
	// Request to bulk upsert products.
	BulkUpsertProductsRequest BulkUpsertProductsRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogProductActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertProductsRequest)
}
func (r *CatalogProductActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductActionBulkUpsertParams]'s query parameters as
// `url.Values`.
func (r CatalogProductActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
