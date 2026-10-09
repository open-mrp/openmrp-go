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

// List and manage locations.
//
// OperationLocationActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationLocationActionService] method instead.
type OperationLocationActionService struct {
	options []option.RequestOption
}

// NewOperationLocationActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationLocationActionService(opts ...option.RequestOption) (r OperationLocationActionService) {
	r = OperationLocationActionService{}
	r.options = opts
	return
}

// Creates or updates multiple locations for the account, matched by name
// (case-insensitive), then writes asynchronously — 202 with a job to poll.
//
// This endpoint requires the permissions: `locations:create` and
// `locations:update`.
func (r *OperationLocationActionService) BulkUpsert(ctx context.Context, params OperationLocationActionBulkUpsertParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/locations/actions/bulk-upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// BulkUpsertLocationsRequest is the request to bulk upsert locations.
//
// The property Locations is required.
type BulkUpsertLocationsRequestParam struct {
	// Locations to create or update, matched by name within the account.
	Locations []UpsertLocationInputParam `json:"locations,omitzero" api:"required"`
	paramObj
}

func (r BulkUpsertLocationsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkUpsertLocationsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkUpsertLocationsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpsertLocationInput is the input for a single location in a bulk upsert
// operation.
//
// The properties Name, Type are required.
type UpsertLocationInputParam struct {
	// Display name of the location, used to match existing locations.
	Name string `json:"name" api:"required"`
	// Location type code.
	//
	// Any of "building", "section", "aisle", "rack", "shelf", "bin".
	Type LocationTypeCode `json:"type,omitzero" api:"required"`
	// Child locations to re-parent under this one, referenced by `id` or `name`, or by
	// name for a location in the same batch. Redundant with `parent` on each child.
	Children []ObjectIdentifierParam `json:"children,omitzero"`
	// -------------------------- Named Object -------------------------- Identifies an
	// object by its id or its name. An id wins when both are given.
	Parent ObjectIdentifierParam `json:"parent,omitzero"`
	paramObj
}

func (r UpsertLocationInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertLocationInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertLocationInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationLocationActionBulkUpsertParams struct {
	// BulkUpsertLocationsRequest is the request to bulk upsert locations.
	BulkUpsertLocationsRequest BulkUpsertLocationsRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationLocationActionBulkUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkUpsertLocationsRequest)
}
func (r *OperationLocationActionBulkUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationLocationActionBulkUpsertParams]'s query parameters
// as `url.Values`.
func (r OperationLocationActionBulkUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
