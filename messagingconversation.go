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

// Create conversations, send and read messages (1:1 direct messages).
//
// MessagingConversationService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationService] method instead.
type MessagingConversationService struct {
	options []option.RequestOption
	// Create conversations, send and read messages (1:1 direct messages).
	Actions MessagingConversationActionService
	// Create conversations, send and read messages (1:1 direct messages).
	Links MessagingConversationLinkService
	// Send, list, edit, and delete chat messages.
	Messages MessagingConversationMessageService
	// Add, remove, and manage participants (including agents) in a conversation.
	Participants MessagingConversationParticipantService
	Attachments  MessagingConversationAttachmentService
}

// NewMessagingConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingConversationService(opts ...option.RequestOption) (r MessagingConversationService) {
	r = MessagingConversationService{}
	r.options = opts
	r.Actions = NewMessagingConversationActionService(opts...)
	r.Links = NewMessagingConversationLinkService(opts...)
	r.Messages = NewMessagingConversationMessageService(opts...)
	r.Participants = NewMessagingConversationParticipantService(opts...)
	r.Attachments = NewMessagingConversationAttachmentService(opts...)
	return
}

// Starts a direct message or group conversation.
//
// Requesting a direct message that already exists returns the existing thread
// instead of creating a duplicate, and a direct message is refused when either
// user has blocked the other. Conversation creation is rate limited per user.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingConversationService) New(ctx context.Context, params MessagingConversationNewParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns a single conversation the caller participates in.
//
// Someone who has left the conversation can still read it back; it comes back
// marked hidden for them. A team member who opens a customer-facing case they are
// not yet part of is seated in it as a participant.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingConversationService) Get(ctx context.Context, id string, query MessagingConversationGetParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Renames a group conversation.
//
// Only an owner or admin of the conversation can rename it, and direct messages
// cannot be renamed.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationService) Update(ctx context.Context, id string, params MessagingConversationUpdateParams, opts ...option.RequestOption) (res *Conversation, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns the caller's conversations, most recently active first.
//
// A customer portal user sees only their own support case with the vendor, and an
// empty list until they have contacted support.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingConversationService) List(ctx context.Context, query MessagingConversationListParams, opts ...option.RequestOption) (res *ListConversation, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A conversation thread the caller participates in.
type Conversation struct {
	// Conversation ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Assignee Actor `json:"assignee" api:"required"`
	// Whether this is a team-only conversation (`internal`) or a customer-facing case
	// (`customer`).
	//
	// A customer never sees an `internal` conversation, even one that is about them;
	// within a `customer` case they see only the messages that were sent to them, not
	// the team's internal notes on the case.
	//
	// Any of "internal", "customer".
	Audience ConversationAudience `json:"audience" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A reusable roster: a named set of members (users and/or agents) that seeds new
	// conversations.
	//
	// Starting a conversation from a group snapshots its current members into that
	// conversation, so the same group can back many conversations (each with its own
	// title); later edits to the group never change conversations already created from
	// it.
	Group MessagingGroup `json:"group" api:"required"`
	// A chat message within a conversation.
	//
	// One resource covers every stage of a message's life: a delivered timeline
	// message, a message queued for a future send, and a customer-reply draft awaiting
	// approval. Read `status` to tell them apart.
	LastMessage *Message `json:"last_message" api:"required"`
	// When the most recent message was sent.
	LastMessageAt time.Time `json:"last_message_at" api:"required" format:"date-time"`
	// Whether the conversation is under legal hold.
	//
	// While held, the conversation is exempt from automatic retention purging and from
	// redaction until the hold is released.
	//
	// Any of "released", "held".
	LegalHold ConversationLegalHold `json:"legal_hold" api:"required"`
	// Resource type identifier.
	//
	// Any of "conversation".
	Object ConversationObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Participants ListConversationParticipant `json:"participants" api:"required"`
	// The conversation's state from the caller's point of view.
	//
	//   - `active`: a normal, visible conversation.
	//   - `archived`: archived for the whole account.
	//   - `hidden`: the caller dismissed the conversation from their own list while
	//     everyone else still sees it, which takes precedence over an account-level
	//     archive.
	//
	// Any of "active", "archived", "hidden".
	Status ConversationStatus `json:"status" api:"required"`
	// The display title of a group conversation.
	//
	// Direct messages carry no stored title; clients derive one from the participants.
	Title string `json:"title" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Topic Entity `json:"topic" api:"required"`
	// What kind of conversation this is.
	//
	//   - `direct_message`: a 1:1 thread between two users.
	//   - `group`: a named thread with multiple user or agent members (including
	//     customer-facing support cases).
	//   - `system`: a system channel that delivers automated account alerts.
	//
	// Any of "direct_message", "group", "system".
	Type ConversationType `json:"type" api:"required"`
	// Number of messages the caller has not yet read.
	Unread int64 `json:"unread" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The triage lane of a customer-facing case.
	//
	// Only conversations with a `customer` audience have a triage lane. It drives the
	// support inbox and is independent of `status`, which is about visibility rather
	// than progress.
	//
	// - `new`: opened but not yet triaged.
	// - `open`: actively being worked.
	// - `waiting_internal`: blocked on the internal team.
	// - `waiting_external`: blocked on an external reply.
	// - `needs_approval`: a drafted reply is awaiting human approval.
	// - `resolved`: closed out.
	//
	// Any of "new", "open", "waiting_internal", "waiting_external", "needs_approval",
	// "resolved".
	WorkflowStatus ConversationWorkflowStatus `json:"workflow_status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Assignee       respjson.Field
		Audience       respjson.Field
		CreatedAt      respjson.Field
		Group          respjson.Field
		LastMessage    respjson.Field
		LastMessageAt  respjson.Field
		LegalHold      respjson.Field
		Object         respjson.Field
		Participants   respjson.Field
		Status         respjson.Field
		Title          respjson.Field
		Topic          respjson.Field
		Type           respjson.Field
		Unread         respjson.Field
		UpdatedAt      respjson.Field
		WorkflowStatus respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Conversation) RawJSON() string { return r.JSON.raw }
func (r *Conversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this is a team-only conversation (`internal`) or a customer-facing case
// (`customer`).
//
// A customer never sees an `internal` conversation, even one that is about them;
// within a `customer` case they see only the messages that were sent to them, not
// the team's internal notes on the case.
type ConversationAudience string

const (
	ConversationAudienceInternal ConversationAudience = "internal"
	ConversationAudienceCustomer ConversationAudience = "customer"
)

// Whether the conversation is under legal hold.
//
// While held, the conversation is exempt from automatic retention purging and from
// redaction until the hold is released.
type ConversationLegalHold string

const (
	ConversationLegalHoldReleased ConversationLegalHold = "released"
	ConversationLegalHoldHeld     ConversationLegalHold = "held"
)

// Resource type identifier.
type ConversationObject string

const (
	ConversationObjectConversation ConversationObject = "conversation"
)

// The conversation's state from the caller's point of view.
//
//   - `active`: a normal, visible conversation.
//   - `archived`: archived for the whole account.
//   - `hidden`: the caller dismissed the conversation from their own list while
//     everyone else still sees it, which takes precedence over an account-level
//     archive.
type ConversationStatus string

const (
	ConversationStatusActive   ConversationStatus = "active"
	ConversationStatusArchived ConversationStatus = "archived"
	ConversationStatusHidden   ConversationStatus = "hidden"
)

// What kind of conversation this is.
//
//   - `direct_message`: a 1:1 thread between two users.
//   - `group`: a named thread with multiple user or agent members (including
//     customer-facing support cases).
//   - `system`: a system channel that delivers automated account alerts.
type ConversationType string

const (
	ConversationTypeDirectMessage ConversationType = "direct_message"
	ConversationTypeGroup         ConversationType = "group"
	ConversationTypeSystem        ConversationType = "system"
)

// The triage lane of a customer-facing case.
//
// Only conversations with a `customer` audience have a triage lane. It drives the
// support inbox and is independent of `status`, which is about visibility rather
// than progress.
//
// - `new`: opened but not yet triaged.
// - `open`: actively being worked.
// - `waiting_internal`: blocked on the internal team.
// - `waiting_external`: blocked on an external reply.
// - `needs_approval`: a drafted reply is awaiting human approval.
// - `resolved`: closed out.
type ConversationWorkflowStatus string

const (
	ConversationWorkflowStatusNew             ConversationWorkflowStatus = "new"
	ConversationWorkflowStatusOpen            ConversationWorkflowStatus = "open"
	ConversationWorkflowStatusWaitingInternal ConversationWorkflowStatus = "waiting_internal"
	ConversationWorkflowStatusWaitingExternal ConversationWorkflowStatus = "waiting_external"
	ConversationWorkflowStatusNeedsApproval   ConversationWorkflowStatus = "needs_approval"
	ConversationWorkflowStatusResolved        ConversationWorkflowStatus = "resolved"
)

// A participant (membership) in a conversation.
type ConversationParticipant struct {
	// Participant ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// For agent participants with a keyword or mention policy, the keywords that
	// trigger it.
	//
	// Matching is case-insensitive and looks anywhere in the message body: under
	// `keyword` the bare word is matched, under `mention` it must appear as
	// `@keyword`. Replying directly to one of the agent's own messages always reaches
	// it, so an agent with no keywords still answers replies but nothing else.
	AgentTriggerKeywords []string `json:"agent_trigger_keywords" api:"required"`
	// For agent participants, when the agent is invoked in response to messages.
	//
	// - `mention`: only when the agent is @mentioned.
	// - `keyword`: when a message contains one of the agent's trigger keywords.
	// - `always`: on every human message in the conversation.
	//
	// Any of "mention", "keyword", "always".
	AgentTriggerPolicy ConversationParticipantAgentTriggerPolicy `json:"agent_trigger_policy" api:"required"`
	// The participant's membership in the conversation.
	//
	// - `active`: currently a member.
	// - `left`: voluntarily left the conversation.
	// - `removed`: removed by an admin.
	// - `hidden`: still a member but has hidden the conversation from their own list.
	//
	// Membership records are kept rather than deleted, so re-adding someone who left
	// or was removed reactivates their original record and their earlier messages stay
	// attributed to them.
	//
	// Any of "active", "left", "removed", "hidden".
	Membership ConversationParticipantMembership `json:"membership" api:"required"`
	// The participant's notification preference for the conversation.
	//
	//   - `unmuted`: receives notifications for new messages.
	//   - `muted`: new-message notifications are suppressed, though a direct @mention
	//     still raises an in-app alert (never an email), and the conversation still
	//     counts toward the unread total.
	//
	// Any of "unmuted", "muted".
	Notifications ConversationParticipantNotifications `json:"notifications" api:"required"`
	// Resource type identifier.
	//
	// Any of "conversation_participant".
	Object ConversationParticipantObject `json:"object" api:"required"`
	// A participant's read position in a conversation — the basis for read receipts
	// ("who has seen this").
	ReadCursor ReadCursor `json:"read_cursor" api:"required"`
	// The participant's permission level in the conversation.
	//
	//   - `owner`: can rename or delete the conversation and manage its members and
	//     their roles.
	//   - `admin`: can add or remove members and rename the conversation.
	//   - `member`: can post, react, mute, and leave.
	//   - `viewer`: read-only access.
	//
	// Any of "owner", "admin", "member", "viewer".
	Role ConversationParticipantRole `json:"role" api:"required"`
	// The kind of participant.
	//
	// - `user`: an account user (a teammate).
	// - `agent`: an AI agent.
	// - `system`: the system itself, which posts automated messages.
	// - `customer`: an external customer in a support case.
	//
	// Any of "user", "agent", "system", "customer".
	Type ConversationParticipantType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Actor                respjson.Field
		AgentTriggerKeywords respjson.Field
		AgentTriggerPolicy   respjson.Field
		Membership           respjson.Field
		Notifications        respjson.Field
		Object               respjson.Field
		ReadCursor           respjson.Field
		Role                 respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationParticipant) RawJSON() string { return r.JSON.raw }
func (r *ConversationParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// For agent participants, when the agent is invoked in response to messages.
//
// - `mention`: only when the agent is @mentioned.
// - `keyword`: when a message contains one of the agent's trigger keywords.
// - `always`: on every human message in the conversation.
type ConversationParticipantAgentTriggerPolicy string

const (
	ConversationParticipantAgentTriggerPolicyMention ConversationParticipantAgentTriggerPolicy = "mention"
	ConversationParticipantAgentTriggerPolicyKeyword ConversationParticipantAgentTriggerPolicy = "keyword"
	ConversationParticipantAgentTriggerPolicyAlways  ConversationParticipantAgentTriggerPolicy = "always"
)

// The participant's membership in the conversation.
//
// - `active`: currently a member.
// - `left`: voluntarily left the conversation.
// - `removed`: removed by an admin.
// - `hidden`: still a member but has hidden the conversation from their own list.
//
// Membership records are kept rather than deleted, so re-adding someone who left
// or was removed reactivates their original record and their earlier messages stay
// attributed to them.
type ConversationParticipantMembership string

const (
	ConversationParticipantMembershipActive  ConversationParticipantMembership = "active"
	ConversationParticipantMembershipLeft    ConversationParticipantMembership = "left"
	ConversationParticipantMembershipRemoved ConversationParticipantMembership = "removed"
	ConversationParticipantMembershipHidden  ConversationParticipantMembership = "hidden"
)

// The participant's notification preference for the conversation.
//
//   - `unmuted`: receives notifications for new messages.
//   - `muted`: new-message notifications are suppressed, though a direct @mention
//     still raises an in-app alert (never an email), and the conversation still
//     counts toward the unread total.
type ConversationParticipantNotifications string

const (
	ConversationParticipantNotificationsUnmuted ConversationParticipantNotifications = "unmuted"
	ConversationParticipantNotificationsMuted   ConversationParticipantNotifications = "muted"
)

// Resource type identifier.
type ConversationParticipantObject string

const (
	ConversationParticipantObjectConversationParticipant ConversationParticipantObject = "conversation_participant"
)

// The participant's permission level in the conversation.
//
//   - `owner`: can rename or delete the conversation and manage its members and
//     their roles.
//   - `admin`: can add or remove members and rename the conversation.
//   - `member`: can post, react, mute, and leave.
//   - `viewer`: read-only access.
type ConversationParticipantRole string

const (
	ConversationParticipantRoleOwner  ConversationParticipantRole = "owner"
	ConversationParticipantRoleAdmin  ConversationParticipantRole = "admin"
	ConversationParticipantRoleMember ConversationParticipantRole = "member"
	ConversationParticipantRoleViewer ConversationParticipantRole = "viewer"
)

// The kind of participant.
//
// - `user`: an account user (a teammate).
// - `agent`: an AI agent.
// - `system`: the system itself, which posts automated messages.
// - `customer`: an external customer in a support case.
type ConversationParticipantType string

const (
	ConversationParticipantTypeUser     ConversationParticipantType = "user"
	ConversationParticipantTypeAgent    ConversationParticipantType = "agent"
	ConversationParticipantTypeSystem   ConversationParticipantType = "system"
	ConversationParticipantTypeCustomer ConversationParticipantType = "customer"
)

// Request to create a conversation.
//
// The properties ParticipantAccountUserIDs, Type are required.
type CreateConversationRequestParam struct {
	// The other participants to add.
	//
	// For a direct message, exactly one account user. For a group, the members to seed
	// — these can be omitted when `group_id` supplies a roster, or when the
	// conversation is anchored to a topic resource, since a record discussion may
	// start solo and pull people in later.
	//
	// The caller is always a participant and does not need to be listed; on a group
	// they become its owner and every other member seeded at creation is notified.
	ParticipantAccountUserIDs []string `json:"participant_account_user_ids,omitzero" api:"required"`
	// The kind of conversation to create.
	//
	//   - `direct_message`: a 1:1 thread with exactly one other user. Addressing
	//     yourself is allowed and gives you a private notes thread.
	//   - `group`: a named thread with any number of user and agent members.
	//
	// `system` channels are created by the platform and cannot be requested here.
	//
	// Any of "direct_message", "group", "system".
	Type CreateConversationRequestType `json:"type,omitzero" api:"required"`
	// Seed a group conversation from a reusable roster.
	//
	// The roster's current members are copied into this conversation (in addition to
	// any `participant_account_user_ids`); the conversation is independent afterward.
	// Ignored for direct messages.
	GroupID param.Opt[string] `json:"group_id,omitzero"`
	// Title for a group conversation.
	//
	// A direct message is identified by its participants rather than by a title.
	Title param.Opt[string] `json:"title,omitzero"`
	// The id of the business record to anchor this conversation to.
	TopicResourceID param.Opt[string] `json:"topic_resource_id,omitzero"`
	// The type of business record to anchor this conversation to.
	//
	// An anchored conversation is returned when conversations are listed for that
	// record, which is how a discussion shows up on an order or invoice.
	//
	// Any of "account", "actor", "entity", "record", "freight", "sales_order_totals",
	// "sales_order_stage_total", "sales_order_related", "order_contact", "user",
	// "address", "api_key", "created_api_key", "refresh_token", "list", "sandbox",
	// "registration_session", "pricing_plan", "account_plan", "plan_change",
	// "enterprise_inquiry", "request_log", "audit_event", "audit_field_change",
	// "role", "unit", "account_affiliation", "agent_definition", "available_tool",
	// "agent_definition_tool", "agent_account_status", "agent_run", "agent_action",
	// "agent_run_step", "agent_token_usage", "agent_memory", "notification",
	// "notification_unread_count", "notification_send_result",
	// "notification_unread_summary", "announcement", "conversation", "support_case",
	// "conversation_participant", "read_cursor", "chat_message",
	// "notification_unread_summary_account", "messaging_block",
	// "notification_preference", "message_attachment", "attachment_upload_target",
	// "scheduled_message", "messaging_contact", "message_report", "tool_group",
	// "model", "payment_term", "shipping_term", "quantity", "account_group",
	// "support_route", "support_availability", "account_status", "geolocation",
	// "account_user", "department", "account_integration", "account_price",
	// "product_line", "item_category", "attribute", "rate",
	// "account_group_product_line_access", "sales_target", "adjustment_type",
	// "account_branding", "account_portal", "account_logo_url", "account_favicon_url",
	// "public_account", "property", "carrier", "service_level", "item",
	// "item_lot_default", "item_inventory", "product", "batch", "batch_flow_node",
	// "scanning_consumption", "open_batch_summary", "scanning_production_step_info",
	// "scanning_station", "production_step", "production_run", "machine",
	// "machine_status", "machine_downtime_event", "demand_override",
	// "demand_override_type", "machine_downtime_reason",
	// "production_schedule_preview", "production_schedule_regenerate_preview",
	// "production_schedule", "production_schedule_line",
	// "production_schedule_deviation", "production_schedule_derived_line",
	// "production_schedule_settings", "production_schedule_resource_setting",
	// "production_schedule_item_setting", "fulfillment_recommendation",
	// "analyze_delivery_performance_response", "delivery_performance",
	// "delivery_backlog_bucket", "delivery_lateness_bucket", "delivery_breakdown",
	// "schedule_order_coverage", "schedule_order_coverage_line",
	// "schedule_deviation_type", "schedule_at_risk_order",
	// "production_schedule_finished_policy", "production_schedule_finishing_line",
	// "production_schedule_week_release", "production_schedule_week_release_preview",
	// "production_schedule_item_policy", "child_account", "unit_group",
	// "unit_group_unit", "consumption", "customer_product_line_access", "customer",
	// "frequently_ordered_product", "priority", "delivery", "delivery_line",
	// "sales_order", "location", "location_type", "lot", "email_log", "email_domain",
	// "email_inbox", "portal_domain", "dns_record", "inventory_change_log", "invoice",
	// "invoice_summary", "invoice_line", "invoice_allocation", "invoice_for_payment",
	// "shipment", "shipment_summary", "shipment_line", "shipping_case",
	// "shipping_case_label_url", "settlement", "settlement_summary",
	// "role_permission", "registration_flow", "registration_flow_option",
	// "transaction", "transaction_summary", "transaction_method", "transaction_type",
	// "transaction_allocation", "usage_item", "account_usage_response",
	// "subscription_info", "billing_portal_session_response", "switch_plan_response",
	// "ensure_billing_customer_response", "spending_cap_response", "agent_spend_info",
	// "webhook_response", "address_suggestion", "address_components",
	// "address_details_result", "validated_address", "plan_limit",
	// "plan_change_proration", "plan_change_line_item", "setup_billing_response",
	// "confirm_payment_response", "oauth_response", "oauth_status_response",
	// "stripe_publishable_key", "stripe_status", "healthcheck",
	// "agent_definition_config", "trigger_config", "customer_contact_info",
	// "customer_freight_preferences", "customer_defaults", "customer_lead_time",
	// "customer_notification_preferences", "order_notification_recipient",
	// "order_discount", "sales_order_line", "sales_order_type", "sales_order_status",
	// "material", "supplier_material", "part", "permission_group", "permission",
	// "pick", "pick_line", "product_type", "production", "production_flow", "map",
	// "purchase_order", "purchase_order_line", "supplier", "supplier_summary",
	// "receivable_entry", "receiving_order", "receiving_order_line", "email_contact",
	// "allocation_entry", "open_credit_entry", "volume_discount",
	// "volume_discount_tier", "analyze_deliveries_response",
	// "analyze_manufacturing_response", "analyze_manufacturing_batch_response",
	// "analyze_quarterly_orders_response", "analyze_new_customers_response",
	// "analyze_demand_forecast_response", "analyze_oee_response",
	// "analyze_oee_trend_response", "analyze_schedule_attainment_response",
	// "catalog_product_line", "catalog_category", "catalog_product",
	// "catalog_property", "catalog_attribute", "dc_location", "edi_run",
	// "inventory_item", "analyze_weeks_of_sales_response",
	// "bulk_reconcile_items_response", "sys_property", "sys_property_type",
	// "sys_property_value", "territory", "tenancy", "checkout_session",
	// "estimate_rate_result", "rate_shop_option", "rate_shop_result", "owner",
	// "created_by", "message", "account_photo_upload_result",
	// "user_photo_upload_result", "user_photo_url", "batch_lot",
	// "check_duplicate_result", "item_trend_point", "tenancy_pending_registration",
	// "invoice_allocation_entry", "allocation_customer", "checkout_sales_order",
	// "sales_order_price_quote", "sales_order_freight_quote",
	// "sales_order_commitment_quote", "operating_calendar",
	// "operating_calendar_closure", "sales_order_price_quote_line",
	// "hubspot_sync_job", "hubspot_sync_report", "hubspot_company_review",
	// "hubspot_company_candidate", "hubspot_sync_record", "contact_match",
	// "reply_draft", "conversation_link", "messaging_group", "messaging_group_member",
	// "portal_profile", "portal_registration_session",
	// "portal_registration_session_data", "pack_list", "pack_list_party",
	// "pack_list_line_item", "pack_list_back_order", "pack_list_case", "job",
	// "job_result", "job_export", "analyze_customer_pricing_response",
	// "customer_pricing_finding", "customer_pricing_summary", "computed_rate",
	// "computed_quantity", "analyze_realized_margins_response",
	// "realized_margin_finding", "realized_margin_summary", "shipment_related",
	// "invoice_related", "pick_related", "pick_shipments_response", "pick_totals",
	// "pick_stage_total".
	TopicResourceType CreateConversationRequestTopicResourceType `json:"topic_resource_type,omitzero"`
	paramObj
}

func (r CreateConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of conversation to create.
//
//   - `direct_message`: a 1:1 thread with exactly one other user. Addressing
//     yourself is allowed and gives you a private notes thread.
//   - `group`: a named thread with any number of user and agent members.
//
// `system` channels are created by the platform and cannot be requested here.
type CreateConversationRequestType string

const (
	CreateConversationRequestTypeDirectMessage CreateConversationRequestType = "direct_message"
	CreateConversationRequestTypeGroup         CreateConversationRequestType = "group"
	CreateConversationRequestTypeSystem        CreateConversationRequestType = "system"
)

// The type of business record to anchor this conversation to.
//
// An anchored conversation is returned when conversations are listed for that
// record, which is how a discussion shows up on an order or invoice.
type CreateConversationRequestTopicResourceType string

const (
	CreateConversationRequestTopicResourceTypeAccount                              CreateConversationRequestTopicResourceType = "account"
	CreateConversationRequestTopicResourceTypeActor                                CreateConversationRequestTopicResourceType = "actor"
	CreateConversationRequestTopicResourceTypeEntity                               CreateConversationRequestTopicResourceType = "entity"
	CreateConversationRequestTopicResourceTypeRecord                               CreateConversationRequestTopicResourceType = "record"
	CreateConversationRequestTopicResourceTypeFreight                              CreateConversationRequestTopicResourceType = "freight"
	CreateConversationRequestTopicResourceTypeSalesOrderTotals                     CreateConversationRequestTopicResourceType = "sales_order_totals"
	CreateConversationRequestTopicResourceTypeSalesOrderStageTotal                 CreateConversationRequestTopicResourceType = "sales_order_stage_total"
	CreateConversationRequestTopicResourceTypeSalesOrderRelated                    CreateConversationRequestTopicResourceType = "sales_order_related"
	CreateConversationRequestTopicResourceTypeOrderContact                         CreateConversationRequestTopicResourceType = "order_contact"
	CreateConversationRequestTopicResourceTypeUser                                 CreateConversationRequestTopicResourceType = "user"
	CreateConversationRequestTopicResourceTypeAddress                              CreateConversationRequestTopicResourceType = "address"
	CreateConversationRequestTopicResourceTypeAPIKey                               CreateConversationRequestTopicResourceType = "api_key"
	CreateConversationRequestTopicResourceTypeCreatedAPIKey                        CreateConversationRequestTopicResourceType = "created_api_key"
	CreateConversationRequestTopicResourceTypeRefreshToken                         CreateConversationRequestTopicResourceType = "refresh_token"
	CreateConversationRequestTopicResourceTypeList                                 CreateConversationRequestTopicResourceType = "list"
	CreateConversationRequestTopicResourceTypeSandbox                              CreateConversationRequestTopicResourceType = "sandbox"
	CreateConversationRequestTopicResourceTypeRegistrationSession                  CreateConversationRequestTopicResourceType = "registration_session"
	CreateConversationRequestTopicResourceTypePricingPlan                          CreateConversationRequestTopicResourceType = "pricing_plan"
	CreateConversationRequestTopicResourceTypeAccountPlan                          CreateConversationRequestTopicResourceType = "account_plan"
	CreateConversationRequestTopicResourceTypePlanChange                           CreateConversationRequestTopicResourceType = "plan_change"
	CreateConversationRequestTopicResourceTypeEnterpriseInquiry                    CreateConversationRequestTopicResourceType = "enterprise_inquiry"
	CreateConversationRequestTopicResourceTypeRequestLog                           CreateConversationRequestTopicResourceType = "request_log"
	CreateConversationRequestTopicResourceTypeAuditEvent                           CreateConversationRequestTopicResourceType = "audit_event"
	CreateConversationRequestTopicResourceTypeAuditFieldChange                     CreateConversationRequestTopicResourceType = "audit_field_change"
	CreateConversationRequestTopicResourceTypeRole                                 CreateConversationRequestTopicResourceType = "role"
	CreateConversationRequestTopicResourceTypeUnit                                 CreateConversationRequestTopicResourceType = "unit"
	CreateConversationRequestTopicResourceTypeAccountAffiliation                   CreateConversationRequestTopicResourceType = "account_affiliation"
	CreateConversationRequestTopicResourceTypeAgentDefinition                      CreateConversationRequestTopicResourceType = "agent_definition"
	CreateConversationRequestTopicResourceTypeAvailableTool                        CreateConversationRequestTopicResourceType = "available_tool"
	CreateConversationRequestTopicResourceTypeAgentDefinitionTool                  CreateConversationRequestTopicResourceType = "agent_definition_tool"
	CreateConversationRequestTopicResourceTypeAgentAccountStatus                   CreateConversationRequestTopicResourceType = "agent_account_status"
	CreateConversationRequestTopicResourceTypeAgentRun                             CreateConversationRequestTopicResourceType = "agent_run"
	CreateConversationRequestTopicResourceTypeAgentAction                          CreateConversationRequestTopicResourceType = "agent_action"
	CreateConversationRequestTopicResourceTypeAgentRunStep                         CreateConversationRequestTopicResourceType = "agent_run_step"
	CreateConversationRequestTopicResourceTypeAgentTokenUsage                      CreateConversationRequestTopicResourceType = "agent_token_usage"
	CreateConversationRequestTopicResourceTypeAgentMemory                          CreateConversationRequestTopicResourceType = "agent_memory"
	CreateConversationRequestTopicResourceTypeNotification                         CreateConversationRequestTopicResourceType = "notification"
	CreateConversationRequestTopicResourceTypeNotificationUnreadCount              CreateConversationRequestTopicResourceType = "notification_unread_count"
	CreateConversationRequestTopicResourceTypeNotificationSendResult               CreateConversationRequestTopicResourceType = "notification_send_result"
	CreateConversationRequestTopicResourceTypeNotificationUnreadSummary            CreateConversationRequestTopicResourceType = "notification_unread_summary"
	CreateConversationRequestTopicResourceTypeAnnouncement                         CreateConversationRequestTopicResourceType = "announcement"
	CreateConversationRequestTopicResourceTypeConversation                         CreateConversationRequestTopicResourceType = "conversation"
	CreateConversationRequestTopicResourceTypeSupportCase                          CreateConversationRequestTopicResourceType = "support_case"
	CreateConversationRequestTopicResourceTypeConversationParticipant              CreateConversationRequestTopicResourceType = "conversation_participant"
	CreateConversationRequestTopicResourceTypeReadCursor                           CreateConversationRequestTopicResourceType = "read_cursor"
	CreateConversationRequestTopicResourceTypeChatMessage                          CreateConversationRequestTopicResourceType = "chat_message"
	CreateConversationRequestTopicResourceTypeNotificationUnreadSummaryAccount     CreateConversationRequestTopicResourceType = "notification_unread_summary_account"
	CreateConversationRequestTopicResourceTypeMessagingBlock                       CreateConversationRequestTopicResourceType = "messaging_block"
	CreateConversationRequestTopicResourceTypeNotificationPreference               CreateConversationRequestTopicResourceType = "notification_preference"
	CreateConversationRequestTopicResourceTypeMessageAttachment                    CreateConversationRequestTopicResourceType = "message_attachment"
	CreateConversationRequestTopicResourceTypeAttachmentUploadTarget               CreateConversationRequestTopicResourceType = "attachment_upload_target"
	CreateConversationRequestTopicResourceTypeScheduledMessage                     CreateConversationRequestTopicResourceType = "scheduled_message"
	CreateConversationRequestTopicResourceTypeMessagingContact                     CreateConversationRequestTopicResourceType = "messaging_contact"
	CreateConversationRequestTopicResourceTypeMessageReport                        CreateConversationRequestTopicResourceType = "message_report"
	CreateConversationRequestTopicResourceTypeToolGroup                            CreateConversationRequestTopicResourceType = "tool_group"
	CreateConversationRequestTopicResourceTypeModel                                CreateConversationRequestTopicResourceType = "model"
	CreateConversationRequestTopicResourceTypePaymentTerm                          CreateConversationRequestTopicResourceType = "payment_term"
	CreateConversationRequestTopicResourceTypeShippingTerm                         CreateConversationRequestTopicResourceType = "shipping_term"
	CreateConversationRequestTopicResourceTypeQuantity                             CreateConversationRequestTopicResourceType = "quantity"
	CreateConversationRequestTopicResourceTypeAccountGroup                         CreateConversationRequestTopicResourceType = "account_group"
	CreateConversationRequestTopicResourceTypeSupportRoute                         CreateConversationRequestTopicResourceType = "support_route"
	CreateConversationRequestTopicResourceTypeSupportAvailability                  CreateConversationRequestTopicResourceType = "support_availability"
	CreateConversationRequestTopicResourceTypeAccountStatus                        CreateConversationRequestTopicResourceType = "account_status"
	CreateConversationRequestTopicResourceTypeGeolocation                          CreateConversationRequestTopicResourceType = "geolocation"
	CreateConversationRequestTopicResourceTypeAccountUser                          CreateConversationRequestTopicResourceType = "account_user"
	CreateConversationRequestTopicResourceTypeDepartment                           CreateConversationRequestTopicResourceType = "department"
	CreateConversationRequestTopicResourceTypeAccountIntegration                   CreateConversationRequestTopicResourceType = "account_integration"
	CreateConversationRequestTopicResourceTypeAccountPrice                         CreateConversationRequestTopicResourceType = "account_price"
	CreateConversationRequestTopicResourceTypeProductLine                          CreateConversationRequestTopicResourceType = "product_line"
	CreateConversationRequestTopicResourceTypeItemCategory                         CreateConversationRequestTopicResourceType = "item_category"
	CreateConversationRequestTopicResourceTypeAttribute                            CreateConversationRequestTopicResourceType = "attribute"
	CreateConversationRequestTopicResourceTypeRate                                 CreateConversationRequestTopicResourceType = "rate"
	CreateConversationRequestTopicResourceTypeAccountGroupProductLineAccess        CreateConversationRequestTopicResourceType = "account_group_product_line_access"
	CreateConversationRequestTopicResourceTypeSalesTarget                          CreateConversationRequestTopicResourceType = "sales_target"
	CreateConversationRequestTopicResourceTypeAdjustmentType                       CreateConversationRequestTopicResourceType = "adjustment_type"
	CreateConversationRequestTopicResourceTypeAccountBranding                      CreateConversationRequestTopicResourceType = "account_branding"
	CreateConversationRequestTopicResourceTypeAccountPortal                        CreateConversationRequestTopicResourceType = "account_portal"
	CreateConversationRequestTopicResourceTypeAccountLogoURL                       CreateConversationRequestTopicResourceType = "account_logo_url"
	CreateConversationRequestTopicResourceTypeAccountFaviconURL                    CreateConversationRequestTopicResourceType = "account_favicon_url"
	CreateConversationRequestTopicResourceTypePublicAccount                        CreateConversationRequestTopicResourceType = "public_account"
	CreateConversationRequestTopicResourceTypeProperty                             CreateConversationRequestTopicResourceType = "property"
	CreateConversationRequestTopicResourceTypeCarrier                              CreateConversationRequestTopicResourceType = "carrier"
	CreateConversationRequestTopicResourceTypeServiceLevel                         CreateConversationRequestTopicResourceType = "service_level"
	CreateConversationRequestTopicResourceTypeItem                                 CreateConversationRequestTopicResourceType = "item"
	CreateConversationRequestTopicResourceTypeItemLotDefault                       CreateConversationRequestTopicResourceType = "item_lot_default"
	CreateConversationRequestTopicResourceTypeItemInventory                        CreateConversationRequestTopicResourceType = "item_inventory"
	CreateConversationRequestTopicResourceTypeProduct                              CreateConversationRequestTopicResourceType = "product"
	CreateConversationRequestTopicResourceTypeBatch                                CreateConversationRequestTopicResourceType = "batch"
	CreateConversationRequestTopicResourceTypeBatchFlowNode                        CreateConversationRequestTopicResourceType = "batch_flow_node"
	CreateConversationRequestTopicResourceTypeScanningConsumption                  CreateConversationRequestTopicResourceType = "scanning_consumption"
	CreateConversationRequestTopicResourceTypeOpenBatchSummary                     CreateConversationRequestTopicResourceType = "open_batch_summary"
	CreateConversationRequestTopicResourceTypeScanningProductionStepInfo           CreateConversationRequestTopicResourceType = "scanning_production_step_info"
	CreateConversationRequestTopicResourceTypeScanningStation                      CreateConversationRequestTopicResourceType = "scanning_station"
	CreateConversationRequestTopicResourceTypeProductionStep                       CreateConversationRequestTopicResourceType = "production_step"
	CreateConversationRequestTopicResourceTypeProductionRun                        CreateConversationRequestTopicResourceType = "production_run"
	CreateConversationRequestTopicResourceTypeMachine                              CreateConversationRequestTopicResourceType = "machine"
	CreateConversationRequestTopicResourceTypeMachineStatus                        CreateConversationRequestTopicResourceType = "machine_status"
	CreateConversationRequestTopicResourceTypeMachineDowntimeEvent                 CreateConversationRequestTopicResourceType = "machine_downtime_event"
	CreateConversationRequestTopicResourceTypeDemandOverride                       CreateConversationRequestTopicResourceType = "demand_override"
	CreateConversationRequestTopicResourceTypeDemandOverrideType                   CreateConversationRequestTopicResourceType = "demand_override_type"
	CreateConversationRequestTopicResourceTypeMachineDowntimeReason                CreateConversationRequestTopicResourceType = "machine_downtime_reason"
	CreateConversationRequestTopicResourceTypeProductionSchedulePreview            CreateConversationRequestTopicResourceType = "production_schedule_preview"
	CreateConversationRequestTopicResourceTypeProductionScheduleRegeneratePreview  CreateConversationRequestTopicResourceType = "production_schedule_regenerate_preview"
	CreateConversationRequestTopicResourceTypeProductionSchedule                   CreateConversationRequestTopicResourceType = "production_schedule"
	CreateConversationRequestTopicResourceTypeProductionScheduleLine               CreateConversationRequestTopicResourceType = "production_schedule_line"
	CreateConversationRequestTopicResourceTypeProductionScheduleDeviation          CreateConversationRequestTopicResourceType = "production_schedule_deviation"
	CreateConversationRequestTopicResourceTypeProductionScheduleDerivedLine        CreateConversationRequestTopicResourceType = "production_schedule_derived_line"
	CreateConversationRequestTopicResourceTypeProductionScheduleSettings           CreateConversationRequestTopicResourceType = "production_schedule_settings"
	CreateConversationRequestTopicResourceTypeProductionScheduleResourceSetting    CreateConversationRequestTopicResourceType = "production_schedule_resource_setting"
	CreateConversationRequestTopicResourceTypeProductionScheduleItemSetting        CreateConversationRequestTopicResourceType = "production_schedule_item_setting"
	CreateConversationRequestTopicResourceTypeFulfillmentRecommendation            CreateConversationRequestTopicResourceType = "fulfillment_recommendation"
	CreateConversationRequestTopicResourceTypeAnalyzeDeliveryPerformanceResponse   CreateConversationRequestTopicResourceType = "analyze_delivery_performance_response"
	CreateConversationRequestTopicResourceTypeDeliveryPerformance                  CreateConversationRequestTopicResourceType = "delivery_performance"
	CreateConversationRequestTopicResourceTypeDeliveryBacklogBucket                CreateConversationRequestTopicResourceType = "delivery_backlog_bucket"
	CreateConversationRequestTopicResourceTypeDeliveryLatenessBucket               CreateConversationRequestTopicResourceType = "delivery_lateness_bucket"
	CreateConversationRequestTopicResourceTypeDeliveryBreakdown                    CreateConversationRequestTopicResourceType = "delivery_breakdown"
	CreateConversationRequestTopicResourceTypeScheduleOrderCoverage                CreateConversationRequestTopicResourceType = "schedule_order_coverage"
	CreateConversationRequestTopicResourceTypeScheduleOrderCoverageLine            CreateConversationRequestTopicResourceType = "schedule_order_coverage_line"
	CreateConversationRequestTopicResourceTypeScheduleDeviationType                CreateConversationRequestTopicResourceType = "schedule_deviation_type"
	CreateConversationRequestTopicResourceTypeScheduleAtRiskOrder                  CreateConversationRequestTopicResourceType = "schedule_at_risk_order"
	CreateConversationRequestTopicResourceTypeProductionScheduleFinishedPolicy     CreateConversationRequestTopicResourceType = "production_schedule_finished_policy"
	CreateConversationRequestTopicResourceTypeProductionScheduleFinishingLine      CreateConversationRequestTopicResourceType = "production_schedule_finishing_line"
	CreateConversationRequestTopicResourceTypeProductionScheduleWeekRelease        CreateConversationRequestTopicResourceType = "production_schedule_week_release"
	CreateConversationRequestTopicResourceTypeProductionScheduleWeekReleasePreview CreateConversationRequestTopicResourceType = "production_schedule_week_release_preview"
	CreateConversationRequestTopicResourceTypeProductionScheduleItemPolicy         CreateConversationRequestTopicResourceType = "production_schedule_item_policy"
	CreateConversationRequestTopicResourceTypeChildAccount                         CreateConversationRequestTopicResourceType = "child_account"
	CreateConversationRequestTopicResourceTypeUnitGroup                            CreateConversationRequestTopicResourceType = "unit_group"
	CreateConversationRequestTopicResourceTypeUnitGroupUnit                        CreateConversationRequestTopicResourceType = "unit_group_unit"
	CreateConversationRequestTopicResourceTypeConsumption                          CreateConversationRequestTopicResourceType = "consumption"
	CreateConversationRequestTopicResourceTypeCustomerProductLineAccess            CreateConversationRequestTopicResourceType = "customer_product_line_access"
	CreateConversationRequestTopicResourceTypeCustomer                             CreateConversationRequestTopicResourceType = "customer"
	CreateConversationRequestTopicResourceTypeFrequentlyOrderedProduct             CreateConversationRequestTopicResourceType = "frequently_ordered_product"
	CreateConversationRequestTopicResourceTypePriority                             CreateConversationRequestTopicResourceType = "priority"
	CreateConversationRequestTopicResourceTypeDelivery                             CreateConversationRequestTopicResourceType = "delivery"
	CreateConversationRequestTopicResourceTypeDeliveryLine                         CreateConversationRequestTopicResourceType = "delivery_line"
	CreateConversationRequestTopicResourceTypeSalesOrder                           CreateConversationRequestTopicResourceType = "sales_order"
	CreateConversationRequestTopicResourceTypeLocation                             CreateConversationRequestTopicResourceType = "location"
	CreateConversationRequestTopicResourceTypeLocationType                         CreateConversationRequestTopicResourceType = "location_type"
	CreateConversationRequestTopicResourceTypeLot                                  CreateConversationRequestTopicResourceType = "lot"
	CreateConversationRequestTopicResourceTypeEmailLog                             CreateConversationRequestTopicResourceType = "email_log"
	CreateConversationRequestTopicResourceTypeEmailDomain                          CreateConversationRequestTopicResourceType = "email_domain"
	CreateConversationRequestTopicResourceTypeEmailInbox                           CreateConversationRequestTopicResourceType = "email_inbox"
	CreateConversationRequestTopicResourceTypePortalDomain                         CreateConversationRequestTopicResourceType = "portal_domain"
	CreateConversationRequestTopicResourceTypeDNSRecord                            CreateConversationRequestTopicResourceType = "dns_record"
	CreateConversationRequestTopicResourceTypeInventoryChangeLog                   CreateConversationRequestTopicResourceType = "inventory_change_log"
	CreateConversationRequestTopicResourceTypeInvoice                              CreateConversationRequestTopicResourceType = "invoice"
	CreateConversationRequestTopicResourceTypeInvoiceSummary                       CreateConversationRequestTopicResourceType = "invoice_summary"
	CreateConversationRequestTopicResourceTypeInvoiceLine                          CreateConversationRequestTopicResourceType = "invoice_line"
	CreateConversationRequestTopicResourceTypeInvoiceAllocation                    CreateConversationRequestTopicResourceType = "invoice_allocation"
	CreateConversationRequestTopicResourceTypeInvoiceForPayment                    CreateConversationRequestTopicResourceType = "invoice_for_payment"
	CreateConversationRequestTopicResourceTypeShipment                             CreateConversationRequestTopicResourceType = "shipment"
	CreateConversationRequestTopicResourceTypeShipmentSummary                      CreateConversationRequestTopicResourceType = "shipment_summary"
	CreateConversationRequestTopicResourceTypeShipmentLine                         CreateConversationRequestTopicResourceType = "shipment_line"
	CreateConversationRequestTopicResourceTypeShippingCase                         CreateConversationRequestTopicResourceType = "shipping_case"
	CreateConversationRequestTopicResourceTypeShippingCaseLabelURL                 CreateConversationRequestTopicResourceType = "shipping_case_label_url"
	CreateConversationRequestTopicResourceTypeSettlement                           CreateConversationRequestTopicResourceType = "settlement"
	CreateConversationRequestTopicResourceTypeSettlementSummary                    CreateConversationRequestTopicResourceType = "settlement_summary"
	CreateConversationRequestTopicResourceTypeRolePermission                       CreateConversationRequestTopicResourceType = "role_permission"
	CreateConversationRequestTopicResourceTypeRegistrationFlow                     CreateConversationRequestTopicResourceType = "registration_flow"
	CreateConversationRequestTopicResourceTypeRegistrationFlowOption               CreateConversationRequestTopicResourceType = "registration_flow_option"
	CreateConversationRequestTopicResourceTypeTransaction                          CreateConversationRequestTopicResourceType = "transaction"
	CreateConversationRequestTopicResourceTypeTransactionSummary                   CreateConversationRequestTopicResourceType = "transaction_summary"
	CreateConversationRequestTopicResourceTypeTransactionMethod                    CreateConversationRequestTopicResourceType = "transaction_method"
	CreateConversationRequestTopicResourceTypeTransactionType                      CreateConversationRequestTopicResourceType = "transaction_type"
	CreateConversationRequestTopicResourceTypeTransactionAllocation                CreateConversationRequestTopicResourceType = "transaction_allocation"
	CreateConversationRequestTopicResourceTypeUsageItem                            CreateConversationRequestTopicResourceType = "usage_item"
	CreateConversationRequestTopicResourceTypeAccountUsageResponse                 CreateConversationRequestTopicResourceType = "account_usage_response"
	CreateConversationRequestTopicResourceTypeSubscriptionInfo                     CreateConversationRequestTopicResourceType = "subscription_info"
	CreateConversationRequestTopicResourceTypeBillingPortalSessionResponse         CreateConversationRequestTopicResourceType = "billing_portal_session_response"
	CreateConversationRequestTopicResourceTypeSwitchPlanResponse                   CreateConversationRequestTopicResourceType = "switch_plan_response"
	CreateConversationRequestTopicResourceTypeEnsureBillingCustomerResponse        CreateConversationRequestTopicResourceType = "ensure_billing_customer_response"
	CreateConversationRequestTopicResourceTypeSpendingCapResponse                  CreateConversationRequestTopicResourceType = "spending_cap_response"
	CreateConversationRequestTopicResourceTypeAgentSpendInfo                       CreateConversationRequestTopicResourceType = "agent_spend_info"
	CreateConversationRequestTopicResourceTypeWebhookResponse                      CreateConversationRequestTopicResourceType = "webhook_response"
	CreateConversationRequestTopicResourceTypeAddressSuggestion                    CreateConversationRequestTopicResourceType = "address_suggestion"
	CreateConversationRequestTopicResourceTypeAddressComponents                    CreateConversationRequestTopicResourceType = "address_components"
	CreateConversationRequestTopicResourceTypeAddressDetailsResult                 CreateConversationRequestTopicResourceType = "address_details_result"
	CreateConversationRequestTopicResourceTypeValidatedAddress                     CreateConversationRequestTopicResourceType = "validated_address"
	CreateConversationRequestTopicResourceTypePlanLimit                            CreateConversationRequestTopicResourceType = "plan_limit"
	CreateConversationRequestTopicResourceTypePlanChangeProration                  CreateConversationRequestTopicResourceType = "plan_change_proration"
	CreateConversationRequestTopicResourceTypePlanChangeLineItem                   CreateConversationRequestTopicResourceType = "plan_change_line_item"
	CreateConversationRequestTopicResourceTypeSetupBillingResponse                 CreateConversationRequestTopicResourceType = "setup_billing_response"
	CreateConversationRequestTopicResourceTypeConfirmPaymentResponse               CreateConversationRequestTopicResourceType = "confirm_payment_response"
	CreateConversationRequestTopicResourceTypeOAuthResponse                        CreateConversationRequestTopicResourceType = "oauth_response"
	CreateConversationRequestTopicResourceTypeOAuthStatusResponse                  CreateConversationRequestTopicResourceType = "oauth_status_response"
	CreateConversationRequestTopicResourceTypeStripePublishableKey                 CreateConversationRequestTopicResourceType = "stripe_publishable_key"
	CreateConversationRequestTopicResourceTypeStripeStatus                         CreateConversationRequestTopicResourceType = "stripe_status"
	CreateConversationRequestTopicResourceTypeHealthcheck                          CreateConversationRequestTopicResourceType = "healthcheck"
	CreateConversationRequestTopicResourceTypeAgentDefinitionConfig                CreateConversationRequestTopicResourceType = "agent_definition_config"
	CreateConversationRequestTopicResourceTypeTriggerConfig                        CreateConversationRequestTopicResourceType = "trigger_config"
	CreateConversationRequestTopicResourceTypeCustomerContactInfo                  CreateConversationRequestTopicResourceType = "customer_contact_info"
	CreateConversationRequestTopicResourceTypeCustomerFreightPreferences           CreateConversationRequestTopicResourceType = "customer_freight_preferences"
	CreateConversationRequestTopicResourceTypeCustomerDefaults                     CreateConversationRequestTopicResourceType = "customer_defaults"
	CreateConversationRequestTopicResourceTypeCustomerLeadTime                     CreateConversationRequestTopicResourceType = "customer_lead_time"
	CreateConversationRequestTopicResourceTypeCustomerNotificationPreferences      CreateConversationRequestTopicResourceType = "customer_notification_preferences"
	CreateConversationRequestTopicResourceTypeOrderNotificationRecipient           CreateConversationRequestTopicResourceType = "order_notification_recipient"
	CreateConversationRequestTopicResourceTypeOrderDiscount                        CreateConversationRequestTopicResourceType = "order_discount"
	CreateConversationRequestTopicResourceTypeSalesOrderLine                       CreateConversationRequestTopicResourceType = "sales_order_line"
	CreateConversationRequestTopicResourceTypeSalesOrderType                       CreateConversationRequestTopicResourceType = "sales_order_type"
	CreateConversationRequestTopicResourceTypeSalesOrderStatus                     CreateConversationRequestTopicResourceType = "sales_order_status"
	CreateConversationRequestTopicResourceTypeMaterial                             CreateConversationRequestTopicResourceType = "material"
	CreateConversationRequestTopicResourceTypeSupplierMaterial                     CreateConversationRequestTopicResourceType = "supplier_material"
	CreateConversationRequestTopicResourceTypePart                                 CreateConversationRequestTopicResourceType = "part"
	CreateConversationRequestTopicResourceTypePermissionGroup                      CreateConversationRequestTopicResourceType = "permission_group"
	CreateConversationRequestTopicResourceTypePermission                           CreateConversationRequestTopicResourceType = "permission"
	CreateConversationRequestTopicResourceTypePick                                 CreateConversationRequestTopicResourceType = "pick"
	CreateConversationRequestTopicResourceTypePickLine                             CreateConversationRequestTopicResourceType = "pick_line"
	CreateConversationRequestTopicResourceTypeProductType                          CreateConversationRequestTopicResourceType = "product_type"
	CreateConversationRequestTopicResourceTypeProduction                           CreateConversationRequestTopicResourceType = "production"
	CreateConversationRequestTopicResourceTypeProductionFlow                       CreateConversationRequestTopicResourceType = "production_flow"
	CreateConversationRequestTopicResourceTypeMap                                  CreateConversationRequestTopicResourceType = "map"
	CreateConversationRequestTopicResourceTypePurchaseOrder                        CreateConversationRequestTopicResourceType = "purchase_order"
	CreateConversationRequestTopicResourceTypePurchaseOrderLine                    CreateConversationRequestTopicResourceType = "purchase_order_line"
	CreateConversationRequestTopicResourceTypeSupplier                             CreateConversationRequestTopicResourceType = "supplier"
	CreateConversationRequestTopicResourceTypeSupplierSummary                      CreateConversationRequestTopicResourceType = "supplier_summary"
	CreateConversationRequestTopicResourceTypeReceivableEntry                      CreateConversationRequestTopicResourceType = "receivable_entry"
	CreateConversationRequestTopicResourceTypeReceivingOrder                       CreateConversationRequestTopicResourceType = "receiving_order"
	CreateConversationRequestTopicResourceTypeReceivingOrderLine                   CreateConversationRequestTopicResourceType = "receiving_order_line"
	CreateConversationRequestTopicResourceTypeEmailContact                         CreateConversationRequestTopicResourceType = "email_contact"
	CreateConversationRequestTopicResourceTypeAllocationEntry                      CreateConversationRequestTopicResourceType = "allocation_entry"
	CreateConversationRequestTopicResourceTypeOpenCreditEntry                      CreateConversationRequestTopicResourceType = "open_credit_entry"
	CreateConversationRequestTopicResourceTypeVolumeDiscount                       CreateConversationRequestTopicResourceType = "volume_discount"
	CreateConversationRequestTopicResourceTypeVolumeDiscountTier                   CreateConversationRequestTopicResourceType = "volume_discount_tier"
	CreateConversationRequestTopicResourceTypeAnalyzeDeliveriesResponse            CreateConversationRequestTopicResourceType = "analyze_deliveries_response"
	CreateConversationRequestTopicResourceTypeAnalyzeManufacturingResponse         CreateConversationRequestTopicResourceType = "analyze_manufacturing_response"
	CreateConversationRequestTopicResourceTypeAnalyzeManufacturingBatchResponse    CreateConversationRequestTopicResourceType = "analyze_manufacturing_batch_response"
	CreateConversationRequestTopicResourceTypeAnalyzeQuarterlyOrdersResponse       CreateConversationRequestTopicResourceType = "analyze_quarterly_orders_response"
	CreateConversationRequestTopicResourceTypeAnalyzeNewCustomersResponse          CreateConversationRequestTopicResourceType = "analyze_new_customers_response"
	CreateConversationRequestTopicResourceTypeAnalyzeDemandForecastResponse        CreateConversationRequestTopicResourceType = "analyze_demand_forecast_response"
	CreateConversationRequestTopicResourceTypeAnalyzeOeeResponse                   CreateConversationRequestTopicResourceType = "analyze_oee_response"
	CreateConversationRequestTopicResourceTypeAnalyzeOeeTrendResponse              CreateConversationRequestTopicResourceType = "analyze_oee_trend_response"
	CreateConversationRequestTopicResourceTypeAnalyzeScheduleAttainmentResponse    CreateConversationRequestTopicResourceType = "analyze_schedule_attainment_response"
	CreateConversationRequestTopicResourceTypeCatalogProductLine                   CreateConversationRequestTopicResourceType = "catalog_product_line"
	CreateConversationRequestTopicResourceTypeCatalogCategory                      CreateConversationRequestTopicResourceType = "catalog_category"
	CreateConversationRequestTopicResourceTypeCatalogProduct                       CreateConversationRequestTopicResourceType = "catalog_product"
	CreateConversationRequestTopicResourceTypeCatalogProperty                      CreateConversationRequestTopicResourceType = "catalog_property"
	CreateConversationRequestTopicResourceTypeCatalogAttribute                     CreateConversationRequestTopicResourceType = "catalog_attribute"
	CreateConversationRequestTopicResourceTypeDcLocation                           CreateConversationRequestTopicResourceType = "dc_location"
	CreateConversationRequestTopicResourceTypeEdiRun                               CreateConversationRequestTopicResourceType = "edi_run"
	CreateConversationRequestTopicResourceTypeInventoryItem                        CreateConversationRequestTopicResourceType = "inventory_item"
	CreateConversationRequestTopicResourceTypeAnalyzeWeeksOfSalesResponse          CreateConversationRequestTopicResourceType = "analyze_weeks_of_sales_response"
	CreateConversationRequestTopicResourceTypeBulkReconcileItemsResponse           CreateConversationRequestTopicResourceType = "bulk_reconcile_items_response"
	CreateConversationRequestTopicResourceTypeSysProperty                          CreateConversationRequestTopicResourceType = "sys_property"
	CreateConversationRequestTopicResourceTypeSysPropertyType                      CreateConversationRequestTopicResourceType = "sys_property_type"
	CreateConversationRequestTopicResourceTypeSysPropertyValue                     CreateConversationRequestTopicResourceType = "sys_property_value"
	CreateConversationRequestTopicResourceTypeTerritory                            CreateConversationRequestTopicResourceType = "territory"
	CreateConversationRequestTopicResourceTypeTenancy                              CreateConversationRequestTopicResourceType = "tenancy"
	CreateConversationRequestTopicResourceTypeCheckoutSession                      CreateConversationRequestTopicResourceType = "checkout_session"
	CreateConversationRequestTopicResourceTypeEstimateRateResult                   CreateConversationRequestTopicResourceType = "estimate_rate_result"
	CreateConversationRequestTopicResourceTypeRateShopOption                       CreateConversationRequestTopicResourceType = "rate_shop_option"
	CreateConversationRequestTopicResourceTypeRateShopResult                       CreateConversationRequestTopicResourceType = "rate_shop_result"
	CreateConversationRequestTopicResourceTypeOwner                                CreateConversationRequestTopicResourceType = "owner"
	CreateConversationRequestTopicResourceTypeCreatedBy                            CreateConversationRequestTopicResourceType = "created_by"
	CreateConversationRequestTopicResourceTypeMessage                              CreateConversationRequestTopicResourceType = "message"
	CreateConversationRequestTopicResourceTypeAccountPhotoUploadResult             CreateConversationRequestTopicResourceType = "account_photo_upload_result"
	CreateConversationRequestTopicResourceTypeUserPhotoUploadResult                CreateConversationRequestTopicResourceType = "user_photo_upload_result"
	CreateConversationRequestTopicResourceTypeUserPhotoURL                         CreateConversationRequestTopicResourceType = "user_photo_url"
	CreateConversationRequestTopicResourceTypeBatchLot                             CreateConversationRequestTopicResourceType = "batch_lot"
	CreateConversationRequestTopicResourceTypeCheckDuplicateResult                 CreateConversationRequestTopicResourceType = "check_duplicate_result"
	CreateConversationRequestTopicResourceTypeItemTrendPoint                       CreateConversationRequestTopicResourceType = "item_trend_point"
	CreateConversationRequestTopicResourceTypeTenancyPendingRegistration           CreateConversationRequestTopicResourceType = "tenancy_pending_registration"
	CreateConversationRequestTopicResourceTypeInvoiceAllocationEntry               CreateConversationRequestTopicResourceType = "invoice_allocation_entry"
	CreateConversationRequestTopicResourceTypeAllocationCustomer                   CreateConversationRequestTopicResourceType = "allocation_customer"
	CreateConversationRequestTopicResourceTypeCheckoutSalesOrder                   CreateConversationRequestTopicResourceType = "checkout_sales_order"
	CreateConversationRequestTopicResourceTypeSalesOrderPriceQuote                 CreateConversationRequestTopicResourceType = "sales_order_price_quote"
	CreateConversationRequestTopicResourceTypeSalesOrderFreightQuote               CreateConversationRequestTopicResourceType = "sales_order_freight_quote"
	CreateConversationRequestTopicResourceTypeSalesOrderCommitmentQuote            CreateConversationRequestTopicResourceType = "sales_order_commitment_quote"
	CreateConversationRequestTopicResourceTypeOperatingCalendar                    CreateConversationRequestTopicResourceType = "operating_calendar"
	CreateConversationRequestTopicResourceTypeOperatingCalendarClosure             CreateConversationRequestTopicResourceType = "operating_calendar_closure"
	CreateConversationRequestTopicResourceTypeSalesOrderPriceQuoteLine             CreateConversationRequestTopicResourceType = "sales_order_price_quote_line"
	CreateConversationRequestTopicResourceTypeHubspotSyncJob                       CreateConversationRequestTopicResourceType = "hubspot_sync_job"
	CreateConversationRequestTopicResourceTypeHubspotSyncReport                    CreateConversationRequestTopicResourceType = "hubspot_sync_report"
	CreateConversationRequestTopicResourceTypeHubspotCompanyReview                 CreateConversationRequestTopicResourceType = "hubspot_company_review"
	CreateConversationRequestTopicResourceTypeHubspotCompanyCandidate              CreateConversationRequestTopicResourceType = "hubspot_company_candidate"
	CreateConversationRequestTopicResourceTypeHubspotSyncRecord                    CreateConversationRequestTopicResourceType = "hubspot_sync_record"
	CreateConversationRequestTopicResourceTypeContactMatch                         CreateConversationRequestTopicResourceType = "contact_match"
	CreateConversationRequestTopicResourceTypeReplyDraft                           CreateConversationRequestTopicResourceType = "reply_draft"
	CreateConversationRequestTopicResourceTypeConversationLink                     CreateConversationRequestTopicResourceType = "conversation_link"
	CreateConversationRequestTopicResourceTypeMessagingGroup                       CreateConversationRequestTopicResourceType = "messaging_group"
	CreateConversationRequestTopicResourceTypeMessagingGroupMember                 CreateConversationRequestTopicResourceType = "messaging_group_member"
	CreateConversationRequestTopicResourceTypePortalProfile                        CreateConversationRequestTopicResourceType = "portal_profile"
	CreateConversationRequestTopicResourceTypePortalRegistrationSession            CreateConversationRequestTopicResourceType = "portal_registration_session"
	CreateConversationRequestTopicResourceTypePortalRegistrationSessionData        CreateConversationRequestTopicResourceType = "portal_registration_session_data"
	CreateConversationRequestTopicResourceTypePackList                             CreateConversationRequestTopicResourceType = "pack_list"
	CreateConversationRequestTopicResourceTypePackListParty                        CreateConversationRequestTopicResourceType = "pack_list_party"
	CreateConversationRequestTopicResourceTypePackListLineItem                     CreateConversationRequestTopicResourceType = "pack_list_line_item"
	CreateConversationRequestTopicResourceTypePackListBackOrder                    CreateConversationRequestTopicResourceType = "pack_list_back_order"
	CreateConversationRequestTopicResourceTypePackListCase                         CreateConversationRequestTopicResourceType = "pack_list_case"
	CreateConversationRequestTopicResourceTypeJob                                  CreateConversationRequestTopicResourceType = "job"
	CreateConversationRequestTopicResourceTypeJobResult                            CreateConversationRequestTopicResourceType = "job_result"
	CreateConversationRequestTopicResourceTypeJobExport                            CreateConversationRequestTopicResourceType = "job_export"
	CreateConversationRequestTopicResourceTypeAnalyzeCustomerPricingResponse       CreateConversationRequestTopicResourceType = "analyze_customer_pricing_response"
	CreateConversationRequestTopicResourceTypeCustomerPricingFinding               CreateConversationRequestTopicResourceType = "customer_pricing_finding"
	CreateConversationRequestTopicResourceTypeCustomerPricingSummary               CreateConversationRequestTopicResourceType = "customer_pricing_summary"
	CreateConversationRequestTopicResourceTypeComputedRate                         CreateConversationRequestTopicResourceType = "computed_rate"
	CreateConversationRequestTopicResourceTypeComputedQuantity                     CreateConversationRequestTopicResourceType = "computed_quantity"
	CreateConversationRequestTopicResourceTypeAnalyzeRealizedMarginsResponse       CreateConversationRequestTopicResourceType = "analyze_realized_margins_response"
	CreateConversationRequestTopicResourceTypeRealizedMarginFinding                CreateConversationRequestTopicResourceType = "realized_margin_finding"
	CreateConversationRequestTopicResourceTypeRealizedMarginSummary                CreateConversationRequestTopicResourceType = "realized_margin_summary"
	CreateConversationRequestTopicResourceTypeShipmentRelated                      CreateConversationRequestTopicResourceType = "shipment_related"
	CreateConversationRequestTopicResourceTypeInvoiceRelated                       CreateConversationRequestTopicResourceType = "invoice_related"
	CreateConversationRequestTopicResourceTypePickRelated                          CreateConversationRequestTopicResourceType = "pick_related"
	CreateConversationRequestTopicResourceTypePickShipmentsResponse                CreateConversationRequestTopicResourceType = "pick_shipments_response"
	CreateConversationRequestTopicResourceTypePickTotals                           CreateConversationRequestTopicResourceType = "pick_totals"
	CreateConversationRequestTopicResourceTypePickStageTotal                       CreateConversationRequestTopicResourceType = "pick_stage_total"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListConversation struct {
	// Resources in this page.
	Data []Conversation `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConversationObject `json:"object" api:"required"`
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
func (r ListConversation) RawJSON() string { return r.JSON.raw }
func (r *ListConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConversationObject string

const (
	ListConversationObjectList ListConversationObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListConversationParticipant struct {
	// Resources in this page.
	Data []ConversationParticipant `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConversationParticipantObject `json:"object" api:"required"`
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
func (r ListConversationParticipant) RawJSON() string { return r.JSON.raw }
func (r *ListConversationParticipant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConversationParticipantObject string

const (
	ListConversationParticipantObjectList ListConversationParticipantObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMessageAttachment struct {
	// Resources in this page.
	Data []MessageAttachment `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessageAttachmentObject `json:"object" api:"required"`
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
func (r ListMessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *ListMessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessageAttachmentObject string

const (
	ListMessageAttachmentObjectList ListMessageAttachmentObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMessagingGroupMember struct {
	// Resources in this page.
	Data []MessagingGroupMember `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessagingGroupMemberObject `json:"object" api:"required"`
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
func (r ListMessagingGroupMember) RawJSON() string { return r.JSON.raw }
func (r *ListMessagingGroupMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessagingGroupMemberObject string

const (
	ListMessagingGroupMemberObjectList ListMessagingGroupMemberObject = "list"
)

// A chat message within a conversation.
//
// One resource covers every stage of a message's life: a delivered timeline
// message, a message queued for a future send, and a customer-reply draft awaiting
// approval. Read `status` to tell them apart.
type Message struct {
	// Message ID.
	ID string `json:"id" api:"required"`
	// Machine-readable reason an agent reply failed.
	//
	// A client can react to the specific code rather than just showing the body —
	// `agent_spending_cap_reached`, for example, is a cue to offer raising the agent
	// spending limit.
	//
	// Any of "expired_token", "api_key_expired", "api_key_revoked",
	// "invalid_credentials", "insufficient_permissions", "payment_required",
	// "agent_spending_cap_reached", "validation_failed", "missing_field",
	// "invalid_format", "method_not_allowed", "resource_not_found", "resource_exists",
	// "resource_conflict", "resource_gone", "idempotency_in_progress",
	// "limit_exceeded", "registration_closed", "rate_limit_exceeded",
	// "parameter_missing", "parameter_invalid", "parameter_unknown",
	// "parameters_exclusive", "internal_error", "service_unavailable",
	// "external_service_error", "timeout", "connection_error", "request_timeout",
	// "client_closed_request", "api_version_required", "api_version_invalid",
	// "api_version_too_old".
	AgentErrorCode MessageAgentErrorCode `json:"agent_error_code" api:"required"`
	// A single execution of an agent, from trigger through completion.
	AgentRun AgentRun `json:"agent_run" api:"required"`
	// Whether this message is an agent reply reporting that the agent's run failed.
	//
	// The body explains the failure to the reader rather than answering the request.
	AgentRunFailed bool `json:"agent_run_failed" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Attachments ListMessageAttachment `json:"attachments" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Author Actor `json:"author" api:"required"`
	// Message body.
	//
	// A message made up of nothing but attachments or a linked record carries no body,
	// and a deleted message has its body cleared.
	Body string `json:"body" api:"required"`
	// How the message reached its audience, or how a draft will be sent once it is
	// approved.
	//
	// - `message`: appears in the conversation itself.
	// - `email`: goes out as email on the thread of the inbox the case is bridged to.
	//
	// Any of "message", "email".
	Channel MessageChannel `json:"channel" api:"required"`
	// The dedupe key the client supplied when sending, echoed back so an optimistic
	// local copy can be matched to the stored message.
	ClientMessageID string `json:"client_message_id" api:"required"`
	// A conversation thread the caller participates in.
	Conversation *Conversation `json:"conversation" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the message was deleted.
	//
	// A deleted message keeps its place in the timeline with its body cleared, so
	// surrounding ordering and replies stay intact.
	DeletedAt time.Time `json:"deleted_at" api:"required" format:"date-time"`
	// When the message was last edited.
	EditedAt time.Time `json:"edited_at" api:"required" format:"date-time"`
	// What this message represents.
	//
	//   - `chat`: written by a person.
	//   - `system_event`: a record of something that happened in the conversation, such
	//     as someone joining or a record being linked.
	//   - `agent`: written by an AI agent taking part in the conversation.
	//   - `scheduled`: came from a send queued ahead of time.
	//   - `alert`: an automated alert surfaced in the conversation.
	//   - `email`: a message carried over the case's bridged email thread, either one
	//     that arrived from the customer or a reply sent back out to them.
	//
	// Any of "chat", "system_event", "agent", "scheduled", "alert", "email".
	Kind MessageKind `json:"kind" api:"required"`
	// Resource type identifier.
	//
	// Any of "chat_message".
	Object MessageObject `json:"object" api:"required"`
	// A chat message within a conversation.
	//
	// One resource covers every stage of a message's life: a delivered timeline
	// message, a message queued for a future send, and a customer-reply draft awaiting
	// approval. Read `status` to tell them apart.
	ReplyTo *Message `json:"reply_to" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// When a message queued for a future send is due to go out.
	ScheduledAt time.Time `json:"scheduled_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Sender Actor `json:"sender" api:"required"`
	// The message's position in the conversation timeline, counting up from the first
	// message.
	//
	// A sequence is assigned only when a message is delivered, so a draft or a
	// not-yet-sent scheduled message reports `0`. Listing a conversation's messages
	// pages backwards through this ordering.
	Sequence int64 `json:"sequence" api:"required"`
	// Where the message stands in its life.
	//
	//   - `draft`: a proposed reply to the customer, still editable and waiting for
	//     approval before anyone outside sees it.
	//   - `scheduled`: queued to go out at a future time.
	//   - `sent`: delivered, and part of the conversation everyone reads.
	//   - `canceled`: a scheduled message stopped before it went out.
	//   - `rejected`: a draft discarded instead of being sent.
	//   - `failed`: a scheduled message that could not be delivered.
	//   - `superseded`: a draft replaced by a newer one for the same thread.
	//
	// Only a `sent` message occupies a place in the conversation; the others are
	// records of messages that never reached it.
	//
	// Any of "draft", "scheduled", "sent", "canceled", "rejected", "failed",
	// "superseded".
	Status MessageStatus `json:"status" api:"required"`
	// The streaming state of an agent reply.
	//
	// `streaming` means the body is still being generated and keeps growing as
	// realtime updates arrive; `complete` means it is final.
	//
	// Any of "streaming", "complete".
	StreamingState MessageStreamingState `json:"streaming_state" api:"required"`
	// The email subject line.
	//
	// On an email-bridged case, this is the subject of the inbound email, or the
	// subject a customer reply is sent out with.
	Subject string `json:"subject" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Who can see this message.
	//
	//   - `internal`: a note only your team can see.
	//   - `external`: sent to or received from an outside party, such as the customer on
	//     a support case, and part of the official record of that exchange.
	//   - `system`: an event both your team and the customer see.
	//
	// A customer reading their own case is never served `internal` messages.
	//
	// Any of "internal", "external", "system".
	Visibility MessageVisibility `json:"visibility" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AgentErrorCode  respjson.Field
		AgentRun        respjson.Field
		AgentRunFailed  respjson.Field
		Attachments     respjson.Field
		Author          respjson.Field
		Body            respjson.Field
		Channel         respjson.Field
		ClientMessageID respjson.Field
		Conversation    respjson.Field
		CreatedAt       respjson.Field
		DeletedAt       respjson.Field
		EditedAt        respjson.Field
		Kind            respjson.Field
		Object          respjson.Field
		ReplyTo         respjson.Field
		Resource        respjson.Field
		ScheduledAt     respjson.Field
		Sender          respjson.Field
		Sequence        respjson.Field
		Status          respjson.Field
		StreamingState  respjson.Field
		Subject         respjson.Field
		UpdatedAt       respjson.Field
		Visibility      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable reason an agent reply failed.
//
// A client can react to the specific code rather than just showing the body —
// `agent_spending_cap_reached`, for example, is a cue to offer raising the agent
// spending limit.
type MessageAgentErrorCode string

const (
	MessageAgentErrorCodeExpiredToken            MessageAgentErrorCode = "expired_token"
	MessageAgentErrorCodeAPIKeyExpired           MessageAgentErrorCode = "api_key_expired"
	MessageAgentErrorCodeAPIKeyRevoked           MessageAgentErrorCode = "api_key_revoked"
	MessageAgentErrorCodeInvalidCredentials      MessageAgentErrorCode = "invalid_credentials"
	MessageAgentErrorCodeInsufficientPermissions MessageAgentErrorCode = "insufficient_permissions"
	MessageAgentErrorCodePaymentRequired         MessageAgentErrorCode = "payment_required"
	MessageAgentErrorCodeAgentSpendingCapReached MessageAgentErrorCode = "agent_spending_cap_reached"
	MessageAgentErrorCodeValidationFailed        MessageAgentErrorCode = "validation_failed"
	MessageAgentErrorCodeMissingField            MessageAgentErrorCode = "missing_field"
	MessageAgentErrorCodeInvalidFormat           MessageAgentErrorCode = "invalid_format"
	MessageAgentErrorCodeMethodNotAllowed        MessageAgentErrorCode = "method_not_allowed"
	MessageAgentErrorCodeResourceNotFound        MessageAgentErrorCode = "resource_not_found"
	MessageAgentErrorCodeResourceExists          MessageAgentErrorCode = "resource_exists"
	MessageAgentErrorCodeResourceConflict        MessageAgentErrorCode = "resource_conflict"
	MessageAgentErrorCodeResourceGone            MessageAgentErrorCode = "resource_gone"
	MessageAgentErrorCodeIdempotencyInProgress   MessageAgentErrorCode = "idempotency_in_progress"
	MessageAgentErrorCodeLimitExceeded           MessageAgentErrorCode = "limit_exceeded"
	MessageAgentErrorCodeRegistrationClosed      MessageAgentErrorCode = "registration_closed"
	MessageAgentErrorCodeRateLimitExceeded       MessageAgentErrorCode = "rate_limit_exceeded"
	MessageAgentErrorCodeParameterMissing        MessageAgentErrorCode = "parameter_missing"
	MessageAgentErrorCodeParameterInvalid        MessageAgentErrorCode = "parameter_invalid"
	MessageAgentErrorCodeParameterUnknown        MessageAgentErrorCode = "parameter_unknown"
	MessageAgentErrorCodeParametersExclusive     MessageAgentErrorCode = "parameters_exclusive"
	MessageAgentErrorCodeInternalError           MessageAgentErrorCode = "internal_error"
	MessageAgentErrorCodeServiceUnavailable      MessageAgentErrorCode = "service_unavailable"
	MessageAgentErrorCodeExternalServiceError    MessageAgentErrorCode = "external_service_error"
	MessageAgentErrorCodeTimeout                 MessageAgentErrorCode = "timeout"
	MessageAgentErrorCodeConnectionError         MessageAgentErrorCode = "connection_error"
	MessageAgentErrorCodeRequestTimeout          MessageAgentErrorCode = "request_timeout"
	MessageAgentErrorCodeClientClosedRequest     MessageAgentErrorCode = "client_closed_request"
	MessageAgentErrorCodeAPIVersionRequired      MessageAgentErrorCode = "api_version_required"
	MessageAgentErrorCodeAPIVersionInvalid       MessageAgentErrorCode = "api_version_invalid"
	MessageAgentErrorCodeAPIVersionTooOld        MessageAgentErrorCode = "api_version_too_old"
)

// How the message reached its audience, or how a draft will be sent once it is
// approved.
//
// - `message`: appears in the conversation itself.
// - `email`: goes out as email on the thread of the inbox the case is bridged to.
type MessageChannel string

const (
	MessageChannelMessage MessageChannel = "message"
	MessageChannelEmail   MessageChannel = "email"
)

// What this message represents.
//
//   - `chat`: written by a person.
//   - `system_event`: a record of something that happened in the conversation, such
//     as someone joining or a record being linked.
//   - `agent`: written by an AI agent taking part in the conversation.
//   - `scheduled`: came from a send queued ahead of time.
//   - `alert`: an automated alert surfaced in the conversation.
//   - `email`: a message carried over the case's bridged email thread, either one
//     that arrived from the customer or a reply sent back out to them.
type MessageKind string

const (
	MessageKindChat        MessageKind = "chat"
	MessageKindSystemEvent MessageKind = "system_event"
	MessageKindAgent       MessageKind = "agent"
	MessageKindScheduled   MessageKind = "scheduled"
	MessageKindAlert       MessageKind = "alert"
	MessageKindEmail       MessageKind = "email"
)

// Resource type identifier.
type MessageObject string

const (
	MessageObjectChatMessage MessageObject = "chat_message"
)

// Where the message stands in its life.
//
//   - `draft`: a proposed reply to the customer, still editable and waiting for
//     approval before anyone outside sees it.
//   - `scheduled`: queued to go out at a future time.
//   - `sent`: delivered, and part of the conversation everyone reads.
//   - `canceled`: a scheduled message stopped before it went out.
//   - `rejected`: a draft discarded instead of being sent.
//   - `failed`: a scheduled message that could not be delivered.
//   - `superseded`: a draft replaced by a newer one for the same thread.
//
// Only a `sent` message occupies a place in the conversation; the others are
// records of messages that never reached it.
type MessageStatus string

const (
	MessageStatusDraft      MessageStatus = "draft"
	MessageStatusScheduled  MessageStatus = "scheduled"
	MessageStatusSent       MessageStatus = "sent"
	MessageStatusCanceled   MessageStatus = "canceled"
	MessageStatusRejected   MessageStatus = "rejected"
	MessageStatusFailed     MessageStatus = "failed"
	MessageStatusSuperseded MessageStatus = "superseded"
)

// The streaming state of an agent reply.
//
// `streaming` means the body is still being generated and keeps growing as
// realtime updates arrive; `complete` means it is final.
type MessageStreamingState string

const (
	MessageStreamingStateStreaming MessageStreamingState = "streaming"
	MessageStreamingStateComplete  MessageStreamingState = "complete"
)

// Who can see this message.
//
//   - `internal`: a note only your team can see.
//   - `external`: sent to or received from an outside party, such as the customer on
//     a support case, and part of the official record of that exchange.
//   - `system`: an event both your team and the customer see.
//
// A customer reading their own case is never served `internal` messages.
type MessageVisibility string

const (
	MessageVisibilityInternal MessageVisibility = "internal"
	MessageVisibilityExternal MessageVisibility = "external"
	MessageVisibilitySystem   MessageVisibility = "system"
)

// A file, image, link, or resource attached to a message.
type MessageAttachment struct {
	// Attachment ID.
	ID string `json:"id" api:"required"`
	// The MIME type of the uploaded content.
	//
	// Carried only by `file` and `image` attachments.
	ContentType string `json:"content_type" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The filename the attachment was uploaded under.
	//
	// Carried only by `file` and `image` attachments.
	Filename string `json:"filename" api:"required"`
	// The kind of attachment, which determines how it is stored and which of the
	// fields below are populated.
	//
	// - `file`: an uploaded non-image file.
	// - `image`: an uploaded image.
	// - `link`: an external URL reference, with no stored file.
	// - `resource`: a reference to an in-app resource, such as an order.
	//
	// Any of "file", "image", "link", "resource".
	Kind MessageAttachmentKind `json:"kind" api:"required"`
	// Resource type identifier.
	//
	// Any of "message_attachment".
	Object MessageAttachmentObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// The size of the uploaded content in bytes.
	//
	// Carried only by `file` and `image` attachments, and only when the sender
	// supplied it with the message.
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Where to fetch the attachment: a signed download URL for `file` and `image`
	// attachments, or the target address for `link` attachments.
	//
	// Download URLs are signed for one hour and regenerated each time the message is
	// read, so follow the URL promptly instead of persisting it. `resource`
	// attachments have no URL — use `resource` to resolve them.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		CreatedAt   respjson.Field
		Filename    respjson.Field
		Kind        respjson.Field
		Object      respjson.Field
		Resource    respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessageAttachment) RawJSON() string { return r.JSON.raw }
func (r *MessageAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of attachment, which determines how it is stored and which of the
// fields below are populated.
//
// - `file`: an uploaded non-image file.
// - `image`: an uploaded image.
// - `link`: an external URL reference, with no stored file.
// - `resource`: a reference to an in-app resource, such as an order.
type MessageAttachmentKind string

const (
	MessageAttachmentKindFile     MessageAttachmentKind = "file"
	MessageAttachmentKindImage    MessageAttachmentKind = "image"
	MessageAttachmentKindLink     MessageAttachmentKind = "link"
	MessageAttachmentKindResource MessageAttachmentKind = "resource"
)

// Resource type identifier.
type MessageAttachmentObject string

const (
	MessageAttachmentObjectMessageAttachment MessageAttachmentObject = "message_attachment"
)

// A reusable roster: a named set of members (users and/or agents) that seeds new
// conversations.
//
// Starting a conversation from a group snapshots its current members into that
// conversation, so the same group can back many conversations (each with its own
// title); later edits to the group never change conversations already created from
// it.
type MessagingGroup struct {
	// Messaging group ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Members ListMessagingGroupMember `json:"members" api:"required"`
	// The roster's display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "messaging_group".
	Object MessagingGroupObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Members     respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingGroup) RawJSON() string { return r.JSON.raw }
func (r *MessagingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MessagingGroupObject string

const (
	MessagingGroupObjectMessagingGroup MessagingGroupObject = "messaging_group"
)

// A member of a reusable roster: either a user or an agent, represented by its
// actor.
type MessagingGroupMember struct {
	// Membership ID.
	//
	// This identifies the member's place on the roster, not the user or agent
	// themselves; it is the id to pass when removing them from the roster.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Resource type identifier.
	//
	// Any of "messaging_group_member".
	Object MessagingGroupMemberObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actor       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingGroupMember) RawJSON() string { return r.JSON.raw }
func (r *MessagingGroupMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MessagingGroupMemberObject string

const (
	MessagingGroupMemberObjectMessagingGroupMember MessagingGroupMemberObject = "messaging_group_member"
)

// A participant's read position in a conversation — the basis for read receipts
// ("who has seen this").
type ReadCursor struct {
	// The id of the last message the participant has read.
	MessageID string `json:"message_id" api:"required"`
	// Resource type identifier.
	//
	// Any of "read_cursor".
	Object ReadCursorObject `json:"object" api:"required"`
	// When the participant last advanced their read cursor.
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// The sequence number of the last message the participant has read in the
	// conversation.
	//
	// A message is "seen" by this participant when its `sequence` is `<=` this value.
	// `0` means they have not read any message in the conversation yet.
	Sequence int64 `json:"sequence" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessageID   respjson.Field
		Object      respjson.Field
		ReadAt      respjson.Field
		Sequence    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReadCursor) RawJSON() string { return r.JSON.raw }
func (r *ReadCursor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ReadCursorObject string

const (
	ReadCursorObjectReadCursor ReadCursorObject = "read_cursor"
)

// Request to rename a conversation.
type UpdateConversationRequestParam struct {
	// The group conversation's new display title.
	//
	// Send `null` to clear the title and leave the conversation unnamed.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r UpdateConversationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateConversationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateConversationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingConversationNewParams struct {
	// Request to create a conversation.
	CreateConversationRequest CreateConversationRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateConversationRequest)
}
func (r *MessagingConversationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationNewParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationGetParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to rename a conversation.
	UpdateConversationRequest UpdateConversationRequestParam
	paramObj
}

func (r MessagingConversationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateConversationRequest)
}
func (r *MessagingConversationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationUpdateParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationListParams struct {
	// Filter the support inbox to cases owned by this assignee, an account user or an
	// account group.
	AssigneeResourceID param.Opt[string] `query:"assignee_resource_id,omitzero" json:"-"`
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Return the archived support inbox instead of the working one.
	//
	// This swaps the view rather than widening it: archived cases are returned and
	// unarchived ones are left out.
	IncludeArchived param.Opt[bool] `query:"include_archived,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// The id of the business record, together with `topic_resource_type`.
	TopicResourceID param.Opt[string] `query:"topic_resource_id,omitzero" json:"-"`
	// Restrict the support inbox to cases nobody has been assigned yet.
	Unassigned param.Opt[bool] `query:"unassigned,omitzero" json:"-"`
	// Filter by whether the conversation is team-only or customer-facing.
	//
	//   - `internal`: threads the customer never sees — direct messages, group threads,
	//     and record discussions.
	//   - `customer`: external customer-service cases the customer takes part in, from
	//     the portal or a bridged email thread.
	//
	// Any of "internal", "customer".
	Audience MessagingConversationListParamsAudience `query:"audience,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "assignee", "group", "participants", "topic", "last_message",
	// "last_message.sender", "last_message.author", "last_message.resource",
	// "last_message.attachments", "last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by whether the caller has hidden the conversation from their own list.
	//
	// Any of "active", "hidden".
	Status MessagingConversationListParamsStatus `query:"status,omitzero" json:"-"`
	// Restrict to conversations attached to a business record of this type, together
	// with `topic_resource_id`.
	//
	// Matches both conversations anchored to the record and conversations that merely
	// link it, which is what powers the "discussions on this record" view.
	//
	// Any of "account", "actor", "entity", "record", "freight", "sales_order_totals",
	// "sales_order_stage_total", "sales_order_related", "order_contact", "user",
	// "address", "api_key", "created_api_key", "refresh_token", "list", "sandbox",
	// "registration_session", "pricing_plan", "account_plan", "plan_change",
	// "enterprise_inquiry", "request_log", "audit_event", "audit_field_change",
	// "role", "unit", "account_affiliation", "agent_definition", "available_tool",
	// "agent_definition_tool", "agent_account_status", "agent_run", "agent_action",
	// "agent_run_step", "agent_token_usage", "agent_memory", "notification",
	// "notification_unread_count", "notification_send_result",
	// "notification_unread_summary", "announcement", "conversation", "support_case",
	// "conversation_participant", "read_cursor", "chat_message",
	// "notification_unread_summary_account", "messaging_block",
	// "notification_preference", "message_attachment", "attachment_upload_target",
	// "scheduled_message", "messaging_contact", "message_report", "tool_group",
	// "model", "payment_term", "shipping_term", "quantity", "account_group",
	// "support_route", "support_availability", "account_status", "geolocation",
	// "account_user", "department", "account_integration", "account_price",
	// "product_line", "item_category", "attribute", "rate",
	// "account_group_product_line_access", "sales_target", "adjustment_type",
	// "account_branding", "account_portal", "account_logo_url", "account_favicon_url",
	// "public_account", "property", "carrier", "service_level", "item",
	// "item_lot_default", "item_inventory", "product", "batch", "batch_flow_node",
	// "scanning_consumption", "open_batch_summary", "scanning_production_step_info",
	// "scanning_station", "production_step", "production_run", "machine",
	// "machine_status", "machine_downtime_event", "demand_override",
	// "demand_override_type", "machine_downtime_reason",
	// "production_schedule_preview", "production_schedule_regenerate_preview",
	// "production_schedule", "production_schedule_line",
	// "production_schedule_deviation", "production_schedule_derived_line",
	// "production_schedule_settings", "production_schedule_resource_setting",
	// "production_schedule_item_setting", "fulfillment_recommendation",
	// "analyze_delivery_performance_response", "delivery_performance",
	// "delivery_backlog_bucket", "delivery_lateness_bucket", "delivery_breakdown",
	// "schedule_order_coverage", "schedule_order_coverage_line",
	// "schedule_deviation_type", "schedule_at_risk_order",
	// "production_schedule_finished_policy", "production_schedule_finishing_line",
	// "production_schedule_week_release", "production_schedule_week_release_preview",
	// "production_schedule_item_policy", "child_account", "unit_group",
	// "unit_group_unit", "consumption", "customer_product_line_access", "customer",
	// "frequently_ordered_product", "priority", "delivery", "delivery_line",
	// "sales_order", "location", "location_type", "lot", "email_log", "email_domain",
	// "email_inbox", "portal_domain", "dns_record", "inventory_change_log", "invoice",
	// "invoice_summary", "invoice_line", "invoice_allocation", "invoice_for_payment",
	// "shipment", "shipment_summary", "shipment_line", "shipping_case",
	// "shipping_case_label_url", "settlement", "settlement_summary",
	// "role_permission", "registration_flow", "registration_flow_option",
	// "transaction", "transaction_summary", "transaction_method", "transaction_type",
	// "transaction_allocation", "usage_item", "account_usage_response",
	// "subscription_info", "billing_portal_session_response", "switch_plan_response",
	// "ensure_billing_customer_response", "spending_cap_response", "agent_spend_info",
	// "webhook_response", "address_suggestion", "address_components",
	// "address_details_result", "validated_address", "plan_limit",
	// "plan_change_proration", "plan_change_line_item", "setup_billing_response",
	// "confirm_payment_response", "oauth_response", "oauth_status_response",
	// "stripe_publishable_key", "stripe_status", "healthcheck",
	// "agent_definition_config", "trigger_config", "customer_contact_info",
	// "customer_freight_preferences", "customer_defaults", "customer_lead_time",
	// "customer_notification_preferences", "order_notification_recipient",
	// "order_discount", "sales_order_line", "sales_order_type", "sales_order_status",
	// "material", "supplier_material", "part", "permission_group", "permission",
	// "pick", "pick_line", "product_type", "production", "production_flow", "map",
	// "purchase_order", "purchase_order_line", "supplier", "supplier_summary",
	// "receivable_entry", "receiving_order", "receiving_order_line", "email_contact",
	// "allocation_entry", "open_credit_entry", "volume_discount",
	// "volume_discount_tier", "analyze_deliveries_response",
	// "analyze_manufacturing_response", "analyze_manufacturing_batch_response",
	// "analyze_quarterly_orders_response", "analyze_new_customers_response",
	// "analyze_demand_forecast_response", "analyze_oee_response",
	// "analyze_oee_trend_response", "analyze_schedule_attainment_response",
	// "catalog_product_line", "catalog_category", "catalog_product",
	// "catalog_property", "catalog_attribute", "dc_location", "edi_run",
	// "inventory_item", "analyze_weeks_of_sales_response",
	// "bulk_reconcile_items_response", "sys_property", "sys_property_type",
	// "sys_property_value", "territory", "tenancy", "checkout_session",
	// "estimate_rate_result", "rate_shop_option", "rate_shop_result", "owner",
	// "created_by", "message", "account_photo_upload_result",
	// "user_photo_upload_result", "user_photo_url", "batch_lot",
	// "check_duplicate_result", "item_trend_point", "tenancy_pending_registration",
	// "invoice_allocation_entry", "allocation_customer", "checkout_sales_order",
	// "sales_order_price_quote", "sales_order_freight_quote",
	// "sales_order_commitment_quote", "operating_calendar",
	// "operating_calendar_closure", "sales_order_price_quote_line",
	// "hubspot_sync_job", "hubspot_sync_report", "hubspot_company_review",
	// "hubspot_company_candidate", "hubspot_sync_record", "contact_match",
	// "reply_draft", "conversation_link", "messaging_group", "messaging_group_member",
	// "portal_profile", "portal_registration_session",
	// "portal_registration_session_data", "pack_list", "pack_list_party",
	// "pack_list_line_item", "pack_list_back_order", "pack_list_case", "job",
	// "job_result", "job_export", "analyze_customer_pricing_response",
	// "customer_pricing_finding", "customer_pricing_summary", "computed_rate",
	// "computed_quantity", "analyze_realized_margins_response",
	// "realized_margin_finding", "realized_margin_summary", "shipment_related",
	// "invoice_related", "pick_related", "pick_shipments_response", "pick_totals",
	// "pick_stage_total".
	TopicResourceType MessagingConversationListParamsTopicResourceType `query:"topic_resource_type,omitzero" json:"-"`
	// Filter by conversation type.
	//
	// Any of "direct_message", "group", "system".
	Type MessagingConversationListParamsType `query:"type,omitzero" json:"-"`
	// Filter the support inbox to a single triage lane.
	//
	// - `new`: opened but nobody has triaged it yet.
	// - `open`: actively being worked.
	// - `waiting_internal`: blocked on the internal team.
	// - `waiting_external`: blocked on a reply from the customer.
	// - `needs_approval`: a drafted reply is waiting for a human to approve it.
	// - `resolved`: closed out.
	//
	// The working inbox hides resolved cases unless you ask for this lane explicitly.
	//
	// Any of "new", "open", "waiting_internal", "waiting_external", "needs_approval",
	// "resolved".
	WorkflowStatus MessagingConversationListParamsWorkflowStatus `query:"workflow_status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationListParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by whether the conversation is team-only or customer-facing.
//
//   - `internal`: threads the customer never sees — direct messages, group threads,
//     and record discussions.
//   - `customer`: external customer-service cases the customer takes part in, from
//     the portal or a bridged email thread.
type MessagingConversationListParamsAudience string

const (
	MessagingConversationListParamsAudienceInternal MessagingConversationListParamsAudience = "internal"
	MessagingConversationListParamsAudienceCustomer MessagingConversationListParamsAudience = "customer"
)

// Filter by whether the caller has hidden the conversation from their own list.
type MessagingConversationListParamsStatus string

const (
	MessagingConversationListParamsStatusActive MessagingConversationListParamsStatus = "active"
	MessagingConversationListParamsStatusHidden MessagingConversationListParamsStatus = "hidden"
)

// Restrict to conversations attached to a business record of this type, together
// with `topic_resource_id`.
//
// Matches both conversations anchored to the record and conversations that merely
// link it, which is what powers the "discussions on this record" view.
type MessagingConversationListParamsTopicResourceType string

const (
	MessagingConversationListParamsTopicResourceTypeAccount                              MessagingConversationListParamsTopicResourceType = "account"
	MessagingConversationListParamsTopicResourceTypeActor                                MessagingConversationListParamsTopicResourceType = "actor"
	MessagingConversationListParamsTopicResourceTypeEntity                               MessagingConversationListParamsTopicResourceType = "entity"
	MessagingConversationListParamsTopicResourceTypeRecord                               MessagingConversationListParamsTopicResourceType = "record"
	MessagingConversationListParamsTopicResourceTypeFreight                              MessagingConversationListParamsTopicResourceType = "freight"
	MessagingConversationListParamsTopicResourceTypeSalesOrderTotals                     MessagingConversationListParamsTopicResourceType = "sales_order_totals"
	MessagingConversationListParamsTopicResourceTypeSalesOrderStageTotal                 MessagingConversationListParamsTopicResourceType = "sales_order_stage_total"
	MessagingConversationListParamsTopicResourceTypeSalesOrderRelated                    MessagingConversationListParamsTopicResourceType = "sales_order_related"
	MessagingConversationListParamsTopicResourceTypeOrderContact                         MessagingConversationListParamsTopicResourceType = "order_contact"
	MessagingConversationListParamsTopicResourceTypeUser                                 MessagingConversationListParamsTopicResourceType = "user"
	MessagingConversationListParamsTopicResourceTypeAddress                              MessagingConversationListParamsTopicResourceType = "address"
	MessagingConversationListParamsTopicResourceTypeAPIKey                               MessagingConversationListParamsTopicResourceType = "api_key"
	MessagingConversationListParamsTopicResourceTypeCreatedAPIKey                        MessagingConversationListParamsTopicResourceType = "created_api_key"
	MessagingConversationListParamsTopicResourceTypeRefreshToken                         MessagingConversationListParamsTopicResourceType = "refresh_token"
	MessagingConversationListParamsTopicResourceTypeList                                 MessagingConversationListParamsTopicResourceType = "list"
	MessagingConversationListParamsTopicResourceTypeSandbox                              MessagingConversationListParamsTopicResourceType = "sandbox"
	MessagingConversationListParamsTopicResourceTypeRegistrationSession                  MessagingConversationListParamsTopicResourceType = "registration_session"
	MessagingConversationListParamsTopicResourceTypePricingPlan                          MessagingConversationListParamsTopicResourceType = "pricing_plan"
	MessagingConversationListParamsTopicResourceTypeAccountPlan                          MessagingConversationListParamsTopicResourceType = "account_plan"
	MessagingConversationListParamsTopicResourceTypePlanChange                           MessagingConversationListParamsTopicResourceType = "plan_change"
	MessagingConversationListParamsTopicResourceTypeEnterpriseInquiry                    MessagingConversationListParamsTopicResourceType = "enterprise_inquiry"
	MessagingConversationListParamsTopicResourceTypeRequestLog                           MessagingConversationListParamsTopicResourceType = "request_log"
	MessagingConversationListParamsTopicResourceTypeAuditEvent                           MessagingConversationListParamsTopicResourceType = "audit_event"
	MessagingConversationListParamsTopicResourceTypeAuditFieldChange                     MessagingConversationListParamsTopicResourceType = "audit_field_change"
	MessagingConversationListParamsTopicResourceTypeRole                                 MessagingConversationListParamsTopicResourceType = "role"
	MessagingConversationListParamsTopicResourceTypeUnit                                 MessagingConversationListParamsTopicResourceType = "unit"
	MessagingConversationListParamsTopicResourceTypeAccountAffiliation                   MessagingConversationListParamsTopicResourceType = "account_affiliation"
	MessagingConversationListParamsTopicResourceTypeAgentDefinition                      MessagingConversationListParamsTopicResourceType = "agent_definition"
	MessagingConversationListParamsTopicResourceTypeAvailableTool                        MessagingConversationListParamsTopicResourceType = "available_tool"
	MessagingConversationListParamsTopicResourceTypeAgentDefinitionTool                  MessagingConversationListParamsTopicResourceType = "agent_definition_tool"
	MessagingConversationListParamsTopicResourceTypeAgentAccountStatus                   MessagingConversationListParamsTopicResourceType = "agent_account_status"
	MessagingConversationListParamsTopicResourceTypeAgentRun                             MessagingConversationListParamsTopicResourceType = "agent_run"
	MessagingConversationListParamsTopicResourceTypeAgentAction                          MessagingConversationListParamsTopicResourceType = "agent_action"
	MessagingConversationListParamsTopicResourceTypeAgentRunStep                         MessagingConversationListParamsTopicResourceType = "agent_run_step"
	MessagingConversationListParamsTopicResourceTypeAgentTokenUsage                      MessagingConversationListParamsTopicResourceType = "agent_token_usage"
	MessagingConversationListParamsTopicResourceTypeAgentMemory                          MessagingConversationListParamsTopicResourceType = "agent_memory"
	MessagingConversationListParamsTopicResourceTypeNotification                         MessagingConversationListParamsTopicResourceType = "notification"
	MessagingConversationListParamsTopicResourceTypeNotificationUnreadCount              MessagingConversationListParamsTopicResourceType = "notification_unread_count"
	MessagingConversationListParamsTopicResourceTypeNotificationSendResult               MessagingConversationListParamsTopicResourceType = "notification_send_result"
	MessagingConversationListParamsTopicResourceTypeNotificationUnreadSummary            MessagingConversationListParamsTopicResourceType = "notification_unread_summary"
	MessagingConversationListParamsTopicResourceTypeAnnouncement                         MessagingConversationListParamsTopicResourceType = "announcement"
	MessagingConversationListParamsTopicResourceTypeConversation                         MessagingConversationListParamsTopicResourceType = "conversation"
	MessagingConversationListParamsTopicResourceTypeSupportCase                          MessagingConversationListParamsTopicResourceType = "support_case"
	MessagingConversationListParamsTopicResourceTypeConversationParticipant              MessagingConversationListParamsTopicResourceType = "conversation_participant"
	MessagingConversationListParamsTopicResourceTypeReadCursor                           MessagingConversationListParamsTopicResourceType = "read_cursor"
	MessagingConversationListParamsTopicResourceTypeChatMessage                          MessagingConversationListParamsTopicResourceType = "chat_message"
	MessagingConversationListParamsTopicResourceTypeNotificationUnreadSummaryAccount     MessagingConversationListParamsTopicResourceType = "notification_unread_summary_account"
	MessagingConversationListParamsTopicResourceTypeMessagingBlock                       MessagingConversationListParamsTopicResourceType = "messaging_block"
	MessagingConversationListParamsTopicResourceTypeNotificationPreference               MessagingConversationListParamsTopicResourceType = "notification_preference"
	MessagingConversationListParamsTopicResourceTypeMessageAttachment                    MessagingConversationListParamsTopicResourceType = "message_attachment"
	MessagingConversationListParamsTopicResourceTypeAttachmentUploadTarget               MessagingConversationListParamsTopicResourceType = "attachment_upload_target"
	MessagingConversationListParamsTopicResourceTypeScheduledMessage                     MessagingConversationListParamsTopicResourceType = "scheduled_message"
	MessagingConversationListParamsTopicResourceTypeMessagingContact                     MessagingConversationListParamsTopicResourceType = "messaging_contact"
	MessagingConversationListParamsTopicResourceTypeMessageReport                        MessagingConversationListParamsTopicResourceType = "message_report"
	MessagingConversationListParamsTopicResourceTypeToolGroup                            MessagingConversationListParamsTopicResourceType = "tool_group"
	MessagingConversationListParamsTopicResourceTypeModel                                MessagingConversationListParamsTopicResourceType = "model"
	MessagingConversationListParamsTopicResourceTypePaymentTerm                          MessagingConversationListParamsTopicResourceType = "payment_term"
	MessagingConversationListParamsTopicResourceTypeShippingTerm                         MessagingConversationListParamsTopicResourceType = "shipping_term"
	MessagingConversationListParamsTopicResourceTypeQuantity                             MessagingConversationListParamsTopicResourceType = "quantity"
	MessagingConversationListParamsTopicResourceTypeAccountGroup                         MessagingConversationListParamsTopicResourceType = "account_group"
	MessagingConversationListParamsTopicResourceTypeSupportRoute                         MessagingConversationListParamsTopicResourceType = "support_route"
	MessagingConversationListParamsTopicResourceTypeSupportAvailability                  MessagingConversationListParamsTopicResourceType = "support_availability"
	MessagingConversationListParamsTopicResourceTypeAccountStatus                        MessagingConversationListParamsTopicResourceType = "account_status"
	MessagingConversationListParamsTopicResourceTypeGeolocation                          MessagingConversationListParamsTopicResourceType = "geolocation"
	MessagingConversationListParamsTopicResourceTypeAccountUser                          MessagingConversationListParamsTopicResourceType = "account_user"
	MessagingConversationListParamsTopicResourceTypeDepartment                           MessagingConversationListParamsTopicResourceType = "department"
	MessagingConversationListParamsTopicResourceTypeAccountIntegration                   MessagingConversationListParamsTopicResourceType = "account_integration"
	MessagingConversationListParamsTopicResourceTypeAccountPrice                         MessagingConversationListParamsTopicResourceType = "account_price"
	MessagingConversationListParamsTopicResourceTypeProductLine                          MessagingConversationListParamsTopicResourceType = "product_line"
	MessagingConversationListParamsTopicResourceTypeItemCategory                         MessagingConversationListParamsTopicResourceType = "item_category"
	MessagingConversationListParamsTopicResourceTypeAttribute                            MessagingConversationListParamsTopicResourceType = "attribute"
	MessagingConversationListParamsTopicResourceTypeRate                                 MessagingConversationListParamsTopicResourceType = "rate"
	MessagingConversationListParamsTopicResourceTypeAccountGroupProductLineAccess        MessagingConversationListParamsTopicResourceType = "account_group_product_line_access"
	MessagingConversationListParamsTopicResourceTypeSalesTarget                          MessagingConversationListParamsTopicResourceType = "sales_target"
	MessagingConversationListParamsTopicResourceTypeAdjustmentType                       MessagingConversationListParamsTopicResourceType = "adjustment_type"
	MessagingConversationListParamsTopicResourceTypeAccountBranding                      MessagingConversationListParamsTopicResourceType = "account_branding"
	MessagingConversationListParamsTopicResourceTypeAccountPortal                        MessagingConversationListParamsTopicResourceType = "account_portal"
	MessagingConversationListParamsTopicResourceTypeAccountLogoURL                       MessagingConversationListParamsTopicResourceType = "account_logo_url"
	MessagingConversationListParamsTopicResourceTypeAccountFaviconURL                    MessagingConversationListParamsTopicResourceType = "account_favicon_url"
	MessagingConversationListParamsTopicResourceTypePublicAccount                        MessagingConversationListParamsTopicResourceType = "public_account"
	MessagingConversationListParamsTopicResourceTypeProperty                             MessagingConversationListParamsTopicResourceType = "property"
	MessagingConversationListParamsTopicResourceTypeCarrier                              MessagingConversationListParamsTopicResourceType = "carrier"
	MessagingConversationListParamsTopicResourceTypeServiceLevel                         MessagingConversationListParamsTopicResourceType = "service_level"
	MessagingConversationListParamsTopicResourceTypeItem                                 MessagingConversationListParamsTopicResourceType = "item"
	MessagingConversationListParamsTopicResourceTypeItemLotDefault                       MessagingConversationListParamsTopicResourceType = "item_lot_default"
	MessagingConversationListParamsTopicResourceTypeItemInventory                        MessagingConversationListParamsTopicResourceType = "item_inventory"
	MessagingConversationListParamsTopicResourceTypeProduct                              MessagingConversationListParamsTopicResourceType = "product"
	MessagingConversationListParamsTopicResourceTypeBatch                                MessagingConversationListParamsTopicResourceType = "batch"
	MessagingConversationListParamsTopicResourceTypeBatchFlowNode                        MessagingConversationListParamsTopicResourceType = "batch_flow_node"
	MessagingConversationListParamsTopicResourceTypeScanningConsumption                  MessagingConversationListParamsTopicResourceType = "scanning_consumption"
	MessagingConversationListParamsTopicResourceTypeOpenBatchSummary                     MessagingConversationListParamsTopicResourceType = "open_batch_summary"
	MessagingConversationListParamsTopicResourceTypeScanningProductionStepInfo           MessagingConversationListParamsTopicResourceType = "scanning_production_step_info"
	MessagingConversationListParamsTopicResourceTypeScanningStation                      MessagingConversationListParamsTopicResourceType = "scanning_station"
	MessagingConversationListParamsTopicResourceTypeProductionStep                       MessagingConversationListParamsTopicResourceType = "production_step"
	MessagingConversationListParamsTopicResourceTypeProductionRun                        MessagingConversationListParamsTopicResourceType = "production_run"
	MessagingConversationListParamsTopicResourceTypeMachine                              MessagingConversationListParamsTopicResourceType = "machine"
	MessagingConversationListParamsTopicResourceTypeMachineStatus                        MessagingConversationListParamsTopicResourceType = "machine_status"
	MessagingConversationListParamsTopicResourceTypeMachineDowntimeEvent                 MessagingConversationListParamsTopicResourceType = "machine_downtime_event"
	MessagingConversationListParamsTopicResourceTypeDemandOverride                       MessagingConversationListParamsTopicResourceType = "demand_override"
	MessagingConversationListParamsTopicResourceTypeDemandOverrideType                   MessagingConversationListParamsTopicResourceType = "demand_override_type"
	MessagingConversationListParamsTopicResourceTypeMachineDowntimeReason                MessagingConversationListParamsTopicResourceType = "machine_downtime_reason"
	MessagingConversationListParamsTopicResourceTypeProductionSchedulePreview            MessagingConversationListParamsTopicResourceType = "production_schedule_preview"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleRegeneratePreview  MessagingConversationListParamsTopicResourceType = "production_schedule_regenerate_preview"
	MessagingConversationListParamsTopicResourceTypeProductionSchedule                   MessagingConversationListParamsTopicResourceType = "production_schedule"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleLine               MessagingConversationListParamsTopicResourceType = "production_schedule_line"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleDeviation          MessagingConversationListParamsTopicResourceType = "production_schedule_deviation"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleDerivedLine        MessagingConversationListParamsTopicResourceType = "production_schedule_derived_line"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleSettings           MessagingConversationListParamsTopicResourceType = "production_schedule_settings"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleResourceSetting    MessagingConversationListParamsTopicResourceType = "production_schedule_resource_setting"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleItemSetting        MessagingConversationListParamsTopicResourceType = "production_schedule_item_setting"
	MessagingConversationListParamsTopicResourceTypeFulfillmentRecommendation            MessagingConversationListParamsTopicResourceType = "fulfillment_recommendation"
	MessagingConversationListParamsTopicResourceTypeAnalyzeDeliveryPerformanceResponse   MessagingConversationListParamsTopicResourceType = "analyze_delivery_performance_response"
	MessagingConversationListParamsTopicResourceTypeDeliveryPerformance                  MessagingConversationListParamsTopicResourceType = "delivery_performance"
	MessagingConversationListParamsTopicResourceTypeDeliveryBacklogBucket                MessagingConversationListParamsTopicResourceType = "delivery_backlog_bucket"
	MessagingConversationListParamsTopicResourceTypeDeliveryLatenessBucket               MessagingConversationListParamsTopicResourceType = "delivery_lateness_bucket"
	MessagingConversationListParamsTopicResourceTypeDeliveryBreakdown                    MessagingConversationListParamsTopicResourceType = "delivery_breakdown"
	MessagingConversationListParamsTopicResourceTypeScheduleOrderCoverage                MessagingConversationListParamsTopicResourceType = "schedule_order_coverage"
	MessagingConversationListParamsTopicResourceTypeScheduleOrderCoverageLine            MessagingConversationListParamsTopicResourceType = "schedule_order_coverage_line"
	MessagingConversationListParamsTopicResourceTypeScheduleDeviationType                MessagingConversationListParamsTopicResourceType = "schedule_deviation_type"
	MessagingConversationListParamsTopicResourceTypeScheduleAtRiskOrder                  MessagingConversationListParamsTopicResourceType = "schedule_at_risk_order"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleFinishedPolicy     MessagingConversationListParamsTopicResourceType = "production_schedule_finished_policy"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleFinishingLine      MessagingConversationListParamsTopicResourceType = "production_schedule_finishing_line"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleWeekRelease        MessagingConversationListParamsTopicResourceType = "production_schedule_week_release"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleWeekReleasePreview MessagingConversationListParamsTopicResourceType = "production_schedule_week_release_preview"
	MessagingConversationListParamsTopicResourceTypeProductionScheduleItemPolicy         MessagingConversationListParamsTopicResourceType = "production_schedule_item_policy"
	MessagingConversationListParamsTopicResourceTypeChildAccount                         MessagingConversationListParamsTopicResourceType = "child_account"
	MessagingConversationListParamsTopicResourceTypeUnitGroup                            MessagingConversationListParamsTopicResourceType = "unit_group"
	MessagingConversationListParamsTopicResourceTypeUnitGroupUnit                        MessagingConversationListParamsTopicResourceType = "unit_group_unit"
	MessagingConversationListParamsTopicResourceTypeConsumption                          MessagingConversationListParamsTopicResourceType = "consumption"
	MessagingConversationListParamsTopicResourceTypeCustomerProductLineAccess            MessagingConversationListParamsTopicResourceType = "customer_product_line_access"
	MessagingConversationListParamsTopicResourceTypeCustomer                             MessagingConversationListParamsTopicResourceType = "customer"
	MessagingConversationListParamsTopicResourceTypeFrequentlyOrderedProduct             MessagingConversationListParamsTopicResourceType = "frequently_ordered_product"
	MessagingConversationListParamsTopicResourceTypePriority                             MessagingConversationListParamsTopicResourceType = "priority"
	MessagingConversationListParamsTopicResourceTypeDelivery                             MessagingConversationListParamsTopicResourceType = "delivery"
	MessagingConversationListParamsTopicResourceTypeDeliveryLine                         MessagingConversationListParamsTopicResourceType = "delivery_line"
	MessagingConversationListParamsTopicResourceTypeSalesOrder                           MessagingConversationListParamsTopicResourceType = "sales_order"
	MessagingConversationListParamsTopicResourceTypeLocation                             MessagingConversationListParamsTopicResourceType = "location"
	MessagingConversationListParamsTopicResourceTypeLocationType                         MessagingConversationListParamsTopicResourceType = "location_type"
	MessagingConversationListParamsTopicResourceTypeLot                                  MessagingConversationListParamsTopicResourceType = "lot"
	MessagingConversationListParamsTopicResourceTypeEmailLog                             MessagingConversationListParamsTopicResourceType = "email_log"
	MessagingConversationListParamsTopicResourceTypeEmailDomain                          MessagingConversationListParamsTopicResourceType = "email_domain"
	MessagingConversationListParamsTopicResourceTypeEmailInbox                           MessagingConversationListParamsTopicResourceType = "email_inbox"
	MessagingConversationListParamsTopicResourceTypePortalDomain                         MessagingConversationListParamsTopicResourceType = "portal_domain"
	MessagingConversationListParamsTopicResourceTypeDNSRecord                            MessagingConversationListParamsTopicResourceType = "dns_record"
	MessagingConversationListParamsTopicResourceTypeInventoryChangeLog                   MessagingConversationListParamsTopicResourceType = "inventory_change_log"
	MessagingConversationListParamsTopicResourceTypeInvoice                              MessagingConversationListParamsTopicResourceType = "invoice"
	MessagingConversationListParamsTopicResourceTypeInvoiceSummary                       MessagingConversationListParamsTopicResourceType = "invoice_summary"
	MessagingConversationListParamsTopicResourceTypeInvoiceLine                          MessagingConversationListParamsTopicResourceType = "invoice_line"
	MessagingConversationListParamsTopicResourceTypeInvoiceAllocation                    MessagingConversationListParamsTopicResourceType = "invoice_allocation"
	MessagingConversationListParamsTopicResourceTypeInvoiceForPayment                    MessagingConversationListParamsTopicResourceType = "invoice_for_payment"
	MessagingConversationListParamsTopicResourceTypeShipment                             MessagingConversationListParamsTopicResourceType = "shipment"
	MessagingConversationListParamsTopicResourceTypeShipmentSummary                      MessagingConversationListParamsTopicResourceType = "shipment_summary"
	MessagingConversationListParamsTopicResourceTypeShipmentLine                         MessagingConversationListParamsTopicResourceType = "shipment_line"
	MessagingConversationListParamsTopicResourceTypeShippingCase                         MessagingConversationListParamsTopicResourceType = "shipping_case"
	MessagingConversationListParamsTopicResourceTypeShippingCaseLabelURL                 MessagingConversationListParamsTopicResourceType = "shipping_case_label_url"
	MessagingConversationListParamsTopicResourceTypeSettlement                           MessagingConversationListParamsTopicResourceType = "settlement"
	MessagingConversationListParamsTopicResourceTypeSettlementSummary                    MessagingConversationListParamsTopicResourceType = "settlement_summary"
	MessagingConversationListParamsTopicResourceTypeRolePermission                       MessagingConversationListParamsTopicResourceType = "role_permission"
	MessagingConversationListParamsTopicResourceTypeRegistrationFlow                     MessagingConversationListParamsTopicResourceType = "registration_flow"
	MessagingConversationListParamsTopicResourceTypeRegistrationFlowOption               MessagingConversationListParamsTopicResourceType = "registration_flow_option"
	MessagingConversationListParamsTopicResourceTypeTransaction                          MessagingConversationListParamsTopicResourceType = "transaction"
	MessagingConversationListParamsTopicResourceTypeTransactionSummary                   MessagingConversationListParamsTopicResourceType = "transaction_summary"
	MessagingConversationListParamsTopicResourceTypeTransactionMethod                    MessagingConversationListParamsTopicResourceType = "transaction_method"
	MessagingConversationListParamsTopicResourceTypeTransactionType                      MessagingConversationListParamsTopicResourceType = "transaction_type"
	MessagingConversationListParamsTopicResourceTypeTransactionAllocation                MessagingConversationListParamsTopicResourceType = "transaction_allocation"
	MessagingConversationListParamsTopicResourceTypeUsageItem                            MessagingConversationListParamsTopicResourceType = "usage_item"
	MessagingConversationListParamsTopicResourceTypeAccountUsageResponse                 MessagingConversationListParamsTopicResourceType = "account_usage_response"
	MessagingConversationListParamsTopicResourceTypeSubscriptionInfo                     MessagingConversationListParamsTopicResourceType = "subscription_info"
	MessagingConversationListParamsTopicResourceTypeBillingPortalSessionResponse         MessagingConversationListParamsTopicResourceType = "billing_portal_session_response"
	MessagingConversationListParamsTopicResourceTypeSwitchPlanResponse                   MessagingConversationListParamsTopicResourceType = "switch_plan_response"
	MessagingConversationListParamsTopicResourceTypeEnsureBillingCustomerResponse        MessagingConversationListParamsTopicResourceType = "ensure_billing_customer_response"
	MessagingConversationListParamsTopicResourceTypeSpendingCapResponse                  MessagingConversationListParamsTopicResourceType = "spending_cap_response"
	MessagingConversationListParamsTopicResourceTypeAgentSpendInfo                       MessagingConversationListParamsTopicResourceType = "agent_spend_info"
	MessagingConversationListParamsTopicResourceTypeWebhookResponse                      MessagingConversationListParamsTopicResourceType = "webhook_response"
	MessagingConversationListParamsTopicResourceTypeAddressSuggestion                    MessagingConversationListParamsTopicResourceType = "address_suggestion"
	MessagingConversationListParamsTopicResourceTypeAddressComponents                    MessagingConversationListParamsTopicResourceType = "address_components"
	MessagingConversationListParamsTopicResourceTypeAddressDetailsResult                 MessagingConversationListParamsTopicResourceType = "address_details_result"
	MessagingConversationListParamsTopicResourceTypeValidatedAddress                     MessagingConversationListParamsTopicResourceType = "validated_address"
	MessagingConversationListParamsTopicResourceTypePlanLimit                            MessagingConversationListParamsTopicResourceType = "plan_limit"
	MessagingConversationListParamsTopicResourceTypePlanChangeProration                  MessagingConversationListParamsTopicResourceType = "plan_change_proration"
	MessagingConversationListParamsTopicResourceTypePlanChangeLineItem                   MessagingConversationListParamsTopicResourceType = "plan_change_line_item"
	MessagingConversationListParamsTopicResourceTypeSetupBillingResponse                 MessagingConversationListParamsTopicResourceType = "setup_billing_response"
	MessagingConversationListParamsTopicResourceTypeConfirmPaymentResponse               MessagingConversationListParamsTopicResourceType = "confirm_payment_response"
	MessagingConversationListParamsTopicResourceTypeOAuthResponse                        MessagingConversationListParamsTopicResourceType = "oauth_response"
	MessagingConversationListParamsTopicResourceTypeOAuthStatusResponse                  MessagingConversationListParamsTopicResourceType = "oauth_status_response"
	MessagingConversationListParamsTopicResourceTypeStripePublishableKey                 MessagingConversationListParamsTopicResourceType = "stripe_publishable_key"
	MessagingConversationListParamsTopicResourceTypeStripeStatus                         MessagingConversationListParamsTopicResourceType = "stripe_status"
	MessagingConversationListParamsTopicResourceTypeHealthcheck                          MessagingConversationListParamsTopicResourceType = "healthcheck"
	MessagingConversationListParamsTopicResourceTypeAgentDefinitionConfig                MessagingConversationListParamsTopicResourceType = "agent_definition_config"
	MessagingConversationListParamsTopicResourceTypeTriggerConfig                        MessagingConversationListParamsTopicResourceType = "trigger_config"
	MessagingConversationListParamsTopicResourceTypeCustomerContactInfo                  MessagingConversationListParamsTopicResourceType = "customer_contact_info"
	MessagingConversationListParamsTopicResourceTypeCustomerFreightPreferences           MessagingConversationListParamsTopicResourceType = "customer_freight_preferences"
	MessagingConversationListParamsTopicResourceTypeCustomerDefaults                     MessagingConversationListParamsTopicResourceType = "customer_defaults"
	MessagingConversationListParamsTopicResourceTypeCustomerLeadTime                     MessagingConversationListParamsTopicResourceType = "customer_lead_time"
	MessagingConversationListParamsTopicResourceTypeCustomerNotificationPreferences      MessagingConversationListParamsTopicResourceType = "customer_notification_preferences"
	MessagingConversationListParamsTopicResourceTypeOrderNotificationRecipient           MessagingConversationListParamsTopicResourceType = "order_notification_recipient"
	MessagingConversationListParamsTopicResourceTypeOrderDiscount                        MessagingConversationListParamsTopicResourceType = "order_discount"
	MessagingConversationListParamsTopicResourceTypeSalesOrderLine                       MessagingConversationListParamsTopicResourceType = "sales_order_line"
	MessagingConversationListParamsTopicResourceTypeSalesOrderType                       MessagingConversationListParamsTopicResourceType = "sales_order_type"
	MessagingConversationListParamsTopicResourceTypeSalesOrderStatus                     MessagingConversationListParamsTopicResourceType = "sales_order_status"
	MessagingConversationListParamsTopicResourceTypeMaterial                             MessagingConversationListParamsTopicResourceType = "material"
	MessagingConversationListParamsTopicResourceTypeSupplierMaterial                     MessagingConversationListParamsTopicResourceType = "supplier_material"
	MessagingConversationListParamsTopicResourceTypePart                                 MessagingConversationListParamsTopicResourceType = "part"
	MessagingConversationListParamsTopicResourceTypePermissionGroup                      MessagingConversationListParamsTopicResourceType = "permission_group"
	MessagingConversationListParamsTopicResourceTypePermission                           MessagingConversationListParamsTopicResourceType = "permission"
	MessagingConversationListParamsTopicResourceTypePick                                 MessagingConversationListParamsTopicResourceType = "pick"
	MessagingConversationListParamsTopicResourceTypePickLine                             MessagingConversationListParamsTopicResourceType = "pick_line"
	MessagingConversationListParamsTopicResourceTypeProductType                          MessagingConversationListParamsTopicResourceType = "product_type"
	MessagingConversationListParamsTopicResourceTypeProduction                           MessagingConversationListParamsTopicResourceType = "production"
	MessagingConversationListParamsTopicResourceTypeProductionFlow                       MessagingConversationListParamsTopicResourceType = "production_flow"
	MessagingConversationListParamsTopicResourceTypeMap                                  MessagingConversationListParamsTopicResourceType = "map"
	MessagingConversationListParamsTopicResourceTypePurchaseOrder                        MessagingConversationListParamsTopicResourceType = "purchase_order"
	MessagingConversationListParamsTopicResourceTypePurchaseOrderLine                    MessagingConversationListParamsTopicResourceType = "purchase_order_line"
	MessagingConversationListParamsTopicResourceTypeSupplier                             MessagingConversationListParamsTopicResourceType = "supplier"
	MessagingConversationListParamsTopicResourceTypeSupplierSummary                      MessagingConversationListParamsTopicResourceType = "supplier_summary"
	MessagingConversationListParamsTopicResourceTypeReceivableEntry                      MessagingConversationListParamsTopicResourceType = "receivable_entry"
	MessagingConversationListParamsTopicResourceTypeReceivingOrder                       MessagingConversationListParamsTopicResourceType = "receiving_order"
	MessagingConversationListParamsTopicResourceTypeReceivingOrderLine                   MessagingConversationListParamsTopicResourceType = "receiving_order_line"
	MessagingConversationListParamsTopicResourceTypeEmailContact                         MessagingConversationListParamsTopicResourceType = "email_contact"
	MessagingConversationListParamsTopicResourceTypeAllocationEntry                      MessagingConversationListParamsTopicResourceType = "allocation_entry"
	MessagingConversationListParamsTopicResourceTypeOpenCreditEntry                      MessagingConversationListParamsTopicResourceType = "open_credit_entry"
	MessagingConversationListParamsTopicResourceTypeVolumeDiscount                       MessagingConversationListParamsTopicResourceType = "volume_discount"
	MessagingConversationListParamsTopicResourceTypeVolumeDiscountTier                   MessagingConversationListParamsTopicResourceType = "volume_discount_tier"
	MessagingConversationListParamsTopicResourceTypeAnalyzeDeliveriesResponse            MessagingConversationListParamsTopicResourceType = "analyze_deliveries_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeManufacturingResponse         MessagingConversationListParamsTopicResourceType = "analyze_manufacturing_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeManufacturingBatchResponse    MessagingConversationListParamsTopicResourceType = "analyze_manufacturing_batch_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeQuarterlyOrdersResponse       MessagingConversationListParamsTopicResourceType = "analyze_quarterly_orders_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeNewCustomersResponse          MessagingConversationListParamsTopicResourceType = "analyze_new_customers_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeDemandForecastResponse        MessagingConversationListParamsTopicResourceType = "analyze_demand_forecast_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeOeeResponse                   MessagingConversationListParamsTopicResourceType = "analyze_oee_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeOeeTrendResponse              MessagingConversationListParamsTopicResourceType = "analyze_oee_trend_response"
	MessagingConversationListParamsTopicResourceTypeAnalyzeScheduleAttainmentResponse    MessagingConversationListParamsTopicResourceType = "analyze_schedule_attainment_response"
	MessagingConversationListParamsTopicResourceTypeCatalogProductLine                   MessagingConversationListParamsTopicResourceType = "catalog_product_line"
	MessagingConversationListParamsTopicResourceTypeCatalogCategory                      MessagingConversationListParamsTopicResourceType = "catalog_category"
	MessagingConversationListParamsTopicResourceTypeCatalogProduct                       MessagingConversationListParamsTopicResourceType = "catalog_product"
	MessagingConversationListParamsTopicResourceTypeCatalogProperty                      MessagingConversationListParamsTopicResourceType = "catalog_property"
	MessagingConversationListParamsTopicResourceTypeCatalogAttribute                     MessagingConversationListParamsTopicResourceType = "catalog_attribute"
	MessagingConversationListParamsTopicResourceTypeDcLocation                           MessagingConversationListParamsTopicResourceType = "dc_location"
	MessagingConversationListParamsTopicResourceTypeEdiRun                               MessagingConversationListParamsTopicResourceType = "edi_run"
	MessagingConversationListParamsTopicResourceTypeInventoryItem                        MessagingConversationListParamsTopicResourceType = "inventory_item"
	MessagingConversationListParamsTopicResourceTypeAnalyzeWeeksOfSalesResponse          MessagingConversationListParamsTopicResourceType = "analyze_weeks_of_sales_response"
	MessagingConversationListParamsTopicResourceTypeBulkReconcileItemsResponse           MessagingConversationListParamsTopicResourceType = "bulk_reconcile_items_response"
	MessagingConversationListParamsTopicResourceTypeSysProperty                          MessagingConversationListParamsTopicResourceType = "sys_property"
	MessagingConversationListParamsTopicResourceTypeSysPropertyType                      MessagingConversationListParamsTopicResourceType = "sys_property_type"
	MessagingConversationListParamsTopicResourceTypeSysPropertyValue                     MessagingConversationListParamsTopicResourceType = "sys_property_value"
	MessagingConversationListParamsTopicResourceTypeTerritory                            MessagingConversationListParamsTopicResourceType = "territory"
	MessagingConversationListParamsTopicResourceTypeTenancy                              MessagingConversationListParamsTopicResourceType = "tenancy"
	MessagingConversationListParamsTopicResourceTypeCheckoutSession                      MessagingConversationListParamsTopicResourceType = "checkout_session"
	MessagingConversationListParamsTopicResourceTypeEstimateRateResult                   MessagingConversationListParamsTopicResourceType = "estimate_rate_result"
	MessagingConversationListParamsTopicResourceTypeRateShopOption                       MessagingConversationListParamsTopicResourceType = "rate_shop_option"
	MessagingConversationListParamsTopicResourceTypeRateShopResult                       MessagingConversationListParamsTopicResourceType = "rate_shop_result"
	MessagingConversationListParamsTopicResourceTypeOwner                                MessagingConversationListParamsTopicResourceType = "owner"
	MessagingConversationListParamsTopicResourceTypeCreatedBy                            MessagingConversationListParamsTopicResourceType = "created_by"
	MessagingConversationListParamsTopicResourceTypeMessage                              MessagingConversationListParamsTopicResourceType = "message"
	MessagingConversationListParamsTopicResourceTypeAccountPhotoUploadResult             MessagingConversationListParamsTopicResourceType = "account_photo_upload_result"
	MessagingConversationListParamsTopicResourceTypeUserPhotoUploadResult                MessagingConversationListParamsTopicResourceType = "user_photo_upload_result"
	MessagingConversationListParamsTopicResourceTypeUserPhotoURL                         MessagingConversationListParamsTopicResourceType = "user_photo_url"
	MessagingConversationListParamsTopicResourceTypeBatchLot                             MessagingConversationListParamsTopicResourceType = "batch_lot"
	MessagingConversationListParamsTopicResourceTypeCheckDuplicateResult                 MessagingConversationListParamsTopicResourceType = "check_duplicate_result"
	MessagingConversationListParamsTopicResourceTypeItemTrendPoint                       MessagingConversationListParamsTopicResourceType = "item_trend_point"
	MessagingConversationListParamsTopicResourceTypeTenancyPendingRegistration           MessagingConversationListParamsTopicResourceType = "tenancy_pending_registration"
	MessagingConversationListParamsTopicResourceTypeInvoiceAllocationEntry               MessagingConversationListParamsTopicResourceType = "invoice_allocation_entry"
	MessagingConversationListParamsTopicResourceTypeAllocationCustomer                   MessagingConversationListParamsTopicResourceType = "allocation_customer"
	MessagingConversationListParamsTopicResourceTypeCheckoutSalesOrder                   MessagingConversationListParamsTopicResourceType = "checkout_sales_order"
	MessagingConversationListParamsTopicResourceTypeSalesOrderPriceQuote                 MessagingConversationListParamsTopicResourceType = "sales_order_price_quote"
	MessagingConversationListParamsTopicResourceTypeSalesOrderFreightQuote               MessagingConversationListParamsTopicResourceType = "sales_order_freight_quote"
	MessagingConversationListParamsTopicResourceTypeSalesOrderCommitmentQuote            MessagingConversationListParamsTopicResourceType = "sales_order_commitment_quote"
	MessagingConversationListParamsTopicResourceTypeOperatingCalendar                    MessagingConversationListParamsTopicResourceType = "operating_calendar"
	MessagingConversationListParamsTopicResourceTypeOperatingCalendarClosure             MessagingConversationListParamsTopicResourceType = "operating_calendar_closure"
	MessagingConversationListParamsTopicResourceTypeSalesOrderPriceQuoteLine             MessagingConversationListParamsTopicResourceType = "sales_order_price_quote_line"
	MessagingConversationListParamsTopicResourceTypeHubspotSyncJob                       MessagingConversationListParamsTopicResourceType = "hubspot_sync_job"
	MessagingConversationListParamsTopicResourceTypeHubspotSyncReport                    MessagingConversationListParamsTopicResourceType = "hubspot_sync_report"
	MessagingConversationListParamsTopicResourceTypeHubspotCompanyReview                 MessagingConversationListParamsTopicResourceType = "hubspot_company_review"
	MessagingConversationListParamsTopicResourceTypeHubspotCompanyCandidate              MessagingConversationListParamsTopicResourceType = "hubspot_company_candidate"
	MessagingConversationListParamsTopicResourceTypeHubspotSyncRecord                    MessagingConversationListParamsTopicResourceType = "hubspot_sync_record"
	MessagingConversationListParamsTopicResourceTypeContactMatch                         MessagingConversationListParamsTopicResourceType = "contact_match"
	MessagingConversationListParamsTopicResourceTypeReplyDraft                           MessagingConversationListParamsTopicResourceType = "reply_draft"
	MessagingConversationListParamsTopicResourceTypeConversationLink                     MessagingConversationListParamsTopicResourceType = "conversation_link"
	MessagingConversationListParamsTopicResourceTypeMessagingGroup                       MessagingConversationListParamsTopicResourceType = "messaging_group"
	MessagingConversationListParamsTopicResourceTypeMessagingGroupMember                 MessagingConversationListParamsTopicResourceType = "messaging_group_member"
	MessagingConversationListParamsTopicResourceTypePortalProfile                        MessagingConversationListParamsTopicResourceType = "portal_profile"
	MessagingConversationListParamsTopicResourceTypePortalRegistrationSession            MessagingConversationListParamsTopicResourceType = "portal_registration_session"
	MessagingConversationListParamsTopicResourceTypePortalRegistrationSessionData        MessagingConversationListParamsTopicResourceType = "portal_registration_session_data"
	MessagingConversationListParamsTopicResourceTypePackList                             MessagingConversationListParamsTopicResourceType = "pack_list"
	MessagingConversationListParamsTopicResourceTypePackListParty                        MessagingConversationListParamsTopicResourceType = "pack_list_party"
	MessagingConversationListParamsTopicResourceTypePackListLineItem                     MessagingConversationListParamsTopicResourceType = "pack_list_line_item"
	MessagingConversationListParamsTopicResourceTypePackListBackOrder                    MessagingConversationListParamsTopicResourceType = "pack_list_back_order"
	MessagingConversationListParamsTopicResourceTypePackListCase                         MessagingConversationListParamsTopicResourceType = "pack_list_case"
	MessagingConversationListParamsTopicResourceTypeJob                                  MessagingConversationListParamsTopicResourceType = "job"
	MessagingConversationListParamsTopicResourceTypeJobResult                            MessagingConversationListParamsTopicResourceType = "job_result"
	MessagingConversationListParamsTopicResourceTypeJobExport                            MessagingConversationListParamsTopicResourceType = "job_export"
	MessagingConversationListParamsTopicResourceTypeAnalyzeCustomerPricingResponse       MessagingConversationListParamsTopicResourceType = "analyze_customer_pricing_response"
	MessagingConversationListParamsTopicResourceTypeCustomerPricingFinding               MessagingConversationListParamsTopicResourceType = "customer_pricing_finding"
	MessagingConversationListParamsTopicResourceTypeCustomerPricingSummary               MessagingConversationListParamsTopicResourceType = "customer_pricing_summary"
	MessagingConversationListParamsTopicResourceTypeComputedRate                         MessagingConversationListParamsTopicResourceType = "computed_rate"
	MessagingConversationListParamsTopicResourceTypeComputedQuantity                     MessagingConversationListParamsTopicResourceType = "computed_quantity"
	MessagingConversationListParamsTopicResourceTypeAnalyzeRealizedMarginsResponse       MessagingConversationListParamsTopicResourceType = "analyze_realized_margins_response"
	MessagingConversationListParamsTopicResourceTypeRealizedMarginFinding                MessagingConversationListParamsTopicResourceType = "realized_margin_finding"
	MessagingConversationListParamsTopicResourceTypeRealizedMarginSummary                MessagingConversationListParamsTopicResourceType = "realized_margin_summary"
	MessagingConversationListParamsTopicResourceTypeShipmentRelated                      MessagingConversationListParamsTopicResourceType = "shipment_related"
	MessagingConversationListParamsTopicResourceTypeInvoiceRelated                       MessagingConversationListParamsTopicResourceType = "invoice_related"
	MessagingConversationListParamsTopicResourceTypePickRelated                          MessagingConversationListParamsTopicResourceType = "pick_related"
	MessagingConversationListParamsTopicResourceTypePickShipmentsResponse                MessagingConversationListParamsTopicResourceType = "pick_shipments_response"
	MessagingConversationListParamsTopicResourceTypePickTotals                           MessagingConversationListParamsTopicResourceType = "pick_totals"
	MessagingConversationListParamsTopicResourceTypePickStageTotal                       MessagingConversationListParamsTopicResourceType = "pick_stage_total"
)

// Filter by conversation type.
type MessagingConversationListParamsType string

const (
	MessagingConversationListParamsTypeDirectMessage MessagingConversationListParamsType = "direct_message"
	MessagingConversationListParamsTypeGroup         MessagingConversationListParamsType = "group"
	MessagingConversationListParamsTypeSystem        MessagingConversationListParamsType = "system"
)

// Filter the support inbox to a single triage lane.
//
// - `new`: opened but nobody has triaged it yet.
// - `open`: actively being worked.
// - `waiting_internal`: blocked on the internal team.
// - `waiting_external`: blocked on a reply from the customer.
// - `needs_approval`: a drafted reply is waiting for a human to approve it.
// - `resolved`: closed out.
//
// The working inbox hides resolved cases unless you ask for this lane explicitly.
type MessagingConversationListParamsWorkflowStatus string

const (
	MessagingConversationListParamsWorkflowStatusNew             MessagingConversationListParamsWorkflowStatus = "new"
	MessagingConversationListParamsWorkflowStatusOpen            MessagingConversationListParamsWorkflowStatus = "open"
	MessagingConversationListParamsWorkflowStatusWaitingInternal MessagingConversationListParamsWorkflowStatus = "waiting_internal"
	MessagingConversationListParamsWorkflowStatusWaitingExternal MessagingConversationListParamsWorkflowStatus = "waiting_external"
	MessagingConversationListParamsWorkflowStatusNeedsApproval   MessagingConversationListParamsWorkflowStatus = "needs_approval"
	MessagingConversationListParamsWorkflowStatusResolved        MessagingConversationListParamsWorkflowStatus = "resolved"
)
