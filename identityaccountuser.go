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

// List and manage account users.
//
// IdentityAccountUserService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityAccountUserService] method instead.
type IdentityAccountUserService struct {
	options []option.RequestOption
	// List and manage account users.
	Actions IdentityAccountUserActionService
}

// NewIdentityAccountUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIdentityAccountUserService(opts ...option.RequestOption) (r IdentityAccountUserService) {
	r = IdentityAccountUserService{}
	r.options = opts
	r.Actions = NewIdentityAccountUserActionService(opts...)
	return
}

// Adds a user to the account you are acting in.
//
// If no user with the given email or username exists, a new user is created; a
// user created with an email address is sent a welcome email containing a
// generated password and a link to sign in, unless they are being added to a
// supplier account, since suppliers have no portal to sign in to. A user added to
// a customer account signs in through your customer portal, on its verified custom
// domain when you have one. If a matching user already exists, that user is added
// to the account instead, and a user you previously removed is restored rather
// than duplicated. Adding a user to your own account consumes a seat and is
// rejected once your plan's seat limit is reached.
//
// You can add users to a customer or supplier account only while you manage it
// alone: an account that runs its own OpenMRP subscription, or that another
// account also has as a customer or supplier, manages its own users, and the
// request is refused.
//
// Acting in a customer's account requires `customers:create`, and acting in a
// supplier's account requires `suppliers:create`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:create`.
func (r *IdentityAccountUserService) New(ctx context.Context, params IdentityAccountUserNewParams, opts ...option.RequestOption) (res *AccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/account-users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns an account user by ID.
//
// The lookup is scoped to the account you are acting in, so an ID belonging to
// another account is reported as not found.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:read`.
func (r *IdentityAccountUserService) Get(ctx context.Context, id string, query IdentityAccountUserGetParams, opts ...option.RequestOption) (res *AccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Partially updates an account user.
//
// Omitted fields are left unchanged. Profile fields (`name`, `email`, `username`)
// update the underlying user, which is shared across every account the user
// belongs to, so the change is visible everywhere that person works.
//
// Acting in a customer's account requires `customers:update`, and acting in a
// supplier's account requires `suppliers:update`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:update`.
func (r *IdentityAccountUserService) Update(ctx context.Context, id string, params IdentityAccountUserUpdateParams, opts ...option.RequestOption) (res *AccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/identity/account-users/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns a paginated list of the users who belong to the account you are acting
// in.
//
// When the account you are acting in is a customer or supplier account you manage,
// this lists that account's users rather than your own team, each with the
// `notification_types` you send them.
//
// Acting in a customer's account requires `customers:read`, and acting in a
// supplier's account requires `suppliers:read`, instead of the permission this
// endpoint requires in your own account.
//
// This endpoint requires the permission: `team:read`.
func (r *IdentityAccountUserService) List(ctx context.Context, query IdentityAccountUserListParams, opts ...option.RequestOption) (res *ListAccountUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/identity/account-users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Request to create an account user.
type CreateAccountUserRequestParam struct {
	// ID of the department to assign to the user.
	//
	// The department must already exist in the account you are acting in.
	DepartmentID param.Opt[string] `json:"department_id,omitzero"`
	// User email address.
	//
	// Either `email` or `username` must be provided. If a user with this email already
	// exists, that user is added to the account instead of a new user being created,
	// and the request fails with a conflict if they are already an active member of
	// it.
	Email param.Opt[string] `json:"email,omitzero"`
	// Whether the user can be assigned as a sales representative on orders,
	// territories, and targets.
	//
	// Defaults to false. Forced true for the `sales_rep` role type and rejected for
	// scanner and agent roles.
	IsCommissionEligible param.Opt[bool] `json:"is_commission_eligible,omitzero"`
	// User display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Password for scanning station users.
	//
	// Required when creating a scanning station user (username without email) and
	// rejected for all other users, who instead receive a generated password in their
	// welcome email. Must be 8–72 characters and include an uppercase letter, a
	// lowercase letter, a number, and a special character.
	Password param.Opt[string] `json:"password,omitzero"`
	// ID of the role to assign to the user.
	//
	// The role you supply can be overridden: users added to a customer account always
	// receive the shared customer role so their portal capabilities stay
	// permission-driven, and scanning station users in any other account receive the
	// scanner role. Supplying a role whose type is `sales_rep` normalizes to the
	// account's canonical sales-rep role.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// Unique username.
	//
	// 3–255 characters; letters, numbers, underscores, and hyphens. Either `email` or
	// `username` must be provided. Providing a username without an email creates a
	// scanning station user.
	Username param.Opt[string] `json:"username,omitzero"`
	// Notification preference toggles for the new user.
	//
	// Only applies when adding a user to a customer or supplier account you manage;
	// ignored when adding a user to your own account. The enabled types are returned
	// in the user's `notification_types`. When the user is one you previously removed,
	// the toggles apply over the preferences they had before removal, and types you
	// leave out keep their previous state.
	Preferences []NotificationPreferenceItemParam `json:"preferences,omitzero"`
	paramObj
}

func (r CreateAccountUserRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAccountUserRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAccountUserRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAccountUser struct {
	// Resources in this page.
	Data []AccountUser `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountUserObject `json:"object" api:"required"`
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
func (r ListAccountUser) RawJSON() string { return r.JSON.raw }
func (r *ListAccountUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountUserObject string

const (
	ListAccountUserObjectList ListAccountUserObject = "list"
)

// NotificationPreferenceItem toggles a single account-relation notification type.
//
// The properties Enabled, NotificationType are required.
type NotificationPreferenceItemParam struct {
	// Whether this notification type is enabled for the account user.
	Enabled bool `json:"enabled" api:"required"`
	// Notification type.
	//
	// Any of "invoice", "order_acknowledgement", "purchase_order_submission".
	NotificationType NotificationPreferenceItemNotificationType `json:"notification_type,omitzero" api:"required"`
	paramObj
}

func (r NotificationPreferenceItemParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationPreferenceItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationPreferenceItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notification type.
type NotificationPreferenceItemNotificationType string

const (
	NotificationPreferenceItemNotificationTypeInvoice                 NotificationPreferenceItemNotificationType = "invoice"
	NotificationPreferenceItemNotificationTypeOrderAcknowledgement    NotificationPreferenceItemNotificationType = "order_acknowledgement"
	NotificationPreferenceItemNotificationTypePurchaseOrderSubmission NotificationPreferenceItemNotificationType = "purchase_order_submission"
)

// Request to partially update an account user.
type UpdateAccountUserRequestParam struct {
	// ID of the department to assign to the user.
	//
	// Set to `null` to clear the department. The department must already exist in the
	// account.
	DepartmentID param.Opt[string] `json:"department_id,omitzero"`
	// ID of the role to assign to the user.
	//
	// Set to `null` to clear the role.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// User email address.
	//
	// Must not already be in use by another user.
	Email param.Opt[string] `json:"email,omitzero"`
	// Whether the user can be assigned as a sales representative on orders,
	// territories, and targets.
	//
	// Forced true for the `sales_rep` role type and rejected for scanner and agent
	// roles. Cannot be turned off while the user stays on a `sales_rep` role.
	IsCommissionEligible param.Opt[bool] `json:"is_commission_eligible,omitzero"`
	// User display name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Unique username.
	//
	// 3–255 characters; letters, numbers, underscores, and hyphens. Must not already
	// be in use by another user.
	Username param.Opt[string] `json:"username,omitzero"`
	// Notification preference toggles to apply.
	//
	// Only allowed when updating a user in a customer or supplier account you manage;
	// rejected otherwise. Notification types omitted from the list are left unchanged,
	// and the result is returned in `notification_types`.
	Preferences []NotificationPreferenceItemParam `json:"preferences,omitzero"`
	paramObj
}

func (r UpdateAccountUserRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAccountUserRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAccountUserRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityAccountUserNewParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to create an account user.
	CreateAccountUserRequest CreateAccountUserRequestParam
	paramObj
}

func (r IdentityAccountUserNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAccountUserRequest)
}
func (r *IdentityAccountUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IdentityAccountUserNewParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityAccountUserGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IdentityAccountUserGetParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityAccountUserUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update an account user.
	UpdateAccountUserRequest UpdateAccountUserRequestParam
	paramObj
}

func (r IdentityAccountUserUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAccountUserRequest)
}
func (r *IdentityAccountUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [IdentityAccountUserUpdateParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IdentityAccountUserListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter by commission eligibility.
	//
	// Exact match on the column. Pass `true` to list users who can be assigned as
	// sales representatives, including dedicated `sales_rep` users.
	IsCommissionEligible param.Opt[bool] `query:"is_commission_eligible,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "user", "role", "department".
	Include []string `query:"include,omitzero" json:"-"`
	// Controls whether removed (soft-deleted) account users appear in the list.
	//
	// Removed users are left out unless you pass `included`, so a user removed with
	// the remove action disappears from the default listing.
	//
	// Any of "excluded", "included".
	RemovedScope IdentityAccountUserListParamsRemovedScope `query:"removed_scope,omitzero" json:"-"`
	// Filter by role type.
	//
	// - `admin`: account administrators.
	// - `user`: users with a custom role.
	// - `scanner`: scanning station users.
	// - `sales_rep`: sales representatives.
	// - `agent`: automated agents.
	//
	// Any of "admin", "user", "scanner", "sales_rep", "agent".
	RoleType IdentityAccountUserListParamsRoleType `query:"role_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IdentityAccountUserListParams]'s query parameters as
// `url.Values`.
func (r IdentityAccountUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Controls whether removed (soft-deleted) account users appear in the list.
//
// Removed users are left out unless you pass `included`, so a user removed with
// the remove action disappears from the default listing.
type IdentityAccountUserListParamsRemovedScope string

const (
	IdentityAccountUserListParamsRemovedScopeExcluded IdentityAccountUserListParamsRemovedScope = "excluded"
	IdentityAccountUserListParamsRemovedScopeIncluded IdentityAccountUserListParamsRemovedScope = "included"
)

// Filter by role type.
//
// - `admin`: account administrators.
// - `user`: users with a custom role.
// - `scanner`: scanning station users.
// - `sales_rep`: sales representatives.
// - `agent`: automated agents.
type IdentityAccountUserListParamsRoleType string

const (
	IdentityAccountUserListParamsRoleTypeAdmin    IdentityAccountUserListParamsRoleType = "admin"
	IdentityAccountUserListParamsRoleTypeUser     IdentityAccountUserListParamsRoleType = "user"
	IdentityAccountUserListParamsRoleTypeScanner  IdentityAccountUserListParamsRoleType = "scanner"
	IdentityAccountUserListParamsRoleTypeSalesRep IdentityAccountUserListParamsRoleType = "sales_rep"
	IdentityAccountUserListParamsRoleTypeAgent    IdentityAccountUserListParamsRoleType = "agent"
)
