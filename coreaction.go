// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Utility action endpoints for checking duplicates and emailing records.
//
// CoreActionService contains methods and other services that help with interacting
// with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreActionService] method instead.
type CoreActionService struct {
	options []option.RequestOption
}

// NewCoreActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreActionService(opts ...option.RequestOption) (r CoreActionService) {
	r = CoreActionService{}
	r.options = opts
	return
}

// Emails a record (invoice, sales order, or purchase order) to its configured
// recipients and marks the record as sent.
//
// Delivery is asynchronous: the endpoint returns `202 Accepted` once the email is
// queued, so a `202` means the send was accepted, not that it reached the
// recipients. If the record has no configured recipients the request still
// succeeds and nothing is sent; in that case a sales order or purchase order is
// also left unmarked, while an invoice is still marked as sent.
//
// An `id` that is not one of the account's records of the given `type` answers
// `404`.
//
// This endpoint requires the permissions: `invoices:read`, `sales_orders:read`,
// `purchase_orders:read`.
func (r *CoreActionService) EmailRecord(ctx context.Context, body CoreActionEmailRecordParams, opts ...option.RequestOption) (res *CoreActionEmailRecordResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/actions/email-record"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request to email a record to its configured recipients.
//
// The properties ID, Type are required.
type EmailRecordRequestParam struct {
	// ID of the record to email.
	ID string `json:"id" api:"required"`
	// The type of record to email.
	//
	//   - `invoice`: emails the invoice to the contacts on its sales order that are set
	//     to receive invoice emails.
	//   - `sales_order`: sends an order acknowledgement to the order's acknowledgement
	//     recipients.
	//   - `purchase_order`: sends the purchase order submission to the order's
	//     submission recipients.
	//
	// Any of "invoice", "sales_order", "purchase_order".
	Type EmailRecordRequestType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r EmailRecordRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EmailRecordRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailRecordRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of record to email.
//
//   - `invoice`: emails the invoice to the contacts on its sales order that are set
//     to receive invoice emails.
//   - `sales_order`: sends an order acknowledgement to the order's acknowledgement
//     recipients.
//   - `purchase_order`: sends the purchase order submission to the order's
//     submission recipients.
type EmailRecordRequestType string

const (
	EmailRecordRequestTypeInvoice       EmailRecordRequestType = "invoice"
	EmailRecordRequestTypeSalesOrder    EmailRecordRequestType = "sales_order"
	EmailRecordRequestTypePurchaseOrder EmailRecordRequestType = "purchase_order"
)

type CoreActionEmailRecordResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CoreActionEmailRecordResponse) RawJSON() string { return r.JSON.raw }
func (r *CoreActionEmailRecordResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CoreActionEmailRecordParams struct {
	// Request to email a record to its configured recipients.
	EmailRecordRequest EmailRecordRequestParam
	paramObj
}

func (r CoreActionEmailRecordParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmailRecordRequest)
}
func (r *CoreActionEmailRecordParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
