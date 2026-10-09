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

// List and manage item categories.
//
// CatalogItemCategoryActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemCategoryActionService] method instead.
type CatalogItemCategoryActionService struct {
	options []option.RequestOption
}

// NewCatalogItemCategoryActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCatalogItemCategoryActionService(opts ...option.RequestOption) (r CatalogItemCategoryActionService) {
	r = CatalogItemCategoryActionService{}
	r.options = opts
	return
}

// Creates or updates multiple item categories for the account, matched by name
// (case-insensitive), then writes asynchronously — 202 with a job to poll.
//
// This endpoint requires the permissions: `item_categories:create` and
// `item_categories:update`.
func (r *CatalogItemCategoryActionService) BulkUpsert(ctx context.Context, params CatalogItemCategoryActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/item-categories/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// BulkUpsertItemCategoriesRequest is the request to bulk upsert item categories.
//
// The property ItemCategories is required.
type BulkUpsertItemCategoriesRequestParam struct {
	// Item categories to create or update, matched by name within the account.
	ItemCategories []UpsertItemCategoryInputParam `json:"item_categories,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertItemCategoriesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertItemCategoriesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertItemCategoriesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// -------------------------- Named Object -------------------------- Identifies an
// object by its id or its name. An id wins when both are given.
//
// The properties ID, Name are required.
type ObjectIdentifierParam struct {
	// Object ID.
	ID string `json:"id" api:"required"`
	// Object name, matched case-insensitively.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r ObjectIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertItemCategoryInput is the input for a single item category in a bulk upsert
// operation.
//
// The properties Name, Notes, PropertyNames, Type, UnitGroup are required.
type UpsertItemCategoryInputParam struct {
	// Optional notes.
	Notes param.Opt[string] `json:"notes,omitzero" api:"required"`
	// Display name of the item category, used to match existing categories.
	Name string `json:"name" api:"required"`
	// Optional list of property names to attach to this category. Properties are
	// matched by name (case-insensitive) within the account; names not found are
	// created automatically. Relations are additive — existing relations are not
	// removed.
	PropertyNames []string `json:"property_names,omitzero" api:"required"`
	// Item category type code. Create-only.
	//
	// Any of "material_category", "product_category".
	Type UpsertItemCategoryInputType `json:"type,omitzero" api:"required"`
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	UnitGroup ObjectIdentifierParam `json:"unit_group,omitzero" api:"required"`
	paramObj
}

func (r UpsertItemCategoryInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertItemCategoryInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertItemCategoryInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Item category type code. Create-only.
type UpsertItemCategoryInputType string

const (
	UpsertItemCategoryInputTypeMaterialCategory UpsertItemCategoryInputType = "material_category"
	UpsertItemCategoryInputTypeProductCategory  UpsertItemCategoryInputType = "product_category"
)

type CatalogItemCategoryActionBulkUpsertParams struct {
	// BulkUpsertItemCategoriesRequest is the request to bulk upsert item categories.
	BulkUpsertItemCategoriesRequest BulkUpsertItemCategoriesRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogItemCategoryActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertItemCategoriesRequest)
}
func (r *CatalogItemCategoryActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogItemCategoryActionBulkUpsertParams]'s query
// parameters as `url.Values`.
func (r CatalogItemCategoryActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
