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

// List, view, create, update, and delete sales orders.
//
// SaleSalesOrderLineService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleSalesOrderLineService] method instead.
type SaleSalesOrderLineService struct {
	options []option.RequestOption
	// List, view, create, update, and delete sales orders.
	Actions SaleSalesOrderLineActionService
}

// NewSaleSalesOrderLineService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleSalesOrderLineService(opts ...option.RequestOption) (r SaleSalesOrderLineService) {
	r = SaleSalesOrderLineService{}
	r.options = opts
	r.Actions = NewSaleSalesOrderLineActionService(opts...)
	return
}

// Adds a line item to a sales order.
//
// The new line is appended below the existing product lines, keeping the order's
// freight and discount lines at the bottom. When the order has already been
// issued, the line is added to its pick as outstanding work and the pick is
// reopened if it had been finished.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderLineService) New(ctx context.Context, id string, params SaleSalesOrderLineNewParams, opts ...option.RequestOption) (res *SalesOrderLine, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Partially updates a sales order line item.
//
// Changing the quantity flows through to fulfillment: the order's pick is
// reconciled against what is still outstanding — reopening it when the new
// quantity leaves work to do, or dropping the surplus pick line and finishing it
// when everything ordered is already packed. Shipment and invoice lines that still
// carry the full previously ordered quantity follow the new value, while partial
// ones keep the amount that actually moved.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderLineService) Update(ctx context.Context, lineID string, params SaleSalesOrderLineUpdateParams, opts ...option.RequestOption) (res *SalesOrderLine, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if lineID == "" {
		err = errors.New("missing required line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines/%s", url.PathEscape(params.ID), url.PathEscape(lineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Deletes a sales order line and its pick lines.
//
// A line cannot be removed once it has been packed onto a shipment, or once the
// order is fulfilled, and removing one from an order that is already completed or
// has a shipped shipment requires an admin. The remaining lines are renumbered so
// the sequence stays contiguous, and if this was the last line left to pick, the
// order's pick is deleted and the order falls back to `estimate` with its reserved
// inventory released.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `sales_orders:update`.
func (r *SaleSalesOrderLineService) Delete(ctx context.Context, lineID string, body SaleSalesOrderLineDeleteParams, opts ...option.RequestOption) (res *SaleSalesOrderLineDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if lineID == "" {
		err = errors.New("missing required line_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/sales-orders/%s/lines/%s", url.PathEscape(body.ID), url.PathEscape(lineID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a line on a sales order.
//
// The properties ProductID, ProductSKU, Quantity are required.
type CreateSalesOrderLineRequestParam struct {
	// ID of the product being ordered.
	ProductID string `json:"product_id" api:"required"`
	// The product SKU recorded on the line.
	ProductSKU string `json:"product_sku" api:"required"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	Quantity QuantityInputParam `json:"quantity,omitzero" api:"required"`
	// The product description recorded on the line.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r CreateSalesOrderLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSalesOrderLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSalesOrderLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to update a sales order line.
type UpdateSalesOrderLineRequestParam struct {
	// Description recorded on the line.
	ProductDescription param.Opt[string] `json:"product_description,omitzero"`
	// SKU recorded on the line.
	ProductSKU param.Opt[string] `json:"product_sku,omitzero"`
	// An amount together with the unit it is expressed in.
	//
	// The unit may be a currency, so money amounts such as a credit limit are written
	// the same way as physical amounts like weights or counts.
	Quantity QuantityInputParam `json:"quantity,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitCost RateInputParam `json:"unit_cost,omitzero"`
	// A value expressed as a ratio of two units, supplied on create and update
	// requests.
	//
	// A unit price, for example, has a currency as its numerator unit and the unit the
	// product is bought or sold by as its denominator.
	UnitPrice RateInputParam `json:"unit_price,omitzero"`
	paramObj
}

func (r UpdateSalesOrderLineRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateSalesOrderLineRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateSalesOrderLineRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderLineDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SaleSalesOrderLineDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SaleSalesOrderLineDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleSalesOrderLineNewParams struct {
	// Request to create a line on a sales order.
	CreateSalesOrderLineRequest CreateSalesOrderLineRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product", "quantity_ordered", "unit_price", "unit_cost", "totals".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleSalesOrderLineNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSalesOrderLineRequest)
}
func (r *SaleSalesOrderLineNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderLineNewParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderLineNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderLineUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "product", "quantity_ordered", "unit_price", "unit_cost", "totals".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to update a sales order line.
	UpdateSalesOrderLineRequest UpdateSalesOrderLineRequestParam
	paramObj
}

func (r SaleSalesOrderLineUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateSalesOrderLineRequest)
}
func (r *SaleSalesOrderLineUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleSalesOrderLineUpdateParams]'s query parameters as
// `url.Values`.
func (r SaleSalesOrderLineUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SaleSalesOrderLineDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
