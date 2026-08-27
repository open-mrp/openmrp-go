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
// OperationPickActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationPickActionService] method instead.
type OperationPickActionService struct {
	options []option.RequestOption
}

// NewOperationPickActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationPickActionService(opts ...option.RequestOption) (r OperationPickActionService) {
	r = OperationPickActionService{}
	r.options = opts
	return
}

// Packs a pick, creating a shipment from the picked lines.
//
// Returns `202 Accepted` with a job, because packing writes a shipment, one
// shipment line per packed pick line, and the requested shipping cases. Poll the
// job at the returned `Location`; once it reports `completed`, its first result
// carries the new shipment's `id`, with the shipment line and shipping case ids in
// `sub_resource_ids`. Every unpacked line with a picked quantity greater than zero
// is marked as packed and added to a new shipment in `packed` status, which
// inherits the sales order's carrier, service level, and shipping address. When a
// sales order line still has outstanding quantity afterward and no unpacked pick
// line is already open for it, a new zero-quantity pick line is created for the
// remainder, so packing a partial pick leaves the pick open for the next round.
// The pick is marked finished only once every one of its lines is packed.
//
// Returns a validation error if no line on the pick has a picked quantity greater
// than zero.
//
// This endpoint requires the permission: `picks:update`.
func (r *OperationPickActionService) Pack(ctx context.Context, id string, params OperationPickActionPackParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s/actions/pack", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Marks all lines on a pick as picked.
//
// Sets each unpacked line's picked quantity to the quantity still outstanding on
// its sales order line, after accounting for what other pick lines for that order
// line have already picked. Lines that have already been packed are unaffected.
// Use this to fill in a full pick in one call instead of picking each line
// individually; nothing is shipped until the pick is packed.
//
// This endpoint requires the permission: `picks:update`.
func (r *OperationPickActionService) Pick(ctx context.Context, id string, opts ...option.RequestOption) (res *Pick, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s/actions/pick", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Voids a pick, undoing all picking work recorded on it.
//
// Resets the picked quantity on every unpacked line to zero and clears the pick's
// `finished_at` timestamp, so the pick starts over as open with nothing picked.
// The pick itself is not deleted, and the sales order is unaffected.
//
// Returns a validation error if any shipment exists for the pick's sales order.
// Voiding those shipments is not enough — they must be deleted, since a voided
// shipment still exists.
//
// This endpoint requires the permission: `picks:update`.
func (r *OperationPickActionService) Void(ctx context.Context, id string, opts ...option.RequestOption) (res *Pick, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s/actions/void", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Request to pack a pick, creating a shipment from the picked lines.
//
// The property ShipmentCaseCount is required.
type PackPickRequestParam struct {
	// Number of shipping cases to create on the new shipment.
	//
	// Must be at least 1. Cases are numbered sequentially from the shipment number
	// (e.g. `SO-001-1`, `SO-001-2`), and each starts with zero freight weight and
	// freight cost for you to fill in later.
	ShipmentCaseCount int64 `json:"shipment_case_count" api:"required"`
	paramObj
}

func (r PackPickRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PackPickRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PackPickRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationPickActionPackParams struct {
	// Request to pack a pick, creating a shipment from the picked lines.
	PackPickRequest PackPickRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r OperationPickActionPackParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PackPickRequest)
}
func (r *OperationPickActionPackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OperationPickActionPackParams]'s query parameters as
// `url.Values`.
func (r OperationPickActionPackParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
