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

// List, read, and manage in-app notifications.
//
// MessagingNotificationService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingNotificationService] method instead.
type MessagingNotificationService struct {
	options []option.RequestOption
	// List, read, and manage in-app notifications.
	Actions MessagingNotificationActionService
}

// NewMessagingNotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingNotificationService(opts ...option.RequestOption) (r MessagingNotificationService) {
	r = MessagingNotificationService{}
	r.options = opts
	r.Actions = NewMessagingNotificationActionService(opts...)
	return
}

// Sends an in-app notification to a single member of an account, or announces it
// to everyone in the account.
//
// A send to one member is attributed to the authenticated caller, so the recipient
// sees who sent it. It is accepted and then fanned out, so it reaches the
// recipient's feed and their connected clients shortly after the response.
//
// An announcement to the whole account is stored as the request is accepted,
// carries no sender, and may only target the account you are currently acting in.
//
// This endpoint requires the permission: `alerts:create`.
func (r *MessagingNotificationService) New(ctx context.Context, body MessagingNotificationNewParams, opts ...option.RequestOption) (res *NotificationSendResult, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single notification by ID.
//
// Only notifications addressed to the current user are visible; another user's
// notification is reported as not found. Dismissed notifications remain
// retrievable.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingNotificationService) Get(ctx context.Context, id string, query MessagingNotificationGetParams, opts ...option.RequestOption) (res *Notification, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/notifications/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists the notifications addressed to the current user, newest first.
//
// The feed is personal and scoped to the account being acted in, so it never
// includes another user's notifications. Callers with no user membership in that
// account, such as an API key, get an empty list rather than an error.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingNotificationService) List(ctx context.Context, query MessagingNotificationListParams, opts ...option.RequestOption) (res *ListNotification, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns the current user's unread tallies for the account they are acting in,
// for driving a notification badge.
//
// The total also counts account announcements the user has not seen, so it can be
// higher than the notification count alone.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingNotificationService) GetUnreadCount(ctx context.Context, opts ...option.RequestOption) (res *NotificationUnreadCount, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/notifications/unread-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the caller's unread totals broken down by account, covering every
// account they belong to and not just the one they are acting in.
//
// Use it to show a user that activity is waiting for them elsewhere before they
// switch accounts. Each tally counts unseen notifications and unseen account
// announcements together.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingNotificationService) GetUnreadSummary(ctx context.Context, opts ...option.RequestOption) (res *NotificationUnreadSummary, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/notifications/unread-summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListNotification struct {
	// Resources in this page.
	Data []Notification `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListNotificationObject `json:"object" api:"required"`
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
func (r ListNotification) RawJSON() string { return r.JSON.raw }
func (r *ListNotification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListNotificationObject string

const (
	ListNotificationObjectList ListNotificationObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListNotificationUnreadSummaryAccount struct {
	// Resources in this page.
	Data []NotificationUnreadSummaryAccount `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListNotificationUnreadSummaryAccountObject `json:"object" api:"required"`
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
func (r ListNotificationUnreadSummaryAccount) RawJSON() string { return r.JSON.raw }
func (r *ListNotificationUnreadSummaryAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListNotificationUnreadSummaryAccountObject string

const (
	ListNotificationUnreadSummaryAccountObjectList ListNotificationUnreadSummaryAccountObject = "list"
)

// An in-app notification addressed to a single user, shown in their notification
// (bell) feed.
//
// A notification belongs to one user in one account, so the feed you read is
// always that of the authenticated caller in the account they are acting in.
// Announcements broadcast to a whole account are a separate resource.
type Notification struct {
	// Notification ID.
	ID string `json:"id" api:"required"`
	// Supporting detail shown beneath the title, such as a preview of the message that
	// triggered the notification.
	Body string `json:"body" api:"required"`
	// The kind of event this notification represents.
	//
	// The set is open-ended and may grow over time. Common first-party categories are:
	//
	//   - `chat.message`: a new message in a conversation.
	//   - `chat.mention`: a direct @mention, delivered even when the conversation is
	//     muted.
	//   - `chat.added`: the user was added to a conversation.
	//   - `order.updated`: an order the user is involved with changed.
	//   - `agent.run_completed`: an agent run the user triggered finished.
	//   - `agent.alert`: an agent raised an alert during a run.
	//   - `system.broadcast`: a targeted system message.
	//   - `customer.registered`: a buyer completed registration on your customer portal.
	//
	// Any of "chat.message", "chat.mention", "chat.added", "order.updated",
	// "agent.run_completed", "agent.alert", "system.broadcast", "customer.registered".
	Category NotificationCategory `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the notification was dismissed.
	DismissedAt time.Time `json:"dismissed_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "notification".
	Object NotificationObject `json:"object" api:"required"`
	// How prominently the notification should be surfaced, from `low` through
	// `urgent`.
	//
	// Any of "low", "normal", "high", "urgent".
	Priority NotificationPriority `json:"priority" api:"required"`
	// When the notification was explicitly opened.
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// When the notification was first surfaced to the user.
	SeenAt time.Time `json:"seen_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Sender Actor `json:"sender" api:"required"`
	// Where the notification is in its lifecycle.
	//
	// - `unseen`: delivered but not yet surfaced to the user.
	// - `seen`: surfaced in the feed but not yet opened.
	// - `read`: explicitly opened by the user.
	// - `dismissed`: removed from the active feed.
	//
	// The status is derived from the seen, read, and dismissed timestamps, and only
	// ever moves forward — a notification can never become unseen again.
	//
	// Any of "unseen", "seen", "read", "dismissed".
	Status NotificationStatus `json:"status" api:"required"`
	// Short headline shown in the feed.
	Title string `json:"title" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Body        respjson.Field
		Category    respjson.Field
		CreatedAt   respjson.Field
		DismissedAt respjson.Field
		Object      respjson.Field
		Priority    respjson.Field
		ReadAt      respjson.Field
		Resource    respjson.Field
		SeenAt      respjson.Field
		Sender      respjson.Field
		Status      respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Notification) RawJSON() string { return r.JSON.raw }
func (r *Notification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of event this notification represents.
//
// The set is open-ended and may grow over time. Common first-party categories are:
//
//   - `chat.message`: a new message in a conversation.
//   - `chat.mention`: a direct @mention, delivered even when the conversation is
//     muted.
//   - `chat.added`: the user was added to a conversation.
//   - `order.updated`: an order the user is involved with changed.
//   - `agent.run_completed`: an agent run the user triggered finished.
//   - `agent.alert`: an agent raised an alert during a run.
//   - `system.broadcast`: a targeted system message.
//   - `customer.registered`: a buyer completed registration on your customer portal.
type NotificationCategory string

const (
	NotificationCategoryChatMessage        NotificationCategory = "chat.message"
	NotificationCategoryChatMention        NotificationCategory = "chat.mention"
	NotificationCategoryChatAdded          NotificationCategory = "chat.added"
	NotificationCategoryOrderUpdated       NotificationCategory = "order.updated"
	NotificationCategoryAgentRunCompleted  NotificationCategory = "agent.run_completed"
	NotificationCategoryAgentAlert         NotificationCategory = "agent.alert"
	NotificationCategorySystemBroadcast    NotificationCategory = "system.broadcast"
	NotificationCategoryCustomerRegistered NotificationCategory = "customer.registered"
)

// Resource type identifier.
type NotificationObject string

const (
	NotificationObjectNotification NotificationObject = "notification"
)

// How prominently the notification should be surfaced, from `low` through
// `urgent`.
type NotificationPriority string

const (
	NotificationPriorityLow    NotificationPriority = "low"
	NotificationPriorityNormal NotificationPriority = "normal"
	NotificationPriorityHigh   NotificationPriority = "high"
	NotificationPriorityUrgent NotificationPriority = "urgent"
)

// Where the notification is in its lifecycle.
//
// - `unseen`: delivered but not yet surfaced to the user.
// - `seen`: surfaced in the feed but not yet opened.
// - `read`: explicitly opened by the user.
// - `dismissed`: removed from the active feed.
//
// The status is derived from the seen, read, and dismissed timestamps, and only
// ever moves forward — a notification can never become unseen again.
type NotificationStatus string

const (
	NotificationStatusUnseen    NotificationStatus = "unseen"
	NotificationStatusSeen      NotificationStatus = "seen"
	NotificationStatusRead      NotificationStatus = "read"
	NotificationStatusDismissed NotificationStatus = "dismissed"
)

// The acknowledgement returned when a notification is accepted for delivery.
type NotificationSendResult struct {
	// Number of deliveries accepted for the notification.
	//
	// An account broadcast is stored once as a single announcement that serves
	// everyone in the account, so it reports `1` rather than a per-user count.
	// Acceptance is not delivery: recipients who cannot be resolved are skipped when
	// the notification is fanned out.
	Enqueued int64 `json:"enqueued" api:"required"`
	// Resource type identifier.
	//
	// Any of "notification_send_result".
	Object NotificationSendResultObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enqueued    respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationSendResult) RawJSON() string { return r.JSON.raw }
func (r *NotificationSendResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type NotificationSendResultObject string

const (
	NotificationSendResultObjectNotificationSendResult NotificationSendResultObject = "notification_send_result"
)

// Who a notification is aimed at.
//
// The properties ID, Type are required.
type NotificationTargetInputParam struct {
	// The id of the recipient, matching `type`: an account user id, or an account id.
	//
	// An account target must be the account you are currently acting in — you cannot
	// broadcast into another account.
	ID string `json:"id" api:"required"`
	// The kind of recipient being addressed.
	//
	//   - `account_user`: one member of the account, who receives a personal
	//     notification in their feed.
	//   - `account`: every member of the account, who all receive a single shared
	//     announcement.
	//
	// Any of "account_user", "account".
	Type NotificationTargetInputType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r NotificationTargetInputParam) MarshalJSON() (data []byte, err error) {
	type shadow NotificationTargetInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NotificationTargetInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of recipient being addressed.
//
//   - `account_user`: one member of the account, who receives a personal
//     notification in their feed.
//   - `account`: every member of the account, who all receive a single shared
//     announcement.
type NotificationTargetInputType string

const (
	NotificationTargetInputTypeAccountUser NotificationTargetInputType = "account_user"
	NotificationTargetInputTypeAccount     NotificationTargetInputType = "account"
)

// The caller's unread tallies in one account, used to drive the notification bell
// badge.
type NotificationUnreadCount struct {
	// Number of conversations with unread messages.
	//
	// Always `0` today — conversation unread counts are not yet folded into the bell.
	Conversations int64 `json:"conversations" api:"required"`
	// Number of the caller's notifications that have not been seen yet.
	//
	// Dismissed notifications are never counted, and marking all notifications seen
	// drops this to zero.
	Notifications int64 `json:"notifications" api:"required"`
	// Resource type identifier.
	//
	// Any of "notification_unread_count".
	Object NotificationUnreadCountObject `json:"object" api:"required"`
	// Combined unread total for the bell badge.
	//
	// This is the unseen notification count plus any account announcements the caller
	// has not seen, so it can exceed `notifications`. Announcements are cleared
	// individually rather than by marking all notifications seen.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversations respjson.Field
		Notifications respjson.Field
		Object        respjson.Field
		Total         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationUnreadCount) RawJSON() string { return r.JSON.raw }
func (r *NotificationUnreadCount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type NotificationUnreadCountObject string

const (
	NotificationUnreadCountObjectNotificationUnreadCount NotificationUnreadCountObject = "notification_unread_count"
)

// The caller's unread totals across every account they belong to, used to show
// unread activity waiting in accounts they are not currently working in.
type NotificationUnreadSummary struct {
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Accounts ListNotificationUnreadSummaryAccount `json:"accounts" api:"required"`
	// Resource type identifier.
	//
	// Any of "notification_unread_summary".
	Object NotificationUnreadSummaryObject `json:"object" api:"required"`
	// Combined unread total across all of the caller's accounts.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		Object      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationUnreadSummary) RawJSON() string { return r.JSON.raw }
func (r *NotificationUnreadSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type NotificationUnreadSummaryObject string

const (
	NotificationUnreadSummaryObjectNotificationUnreadSummary NotificationUnreadSummaryObject = "notification_unread_summary"
)

// One account's unread tally within the caller's cross-account summary.
type NotificationUnreadSummaryAccount struct {
	// Entity is a polymorphic reference to any resource in the system.
	Account Entity `json:"account" api:"required"`
	// Resource type identifier.
	//
	// Any of "notification_unread_summary_account".
	Object NotificationUnreadSummaryAccountObject `json:"object" api:"required"`
	// Number of unseen notifications and account announcements the caller has in this
	// account.
	Unread int64 `json:"unread" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Object      respjson.Field
		Unread      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationUnreadSummaryAccount) RawJSON() string { return r.JSON.raw }
func (r *NotificationUnreadSummaryAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type NotificationUnreadSummaryAccountObject string

const (
	NotificationUnreadSummaryAccountObjectNotificationUnreadSummaryAccount NotificationUnreadSummaryAccountObject = "notification_unread_summary_account"
)

// Request to send an in-app notification.
//
// The target decides whether the notification goes to one member of the account or
// to everyone in it.
//
// The properties Category, Target, Title are required.
type SendNotificationRequestParam struct {
	// The kind of event the notification represents, such as `order.updated`.
	//
	// Categories are how clients group and filter the feed, so reuse an existing one
	// where it fits.
	//
	// Any of "chat.message", "chat.mention", "chat.added", "order.updated",
	// "agent.run_completed", "agent.alert", "system.broadcast", "customer.registered".
	Category SendNotificationRequestCategory `json:"category,omitzero" api:"required"`
	// Who a notification is aimed at.
	Target NotificationTargetInputParam `json:"target,omitzero" api:"required"`
	// Short headline shown in the recipient's feed.
	Title string `json:"title" api:"required"`
	// Supporting detail shown beneath the title.
	Body param.Opt[string] `json:"body,omitzero"`
	// ID of the resource the notification should link to.
	LinkResourceID param.Opt[string] `json:"link_resource_id,omitzero"`
	// Type of the resource the notification should link to, such as `sales_order`.
	//
	// Set it together with `link_resource_id` to point the notification at something
	// the recipient can open; supplying only one of the two produces a notification
	// with no link.
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
	LinkResourceType SendNotificationRequestLinkResourceType `json:"link_resource_type,omitzero"`
	// How prominently the notification should be surfaced, from `low` through
	// `urgent`.
	//
	// Any of "low", "normal", "high", "urgent".
	Priority SendNotificationRequestPriority `json:"priority,omitzero"`
	paramObj
}

func (r SendNotificationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SendNotificationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SendNotificationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of event the notification represents, such as `order.updated`.
//
// Categories are how clients group and filter the feed, so reuse an existing one
// where it fits.
type SendNotificationRequestCategory string

const (
	SendNotificationRequestCategoryChatMessage        SendNotificationRequestCategory = "chat.message"
	SendNotificationRequestCategoryChatMention        SendNotificationRequestCategory = "chat.mention"
	SendNotificationRequestCategoryChatAdded          SendNotificationRequestCategory = "chat.added"
	SendNotificationRequestCategoryOrderUpdated       SendNotificationRequestCategory = "order.updated"
	SendNotificationRequestCategoryAgentRunCompleted  SendNotificationRequestCategory = "agent.run_completed"
	SendNotificationRequestCategoryAgentAlert         SendNotificationRequestCategory = "agent.alert"
	SendNotificationRequestCategorySystemBroadcast    SendNotificationRequestCategory = "system.broadcast"
	SendNotificationRequestCategoryCustomerRegistered SendNotificationRequestCategory = "customer.registered"
)

// Type of the resource the notification should link to, such as `sales_order`.
//
// Set it together with `link_resource_id` to point the notification at something
// the recipient can open; supplying only one of the two produces a notification
// with no link.
type SendNotificationRequestLinkResourceType string

const (
	SendNotificationRequestLinkResourceTypeAccount                              SendNotificationRequestLinkResourceType = "account"
	SendNotificationRequestLinkResourceTypeActor                                SendNotificationRequestLinkResourceType = "actor"
	SendNotificationRequestLinkResourceTypeEntity                               SendNotificationRequestLinkResourceType = "entity"
	SendNotificationRequestLinkResourceTypeRecord                               SendNotificationRequestLinkResourceType = "record"
	SendNotificationRequestLinkResourceTypeFreight                              SendNotificationRequestLinkResourceType = "freight"
	SendNotificationRequestLinkResourceTypeCommitment                           SendNotificationRequestLinkResourceType = "commitment"
	SendNotificationRequestLinkResourceTypeSalesOrderTotals                     SendNotificationRequestLinkResourceType = "sales_order_totals"
	SendNotificationRequestLinkResourceTypeSalesOrderStageTotal                 SendNotificationRequestLinkResourceType = "sales_order_stage_total"
	SendNotificationRequestLinkResourceTypeSalesOrderRelated                    SendNotificationRequestLinkResourceType = "sales_order_related"
	SendNotificationRequestLinkResourceTypeOrderContact                         SendNotificationRequestLinkResourceType = "order_contact"
	SendNotificationRequestLinkResourceTypeUser                                 SendNotificationRequestLinkResourceType = "user"
	SendNotificationRequestLinkResourceTypeAddress                              SendNotificationRequestLinkResourceType = "address"
	SendNotificationRequestLinkResourceTypeAPIKey                               SendNotificationRequestLinkResourceType = "api_key"
	SendNotificationRequestLinkResourceTypeCreatedAPIKey                        SendNotificationRequestLinkResourceType = "created_api_key"
	SendNotificationRequestLinkResourceTypeRefreshToken                         SendNotificationRequestLinkResourceType = "refresh_token"
	SendNotificationRequestLinkResourceTypeList                                 SendNotificationRequestLinkResourceType = "list"
	SendNotificationRequestLinkResourceTypeSandbox                              SendNotificationRequestLinkResourceType = "sandbox"
	SendNotificationRequestLinkResourceTypeRegistrationSession                  SendNotificationRequestLinkResourceType = "registration_session"
	SendNotificationRequestLinkResourceTypePricingPlan                          SendNotificationRequestLinkResourceType = "pricing_plan"
	SendNotificationRequestLinkResourceTypeAccountPlan                          SendNotificationRequestLinkResourceType = "account_plan"
	SendNotificationRequestLinkResourceTypePlanChange                           SendNotificationRequestLinkResourceType = "plan_change"
	SendNotificationRequestLinkResourceTypeEnterpriseInquiry                    SendNotificationRequestLinkResourceType = "enterprise_inquiry"
	SendNotificationRequestLinkResourceTypeRequestLog                           SendNotificationRequestLinkResourceType = "request_log"
	SendNotificationRequestLinkResourceTypeAuditEvent                           SendNotificationRequestLinkResourceType = "audit_event"
	SendNotificationRequestLinkResourceTypeAuditFieldChange                     SendNotificationRequestLinkResourceType = "audit_field_change"
	SendNotificationRequestLinkResourceTypeRole                                 SendNotificationRequestLinkResourceType = "role"
	SendNotificationRequestLinkResourceTypeUnit                                 SendNotificationRequestLinkResourceType = "unit"
	SendNotificationRequestLinkResourceTypeAccountAffiliation                   SendNotificationRequestLinkResourceType = "account_affiliation"
	SendNotificationRequestLinkResourceTypeAgentDefinition                      SendNotificationRequestLinkResourceType = "agent_definition"
	SendNotificationRequestLinkResourceTypeAvailableTool                        SendNotificationRequestLinkResourceType = "available_tool"
	SendNotificationRequestLinkResourceTypeAgentDefinitionTool                  SendNotificationRequestLinkResourceType = "agent_definition_tool"
	SendNotificationRequestLinkResourceTypeAgentAccountStatus                   SendNotificationRequestLinkResourceType = "agent_account_status"
	SendNotificationRequestLinkResourceTypeAgentRun                             SendNotificationRequestLinkResourceType = "agent_run"
	SendNotificationRequestLinkResourceTypeAgentAction                          SendNotificationRequestLinkResourceType = "agent_action"
	SendNotificationRequestLinkResourceTypeAgentRunStep                         SendNotificationRequestLinkResourceType = "agent_run_step"
	SendNotificationRequestLinkResourceTypeAgentTokenUsage                      SendNotificationRequestLinkResourceType = "agent_token_usage"
	SendNotificationRequestLinkResourceTypeAgentMemory                          SendNotificationRequestLinkResourceType = "agent_memory"
	SendNotificationRequestLinkResourceTypeNotification                         SendNotificationRequestLinkResourceType = "notification"
	SendNotificationRequestLinkResourceTypeNotificationUnreadCount              SendNotificationRequestLinkResourceType = "notification_unread_count"
	SendNotificationRequestLinkResourceTypeNotificationSendResult               SendNotificationRequestLinkResourceType = "notification_send_result"
	SendNotificationRequestLinkResourceTypeNotificationUnreadSummary            SendNotificationRequestLinkResourceType = "notification_unread_summary"
	SendNotificationRequestLinkResourceTypeAnnouncement                         SendNotificationRequestLinkResourceType = "announcement"
	SendNotificationRequestLinkResourceTypeConversation                         SendNotificationRequestLinkResourceType = "conversation"
	SendNotificationRequestLinkResourceTypeSupportCase                          SendNotificationRequestLinkResourceType = "support_case"
	SendNotificationRequestLinkResourceTypeConversationParticipant              SendNotificationRequestLinkResourceType = "conversation_participant"
	SendNotificationRequestLinkResourceTypeReadCursor                           SendNotificationRequestLinkResourceType = "read_cursor"
	SendNotificationRequestLinkResourceTypeChatMessage                          SendNotificationRequestLinkResourceType = "chat_message"
	SendNotificationRequestLinkResourceTypeNotificationUnreadSummaryAccount     SendNotificationRequestLinkResourceType = "notification_unread_summary_account"
	SendNotificationRequestLinkResourceTypeMessagingBlock                       SendNotificationRequestLinkResourceType = "messaging_block"
	SendNotificationRequestLinkResourceTypeNotificationPreference               SendNotificationRequestLinkResourceType = "notification_preference"
	SendNotificationRequestLinkResourceTypeMessageAttachment                    SendNotificationRequestLinkResourceType = "message_attachment"
	SendNotificationRequestLinkResourceTypeAttachmentUploadTarget               SendNotificationRequestLinkResourceType = "attachment_upload_target"
	SendNotificationRequestLinkResourceTypeScheduledMessage                     SendNotificationRequestLinkResourceType = "scheduled_message"
	SendNotificationRequestLinkResourceTypeMessagingContact                     SendNotificationRequestLinkResourceType = "messaging_contact"
	SendNotificationRequestLinkResourceTypeMessageReport                        SendNotificationRequestLinkResourceType = "message_report"
	SendNotificationRequestLinkResourceTypeToolGroup                            SendNotificationRequestLinkResourceType = "tool_group"
	SendNotificationRequestLinkResourceTypeModel                                SendNotificationRequestLinkResourceType = "model"
	SendNotificationRequestLinkResourceTypePaymentTerm                          SendNotificationRequestLinkResourceType = "payment_term"
	SendNotificationRequestLinkResourceTypeShippingTerm                         SendNotificationRequestLinkResourceType = "shipping_term"
	SendNotificationRequestLinkResourceTypeQuantity                             SendNotificationRequestLinkResourceType = "quantity"
	SendNotificationRequestLinkResourceTypeAccountGroup                         SendNotificationRequestLinkResourceType = "account_group"
	SendNotificationRequestLinkResourceTypeSupportRoute                         SendNotificationRequestLinkResourceType = "support_route"
	SendNotificationRequestLinkResourceTypeSupportAvailability                  SendNotificationRequestLinkResourceType = "support_availability"
	SendNotificationRequestLinkResourceTypeAccountStatus                        SendNotificationRequestLinkResourceType = "account_status"
	SendNotificationRequestLinkResourceTypeGeolocation                          SendNotificationRequestLinkResourceType = "geolocation"
	SendNotificationRequestLinkResourceTypeAccountUser                          SendNotificationRequestLinkResourceType = "account_user"
	SendNotificationRequestLinkResourceTypeDepartment                           SendNotificationRequestLinkResourceType = "department"
	SendNotificationRequestLinkResourceTypeAccountIntegration                   SendNotificationRequestLinkResourceType = "account_integration"
	SendNotificationRequestLinkResourceTypeAccountPrice                         SendNotificationRequestLinkResourceType = "account_price"
	SendNotificationRequestLinkResourceTypeProductLine                          SendNotificationRequestLinkResourceType = "product_line"
	SendNotificationRequestLinkResourceTypeItemCategory                         SendNotificationRequestLinkResourceType = "item_category"
	SendNotificationRequestLinkResourceTypeAttribute                            SendNotificationRequestLinkResourceType = "attribute"
	SendNotificationRequestLinkResourceTypeRate                                 SendNotificationRequestLinkResourceType = "rate"
	SendNotificationRequestLinkResourceTypeAccountGroupProductLineAccess        SendNotificationRequestLinkResourceType = "account_group_product_line_access"
	SendNotificationRequestLinkResourceTypeSalesTarget                          SendNotificationRequestLinkResourceType = "sales_target"
	SendNotificationRequestLinkResourceTypeAdjustmentType                       SendNotificationRequestLinkResourceType = "adjustment_type"
	SendNotificationRequestLinkResourceTypeAccountBranding                      SendNotificationRequestLinkResourceType = "account_branding"
	SendNotificationRequestLinkResourceTypeAccountPortal                        SendNotificationRequestLinkResourceType = "account_portal"
	SendNotificationRequestLinkResourceTypeAccountLogoURL                       SendNotificationRequestLinkResourceType = "account_logo_url"
	SendNotificationRequestLinkResourceTypeAccountFaviconURL                    SendNotificationRequestLinkResourceType = "account_favicon_url"
	SendNotificationRequestLinkResourceTypePublicAccount                        SendNotificationRequestLinkResourceType = "public_account"
	SendNotificationRequestLinkResourceTypeProperty                             SendNotificationRequestLinkResourceType = "property"
	SendNotificationRequestLinkResourceTypeCarrier                              SendNotificationRequestLinkResourceType = "carrier"
	SendNotificationRequestLinkResourceTypeServiceLevel                         SendNotificationRequestLinkResourceType = "service_level"
	SendNotificationRequestLinkResourceTypeItem                                 SendNotificationRequestLinkResourceType = "item"
	SendNotificationRequestLinkResourceTypeItemLotDefault                       SendNotificationRequestLinkResourceType = "item_lot_default"
	SendNotificationRequestLinkResourceTypeItemInventory                        SendNotificationRequestLinkResourceType = "item_inventory"
	SendNotificationRequestLinkResourceTypeProduct                              SendNotificationRequestLinkResourceType = "product"
	SendNotificationRequestLinkResourceTypeBatch                                SendNotificationRequestLinkResourceType = "batch"
	SendNotificationRequestLinkResourceTypeBatchFlowNode                        SendNotificationRequestLinkResourceType = "batch_flow_node"
	SendNotificationRequestLinkResourceTypeScanningConsumption                  SendNotificationRequestLinkResourceType = "scanning_consumption"
	SendNotificationRequestLinkResourceTypeOpenBatchSummary                     SendNotificationRequestLinkResourceType = "open_batch_summary"
	SendNotificationRequestLinkResourceTypeScanningProductionStepInfo           SendNotificationRequestLinkResourceType = "scanning_production_step_info"
	SendNotificationRequestLinkResourceTypeScanningStation                      SendNotificationRequestLinkResourceType = "scanning_station"
	SendNotificationRequestLinkResourceTypeProductionStep                       SendNotificationRequestLinkResourceType = "production_step"
	SendNotificationRequestLinkResourceTypeProductionRun                        SendNotificationRequestLinkResourceType = "production_run"
	SendNotificationRequestLinkResourceTypeMachine                              SendNotificationRequestLinkResourceType = "machine"
	SendNotificationRequestLinkResourceTypeMachineStatus                        SendNotificationRequestLinkResourceType = "machine_status"
	SendNotificationRequestLinkResourceTypeMachineDowntimeEvent                 SendNotificationRequestLinkResourceType = "machine_downtime_event"
	SendNotificationRequestLinkResourceTypeDemandOverride                       SendNotificationRequestLinkResourceType = "demand_override"
	SendNotificationRequestLinkResourceTypeDemandOverrideType                   SendNotificationRequestLinkResourceType = "demand_override_type"
	SendNotificationRequestLinkResourceTypeMachineDowntimeReason                SendNotificationRequestLinkResourceType = "machine_downtime_reason"
	SendNotificationRequestLinkResourceTypeProductionSchedulePreview            SendNotificationRequestLinkResourceType = "production_schedule_preview"
	SendNotificationRequestLinkResourceTypeProductionScheduleRegeneratePreview  SendNotificationRequestLinkResourceType = "production_schedule_regenerate_preview"
	SendNotificationRequestLinkResourceTypeProductionSchedule                   SendNotificationRequestLinkResourceType = "production_schedule"
	SendNotificationRequestLinkResourceTypeProductionScheduleLine               SendNotificationRequestLinkResourceType = "production_schedule_line"
	SendNotificationRequestLinkResourceTypeProductionScheduleDeviation          SendNotificationRequestLinkResourceType = "production_schedule_deviation"
	SendNotificationRequestLinkResourceTypeProductionScheduleDerivedLine        SendNotificationRequestLinkResourceType = "production_schedule_derived_line"
	SendNotificationRequestLinkResourceTypeProductionScheduleSettings           SendNotificationRequestLinkResourceType = "production_schedule_settings"
	SendNotificationRequestLinkResourceTypeProductionScheduleResourceSetting    SendNotificationRequestLinkResourceType = "production_schedule_resource_setting"
	SendNotificationRequestLinkResourceTypeProductionScheduleItemSetting        SendNotificationRequestLinkResourceType = "production_schedule_item_setting"
	SendNotificationRequestLinkResourceTypeFulfillmentRecommendation            SendNotificationRequestLinkResourceType = "fulfillment_recommendation"
	SendNotificationRequestLinkResourceTypeAnalyzeDeliveryPerformanceResponse   SendNotificationRequestLinkResourceType = "analyze_delivery_performance_response"
	SendNotificationRequestLinkResourceTypeDeliveryPerformance                  SendNotificationRequestLinkResourceType = "delivery_performance"
	SendNotificationRequestLinkResourceTypeDeliveryBacklogBucket                SendNotificationRequestLinkResourceType = "delivery_backlog_bucket"
	SendNotificationRequestLinkResourceTypeDeliveryLatenessBucket               SendNotificationRequestLinkResourceType = "delivery_lateness_bucket"
	SendNotificationRequestLinkResourceTypeDeliveryBreakdown                    SendNotificationRequestLinkResourceType = "delivery_breakdown"
	SendNotificationRequestLinkResourceTypeAnalyzeSalesBreakdownResponse        SendNotificationRequestLinkResourceType = "analyze_sales_breakdown_response"
	SendNotificationRequestLinkResourceTypeSalesTotals                          SendNotificationRequestLinkResourceType = "sales_totals"
	SendNotificationRequestLinkResourceTypeSalesBreakdown                       SendNotificationRequestLinkResourceType = "sales_breakdown"
	SendNotificationRequestLinkResourceTypeScheduleOrderCoverage                SendNotificationRequestLinkResourceType = "schedule_order_coverage"
	SendNotificationRequestLinkResourceTypeScheduleOrderCoverageLine            SendNotificationRequestLinkResourceType = "schedule_order_coverage_line"
	SendNotificationRequestLinkResourceTypeScheduleDeviationType                SendNotificationRequestLinkResourceType = "schedule_deviation_type"
	SendNotificationRequestLinkResourceTypeScheduleAtRiskOrder                  SendNotificationRequestLinkResourceType = "schedule_at_risk_order"
	SendNotificationRequestLinkResourceTypeProductionScheduleFinishedPolicy     SendNotificationRequestLinkResourceType = "production_schedule_finished_policy"
	SendNotificationRequestLinkResourceTypeProductionScheduleFinishingLine      SendNotificationRequestLinkResourceType = "production_schedule_finishing_line"
	SendNotificationRequestLinkResourceTypeProductionScheduleWeekRelease        SendNotificationRequestLinkResourceType = "production_schedule_week_release"
	SendNotificationRequestLinkResourceTypeProductionScheduleWeekReleasePreview SendNotificationRequestLinkResourceType = "production_schedule_week_release_preview"
	SendNotificationRequestLinkResourceTypeProductionScheduleItemPolicy         SendNotificationRequestLinkResourceType = "production_schedule_item_policy"
	SendNotificationRequestLinkResourceTypeChildAccount                         SendNotificationRequestLinkResourceType = "child_account"
	SendNotificationRequestLinkResourceTypeUnitGroup                            SendNotificationRequestLinkResourceType = "unit_group"
	SendNotificationRequestLinkResourceTypeUnitGroupUnit                        SendNotificationRequestLinkResourceType = "unit_group_unit"
	SendNotificationRequestLinkResourceTypeConsumption                          SendNotificationRequestLinkResourceType = "consumption"
	SendNotificationRequestLinkResourceTypeCustomerProductLineAccess            SendNotificationRequestLinkResourceType = "customer_product_line_access"
	SendNotificationRequestLinkResourceTypeCustomer                             SendNotificationRequestLinkResourceType = "customer"
	SendNotificationRequestLinkResourceTypeFrequentlyOrderedProduct             SendNotificationRequestLinkResourceType = "frequently_ordered_product"
	SendNotificationRequestLinkResourceTypePriority                             SendNotificationRequestLinkResourceType = "priority"
	SendNotificationRequestLinkResourceTypeDelivery                             SendNotificationRequestLinkResourceType = "delivery"
	SendNotificationRequestLinkResourceTypeDeliveryLine                         SendNotificationRequestLinkResourceType = "delivery_line"
	SendNotificationRequestLinkResourceTypeSalesOrder                           SendNotificationRequestLinkResourceType = "sales_order"
	SendNotificationRequestLinkResourceTypeLocation                             SendNotificationRequestLinkResourceType = "location"
	SendNotificationRequestLinkResourceTypeLocationType                         SendNotificationRequestLinkResourceType = "location_type"
	SendNotificationRequestLinkResourceTypeLot                                  SendNotificationRequestLinkResourceType = "lot"
	SendNotificationRequestLinkResourceTypeEmailLog                             SendNotificationRequestLinkResourceType = "email_log"
	SendNotificationRequestLinkResourceTypeEmailDomain                          SendNotificationRequestLinkResourceType = "email_domain"
	SendNotificationRequestLinkResourceTypeEmailInbox                           SendNotificationRequestLinkResourceType = "email_inbox"
	SendNotificationRequestLinkResourceTypePortalDomain                         SendNotificationRequestLinkResourceType = "portal_domain"
	SendNotificationRequestLinkResourceTypeDNSRecord                            SendNotificationRequestLinkResourceType = "dns_record"
	SendNotificationRequestLinkResourceTypeInventoryChangeLog                   SendNotificationRequestLinkResourceType = "inventory_change_log"
	SendNotificationRequestLinkResourceTypeInvoice                              SendNotificationRequestLinkResourceType = "invoice"
	SendNotificationRequestLinkResourceTypeInvoiceSummary                       SendNotificationRequestLinkResourceType = "invoice_summary"
	SendNotificationRequestLinkResourceTypeInvoiceLine                          SendNotificationRequestLinkResourceType = "invoice_line"
	SendNotificationRequestLinkResourceTypeInvoiceAllocation                    SendNotificationRequestLinkResourceType = "invoice_allocation"
	SendNotificationRequestLinkResourceTypeInvoiceForPayment                    SendNotificationRequestLinkResourceType = "invoice_for_payment"
	SendNotificationRequestLinkResourceTypeShipment                             SendNotificationRequestLinkResourceType = "shipment"
	SendNotificationRequestLinkResourceTypeShipmentSummary                      SendNotificationRequestLinkResourceType = "shipment_summary"
	SendNotificationRequestLinkResourceTypeShipmentLine                         SendNotificationRequestLinkResourceType = "shipment_line"
	SendNotificationRequestLinkResourceTypeShippingCase                         SendNotificationRequestLinkResourceType = "shipping_case"
	SendNotificationRequestLinkResourceTypeShippingCaseLabelURL                 SendNotificationRequestLinkResourceType = "shipping_case_label_url"
	SendNotificationRequestLinkResourceTypeSettlement                           SendNotificationRequestLinkResourceType = "settlement"
	SendNotificationRequestLinkResourceTypeSettlementSummary                    SendNotificationRequestLinkResourceType = "settlement_summary"
	SendNotificationRequestLinkResourceTypeRolePermission                       SendNotificationRequestLinkResourceType = "role_permission"
	SendNotificationRequestLinkResourceTypeRegistrationFlow                     SendNotificationRequestLinkResourceType = "registration_flow"
	SendNotificationRequestLinkResourceTypeRegistrationFlowOption               SendNotificationRequestLinkResourceType = "registration_flow_option"
	SendNotificationRequestLinkResourceTypeTransaction                          SendNotificationRequestLinkResourceType = "transaction"
	SendNotificationRequestLinkResourceTypeTransactionSummary                   SendNotificationRequestLinkResourceType = "transaction_summary"
	SendNotificationRequestLinkResourceTypeTransactionMethod                    SendNotificationRequestLinkResourceType = "transaction_method"
	SendNotificationRequestLinkResourceTypeTransactionType                      SendNotificationRequestLinkResourceType = "transaction_type"
	SendNotificationRequestLinkResourceTypeTransactionAllocation                SendNotificationRequestLinkResourceType = "transaction_allocation"
	SendNotificationRequestLinkResourceTypeUsageItem                            SendNotificationRequestLinkResourceType = "usage_item"
	SendNotificationRequestLinkResourceTypeAccountUsageResponse                 SendNotificationRequestLinkResourceType = "account_usage_response"
	SendNotificationRequestLinkResourceTypeSubscriptionInfo                     SendNotificationRequestLinkResourceType = "subscription_info"
	SendNotificationRequestLinkResourceTypeBillingPortalSessionResponse         SendNotificationRequestLinkResourceType = "billing_portal_session_response"
	SendNotificationRequestLinkResourceTypeSwitchPlanResponse                   SendNotificationRequestLinkResourceType = "switch_plan_response"
	SendNotificationRequestLinkResourceTypeEnsureBillingCustomerResponse        SendNotificationRequestLinkResourceType = "ensure_billing_customer_response"
	SendNotificationRequestLinkResourceTypeSpendingCapResponse                  SendNotificationRequestLinkResourceType = "spending_cap_response"
	SendNotificationRequestLinkResourceTypeAgentSpendInfo                       SendNotificationRequestLinkResourceType = "agent_spend_info"
	SendNotificationRequestLinkResourceTypeWebhookResponse                      SendNotificationRequestLinkResourceType = "webhook_response"
	SendNotificationRequestLinkResourceTypeAddressSuggestion                    SendNotificationRequestLinkResourceType = "address_suggestion"
	SendNotificationRequestLinkResourceTypeAddressComponents                    SendNotificationRequestLinkResourceType = "address_components"
	SendNotificationRequestLinkResourceTypeAddressDetailsResult                 SendNotificationRequestLinkResourceType = "address_details_result"
	SendNotificationRequestLinkResourceTypeValidatedAddress                     SendNotificationRequestLinkResourceType = "validated_address"
	SendNotificationRequestLinkResourceTypePlanLimit                            SendNotificationRequestLinkResourceType = "plan_limit"
	SendNotificationRequestLinkResourceTypePlanChangeProration                  SendNotificationRequestLinkResourceType = "plan_change_proration"
	SendNotificationRequestLinkResourceTypePlanChangeLineItem                   SendNotificationRequestLinkResourceType = "plan_change_line_item"
	SendNotificationRequestLinkResourceTypeSetupBillingResponse                 SendNotificationRequestLinkResourceType = "setup_billing_response"
	SendNotificationRequestLinkResourceTypeConfirmPaymentResponse               SendNotificationRequestLinkResourceType = "confirm_payment_response"
	SendNotificationRequestLinkResourceTypeOAuthResponse                        SendNotificationRequestLinkResourceType = "oauth_response"
	SendNotificationRequestLinkResourceTypeOAuthStatusResponse                  SendNotificationRequestLinkResourceType = "oauth_status_response"
	SendNotificationRequestLinkResourceTypeStripePublishableKey                 SendNotificationRequestLinkResourceType = "stripe_publishable_key"
	SendNotificationRequestLinkResourceTypeStripeStatus                         SendNotificationRequestLinkResourceType = "stripe_status"
	SendNotificationRequestLinkResourceTypeHealthcheck                          SendNotificationRequestLinkResourceType = "healthcheck"
	SendNotificationRequestLinkResourceTypeAgentDefinitionConfig                SendNotificationRequestLinkResourceType = "agent_definition_config"
	SendNotificationRequestLinkResourceTypeTriggerConfig                        SendNotificationRequestLinkResourceType = "trigger_config"
	SendNotificationRequestLinkResourceTypeCustomerContactInfo                  SendNotificationRequestLinkResourceType = "customer_contact_info"
	SendNotificationRequestLinkResourceTypeCustomerFreightPreferences           SendNotificationRequestLinkResourceType = "customer_freight_preferences"
	SendNotificationRequestLinkResourceTypeCustomerDefaults                     SendNotificationRequestLinkResourceType = "customer_defaults"
	SendNotificationRequestLinkResourceTypeCustomerLeadTime                     SendNotificationRequestLinkResourceType = "customer_lead_time"
	SendNotificationRequestLinkResourceTypeCustomerNotificationPreferences      SendNotificationRequestLinkResourceType = "customer_notification_preferences"
	SendNotificationRequestLinkResourceTypeOrderNotificationRecipient           SendNotificationRequestLinkResourceType = "order_notification_recipient"
	SendNotificationRequestLinkResourceTypeOrderDiscount                        SendNotificationRequestLinkResourceType = "order_discount"
	SendNotificationRequestLinkResourceTypeSalesOrderLine                       SendNotificationRequestLinkResourceType = "sales_order_line"
	SendNotificationRequestLinkResourceTypeSalesOrderType                       SendNotificationRequestLinkResourceType = "sales_order_type"
	SendNotificationRequestLinkResourceTypeSalesOrderStatus                     SendNotificationRequestLinkResourceType = "sales_order_status"
	SendNotificationRequestLinkResourceTypeMaterial                             SendNotificationRequestLinkResourceType = "material"
	SendNotificationRequestLinkResourceTypeSupplierMaterial                     SendNotificationRequestLinkResourceType = "supplier_material"
	SendNotificationRequestLinkResourceTypePart                                 SendNotificationRequestLinkResourceType = "part"
	SendNotificationRequestLinkResourceTypePermissionGroup                      SendNotificationRequestLinkResourceType = "permission_group"
	SendNotificationRequestLinkResourceTypePermission                           SendNotificationRequestLinkResourceType = "permission"
	SendNotificationRequestLinkResourceTypePick                                 SendNotificationRequestLinkResourceType = "pick"
	SendNotificationRequestLinkResourceTypePickLine                             SendNotificationRequestLinkResourceType = "pick_line"
	SendNotificationRequestLinkResourceTypeProductType                          SendNotificationRequestLinkResourceType = "product_type"
	SendNotificationRequestLinkResourceTypeProduction                           SendNotificationRequestLinkResourceType = "production"
	SendNotificationRequestLinkResourceTypeProductionFlow                       SendNotificationRequestLinkResourceType = "production_flow"
	SendNotificationRequestLinkResourceTypeMap                                  SendNotificationRequestLinkResourceType = "map"
	SendNotificationRequestLinkResourceTypePurchaseOrder                        SendNotificationRequestLinkResourceType = "purchase_order"
	SendNotificationRequestLinkResourceTypePurchaseOrderLine                    SendNotificationRequestLinkResourceType = "purchase_order_line"
	SendNotificationRequestLinkResourceTypeSupplier                             SendNotificationRequestLinkResourceType = "supplier"
	SendNotificationRequestLinkResourceTypeSupplierSummary                      SendNotificationRequestLinkResourceType = "supplier_summary"
	SendNotificationRequestLinkResourceTypeReceivableEntry                      SendNotificationRequestLinkResourceType = "receivable_entry"
	SendNotificationRequestLinkResourceTypeReceivingOrder                       SendNotificationRequestLinkResourceType = "receiving_order"
	SendNotificationRequestLinkResourceTypeReceivingOrderLine                   SendNotificationRequestLinkResourceType = "receiving_order_line"
	SendNotificationRequestLinkResourceTypeEmailContact                         SendNotificationRequestLinkResourceType = "email_contact"
	SendNotificationRequestLinkResourceTypeAllocationEntry                      SendNotificationRequestLinkResourceType = "allocation_entry"
	SendNotificationRequestLinkResourceTypeOpenCreditEntry                      SendNotificationRequestLinkResourceType = "open_credit_entry"
	SendNotificationRequestLinkResourceTypeVolumeDiscount                       SendNotificationRequestLinkResourceType = "volume_discount"
	SendNotificationRequestLinkResourceTypeVolumeDiscountTier                   SendNotificationRequestLinkResourceType = "volume_discount_tier"
	SendNotificationRequestLinkResourceTypeAnalyzeDeliveriesResponse            SendNotificationRequestLinkResourceType = "analyze_deliveries_response"
	SendNotificationRequestLinkResourceTypeAnalyzeManufacturingResponse         SendNotificationRequestLinkResourceType = "analyze_manufacturing_response"
	SendNotificationRequestLinkResourceTypeAnalyzeManufacturingBatchResponse    SendNotificationRequestLinkResourceType = "analyze_manufacturing_batch_response"
	SendNotificationRequestLinkResourceTypeAnalyzeQuarterlyOrdersResponse       SendNotificationRequestLinkResourceType = "analyze_quarterly_orders_response"
	SendNotificationRequestLinkResourceTypeAnalyzeNewCustomersResponse          SendNotificationRequestLinkResourceType = "analyze_new_customers_response"
	SendNotificationRequestLinkResourceTypeAnalyzeDemandForecastResponse        SendNotificationRequestLinkResourceType = "analyze_demand_forecast_response"
	SendNotificationRequestLinkResourceTypeAnalyzeOeeResponse                   SendNotificationRequestLinkResourceType = "analyze_oee_response"
	SendNotificationRequestLinkResourceTypeAnalyzeOeeTrendResponse              SendNotificationRequestLinkResourceType = "analyze_oee_trend_response"
	SendNotificationRequestLinkResourceTypeAnalyzeScheduleAttainmentResponse    SendNotificationRequestLinkResourceType = "analyze_schedule_attainment_response"
	SendNotificationRequestLinkResourceTypeCatalogProductLine                   SendNotificationRequestLinkResourceType = "catalog_product_line"
	SendNotificationRequestLinkResourceTypeCatalogCategory                      SendNotificationRequestLinkResourceType = "catalog_category"
	SendNotificationRequestLinkResourceTypeCatalogProduct                       SendNotificationRequestLinkResourceType = "catalog_product"
	SendNotificationRequestLinkResourceTypeCatalogProperty                      SendNotificationRequestLinkResourceType = "catalog_property"
	SendNotificationRequestLinkResourceTypeCatalogAttribute                     SendNotificationRequestLinkResourceType = "catalog_attribute"
	SendNotificationRequestLinkResourceTypeDcLocation                           SendNotificationRequestLinkResourceType = "dc_location"
	SendNotificationRequestLinkResourceTypeEdiRun                               SendNotificationRequestLinkResourceType = "edi_run"
	SendNotificationRequestLinkResourceTypeInventoryItem                        SendNotificationRequestLinkResourceType = "inventory_item"
	SendNotificationRequestLinkResourceTypeAnalyzeWeeksOfSalesResponse          SendNotificationRequestLinkResourceType = "analyze_weeks_of_sales_response"
	SendNotificationRequestLinkResourceTypeBulkReconcileItemsResponse           SendNotificationRequestLinkResourceType = "bulk_reconcile_items_response"
	SendNotificationRequestLinkResourceTypeSysProperty                          SendNotificationRequestLinkResourceType = "sys_property"
	SendNotificationRequestLinkResourceTypeSysPropertyType                      SendNotificationRequestLinkResourceType = "sys_property_type"
	SendNotificationRequestLinkResourceTypeSysPropertyValue                     SendNotificationRequestLinkResourceType = "sys_property_value"
	SendNotificationRequestLinkResourceTypeTerritory                            SendNotificationRequestLinkResourceType = "territory"
	SendNotificationRequestLinkResourceTypeTenancy                              SendNotificationRequestLinkResourceType = "tenancy"
	SendNotificationRequestLinkResourceTypeCheckoutSession                      SendNotificationRequestLinkResourceType = "checkout_session"
	SendNotificationRequestLinkResourceTypeEstimateRateResult                   SendNotificationRequestLinkResourceType = "estimate_rate_result"
	SendNotificationRequestLinkResourceTypeRateShopOption                       SendNotificationRequestLinkResourceType = "rate_shop_option"
	SendNotificationRequestLinkResourceTypeRateShopResult                       SendNotificationRequestLinkResourceType = "rate_shop_result"
	SendNotificationRequestLinkResourceTypeOwner                                SendNotificationRequestLinkResourceType = "owner"
	SendNotificationRequestLinkResourceTypeCreatedBy                            SendNotificationRequestLinkResourceType = "created_by"
	SendNotificationRequestLinkResourceTypeMessage                              SendNotificationRequestLinkResourceType = "message"
	SendNotificationRequestLinkResourceTypeAccountPhotoUploadResult             SendNotificationRequestLinkResourceType = "account_photo_upload_result"
	SendNotificationRequestLinkResourceTypeUserPhotoUploadResult                SendNotificationRequestLinkResourceType = "user_photo_upload_result"
	SendNotificationRequestLinkResourceTypeUserPhotoURL                         SendNotificationRequestLinkResourceType = "user_photo_url"
	SendNotificationRequestLinkResourceTypeBatchLot                             SendNotificationRequestLinkResourceType = "batch_lot"
	SendNotificationRequestLinkResourceTypeCheckDuplicateResult                 SendNotificationRequestLinkResourceType = "check_duplicate_result"
	SendNotificationRequestLinkResourceTypeItemTrendPoint                       SendNotificationRequestLinkResourceType = "item_trend_point"
	SendNotificationRequestLinkResourceTypeTenancyPendingRegistration           SendNotificationRequestLinkResourceType = "tenancy_pending_registration"
	SendNotificationRequestLinkResourceTypeInvoiceAllocationEntry               SendNotificationRequestLinkResourceType = "invoice_allocation_entry"
	SendNotificationRequestLinkResourceTypeAllocationCustomer                   SendNotificationRequestLinkResourceType = "allocation_customer"
	SendNotificationRequestLinkResourceTypeCheckoutSalesOrder                   SendNotificationRequestLinkResourceType = "checkout_sales_order"
	SendNotificationRequestLinkResourceTypeSalesOrderPriceQuote                 SendNotificationRequestLinkResourceType = "sales_order_price_quote"
	SendNotificationRequestLinkResourceTypeSalesOrderFreightQuote               SendNotificationRequestLinkResourceType = "sales_order_freight_quote"
	SendNotificationRequestLinkResourceTypeSalesOrderCommitmentQuote            SendNotificationRequestLinkResourceType = "sales_order_commitment_quote"
	SendNotificationRequestLinkResourceTypeOperatingCalendar                    SendNotificationRequestLinkResourceType = "operating_calendar"
	SendNotificationRequestLinkResourceTypeOperatingCalendarClosure             SendNotificationRequestLinkResourceType = "operating_calendar_closure"
	SendNotificationRequestLinkResourceTypeSalesOrderPriceQuoteLine             SendNotificationRequestLinkResourceType = "sales_order_price_quote_line"
	SendNotificationRequestLinkResourceTypeHubspotSyncJob                       SendNotificationRequestLinkResourceType = "hubspot_sync_job"
	SendNotificationRequestLinkResourceTypeHubspotSyncReport                    SendNotificationRequestLinkResourceType = "hubspot_sync_report"
	SendNotificationRequestLinkResourceTypeHubspotCompanyReview                 SendNotificationRequestLinkResourceType = "hubspot_company_review"
	SendNotificationRequestLinkResourceTypeHubspotCompanyCandidate              SendNotificationRequestLinkResourceType = "hubspot_company_candidate"
	SendNotificationRequestLinkResourceTypeHubspotSyncRecord                    SendNotificationRequestLinkResourceType = "hubspot_sync_record"
	SendNotificationRequestLinkResourceTypeContactMatch                         SendNotificationRequestLinkResourceType = "contact_match"
	SendNotificationRequestLinkResourceTypeReplyDraft                           SendNotificationRequestLinkResourceType = "reply_draft"
	SendNotificationRequestLinkResourceTypeConversationLink                     SendNotificationRequestLinkResourceType = "conversation_link"
	SendNotificationRequestLinkResourceTypeMessagingGroup                       SendNotificationRequestLinkResourceType = "messaging_group"
	SendNotificationRequestLinkResourceTypeMessagingGroupMember                 SendNotificationRequestLinkResourceType = "messaging_group_member"
	SendNotificationRequestLinkResourceTypePortalProfile                        SendNotificationRequestLinkResourceType = "portal_profile"
	SendNotificationRequestLinkResourceTypePortalRegistrationSession            SendNotificationRequestLinkResourceType = "portal_registration_session"
	SendNotificationRequestLinkResourceTypePortalRegistrationSessionData        SendNotificationRequestLinkResourceType = "portal_registration_session_data"
	SendNotificationRequestLinkResourceTypePackList                             SendNotificationRequestLinkResourceType = "pack_list"
	SendNotificationRequestLinkResourceTypePackListParty                        SendNotificationRequestLinkResourceType = "pack_list_party"
	SendNotificationRequestLinkResourceTypePackListLineItem                     SendNotificationRequestLinkResourceType = "pack_list_line_item"
	SendNotificationRequestLinkResourceTypePackListBackOrder                    SendNotificationRequestLinkResourceType = "pack_list_back_order"
	SendNotificationRequestLinkResourceTypePackListCase                         SendNotificationRequestLinkResourceType = "pack_list_case"
	SendNotificationRequestLinkResourceTypeJob                                  SendNotificationRequestLinkResourceType = "job"
	SendNotificationRequestLinkResourceTypeJobResult                            SendNotificationRequestLinkResourceType = "job_result"
	SendNotificationRequestLinkResourceTypeJobExport                            SendNotificationRequestLinkResourceType = "job_export"
	SendNotificationRequestLinkResourceTypeAnalyzeCustomerPricingResponse       SendNotificationRequestLinkResourceType = "analyze_customer_pricing_response"
	SendNotificationRequestLinkResourceTypeCustomerPricingFinding               SendNotificationRequestLinkResourceType = "customer_pricing_finding"
	SendNotificationRequestLinkResourceTypeCustomerPricingSummary               SendNotificationRequestLinkResourceType = "customer_pricing_summary"
	SendNotificationRequestLinkResourceTypeComputedRate                         SendNotificationRequestLinkResourceType = "computed_rate"
	SendNotificationRequestLinkResourceTypeComputedQuantity                     SendNotificationRequestLinkResourceType = "computed_quantity"
	SendNotificationRequestLinkResourceTypeAnalyzeRealizedMarginsResponse       SendNotificationRequestLinkResourceType = "analyze_realized_margins_response"
	SendNotificationRequestLinkResourceTypeRealizedMarginFinding                SendNotificationRequestLinkResourceType = "realized_margin_finding"
	SendNotificationRequestLinkResourceTypeRealizedMarginSummary                SendNotificationRequestLinkResourceType = "realized_margin_summary"
	SendNotificationRequestLinkResourceTypeShipmentRelated                      SendNotificationRequestLinkResourceType = "shipment_related"
	SendNotificationRequestLinkResourceTypeInvoiceRelated                       SendNotificationRequestLinkResourceType = "invoice_related"
	SendNotificationRequestLinkResourceTypePickRelated                          SendNotificationRequestLinkResourceType = "pick_related"
	SendNotificationRequestLinkResourceTypePickTotals                           SendNotificationRequestLinkResourceType = "pick_totals"
	SendNotificationRequestLinkResourceTypePickStageTotal                       SendNotificationRequestLinkResourceType = "pick_stage_total"
)

// How prominently the notification should be surfaced, from `low` through
// `urgent`.
type SendNotificationRequestPriority string

const (
	SendNotificationRequestPriorityLow    SendNotificationRequestPriority = "low"
	SendNotificationRequestPriorityNormal SendNotificationRequestPriority = "normal"
	SendNotificationRequestPriorityHigh   SendNotificationRequestPriority = "high"
	SendNotificationRequestPriorityUrgent SendNotificationRequestPriority = "urgent"
)

type MessagingNotificationNewParams struct {
	// Request to send an in-app notification.
	//
	// The target decides whether the notification goes to one member of the account or
	// to everyone in it.
	SendNotificationRequest SendNotificationRequestParam
	paramObj
}

func (r MessagingNotificationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SendNotificationRequest)
}
func (r *MessagingNotificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingNotificationGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingNotificationGetParams]'s query parameters as
// `url.Values`.
func (r MessagingNotificationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingNotificationListParams struct {
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
	// Return only notifications of this category, such as `chat.mention` or
	// `order.updated`.
	//
	// Any of "chat.message", "chat.mention", "chat.added", "order.updated",
	// "agent.run_completed", "agent.alert", "system.broadcast", "customer.registered".
	Category MessagingNotificationListParamsCategory `query:"category,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "sender", "resource".
	Include []string `query:"include,omitzero" json:"-"`
	// Return only notifications sent by these actors.
	//
	// A notification sent by a person is attributed to their account user id, not
	// their user id.
	SenderIDs []string `query:"sender_ids,omitzero" json:"-"`
	// Return only notifications sent by these kinds of actor.
	//
	// Notifications raised by the platform itself are attributed to the `system`
	// sender type but are returned without a sender.
	//
	// Any of "user", "group", "system", "agent", "apikey".
	SenderTypes []string `query:"sender_types,omitzero" json:"-"`
	// Return only notifications in this lifecycle state.
	//
	// When omitted, the response is the active feed: every notification that has not
	// been dismissed, whatever its seen or read state. Pass `dismissed` to review
	// notifications that were cleared out of the feed.
	//
	// Any of "unseen", "seen", "read", "dismissed".
	Status MessagingNotificationListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingNotificationListParams]'s query parameters as
// `url.Values`.
func (r MessagingNotificationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Return only notifications of this category, such as `chat.mention` or
// `order.updated`.
type MessagingNotificationListParamsCategory string

const (
	MessagingNotificationListParamsCategoryChatMessage        MessagingNotificationListParamsCategory = "chat.message"
	MessagingNotificationListParamsCategoryChatMention        MessagingNotificationListParamsCategory = "chat.mention"
	MessagingNotificationListParamsCategoryChatAdded          MessagingNotificationListParamsCategory = "chat.added"
	MessagingNotificationListParamsCategoryOrderUpdated       MessagingNotificationListParamsCategory = "order.updated"
	MessagingNotificationListParamsCategoryAgentRunCompleted  MessagingNotificationListParamsCategory = "agent.run_completed"
	MessagingNotificationListParamsCategoryAgentAlert         MessagingNotificationListParamsCategory = "agent.alert"
	MessagingNotificationListParamsCategorySystemBroadcast    MessagingNotificationListParamsCategory = "system.broadcast"
	MessagingNotificationListParamsCategoryCustomerRegistered MessagingNotificationListParamsCategory = "customer.registered"
)

// Return only notifications in this lifecycle state.
//
// When omitted, the response is the active feed: every notification that has not
// been dismissed, whatever its seen or read state. Pass `dismissed` to review
// notifications that were cleared out of the feed.
type MessagingNotificationListParamsStatus string

const (
	MessagingNotificationListParamsStatusUnseen    MessagingNotificationListParamsStatus = "unseen"
	MessagingNotificationListParamsStatusSeen      MessagingNotificationListParamsStatus = "seen"
	MessagingNotificationListParamsStatusRead      MessagingNotificationListParamsStatus = "read"
	MessagingNotificationListParamsStatusDismissed MessagingNotificationListParamsStatus = "dismissed"
)
