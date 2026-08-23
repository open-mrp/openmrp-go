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

// List and manage inventory items.
//
// CatalogItemActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemActionService] method instead.
type CatalogItemActionService struct {
	options []option.RequestOption
}

// NewCatalogItemActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCatalogItemActionService(opts ...option.RequestOption) (r CatalogItemActionService) {
	r = CatalogItemActionService{}
	r.options = opts
	return
}

// Reconciles inventory for multiple items by SKU in one call, the bulk equivalent
// of counting stock and correcting the books.
//
// `reconcile_type` controls whether each quantity is added to the item's current
// quantity (`addition`) or replaces it (`force`). The figure a `force` measures
// against is what is on hand net of demand nothing has covered, the same basis the
// single-item endpoint uses. The response reports each item as reconciled, skipped
// (e.g. unknown SKU), or errored (e.g. unknown unit), so a problem with one item
// does not fail the rest of the batch.
//
// Each correction is written to the item's inventory audit trail as a user
// correction, attributed to the caller.
//
// This endpoint requires the permission: `items:create`.
func (r *CatalogItemActionService) BulkReconcile(ctx context.Context, body CatalogItemActionBulkReconcileParams, opts ...option.RequestOption) (res *BulkReconcileItemsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/catalog/items/actions/bulk-reconcile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// One item to reconcile in a bulk reconcile request.
//
// The properties Quantity, SKU, Unit are required.
type BulkReconcileItemInputParam struct {
	// Quantity to apply, interpreted according to the request's `reconcile_type`.
	//
	// A decimal string rather than a number: a quantity that has been through a binary
	// float is not the quantity you sent.
	Quantity string `json:"quantity" api:"required" format:"decimal"`
	// SKU of the item to reconcile.
	//
	// Items whose SKU does not match an existing item are reported in the response's
	// `skipped_items` rather than failing the request.
	SKU string `json:"sku" api:"required"`
	// Abbreviation of a unit available to your account (e.g. `kg`).
	//
	// The unit is checked for existence only: the quantity is always recorded in the
	// item's own base unit, so send figures already expressed in that unit. Rows
	// naming an abbreviation that matches no built-in or account-defined unit are
	// reported in the response's `errors`.
	Unit string `json:"unit" api:"required"`
	paramObj
}

func (r BulkReconcileItemInputParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkReconcileItemInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkReconcileItemInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to reconcile inventory for many items at once.
//
// The properties Data, ReconcileType are required.
type BulkReconcileItemsRequestParam struct {
	// Items to reconcile.
	Data []BulkReconcileItemInputParam `json:"data,omitzero" api:"required"`
	// How each item's quantity is applied to its current quantity.
	//
	// - `addition`: adds the quantity to the item's current quantity.
	// - `force`: sets the item's current quantity to exactly the given quantity.
	//
	// Any of "addition", "force".
	ReconcileType BulkReconcileItemsRequestReconcileType `json:"reconcile_type,omitzero" api:"required"`
	paramObj
}

func (r BulkReconcileItemsRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BulkReconcileItemsRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkReconcileItemsRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How each item's quantity is applied to its current quantity.
//
// - `addition`: adds the quantity to the item's current quantity.
// - `force`: sets the item's current quantity to exactly the given quantity.
type BulkReconcileItemsRequestReconcileType string

const (
	BulkReconcileItemsRequestReconcileTypeAddition BulkReconcileItemsRequestReconcileType = "addition"
	BulkReconcileItemsRequestReconcileTypeForce    BulkReconcileItemsRequestReconcileType = "force"
)

// The outcome of a bulk inventory reconciliation, reported as three separate
// lists.
type BulkReconcileItemsResponse struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Errors ListReconcileErrorResult `json:"errors" api:"required"`
	// Resource type identifier.
	//
	// Any of "bulk_reconcile_items_response".
	Object BulkReconcileItemsResponseObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ReconciledItems ListReconciledItemResult `json:"reconciled_items" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	SkippedItems ListSkippedItemResult `json:"skipped_items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Errors          respjson.Field
		Object          respjson.Field
		ReconciledItems respjson.Field
		SkippedItems    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkReconcileItemsResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkReconcileItemsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type BulkReconcileItemsResponseObject string

const (
	BulkReconcileItemsResponseObjectBulkReconcileItemsResponse BulkReconcileItemsResponseObject = "bulk_reconcile_items_response"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListReconcileErrorResult struct {
	// Resources in this page.
	Data []ReconcileErrorResult `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListReconcileErrorResultObject `json:"object" api:"required"`
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
func (r ListReconcileErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ListReconcileErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListReconcileErrorResultObject string

const (
	ListReconcileErrorResultObjectList ListReconcileErrorResultObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListReconciledItemResult struct {
	// Resources in this page.
	Data []ReconciledItemResult `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListReconciledItemResultObject `json:"object" api:"required"`
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
func (r ListReconciledItemResult) RawJSON() string { return r.JSON.raw }
func (r *ListReconciledItemResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListReconciledItemResultObject string

const (
	ListReconciledItemResultObjectList ListReconciledItemResultObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListSkippedItemResult struct {
	// Resources in this page.
	Data []SkippedItemResult `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListSkippedItemResultObject `json:"object" api:"required"`
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
func (r ListSkippedItemResult) RawJSON() string { return r.JSON.raw }
func (r *ListSkippedItemResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListSkippedItemResultObject string

const (
	ListSkippedItemResultObjectList ListSkippedItemResultObject = "list"
)

// A submitted row that could not be reconciled.
type ReconcileErrorResult struct {
	// Error message.
	Error string `json:"error" api:"required"`
	// Item SKU.
	SKU string `json:"sku" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		SKU         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReconcileErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ReconcileErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An item whose on-hand quantity was successfully reconciled.
//
// Both quantities are expressed in the item's own base unit, not in the unit
// submitted with the request.
type ReconciledItemResult struct {
	// Item ID.
	ItemID string `json:"item_id" api:"required"`
	// Quantity after the reconciliation, as a decimal string.
	NewQuantity string `json:"new_quantity" api:"required" format:"decimal"`
	// Quantity before the reconciliation, as a decimal string.
	PreviousQuantity string `json:"previous_quantity" api:"required" format:"decimal"`
	// Item SKU.
	SKU string `json:"sku" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ItemID           respjson.Field
		NewQuantity      respjson.Field
		PreviousQuantity respjson.Field
		SKU              respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReconciledItemResult) RawJSON() string { return r.JSON.raw }
func (r *ReconciledItemResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A submitted row that was skipped rather than reconciled.
type SkippedItemResult struct {
	// Human-readable reason the item was skipped.
	Reason string `json:"reason" api:"required"`
	// Item SKU.
	SKU string `json:"sku" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		SKU         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkippedItemResult) RawJSON() string { return r.JSON.raw }
func (r *SkippedItemResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemActionBulkReconcileParams struct {
	// Request to reconcile inventory for many items at once.
	BulkReconcileItemsRequest BulkReconcileItemsRequestParam
	paramObj
}

func (r CatalogItemActionBulkReconcileParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BulkReconcileItemsRequest)
}
func (r *CatalogItemActionBulkReconcileParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
