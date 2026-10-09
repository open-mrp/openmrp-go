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

// List and manage product lines.
//
// CatalogProductLineActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogProductLineActionService] method instead.
type CatalogProductLineActionService struct {
	options []option.RequestOption
}

// NewCatalogProductLineActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCatalogProductLineActionService(opts ...option.RequestOption) (r CatalogProductLineActionService) {
	r = CatalogProductLineActionService{}
	r.options = opts
	return
}

// Creates or updates multiple product lines for the account, matched by name
// (case-insensitive), then writes asynchronously — 202 with a job to poll.
//
// This endpoint requires the permissions: `product_lines:create` and
// `product_lines:update`.
func (r *CatalogProductLineActionService) BulkUpsert(ctx context.Context, params CatalogProductLineActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/product-lines/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// BulkUpsertProductLinesRequest is the request to bulk upsert product lines.
//
// The property ProductLines is required.
type BulkUpsertProductLinesRequestParam struct {
	// Product lines to create or update, matched by name within the account.
	ProductLines []UpsertProductLineInputParam `json:"product_lines,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertProductLinesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertProductLinesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertProductLinesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertProductLineInput is the input for a single product line in a bulk upsert
// operation.
//
// The properties CommissionPolicy, FreightPolicy, Name, UnitGroup are required.
type UpsertProductLineInputParam struct {
	// Default commission policy for products in this product line.
	//
	//   - `commission_exempt`: no commission applies to these products.
	//   - `commission_applied`: commission applies to these products, unless overridden
	//     elsewhere.
	//
	// Any of "commission_applied", "commission_exempt".
	CommissionPolicy UpsertProductLineInputCommissionPolicy `json:"commission_policy,omitzero" api:"required"`
	// Default freight policy for products in this product line.
	//
	//   - `free_freight`: these products do not incur a freight charge.
	//   - `billed_freight`: freight is billed for these products, unless overridden
	//     elsewhere.
	//
	// Any of "free_freight", "billed_freight".
	FreightPolicy UpsertProductLineInputFreightPolicy `json:"freight_policy,omitzero" api:"required"`
	// Display name of the product line, matched case-insensitively against existing
	// lines. A row matching a system product line fails — system lines cannot be
	// modified.
	Name string `json:"name" api:"required"`
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	UnitGroup ObjectIdentifierParam `json:"unit_group,omitzero" api:"required"`
	paramObj
}

func (r UpsertProductLineInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertProductLineInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertProductLineInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default commission policy for products in this product line.
//
//   - `commission_exempt`: no commission applies to these products.
//   - `commission_applied`: commission applies to these products, unless overridden
//     elsewhere.
type UpsertProductLineInputCommissionPolicy string

const (
	UpsertProductLineInputCommissionPolicyCommissionApplied UpsertProductLineInputCommissionPolicy = "commission_applied"
	UpsertProductLineInputCommissionPolicyCommissionExempt  UpsertProductLineInputCommissionPolicy = "commission_exempt"
)

// Default freight policy for products in this product line.
//
//   - `free_freight`: these products do not incur a freight charge.
//   - `billed_freight`: freight is billed for these products, unless overridden
//     elsewhere.
type UpsertProductLineInputFreightPolicy string

const (
	UpsertProductLineInputFreightPolicyFreeFreight   UpsertProductLineInputFreightPolicy = "free_freight"
	UpsertProductLineInputFreightPolicyBilledFreight UpsertProductLineInputFreightPolicy = "billed_freight"
)

type CatalogProductLineActionBulkUpsertParams struct {
	// BulkUpsertProductLinesRequest is the request to bulk upsert product lines.
	BulkUpsertProductLinesRequest BulkUpsertProductLinesRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogProductLineActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertProductLinesRequest)
}
func (r *CatalogProductLineActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogProductLineActionBulkUpsertParams]'s query
// parameters as `url.Values`.
func (r CatalogProductLineActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
