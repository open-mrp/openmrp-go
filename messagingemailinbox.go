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
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Provision and manage routable email inboxes that bind inbound mail to chat
// conversations and send agent replies.
//
// MessagingEmailInboxService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingEmailInboxService] method instead.
type MessagingEmailInboxService struct {
	options []option.RequestOption
}

// NewMessagingEmailInboxService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingEmailInboxService(opts ...option.RequestOption) (r MessagingEmailInboxService) {
	r = MessagingEmailInboxService{}
	r.options = opts
	return
}

// Provisions a routable inbox address on a verified domain.
//
// Once created, mail arriving at the address opens a customer case conversation
// and seats the bound agent and the group's members on it; a reply in a thread
// that already opened one joins that conversation instead.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingEmailInboxService) New(ctx context.Context, params MessagingEmailInboxNewParams, opts ...option.RequestOption) (res *EmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single email inbox owned by the account.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailInboxService) Get(ctx context.Context, id string, query MessagingEmailInboxGetParams, opts ...option.RequestOption) (res *EmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-inboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Edits an email inbox's from-name, status, agent configuration, and roster.
//
// Every field except `status` is merged into the inbox's current settings: a field
// you omit — and an empty array you send — keeps the value it already has, so this
// endpoint can change a setting but cannot clear one back to unset. The inbox's
// address and domain are fixed at creation and cannot be changed here.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingEmailInboxService) Update(ctx context.Context, id string, params MessagingEmailInboxUpdateParams, opts ...option.RequestOption) (res *EmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-inboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns the account's email inboxes across every registered domain.
//
// Every inbox is returned in a single response; this list is not paginated.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingEmailInboxService) List(ctx context.Context, query MessagingEmailInboxListParams, opts ...option.RequestOption) (res *ListEmailInbox, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/email-inboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes an email inbox.
//
// Mail sent to its address is no longer routed. Conversations the inbox already
// opened are kept, but replies can no longer be sent on them, so disable the inbox
// instead of deleting it if you still need to answer open threads.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingEmailInboxService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingEmailInboxDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/email-inboxes/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to provision a routable inbox on a verified domain.
//
// The properties Address, EmailDomainID are required.
type CreateEmailInboxRequestParam struct {
	// The full inbox address (e.g. `support@acme.com`).
	//
	// Its domain part must match the selected domain, which must already be verified.
	// The address is lowercased before it is stored, and it must not already be in use
	// by another inbox.
	Address string `json:"address" api:"required"`
	// The verified domain this inbox belongs to.
	EmailDomainID string `json:"email_domain_id" api:"required"`
	// The agent to bind to this inbox to handle incoming mail.
	//
	// With no agent bound, mail is still threaded into a conversation for your team,
	// but nothing runs on it automatically.
	AgentConfigID param.Opt[string] `json:"agent_config_id,omitzero"`
	// Display name for the `From` header of outbound mail.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// The messaging group (roster) whose members are seated on every conversation this
	// inbox opens.
	//
	// Must name a group in your own account. Agents in the group are seated to run
	// only when @mentioned, so they do not all fire alongside the inbox's own agent.
	GroupID param.Opt[string] `json:"group_id,omitzero"`
	// The keywords that decide whether the agent runs on an incoming message.
	//
	// Under the `keyword` policy a keyword matches anywhere in the message; under
	// `mention` it only counts where it is prefixed with `@`.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords,omitzero"`
	// How the bound agent decides whether to run on incoming mail.
	//
	//   - `mention`: runs only when the agent is @mentioned, matched against the trigger
	//     keywords below.
	//   - `keyword`: runs when the message contains any of the trigger keywords.
	//   - `always`: runs on every incoming message.
	//
	// Leaving this unset makes the agent run on every incoming message, since email
	// has no reliable @mention convention.
	//
	// Any of "mention", "keyword", "always".
	AgentTriggerPolicy CreateEmailInboxRequestAgentTriggerPolicy `json:"agent_trigger_policy,omitzero"`
	paramObj
}

func (r CreateEmailInboxRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateEmailInboxRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateEmailInboxRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How the bound agent decides whether to run on incoming mail.
//
//   - `mention`: runs only when the agent is @mentioned, matched against the trigger
//     keywords below.
//   - `keyword`: runs when the message contains any of the trigger keywords.
//   - `always`: runs on every incoming message.
//
// Leaving this unset makes the agent run on every incoming message, since email
// has no reliable @mention convention.
type CreateEmailInboxRequestAgentTriggerPolicy string

const (
	CreateEmailInboxRequestAgentTriggerPolicyMention CreateEmailInboxRequestAgentTriggerPolicy = "mention"
	CreateEmailInboxRequestAgentTriggerPolicyKeyword CreateEmailInboxRequestAgentTriggerPolicy = "keyword"
	CreateEmailInboxRequestAgentTriggerPolicyAlways  CreateEmailInboxRequestAgentTriggerPolicy = "always"
)

// A routable email inbox on a verified domain.
//
// Mail sent to this address is threaded into a conversation: the first message of
// a thread opens a new customer case, and later messages in the same thread join
// the conversation it already created. Replies to the customer go back out from
// this address, and the bound agent — if there is one — can draft or send them.
type EmailInbox struct {
	// Email inbox ID.
	ID string `json:"id" api:"required"`
	// The full inbox address (e.g. `support@acme.com`).
	Address string `json:"address" api:"required"`
	// An AI agent available to the account.
	//
	// The definition describes what the agent does, how its runs are triggered, the
	// tools it can use, and whether it is currently enabled for the account.
	AgentConfig AgentDefinition `json:"agent_config" api:"required"`
	// The keywords that decide whether the agent runs on an incoming message.
	//
	// Under the `keyword` policy a keyword matches anywhere in the message; under
	// `mention` it only counts where it is prefixed with `@`.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords" api:"required"`
	// When the bound agent runs on incoming mail.
	//
	//   - `mention`: only when the agent is @mentioned, matched against its trigger
	//     keywords.
	//   - `keyword`: when the mail contains any of the configured trigger keywords.
	//   - `always`: on every incoming message.
	//
	// When no policy is set the agent runs on every incoming message, since email has
	// no reliable @mention convention.
	//
	// Any of "mention", "keyword", "always".
	AgentTriggerPolicy EmailInboxAgentTriggerPolicy `json:"agent_trigger_policy" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A domain registered with the email bridge for sending and receiving mail.
	//
	// After registration the domain starts in `pending`; publish the returned DKIM
	// records, then poll the verify action until it flips to `verified`.
	EmailDomain EmailDomain `json:"email_domain" api:"required"`
	// A forwarding address on an OpenMRP-owned domain that also routes to this inbox.
	//
	// Use this when your domain's mail is hosted elsewhere (e.g. Google Workspace,
	// Microsoft 365) and you cannot point its MX records at OpenMRP: forward mail from
	// `address` to this address instead, and it will still be threaded into a
	// conversation.
	ForwardingAddress string `json:"forwarding_address" api:"required"`
	// The display name used in the `From` header of outbound mail.
	FromName string `json:"from_name" api:"required"`
	// The messaging group (roster) whose members are added to every conversation this
	// inbox opens.
	//
	// Its members join each new email thread so the team can read, edit, and approve
	// replies alongside the bound agent. Membership is captured when the thread opens,
	// so later edits to the group only affect conversations opened after the change.
	GroupID string `json:"group_id" api:"required"`
	// Resource type identifier.
	//
	// Any of "email_inbox".
	Object EmailInboxObject `json:"object" api:"required"`
	// Whether the inbox is currently accepting mail.
	//
	//   - `active`: inbound mail is threaded into a conversation.
	//   - `disabled`: the inbox stays provisioned and keeps its history, but inbound
	//     mail is dropped without being threaded.
	//
	// Any of "active", "disabled".
	Status EmailInboxStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Address              respjson.Field
		AgentConfig          respjson.Field
		AgentTriggerKeywords respjson.Field
		AgentTriggerPolicy   respjson.Field
		CreatedAt            respjson.Field
		EmailDomain          respjson.Field
		ForwardingAddress    respjson.Field
		FromName             respjson.Field
		GroupID              respjson.Field
		Object               respjson.Field
		Status               respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInbox) RawJSON() string { return r.JSON.raw }
func (r *EmailInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// When the bound agent runs on incoming mail.
//
//   - `mention`: only when the agent is @mentioned, matched against its trigger
//     keywords.
//   - `keyword`: when the mail contains any of the configured trigger keywords.
//   - `always`: on every incoming message.
//
// When no policy is set the agent runs on every incoming message, since email has
// no reliable @mention convention.
type EmailInboxAgentTriggerPolicy string

const (
	EmailInboxAgentTriggerPolicyMention EmailInboxAgentTriggerPolicy = "mention"
	EmailInboxAgentTriggerPolicyKeyword EmailInboxAgentTriggerPolicy = "keyword"
	EmailInboxAgentTriggerPolicyAlways  EmailInboxAgentTriggerPolicy = "always"
)

// Resource type identifier.
type EmailInboxObject string

const (
	EmailInboxObjectEmailInbox EmailInboxObject = "email_inbox"
)

// Whether the inbox is currently accepting mail.
//
//   - `active`: inbound mail is threaded into a conversation.
//   - `disabled`: the inbox stays provisioned and keeps its history, but inbound
//     mail is dropped without being threaded.
type EmailInboxStatus string

const (
	EmailInboxStatusActive   EmailInboxStatus = "active"
	EmailInboxStatusDisabled EmailInboxStatus = "disabled"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListEmailInbox struct {
	// Resources in this page.
	Data []EmailInbox `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListEmailInboxObject `json:"object" api:"required"`
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
func (r ListEmailInbox) RawJSON() string { return r.JSON.raw }
func (r *ListEmailInbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListEmailInboxObject string

const (
	ListEmailInboxObjectList ListEmailInboxObject = "list"
)

// Request to edit an email inbox's from-name, status, agent configuration, and
// roster.
//
// The property Status is required.
type UpdateEmailInboxRequestParam struct {
	// Whether the inbox accepts mail.
	//
	//   - `active`: inbound mail is threaded into a conversation.
	//   - `disabled`: the inbox stays provisioned and keeps its history, but inbound
	//     mail is dropped without being threaded.
	//
	// Any of "active", "disabled".
	Status UpdateEmailInboxRequestStatus `json:"status,omitzero" api:"required"`
	// The agent to bind to this inbox to handle incoming mail.
	AgentConfigID param.Opt[string] `json:"agent_config_id,omitzero"`
	// Display name for the `From` header of outbound mail.
	FromName param.Opt[string] `json:"from_name,omitzero"`
	// The messaging group (roster) whose members are seated on every conversation this
	// inbox opens.
	//
	// Must name a group in your own account. Changing it only affects conversations
	// opened afterwards.
	GroupID param.Opt[string] `json:"group_id,omitzero"`
	// The keywords that decide whether the agent runs on an incoming message.
	//
	// Under the `keyword` policy a keyword matches anywhere in the message; under
	// `mention` it only counts where it is prefixed with `@`.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords,omitzero"`
	// How the bound agent decides whether to run on incoming mail.
	//
	//   - `mention`: runs only when the agent is @mentioned, matched against the trigger
	//     keywords below.
	//   - `keyword`: runs when the message contains any of the trigger keywords.
	//   - `always`: runs on every incoming message.
	//
	// While no policy has been set, the agent runs on every incoming message, since
	// email has no reliable @mention convention.
	//
	// Any of "mention", "keyword", "always".
	AgentTriggerPolicy UpdateEmailInboxRequestAgentTriggerPolicy `json:"agent_trigger_policy,omitzero"`
	paramObj
}

func (r UpdateEmailInboxRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateEmailInboxRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateEmailInboxRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the inbox accepts mail.
//
//   - `active`: inbound mail is threaded into a conversation.
//   - `disabled`: the inbox stays provisioned and keeps its history, but inbound
//     mail is dropped without being threaded.
type UpdateEmailInboxRequestStatus string

const (
	UpdateEmailInboxRequestStatusActive   UpdateEmailInboxRequestStatus = "active"
	UpdateEmailInboxRequestStatusDisabled UpdateEmailInboxRequestStatus = "disabled"
)

// How the bound agent decides whether to run on incoming mail.
//
//   - `mention`: runs only when the agent is @mentioned, matched against the trigger
//     keywords below.
//   - `keyword`: runs when the message contains any of the trigger keywords.
//   - `always`: runs on every incoming message.
//
// While no policy has been set, the agent runs on every incoming message, since
// email has no reliable @mention convention.
type UpdateEmailInboxRequestAgentTriggerPolicy string

const (
	UpdateEmailInboxRequestAgentTriggerPolicyMention UpdateEmailInboxRequestAgentTriggerPolicy = "mention"
	UpdateEmailInboxRequestAgentTriggerPolicyKeyword UpdateEmailInboxRequestAgentTriggerPolicy = "keyword"
	UpdateEmailInboxRequestAgentTriggerPolicyAlways  UpdateEmailInboxRequestAgentTriggerPolicy = "always"
)

type MessagingEmailInboxDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingEmailInboxDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingEmailInboxDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingEmailInboxNewParams struct {
	// Request to provision a routable inbox on a verified domain.
	CreateEmailInboxRequest CreateEmailInboxRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingEmailInboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateEmailInboxRequest)
}
func (r *MessagingEmailInboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingEmailInboxNewParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingEmailInboxGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingEmailInboxGetParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingEmailInboxUpdateParams struct {
	// Request to edit an email inbox's from-name, status, agent configuration, and
	// roster.
	UpdateEmailInboxRequest UpdateEmailInboxRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingEmailInboxUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateEmailInboxRequest)
}
func (r *MessagingEmailInboxUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingEmailInboxUpdateParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingEmailInboxListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "email_domain", "agent_config".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingEmailInboxListParams]'s query parameters as
// `url.Values`.
func (r MessagingEmailInboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
