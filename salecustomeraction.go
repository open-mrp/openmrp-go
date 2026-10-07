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

// Manage customer accounts.
//
// SaleCustomerActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSaleCustomerActionService] method instead.
type SaleCustomerActionService struct {
	options []option.RequestOption
}

// NewSaleCustomerActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSaleCustomerActionService(opts ...option.RequestOption) (r SaleCustomerActionService) {
	r = SaleCustomerActionService{}
	r.options = opts
	return
}

// Merges one or more source customers into a target customer.
//
// Sales orders, invoices, shipments, deliveries, and other transaction records
// from the source customers are reassigned to the target; price groups, product
// line access, addresses, and users are consolidated without duplicates; child
// accounts of the sources are re-parented to the target; the source customers are
// then deleted.
//
// The target keeps its own name, number, default addresses, and default settings —
// none of those are copied over from the sources, and the sources' notification
// recipients are discarded rather than transferred.
//
// This endpoint requires the permissions: `customers:update` and
// `customers:delete`.
func (r *SaleCustomerActionService) Merge(ctx context.Context, id string, params SaleCustomerActionMergeParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/sales/customers/%s/actions/merge", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to merge source customers into a target customer.
//
// The property SourceCustomerIDs is required.
type MergeCustomersRequestParam struct {
	// IDs of the source customers to merge into the target.
	//
	// Sources are deleted after the merge. The list must not contain duplicates or the
	// target customer's ID.
	SourceCustomerIDs []string `json:"source_customer_ids,omitzero" api:"required"`
	paramObj
}

func (r MergeCustomersRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MergeCustomersRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MergeCustomersRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SaleCustomerActionMergeParams struct {
	// Request to merge source customers into a target customer.
	MergeCustomersRequest MergeCustomersRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "bill_to_address", "ship_to_address", "type", "parent_account",
	// "freight_preferences.carrier", "freight_preferences.carrier.service_levels",
	// "freight_preferences.service_level", "defaults.payment_term",
	// "defaults.shipping_term", "defaults.sales_rep", "defaults.sales_rep.user",
	// "defaults.priority", "contact_info", "freight_preferences", "defaults",
	// "notification_preferences", "price_groups", "child_accounts", "credit_limit",
	// "credit_limit.unit".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r SaleCustomerActionMergeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MergeCustomersRequest)
}
func (r *SaleCustomerActionMergeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SaleCustomerActionMergeParams]'s query parameters as
// `url.Values`.
func (r SaleCustomerActionMergeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
