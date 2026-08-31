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
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and retrieve audit events.
//
// CoreAuditEventService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAuditEventService] method instead.
type CoreAuditEventService struct {
	options []option.RequestOption
}

// NewCoreAuditEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreAuditEventService(opts ...option.RequestOption) (r CoreAuditEventService) {
	r = CoreAuditEventService{}
	r.options = opts
	return
}

// Returns a single audit event by ID.
//
// The event is readable when your account is either the acting account or the
// account that was acted upon.
//
// This endpoint requires the permission: `audit_events:read`.
func (r *CoreAuditEventService) Get(ctx context.Context, id string, query CoreAuditEventGetParams, opts ...option.RequestOption) (res *AuditEvent, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/audit-events/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of audit events, newest first.
//
// Results cover every change where your account is either the acting account or
// the account that was acted upon, so a customer's or supplier's changes to your
// records appear alongside your own. The `q` parameter searches the resource type,
// action, resource ID, and originating request ID.
//
// This endpoint requires the permission: `audit_events:read`.
func (r *CoreAuditEventService) List(ctx context.Context, query CoreAuditEventListParams, opts ...option.RequestOption) (res *ListAuditEvent, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/audit-events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns every resource type an audit event can refer to, as plain strings.
//
// This is the accepted vocabulary for the `resource_types` filter when listing
// audit events. It is the API's complete resource-type list rather than a list
// derived from your account's data, so it includes types you may never have
// recorded events for.
//
// This endpoint requires the permission: `audit_events:read`.
func (r *CoreAuditEventService) GetResourceTypes(ctx context.Context, opts ...option.RequestOption) (res *ListObjectType, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/audit-events/resource-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// An immutable record of a single change to a resource, capturing who made the
// change, what changed, and when.
//
// Audit events are recorded automatically as mutations happen; they cannot be
// created, edited, or deleted through the API. Recording is asynchronous, so an
// event may take a moment to become readable after the request that caused it has
// returned. An update that leaves every tracked field at its existing value
// records no event unless the mutation attaches metadata of its own — a password
// rotation, for example, records metadata and no field changes.
type AuditEvent struct {
	// Audit event ID.
	ID string `json:"id" api:"required"`
	// An organization on OpenMRP, including its branding and customer portal
	// sub-resources.
	//
	// Your own account and any customer or supplier account you trade with are both
	// represented by this object.
	Account Account `json:"account" api:"required"`
	// The type of action this event records.
	//
	//   - `create`: the resource was created.
	//   - `update`: one or more fields were changed.
	//   - `delete`: the resource was deleted.
	//   - `restore`: a previously deleted resource was restored.
	//   - `archive`: the resource was archived.
	//   - `approve`: a human approved a gated action, such as allowing a review-gated
	//     agent tool to run.
	//   - `deny`: a human denied a gated action, such as rejecting a review-gated agent
	//     tool.
	//
	// Any of "create", "update", "upsert", "delete", "restore", "archive", "approve",
	// "deny".
	Action AuditEventAction `json:"action" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Changes ListAuditFieldChange `json:"changes" api:"required"`
	// When the audit event record was written.
	//
	// Slightly later than `occurred_at`, since events are recorded out of band from
	// the request that caused them.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Idempotency key of the originating request.
	IdempotencyKey string `json:"idempotency_key" api:"required"`
	// Arbitrary JSON metadata for the mutation (e.g. reason, source, tags). Encoded as
	// a JSON value (object, array, string, number, boolean, or null), not a
	// JSON-encoded string.
	Metadata any `json:"metadata" api:"required"`
	// Resource type identifier.
	//
	// Any of "audit_event".
	Object AuditEventObject `json:"object" api:"required"`
	// When the audited mutation occurred.
	//
	// Audit events are ordered and date-filtered by this timestamp rather than by
	// `created_at`.
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// A log of a single API request, capturing its route, outcome, latency, and actor.
	//
	// Logs are written after the response has been sent, so a new entry may take a
	// moment to become readable.
	Request RequestLog `json:"request" api:"required"`
	// Audited resource ID.
	ResourceID string `json:"resource_id" api:"required"`
	// Resource type of the audited entity.
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
	// "email_inbox", "email_sender", "portal_domain", "dns_record",
	// "inventory_change_log", "invoice", "invoice_summary", "invoice_line",
	// "invoice_allocation", "invoice_for_payment", "shipment", "shipment_summary",
	// "shipment_line", "shipping_case", "shipping_case_label_url", "settlement",
	// "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	ResourceType AuditEventResourceType `json:"resource_type" api:"required"`
	// Originating client IP address.
	SourceIP string `json:"source_ip" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Account        respjson.Field
		Action         respjson.Field
		Actor          respjson.Field
		Changes        respjson.Field
		CreatedAt      respjson.Field
		IdempotencyKey respjson.Field
		Metadata       respjson.Field
		Object         respjson.Field
		OccurredAt     respjson.Field
		Request        respjson.Field
		ResourceID     respjson.Field
		ResourceType   respjson.Field
		SourceIP       respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEvent) RawJSON() string { return r.JSON.raw }
func (r *AuditEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of action this event records.
//
//   - `create`: the resource was created.
//   - `update`: one or more fields were changed.
//   - `delete`: the resource was deleted.
//   - `restore`: a previously deleted resource was restored.
//   - `archive`: the resource was archived.
//   - `approve`: a human approved a gated action, such as allowing a review-gated
//     agent tool to run.
//   - `deny`: a human denied a gated action, such as rejecting a review-gated agent
//     tool.
type AuditEventAction string

const (
	AuditEventActionCreate  AuditEventAction = "create"
	AuditEventActionUpdate  AuditEventAction = "update"
	AuditEventActionUpsert  AuditEventAction = "upsert"
	AuditEventActionDelete  AuditEventAction = "delete"
	AuditEventActionRestore AuditEventAction = "restore"
	AuditEventActionArchive AuditEventAction = "archive"
	AuditEventActionApprove AuditEventAction = "approve"
	AuditEventActionDeny    AuditEventAction = "deny"
)

// Resource type identifier.
type AuditEventObject string

const (
	AuditEventObjectAuditEvent AuditEventObject = "audit_event"
)

// Resource type of the audited entity.
type AuditEventResourceType string

const (
	AuditEventResourceTypeAccount                              AuditEventResourceType = "account"
	AuditEventResourceTypeActor                                AuditEventResourceType = "actor"
	AuditEventResourceTypeEntity                               AuditEventResourceType = "entity"
	AuditEventResourceTypeRecord                               AuditEventResourceType = "record"
	AuditEventResourceTypeFreight                              AuditEventResourceType = "freight"
	AuditEventResourceTypeCommitment                           AuditEventResourceType = "commitment"
	AuditEventResourceTypeSalesOrderTotals                     AuditEventResourceType = "sales_order_totals"
	AuditEventResourceTypeSalesOrderStageTotal                 AuditEventResourceType = "sales_order_stage_total"
	AuditEventResourceTypeSalesOrderRelated                    AuditEventResourceType = "sales_order_related"
	AuditEventResourceTypeOrderContact                         AuditEventResourceType = "order_contact"
	AuditEventResourceTypeUser                                 AuditEventResourceType = "user"
	AuditEventResourceTypeAddress                              AuditEventResourceType = "address"
	AuditEventResourceTypeAPIKey                               AuditEventResourceType = "api_key"
	AuditEventResourceTypeCreatedAPIKey                        AuditEventResourceType = "created_api_key"
	AuditEventResourceTypeRefreshToken                         AuditEventResourceType = "refresh_token"
	AuditEventResourceTypeList                                 AuditEventResourceType = "list"
	AuditEventResourceTypeSandbox                              AuditEventResourceType = "sandbox"
	AuditEventResourceTypeRegistrationSession                  AuditEventResourceType = "registration_session"
	AuditEventResourceTypePricingPlan                          AuditEventResourceType = "pricing_plan"
	AuditEventResourceTypeAccountPlan                          AuditEventResourceType = "account_plan"
	AuditEventResourceTypePlanChange                           AuditEventResourceType = "plan_change"
	AuditEventResourceTypeEnterpriseInquiry                    AuditEventResourceType = "enterprise_inquiry"
	AuditEventResourceTypeRequestLog                           AuditEventResourceType = "request_log"
	AuditEventResourceTypeAuditEvent                           AuditEventResourceType = "audit_event"
	AuditEventResourceTypeAuditFieldChange                     AuditEventResourceType = "audit_field_change"
	AuditEventResourceTypeRole                                 AuditEventResourceType = "role"
	AuditEventResourceTypeUnit                                 AuditEventResourceType = "unit"
	AuditEventResourceTypeAccountAffiliation                   AuditEventResourceType = "account_affiliation"
	AuditEventResourceTypeAgentDefinition                      AuditEventResourceType = "agent_definition"
	AuditEventResourceTypeAvailableTool                        AuditEventResourceType = "available_tool"
	AuditEventResourceTypeAgentDefinitionTool                  AuditEventResourceType = "agent_definition_tool"
	AuditEventResourceTypeAgentAccountStatus                   AuditEventResourceType = "agent_account_status"
	AuditEventResourceTypeAgentRun                             AuditEventResourceType = "agent_run"
	AuditEventResourceTypeAgentAction                          AuditEventResourceType = "agent_action"
	AuditEventResourceTypeAgentRunStep                         AuditEventResourceType = "agent_run_step"
	AuditEventResourceTypeAgentTokenUsage                      AuditEventResourceType = "agent_token_usage"
	AuditEventResourceTypeAgentMemory                          AuditEventResourceType = "agent_memory"
	AuditEventResourceTypeNotification                         AuditEventResourceType = "notification"
	AuditEventResourceTypeNotificationUnreadCount              AuditEventResourceType = "notification_unread_count"
	AuditEventResourceTypeNotificationSendResult               AuditEventResourceType = "notification_send_result"
	AuditEventResourceTypeNotificationUnreadSummary            AuditEventResourceType = "notification_unread_summary"
	AuditEventResourceTypeAnnouncement                         AuditEventResourceType = "announcement"
	AuditEventResourceTypeConversation                         AuditEventResourceType = "conversation"
	AuditEventResourceTypeSupportCase                          AuditEventResourceType = "support_case"
	AuditEventResourceTypeConversationParticipant              AuditEventResourceType = "conversation_participant"
	AuditEventResourceTypeReadCursor                           AuditEventResourceType = "read_cursor"
	AuditEventResourceTypeChatMessage                          AuditEventResourceType = "chat_message"
	AuditEventResourceTypeNotificationUnreadSummaryAccount     AuditEventResourceType = "notification_unread_summary_account"
	AuditEventResourceTypeMessagingBlock                       AuditEventResourceType = "messaging_block"
	AuditEventResourceTypeNotificationPreference               AuditEventResourceType = "notification_preference"
	AuditEventResourceTypeMessageAttachment                    AuditEventResourceType = "message_attachment"
	AuditEventResourceTypeAttachmentUploadTarget               AuditEventResourceType = "attachment_upload_target"
	AuditEventResourceTypeScheduledMessage                     AuditEventResourceType = "scheduled_message"
	AuditEventResourceTypeMessagingContact                     AuditEventResourceType = "messaging_contact"
	AuditEventResourceTypeMessageReport                        AuditEventResourceType = "message_report"
	AuditEventResourceTypeToolGroup                            AuditEventResourceType = "tool_group"
	AuditEventResourceTypeModel                                AuditEventResourceType = "model"
	AuditEventResourceTypePaymentTerm                          AuditEventResourceType = "payment_term"
	AuditEventResourceTypeShippingTerm                         AuditEventResourceType = "shipping_term"
	AuditEventResourceTypeQuantity                             AuditEventResourceType = "quantity"
	AuditEventResourceTypeAccountGroup                         AuditEventResourceType = "account_group"
	AuditEventResourceTypeSupportRoute                         AuditEventResourceType = "support_route"
	AuditEventResourceTypeSupportAvailability                  AuditEventResourceType = "support_availability"
	AuditEventResourceTypeAccountStatus                        AuditEventResourceType = "account_status"
	AuditEventResourceTypeGeolocation                          AuditEventResourceType = "geolocation"
	AuditEventResourceTypeAccountUser                          AuditEventResourceType = "account_user"
	AuditEventResourceTypeDepartment                           AuditEventResourceType = "department"
	AuditEventResourceTypeAccountIntegration                   AuditEventResourceType = "account_integration"
	AuditEventResourceTypeAccountPrice                         AuditEventResourceType = "account_price"
	AuditEventResourceTypeProductLine                          AuditEventResourceType = "product_line"
	AuditEventResourceTypeItemCategory                         AuditEventResourceType = "item_category"
	AuditEventResourceTypeAttribute                            AuditEventResourceType = "attribute"
	AuditEventResourceTypeRate                                 AuditEventResourceType = "rate"
	AuditEventResourceTypeAccountGroupProductLineAccess        AuditEventResourceType = "account_group_product_line_access"
	AuditEventResourceTypeSalesTarget                          AuditEventResourceType = "sales_target"
	AuditEventResourceTypeAdjustmentType                       AuditEventResourceType = "adjustment_type"
	AuditEventResourceTypeAccountBranding                      AuditEventResourceType = "account_branding"
	AuditEventResourceTypeAccountPortal                        AuditEventResourceType = "account_portal"
	AuditEventResourceTypeAccountLogoURL                       AuditEventResourceType = "account_logo_url"
	AuditEventResourceTypeAccountFaviconURL                    AuditEventResourceType = "account_favicon_url"
	AuditEventResourceTypePublicAccount                        AuditEventResourceType = "public_account"
	AuditEventResourceTypeProperty                             AuditEventResourceType = "property"
	AuditEventResourceTypeCarrier                              AuditEventResourceType = "carrier"
	AuditEventResourceTypeServiceLevel                         AuditEventResourceType = "service_level"
	AuditEventResourceTypeItem                                 AuditEventResourceType = "item"
	AuditEventResourceTypeItemLotDefault                       AuditEventResourceType = "item_lot_default"
	AuditEventResourceTypeItemInventory                        AuditEventResourceType = "item_inventory"
	AuditEventResourceTypeProduct                              AuditEventResourceType = "product"
	AuditEventResourceTypeBatch                                AuditEventResourceType = "batch"
	AuditEventResourceTypeBatchFlowNode                        AuditEventResourceType = "batch_flow_node"
	AuditEventResourceTypeScanningConsumption                  AuditEventResourceType = "scanning_consumption"
	AuditEventResourceTypeOpenBatchSummary                     AuditEventResourceType = "open_batch_summary"
	AuditEventResourceTypeScanningProductionStepInfo           AuditEventResourceType = "scanning_production_step_info"
	AuditEventResourceTypeScanningStation                      AuditEventResourceType = "scanning_station"
	AuditEventResourceTypeProductionStep                       AuditEventResourceType = "production_step"
	AuditEventResourceTypeProductionRun                        AuditEventResourceType = "production_run"
	AuditEventResourceTypeMachine                              AuditEventResourceType = "machine"
	AuditEventResourceTypeMachineStatus                        AuditEventResourceType = "machine_status"
	AuditEventResourceTypeMachineDowntimeEvent                 AuditEventResourceType = "machine_downtime_event"
	AuditEventResourceTypeDemandOverride                       AuditEventResourceType = "demand_override"
	AuditEventResourceTypeDemandOverrideType                   AuditEventResourceType = "demand_override_type"
	AuditEventResourceTypeMachineDowntimeReason                AuditEventResourceType = "machine_downtime_reason"
	AuditEventResourceTypeProductionSchedulePreview            AuditEventResourceType = "production_schedule_preview"
	AuditEventResourceTypeProductionScheduleRegeneratePreview  AuditEventResourceType = "production_schedule_regenerate_preview"
	AuditEventResourceTypeProductionSchedule                   AuditEventResourceType = "production_schedule"
	AuditEventResourceTypeProductionScheduleLine               AuditEventResourceType = "production_schedule_line"
	AuditEventResourceTypeProductionScheduleDeviation          AuditEventResourceType = "production_schedule_deviation"
	AuditEventResourceTypeProductionScheduleDerivedLine        AuditEventResourceType = "production_schedule_derived_line"
	AuditEventResourceTypeProductionScheduleSettings           AuditEventResourceType = "production_schedule_settings"
	AuditEventResourceTypeProductionScheduleResourceSetting    AuditEventResourceType = "production_schedule_resource_setting"
	AuditEventResourceTypeProductionScheduleItemSetting        AuditEventResourceType = "production_schedule_item_setting"
	AuditEventResourceTypeFulfillmentRecommendation            AuditEventResourceType = "fulfillment_recommendation"
	AuditEventResourceTypeAnalyzeDeliveryPerformanceResponse   AuditEventResourceType = "analyze_delivery_performance_response"
	AuditEventResourceTypeDeliveryPerformance                  AuditEventResourceType = "delivery_performance"
	AuditEventResourceTypeDeliveryBacklogBucket                AuditEventResourceType = "delivery_backlog_bucket"
	AuditEventResourceTypeDeliveryLatenessBucket               AuditEventResourceType = "delivery_lateness_bucket"
	AuditEventResourceTypeDeliveryBreakdown                    AuditEventResourceType = "delivery_breakdown"
	AuditEventResourceTypeAnalyzeSalesBreakdownResponse        AuditEventResourceType = "analyze_sales_breakdown_response"
	AuditEventResourceTypeSalesTotals                          AuditEventResourceType = "sales_totals"
	AuditEventResourceTypeSalesBreakdown                       AuditEventResourceType = "sales_breakdown"
	AuditEventResourceTypeScheduleOrderCoverage                AuditEventResourceType = "schedule_order_coverage"
	AuditEventResourceTypeScheduleOrderCoverageLine            AuditEventResourceType = "schedule_order_coverage_line"
	AuditEventResourceTypeScheduleDeviationType                AuditEventResourceType = "schedule_deviation_type"
	AuditEventResourceTypeScheduleAtRiskOrder                  AuditEventResourceType = "schedule_at_risk_order"
	AuditEventResourceTypeProductionScheduleFinishedPolicy     AuditEventResourceType = "production_schedule_finished_policy"
	AuditEventResourceTypeProductionScheduleFinishingLine      AuditEventResourceType = "production_schedule_finishing_line"
	AuditEventResourceTypeProductionScheduleWeekRelease        AuditEventResourceType = "production_schedule_week_release"
	AuditEventResourceTypeProductionScheduleWeekReleasePreview AuditEventResourceType = "production_schedule_week_release_preview"
	AuditEventResourceTypeProductionScheduleItemPolicy         AuditEventResourceType = "production_schedule_item_policy"
	AuditEventResourceTypeChildAccount                         AuditEventResourceType = "child_account"
	AuditEventResourceTypeUnitGroup                            AuditEventResourceType = "unit_group"
	AuditEventResourceTypeUnitGroupUnit                        AuditEventResourceType = "unit_group_unit"
	AuditEventResourceTypeConsumption                          AuditEventResourceType = "consumption"
	AuditEventResourceTypeCustomerProductLineAccess            AuditEventResourceType = "customer_product_line_access"
	AuditEventResourceTypeCustomer                             AuditEventResourceType = "customer"
	AuditEventResourceTypeFrequentlyOrderedProduct             AuditEventResourceType = "frequently_ordered_product"
	AuditEventResourceTypePriority                             AuditEventResourceType = "priority"
	AuditEventResourceTypeDelivery                             AuditEventResourceType = "delivery"
	AuditEventResourceTypeDeliveryLine                         AuditEventResourceType = "delivery_line"
	AuditEventResourceTypeSalesOrder                           AuditEventResourceType = "sales_order"
	AuditEventResourceTypeLocation                             AuditEventResourceType = "location"
	AuditEventResourceTypeLocationType                         AuditEventResourceType = "location_type"
	AuditEventResourceTypeLot                                  AuditEventResourceType = "lot"
	AuditEventResourceTypeEmailLog                             AuditEventResourceType = "email_log"
	AuditEventResourceTypeEmailDomain                          AuditEventResourceType = "email_domain"
	AuditEventResourceTypeEmailInbox                           AuditEventResourceType = "email_inbox"
	AuditEventResourceTypeEmailSender                          AuditEventResourceType = "email_sender"
	AuditEventResourceTypePortalDomain                         AuditEventResourceType = "portal_domain"
	AuditEventResourceTypeDNSRecord                            AuditEventResourceType = "dns_record"
	AuditEventResourceTypeInventoryChangeLog                   AuditEventResourceType = "inventory_change_log"
	AuditEventResourceTypeInvoice                              AuditEventResourceType = "invoice"
	AuditEventResourceTypeInvoiceSummary                       AuditEventResourceType = "invoice_summary"
	AuditEventResourceTypeInvoiceLine                          AuditEventResourceType = "invoice_line"
	AuditEventResourceTypeInvoiceAllocation                    AuditEventResourceType = "invoice_allocation"
	AuditEventResourceTypeInvoiceForPayment                    AuditEventResourceType = "invoice_for_payment"
	AuditEventResourceTypeShipment                             AuditEventResourceType = "shipment"
	AuditEventResourceTypeShipmentSummary                      AuditEventResourceType = "shipment_summary"
	AuditEventResourceTypeShipmentLine                         AuditEventResourceType = "shipment_line"
	AuditEventResourceTypeShippingCase                         AuditEventResourceType = "shipping_case"
	AuditEventResourceTypeShippingCaseLabelURL                 AuditEventResourceType = "shipping_case_label_url"
	AuditEventResourceTypeSettlement                           AuditEventResourceType = "settlement"
	AuditEventResourceTypeSettlementSummary                    AuditEventResourceType = "settlement_summary"
	AuditEventResourceTypeRolePermission                       AuditEventResourceType = "role_permission"
	AuditEventResourceTypeRegistrationFlow                     AuditEventResourceType = "registration_flow"
	AuditEventResourceTypeRegistrationFlowOption               AuditEventResourceType = "registration_flow_option"
	AuditEventResourceTypeTransaction                          AuditEventResourceType = "transaction"
	AuditEventResourceTypeTransactionSummary                   AuditEventResourceType = "transaction_summary"
	AuditEventResourceTypeTransactionMethod                    AuditEventResourceType = "transaction_method"
	AuditEventResourceTypeTransactionType                      AuditEventResourceType = "transaction_type"
	AuditEventResourceTypeTransactionAllocation                AuditEventResourceType = "transaction_allocation"
	AuditEventResourceTypeUsageItem                            AuditEventResourceType = "usage_item"
	AuditEventResourceTypeAccountUsageResponse                 AuditEventResourceType = "account_usage_response"
	AuditEventResourceTypeSubscriptionInfo                     AuditEventResourceType = "subscription_info"
	AuditEventResourceTypeBillingPortalSessionResponse         AuditEventResourceType = "billing_portal_session_response"
	AuditEventResourceTypeSwitchPlanResponse                   AuditEventResourceType = "switch_plan_response"
	AuditEventResourceTypeEnsureBillingCustomerResponse        AuditEventResourceType = "ensure_billing_customer_response"
	AuditEventResourceTypeSpendingCapResponse                  AuditEventResourceType = "spending_cap_response"
	AuditEventResourceTypeAgentSpendInfo                       AuditEventResourceType = "agent_spend_info"
	AuditEventResourceTypeWebhookResponse                      AuditEventResourceType = "webhook_response"
	AuditEventResourceTypeAddressSuggestion                    AuditEventResourceType = "address_suggestion"
	AuditEventResourceTypeAddressComponents                    AuditEventResourceType = "address_components"
	AuditEventResourceTypeAddressDetailsResult                 AuditEventResourceType = "address_details_result"
	AuditEventResourceTypeValidatedAddress                     AuditEventResourceType = "validated_address"
	AuditEventResourceTypePlanLimit                            AuditEventResourceType = "plan_limit"
	AuditEventResourceTypePlanChangeProration                  AuditEventResourceType = "plan_change_proration"
	AuditEventResourceTypePlanChangeLineItem                   AuditEventResourceType = "plan_change_line_item"
	AuditEventResourceTypeSetupBillingResponse                 AuditEventResourceType = "setup_billing_response"
	AuditEventResourceTypeConfirmPaymentResponse               AuditEventResourceType = "confirm_payment_response"
	AuditEventResourceTypeOAuthResponse                        AuditEventResourceType = "oauth_response"
	AuditEventResourceTypeOAuthStatusResponse                  AuditEventResourceType = "oauth_status_response"
	AuditEventResourceTypeStripePublishableKey                 AuditEventResourceType = "stripe_publishable_key"
	AuditEventResourceTypeStripeStatus                         AuditEventResourceType = "stripe_status"
	AuditEventResourceTypeHealthcheck                          AuditEventResourceType = "healthcheck"
	AuditEventResourceTypeAgentDefinitionConfig                AuditEventResourceType = "agent_definition_config"
	AuditEventResourceTypeTriggerConfig                        AuditEventResourceType = "trigger_config"
	AuditEventResourceTypeCustomerContactInfo                  AuditEventResourceType = "customer_contact_info"
	AuditEventResourceTypeCustomerFreightPreferences           AuditEventResourceType = "customer_freight_preferences"
	AuditEventResourceTypeCustomerDefaults                     AuditEventResourceType = "customer_defaults"
	AuditEventResourceTypeCustomerLeadTime                     AuditEventResourceType = "customer_lead_time"
	AuditEventResourceTypeCustomerNotificationPreferences      AuditEventResourceType = "customer_notification_preferences"
	AuditEventResourceTypeOrderNotificationRecipient           AuditEventResourceType = "order_notification_recipient"
	AuditEventResourceTypeOrderDiscount                        AuditEventResourceType = "order_discount"
	AuditEventResourceTypeSalesOrderLine                       AuditEventResourceType = "sales_order_line"
	AuditEventResourceTypeSalesOrderType                       AuditEventResourceType = "sales_order_type"
	AuditEventResourceTypeSalesOrderStatus                     AuditEventResourceType = "sales_order_status"
	AuditEventResourceTypeMaterial                             AuditEventResourceType = "material"
	AuditEventResourceTypeSupplierMaterial                     AuditEventResourceType = "supplier_material"
	AuditEventResourceTypePart                                 AuditEventResourceType = "part"
	AuditEventResourceTypePermissionGroup                      AuditEventResourceType = "permission_group"
	AuditEventResourceTypePermission                           AuditEventResourceType = "permission"
	AuditEventResourceTypePick                                 AuditEventResourceType = "pick"
	AuditEventResourceTypePickLine                             AuditEventResourceType = "pick_line"
	AuditEventResourceTypeProductType                          AuditEventResourceType = "product_type"
	AuditEventResourceTypeProduction                           AuditEventResourceType = "production"
	AuditEventResourceTypeProductionFlow                       AuditEventResourceType = "production_flow"
	AuditEventResourceTypeMap                                  AuditEventResourceType = "map"
	AuditEventResourceTypePurchaseOrder                        AuditEventResourceType = "purchase_order"
	AuditEventResourceTypePurchaseOrderLine                    AuditEventResourceType = "purchase_order_line"
	AuditEventResourceTypeSupplier                             AuditEventResourceType = "supplier"
	AuditEventResourceTypeSupplierSummary                      AuditEventResourceType = "supplier_summary"
	AuditEventResourceTypeReceivableEntry                      AuditEventResourceType = "receivable_entry"
	AuditEventResourceTypeReceivingOrder                       AuditEventResourceType = "receiving_order"
	AuditEventResourceTypeReceivingOrderLine                   AuditEventResourceType = "receiving_order_line"
	AuditEventResourceTypeEmailContact                         AuditEventResourceType = "email_contact"
	AuditEventResourceTypeAllocationEntry                      AuditEventResourceType = "allocation_entry"
	AuditEventResourceTypeOpenCreditEntry                      AuditEventResourceType = "open_credit_entry"
	AuditEventResourceTypeVolumeDiscount                       AuditEventResourceType = "volume_discount"
	AuditEventResourceTypeVolumeDiscountTier                   AuditEventResourceType = "volume_discount_tier"
	AuditEventResourceTypeAnalyzeDeliveriesResponse            AuditEventResourceType = "analyze_deliveries_response"
	AuditEventResourceTypeAnalyzeManufacturingResponse         AuditEventResourceType = "analyze_manufacturing_response"
	AuditEventResourceTypeAnalyzeManufacturingBatchResponse    AuditEventResourceType = "analyze_manufacturing_batch_response"
	AuditEventResourceTypeAnalyzeQuarterlyOrdersResponse       AuditEventResourceType = "analyze_quarterly_orders_response"
	AuditEventResourceTypeAnalyzeNewCustomersResponse          AuditEventResourceType = "analyze_new_customers_response"
	AuditEventResourceTypeAnalyzeDemandForecastResponse        AuditEventResourceType = "analyze_demand_forecast_response"
	AuditEventResourceTypeAnalyzeOeeResponse                   AuditEventResourceType = "analyze_oee_response"
	AuditEventResourceTypeAnalyzeOeeTrendResponse              AuditEventResourceType = "analyze_oee_trend_response"
	AuditEventResourceTypeAnalyzeScheduleAttainmentResponse    AuditEventResourceType = "analyze_schedule_attainment_response"
	AuditEventResourceTypeCatalogProductLine                   AuditEventResourceType = "catalog_product_line"
	AuditEventResourceTypeCatalogCategory                      AuditEventResourceType = "catalog_category"
	AuditEventResourceTypeCatalogProduct                       AuditEventResourceType = "catalog_product"
	AuditEventResourceTypeCatalogProperty                      AuditEventResourceType = "catalog_property"
	AuditEventResourceTypeCatalogAttribute                     AuditEventResourceType = "catalog_attribute"
	AuditEventResourceTypeDcLocation                           AuditEventResourceType = "dc_location"
	AuditEventResourceTypeEdiRun                               AuditEventResourceType = "edi_run"
	AuditEventResourceTypeInventoryItem                        AuditEventResourceType = "inventory_item"
	AuditEventResourceTypeAnalyzeWeeksOfSalesResponse          AuditEventResourceType = "analyze_weeks_of_sales_response"
	AuditEventResourceTypeBulkReconcileItemsResponse           AuditEventResourceType = "bulk_reconcile_items_response"
	AuditEventResourceTypeSysProperty                          AuditEventResourceType = "sys_property"
	AuditEventResourceTypeSysPropertyType                      AuditEventResourceType = "sys_property_type"
	AuditEventResourceTypeSysPropertyValue                     AuditEventResourceType = "sys_property_value"
	AuditEventResourceTypeTerritory                            AuditEventResourceType = "territory"
	AuditEventResourceTypeTenancy                              AuditEventResourceType = "tenancy"
	AuditEventResourceTypeCheckoutSession                      AuditEventResourceType = "checkout_session"
	AuditEventResourceTypeEstimateRateResult                   AuditEventResourceType = "estimate_rate_result"
	AuditEventResourceTypeRateShopOption                       AuditEventResourceType = "rate_shop_option"
	AuditEventResourceTypeRateShopResult                       AuditEventResourceType = "rate_shop_result"
	AuditEventResourceTypeOwner                                AuditEventResourceType = "owner"
	AuditEventResourceTypeCreatedBy                            AuditEventResourceType = "created_by"
	AuditEventResourceTypeMessage                              AuditEventResourceType = "message"
	AuditEventResourceTypeAccountPhotoUploadResult             AuditEventResourceType = "account_photo_upload_result"
	AuditEventResourceTypeUserPhotoUploadResult                AuditEventResourceType = "user_photo_upload_result"
	AuditEventResourceTypeUserPhotoURL                         AuditEventResourceType = "user_photo_url"
	AuditEventResourceTypeBatchLot                             AuditEventResourceType = "batch_lot"
	AuditEventResourceTypeCheckDuplicateResult                 AuditEventResourceType = "check_duplicate_result"
	AuditEventResourceTypeItemTrendPoint                       AuditEventResourceType = "item_trend_point"
	AuditEventResourceTypeTenancyPendingRegistration           AuditEventResourceType = "tenancy_pending_registration"
	AuditEventResourceTypeInvoiceAllocationEntry               AuditEventResourceType = "invoice_allocation_entry"
	AuditEventResourceTypeAllocationCustomer                   AuditEventResourceType = "allocation_customer"
	AuditEventResourceTypeCheckoutSalesOrder                   AuditEventResourceType = "checkout_sales_order"
	AuditEventResourceTypeSalesOrderPriceQuote                 AuditEventResourceType = "sales_order_price_quote"
	AuditEventResourceTypeSalesOrderFreightQuote               AuditEventResourceType = "sales_order_freight_quote"
	AuditEventResourceTypeSalesOrderCommitmentQuote            AuditEventResourceType = "sales_order_commitment_quote"
	AuditEventResourceTypeOperatingCalendar                    AuditEventResourceType = "operating_calendar"
	AuditEventResourceTypeOperatingCalendarClosure             AuditEventResourceType = "operating_calendar_closure"
	AuditEventResourceTypeSalesOrderPriceQuoteLine             AuditEventResourceType = "sales_order_price_quote_line"
	AuditEventResourceTypeHubspotSyncJob                       AuditEventResourceType = "hubspot_sync_job"
	AuditEventResourceTypeHubspotSyncReport                    AuditEventResourceType = "hubspot_sync_report"
	AuditEventResourceTypeHubspotCompanyReview                 AuditEventResourceType = "hubspot_company_review"
	AuditEventResourceTypeHubspotCompanyCandidate              AuditEventResourceType = "hubspot_company_candidate"
	AuditEventResourceTypeHubspotSyncRecord                    AuditEventResourceType = "hubspot_sync_record"
	AuditEventResourceTypeContactMatch                         AuditEventResourceType = "contact_match"
	AuditEventResourceTypeReplyDraft                           AuditEventResourceType = "reply_draft"
	AuditEventResourceTypeConversationLink                     AuditEventResourceType = "conversation_link"
	AuditEventResourceTypeMessagingGroup                       AuditEventResourceType = "messaging_group"
	AuditEventResourceTypeMessagingGroupMember                 AuditEventResourceType = "messaging_group_member"
	AuditEventResourceTypePortalProfile                        AuditEventResourceType = "portal_profile"
	AuditEventResourceTypePortalRegistrationSession            AuditEventResourceType = "portal_registration_session"
	AuditEventResourceTypePortalRegistrationSessionData        AuditEventResourceType = "portal_registration_session_data"
	AuditEventResourceTypePackList                             AuditEventResourceType = "pack_list"
	AuditEventResourceTypePackListParty                        AuditEventResourceType = "pack_list_party"
	AuditEventResourceTypePackListLineItem                     AuditEventResourceType = "pack_list_line_item"
	AuditEventResourceTypePackListBackOrder                    AuditEventResourceType = "pack_list_back_order"
	AuditEventResourceTypePackListCase                         AuditEventResourceType = "pack_list_case"
	AuditEventResourceTypeJob                                  AuditEventResourceType = "job"
	AuditEventResourceTypeJobResult                            AuditEventResourceType = "job_result"
	AuditEventResourceTypeJobExport                            AuditEventResourceType = "job_export"
	AuditEventResourceTypeAnalyzeCustomerPricingResponse       AuditEventResourceType = "analyze_customer_pricing_response"
	AuditEventResourceTypeCustomerPricingFinding               AuditEventResourceType = "customer_pricing_finding"
	AuditEventResourceTypeCustomerPricingSummary               AuditEventResourceType = "customer_pricing_summary"
	AuditEventResourceTypeComputedRate                         AuditEventResourceType = "computed_rate"
	AuditEventResourceTypeComputedQuantity                     AuditEventResourceType = "computed_quantity"
	AuditEventResourceTypeAnalyzeRealizedMarginsResponse       AuditEventResourceType = "analyze_realized_margins_response"
	AuditEventResourceTypeRealizedMarginFinding                AuditEventResourceType = "realized_margin_finding"
	AuditEventResourceTypeRealizedMarginSummary                AuditEventResourceType = "realized_margin_summary"
	AuditEventResourceTypeShipmentRelated                      AuditEventResourceType = "shipment_related"
	AuditEventResourceTypeInvoiceRelated                       AuditEventResourceType = "invoice_related"
	AuditEventResourceTypePickRelated                          AuditEventResourceType = "pick_related"
	AuditEventResourceTypePickTotals                           AuditEventResourceType = "pick_totals"
	AuditEventResourceTypePickStageTotal                       AuditEventResourceType = "pick_stage_total"
)

// Field-level before/after transition recorded during a mutation.
type AuditFieldChange struct {
	// Name of the changed field.
	//
	// Field names come from the audited record's stored representation and can differ
	// slightly from the corresponding field on the API resource — for example
	// `commission_policy_code` rather than `commission_policy`.
	Field string `json:"field" api:"required"`
	// New value as a JSON fragment.
	//
	// `null` on `delete` events, where the field has no remaining value. Encoded as a
	// JSON value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	NewValue any `json:"new_value" api:"required"`
	// Resource type identifier.
	//
	// Any of "audit_field_change".
	Object AuditFieldChangeObject `json:"object" api:"required"`
	// Previous value as a JSON fragment.
	//
	// `null` on `create` events, where the field had no prior value. Encoded as a JSON
	// value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	OldValue any `json:"old_value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		NewValue    respjson.Field
		Object      respjson.Field
		OldValue    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditFieldChange) RawJSON() string { return r.JSON.raw }
func (r *AuditFieldChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AuditFieldChangeObject string

const (
	AuditFieldChangeObjectAuditFieldChange AuditFieldChangeObject = "audit_field_change"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAuditEvent struct {
	// Resources in this page.
	Data []AuditEvent `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAuditEventObject `json:"object" api:"required"`
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
func (r ListAuditEvent) RawJSON() string { return r.JSON.raw }
func (r *ListAuditEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAuditEventObject string

const (
	ListAuditEventObjectList ListAuditEventObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAuditFieldChange struct {
	// Resources in this page.
	Data []AuditFieldChange `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAuditFieldChangeObject `json:"object" api:"required"`
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
func (r ListAuditFieldChange) RawJSON() string { return r.JSON.raw }
func (r *ListAuditFieldChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAuditFieldChangeObject string

const (
	ListAuditFieldChangeObjectList ListAuditFieldChangeObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListObjectType struct {
	// Resources in this page.
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
	// "email_inbox", "email_sender", "portal_domain", "dns_record",
	// "inventory_change_log", "invoice", "invoice_summary", "invoice_line",
	// "invoice_allocation", "invoice_for_payment", "shipment", "shipment_summary",
	// "shipment_line", "shipping_case", "shipping_case_label_url", "settlement",
	// "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	Data []string `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListObjectTypeObject `json:"object" api:"required"`
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
func (r ListObjectType) RawJSON() string { return r.JSON.raw }
func (r *ListObjectType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListObjectTypeObject string

const (
	ListObjectTypeObjectList ListObjectTypeObject = "list"
)

type CoreAuditEventGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account", "actor", "changes", "metadata", "request".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreAuditEventGetParams]'s query parameters as
// `url.Values`.
func (r CoreAuditEventGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreAuditEventListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Restricts results to audit events on or before this timestamp.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// ID of the root record whose history tree to return.
	//
	// Only applied when paired with `root_resource_type`.
	RootResourceID param.Opt[string] `query:"root_resource_id,omitzero" json:"-"`
	// Restricts results to audit events on or after this timestamp.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Filter by the mutation type recorded on the event.
	//
	// Any of "create", "update", "upsert", "delete", "restore", "archive", "approve",
	// "deny".
	Actions []string `query:"actions,omitzero" json:"-"`
	// Filter by the _acting_ account: the account that performed the mutation.
	//
	// Results are always scoped to events where your account is either the acting
	// account or the target account; this narrows that set to specific acting accounts
	// — for example a specific customer's account that mutated a resource on your
	// account.
	ActorAccountIDs []string `query:"actor_account_ids,omitzero" json:"-"`
	// Filter by the actor identifier.
	//
	// Matches the event's `actor.id`: a user ID for `user` actors, an API key ID for
	// `api_key` actors, or an agent ID for `agent` actors.
	ActorIDs []string `query:"actor_ids,omitzero" json:"-"`
	// Filter by the actor type.
	//
	// Events are recorded for actors of type `user`, `api_key`, and `agent` — the last
	// covering changes an OpenMRP agent made on your account's behalf.
	//
	// Any of "user", "api_key", "agent", "group".
	ActorTypes []string `query:"actor_types,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account", "actor", "changes", "metadata", "request".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by the audited resource IDs.
	ResourceIDs []string `query:"resource_ids,omitzero" json:"-"`
	// Filter by the resource type of the audited entity.
	//
	// The full set of valid values is available from the List Audit Event Resource
	// Types endpoint.
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
	// "email_inbox", "email_sender", "portal_domain", "dns_record",
	// "inventory_change_log", "invoice", "invoice_summary", "invoice_line",
	// "invoice_allocation", "invoice_for_payment", "shipment", "shipment_summary",
	// "shipment_line", "shipping_case", "shipping_case_label_url", "settlement",
	// "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	ResourceTypes []string `query:"resource_types,omitzero" json:"-"`
	// Scope results to a root record's entire history tree.
	//
	// Returns every event whose root resource matches, covering the root record itself
	// and all of its descendants — for example a sales order together with its lines,
	// picks, shipments, and invoices. Both `root_resource_type` and `root_resource_id`
	// must be supplied together; supplying only one has no effect.
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
	// "email_inbox", "email_sender", "portal_domain", "dns_record",
	// "inventory_change_log", "invoice", "invoice_summary", "invoice_line",
	// "invoice_allocation", "invoice_for_payment", "shipment", "shipment_summary",
	// "shipment_line", "shipping_case", "shipping_case_label_url", "settlement",
	// "settlement_summary", "role_permission", "registration_flow",
	// "registration_flow_option", "transaction", "transaction_summary",
	// "transaction_method", "transaction_type", "transaction_allocation",
	// "usage_item", "account_usage_response", "subscription_info",
	// "billing_portal_session_response", "switch_plan_response",
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
	RootResourceType CoreAuditEventListParamsRootResourceType `query:"root_resource_type,omitzero" json:"-"`
	// Filter by the _target_ account the mutation was performed against (the event's
	// `account`).
	//
	// Results are always scoped to events where your account is either the acting
	// account or the target account; this narrows that set to specific target accounts
	// — for example a specific customer's or supplier's account.
	TargetAccountIDs []string `query:"target_account_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreAuditEventListParams]'s query parameters as
// `url.Values`.
func (r CoreAuditEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Scope results to a root record's entire history tree.
//
// Returns every event whose root resource matches, covering the root record itself
// and all of its descendants — for example a sales order together with its lines,
// picks, shipments, and invoices. Both `root_resource_type` and `root_resource_id`
// must be supplied together; supplying only one has no effect.
type CoreAuditEventListParamsRootResourceType string

const (
	CoreAuditEventListParamsRootResourceTypeAccount                              CoreAuditEventListParamsRootResourceType = "account"
	CoreAuditEventListParamsRootResourceTypeActor                                CoreAuditEventListParamsRootResourceType = "actor"
	CoreAuditEventListParamsRootResourceTypeEntity                               CoreAuditEventListParamsRootResourceType = "entity"
	CoreAuditEventListParamsRootResourceTypeRecord                               CoreAuditEventListParamsRootResourceType = "record"
	CoreAuditEventListParamsRootResourceTypeFreight                              CoreAuditEventListParamsRootResourceType = "freight"
	CoreAuditEventListParamsRootResourceTypeCommitment                           CoreAuditEventListParamsRootResourceType = "commitment"
	CoreAuditEventListParamsRootResourceTypeSalesOrderTotals                     CoreAuditEventListParamsRootResourceType = "sales_order_totals"
	CoreAuditEventListParamsRootResourceTypeSalesOrderStageTotal                 CoreAuditEventListParamsRootResourceType = "sales_order_stage_total"
	CoreAuditEventListParamsRootResourceTypeSalesOrderRelated                    CoreAuditEventListParamsRootResourceType = "sales_order_related"
	CoreAuditEventListParamsRootResourceTypeOrderContact                         CoreAuditEventListParamsRootResourceType = "order_contact"
	CoreAuditEventListParamsRootResourceTypeUser                                 CoreAuditEventListParamsRootResourceType = "user"
	CoreAuditEventListParamsRootResourceTypeAddress                              CoreAuditEventListParamsRootResourceType = "address"
	CoreAuditEventListParamsRootResourceTypeAPIKey                               CoreAuditEventListParamsRootResourceType = "api_key"
	CoreAuditEventListParamsRootResourceTypeCreatedAPIKey                        CoreAuditEventListParamsRootResourceType = "created_api_key"
	CoreAuditEventListParamsRootResourceTypeRefreshToken                         CoreAuditEventListParamsRootResourceType = "refresh_token"
	CoreAuditEventListParamsRootResourceTypeList                                 CoreAuditEventListParamsRootResourceType = "list"
	CoreAuditEventListParamsRootResourceTypeSandbox                              CoreAuditEventListParamsRootResourceType = "sandbox"
	CoreAuditEventListParamsRootResourceTypeRegistrationSession                  CoreAuditEventListParamsRootResourceType = "registration_session"
	CoreAuditEventListParamsRootResourceTypePricingPlan                          CoreAuditEventListParamsRootResourceType = "pricing_plan"
	CoreAuditEventListParamsRootResourceTypeAccountPlan                          CoreAuditEventListParamsRootResourceType = "account_plan"
	CoreAuditEventListParamsRootResourceTypePlanChange                           CoreAuditEventListParamsRootResourceType = "plan_change"
	CoreAuditEventListParamsRootResourceTypeEnterpriseInquiry                    CoreAuditEventListParamsRootResourceType = "enterprise_inquiry"
	CoreAuditEventListParamsRootResourceTypeRequestLog                           CoreAuditEventListParamsRootResourceType = "request_log"
	CoreAuditEventListParamsRootResourceTypeAuditEvent                           CoreAuditEventListParamsRootResourceType = "audit_event"
	CoreAuditEventListParamsRootResourceTypeAuditFieldChange                     CoreAuditEventListParamsRootResourceType = "audit_field_change"
	CoreAuditEventListParamsRootResourceTypeRole                                 CoreAuditEventListParamsRootResourceType = "role"
	CoreAuditEventListParamsRootResourceTypeUnit                                 CoreAuditEventListParamsRootResourceType = "unit"
	CoreAuditEventListParamsRootResourceTypeAccountAffiliation                   CoreAuditEventListParamsRootResourceType = "account_affiliation"
	CoreAuditEventListParamsRootResourceTypeAgentDefinition                      CoreAuditEventListParamsRootResourceType = "agent_definition"
	CoreAuditEventListParamsRootResourceTypeAvailableTool                        CoreAuditEventListParamsRootResourceType = "available_tool"
	CoreAuditEventListParamsRootResourceTypeAgentDefinitionTool                  CoreAuditEventListParamsRootResourceType = "agent_definition_tool"
	CoreAuditEventListParamsRootResourceTypeAgentAccountStatus                   CoreAuditEventListParamsRootResourceType = "agent_account_status"
	CoreAuditEventListParamsRootResourceTypeAgentRun                             CoreAuditEventListParamsRootResourceType = "agent_run"
	CoreAuditEventListParamsRootResourceTypeAgentAction                          CoreAuditEventListParamsRootResourceType = "agent_action"
	CoreAuditEventListParamsRootResourceTypeAgentRunStep                         CoreAuditEventListParamsRootResourceType = "agent_run_step"
	CoreAuditEventListParamsRootResourceTypeAgentTokenUsage                      CoreAuditEventListParamsRootResourceType = "agent_token_usage"
	CoreAuditEventListParamsRootResourceTypeAgentMemory                          CoreAuditEventListParamsRootResourceType = "agent_memory"
	CoreAuditEventListParamsRootResourceTypeNotification                         CoreAuditEventListParamsRootResourceType = "notification"
	CoreAuditEventListParamsRootResourceTypeNotificationUnreadCount              CoreAuditEventListParamsRootResourceType = "notification_unread_count"
	CoreAuditEventListParamsRootResourceTypeNotificationSendResult               CoreAuditEventListParamsRootResourceType = "notification_send_result"
	CoreAuditEventListParamsRootResourceTypeNotificationUnreadSummary            CoreAuditEventListParamsRootResourceType = "notification_unread_summary"
	CoreAuditEventListParamsRootResourceTypeAnnouncement                         CoreAuditEventListParamsRootResourceType = "announcement"
	CoreAuditEventListParamsRootResourceTypeConversation                         CoreAuditEventListParamsRootResourceType = "conversation"
	CoreAuditEventListParamsRootResourceTypeSupportCase                          CoreAuditEventListParamsRootResourceType = "support_case"
	CoreAuditEventListParamsRootResourceTypeConversationParticipant              CoreAuditEventListParamsRootResourceType = "conversation_participant"
	CoreAuditEventListParamsRootResourceTypeReadCursor                           CoreAuditEventListParamsRootResourceType = "read_cursor"
	CoreAuditEventListParamsRootResourceTypeChatMessage                          CoreAuditEventListParamsRootResourceType = "chat_message"
	CoreAuditEventListParamsRootResourceTypeNotificationUnreadSummaryAccount     CoreAuditEventListParamsRootResourceType = "notification_unread_summary_account"
	CoreAuditEventListParamsRootResourceTypeMessagingBlock                       CoreAuditEventListParamsRootResourceType = "messaging_block"
	CoreAuditEventListParamsRootResourceTypeNotificationPreference               CoreAuditEventListParamsRootResourceType = "notification_preference"
	CoreAuditEventListParamsRootResourceTypeMessageAttachment                    CoreAuditEventListParamsRootResourceType = "message_attachment"
	CoreAuditEventListParamsRootResourceTypeAttachmentUploadTarget               CoreAuditEventListParamsRootResourceType = "attachment_upload_target"
	CoreAuditEventListParamsRootResourceTypeScheduledMessage                     CoreAuditEventListParamsRootResourceType = "scheduled_message"
	CoreAuditEventListParamsRootResourceTypeMessagingContact                     CoreAuditEventListParamsRootResourceType = "messaging_contact"
	CoreAuditEventListParamsRootResourceTypeMessageReport                        CoreAuditEventListParamsRootResourceType = "message_report"
	CoreAuditEventListParamsRootResourceTypeToolGroup                            CoreAuditEventListParamsRootResourceType = "tool_group"
	CoreAuditEventListParamsRootResourceTypeModel                                CoreAuditEventListParamsRootResourceType = "model"
	CoreAuditEventListParamsRootResourceTypePaymentTerm                          CoreAuditEventListParamsRootResourceType = "payment_term"
	CoreAuditEventListParamsRootResourceTypeShippingTerm                         CoreAuditEventListParamsRootResourceType = "shipping_term"
	CoreAuditEventListParamsRootResourceTypeQuantity                             CoreAuditEventListParamsRootResourceType = "quantity"
	CoreAuditEventListParamsRootResourceTypeAccountGroup                         CoreAuditEventListParamsRootResourceType = "account_group"
	CoreAuditEventListParamsRootResourceTypeSupportRoute                         CoreAuditEventListParamsRootResourceType = "support_route"
	CoreAuditEventListParamsRootResourceTypeSupportAvailability                  CoreAuditEventListParamsRootResourceType = "support_availability"
	CoreAuditEventListParamsRootResourceTypeAccountStatus                        CoreAuditEventListParamsRootResourceType = "account_status"
	CoreAuditEventListParamsRootResourceTypeGeolocation                          CoreAuditEventListParamsRootResourceType = "geolocation"
	CoreAuditEventListParamsRootResourceTypeAccountUser                          CoreAuditEventListParamsRootResourceType = "account_user"
	CoreAuditEventListParamsRootResourceTypeDepartment                           CoreAuditEventListParamsRootResourceType = "department"
	CoreAuditEventListParamsRootResourceTypeAccountIntegration                   CoreAuditEventListParamsRootResourceType = "account_integration"
	CoreAuditEventListParamsRootResourceTypeAccountPrice                         CoreAuditEventListParamsRootResourceType = "account_price"
	CoreAuditEventListParamsRootResourceTypeProductLine                          CoreAuditEventListParamsRootResourceType = "product_line"
	CoreAuditEventListParamsRootResourceTypeItemCategory                         CoreAuditEventListParamsRootResourceType = "item_category"
	CoreAuditEventListParamsRootResourceTypeAttribute                            CoreAuditEventListParamsRootResourceType = "attribute"
	CoreAuditEventListParamsRootResourceTypeRate                                 CoreAuditEventListParamsRootResourceType = "rate"
	CoreAuditEventListParamsRootResourceTypeAccountGroupProductLineAccess        CoreAuditEventListParamsRootResourceType = "account_group_product_line_access"
	CoreAuditEventListParamsRootResourceTypeSalesTarget                          CoreAuditEventListParamsRootResourceType = "sales_target"
	CoreAuditEventListParamsRootResourceTypeAdjustmentType                       CoreAuditEventListParamsRootResourceType = "adjustment_type"
	CoreAuditEventListParamsRootResourceTypeAccountBranding                      CoreAuditEventListParamsRootResourceType = "account_branding"
	CoreAuditEventListParamsRootResourceTypeAccountPortal                        CoreAuditEventListParamsRootResourceType = "account_portal"
	CoreAuditEventListParamsRootResourceTypeAccountLogoURL                       CoreAuditEventListParamsRootResourceType = "account_logo_url"
	CoreAuditEventListParamsRootResourceTypeAccountFaviconURL                    CoreAuditEventListParamsRootResourceType = "account_favicon_url"
	CoreAuditEventListParamsRootResourceTypePublicAccount                        CoreAuditEventListParamsRootResourceType = "public_account"
	CoreAuditEventListParamsRootResourceTypeProperty                             CoreAuditEventListParamsRootResourceType = "property"
	CoreAuditEventListParamsRootResourceTypeCarrier                              CoreAuditEventListParamsRootResourceType = "carrier"
	CoreAuditEventListParamsRootResourceTypeServiceLevel                         CoreAuditEventListParamsRootResourceType = "service_level"
	CoreAuditEventListParamsRootResourceTypeItem                                 CoreAuditEventListParamsRootResourceType = "item"
	CoreAuditEventListParamsRootResourceTypeItemLotDefault                       CoreAuditEventListParamsRootResourceType = "item_lot_default"
	CoreAuditEventListParamsRootResourceTypeItemInventory                        CoreAuditEventListParamsRootResourceType = "item_inventory"
	CoreAuditEventListParamsRootResourceTypeProduct                              CoreAuditEventListParamsRootResourceType = "product"
	CoreAuditEventListParamsRootResourceTypeBatch                                CoreAuditEventListParamsRootResourceType = "batch"
	CoreAuditEventListParamsRootResourceTypeBatchFlowNode                        CoreAuditEventListParamsRootResourceType = "batch_flow_node"
	CoreAuditEventListParamsRootResourceTypeScanningConsumption                  CoreAuditEventListParamsRootResourceType = "scanning_consumption"
	CoreAuditEventListParamsRootResourceTypeOpenBatchSummary                     CoreAuditEventListParamsRootResourceType = "open_batch_summary"
	CoreAuditEventListParamsRootResourceTypeScanningProductionStepInfo           CoreAuditEventListParamsRootResourceType = "scanning_production_step_info"
	CoreAuditEventListParamsRootResourceTypeScanningStation                      CoreAuditEventListParamsRootResourceType = "scanning_station"
	CoreAuditEventListParamsRootResourceTypeProductionStep                       CoreAuditEventListParamsRootResourceType = "production_step"
	CoreAuditEventListParamsRootResourceTypeProductionRun                        CoreAuditEventListParamsRootResourceType = "production_run"
	CoreAuditEventListParamsRootResourceTypeMachine                              CoreAuditEventListParamsRootResourceType = "machine"
	CoreAuditEventListParamsRootResourceTypeMachineStatus                        CoreAuditEventListParamsRootResourceType = "machine_status"
	CoreAuditEventListParamsRootResourceTypeMachineDowntimeEvent                 CoreAuditEventListParamsRootResourceType = "machine_downtime_event"
	CoreAuditEventListParamsRootResourceTypeDemandOverride                       CoreAuditEventListParamsRootResourceType = "demand_override"
	CoreAuditEventListParamsRootResourceTypeDemandOverrideType                   CoreAuditEventListParamsRootResourceType = "demand_override_type"
	CoreAuditEventListParamsRootResourceTypeMachineDowntimeReason                CoreAuditEventListParamsRootResourceType = "machine_downtime_reason"
	CoreAuditEventListParamsRootResourceTypeProductionSchedulePreview            CoreAuditEventListParamsRootResourceType = "production_schedule_preview"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleRegeneratePreview  CoreAuditEventListParamsRootResourceType = "production_schedule_regenerate_preview"
	CoreAuditEventListParamsRootResourceTypeProductionSchedule                   CoreAuditEventListParamsRootResourceType = "production_schedule"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleLine               CoreAuditEventListParamsRootResourceType = "production_schedule_line"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleDeviation          CoreAuditEventListParamsRootResourceType = "production_schedule_deviation"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleDerivedLine        CoreAuditEventListParamsRootResourceType = "production_schedule_derived_line"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleSettings           CoreAuditEventListParamsRootResourceType = "production_schedule_settings"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleResourceSetting    CoreAuditEventListParamsRootResourceType = "production_schedule_resource_setting"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleItemSetting        CoreAuditEventListParamsRootResourceType = "production_schedule_item_setting"
	CoreAuditEventListParamsRootResourceTypeFulfillmentRecommendation            CoreAuditEventListParamsRootResourceType = "fulfillment_recommendation"
	CoreAuditEventListParamsRootResourceTypeAnalyzeDeliveryPerformanceResponse   CoreAuditEventListParamsRootResourceType = "analyze_delivery_performance_response"
	CoreAuditEventListParamsRootResourceTypeDeliveryPerformance                  CoreAuditEventListParamsRootResourceType = "delivery_performance"
	CoreAuditEventListParamsRootResourceTypeDeliveryBacklogBucket                CoreAuditEventListParamsRootResourceType = "delivery_backlog_bucket"
	CoreAuditEventListParamsRootResourceTypeDeliveryLatenessBucket               CoreAuditEventListParamsRootResourceType = "delivery_lateness_bucket"
	CoreAuditEventListParamsRootResourceTypeDeliveryBreakdown                    CoreAuditEventListParamsRootResourceType = "delivery_breakdown"
	CoreAuditEventListParamsRootResourceTypeAnalyzeSalesBreakdownResponse        CoreAuditEventListParamsRootResourceType = "analyze_sales_breakdown_response"
	CoreAuditEventListParamsRootResourceTypeSalesTotals                          CoreAuditEventListParamsRootResourceType = "sales_totals"
	CoreAuditEventListParamsRootResourceTypeSalesBreakdown                       CoreAuditEventListParamsRootResourceType = "sales_breakdown"
	CoreAuditEventListParamsRootResourceTypeScheduleOrderCoverage                CoreAuditEventListParamsRootResourceType = "schedule_order_coverage"
	CoreAuditEventListParamsRootResourceTypeScheduleOrderCoverageLine            CoreAuditEventListParamsRootResourceType = "schedule_order_coverage_line"
	CoreAuditEventListParamsRootResourceTypeScheduleDeviationType                CoreAuditEventListParamsRootResourceType = "schedule_deviation_type"
	CoreAuditEventListParamsRootResourceTypeScheduleAtRiskOrder                  CoreAuditEventListParamsRootResourceType = "schedule_at_risk_order"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleFinishedPolicy     CoreAuditEventListParamsRootResourceType = "production_schedule_finished_policy"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleFinishingLine      CoreAuditEventListParamsRootResourceType = "production_schedule_finishing_line"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleWeekRelease        CoreAuditEventListParamsRootResourceType = "production_schedule_week_release"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleWeekReleasePreview CoreAuditEventListParamsRootResourceType = "production_schedule_week_release_preview"
	CoreAuditEventListParamsRootResourceTypeProductionScheduleItemPolicy         CoreAuditEventListParamsRootResourceType = "production_schedule_item_policy"
	CoreAuditEventListParamsRootResourceTypeChildAccount                         CoreAuditEventListParamsRootResourceType = "child_account"
	CoreAuditEventListParamsRootResourceTypeUnitGroup                            CoreAuditEventListParamsRootResourceType = "unit_group"
	CoreAuditEventListParamsRootResourceTypeUnitGroupUnit                        CoreAuditEventListParamsRootResourceType = "unit_group_unit"
	CoreAuditEventListParamsRootResourceTypeConsumption                          CoreAuditEventListParamsRootResourceType = "consumption"
	CoreAuditEventListParamsRootResourceTypeCustomerProductLineAccess            CoreAuditEventListParamsRootResourceType = "customer_product_line_access"
	CoreAuditEventListParamsRootResourceTypeCustomer                             CoreAuditEventListParamsRootResourceType = "customer"
	CoreAuditEventListParamsRootResourceTypeFrequentlyOrderedProduct             CoreAuditEventListParamsRootResourceType = "frequently_ordered_product"
	CoreAuditEventListParamsRootResourceTypePriority                             CoreAuditEventListParamsRootResourceType = "priority"
	CoreAuditEventListParamsRootResourceTypeDelivery                             CoreAuditEventListParamsRootResourceType = "delivery"
	CoreAuditEventListParamsRootResourceTypeDeliveryLine                         CoreAuditEventListParamsRootResourceType = "delivery_line"
	CoreAuditEventListParamsRootResourceTypeSalesOrder                           CoreAuditEventListParamsRootResourceType = "sales_order"
	CoreAuditEventListParamsRootResourceTypeLocation                             CoreAuditEventListParamsRootResourceType = "location"
	CoreAuditEventListParamsRootResourceTypeLocationType                         CoreAuditEventListParamsRootResourceType = "location_type"
	CoreAuditEventListParamsRootResourceTypeLot                                  CoreAuditEventListParamsRootResourceType = "lot"
	CoreAuditEventListParamsRootResourceTypeEmailLog                             CoreAuditEventListParamsRootResourceType = "email_log"
	CoreAuditEventListParamsRootResourceTypeEmailDomain                          CoreAuditEventListParamsRootResourceType = "email_domain"
	CoreAuditEventListParamsRootResourceTypeEmailInbox                           CoreAuditEventListParamsRootResourceType = "email_inbox"
	CoreAuditEventListParamsRootResourceTypeEmailSender                          CoreAuditEventListParamsRootResourceType = "email_sender"
	CoreAuditEventListParamsRootResourceTypePortalDomain                         CoreAuditEventListParamsRootResourceType = "portal_domain"
	CoreAuditEventListParamsRootResourceTypeDNSRecord                            CoreAuditEventListParamsRootResourceType = "dns_record"
	CoreAuditEventListParamsRootResourceTypeInventoryChangeLog                   CoreAuditEventListParamsRootResourceType = "inventory_change_log"
	CoreAuditEventListParamsRootResourceTypeInvoice                              CoreAuditEventListParamsRootResourceType = "invoice"
	CoreAuditEventListParamsRootResourceTypeInvoiceSummary                       CoreAuditEventListParamsRootResourceType = "invoice_summary"
	CoreAuditEventListParamsRootResourceTypeInvoiceLine                          CoreAuditEventListParamsRootResourceType = "invoice_line"
	CoreAuditEventListParamsRootResourceTypeInvoiceAllocation                    CoreAuditEventListParamsRootResourceType = "invoice_allocation"
	CoreAuditEventListParamsRootResourceTypeInvoiceForPayment                    CoreAuditEventListParamsRootResourceType = "invoice_for_payment"
	CoreAuditEventListParamsRootResourceTypeShipment                             CoreAuditEventListParamsRootResourceType = "shipment"
	CoreAuditEventListParamsRootResourceTypeShipmentSummary                      CoreAuditEventListParamsRootResourceType = "shipment_summary"
	CoreAuditEventListParamsRootResourceTypeShipmentLine                         CoreAuditEventListParamsRootResourceType = "shipment_line"
	CoreAuditEventListParamsRootResourceTypeShippingCase                         CoreAuditEventListParamsRootResourceType = "shipping_case"
	CoreAuditEventListParamsRootResourceTypeShippingCaseLabelURL                 CoreAuditEventListParamsRootResourceType = "shipping_case_label_url"
	CoreAuditEventListParamsRootResourceTypeSettlement                           CoreAuditEventListParamsRootResourceType = "settlement"
	CoreAuditEventListParamsRootResourceTypeSettlementSummary                    CoreAuditEventListParamsRootResourceType = "settlement_summary"
	CoreAuditEventListParamsRootResourceTypeRolePermission                       CoreAuditEventListParamsRootResourceType = "role_permission"
	CoreAuditEventListParamsRootResourceTypeRegistrationFlow                     CoreAuditEventListParamsRootResourceType = "registration_flow"
	CoreAuditEventListParamsRootResourceTypeRegistrationFlowOption               CoreAuditEventListParamsRootResourceType = "registration_flow_option"
	CoreAuditEventListParamsRootResourceTypeTransaction                          CoreAuditEventListParamsRootResourceType = "transaction"
	CoreAuditEventListParamsRootResourceTypeTransactionSummary                   CoreAuditEventListParamsRootResourceType = "transaction_summary"
	CoreAuditEventListParamsRootResourceTypeTransactionMethod                    CoreAuditEventListParamsRootResourceType = "transaction_method"
	CoreAuditEventListParamsRootResourceTypeTransactionType                      CoreAuditEventListParamsRootResourceType = "transaction_type"
	CoreAuditEventListParamsRootResourceTypeTransactionAllocation                CoreAuditEventListParamsRootResourceType = "transaction_allocation"
	CoreAuditEventListParamsRootResourceTypeUsageItem                            CoreAuditEventListParamsRootResourceType = "usage_item"
	CoreAuditEventListParamsRootResourceTypeAccountUsageResponse                 CoreAuditEventListParamsRootResourceType = "account_usage_response"
	CoreAuditEventListParamsRootResourceTypeSubscriptionInfo                     CoreAuditEventListParamsRootResourceType = "subscription_info"
	CoreAuditEventListParamsRootResourceTypeBillingPortalSessionResponse         CoreAuditEventListParamsRootResourceType = "billing_portal_session_response"
	CoreAuditEventListParamsRootResourceTypeSwitchPlanResponse                   CoreAuditEventListParamsRootResourceType = "switch_plan_response"
	CoreAuditEventListParamsRootResourceTypeEnsureBillingCustomerResponse        CoreAuditEventListParamsRootResourceType = "ensure_billing_customer_response"
	CoreAuditEventListParamsRootResourceTypeSpendingCapResponse                  CoreAuditEventListParamsRootResourceType = "spending_cap_response"
	CoreAuditEventListParamsRootResourceTypeAgentSpendInfo                       CoreAuditEventListParamsRootResourceType = "agent_spend_info"
	CoreAuditEventListParamsRootResourceTypeWebhookResponse                      CoreAuditEventListParamsRootResourceType = "webhook_response"
	CoreAuditEventListParamsRootResourceTypeAddressSuggestion                    CoreAuditEventListParamsRootResourceType = "address_suggestion"
	CoreAuditEventListParamsRootResourceTypeAddressComponents                    CoreAuditEventListParamsRootResourceType = "address_components"
	CoreAuditEventListParamsRootResourceTypeAddressDetailsResult                 CoreAuditEventListParamsRootResourceType = "address_details_result"
	CoreAuditEventListParamsRootResourceTypeValidatedAddress                     CoreAuditEventListParamsRootResourceType = "validated_address"
	CoreAuditEventListParamsRootResourceTypePlanLimit                            CoreAuditEventListParamsRootResourceType = "plan_limit"
	CoreAuditEventListParamsRootResourceTypePlanChangeProration                  CoreAuditEventListParamsRootResourceType = "plan_change_proration"
	CoreAuditEventListParamsRootResourceTypePlanChangeLineItem                   CoreAuditEventListParamsRootResourceType = "plan_change_line_item"
	CoreAuditEventListParamsRootResourceTypeSetupBillingResponse                 CoreAuditEventListParamsRootResourceType = "setup_billing_response"
	CoreAuditEventListParamsRootResourceTypeConfirmPaymentResponse               CoreAuditEventListParamsRootResourceType = "confirm_payment_response"
	CoreAuditEventListParamsRootResourceTypeOAuthResponse                        CoreAuditEventListParamsRootResourceType = "oauth_response"
	CoreAuditEventListParamsRootResourceTypeOAuthStatusResponse                  CoreAuditEventListParamsRootResourceType = "oauth_status_response"
	CoreAuditEventListParamsRootResourceTypeStripePublishableKey                 CoreAuditEventListParamsRootResourceType = "stripe_publishable_key"
	CoreAuditEventListParamsRootResourceTypeStripeStatus                         CoreAuditEventListParamsRootResourceType = "stripe_status"
	CoreAuditEventListParamsRootResourceTypeHealthcheck                          CoreAuditEventListParamsRootResourceType = "healthcheck"
	CoreAuditEventListParamsRootResourceTypeAgentDefinitionConfig                CoreAuditEventListParamsRootResourceType = "agent_definition_config"
	CoreAuditEventListParamsRootResourceTypeTriggerConfig                        CoreAuditEventListParamsRootResourceType = "trigger_config"
	CoreAuditEventListParamsRootResourceTypeCustomerContactInfo                  CoreAuditEventListParamsRootResourceType = "customer_contact_info"
	CoreAuditEventListParamsRootResourceTypeCustomerFreightPreferences           CoreAuditEventListParamsRootResourceType = "customer_freight_preferences"
	CoreAuditEventListParamsRootResourceTypeCustomerDefaults                     CoreAuditEventListParamsRootResourceType = "customer_defaults"
	CoreAuditEventListParamsRootResourceTypeCustomerLeadTime                     CoreAuditEventListParamsRootResourceType = "customer_lead_time"
	CoreAuditEventListParamsRootResourceTypeCustomerNotificationPreferences      CoreAuditEventListParamsRootResourceType = "customer_notification_preferences"
	CoreAuditEventListParamsRootResourceTypeOrderNotificationRecipient           CoreAuditEventListParamsRootResourceType = "order_notification_recipient"
	CoreAuditEventListParamsRootResourceTypeOrderDiscount                        CoreAuditEventListParamsRootResourceType = "order_discount"
	CoreAuditEventListParamsRootResourceTypeSalesOrderLine                       CoreAuditEventListParamsRootResourceType = "sales_order_line"
	CoreAuditEventListParamsRootResourceTypeSalesOrderType                       CoreAuditEventListParamsRootResourceType = "sales_order_type"
	CoreAuditEventListParamsRootResourceTypeSalesOrderStatus                     CoreAuditEventListParamsRootResourceType = "sales_order_status"
	CoreAuditEventListParamsRootResourceTypeMaterial                             CoreAuditEventListParamsRootResourceType = "material"
	CoreAuditEventListParamsRootResourceTypeSupplierMaterial                     CoreAuditEventListParamsRootResourceType = "supplier_material"
	CoreAuditEventListParamsRootResourceTypePart                                 CoreAuditEventListParamsRootResourceType = "part"
	CoreAuditEventListParamsRootResourceTypePermissionGroup                      CoreAuditEventListParamsRootResourceType = "permission_group"
	CoreAuditEventListParamsRootResourceTypePermission                           CoreAuditEventListParamsRootResourceType = "permission"
	CoreAuditEventListParamsRootResourceTypePick                                 CoreAuditEventListParamsRootResourceType = "pick"
	CoreAuditEventListParamsRootResourceTypePickLine                             CoreAuditEventListParamsRootResourceType = "pick_line"
	CoreAuditEventListParamsRootResourceTypeProductType                          CoreAuditEventListParamsRootResourceType = "product_type"
	CoreAuditEventListParamsRootResourceTypeProduction                           CoreAuditEventListParamsRootResourceType = "production"
	CoreAuditEventListParamsRootResourceTypeProductionFlow                       CoreAuditEventListParamsRootResourceType = "production_flow"
	CoreAuditEventListParamsRootResourceTypeMap                                  CoreAuditEventListParamsRootResourceType = "map"
	CoreAuditEventListParamsRootResourceTypePurchaseOrder                        CoreAuditEventListParamsRootResourceType = "purchase_order"
	CoreAuditEventListParamsRootResourceTypePurchaseOrderLine                    CoreAuditEventListParamsRootResourceType = "purchase_order_line"
	CoreAuditEventListParamsRootResourceTypeSupplier                             CoreAuditEventListParamsRootResourceType = "supplier"
	CoreAuditEventListParamsRootResourceTypeSupplierSummary                      CoreAuditEventListParamsRootResourceType = "supplier_summary"
	CoreAuditEventListParamsRootResourceTypeReceivableEntry                      CoreAuditEventListParamsRootResourceType = "receivable_entry"
	CoreAuditEventListParamsRootResourceTypeReceivingOrder                       CoreAuditEventListParamsRootResourceType = "receiving_order"
	CoreAuditEventListParamsRootResourceTypeReceivingOrderLine                   CoreAuditEventListParamsRootResourceType = "receiving_order_line"
	CoreAuditEventListParamsRootResourceTypeEmailContact                         CoreAuditEventListParamsRootResourceType = "email_contact"
	CoreAuditEventListParamsRootResourceTypeAllocationEntry                      CoreAuditEventListParamsRootResourceType = "allocation_entry"
	CoreAuditEventListParamsRootResourceTypeOpenCreditEntry                      CoreAuditEventListParamsRootResourceType = "open_credit_entry"
	CoreAuditEventListParamsRootResourceTypeVolumeDiscount                       CoreAuditEventListParamsRootResourceType = "volume_discount"
	CoreAuditEventListParamsRootResourceTypeVolumeDiscountTier                   CoreAuditEventListParamsRootResourceType = "volume_discount_tier"
	CoreAuditEventListParamsRootResourceTypeAnalyzeDeliveriesResponse            CoreAuditEventListParamsRootResourceType = "analyze_deliveries_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeManufacturingResponse         CoreAuditEventListParamsRootResourceType = "analyze_manufacturing_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeManufacturingBatchResponse    CoreAuditEventListParamsRootResourceType = "analyze_manufacturing_batch_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeQuarterlyOrdersResponse       CoreAuditEventListParamsRootResourceType = "analyze_quarterly_orders_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeNewCustomersResponse          CoreAuditEventListParamsRootResourceType = "analyze_new_customers_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeDemandForecastResponse        CoreAuditEventListParamsRootResourceType = "analyze_demand_forecast_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeOeeResponse                   CoreAuditEventListParamsRootResourceType = "analyze_oee_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeOeeTrendResponse              CoreAuditEventListParamsRootResourceType = "analyze_oee_trend_response"
	CoreAuditEventListParamsRootResourceTypeAnalyzeScheduleAttainmentResponse    CoreAuditEventListParamsRootResourceType = "analyze_schedule_attainment_response"
	CoreAuditEventListParamsRootResourceTypeCatalogProductLine                   CoreAuditEventListParamsRootResourceType = "catalog_product_line"
	CoreAuditEventListParamsRootResourceTypeCatalogCategory                      CoreAuditEventListParamsRootResourceType = "catalog_category"
	CoreAuditEventListParamsRootResourceTypeCatalogProduct                       CoreAuditEventListParamsRootResourceType = "catalog_product"
	CoreAuditEventListParamsRootResourceTypeCatalogProperty                      CoreAuditEventListParamsRootResourceType = "catalog_property"
	CoreAuditEventListParamsRootResourceTypeCatalogAttribute                     CoreAuditEventListParamsRootResourceType = "catalog_attribute"
	CoreAuditEventListParamsRootResourceTypeDcLocation                           CoreAuditEventListParamsRootResourceType = "dc_location"
	CoreAuditEventListParamsRootResourceTypeEdiRun                               CoreAuditEventListParamsRootResourceType = "edi_run"
	CoreAuditEventListParamsRootResourceTypeInventoryItem                        CoreAuditEventListParamsRootResourceType = "inventory_item"
	CoreAuditEventListParamsRootResourceTypeAnalyzeWeeksOfSalesResponse          CoreAuditEventListParamsRootResourceType = "analyze_weeks_of_sales_response"
	CoreAuditEventListParamsRootResourceTypeBulkReconcileItemsResponse           CoreAuditEventListParamsRootResourceType = "bulk_reconcile_items_response"
	CoreAuditEventListParamsRootResourceTypeSysProperty                          CoreAuditEventListParamsRootResourceType = "sys_property"
	CoreAuditEventListParamsRootResourceTypeSysPropertyType                      CoreAuditEventListParamsRootResourceType = "sys_property_type"
	CoreAuditEventListParamsRootResourceTypeSysPropertyValue                     CoreAuditEventListParamsRootResourceType = "sys_property_value"
	CoreAuditEventListParamsRootResourceTypeTerritory                            CoreAuditEventListParamsRootResourceType = "territory"
	CoreAuditEventListParamsRootResourceTypeTenancy                              CoreAuditEventListParamsRootResourceType = "tenancy"
	CoreAuditEventListParamsRootResourceTypeCheckoutSession                      CoreAuditEventListParamsRootResourceType = "checkout_session"
	CoreAuditEventListParamsRootResourceTypeEstimateRateResult                   CoreAuditEventListParamsRootResourceType = "estimate_rate_result"
	CoreAuditEventListParamsRootResourceTypeRateShopOption                       CoreAuditEventListParamsRootResourceType = "rate_shop_option"
	CoreAuditEventListParamsRootResourceTypeRateShopResult                       CoreAuditEventListParamsRootResourceType = "rate_shop_result"
	CoreAuditEventListParamsRootResourceTypeOwner                                CoreAuditEventListParamsRootResourceType = "owner"
	CoreAuditEventListParamsRootResourceTypeCreatedBy                            CoreAuditEventListParamsRootResourceType = "created_by"
	CoreAuditEventListParamsRootResourceTypeMessage                              CoreAuditEventListParamsRootResourceType = "message"
	CoreAuditEventListParamsRootResourceTypeAccountPhotoUploadResult             CoreAuditEventListParamsRootResourceType = "account_photo_upload_result"
	CoreAuditEventListParamsRootResourceTypeUserPhotoUploadResult                CoreAuditEventListParamsRootResourceType = "user_photo_upload_result"
	CoreAuditEventListParamsRootResourceTypeUserPhotoURL                         CoreAuditEventListParamsRootResourceType = "user_photo_url"
	CoreAuditEventListParamsRootResourceTypeBatchLot                             CoreAuditEventListParamsRootResourceType = "batch_lot"
	CoreAuditEventListParamsRootResourceTypeCheckDuplicateResult                 CoreAuditEventListParamsRootResourceType = "check_duplicate_result"
	CoreAuditEventListParamsRootResourceTypeItemTrendPoint                       CoreAuditEventListParamsRootResourceType = "item_trend_point"
	CoreAuditEventListParamsRootResourceTypeTenancyPendingRegistration           CoreAuditEventListParamsRootResourceType = "tenancy_pending_registration"
	CoreAuditEventListParamsRootResourceTypeInvoiceAllocationEntry               CoreAuditEventListParamsRootResourceType = "invoice_allocation_entry"
	CoreAuditEventListParamsRootResourceTypeAllocationCustomer                   CoreAuditEventListParamsRootResourceType = "allocation_customer"
	CoreAuditEventListParamsRootResourceTypeCheckoutSalesOrder                   CoreAuditEventListParamsRootResourceType = "checkout_sales_order"
	CoreAuditEventListParamsRootResourceTypeSalesOrderPriceQuote                 CoreAuditEventListParamsRootResourceType = "sales_order_price_quote"
	CoreAuditEventListParamsRootResourceTypeSalesOrderFreightQuote               CoreAuditEventListParamsRootResourceType = "sales_order_freight_quote"
	CoreAuditEventListParamsRootResourceTypeSalesOrderCommitmentQuote            CoreAuditEventListParamsRootResourceType = "sales_order_commitment_quote"
	CoreAuditEventListParamsRootResourceTypeOperatingCalendar                    CoreAuditEventListParamsRootResourceType = "operating_calendar"
	CoreAuditEventListParamsRootResourceTypeOperatingCalendarClosure             CoreAuditEventListParamsRootResourceType = "operating_calendar_closure"
	CoreAuditEventListParamsRootResourceTypeSalesOrderPriceQuoteLine             CoreAuditEventListParamsRootResourceType = "sales_order_price_quote_line"
	CoreAuditEventListParamsRootResourceTypeHubspotSyncJob                       CoreAuditEventListParamsRootResourceType = "hubspot_sync_job"
	CoreAuditEventListParamsRootResourceTypeHubspotSyncReport                    CoreAuditEventListParamsRootResourceType = "hubspot_sync_report"
	CoreAuditEventListParamsRootResourceTypeHubspotCompanyReview                 CoreAuditEventListParamsRootResourceType = "hubspot_company_review"
	CoreAuditEventListParamsRootResourceTypeHubspotCompanyCandidate              CoreAuditEventListParamsRootResourceType = "hubspot_company_candidate"
	CoreAuditEventListParamsRootResourceTypeHubspotSyncRecord                    CoreAuditEventListParamsRootResourceType = "hubspot_sync_record"
	CoreAuditEventListParamsRootResourceTypeContactMatch                         CoreAuditEventListParamsRootResourceType = "contact_match"
	CoreAuditEventListParamsRootResourceTypeReplyDraft                           CoreAuditEventListParamsRootResourceType = "reply_draft"
	CoreAuditEventListParamsRootResourceTypeConversationLink                     CoreAuditEventListParamsRootResourceType = "conversation_link"
	CoreAuditEventListParamsRootResourceTypeMessagingGroup                       CoreAuditEventListParamsRootResourceType = "messaging_group"
	CoreAuditEventListParamsRootResourceTypeMessagingGroupMember                 CoreAuditEventListParamsRootResourceType = "messaging_group_member"
	CoreAuditEventListParamsRootResourceTypePortalProfile                        CoreAuditEventListParamsRootResourceType = "portal_profile"
	CoreAuditEventListParamsRootResourceTypePortalRegistrationSession            CoreAuditEventListParamsRootResourceType = "portal_registration_session"
	CoreAuditEventListParamsRootResourceTypePortalRegistrationSessionData        CoreAuditEventListParamsRootResourceType = "portal_registration_session_data"
	CoreAuditEventListParamsRootResourceTypePackList                             CoreAuditEventListParamsRootResourceType = "pack_list"
	CoreAuditEventListParamsRootResourceTypePackListParty                        CoreAuditEventListParamsRootResourceType = "pack_list_party"
	CoreAuditEventListParamsRootResourceTypePackListLineItem                     CoreAuditEventListParamsRootResourceType = "pack_list_line_item"
	CoreAuditEventListParamsRootResourceTypePackListBackOrder                    CoreAuditEventListParamsRootResourceType = "pack_list_back_order"
	CoreAuditEventListParamsRootResourceTypePackListCase                         CoreAuditEventListParamsRootResourceType = "pack_list_case"
	CoreAuditEventListParamsRootResourceTypeJob                                  CoreAuditEventListParamsRootResourceType = "job"
	CoreAuditEventListParamsRootResourceTypeJobResult                            CoreAuditEventListParamsRootResourceType = "job_result"
	CoreAuditEventListParamsRootResourceTypeJobExport                            CoreAuditEventListParamsRootResourceType = "job_export"
	CoreAuditEventListParamsRootResourceTypeAnalyzeCustomerPricingResponse       CoreAuditEventListParamsRootResourceType = "analyze_customer_pricing_response"
	CoreAuditEventListParamsRootResourceTypeCustomerPricingFinding               CoreAuditEventListParamsRootResourceType = "customer_pricing_finding"
	CoreAuditEventListParamsRootResourceTypeCustomerPricingSummary               CoreAuditEventListParamsRootResourceType = "customer_pricing_summary"
	CoreAuditEventListParamsRootResourceTypeComputedRate                         CoreAuditEventListParamsRootResourceType = "computed_rate"
	CoreAuditEventListParamsRootResourceTypeComputedQuantity                     CoreAuditEventListParamsRootResourceType = "computed_quantity"
	CoreAuditEventListParamsRootResourceTypeAnalyzeRealizedMarginsResponse       CoreAuditEventListParamsRootResourceType = "analyze_realized_margins_response"
	CoreAuditEventListParamsRootResourceTypeRealizedMarginFinding                CoreAuditEventListParamsRootResourceType = "realized_margin_finding"
	CoreAuditEventListParamsRootResourceTypeRealizedMarginSummary                CoreAuditEventListParamsRootResourceType = "realized_margin_summary"
	CoreAuditEventListParamsRootResourceTypeShipmentRelated                      CoreAuditEventListParamsRootResourceType = "shipment_related"
	CoreAuditEventListParamsRootResourceTypeInvoiceRelated                       CoreAuditEventListParamsRootResourceType = "invoice_related"
	CoreAuditEventListParamsRootResourceTypePickRelated                          CoreAuditEventListParamsRootResourceType = "pick_related"
	CoreAuditEventListParamsRootResourceTypePickTotals                           CoreAuditEventListParamsRootResourceType = "pick_totals"
	CoreAuditEventListParamsRootResourceTypePickStageTotal                       CoreAuditEventListParamsRootResourceType = "pick_stage_total"
)
