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

// List and manage materials.
//
// CatalogMaterialActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogMaterialActionService] method instead.
type CatalogMaterialActionService struct {
	options []option.RequestOption
}

// NewCatalogMaterialActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogMaterialActionService(opts ...option.RequestOption) (r CatalogMaterialActionService) {
	r = CatalogMaterialActionService{}
	r.options = opts
	return
}

// Creates or updates multiple materials for the account, matched by SKU. Validates
// and resolves synchronously, then writes asynchronously — 202 with a job to poll.
//
// At most 1,000 materials and an 8 MB request body per call.
//
// This endpoint requires the permissions: `materials:create` and
// `materials:update`.
func (r *CatalogMaterialActionService) BulkUpsert(ctx context.Context, params CatalogMaterialActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/materials/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to bulk upsert materials.
//
// The property Materials is required.
type BulkUpsertMaterialsRequestParam struct {
	// Materials to create or update, matched by SKU within the account.
	Materials []UpsertMaterialInputParam `json:"materials,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertMaterialsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertMaterialsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertMaterialsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input for a single material in a bulk upsert operation.
//
// The properties Category, Properties, SKU are required.
type UpsertMaterialInputParam struct {
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	Category ObjectIdentifierParam `json:"category,omitzero" api:"required"`
	// Properties to attach to the material, matched/created by name + value. Additive
	// — existing attributes are not removed.
	Properties []UpsertMaterialPropertyParam `json:"properties,omitzero" api:"required"`
	// SKU for the material, used to match an existing material within the account. If
	// it exists the material is updated in place; otherwise a new material is created.
	// A SKU already used by a non-material item fails that row.
	SKU string `json:"sku" api:"required"`
	// Material description.
	Description param.Opt[string] `json:"description,omitzero"`
	// Material notes.
	Notes param.Opt[string] `json:"notes,omitzero"`
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

func (r UpsertMaterialInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertMaterialInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertMaterialInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Property name + value pair attached to a material. The property and its value
// (an attribute) are created if they do not yet exist.
//
// The properties Name, Value are required.
type UpsertMaterialPropertyParam struct {
	// Property name (e.g. "Grade"). Matched case-insensitively; created if missing.
	Name string `json:"name" api:"required"`
	// Property value (e.g. "A36"). Matched case-insensitively; created under the
	// property if missing. A value already in use under a different property fails the
	// whole job.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r UpsertMaterialPropertyParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertMaterialPropertyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertMaterialPropertyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogMaterialActionBulkUpsertParams struct {
	// Request to bulk upsert materials.
	BulkUpsertMaterialsRequest BulkUpsertMaterialsRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogMaterialActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertMaterialsRequest)
}
func (r *CatalogMaterialActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogMaterialActionBulkUpsertParams]'s query parameters
// as `url.Values`.
func (r CatalogMaterialActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
