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

// List and manage account users.
//
// IdentityAccountUserActionService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityAccountUserActionService] method instead.
type IdentityAccountUserActionService struct {
	options []option.RequestOption
}

// NewIdentityAccountUserActionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIdentityAccountUserActionService(opts ...option.RequestOption) (r IdentityAccountUserActionService) {
	r = IdentityAccountUserActionService{}
	r.options = opts
	return
}

// Activates a disabled or removed account user, restoring their access to the
// account you are acting in.
//
// Reactivating a user in your own account consumes a seat, so the request fails if
// your plan is at its seat limit; users of a customer or supplier account you
// manage take no seat. Activating an already-active user is a no-op.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:update`.
func (r *IdentityAccountUserActionService) Activate(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUserActionActivateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s/actions/activate", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Disables (locks) an account user.
//
// Disabled users cannot access the account and their active sessions are revoked,
// but the membership and its role assignment are kept so access can be restored
// with the activate action. Disabling frees the seat the user occupied. Admin
// users cannot be disabled, you cannot disable yourself, and removed users must be
// activated before they can be disabled. Disabling an already-disabled user is a
// no-op.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:update`.
func (r *IdentityAccountUserActionService) Disable(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUserActionDisableResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s/actions/disable", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Removes a user from the account you are acting in.
//
// Removal is a soft delete: removed users are excluded from listings unless
// requested via `removed_scope`, they free the seat they occupied, and they can be
// restored with the activate action or by adding them again. Removing an
// already-removed user is a no-op. The user's profile itself is untouched, so
// their access to any other account they belong to is unaffected.
//
// Acting in a customer's account requires `customers:delete`, and acting in a
// supplier's account requires `suppliers:delete`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:delete`.
func (r *IdentityAccountUserActionService) Remove(ctx context.Context, id string, opts ...option.RequestOption) (res *IdentityAccountUserActionRemoveResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s/actions/remove", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type IdentityAccountUserActionActivateResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUserActionActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUserActionActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityAccountUserActionDisableResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUserActionDisableResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUserActionDisableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityAccountUserActionRemoveResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IdentityAccountUserActionRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *IdentityAccountUserActionRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
