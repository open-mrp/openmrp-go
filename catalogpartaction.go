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

// List and manage parts.
//
// CatalogPartActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogPartActionService] method instead.
type CatalogPartActionService struct {
	options []option.RequestOption
}

// NewCatalogPartActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogPartActionService(opts ...option.RequestOption) (r CatalogPartActionService) {
	r = CatalogPartActionService{}
	r.options = opts
	return
}

// Creates or updates multiple parts for the account, matched by SKU, then writes
// asynchronously — 202 with a job to poll.
//
// At most 1,000 parts and an 8 MB request body per call.
//
// This endpoint requires the permissions: `parts:create` and `parts:update`.
func (r *CatalogPartActionService) BulkUpsert(ctx context.Context, params CatalogPartActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/parts/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// BulkUpsertPartsRequest is the request to bulk upsert parts.
//
// The property Parts is required.
type BulkUpsertPartsRequestParam struct {
	// Parts to create or update, matched by SKU within the account.
	Parts []UpsertPartInputParam `json:"parts,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertPartsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertPartsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertPartsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertPartInput is the input for a single part in a bulk upsert operation.
//
// The properties Category, Properties, SKU are required.
type UpsertPartInputParam struct {
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	Category ObjectIdentifierParam `json:"category,omitzero" api:"required"`
	// Properties to attach to the part, matched/created by name + value. Additive —
	// existing attributes are not removed.
	Properties []UpsertPartPropertyParam `json:"properties,omitzero" api:"required"`
	// SKU for the part, matched against existing parts in the account: a match updates
	// in place, otherwise a part is created. A SKU held by a non-part item fails that
	// row.
	SKU string `json:"sku" api:"required"`
	// Free-form description of the part.
	Description param.Opt[string] `json:"description,omitzero"`
	// Free-form notes about the part.
	Notes param.Opt[string] `json:"notes,omitzero"`
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

func (r UpsertPartInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertPartInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertPartInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertPartProperty is a property name + value pair attached to a part. The
// property and its value (an attribute) are created if they do not yet exist.
//
// The properties Name, Value are required.
type UpsertPartPropertyParam struct {
	// Property name (e.g. "Material"). Matched case-insensitively; created if missing.
	Name string `json:"name" api:"required"`
	// Property value (e.g. "Steel"). Matched exactly; created under the property if
	// missing.
	Value string `json:"value" api:"required" format:"decimal"`
	paramObj
}

func (r UpsertPartPropertyParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertPartPropertyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertPartPropertyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogPartActionBulkUpsertParams struct {
	// BulkUpsertPartsRequest is the request to bulk upsert parts.
	BulkUpsertPartsRequest BulkUpsertPartsRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogPartActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertPartsRequest)
}
func (r *CatalogPartActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogPartActionBulkUpsertParams]'s query parameters as
// `url.Values`.
func (r CatalogPartActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
