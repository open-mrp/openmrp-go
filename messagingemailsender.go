// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Choose the address your order, invoice, and statement emails are sent from, on a
// domain you have verified.
//
// MessagingEmailSenderService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingEmailSenderService] method instead.
type MessagingEmailSenderService struct {
	options []option.RequestOption
}

// NewMessagingEmailSenderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingEmailSenderService(opts ...option.RequestOption) (r MessagingEmailSenderService) {
	r = MessagingEmailSenderService{}
	r.options = opts
	return
}

// Sets the address your order, invoice, and statement emails are sent from,
// replacing any address already configured.
//
// The domain must be verified first. Emails about someone's OpenMRP account —
// password resets, verification, plan changes — continue to send from the platform
// address.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingEmailSenderService) Update(ctx context.Context, body MessagingEmailSenderUpdateParams, opts ...option.RequestOption) (res *EmailSender, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-sender"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns the address your order, invoice, and statement emails are sent from, or
// 404 when none is configured and that mail sends from the platform address.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailSenderService) List(ctx context.Context, opts ...option.RequestOption) (res *EmailSender, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-sender"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Clears the configured sending address, returning your customer-facing email to
// the platform address.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingEmailSenderService) Delete(ctx context.Context, opts ...option.RequestOption) (res *MessagingEmailSenderDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-sender"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The address your order, invoice, and statement emails are sent from.
//
// Configure one on a verified email domain and your customers see mail from your
// own address instead of the platform's. Emails about someone's OpenMRP account —
// password resets, verification, plan changes — always send from the platform
// address regardless of this setting.
//
// Mail only sends from this address while the underlying domain stays verified; if
// verification lapses it falls back to the platform address rather than failing to
// send.
type EmailSender struct {
	// Email sender ID.
	ID string `json:"id" api:"required"`
	// The full sending address.
	Address string `json:"address" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The domain name the address sends from.
	Domain string `json:"domain" api:"required"`
	// Verification status of the underlying domain.
	//
	// Any of "pending", "verified", "failed".
	DomainStatus EmailSenderDomainStatus `json:"domain_status" api:"required"`
	// The verified email domain this address belongs to.
	EmailDomainID string `json:"email_domain_id" api:"required"`
	// The name shown in a mail client's sender column. When unset, mail shows the bare
	// address.
	FromName string `json:"from_name" api:"required"`
	// The mailbox name before the `@`.
	LocalPart string `json:"local_part" api:"required"`
	// Resource type identifier.
	//
	// Any of "email_sender".
	Object EmailSenderObject `json:"object" api:"required"`
	// Where customer replies are delivered. When unset, replies go to the sending
	// address.
	ReplyTo string `json:"reply_to" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Address       respjson.Field
		CreatedAt     respjson.Field
		Domain        respjson.Field
		DomainStatus  respjson.Field
		EmailDomainID respjson.Field
		FromName      respjson.Field
		LocalPart     respjson.Field
		Object        respjson.Field
		ReplyTo       respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailSender) RawJSON() string { return r.JSON.raw }
func (r *EmailSender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Verification status of the underlying domain.
type EmailSenderDomainStatus string

const (
	EmailSenderDomainStatusPending  EmailSenderDomainStatus = "pending"
	EmailSenderDomainStatusVerified EmailSenderDomainStatus = "verified"
	EmailSenderDomainStatusFailed   EmailSenderDomainStatus = "failed"
)

// Resource type identifier.
type EmailSenderObject string

const (
	EmailSenderObjectEmailSender EmailSenderObject = "email_sender"
)

// Request to configure the address the account's customer-facing email is sent
// from.
//
// The properties EmailDomainID, LocalPart are required.
type SetEmailSenderRequestParam struct {
	// The verified email domain to send from.
	EmailDomainID string `json:"email_domain_id" api:"required"`
	// The mailbox name before the `@`, for example `orders`.
	LocalPart string `json:"local_part" api:"required"`
	// The name shown in a mail client's sender column. When unset, mail shows the bare
	// address.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// Where customer replies are delivered. When unset, replies go to the sending
	// address.
	ReplyTo param.Opt[string] `json:"reply_to,omitzero"`
	paramObj
}

func (r SetEmailSenderRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SetEmailSenderRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetEmailSenderRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailSenderDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingEmailSenderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingEmailSenderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailSenderUpdateParams struct {
	// Request to configure the address the account's customer-facing email is sent
	// from.
	SetEmailSenderRequest SetEmailSenderRequestParam
	paramObj
}

func (r MessagingEmailSenderUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SetEmailSenderRequest)
}
func (r *MessagingEmailSenderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
