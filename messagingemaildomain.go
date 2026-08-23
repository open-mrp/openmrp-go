// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Register customer-owned domains with the email bridge and verify them for
// sending and receiving mail.
//
// MessagingEmailDomainService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingEmailDomainService] method instead.
type MessagingEmailDomainService struct {
	options []option.RequestOption
	// Register customer-owned domains with the email bridge and verify them for
	// sending and receiving mail.
	Actions MessagingEmailDomainActionService
}

// NewMessagingEmailDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingEmailDomainService(opts ...option.RequestOption) (r MessagingEmailDomainService) {
	r = MessagingEmailDomainService{}
	r.options = opts
	r.Actions = NewMessagingEmailDomainActionService(opts...)
	return
}

// Registers a domain you own with the email bridge and returns the DKIM tokens to
// publish.
//
// The domain starts in `pending`. Publish each returned token as a CNAME record in
// the domain's DNS, then call the verify action to move it to `verified`; only
// then can inboxes be created on it.
//
// A domain can only be registered once across the platform, so registering one
// that is already in use returns a conflict error.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingEmailDomainService) New(ctx context.Context, body MessagingEmailDomainNewParams, opts ...option.RequestOption) (res *EmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a single email domain owned by the account.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailDomainService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the account's registered email domains.
//
// Every domain is returned in a single response; this list is not paginated.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailDomainService) List(ctx context.Context, opts ...option.RequestOption) (res *ListEmailDomain, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-domains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deregisters a domain from the email bridge and removes its sending identity from
// the mail provider.
//
// Delete the domain's inboxes first: while any inbox still exists on it, this
// returns a conflict error.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingEmailDomainService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingEmailDomainDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-domains/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to register a sending/receiving domain with the email bridge.
//
// The property Domain is required.
type CreateEmailDomainRequestParam struct {
	// The fully-qualified domain name to register (e.g. `support.acme.com`).
	//
	// Supply a bare domain, not an email address; the value is lowercased before it is
	// stored.
	Domain string `json:"domain" api:"required"`
	paramObj
}

func (r CreateEmailDomainRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateEmailDomainRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateEmailDomainRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A domain registered with the email bridge for sending and receiving mail.
//
// After registration the domain starts in `pending`; publish the returned DKIM
// records, then poll the verify action until it flips to `verified`.
type EmailDomain struct {
	// Email domain ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The DKIM tokens that must be published in your DNS before the domain can be
	// verified.
	//
	// Publish each token as a CNAME record on the domain, then call the verify action
	// to confirm them.
	DkimTokens []string `json:"dkim_tokens" api:"required"`
	// The fully-qualified domain name (e.g. `support.acme.com`).
	Domain string `json:"domain" api:"required"`
	// Resource type identifier.
	//
	// Any of "email_domain".
	Object EmailDomainObject `json:"object" api:"required"`
	// Verification status.
	//
	// - `pending`: registered and awaiting DKIM confirmation.
	// - `verified`: DKIM confirmed; the domain can send mail.
	// - `failed`: verification could not be completed.
	//
	// Inboxes can only be created on a `verified` domain.
	//
	// Any of "pending", "verified", "failed".
	Status EmailDomainStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// When the domain's DKIM verification was confirmed.
	VerifiedAt time.Time `json:"verified_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DkimTokens  respjson.Field
		Domain      respjson.Field
		Object      respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		VerifiedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailDomain) RawJSON() string { return r.JSON.raw }
func (r *EmailDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type EmailDomainObject string

const (
	EmailDomainObjectEmailDomain EmailDomainObject = "email_domain"
)

// Verification status.
//
// - `pending`: registered and awaiting DKIM confirmation.
// - `verified`: DKIM confirmed; the domain can send mail.
// - `failed`: verification could not be completed.
//
// Inboxes can only be created on a `verified` domain.
type EmailDomainStatus string

const (
	EmailDomainStatusPending  EmailDomainStatus = "pending"
	EmailDomainStatusVerified EmailDomainStatus = "verified"
	EmailDomainStatusFailed   EmailDomainStatus = "failed"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListEmailDomain struct {
	// Resources in this page.
	Data []EmailDomain `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListEmailDomainObject `json:"object" api:"required"`
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
func (r ListEmailDomain) RawJSON() string { return r.JSON.raw }
func (r *ListEmailDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListEmailDomainObject string

const (
	ListEmailDomainObjectList ListEmailDomainObject = "list"
)

type MessagingEmailDomainDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingEmailDomainDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingEmailDomainDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailDomainNewParams struct {
	// Request to register a sending/receiving domain with the email bridge.
	CreateEmailDomainRequest CreateEmailDomainRequestParam
	paramObj
}

func (r MessagingEmailDomainNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateEmailDomainRequest)
}
func (r *MessagingEmailDomainNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
