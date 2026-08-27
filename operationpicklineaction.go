// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
)

// List, view, pick, void, and pack picks and pick lines.
//
// OperationPickLineActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationPickLineActionService] method instead.
type OperationPickLineActionService struct {
	options []option.RequestOption
}

// NewOperationPickLineActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOperationPickLineActionService(opts ...option.RequestOption) (r OperationPickLineActionService) {
	r = OperationPickLineActionService{}
	r.options = opts
	return
}

// Marks a pick line as fully picked.
//
// Sets the line's picked quantity to its sales order line's ordered quantity less
// everything already picked for that order line, including whatever this line had
// picked before the call. To record a short pick instead, set the quantity
// yourself with Update Pick Line. Has no effect on a line that has already been
// packed.
//
// This endpoint requires the permission: `picks:update`.
func (r *OperationPickLineActionService) Pick(ctx context.Context, id string, body OperationPickLineActionPickParams, opts ...option.RequestOption) (res *PickLine, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PickID == "" {
		err = errors.New("missing required pick_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s/lines/%s/actions/pick", url.PathEscape(body.PickID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Voids a pick line, undoing the picking work recorded on it.
//
// Resets the line's picked quantity to zero without deleting the line, so the
// quantity can be picked again. Returns a validation error if the line has already
// been packed.
//
// This endpoint requires the permission: `picks:update`.
func (r *OperationPickLineActionService) Void(ctx context.Context, id string, body OperationPickLineActionVoidParams, opts ...option.RequestOption) (res *PickLine, err error) {
	opts = slices.Concat(r.options, opts)
	if body.PickID == "" {
		err = errors.New("missing required pick_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/operations/picks/%s/lines/%s/actions/void", url.PathEscape(body.PickID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type OperationPickLineActionPickParams struct {
	PickID string `path:"pick_id" api:"required" json:"-"`
	paramObj
}

type OperationPickLineActionVoidParams struct {
	PickID string `path:"pick_id" api:"required" json:"-"`
	paramObj
}
