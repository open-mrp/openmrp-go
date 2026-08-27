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

// Send, list, edit, and delete chat messages.
//
// MessagingConversationMessageService contains methods and other services that
// help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationMessageService] method instead.
type MessagingConversationMessageService struct {
	options []option.RequestOption
}

// NewMessagingConversationMessageService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingConversationMessageService(opts ...option.RequestOption) (r MessagingConversationMessageService) {
	r = MessagingConversationMessageService{}
	r.options = opts
	return
}

// Posts a message to a conversation.
//
// With `mode` = `send` the message is delivered — immediately, or queued when
// `scheduled_at` is set — and a retry of an immediate send with the same
// `client_message_id` returns the original message rather than posting it twice.
// With `mode` = `draft` the message is proposed as a reply to the customer and
// held for a teammate to approve instead of being sent, and `channel` is required.
//
// Sending requires you to be an active participant allowed to post: view-only
// participants cannot post, and in a direct message neither side of a block can.
// On a customer-facing case, replying to the customer moves the case to waiting on
// the customer, and proposing a draft moves it to awaiting approval.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingConversationMessageService) New(ctx context.Context, id string, params MessagingConversationMessageNewParams, opts ...option.RequestOption) (res *Message, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/messages", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the messages in a conversation, newest first.
//
// You must be an active participant. A customer reading their own case receives
// only the messages meant for them — internal team notes are never included.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingConversationMessageService) List(ctx context.Context, id string, query MessagingConversationMessageListParams, opts ...option.RequestOption) (res *ListMessage, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/messages", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMessage struct {
	// Resources in this page.
	Data []Message `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessageObject `json:"object" api:"required"`
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
func (r ListMessage) RawJSON() string { return r.JSON.raw }
func (r *ListMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessageObject string

const (
	ListMessageObjectList ListMessageObject = "list"
)

// A single attachment supplied when sending a message.
//
// For an uploaded file or image, supply the `s3_key` you uploaded to; for a link,
// supply `url`; for a resource reference, supply `resource_type` and
// `resource_id`.
//
// The property Kind is required.
type MessageAttachmentInputParam struct {
	// What is being attached.
	//
	// - `file`: a document you uploaded to object storage first.
	// - `image`: an uploaded image, rendered inline in the conversation.
	// - `link`: an external web address, with nothing stored on our side.
	// - `resource`: a reference to an in-app record, such as an order.
	//
	// Any of "file", "image", "link", "resource".
	Kind MessageAttachmentInputKind `json:"kind,omitzero" api:"required"`
	// The MIME content type of the uploaded file (file and image).
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// The filename to display for the attachment (file and image).
	Filename param.Opt[string] `json:"filename,omitzero"`
	// The id of the record being referenced, paired with `resource_type` (resource).
	ResourceID param.Opt[string] `json:"resource_id,omitzero"`
	// The type of the record being referenced, paired with `resource_id` (resource).
	ResourceType param.Opt[string] `json:"resource_type,omitzero"`
	// The key you uploaded the file to, taken from the upload-url response (file and
	// image).
	//
	// The key must be one minted for this conversation and the file must already be
	// uploaded, otherwise the send is rejected.
	S3Key param.Opt[string] `json:"s3_key,omitzero"`
	// The size of the uploaded file in bytes (file and image).
	SizeBytes param.Opt[int64] `json:"size_bytes,omitzero"`
	// The web address being shared (link).
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r MessageAttachmentInputParam) MarshalJSON() (data []byte, err error) {
	type shadow MessageAttachmentInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessageAttachmentInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What is being attached.
//
// - `file`: a document you uploaded to object storage first.
// - `image`: an uploaded image, rendered inline in the conversation.
// - `link`: an external web address, with nothing stored on our side.
// - `resource`: a reference to an in-app record, such as an order.
type MessageAttachmentInputKind string

const (
	MessageAttachmentInputKindFile     MessageAttachmentInputKind = "file"
	MessageAttachmentInputKindImage    MessageAttachmentInputKind = "image"
	MessageAttachmentInputKindLink     MessageAttachmentInputKind = "link"
	MessageAttachmentInputKindResource MessageAttachmentInputKind = "resource"
)

// Request to post a message to a conversation.
//
// The properties Body, ClientMessageID are required.
type SendMessageRequestParam struct {
	// Message body.
	//
	// Required unless the message carries at least one attachment or a resource link.
	Body string `json:"body" api:"required"`
	// Client-supplied dedupe key.
	//
	// Repeating an immediate send with the same value returns the message created by
	// the first request instead of posting a second one, so a retry after a network
	// failure is safe. Required when sending (`mode` = `send`); ignored for drafts.
	ClientMessageID string `json:"client_message_id" api:"required"`
	// ID of a resource to link in the message, paired with `link_resource_type`.
	LinkResourceID param.Opt[string] `json:"link_resource_id,omitzero"`
	// The message this one is a reply to.
	ReplyToMessageID param.Opt[string] `json:"reply_to_message_id,omitzero"`
	// When set, hold the message and deliver it at this future time instead of sending
	// it now.
	//
	// Only the body is carried into a scheduled send — attachments, mentions, copied
	// recipients, resource links, replies, and audience are dropped, and it is
	// delivered as an ordinary team-visible message. If you are no longer an active
	// participant when it comes due, it is canceled instead of sent.
	ScheduledAt param.Opt[time.Time] `json:"scheduled_at,omitzero" format:"date-time"`
	// The internal thread message a draft is composed from, when drafting from a
	// thread (`mode` = `draft`).
	SourceThreadMessageID param.Opt[string] `json:"source_thread_message_id,omitzero"`
	// The subject line for a customer reply sent by email.
	//
	// When omitted, the reply goes out as "Re:" the case title.
	Subject param.Opt[string] `json:"subject,omitzero"`
	// Attachments to include with the message.
	Attachments []MessageAttachmentInputParam `json:"attachments,omitzero"`
	// Who the message is addressed to on a customer-facing case.
	//
	//   - `customer`: a reply the customer sees, shown to them as coming from "Customer
	//     Service" and delivered as email when the case is bridged to an inbox.
	//   - `internal`: a team-only note the customer never sees.
	//
	// Messages are team-only unless you ask for `customer`, so an internal note can
	// never leak by omission. Asking for `customer` on a conversation that has no
	// customer is rejected.
	//
	// On a case bridged to an email inbox, a customer reply goes out as mail carrying
	// only the body, subject, and copied recipients — attachments, mentions, resource
	// links, and replies are dropped.
	//
	// Any of "internal", "customer".
	Audience SendMessageRequestAudience `json:"audience,omitzero"`
	// Additional email addresses to copy on a customer reply sent by email.
	Cc []string `json:"cc,omitzero"`
	// The channel a draft will be sent over once it is approved (`mode` = `draft`).
	//
	//   - `message`: appears in the customer's conversation timeline.
	//   - `email`: goes out as an email from the inbox the case is bridged to. Falls
	//     back to the conversation timeline if the case has no bridged inbox.
	//
	// Any of "message", "email".
	Channel SendMessageRequestChannel `json:"channel,omitzero"`
	// Type of a resource to link in the message, paired with `link_resource_id`.
	//
	// Linking a record lets clients render the message as a reference to it. A link
	// counts in place of text, so a message may consist of nothing but the link.
	//
	// Any of "account", "actor", "entity", "record", "freight", "commitment",
	// "sales_order_totals", "sales_order_stage_total", "sales_order_related",
	// "order_contact", "user", "address", "api_key", "created_api_key",
	// "refresh_token", "list", "sandbox", "registration_session", "pricing_plan",
	// "account_plan", "plan_change", "enterprise_inquiry", "request_log",
	// "audit_event", "audit_field_change", "role", "unit", "account_affiliation",
	// "agent_definition", "available_tool", "agent_definition_tool",
	// "agent_account_status", "agent_run", "agent_action", "agent_run_step",
	// "agent_token_usage", "agent_memory", "notification",
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
	// "analyze_sales_breakdown_response", "sales_totals", "sales_breakdown",
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
	// "invoice_related", "pick_related", "pick_totals", "pick_stage_total".
	LinkResourceType SendMessageRequestLinkResourceType `json:"link_resource_type,omitzero"`
	// Account user ids explicitly @mentioned in the message.
	//
	// A mention notifies the person even when they have muted the conversation.
	Mentions []string `json:"mentions,omitzero"`
	// Whether to deliver the message now or hold it as a customer-reply draft.
	//
	//   - `send`: delivers the message, immediately or at `scheduled_at`.
	//   - `draft`: proposes a reply to the customer on a customer-facing case and holds
	//     it for a teammate to approve before it goes out. Requires `channel`.
	//
	// A draft is built from `body`, `subject`, `channel`, and
	// `source_thread_message_id` only — attachments, mentions, copied recipients,
	// resource links, replies, and scheduling are not carried onto it.
	//
	// Any of "send", "draft".
	Mode SendMessageRequestMode `json:"mode,omitzero"`
	paramObj
}

func (r SendMessageRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SendMessageRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendMessageRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who the message is addressed to on a customer-facing case.
//
//   - `customer`: a reply the customer sees, shown to them as coming from "Customer
//     Service" and delivered as email when the case is bridged to an inbox.
//   - `internal`: a team-only note the customer never sees.
//
// Messages are team-only unless you ask for `customer`, so an internal note can
// never leak by omission. Asking for `customer` on a conversation that has no
// customer is rejected.
//
// On a case bridged to an email inbox, a customer reply goes out as mail carrying
// only the body, subject, and copied recipients — attachments, mentions, resource
// links, and replies are dropped.
type SendMessageRequestAudience string

const (
	SendMessageRequestAudienceInternal SendMessageRequestAudience = "internal"
	SendMessageRequestAudienceCustomer SendMessageRequestAudience = "customer"
)

// The channel a draft will be sent over once it is approved (`mode` = `draft`).
//
//   - `message`: appears in the customer's conversation timeline.
//   - `email`: goes out as an email from the inbox the case is bridged to. Falls
//     back to the conversation timeline if the case has no bridged inbox.
type SendMessageRequestChannel string

const (
	SendMessageRequestChannelMessage SendMessageRequestChannel = "message"
	SendMessageRequestChannelEmail   SendMessageRequestChannel = "email"
)

// Type of a resource to link in the message, paired with `link_resource_id`.
//
// Linking a record lets clients render the message as a reference to it. A link
// counts in place of text, so a message may consist of nothing but the link.
type SendMessageRequestLinkResourceType string

const (
	SendMessageRequestLinkResourceTypeAccount                              SendMessageRequestLinkResourceType = "account"
	SendMessageRequestLinkResourceTypeActor                                SendMessageRequestLinkResourceType = "actor"
	SendMessageRequestLinkResourceTypeEntity                               SendMessageRequestLinkResourceType = "entity"
	SendMessageRequestLinkResourceTypeRecord                               SendMessageRequestLinkResourceType = "record"
	SendMessageRequestLinkResourceTypeFreight                              SendMessageRequestLinkResourceType = "freight"
	SendMessageRequestLinkResourceTypeCommitment                           SendMessageRequestLinkResourceType = "commitment"
	SendMessageRequestLinkResourceTypeSalesOrderTotals                     SendMessageRequestLinkResourceType = "sales_order_totals"
	SendMessageRequestLinkResourceTypeSalesOrderStageTotal                 SendMessageRequestLinkResourceType = "sales_order_stage_total"
	SendMessageRequestLinkResourceTypeSalesOrderRelated                    SendMessageRequestLinkResourceType = "sales_order_related"
	SendMessageRequestLinkResourceTypeOrderContact                         SendMessageRequestLinkResourceType = "order_contact"
	SendMessageRequestLinkResourceTypeUser                                 SendMessageRequestLinkResourceType = "user"
	SendMessageRequestLinkResourceTypeAddress                              SendMessageRequestLinkResourceType = "address"
	SendMessageRequestLinkResourceTypeAPIKey                               SendMessageRequestLinkResourceType = "api_key"
	SendMessageRequestLinkResourceTypeCreatedAPIKey                        SendMessageRequestLinkResourceType = "created_api_key"
	SendMessageRequestLinkResourceTypeRefreshToken                         SendMessageRequestLinkResourceType = "refresh_token"
	SendMessageRequestLinkResourceTypeList                                 SendMessageRequestLinkResourceType = "list"
	SendMessageRequestLinkResourceTypeSandbox                              SendMessageRequestLinkResourceType = "sandbox"
	SendMessageRequestLinkResourceTypeRegistrationSession                  SendMessageRequestLinkResourceType = "registration_session"
	SendMessageRequestLinkResourceTypePricingPlan                          SendMessageRequestLinkResourceType = "pricing_plan"
	SendMessageRequestLinkResourceTypeAccountPlan                          SendMessageRequestLinkResourceType = "account_plan"
	SendMessageRequestLinkResourceTypePlanChange                           SendMessageRequestLinkResourceType = "plan_change"
	SendMessageRequestLinkResourceTypeEnterpriseInquiry                    SendMessageRequestLinkResourceType = "enterprise_inquiry"
	SendMessageRequestLinkResourceTypeRequestLog                           SendMessageRequestLinkResourceType = "request_log"
	SendMessageRequestLinkResourceTypeAuditEvent                           SendMessageRequestLinkResourceType = "audit_event"
	SendMessageRequestLinkResourceTypeAuditFieldChange                     SendMessageRequestLinkResourceType = "audit_field_change"
	SendMessageRequestLinkResourceTypeRole                                 SendMessageRequestLinkResourceType = "role"
	SendMessageRequestLinkResourceTypeUnit                                 SendMessageRequestLinkResourceType = "unit"
	SendMessageRequestLinkResourceTypeAccountAffiliation                   SendMessageRequestLinkResourceType = "account_affiliation"
	SendMessageRequestLinkResourceTypeAgentDefinition                      SendMessageRequestLinkResourceType = "agent_definition"
	SendMessageRequestLinkResourceTypeAvailableTool                        SendMessageRequestLinkResourceType = "available_tool"
	SendMessageRequestLinkResourceTypeAgentDefinitionTool                  SendMessageRequestLinkResourceType = "agent_definition_tool"
	SendMessageRequestLinkResourceTypeAgentAccountStatus                   SendMessageRequestLinkResourceType = "agent_account_status"
	SendMessageRequestLinkResourceTypeAgentRun                             SendMessageRequestLinkResourceType = "agent_run"
	SendMessageRequestLinkResourceTypeAgentAction                          SendMessageRequestLinkResourceType = "agent_action"
	SendMessageRequestLinkResourceTypeAgentRunStep                         SendMessageRequestLinkResourceType = "agent_run_step"
	SendMessageRequestLinkResourceTypeAgentTokenUsage                      SendMessageRequestLinkResourceType = "agent_token_usage"
	SendMessageRequestLinkResourceTypeAgentMemory                          SendMessageRequestLinkResourceType = "agent_memory"
	SendMessageRequestLinkResourceTypeNotification                         SendMessageRequestLinkResourceType = "notification"
	SendMessageRequestLinkResourceTypeNotificationUnreadCount              SendMessageRequestLinkResourceType = "notification_unread_count"
	SendMessageRequestLinkResourceTypeNotificationSendResult               SendMessageRequestLinkResourceType = "notification_send_result"
	SendMessageRequestLinkResourceTypeNotificationUnreadSummary            SendMessageRequestLinkResourceType = "notification_unread_summary"
	SendMessageRequestLinkResourceTypeAnnouncement                         SendMessageRequestLinkResourceType = "announcement"
	SendMessageRequestLinkResourceTypeConversation                         SendMessageRequestLinkResourceType = "conversation"
	SendMessageRequestLinkResourceTypeSupportCase                          SendMessageRequestLinkResourceType = "support_case"
	SendMessageRequestLinkResourceTypeConversationParticipant              SendMessageRequestLinkResourceType = "conversation_participant"
	SendMessageRequestLinkResourceTypeReadCursor                           SendMessageRequestLinkResourceType = "read_cursor"
	SendMessageRequestLinkResourceTypeChatMessage                          SendMessageRequestLinkResourceType = "chat_message"
	SendMessageRequestLinkResourceTypeNotificationUnreadSummaryAccount     SendMessageRequestLinkResourceType = "notification_unread_summary_account"
	SendMessageRequestLinkResourceTypeMessagingBlock                       SendMessageRequestLinkResourceType = "messaging_block"
	SendMessageRequestLinkResourceTypeNotificationPreference               SendMessageRequestLinkResourceType = "notification_preference"
	SendMessageRequestLinkResourceTypeMessageAttachment                    SendMessageRequestLinkResourceType = "message_attachment"
	SendMessageRequestLinkResourceTypeAttachmentUploadTarget               SendMessageRequestLinkResourceType = "attachment_upload_target"
	SendMessageRequestLinkResourceTypeScheduledMessage                     SendMessageRequestLinkResourceType = "scheduled_message"
	SendMessageRequestLinkResourceTypeMessagingContact                     SendMessageRequestLinkResourceType = "messaging_contact"
	SendMessageRequestLinkResourceTypeMessageReport                        SendMessageRequestLinkResourceType = "message_report"
	SendMessageRequestLinkResourceTypeToolGroup                            SendMessageRequestLinkResourceType = "tool_group"
	SendMessageRequestLinkResourceTypeModel                                SendMessageRequestLinkResourceType = "model"
	SendMessageRequestLinkResourceTypePaymentTerm                          SendMessageRequestLinkResourceType = "payment_term"
	SendMessageRequestLinkResourceTypeShippingTerm                         SendMessageRequestLinkResourceType = "shipping_term"
	SendMessageRequestLinkResourceTypeQuantity                             SendMessageRequestLinkResourceType = "quantity"
	SendMessageRequestLinkResourceTypeAccountGroup                         SendMessageRequestLinkResourceType = "account_group"
	SendMessageRequestLinkResourceTypeSupportRoute                         SendMessageRequestLinkResourceType = "support_route"
	SendMessageRequestLinkResourceTypeSupportAvailability                  SendMessageRequestLinkResourceType = "support_availability"
	SendMessageRequestLinkResourceTypeAccountStatus                        SendMessageRequestLinkResourceType = "account_status"
	SendMessageRequestLinkResourceTypeGeolocation                          SendMessageRequestLinkResourceType = "geolocation"
	SendMessageRequestLinkResourceTypeAccountUser                          SendMessageRequestLinkResourceType = "account_user"
	SendMessageRequestLinkResourceTypeDepartment                           SendMessageRequestLinkResourceType = "department"
	SendMessageRequestLinkResourceTypeAccountIntegration                   SendMessageRequestLinkResourceType = "account_integration"
	SendMessageRequestLinkResourceTypeAccountPrice                         SendMessageRequestLinkResourceType = "account_price"
	SendMessageRequestLinkResourceTypeProductLine                          SendMessageRequestLinkResourceType = "product_line"
	SendMessageRequestLinkResourceTypeItemCategory                         SendMessageRequestLinkResourceType = "item_category"
	SendMessageRequestLinkResourceTypeAttribute                            SendMessageRequestLinkResourceType = "attribute"
	SendMessageRequestLinkResourceTypeRate                                 SendMessageRequestLinkResourceType = "rate"
	SendMessageRequestLinkResourceTypeAccountGroupProductLineAccess        SendMessageRequestLinkResourceType = "account_group_product_line_access"
	SendMessageRequestLinkResourceTypeSalesTarget                          SendMessageRequestLinkResourceType = "sales_target"
	SendMessageRequestLinkResourceTypeAdjustmentType                       SendMessageRequestLinkResourceType = "adjustment_type"
	SendMessageRequestLinkResourceTypeAccountBranding                      SendMessageRequestLinkResourceType = "account_branding"
	SendMessageRequestLinkResourceTypeAccountPortal                        SendMessageRequestLinkResourceType = "account_portal"
	SendMessageRequestLinkResourceTypeAccountLogoURL                       SendMessageRequestLinkResourceType = "account_logo_url"
	SendMessageRequestLinkResourceTypeAccountFaviconURL                    SendMessageRequestLinkResourceType = "account_favicon_url"
	SendMessageRequestLinkResourceTypePublicAccount                        SendMessageRequestLinkResourceType = "public_account"
	SendMessageRequestLinkResourceTypeProperty                             SendMessageRequestLinkResourceType = "property"
	SendMessageRequestLinkResourceTypeCarrier                              SendMessageRequestLinkResourceType = "carrier"
	SendMessageRequestLinkResourceTypeServiceLevel                         SendMessageRequestLinkResourceType = "service_level"
	SendMessageRequestLinkResourceTypeItem                                 SendMessageRequestLinkResourceType = "item"
	SendMessageRequestLinkResourceTypeItemLotDefault                       SendMessageRequestLinkResourceType = "item_lot_default"
	SendMessageRequestLinkResourceTypeItemInventory                        SendMessageRequestLinkResourceType = "item_inventory"
	SendMessageRequestLinkResourceTypeProduct                              SendMessageRequestLinkResourceType = "product"
	SendMessageRequestLinkResourceTypeBatch                                SendMessageRequestLinkResourceType = "batch"
	SendMessageRequestLinkResourceTypeBatchFlowNode                        SendMessageRequestLinkResourceType = "batch_flow_node"
	SendMessageRequestLinkResourceTypeScanningConsumption                  SendMessageRequestLinkResourceType = "scanning_consumption"
	SendMessageRequestLinkResourceTypeOpenBatchSummary                     SendMessageRequestLinkResourceType = "open_batch_summary"
	SendMessageRequestLinkResourceTypeScanningProductionStepInfo           SendMessageRequestLinkResourceType = "scanning_production_step_info"
	SendMessageRequestLinkResourceTypeScanningStation                      SendMessageRequestLinkResourceType = "scanning_station"
	SendMessageRequestLinkResourceTypeProductionStep                       SendMessageRequestLinkResourceType = "production_step"
	SendMessageRequestLinkResourceTypeProductionRun                        SendMessageRequestLinkResourceType = "production_run"
	SendMessageRequestLinkResourceTypeMachine                              SendMessageRequestLinkResourceType = "machine"
	SendMessageRequestLinkResourceTypeMachineStatus                        SendMessageRequestLinkResourceType = "machine_status"
	SendMessageRequestLinkResourceTypeMachineDowntimeEvent                 SendMessageRequestLinkResourceType = "machine_downtime_event"
	SendMessageRequestLinkResourceTypeDemandOverride                       SendMessageRequestLinkResourceType = "demand_override"
	SendMessageRequestLinkResourceTypeDemandOverrideType                   SendMessageRequestLinkResourceType = "demand_override_type"
	SendMessageRequestLinkResourceTypeMachineDowntimeReason                SendMessageRequestLinkResourceType = "machine_downtime_reason"
	SendMessageRequestLinkResourceTypeProductionSchedulePreview            SendMessageRequestLinkResourceType = "production_schedule_preview"
	SendMessageRequestLinkResourceTypeProductionScheduleRegeneratePreview  SendMessageRequestLinkResourceType = "production_schedule_regenerate_preview"
	SendMessageRequestLinkResourceTypeProductionSchedule                   SendMessageRequestLinkResourceType = "production_schedule"
	SendMessageRequestLinkResourceTypeProductionScheduleLine               SendMessageRequestLinkResourceType = "production_schedule_line"
	SendMessageRequestLinkResourceTypeProductionScheduleDeviation          SendMessageRequestLinkResourceType = "production_schedule_deviation"
	SendMessageRequestLinkResourceTypeProductionScheduleDerivedLine        SendMessageRequestLinkResourceType = "production_schedule_derived_line"
	SendMessageRequestLinkResourceTypeProductionScheduleSettings           SendMessageRequestLinkResourceType = "production_schedule_settings"
	SendMessageRequestLinkResourceTypeProductionScheduleResourceSetting    SendMessageRequestLinkResourceType = "production_schedule_resource_setting"
	SendMessageRequestLinkResourceTypeProductionScheduleItemSetting        SendMessageRequestLinkResourceType = "production_schedule_item_setting"
	SendMessageRequestLinkResourceTypeFulfillmentRecommendation            SendMessageRequestLinkResourceType = "fulfillment_recommendation"
	SendMessageRequestLinkResourceTypeAnalyzeDeliveryPerformanceResponse   SendMessageRequestLinkResourceType = "analyze_delivery_performance_response"
	SendMessageRequestLinkResourceTypeDeliveryPerformance                  SendMessageRequestLinkResourceType = "delivery_performance"
	SendMessageRequestLinkResourceTypeDeliveryBacklogBucket                SendMessageRequestLinkResourceType = "delivery_backlog_bucket"
	SendMessageRequestLinkResourceTypeDeliveryLatenessBucket               SendMessageRequestLinkResourceType = "delivery_lateness_bucket"
	SendMessageRequestLinkResourceTypeDeliveryBreakdown                    SendMessageRequestLinkResourceType = "delivery_breakdown"
	SendMessageRequestLinkResourceTypeAnalyzeSalesBreakdownResponse        SendMessageRequestLinkResourceType = "analyze_sales_breakdown_response"
	SendMessageRequestLinkResourceTypeSalesTotals                          SendMessageRequestLinkResourceType = "sales_totals"
	SendMessageRequestLinkResourceTypeSalesBreakdown                       SendMessageRequestLinkResourceType = "sales_breakdown"
	SendMessageRequestLinkResourceTypeScheduleOrderCoverage                SendMessageRequestLinkResourceType = "schedule_order_coverage"
	SendMessageRequestLinkResourceTypeScheduleOrderCoverageLine            SendMessageRequestLinkResourceType = "schedule_order_coverage_line"
	SendMessageRequestLinkResourceTypeScheduleDeviationType                SendMessageRequestLinkResourceType = "schedule_deviation_type"
	SendMessageRequestLinkResourceTypeScheduleAtRiskOrder                  SendMessageRequestLinkResourceType = "schedule_at_risk_order"
	SendMessageRequestLinkResourceTypeProductionScheduleFinishedPolicy     SendMessageRequestLinkResourceType = "production_schedule_finished_policy"
	SendMessageRequestLinkResourceTypeProductionScheduleFinishingLine      SendMessageRequestLinkResourceType = "production_schedule_finishing_line"
	SendMessageRequestLinkResourceTypeProductionScheduleWeekRelease        SendMessageRequestLinkResourceType = "production_schedule_week_release"
	SendMessageRequestLinkResourceTypeProductionScheduleWeekReleasePreview SendMessageRequestLinkResourceType = "production_schedule_week_release_preview"
	SendMessageRequestLinkResourceTypeProductionScheduleItemPolicy         SendMessageRequestLinkResourceType = "production_schedule_item_policy"
	SendMessageRequestLinkResourceTypeChildAccount                         SendMessageRequestLinkResourceType = "child_account"
	SendMessageRequestLinkResourceTypeUnitGroup                            SendMessageRequestLinkResourceType = "unit_group"
	SendMessageRequestLinkResourceTypeUnitGroupUnit                        SendMessageRequestLinkResourceType = "unit_group_unit"
	SendMessageRequestLinkResourceTypeConsumption                          SendMessageRequestLinkResourceType = "consumption"
	SendMessageRequestLinkResourceTypeCustomerProductLineAccess            SendMessageRequestLinkResourceType = "customer_product_line_access"
	SendMessageRequestLinkResourceTypeCustomer                             SendMessageRequestLinkResourceType = "customer"
	SendMessageRequestLinkResourceTypeFrequentlyOrderedProduct             SendMessageRequestLinkResourceType = "frequently_ordered_product"
	SendMessageRequestLinkResourceTypePriority                             SendMessageRequestLinkResourceType = "priority"
	SendMessageRequestLinkResourceTypeDelivery                             SendMessageRequestLinkResourceType = "delivery"
	SendMessageRequestLinkResourceTypeDeliveryLine                         SendMessageRequestLinkResourceType = "delivery_line"
	SendMessageRequestLinkResourceTypeSalesOrder                           SendMessageRequestLinkResourceType = "sales_order"
	SendMessageRequestLinkResourceTypeLocation                             SendMessageRequestLinkResourceType = "location"
	SendMessageRequestLinkResourceTypeLocationType                         SendMessageRequestLinkResourceType = "location_type"
	SendMessageRequestLinkResourceTypeLot                                  SendMessageRequestLinkResourceType = "lot"
	SendMessageRequestLinkResourceTypeEmailLog                             SendMessageRequestLinkResourceType = "email_log"
	SendMessageRequestLinkResourceTypeEmailDomain                          SendMessageRequestLinkResourceType = "email_domain"
	SendMessageRequestLinkResourceTypeEmailInbox                           SendMessageRequestLinkResourceType = "email_inbox"
	SendMessageRequestLinkResourceTypePortalDomain                         SendMessageRequestLinkResourceType = "portal_domain"
	SendMessageRequestLinkResourceTypeDNSRecord                            SendMessageRequestLinkResourceType = "dns_record"
	SendMessageRequestLinkResourceTypeInventoryChangeLog                   SendMessageRequestLinkResourceType = "inventory_change_log"
	SendMessageRequestLinkResourceTypeInvoice                              SendMessageRequestLinkResourceType = "invoice"
	SendMessageRequestLinkResourceTypeInvoiceSummary                       SendMessageRequestLinkResourceType = "invoice_summary"
	SendMessageRequestLinkResourceTypeInvoiceLine                          SendMessageRequestLinkResourceType = "invoice_line"
	SendMessageRequestLinkResourceTypeInvoiceAllocation                    SendMessageRequestLinkResourceType = "invoice_allocation"
	SendMessageRequestLinkResourceTypeInvoiceForPayment                    SendMessageRequestLinkResourceType = "invoice_for_payment"
	SendMessageRequestLinkResourceTypeShipment                             SendMessageRequestLinkResourceType = "shipment"
	SendMessageRequestLinkResourceTypeShipmentSummary                      SendMessageRequestLinkResourceType = "shipment_summary"
	SendMessageRequestLinkResourceTypeShipmentLine                         SendMessageRequestLinkResourceType = "shipment_line"
	SendMessageRequestLinkResourceTypeShippingCase                         SendMessageRequestLinkResourceType = "shipping_case"
	SendMessageRequestLinkResourceTypeShippingCaseLabelURL                 SendMessageRequestLinkResourceType = "shipping_case_label_url"
	SendMessageRequestLinkResourceTypeSettlement                           SendMessageRequestLinkResourceType = "settlement"
	SendMessageRequestLinkResourceTypeSettlementSummary                    SendMessageRequestLinkResourceType = "settlement_summary"
	SendMessageRequestLinkResourceTypeRolePermission                       SendMessageRequestLinkResourceType = "role_permission"
	SendMessageRequestLinkResourceTypeRegistrationFlow                     SendMessageRequestLinkResourceType = "registration_flow"
	SendMessageRequestLinkResourceTypeRegistrationFlowOption               SendMessageRequestLinkResourceType = "registration_flow_option"
	SendMessageRequestLinkResourceTypeTransaction                          SendMessageRequestLinkResourceType = "transaction"
	SendMessageRequestLinkResourceTypeTransactionSummary                   SendMessageRequestLinkResourceType = "transaction_summary"
	SendMessageRequestLinkResourceTypeTransactionMethod                    SendMessageRequestLinkResourceType = "transaction_method"
	SendMessageRequestLinkResourceTypeTransactionType                      SendMessageRequestLinkResourceType = "transaction_type"
	SendMessageRequestLinkResourceTypeTransactionAllocation                SendMessageRequestLinkResourceType = "transaction_allocation"
	SendMessageRequestLinkResourceTypeUsageItem                            SendMessageRequestLinkResourceType = "usage_item"
	SendMessageRequestLinkResourceTypeAccountUsageResponse                 SendMessageRequestLinkResourceType = "account_usage_response"
	SendMessageRequestLinkResourceTypeSubscriptionInfo                     SendMessageRequestLinkResourceType = "subscription_info"
	SendMessageRequestLinkResourceTypeBillingPortalSessionResponse         SendMessageRequestLinkResourceType = "billing_portal_session_response"
	SendMessageRequestLinkResourceTypeSwitchPlanResponse                   SendMessageRequestLinkResourceType = "switch_plan_response"
	SendMessageRequestLinkResourceTypeEnsureBillingCustomerResponse        SendMessageRequestLinkResourceType = "ensure_billing_customer_response"
	SendMessageRequestLinkResourceTypeSpendingCapResponse                  SendMessageRequestLinkResourceType = "spending_cap_response"
	SendMessageRequestLinkResourceTypeAgentSpendInfo                       SendMessageRequestLinkResourceType = "agent_spend_info"
	SendMessageRequestLinkResourceTypeWebhookResponse                      SendMessageRequestLinkResourceType = "webhook_response"
	SendMessageRequestLinkResourceTypeAddressSuggestion                    SendMessageRequestLinkResourceType = "address_suggestion"
	SendMessageRequestLinkResourceTypeAddressComponents                    SendMessageRequestLinkResourceType = "address_components"
	SendMessageRequestLinkResourceTypeAddressDetailsResult                 SendMessageRequestLinkResourceType = "address_details_result"
	SendMessageRequestLinkResourceTypeValidatedAddress                     SendMessageRequestLinkResourceType = "validated_address"
	SendMessageRequestLinkResourceTypePlanLimit                            SendMessageRequestLinkResourceType = "plan_limit"
	SendMessageRequestLinkResourceTypePlanChangeProration                  SendMessageRequestLinkResourceType = "plan_change_proration"
	SendMessageRequestLinkResourceTypePlanChangeLineItem                   SendMessageRequestLinkResourceType = "plan_change_line_item"
	SendMessageRequestLinkResourceTypeSetupBillingResponse                 SendMessageRequestLinkResourceType = "setup_billing_response"
	SendMessageRequestLinkResourceTypeConfirmPaymentResponse               SendMessageRequestLinkResourceType = "confirm_payment_response"
	SendMessageRequestLinkResourceTypeOAuthResponse                        SendMessageRequestLinkResourceType = "oauth_response"
	SendMessageRequestLinkResourceTypeOAuthStatusResponse                  SendMessageRequestLinkResourceType = "oauth_status_response"
	SendMessageRequestLinkResourceTypeStripePublishableKey                 SendMessageRequestLinkResourceType = "stripe_publishable_key"
	SendMessageRequestLinkResourceTypeStripeStatus                         SendMessageRequestLinkResourceType = "stripe_status"
	SendMessageRequestLinkResourceTypeHealthcheck                          SendMessageRequestLinkResourceType = "healthcheck"
	SendMessageRequestLinkResourceTypeAgentDefinitionConfig                SendMessageRequestLinkResourceType = "agent_definition_config"
	SendMessageRequestLinkResourceTypeTriggerConfig                        SendMessageRequestLinkResourceType = "trigger_config"
	SendMessageRequestLinkResourceTypeCustomerContactInfo                  SendMessageRequestLinkResourceType = "customer_contact_info"
	SendMessageRequestLinkResourceTypeCustomerFreightPreferences           SendMessageRequestLinkResourceType = "customer_freight_preferences"
	SendMessageRequestLinkResourceTypeCustomerDefaults                     SendMessageRequestLinkResourceType = "customer_defaults"
	SendMessageRequestLinkResourceTypeCustomerLeadTime                     SendMessageRequestLinkResourceType = "customer_lead_time"
	SendMessageRequestLinkResourceTypeCustomerNotificationPreferences      SendMessageRequestLinkResourceType = "customer_notification_preferences"
	SendMessageRequestLinkResourceTypeOrderNotificationRecipient           SendMessageRequestLinkResourceType = "order_notification_recipient"
	SendMessageRequestLinkResourceTypeOrderDiscount                        SendMessageRequestLinkResourceType = "order_discount"
	SendMessageRequestLinkResourceTypeSalesOrderLine                       SendMessageRequestLinkResourceType = "sales_order_line"
	SendMessageRequestLinkResourceTypeSalesOrderType                       SendMessageRequestLinkResourceType = "sales_order_type"
	SendMessageRequestLinkResourceTypeSalesOrderStatus                     SendMessageRequestLinkResourceType = "sales_order_status"
	SendMessageRequestLinkResourceTypeMaterial                             SendMessageRequestLinkResourceType = "material"
	SendMessageRequestLinkResourceTypeSupplierMaterial                     SendMessageRequestLinkResourceType = "supplier_material"
	SendMessageRequestLinkResourceTypePart                                 SendMessageRequestLinkResourceType = "part"
	SendMessageRequestLinkResourceTypePermissionGroup                      SendMessageRequestLinkResourceType = "permission_group"
	SendMessageRequestLinkResourceTypePermission                           SendMessageRequestLinkResourceType = "permission"
	SendMessageRequestLinkResourceTypePick                                 SendMessageRequestLinkResourceType = "pick"
	SendMessageRequestLinkResourceTypePickLine                             SendMessageRequestLinkResourceType = "pick_line"
	SendMessageRequestLinkResourceTypeProductType                          SendMessageRequestLinkResourceType = "product_type"
	SendMessageRequestLinkResourceTypeProduction                           SendMessageRequestLinkResourceType = "production"
	SendMessageRequestLinkResourceTypeProductionFlow                       SendMessageRequestLinkResourceType = "production_flow"
	SendMessageRequestLinkResourceTypeMap                                  SendMessageRequestLinkResourceType = "map"
	SendMessageRequestLinkResourceTypePurchaseOrder                        SendMessageRequestLinkResourceType = "purchase_order"
	SendMessageRequestLinkResourceTypePurchaseOrderLine                    SendMessageRequestLinkResourceType = "purchase_order_line"
	SendMessageRequestLinkResourceTypeSupplier                             SendMessageRequestLinkResourceType = "supplier"
	SendMessageRequestLinkResourceTypeSupplierSummary                      SendMessageRequestLinkResourceType = "supplier_summary"
	SendMessageRequestLinkResourceTypeReceivableEntry                      SendMessageRequestLinkResourceType = "receivable_entry"
	SendMessageRequestLinkResourceTypeReceivingOrder                       SendMessageRequestLinkResourceType = "receiving_order"
	SendMessageRequestLinkResourceTypeReceivingOrderLine                   SendMessageRequestLinkResourceType = "receiving_order_line"
	SendMessageRequestLinkResourceTypeEmailContact                         SendMessageRequestLinkResourceType = "email_contact"
	SendMessageRequestLinkResourceTypeAllocationEntry                      SendMessageRequestLinkResourceType = "allocation_entry"
	SendMessageRequestLinkResourceTypeOpenCreditEntry                      SendMessageRequestLinkResourceType = "open_credit_entry"
	SendMessageRequestLinkResourceTypeVolumeDiscount                       SendMessageRequestLinkResourceType = "volume_discount"
	SendMessageRequestLinkResourceTypeVolumeDiscountTier                   SendMessageRequestLinkResourceType = "volume_discount_tier"
	SendMessageRequestLinkResourceTypeAnalyzeDeliveriesResponse            SendMessageRequestLinkResourceType = "analyze_deliveries_response"
	SendMessageRequestLinkResourceTypeAnalyzeManufacturingResponse         SendMessageRequestLinkResourceType = "analyze_manufacturing_response"
	SendMessageRequestLinkResourceTypeAnalyzeManufacturingBatchResponse    SendMessageRequestLinkResourceType = "analyze_manufacturing_batch_response"
	SendMessageRequestLinkResourceTypeAnalyzeQuarterlyOrdersResponse       SendMessageRequestLinkResourceType = "analyze_quarterly_orders_response"
	SendMessageRequestLinkResourceTypeAnalyzeNewCustomersResponse          SendMessageRequestLinkResourceType = "analyze_new_customers_response"
	SendMessageRequestLinkResourceTypeAnalyzeDemandForecastResponse        SendMessageRequestLinkResourceType = "analyze_demand_forecast_response"
	SendMessageRequestLinkResourceTypeAnalyzeOeeResponse                   SendMessageRequestLinkResourceType = "analyze_oee_response"
	SendMessageRequestLinkResourceTypeAnalyzeOeeTrendResponse              SendMessageRequestLinkResourceType = "analyze_oee_trend_response"
	SendMessageRequestLinkResourceTypeAnalyzeScheduleAttainmentResponse    SendMessageRequestLinkResourceType = "analyze_schedule_attainment_response"
	SendMessageRequestLinkResourceTypeCatalogProductLine                   SendMessageRequestLinkResourceType = "catalog_product_line"
	SendMessageRequestLinkResourceTypeCatalogCategory                      SendMessageRequestLinkResourceType = "catalog_category"
	SendMessageRequestLinkResourceTypeCatalogProduct                       SendMessageRequestLinkResourceType = "catalog_product"
	SendMessageRequestLinkResourceTypeCatalogProperty                      SendMessageRequestLinkResourceType = "catalog_property"
	SendMessageRequestLinkResourceTypeCatalogAttribute                     SendMessageRequestLinkResourceType = "catalog_attribute"
	SendMessageRequestLinkResourceTypeDcLocation                           SendMessageRequestLinkResourceType = "dc_location"
	SendMessageRequestLinkResourceTypeEdiRun                               SendMessageRequestLinkResourceType = "edi_run"
	SendMessageRequestLinkResourceTypeInventoryItem                        SendMessageRequestLinkResourceType = "inventory_item"
	SendMessageRequestLinkResourceTypeAnalyzeWeeksOfSalesResponse          SendMessageRequestLinkResourceType = "analyze_weeks_of_sales_response"
	SendMessageRequestLinkResourceTypeBulkReconcileItemsResponse           SendMessageRequestLinkResourceType = "bulk_reconcile_items_response"
	SendMessageRequestLinkResourceTypeSysProperty                          SendMessageRequestLinkResourceType = "sys_property"
	SendMessageRequestLinkResourceTypeSysPropertyType                      SendMessageRequestLinkResourceType = "sys_property_type"
	SendMessageRequestLinkResourceTypeSysPropertyValue                     SendMessageRequestLinkResourceType = "sys_property_value"
	SendMessageRequestLinkResourceTypeTerritory                            SendMessageRequestLinkResourceType = "territory"
	SendMessageRequestLinkResourceTypeTenancy                              SendMessageRequestLinkResourceType = "tenancy"
	SendMessageRequestLinkResourceTypeCheckoutSession                      SendMessageRequestLinkResourceType = "checkout_session"
	SendMessageRequestLinkResourceTypeEstimateRateResult                   SendMessageRequestLinkResourceType = "estimate_rate_result"
	SendMessageRequestLinkResourceTypeRateShopOption                       SendMessageRequestLinkResourceType = "rate_shop_option"
	SendMessageRequestLinkResourceTypeRateShopResult                       SendMessageRequestLinkResourceType = "rate_shop_result"
	SendMessageRequestLinkResourceTypeOwner                                SendMessageRequestLinkResourceType = "owner"
	SendMessageRequestLinkResourceTypeCreatedBy                            SendMessageRequestLinkResourceType = "created_by"
	SendMessageRequestLinkResourceTypeMessage                              SendMessageRequestLinkResourceType = "message"
	SendMessageRequestLinkResourceTypeAccountPhotoUploadResult             SendMessageRequestLinkResourceType = "account_photo_upload_result"
	SendMessageRequestLinkResourceTypeUserPhotoUploadResult                SendMessageRequestLinkResourceType = "user_photo_upload_result"
	SendMessageRequestLinkResourceTypeUserPhotoURL                         SendMessageRequestLinkResourceType = "user_photo_url"
	SendMessageRequestLinkResourceTypeBatchLot                             SendMessageRequestLinkResourceType = "batch_lot"
	SendMessageRequestLinkResourceTypeCheckDuplicateResult                 SendMessageRequestLinkResourceType = "check_duplicate_result"
	SendMessageRequestLinkResourceTypeItemTrendPoint                       SendMessageRequestLinkResourceType = "item_trend_point"
	SendMessageRequestLinkResourceTypeTenancyPendingRegistration           SendMessageRequestLinkResourceType = "tenancy_pending_registration"
	SendMessageRequestLinkResourceTypeInvoiceAllocationEntry               SendMessageRequestLinkResourceType = "invoice_allocation_entry"
	SendMessageRequestLinkResourceTypeAllocationCustomer                   SendMessageRequestLinkResourceType = "allocation_customer"
	SendMessageRequestLinkResourceTypeCheckoutSalesOrder                   SendMessageRequestLinkResourceType = "checkout_sales_order"
	SendMessageRequestLinkResourceTypeSalesOrderPriceQuote                 SendMessageRequestLinkResourceType = "sales_order_price_quote"
	SendMessageRequestLinkResourceTypeSalesOrderFreightQuote               SendMessageRequestLinkResourceType = "sales_order_freight_quote"
	SendMessageRequestLinkResourceTypeSalesOrderCommitmentQuote            SendMessageRequestLinkResourceType = "sales_order_commitment_quote"
	SendMessageRequestLinkResourceTypeOperatingCalendar                    SendMessageRequestLinkResourceType = "operating_calendar"
	SendMessageRequestLinkResourceTypeOperatingCalendarClosure             SendMessageRequestLinkResourceType = "operating_calendar_closure"
	SendMessageRequestLinkResourceTypeSalesOrderPriceQuoteLine             SendMessageRequestLinkResourceType = "sales_order_price_quote_line"
	SendMessageRequestLinkResourceTypeHubspotSyncJob                       SendMessageRequestLinkResourceType = "hubspot_sync_job"
	SendMessageRequestLinkResourceTypeHubspotSyncReport                    SendMessageRequestLinkResourceType = "hubspot_sync_report"
	SendMessageRequestLinkResourceTypeHubspotCompanyReview                 SendMessageRequestLinkResourceType = "hubspot_company_review"
	SendMessageRequestLinkResourceTypeHubspotCompanyCandidate              SendMessageRequestLinkResourceType = "hubspot_company_candidate"
	SendMessageRequestLinkResourceTypeHubspotSyncRecord                    SendMessageRequestLinkResourceType = "hubspot_sync_record"
	SendMessageRequestLinkResourceTypeContactMatch                         SendMessageRequestLinkResourceType = "contact_match"
	SendMessageRequestLinkResourceTypeReplyDraft                           SendMessageRequestLinkResourceType = "reply_draft"
	SendMessageRequestLinkResourceTypeConversationLink                     SendMessageRequestLinkResourceType = "conversation_link"
	SendMessageRequestLinkResourceTypeMessagingGroup                       SendMessageRequestLinkResourceType = "messaging_group"
	SendMessageRequestLinkResourceTypeMessagingGroupMember                 SendMessageRequestLinkResourceType = "messaging_group_member"
	SendMessageRequestLinkResourceTypePortalProfile                        SendMessageRequestLinkResourceType = "portal_profile"
	SendMessageRequestLinkResourceTypePortalRegistrationSession            SendMessageRequestLinkResourceType = "portal_registration_session"
	SendMessageRequestLinkResourceTypePortalRegistrationSessionData        SendMessageRequestLinkResourceType = "portal_registration_session_data"
	SendMessageRequestLinkResourceTypePackList                             SendMessageRequestLinkResourceType = "pack_list"
	SendMessageRequestLinkResourceTypePackListParty                        SendMessageRequestLinkResourceType = "pack_list_party"
	SendMessageRequestLinkResourceTypePackListLineItem                     SendMessageRequestLinkResourceType = "pack_list_line_item"
	SendMessageRequestLinkResourceTypePackListBackOrder                    SendMessageRequestLinkResourceType = "pack_list_back_order"
	SendMessageRequestLinkResourceTypePackListCase                         SendMessageRequestLinkResourceType = "pack_list_case"
	SendMessageRequestLinkResourceTypeJob                                  SendMessageRequestLinkResourceType = "job"
	SendMessageRequestLinkResourceTypeJobResult                            SendMessageRequestLinkResourceType = "job_result"
	SendMessageRequestLinkResourceTypeJobExport                            SendMessageRequestLinkResourceType = "job_export"
	SendMessageRequestLinkResourceTypeAnalyzeCustomerPricingResponse       SendMessageRequestLinkResourceType = "analyze_customer_pricing_response"
	SendMessageRequestLinkResourceTypeCustomerPricingFinding               SendMessageRequestLinkResourceType = "customer_pricing_finding"
	SendMessageRequestLinkResourceTypeCustomerPricingSummary               SendMessageRequestLinkResourceType = "customer_pricing_summary"
	SendMessageRequestLinkResourceTypeComputedRate                         SendMessageRequestLinkResourceType = "computed_rate"
	SendMessageRequestLinkResourceTypeComputedQuantity                     SendMessageRequestLinkResourceType = "computed_quantity"
	SendMessageRequestLinkResourceTypeAnalyzeRealizedMarginsResponse       SendMessageRequestLinkResourceType = "analyze_realized_margins_response"
	SendMessageRequestLinkResourceTypeRealizedMarginFinding                SendMessageRequestLinkResourceType = "realized_margin_finding"
	SendMessageRequestLinkResourceTypeRealizedMarginSummary                SendMessageRequestLinkResourceType = "realized_margin_summary"
	SendMessageRequestLinkResourceTypeShipmentRelated                      SendMessageRequestLinkResourceType = "shipment_related"
	SendMessageRequestLinkResourceTypeInvoiceRelated                       SendMessageRequestLinkResourceType = "invoice_related"
	SendMessageRequestLinkResourceTypePickRelated                          SendMessageRequestLinkResourceType = "pick_related"
	SendMessageRequestLinkResourceTypePickTotals                           SendMessageRequestLinkResourceType = "pick_totals"
	SendMessageRequestLinkResourceTypePickStageTotal                       SendMessageRequestLinkResourceType = "pick_stage_total"
)

// Whether to deliver the message now or hold it as a customer-reply draft.
//
//   - `send`: delivers the message, immediately or at `scheduled_at`.
//   - `draft`: proposes a reply to the customer on a customer-facing case and holds
//     it for a teammate to approve before it goes out. Requires `channel`.
//
// A draft is built from `body`, `subject`, `channel`, and
// `source_thread_message_id` only — attachments, mentions, copied recipients,
// resource links, replies, and scheduling are not carried onto it.
type SendMessageRequestMode string

const (
	SendMessageRequestModeSend  SendMessageRequestMode = "send"
	SendMessageRequestModeDraft SendMessageRequestMode = "draft"
)

type MessagingConversationMessageNewParams struct {
	// Request to post a message to a conversation.
	SendMessageRequest SendMessageRequestParam
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

func (r MessagingConversationMessageNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SendMessageRequest)
}
func (r *MessagingConversationMessageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationMessageNewParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationMessageNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationMessageListParams struct {
	// Return only messages that come after this position in the timeline.
	//
	// Use it to catch up after a dropped realtime connection: pass the sequence of the
	// last message you already have to fetch everything since.
	AfterSequence param.Opt[int64] `query:"after_sequence,omitzero" json:"-"`
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "author", "resource", "attachments", "attachments.resource",
	// "conversation", "conversation.participants", "conversation.last_message",
	// "reply_to", "reply_to.sender", "reply_to.author", "reply_to.attachments",
	// "agent_run".
	Include []string `query:"include,omitzero" json:"-"`
	// Which set of the conversation's messages to return.
	//
	// Left unset, you get the delivered timeline. Pass `draft` for the case's reply
	// drafts awaiting approval, or `scheduled` for the messages you yourself have
	// queued for a future send, soonest first. Those two ignore paging and come back
	// in a single response.
	//
	// Any of "draft", "scheduled", "sent", "canceled", "rejected", "failed",
	// "superseded".
	Status MessagingConversationMessageListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationMessageListParams]'s query parameters
// as `url.Values`.
func (r MessagingConversationMessageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which set of the conversation's messages to return.
//
// Left unset, you get the delivered timeline. Pass `draft` for the case's reply
// drafts awaiting approval, or `scheduled` for the messages you yourself have
// queued for a future send, soonest first. Those two ignore paging and come back
// in a single response.
type MessagingConversationMessageListParamsStatus string

const (
	MessagingConversationMessageListParamsStatusDraft      MessagingConversationMessageListParamsStatus = "draft"
	MessagingConversationMessageListParamsStatusScheduled  MessagingConversationMessageListParamsStatus = "scheduled"
	MessagingConversationMessageListParamsStatusSent       MessagingConversationMessageListParamsStatus = "sent"
	MessagingConversationMessageListParamsStatusCanceled   MessagingConversationMessageListParamsStatus = "canceled"
	MessagingConversationMessageListParamsStatusRejected   MessagingConversationMessageListParamsStatus = "rejected"
	MessagingConversationMessageListParamsStatusFailed     MessagingConversationMessageListParamsStatus = "failed"
	MessagingConversationMessageListParamsStatusSuperseded MessagingConversationMessageListParamsStatus = "superseded"
)
