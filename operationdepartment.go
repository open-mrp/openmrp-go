// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage departments.
//
// OperationDepartmentService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationDepartmentService] method instead.
type OperationDepartmentService struct {
	options []option.RequestOption
}

// NewOperationDepartmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationDepartmentService(opts ...option.RequestOption) (r OperationDepartmentService) {
	r = OperationDepartmentService{}
	r.options = opts
	return
}

// Creates a department, optionally assigning scanning stations and machines to it.
//
// Returns a conflict error if a department with the same name already exists.
//
// This endpoint requires the permission: `departments:create`.
func (r *OperationDepartmentService) New(ctx context.Context, params OperationDepartmentNewParams, opts ...option.RequestOption) (res *Department, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/departments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a department by ID.
//
// This endpoint requires the permission: `departments:read`.
func (r *OperationDepartmentService) Get(ctx context.Context, id string, query OperationDepartmentGetParams, opts ...option.RequestOption) (res *Department, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/departments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a department.
//
// Only the fields provided in the request are changed. Assigning scanning stations
// or machines is additive and does not remove existing ones. Returns a conflict
// error if the new name is already in use by another department.
//
// This endpoint requires the permission: `departments:update`.
func (r *OperationDepartmentService) Update(ctx context.Context, id string, params OperationDepartmentUpdateParams, opts ...option.RequestOption) (res *Department, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/departments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of departments in your account, most recently created
// first.
//
// The `q` search term matches the department name.
//
// This endpoint requires the permission: `departments:read`.
func (r *OperationDepartmentService) List(ctx context.Context, query OperationDepartmentListParams, opts ...option.RequestOption) (res *ListDepartment, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/departments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a department.
//
// Scanning stations and machines assigned to the department are not deleted, but
// they keep pointing at it, and a machine whose department is gone can no longer
// be read, updated, or deleted through the machines endpoints. Reassign both to
// another department before deleting this one. Deleting a department that was
// already deleted returns an already-deleted error rather than a not-found error.
//
// This endpoint requires the permission: `departments:delete`.
func (r *OperationDepartmentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationDepartmentDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/departments/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a department.
//
// The property Name is required.
type CreateDepartmentRequestParam struct {
	// Display name of the department.
	//
	// Must be unique within your account; maximum 255 characters.
	Name string `json:"name" api:"required"`
	// ID of the location where this department operates.
	LocationID param.Opt[string] `json:"location_id,omitzero"`
	// Free-form notes about the department.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// A rate, expressed as a value together with the units of its numerator and
	// denominator (for example, `25.00` `$` per `hr`).
	LaborRate DepartmentRateInputParam `json:"labor_rate,omitzero"`
	// IDs of machines to assign to this department.
	//
	// A machine belongs to one department at a time, so listed machines are moved out
	// of their current department.
	MachineIDs []string `json:"machine_ids,omitzero"`
	// IDs of scanning stations to assign to this department.
	//
	// A scanning station belongs to one department at a time, so listed stations are
	// moved out of their current department.
	ScanningStationIDs []string `json:"scanning_station_ids,omitzero"`
	paramObj
}

func (r CreateDepartmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateDepartmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateDepartmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rate, expressed as a value together with the units of its numerator and
// denominator (for example, `25.00` `$` per `hr`).
//
// The properties DenominatorUnitID, NumeratorUnitID, Value are required.
type DepartmentRateInputParam struct {
	// ID of the unit in the rate's denominator (e.g. hours).
	DenominatorUnitID string `json:"denominator_unit_id" api:"required"`
	// ID of the unit in the rate's numerator (a currency, e.g. dollars).
	NumeratorUnitID string `json:"numerator_unit_id" api:"required"`
	// Decimal value of the rate.
	Value string `json:"value" api:"required"`
	paramObj
}

func (r DepartmentRateInputParam) MarshalJSON() (data []byte, err error) {
	type shadow DepartmentRateInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DepartmentRateInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListDepartment struct {
	// Resources in this page.
	Data []Department `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListDepartmentObject `json:"object" api:"required"`
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
func (r ListDepartment) RawJSON() string { return r.JSON.raw }
func (r *ListDepartment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListDepartmentObject string

const (
	ListDepartmentObjectList ListDepartmentObject = "list"
)

// Request to partially update a department.
type UpdateDepartmentRequestParam struct {
	// Free-form notes about the department.
	//
	// Send `null` to clear.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// ID of the location where this department operates.
	LocationID param.Opt[string] `json:"location_id,omitzero"`
	// Display name of the department.
	//
	// Must be unique within your account; maximum 255 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// A rate, expressed as a value together with the units of its numerator and
	// denominator (for example, `25.00` `$` per `hr`).
	LaborRate DepartmentRateInputParam `json:"labor_rate,omitzero"`
	// IDs of machines to assign to this department.
	//
	// Assignment is additive: listed machines are moved into this department and
	// machines already in the department are unaffected.
	MachineIDs []string `json:"machine_ids,omitzero"`
	// IDs of scanning stations to assign to this department.
	//
	// Assignment is additive: listed stations are moved into this department and
	// stations already in the department are unaffected.
	ScanningStationIDs []string `json:"scanning_station_ids,omitzero"`
	paramObj
}

func (r UpdateDepartmentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateDepartmentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateDepartmentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationDepartmentDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationDepartmentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationDepartmentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationDepartmentNewParams struct {
	// Request to create a department.
	CreateDepartmentRequest CreateDepartmentRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "location", "scanning_stations", "machines".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationDepartmentNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateDepartmentRequest)
}
func (r *OperationDepartmentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationDepartmentNewParams]'s query parameters as
// `url.Values`.
func (r OperationDepartmentNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationDepartmentGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "location", "scanning_stations", "machines".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationDepartmentGetParams]'s query parameters as
// `url.Values`.
func (r OperationDepartmentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationDepartmentUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "location", "scanning_stations", "machines".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a department.
	UpdateDepartmentRequest UpdateDepartmentRequestParam
	paramObj
}

func (r OperationDepartmentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateDepartmentRequest)
}
func (r *OperationDepartmentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationDepartmentUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationDepartmentUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationDepartmentListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "location", "scanning_stations", "machines".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationDepartmentListParams]'s query parameters as
// `url.Values`.
func (r OperationDepartmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
