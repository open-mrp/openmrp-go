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
)

// Send, list, edit, and delete chat messages.
//
// MessagingMessageActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingMessageActionService] method instead.
type MessagingMessageActionService struct {
	options []option.RequestOption
}

// NewMessagingMessageActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingMessageActionService(opts ...option.RequestOption) (r MessagingMessageActionService) {
	r = MessagingMessageActionService{}
	r.options = opts
	return
}

// Approves a reply draft and sends it to the customer.
//
// The draft becomes the sent message rather than spawning a copy: it takes its
// place in the case timeline, and the customer sees it as coming from "Customer
// Service". A draft on the email channel goes out as a reply on the case's email
// thread; otherwise it appears in the customer's conversation. Sending also moves
// the case to waiting on the customer.
//
// Only the first approval of a draft sends it — approving one that is no longer
// open fails, so a concurrent double-approve cannot reach the customer twice.
// Customer accounts cannot approve drafts.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) ApproveSend(ctx context.Context, id string, params MessagingMessageActionApproveSendParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/approve-send", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Cancels a message that was scheduled for a future send, so it is never
// delivered.
//
// You can only cancel a message you scheduled yourself, and only while it is still
// waiting to go out — once it has been delivered or has otherwise left the
// scheduled state, the request fails. The canceled message is kept as a record and
// never appears in the conversation.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) Cancel(ctx context.Context, id string, body MessagingMessageActionCancelParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/cancel", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Discards a reply draft without sending it to the customer.
//
// The draft is kept as a rejected record for history and can no longer be edited
// or approved. Because the customer is still owed an answer, the case moves back
// to waiting on your team.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) Reject(ctx context.Context, id string, body MessagingMessageActionRejectParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/reject", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Moves a message you scheduled to a new send time, optionally revising what it
// says, and returns it.
//
// The message keeps its id and is sent once, at the new time; the time it had
// before no longer applies. You can only reschedule a message you scheduled
// yourself, and only until its send time arrives — once it is due, sent or
// canceled the request fails.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingMessageActionService) Reschedule(ctx context.Context, id string, params MessagingMessageActionRescheduleParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/messages/%s/actions/reschedule", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request to approve a customer-reply draft and send it to the customer.
//
// The property ClientMessageID is required.
type ApproveSendDraftRequestParam struct {
	// A unique client-generated key for this approval, such as a UUID.
	ClientMessageID string `json:"client_message_id" api:"required"`
	paramObj
}

func (r ApproveSendDraftRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ApproveSendDraftRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApproveSendDraftRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to move a scheduled message to a new send time.
//
// The property ScheduledAt is required.
type RescheduleMessageRequestParam struct {
	// When the message should now be sent. Must be in the future.
	ScheduledAt time.Time `json:"scheduled_at" api:"required" format:"date-time"`
	// The revised message body, replacing what it said before.
	//
	// Leaving it out keeps the current body.
	Body param.Opt[string] `json:"body,omitzero"`
	paramObj
}

func (r RescheduleMessageRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow RescheduleMessageRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RescheduleMessageRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingMessageActionApproveSendParams struct {
	// Request to approve a customer-reply draft and send it to the customer.
	ApproveSendDraftRequest ApproveSendDraftRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingMessageActionApproveSendParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ApproveSendDraftRequest)
}
func (r *MessagingMessageActionApproveSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingMessageActionApproveSendParams]'s query parameters
// as `url.Values`.
func (r MessagingMessageActionApproveSendParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingMessageActionCancelParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingMessageActionCancelParams]'s query parameters as
// `url.Values`.
func (r MessagingMessageActionCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingMessageActionRejectParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingMessageActionRejectParams]'s query parameters as
// `url.Values`.
func (r MessagingMessageActionRejectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingMessageActionRescheduleParams struct {
	// Request to move a scheduled message to a new send time.
	RescheduleMessageRequest RescheduleMessageRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingMessageActionRescheduleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.RescheduleMessageRequest)
}
func (r *MessagingMessageActionRescheduleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingMessageActionRescheduleParams]'s query parameters
// as `url.Values`.
func (r MessagingMessageActionRescheduleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
