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

// List and manage machines.
//
// OperationMachineService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationMachineService] method instead.
type OperationMachineService struct {
	options []option.RequestOption
}

// NewOperationMachineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationMachineService(opts ...option.RequestOption) (r OperationMachineService) {
	r = OperationMachineService{}
	r.options = opts
	return
}

// Creates a machine and assigns it to a department.
//
// Returns a conflict error if another machine in your account already uses the
// same name, and a not-found error if the department does not belong to your
// account. The department cannot be changed once the machine exists.
//
// This endpoint requires the permission: `machines:create`.
func (r *OperationMachineService) New(ctx context.Context, params OperationMachineNewParams, opts ...option.RequestOption) (res *Machine, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/machines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a machine by ID.
//
// This endpoint requires the permission: `machines:read`.
func (r *OperationMachineService) Get(ctx context.Context, id string, query OperationMachineGetParams, opts ...option.RequestOption) (res *Machine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/machines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates a machine.
//
// Only the fields provided in the request are changed. Returns a conflict error if
// the new name is already in use by another machine in your account. A machine
// cannot be moved to a different department.
//
// This endpoint requires the permission: `machines:update`.
func (r *OperationMachineService) Update(ctx context.Context, id string, params OperationMachineUpdateParams, opts ...option.RequestOption) (res *Machine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/machines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of machines in your account, most recently created
// first.
//
// The search term matches the machine name.
//
// This endpoint requires the permission: `machines:read`.
func (r *OperationMachineService) List(ctx context.Context, query OperationMachineListParams, opts ...option.RequestOption) (res *ListMachine, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/machines"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a machine.
//
// Deletion is permanent, and repeating the call reports that the machine has
// already been deleted. Downtime events and schedule lines already logged against
// the machine are kept rather than removed with it.
//
// This endpoint requires the permission: `machines:delete`.
func (r *OperationMachineService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OperationMachineDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/machines/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a machine.
//
// The properties DepartmentID, Name, SerialNumber are required.
type CreateMachineRequestParam struct {
	// ID of the department this machine belongs to.
	//
	// Must reference a department in your account.
	DepartmentID string `json:"department_id" api:"required"`
	// Display name of the machine.
	//
	// Must be unique within your account; maximum 255 characters.
	Name string `json:"name" api:"required"`
	// Serial number of the machine.
	//
	// Maximum 255 characters.
	SerialNumber string `json:"serial_number" api:"required"`
	// Free-form notes about the machine.
	Notes param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r CreateMachineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMachineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMachineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to partially update a machine.
type UpdateMachineRequestParam struct {
	// Free-form notes about the machine.
	//
	// Send `null` to clear.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Display name of the machine.
	//
	// Must be unique within your account; maximum 255 characters.
	Name param.Opt[string] `json:"name,omitzero"`
	// Serial number of the machine.
	//
	// Maximum 255 characters.
	SerialNumber param.Opt[string] `json:"serial_number,omitzero"`
	paramObj
}

func (r UpdateMachineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateMachineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateMachineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationMachineDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OperationMachineDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OperationMachineDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationMachineNewParams struct {
	// Request to create a machine.
	CreateMachineRequest CreateMachineRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationMachineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMachineRequest)
}
func (r *OperationMachineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationMachineNewParams]'s query parameters as
// `url.Values`.
func (r OperationMachineNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationMachineGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationMachineGetParams]'s query parameters as
// `url.Values`.
func (r OperationMachineGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationMachineUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update a machine.
	UpdateMachineRequest UpdateMachineRequestParam
	paramObj
}

func (r OperationMachineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateMachineRequest)
}
func (r *OperationMachineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationMachineUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationMachineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OperationMachineListParams struct {
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
	// Any of "department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationMachineListParams]'s query parameters as
// `url.Values`.
func (r OperationMachineListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
