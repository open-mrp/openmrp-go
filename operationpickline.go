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
)

// List, view, pick, void, and pack picks and pick lines.
//
// OperationPickLineService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationPickLineService] method instead.
type OperationPickLineService struct {
	options []option.RequestOption
	// List, view, pick, void, and pack picks and pick lines.
	Actions OperationPickLineActionService
}

// NewOperationPickLineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationPickLineService(opts ...option.RequestOption) (r OperationPickLineService) {
	r = OperationPickLineService{}
	r.options = opts
	r.Actions = NewOperationPickLineActionService(opts...)
	return
}

// Updates a pick line's picked quantity.
//
// Use this to record a short or partial pick; Pick Pick Line fills in the full
// outstanding quantity instead.
//
// This endpoint requires the permission: `picks:update`.
func (r *OperationPickLineService) Update(ctx context.Context, id string, params OperationPickLineUpdateParams, opts ...option.RequestOption) (res *PickLine, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PickID == "" {
		err = errors.New("missing required pick_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s/lines/%s", url.PathEscape(params.PickID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Request to update a pick line's picked quantity.
type UpdatePickLineRequestParam struct {
	// New picked quantity for the line, as a decimal string read in the unit the sales
	// order line was sold in, stored as given and not capped at the ordered quantity.
	//
	// Must not be negative. Pulling more than was ordered is a real floor event and is
	// kept as recorded; pulling a negative amount is not.
	QuantityValue param.Opt[string] `json:"quantity_value,omitzero" format:"decimal"`
	paramObj
}

func (r UpdatePickLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdatePickLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdatePickLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationPickLineUpdateParams struct {
	PickID string `path:"pick_id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sales_order_line", "sales_order_line.product", "quantity",
	// "quantity.unit", "ordered_quantity", "ordered_quantity.unit".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a pick line's picked quantity.
	UpdatePickLineRequest UpdatePickLineRequestParam
	paramObj
}

func (r OperationPickLineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdatePickLineRequest)
}
func (r *OperationPickLineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationPickLineUpdateParams]'s query parameters as
// `url.Values`.
func (r OperationPickLineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
