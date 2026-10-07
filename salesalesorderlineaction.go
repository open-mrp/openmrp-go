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
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List, view, create, update, and delete sales orders.
//
// SaleSalesOrderLineActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderLineActionService] method instead.
type SaleSalesOrderLineActionService struct {
	options []option.RequestOption
}

// NewSaleSalesOrderLineActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSaleSalesOrderLineActionService(opts ...option.RequestOption) (r SaleSalesOrderLineActionService) {
	r = SaleSalesOrderLineActionService{}
	r.options = opts
	return
}

// Reorders the product lines on a sales order to match the sequence supplied.
//
// The lines are renumbered from `1` in the given order. Discount and freight lines
// always stay at the bottom of the list regardless of the sequence given here.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderLineActionService) Reorder(ctx context.Context, id string, body SaleSalesOrderLineActionReorderParams, opts ...option.RequestOption) (res *SaleSalesOrderLineActionReorderResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines/actions/reorder", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to reorder a sales order's lines.
//
// The property LineIDs is required.
type ReorderSalesOrderLinesRequestParam struct {
	// The order's product-line IDs in the desired display order.
	//
	// Every product line on the order must be listed exactly once. The automatically
	// generated discount and freight lines are kept at the bottom of the list and must
	// not be included.
	LineIDs []string `json:"line_ids,omitzero" api:"required"`
	paramObj
}

func (r ReorderSalesOrderLinesRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ReorderSalesOrderLinesRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReorderSalesOrderLinesRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderLineActionReorderResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderLineActionReorderResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderLineActionReorderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderLineActionReorderParams struct {
	// Request to reorder a sales order's lines.
	ReorderSalesOrderLinesRequest ReorderSalesOrderLinesRequestParam
	paramObj
}

func (r SaleSalesOrderLineActionReorderParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ReorderSalesOrderLinesRequest)
}
func (r *SaleSalesOrderLineActionReorderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
