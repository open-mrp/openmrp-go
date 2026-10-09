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

// List and manage unit groups and their associated units.
//
// CatalogUnitGroupActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogUnitGroupActionService] method instead.
type CatalogUnitGroupActionService struct {
	options []option.RequestOption
}

// NewCatalogUnitGroupActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogUnitGroupActionService(opts ...option.RequestOption) (r CatalogUnitGroupActionService) {
	r = CatalogUnitGroupActionService{}
	r.options = opts
	return
}

// Creates or updates multiple unit groups for the account, matched by name
// (case-insensitive), then writes asynchronously — 202 with a job to poll.
//
// This endpoint requires the permissions: `unit_groups:create` and
// `unit_groups:update`.
func (r *CatalogUnitGroupActionService) BulkUpsert(ctx context.Context, params CatalogUnitGroupActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/unit-groups/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// BulkUpsertUnitGroupsRequest is the request to bulk upsert unit groups.
//
// The property UnitGroups is required.
type BulkUpsertUnitGroupsRequestParam struct {
	// Unit groups to create or update, matched by name within the account.
	UnitGroups []UpsertUnitGroupInputParam `json:"unit_groups,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertUnitGroupsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertUnitGroupsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertUnitGroupsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// -------------------------- UNIT -------------------------- Identifies a unit by
// its id, its name, or its abbreviation, in that order of precedence.
//
// The properties ID, Abbreviation, Name are required.
type UnitIdentifierParam struct {
	// Unit ID.
	ID string `json:"id" api:"required"`
	// Unit abbreviation, matched case-insensitively against the account's units.
	Abbreviation string `json:"abbreviation" api:"required"`
	// Unit name, matched case-insensitively against the account's units.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r UnitIdentifierParam) MarshalJSON() (data []byte, err error) {
	type shadow UnitIdentifierParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitIdentifierParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertUnitGroupConversionInput is the input for a single unit conversion within
// a bulk upsert unit group.
//
// The property Unit is required.
type UpsertUnitGroupConversionInputParam struct {
	// -------------------------- UNIT -------------------------- Identifies a unit by
	// its id, its name, or its abbreviation, in that order of precedence.
	Unit UnitIdentifierParam `json:"unit,omitzero" api:"required"`
	// Discount percentage to apply for this unit conversion.
	DiscountPercentage param.Opt[float64] `json:"discount_percentage,omitzero"`
	paramObj
}

func (r UpsertUnitGroupConversionInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertUnitGroupConversionInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertUnitGroupConversionInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertUnitGroupInput is the input for a single unit group in a bulk upsert
// operation.
//
// The properties BaseUnit, Name, Type are required.
type UpsertUnitGroupInputParam struct {
	// -------------------------- UNIT -------------------------- Identifies a unit by
	// its id, its name, or its abbreviation, in that order of precedence.
	BaseUnit UnitIdentifierParam `json:"base_unit,omitzero" api:"required"`
	// Display name of the unit group, matched case-insensitively against existing
	// groups. A row matching a system unit group fails — system groups cannot be
	// modified.
	Name string `json:"name" api:"required"`
	// Unit dimension type. Create-only — an existing group keeps its stored type.
	//
	// Any of "currency", "quantity", "time", "mass", "volume", "length",
	// "temperature", "area".
	Type UpsertUnitGroupInputType `json:"type,omitzero" api:"required"`
	// Free-form notes about the unit group. Preserved when omitted on update.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Units to associate with the group. Replaces the existing set on update; the base
	// unit is always kept.
	UnitConversions []UpsertUnitGroupConversionInputParam `json:"unit_conversions,omitzero"`
	paramObj
}

func (r UpsertUnitGroupInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertUnitGroupInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertUnitGroupInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unit dimension type. Create-only — an existing group keeps its stored type.
type UpsertUnitGroupInputType string

const (
	UpsertUnitGroupInputTypeCurrency    UpsertUnitGroupInputType = "currency"
	UpsertUnitGroupInputTypeQuantity    UpsertUnitGroupInputType = "quantity"
	UpsertUnitGroupInputTypeTime        UpsertUnitGroupInputType = "time"
	UpsertUnitGroupInputTypeMass        UpsertUnitGroupInputType = "mass"
	UpsertUnitGroupInputTypeVolume      UpsertUnitGroupInputType = "volume"
	UpsertUnitGroupInputTypeLength      UpsertUnitGroupInputType = "length"
	UpsertUnitGroupInputTypeTemperature UpsertUnitGroupInputType = "temperature"
	UpsertUnitGroupInputTypeArea        UpsertUnitGroupInputType = "area"
)

type CatalogUnitGroupActionBulkUpsertParams struct {
	// BulkUpsertUnitGroupsRequest is the request to bulk upsert unit groups.
	BulkUpsertUnitGroupsRequest BulkUpsertUnitGroupsRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogUnitGroupActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertUnitGroupsRequest)
}
func (r *CatalogUnitGroupActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogUnitGroupActionBulkUpsertParams]'s query parameters
// as `url.Values`.
func (r CatalogUnitGroupActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
