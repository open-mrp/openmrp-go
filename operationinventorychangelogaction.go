// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and export inventory change logs.
//
// OperationInventoryChangeLogActionService contains methods and other services
// that help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationInventoryChangeLogActionService] method instead.
type OperationInventoryChangeLogActionService struct {
	options []option.RequestOption
}

// NewOperationInventoryChangeLogActionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationInventoryChangeLogActionService(opts ...option.RequestOption) (r OperationInventoryChangeLogActionService) {
	r = OperationInventoryChangeLogActionService{}
	r.options = opts
	return
}

// Exports inventory change logs matching the provided filters as an Excel file.
//
// Unlike the list endpoint, results are not paginated — every matching change log
// is included in the download, newest first. The download is named for the date
// range you requested, using `all` in place of a bound you left open.
//
// This endpoint requires the permission: `inventory_logs:read`.
func (r *OperationInventoryChangeLogActionService) Export(ctx context.Context, query OperationInventoryChangeLogActionExportParams, opts ...option.RequestOption) (res *FileDownload, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/operations/inventory-change-logs/actions/export"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// FileDownload is a response type for endpoints that return a file (e.g. Excel
// export). When the service returns \*FileDownload, the handler writes the body
// with Content-Type and Content-Disposition.
type FileDownload struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDownload) RawJSON() string { return r.JSON.raw }
func (r *FileDownload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OperationInventoryChangeLogActionExportParams struct {
	// Restricts results to change logs created on or before this timestamp.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Restricts results to change logs created on or after this timestamp.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Restricts results to these action types.
	//
	// Any of "scan", "user_action", "system_action", "user_correction".
	ActionTypes []string `query:"action_types,omitzero" json:"-"`
	// Restricts results to changes made by these users.
	//
	// Changes that were recorded without a responsible user are excluded whenever this
	// filter is set.
	ChangedByUserIDs []string `query:"changed_by_user_ids,omitzero" json:"-"`
	// Restricts results to changes affecting these items.
	ItemIDs []string `query:"item_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OperationInventoryChangeLogActionExportParams]'s query
// parameters as `url.Values`.
func (r OperationInventoryChangeLogActionExportParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
