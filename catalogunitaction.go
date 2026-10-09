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

// List and manage units.
//
// CatalogUnitActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitActionService] method instead.
type CatalogUnitActionService struct {
	options []option.RequestOption
}

// NewCatalogUnitActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogUnitActionService(opts ...option.RequestOption) (r CatalogUnitActionService) {
	r = CatalogUnitActionService{}
	r.options = opts
	return
}

// Creates or updates multiple units of measure for the account, matched by name or
// abbreviation, then writes asynchronously — 202 with a job to poll.
//
// This endpoint requires the permissions: `units:create` and `units:update`.
func (r *CatalogUnitActionService) BulkUpsert(ctx context.Context, params CatalogUnitActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/units/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// BulkUpsertUnitsRequest is the request to bulk upsert units.
//
// The property Units is required.
type BulkUpsertUnitsRequestParam struct {
	// Units to create or update, matched by name or abbreviation within the account.
	Units []UpsertUnitInputParam `json:"units,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertUnitsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertUnitsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertUnitsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertUnitInput is the input for a single unit in a bulk upsert operation.
//
// The properties Abbreviation, IsBaseUnit, Name, OffsetDenominator,
// OffsetNumerator, RatioDenominator, RatioNumerator, Type are required.
type UpsertUnitInputParam struct {
	// Short abbreviation for the unit (e.g. "g"). Also used for matching — see `name`.
	Abbreviation string `json:"abbreviation" api:"required"`
	// Whether the unit is its dimension's base unit. Bulk upsert never creates a base
	// unit and rejects a change to an existing one.
	IsBaseUnit bool `json:"is_base_unit" api:"required"`
	// Display name of the unit (e.g. "Gram"). A row matching a system unit fails —
	// system units cannot be modified.
	Name string `json:"name" api:"required"`
	// Conversion offset denominator, as a decimal string.
	OffsetDenominator string `json:"offset_denominator" api:"required" format:"decimal"`
	// Conversion offset numerator, as a decimal string.
	OffsetNumerator string `json:"offset_numerator" api:"required" format:"decimal"`
	// Conversion ratio denominator relative to the base unit, as a decimal string.
	RatioDenominator string `json:"ratio_denominator" api:"required" format:"decimal"`
	// Conversion ratio numerator relative to the base unit, as a decimal string.
	RatioNumerator string `json:"ratio_numerator" api:"required" format:"decimal"`
	// Unit dimension code. Create-only — a row that changes an existing unit's
	// dimension fails.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type UpsertUnitInputType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r UpsertUnitInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertUnitInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertUnitInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unit dimension code. Create-only — a row that changes an existing unit's
// dimension fails.
type UpsertUnitInputType string

const (
	UpsertUnitInputTypeCurrency    UpsertUnitInputType = "currency"
	UpsertUnitInputTypeQuantity    UpsertUnitInputType = "quantity"
	UpsertUnitInputTypeTime        UpsertUnitInputType = "time"
	UpsertUnitInputTypeMass        UpsertUnitInputType = "mass"
	UpsertUnitInputTypeVolume      UpsertUnitInputType = "volume"
	UpsertUnitInputTypeLength      UpsertUnitInputType = "length"
	UpsertUnitInputTypeTemperature UpsertUnitInputType = "temperature"
	UpsertUnitInputTypeArea        UpsertUnitInputType = "area"
)

type CatalogUnitActionBulkUpsertParams struct {
	// BulkUpsertUnitsRequest is the request to bulk upsert units.
	BulkUpsertUnitsRequest BulkUpsertUnitsRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogUnitActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertUnitsRequest)
}
func (r *CatalogUnitActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitActionBulkUpsertParams]'s query parameters as
// `url.Values`.
func (r CatalogUnitActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
