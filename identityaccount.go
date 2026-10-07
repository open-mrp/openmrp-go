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
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Manage account details, branding, portal, logo, and favicon.
//
// IdentityAccountService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityAccountService] method instead.
type IdentityAccountService struct {
	options []option.RequestOption
}

// NewIdentityAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIdentityAccountService(opts ...option.RequestOption) (r IdentityAccountService) {
	r = IdentityAccountService{}
	r.options = opts
	return
}

// Uploads a customer-portal favicon.
//
// Send the image as the raw request body, not as multipart form data. Use a small
// square PNG (e.g. 32x32 or 64x64) for the best result in browser tabs. The
// uploaded image replaces any existing favicon and is shown on the account's
// customer portal. You can only upload a favicon for the account you are acting
// in.
//
// A body that is empty or not a PNG, JPEG, GIF, WebP, or ICO image is refused with
// a 400, and one over 10 MB with a 413; the existing favicon is kept either way.
//
// This endpoint requires the permission: `self:update`.
func (r *IdentityAccountService) UpdateFavicon(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUpdateFaviconResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/accounts/%s/favicon", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type IdentityAccountUpdateFaviconResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUpdateFaviconResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUpdateFaviconResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
