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
// MessagingConversationLinkService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingConversationLinkService] method instead.
type MessagingConversationLinkService struct {
	options []option.RequestOption
}

// NewMessagingConversationLinkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingConversationLinkService(opts ...option.RequestOption) (r MessagingConversationLinkService) {
	r = MessagingConversationLinkService{}
	r.options = opts
	return
}

// Links a business record to a conversation, in addition to whatever topic the
// conversation is anchored to.
//
// A conversation can link any number of records, and each linked record surfaces
// the conversation when conversations are listed for that record.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationLinkService) New(ctx context.Context, id string, params MessagingConversationLinkNewParams, opts ...option.RequestOption) (res *ConversationLink, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/links", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the business records linked to a conversation.
//
// Every link is returned in one page. The conversation's primary `topic` anchor is
// not a link and is not listed here.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingConversationLinkService) List(ctx context.Context, id string, query MessagingConversationLinkListParams, opts ...option.RequestOption) (res *ListConversationLink, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/links", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Removes a business-record link from a conversation.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingConversationLinkService) Delete(ctx context.Context, linkID string, body MessagingConversationLinkDeleteParams, opts ...option.RequestOption) (res *MessagingConversationLinkDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if linkID == "" {
		err = errors.New("missing required link_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/conversations/%s/links/%s", url.PathEscape(body.ID), url.PathEscape(linkID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to link a business record to a conversation.
//
// The properties ResourceID, ResourceType are required.
type AddConversationLinkRequestParam struct {
	// The id of the business record to link.
	ResourceID string `json:"resource_id" api:"required"`
	// The kind of business record to link.
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
	ResourceType AddConversationLinkRequestResourceType `json:"resource_type,omitzero" api:"required"`
	paramObj
}

func (r AddConversationLinkRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow AddConversationLinkRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddConversationLinkRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of business record to link.
type AddConversationLinkRequestResourceType string

const (
	AddConversationLinkRequestResourceTypeAccount                              AddConversationLinkRequestResourceType = "account"
	AddConversationLinkRequestResourceTypeActor                                AddConversationLinkRequestResourceType = "actor"
	AddConversationLinkRequestResourceTypeEntity                               AddConversationLinkRequestResourceType = "entity"
	AddConversationLinkRequestResourceTypeRecord                               AddConversationLinkRequestResourceType = "record"
	AddConversationLinkRequestResourceTypeFreight                              AddConversationLinkRequestResourceType = "freight"
	AddConversationLinkRequestResourceTypeCommitment                           AddConversationLinkRequestResourceType = "commitment"
	AddConversationLinkRequestResourceTypeSalesOrderTotals                     AddConversationLinkRequestResourceType = "sales_order_totals"
	AddConversationLinkRequestResourceTypeSalesOrderStageTotal                 AddConversationLinkRequestResourceType = "sales_order_stage_total"
	AddConversationLinkRequestResourceTypeSalesOrderRelated                    AddConversationLinkRequestResourceType = "sales_order_related"
	AddConversationLinkRequestResourceTypeOrderContact                         AddConversationLinkRequestResourceType = "order_contact"
	AddConversationLinkRequestResourceTypeUser                                 AddConversationLinkRequestResourceType = "user"
	AddConversationLinkRequestResourceTypeAddress                              AddConversationLinkRequestResourceType = "address"
	AddConversationLinkRequestResourceTypeAPIKey                               AddConversationLinkRequestResourceType = "api_key"
	AddConversationLinkRequestResourceTypeCreatedAPIKey                        AddConversationLinkRequestResourceType = "created_api_key"
	AddConversationLinkRequestResourceTypeRefreshToken                         AddConversationLinkRequestResourceType = "refresh_token"
	AddConversationLinkRequestResourceTypeList                                 AddConversationLinkRequestResourceType = "list"
	AddConversationLinkRequestResourceTypeSandbox                              AddConversationLinkRequestResourceType = "sandbox"
	AddConversationLinkRequestResourceTypeRegistrationSession                  AddConversationLinkRequestResourceType = "registration_session"
	AddConversationLinkRequestResourceTypePricingPlan                          AddConversationLinkRequestResourceType = "pricing_plan"
	AddConversationLinkRequestResourceTypeAccountPlan                          AddConversationLinkRequestResourceType = "account_plan"
	AddConversationLinkRequestResourceTypePlanChange                           AddConversationLinkRequestResourceType = "plan_change"
	AddConversationLinkRequestResourceTypeEnterpriseInquiry                    AddConversationLinkRequestResourceType = "enterprise_inquiry"
	AddConversationLinkRequestResourceTypeRequestLog                           AddConversationLinkRequestResourceType = "request_log"
	AddConversationLinkRequestResourceTypeAuditEvent                           AddConversationLinkRequestResourceType = "audit_event"
	AddConversationLinkRequestResourceTypeAuditFieldChange                     AddConversationLinkRequestResourceType = "audit_field_change"
	AddConversationLinkRequestResourceTypeRole                                 AddConversationLinkRequestResourceType = "role"
	AddConversationLinkRequestResourceTypeUnit                                 AddConversationLinkRequestResourceType = "unit"
	AddConversationLinkRequestResourceTypeAccountAffiliation                   AddConversationLinkRequestResourceType = "account_affiliation"
	AddConversationLinkRequestResourceTypeAgentDefinition                      AddConversationLinkRequestResourceType = "agent_definition"
	AddConversationLinkRequestResourceTypeAvailableTool                        AddConversationLinkRequestResourceType = "available_tool"
	AddConversationLinkRequestResourceTypeAgentDefinitionTool                  AddConversationLinkRequestResourceType = "agent_definition_tool"
	AddConversationLinkRequestResourceTypeAgentAccountStatus                   AddConversationLinkRequestResourceType = "agent_account_status"
	AddConversationLinkRequestResourceTypeAgentRun                             AddConversationLinkRequestResourceType = "agent_run"
	AddConversationLinkRequestResourceTypeAgentAction                          AddConversationLinkRequestResourceType = "agent_action"
	AddConversationLinkRequestResourceTypeAgentRunStep                         AddConversationLinkRequestResourceType = "agent_run_step"
	AddConversationLinkRequestResourceTypeAgentTokenUsage                      AddConversationLinkRequestResourceType = "agent_token_usage"
	AddConversationLinkRequestResourceTypeAgentMemory                          AddConversationLinkRequestResourceType = "agent_memory"
	AddConversationLinkRequestResourceTypeNotification                         AddConversationLinkRequestResourceType = "notification"
	AddConversationLinkRequestResourceTypeNotificationUnreadCount              AddConversationLinkRequestResourceType = "notification_unread_count"
	AddConversationLinkRequestResourceTypeNotificationSendResult               AddConversationLinkRequestResourceType = "notification_send_result"
	AddConversationLinkRequestResourceTypeNotificationUnreadSummary            AddConversationLinkRequestResourceType = "notification_unread_summary"
	AddConversationLinkRequestResourceTypeAnnouncement                         AddConversationLinkRequestResourceType = "announcement"
	AddConversationLinkRequestResourceTypeConversation                         AddConversationLinkRequestResourceType = "conversation"
	AddConversationLinkRequestResourceTypeSupportCase                          AddConversationLinkRequestResourceType = "support_case"
	AddConversationLinkRequestResourceTypeConversationParticipant              AddConversationLinkRequestResourceType = "conversation_participant"
	AddConversationLinkRequestResourceTypeReadCursor                           AddConversationLinkRequestResourceType = "read_cursor"
	AddConversationLinkRequestResourceTypeChatMessage                          AddConversationLinkRequestResourceType = "chat_message"
	AddConversationLinkRequestResourceTypeNotificationUnreadSummaryAccount     AddConversationLinkRequestResourceType = "notification_unread_summary_account"
	AddConversationLinkRequestResourceTypeMessagingBlock                       AddConversationLinkRequestResourceType = "messaging_block"
	AddConversationLinkRequestResourceTypeNotificationPreference               AddConversationLinkRequestResourceType = "notification_preference"
	AddConversationLinkRequestResourceTypeMessageAttachment                    AddConversationLinkRequestResourceType = "message_attachment"
	AddConversationLinkRequestResourceTypeAttachmentUploadTarget               AddConversationLinkRequestResourceType = "attachment_upload_target"
	AddConversationLinkRequestResourceTypeScheduledMessage                     AddConversationLinkRequestResourceType = "scheduled_message"
	AddConversationLinkRequestResourceTypeMessagingContact                     AddConversationLinkRequestResourceType = "messaging_contact"
	AddConversationLinkRequestResourceTypeMessageReport                        AddConversationLinkRequestResourceType = "message_report"
	AddConversationLinkRequestResourceTypeToolGroup                            AddConversationLinkRequestResourceType = "tool_group"
	AddConversationLinkRequestResourceTypeModel                                AddConversationLinkRequestResourceType = "model"
	AddConversationLinkRequestResourceTypePaymentTerm                          AddConversationLinkRequestResourceType = "payment_term"
	AddConversationLinkRequestResourceTypeShippingTerm                         AddConversationLinkRequestResourceType = "shipping_term"
	AddConversationLinkRequestResourceTypeQuantity                             AddConversationLinkRequestResourceType = "quantity"
	AddConversationLinkRequestResourceTypeAccountGroup                         AddConversationLinkRequestResourceType = "account_group"
	AddConversationLinkRequestResourceTypeSupportRoute                         AddConversationLinkRequestResourceType = "support_route"
	AddConversationLinkRequestResourceTypeSupportAvailability                  AddConversationLinkRequestResourceType = "support_availability"
	AddConversationLinkRequestResourceTypeAccountStatus                        AddConversationLinkRequestResourceType = "account_status"
	AddConversationLinkRequestResourceTypeGeolocation                          AddConversationLinkRequestResourceType = "geolocation"
	AddConversationLinkRequestResourceTypeAccountUser                          AddConversationLinkRequestResourceType = "account_user"
	AddConversationLinkRequestResourceTypeDepartment                           AddConversationLinkRequestResourceType = "department"
	AddConversationLinkRequestResourceTypeAccountIntegration                   AddConversationLinkRequestResourceType = "account_integration"
	AddConversationLinkRequestResourceTypeAccountPrice                         AddConversationLinkRequestResourceType = "account_price"
	AddConversationLinkRequestResourceTypeProductLine                          AddConversationLinkRequestResourceType = "product_line"
	AddConversationLinkRequestResourceTypeItemCategory                         AddConversationLinkRequestResourceType = "item_category"
	AddConversationLinkRequestResourceTypeAttribute                            AddConversationLinkRequestResourceType = "attribute"
	AddConversationLinkRequestResourceTypeRate                                 AddConversationLinkRequestResourceType = "rate"
	AddConversationLinkRequestResourceTypeAccountGroupProductLineAccess        AddConversationLinkRequestResourceType = "account_group_product_line_access"
	AddConversationLinkRequestResourceTypeSalesTarget                          AddConversationLinkRequestResourceType = "sales_target"
	AddConversationLinkRequestResourceTypeAdjustmentType                       AddConversationLinkRequestResourceType = "adjustment_type"
	AddConversationLinkRequestResourceTypeAccountBranding                      AddConversationLinkRequestResourceType = "account_branding"
	AddConversationLinkRequestResourceTypeAccountPortal                        AddConversationLinkRequestResourceType = "account_portal"
	AddConversationLinkRequestResourceTypeAccountLogoURL                       AddConversationLinkRequestResourceType = "account_logo_url"
	AddConversationLinkRequestResourceTypeAccountFaviconURL                    AddConversationLinkRequestResourceType = "account_favicon_url"
	AddConversationLinkRequestResourceTypePublicAccount                        AddConversationLinkRequestResourceType = "public_account"
	AddConversationLinkRequestResourceTypeProperty                             AddConversationLinkRequestResourceType = "property"
	AddConversationLinkRequestResourceTypeCarrier                              AddConversationLinkRequestResourceType = "carrier"
	AddConversationLinkRequestResourceTypeServiceLevel                         AddConversationLinkRequestResourceType = "service_level"
	AddConversationLinkRequestResourceTypeItem                                 AddConversationLinkRequestResourceType = "item"
	AddConversationLinkRequestResourceTypeItemLotDefault                       AddConversationLinkRequestResourceType = "item_lot_default"
	AddConversationLinkRequestResourceTypeItemInventory                        AddConversationLinkRequestResourceType = "item_inventory"
	AddConversationLinkRequestResourceTypeProduct                              AddConversationLinkRequestResourceType = "product"
	AddConversationLinkRequestResourceTypeBatch                                AddConversationLinkRequestResourceType = "batch"
	AddConversationLinkRequestResourceTypeBatchFlowNode                        AddConversationLinkRequestResourceType = "batch_flow_node"
	AddConversationLinkRequestResourceTypeScanningConsumption                  AddConversationLinkRequestResourceType = "scanning_consumption"
	AddConversationLinkRequestResourceTypeOpenBatchSummary                     AddConversationLinkRequestResourceType = "open_batch_summary"
	AddConversationLinkRequestResourceTypeScanningProductionStepInfo           AddConversationLinkRequestResourceType = "scanning_production_step_info"
	AddConversationLinkRequestResourceTypeScanningStation                      AddConversationLinkRequestResourceType = "scanning_station"
	AddConversationLinkRequestResourceTypeProductionStep                       AddConversationLinkRequestResourceType = "production_step"
	AddConversationLinkRequestResourceTypeProductionRun                        AddConversationLinkRequestResourceType = "production_run"
	AddConversationLinkRequestResourceTypeMachine                              AddConversationLinkRequestResourceType = "machine"
	AddConversationLinkRequestResourceTypeMachineStatus                        AddConversationLinkRequestResourceType = "machine_status"
	AddConversationLinkRequestResourceTypeMachineDowntimeEvent                 AddConversationLinkRequestResourceType = "machine_downtime_event"
	AddConversationLinkRequestResourceTypeDemandOverride                       AddConversationLinkRequestResourceType = "demand_override"
	AddConversationLinkRequestResourceTypeDemandOverrideType                   AddConversationLinkRequestResourceType = "demand_override_type"
	AddConversationLinkRequestResourceTypeMachineDowntimeReason                AddConversationLinkRequestResourceType = "machine_downtime_reason"
	AddConversationLinkRequestResourceTypeProductionSchedulePreview            AddConversationLinkRequestResourceType = "production_schedule_preview"
	AddConversationLinkRequestResourceTypeProductionScheduleRegeneratePreview  AddConversationLinkRequestResourceType = "production_schedule_regenerate_preview"
	AddConversationLinkRequestResourceTypeProductionSchedule                   AddConversationLinkRequestResourceType = "production_schedule"
	AddConversationLinkRequestResourceTypeProductionScheduleLine               AddConversationLinkRequestResourceType = "production_schedule_line"
	AddConversationLinkRequestResourceTypeProductionScheduleDeviation          AddConversationLinkRequestResourceType = "production_schedule_deviation"
	AddConversationLinkRequestResourceTypeProductionScheduleDerivedLine        AddConversationLinkRequestResourceType = "production_schedule_derived_line"
	AddConversationLinkRequestResourceTypeProductionScheduleSettings           AddConversationLinkRequestResourceType = "production_schedule_settings"
	AddConversationLinkRequestResourceTypeProductionScheduleResourceSetting    AddConversationLinkRequestResourceType = "production_schedule_resource_setting"
	AddConversationLinkRequestResourceTypeProductionScheduleItemSetting        AddConversationLinkRequestResourceType = "production_schedule_item_setting"
	AddConversationLinkRequestResourceTypeFulfillmentRecommendation            AddConversationLinkRequestResourceType = "fulfillment_recommendation"
	AddConversationLinkRequestResourceTypeAnalyzeDeliveryPerformanceResponse   AddConversationLinkRequestResourceType = "analyze_delivery_performance_response"
	AddConversationLinkRequestResourceTypeDeliveryPerformance                  AddConversationLinkRequestResourceType = "delivery_performance"
	AddConversationLinkRequestResourceTypeDeliveryBacklogBucket                AddConversationLinkRequestResourceType = "delivery_backlog_bucket"
	AddConversationLinkRequestResourceTypeDeliveryLatenessBucket               AddConversationLinkRequestResourceType = "delivery_lateness_bucket"
	AddConversationLinkRequestResourceTypeDeliveryBreakdown                    AddConversationLinkRequestResourceType = "delivery_breakdown"
	AddConversationLinkRequestResourceTypeAnalyzeSalesBreakdownResponse        AddConversationLinkRequestResourceType = "analyze_sales_breakdown_response"
	AddConversationLinkRequestResourceTypeSalesTotals                          AddConversationLinkRequestResourceType = "sales_totals"
	AddConversationLinkRequestResourceTypeSalesBreakdown                       AddConversationLinkRequestResourceType = "sales_breakdown"
	AddConversationLinkRequestResourceTypeScheduleOrderCoverage                AddConversationLinkRequestResourceType = "schedule_order_coverage"
	AddConversationLinkRequestResourceTypeScheduleOrderCoverageLine            AddConversationLinkRequestResourceType = "schedule_order_coverage_line"
	AddConversationLinkRequestResourceTypeScheduleDeviationType                AddConversationLinkRequestResourceType = "schedule_deviation_type"
	AddConversationLinkRequestResourceTypeScheduleAtRiskOrder                  AddConversationLinkRequestResourceType = "schedule_at_risk_order"
	AddConversationLinkRequestResourceTypeProductionScheduleFinishedPolicy     AddConversationLinkRequestResourceType = "production_schedule_finished_policy"
	AddConversationLinkRequestResourceTypeProductionScheduleFinishingLine      AddConversationLinkRequestResourceType = "production_schedule_finishing_line"
	AddConversationLinkRequestResourceTypeProductionScheduleWeekRelease        AddConversationLinkRequestResourceType = "production_schedule_week_release"
	AddConversationLinkRequestResourceTypeProductionScheduleWeekReleasePreview AddConversationLinkRequestResourceType = "production_schedule_week_release_preview"
	AddConversationLinkRequestResourceTypeProductionScheduleItemPolicy         AddConversationLinkRequestResourceType = "production_schedule_item_policy"
	AddConversationLinkRequestResourceTypeChildAccount                         AddConversationLinkRequestResourceType = "child_account"
	AddConversationLinkRequestResourceTypeUnitGroup                            AddConversationLinkRequestResourceType = "unit_group"
	AddConversationLinkRequestResourceTypeUnitGroupUnit                        AddConversationLinkRequestResourceType = "unit_group_unit"
	AddConversationLinkRequestResourceTypeConsumption                          AddConversationLinkRequestResourceType = "consumption"
	AddConversationLinkRequestResourceTypeCustomerProductLineAccess            AddConversationLinkRequestResourceType = "customer_product_line_access"
	AddConversationLinkRequestResourceTypeCustomer                             AddConversationLinkRequestResourceType = "customer"
	AddConversationLinkRequestResourceTypeFrequentlyOrderedProduct             AddConversationLinkRequestResourceType = "frequently_ordered_product"
	AddConversationLinkRequestResourceTypePriority                             AddConversationLinkRequestResourceType = "priority"
	AddConversationLinkRequestResourceTypeDelivery                             AddConversationLinkRequestResourceType = "delivery"
	AddConversationLinkRequestResourceTypeDeliveryLine                         AddConversationLinkRequestResourceType = "delivery_line"
	AddConversationLinkRequestResourceTypeSalesOrder                           AddConversationLinkRequestResourceType = "sales_order"
	AddConversationLinkRequestResourceTypeLocation                             AddConversationLinkRequestResourceType = "location"
	AddConversationLinkRequestResourceTypeLocationType                         AddConversationLinkRequestResourceType = "location_type"
	AddConversationLinkRequestResourceTypeLot                                  AddConversationLinkRequestResourceType = "lot"
	AddConversationLinkRequestResourceTypeEmailLog                             AddConversationLinkRequestResourceType = "email_log"
	AddConversationLinkRequestResourceTypeEmailDomain                          AddConversationLinkRequestResourceType = "email_domain"
	AddConversationLinkRequestResourceTypeEmailInbox                           AddConversationLinkRequestResourceType = "email_inbox"
	AddConversationLinkRequestResourceTypeEmailSender                          AddConversationLinkRequestResourceType = "email_sender"
	AddConversationLinkRequestResourceTypePortalDomain                         AddConversationLinkRequestResourceType = "portal_domain"
	AddConversationLinkRequestResourceTypeDNSRecord                            AddConversationLinkRequestResourceType = "dns_record"
	AddConversationLinkRequestResourceTypeInventoryChangeLog                   AddConversationLinkRequestResourceType = "inventory_change_log"
	AddConversationLinkRequestResourceTypeInvoice                              AddConversationLinkRequestResourceType = "invoice"
	AddConversationLinkRequestResourceTypeInvoiceSummary                       AddConversationLinkRequestResourceType = "invoice_summary"
	AddConversationLinkRequestResourceTypeInvoiceLine                          AddConversationLinkRequestResourceType = "invoice_line"
	AddConversationLinkRequestResourceTypeInvoiceAllocation                    AddConversationLinkRequestResourceType = "invoice_allocation"
	AddConversationLinkRequestResourceTypeInvoiceForPayment                    AddConversationLinkRequestResourceType = "invoice_for_payment"
	AddConversationLinkRequestResourceTypeShipment                             AddConversationLinkRequestResourceType = "shipment"
	AddConversationLinkRequestResourceTypeShipmentSummary                      AddConversationLinkRequestResourceType = "shipment_summary"
	AddConversationLinkRequestResourceTypeShipmentLine                         AddConversationLinkRequestResourceType = "shipment_line"
	AddConversationLinkRequestResourceTypeShippingCase                         AddConversationLinkRequestResourceType = "shipping_case"
	AddConversationLinkRequestResourceTypeShippingCaseLabelURL                 AddConversationLinkRequestResourceType = "shipping_case_label_url"
	AddConversationLinkRequestResourceTypeSettlement                           AddConversationLinkRequestResourceType = "settlement"
	AddConversationLinkRequestResourceTypeSettlementSummary                    AddConversationLinkRequestResourceType = "settlement_summary"
	AddConversationLinkRequestResourceTypeRolePermission                       AddConversationLinkRequestResourceType = "role_permission"
	AddConversationLinkRequestResourceTypeRegistrationFlow                     AddConversationLinkRequestResourceType = "registration_flow"
	AddConversationLinkRequestResourceTypeRegistrationFlowOption               AddConversationLinkRequestResourceType = "registration_flow_option"
	AddConversationLinkRequestResourceTypeTransaction                          AddConversationLinkRequestResourceType = "transaction"
	AddConversationLinkRequestResourceTypeTransactionSummary                   AddConversationLinkRequestResourceType = "transaction_summary"
	AddConversationLinkRequestResourceTypeTransactionMethod                    AddConversationLinkRequestResourceType = "transaction_method"
	AddConversationLinkRequestResourceTypeTransactionType                      AddConversationLinkRequestResourceType = "transaction_type"
	AddConversationLinkRequestResourceTypeTransactionAllocation                AddConversationLinkRequestResourceType = "transaction_allocation"
	AddConversationLinkRequestResourceTypeUsageItem                            AddConversationLinkRequestResourceType = "usage_item"
	AddConversationLinkRequestResourceTypeAccountUsageResponse                 AddConversationLinkRequestResourceType = "account_usage_response"
	AddConversationLinkRequestResourceTypeSubscriptionInfo                     AddConversationLinkRequestResourceType = "subscription_info"
	AddConversationLinkRequestResourceTypeBillingPortalSessionResponse         AddConversationLinkRequestResourceType = "billing_portal_session_response"
	AddConversationLinkRequestResourceTypeSwitchPlanResponse                   AddConversationLinkRequestResourceType = "switch_plan_response"
	AddConversationLinkRequestResourceTypeEnsureBillingCustomerResponse        AddConversationLinkRequestResourceType = "ensure_billing_customer_response"
	AddConversationLinkRequestResourceTypeSpendingCapResponse                  AddConversationLinkRequestResourceType = "spending_cap_response"
	AddConversationLinkRequestResourceTypeAgentSpendInfo                       AddConversationLinkRequestResourceType = "agent_spend_info"
	AddConversationLinkRequestResourceTypeWebhookResponse                      AddConversationLinkRequestResourceType = "webhook_response"
	AddConversationLinkRequestResourceTypeAddressSuggestion                    AddConversationLinkRequestResourceType = "address_suggestion"
	AddConversationLinkRequestResourceTypeAddressComponents                    AddConversationLinkRequestResourceType = "address_components"
	AddConversationLinkRequestResourceTypeAddressDetailsResult                 AddConversationLinkRequestResourceType = "address_details_result"
	AddConversationLinkRequestResourceTypeValidatedAddress                     AddConversationLinkRequestResourceType = "validated_address"
	AddConversationLinkRequestResourceTypePlanLimit                            AddConversationLinkRequestResourceType = "plan_limit"
	AddConversationLinkRequestResourceTypePlanChangeProration                  AddConversationLinkRequestResourceType = "plan_change_proration"
	AddConversationLinkRequestResourceTypePlanChangeLineItem                   AddConversationLinkRequestResourceType = "plan_change_line_item"
	AddConversationLinkRequestResourceTypeSetupBillingResponse                 AddConversationLinkRequestResourceType = "setup_billing_response"
	AddConversationLinkRequestResourceTypeConfirmPaymentResponse               AddConversationLinkRequestResourceType = "confirm_payment_response"
	AddConversationLinkRequestResourceTypeOAuthResponse                        AddConversationLinkRequestResourceType = "oauth_response"
	AddConversationLinkRequestResourceTypeOAuthStatusResponse                  AddConversationLinkRequestResourceType = "oauth_status_response"
	AddConversationLinkRequestResourceTypeStripePublishableKey                 AddConversationLinkRequestResourceType = "stripe_publishable_key"
	AddConversationLinkRequestResourceTypeStripeStatus                         AddConversationLinkRequestResourceType = "stripe_status"
	AddConversationLinkRequestResourceTypeHealthcheck                          AddConversationLinkRequestResourceType = "healthcheck"
	AddConversationLinkRequestResourceTypeAgentDefinitionConfig                AddConversationLinkRequestResourceType = "agent_definition_config"
	AddConversationLinkRequestResourceTypeTriggerConfig                        AddConversationLinkRequestResourceType = "trigger_config"
	AddConversationLinkRequestResourceTypeCustomerContactInfo                  AddConversationLinkRequestResourceType = "customer_contact_info"
	AddConversationLinkRequestResourceTypeCustomerFreightPreferences           AddConversationLinkRequestResourceType = "customer_freight_preferences"
	AddConversationLinkRequestResourceTypeCustomerDefaults                     AddConversationLinkRequestResourceType = "customer_defaults"
	AddConversationLinkRequestResourceTypeCustomerLeadTime                     AddConversationLinkRequestResourceType = "customer_lead_time"
	AddConversationLinkRequestResourceTypeCustomerNotificationPreferences      AddConversationLinkRequestResourceType = "customer_notification_preferences"
	AddConversationLinkRequestResourceTypeOrderNotificationRecipient           AddConversationLinkRequestResourceType = "order_notification_recipient"
	AddConversationLinkRequestResourceTypeOrderDiscount                        AddConversationLinkRequestResourceType = "order_discount"
	AddConversationLinkRequestResourceTypeSalesOrderLine                       AddConversationLinkRequestResourceType = "sales_order_line"
	AddConversationLinkRequestResourceTypeSalesOrderType                       AddConversationLinkRequestResourceType = "sales_order_type"
	AddConversationLinkRequestResourceTypeSalesOrderStatus                     AddConversationLinkRequestResourceType = "sales_order_status"
	AddConversationLinkRequestResourceTypeMaterial                             AddConversationLinkRequestResourceType = "material"
	AddConversationLinkRequestResourceTypeSupplierMaterial                     AddConversationLinkRequestResourceType = "supplier_material"
	AddConversationLinkRequestResourceTypePart                                 AddConversationLinkRequestResourceType = "part"
	AddConversationLinkRequestResourceTypePermissionGroup                      AddConversationLinkRequestResourceType = "permission_group"
	AddConversationLinkRequestResourceTypePermission                           AddConversationLinkRequestResourceType = "permission"
	AddConversationLinkRequestResourceTypePick                                 AddConversationLinkRequestResourceType = "pick"
	AddConversationLinkRequestResourceTypePickLine                             AddConversationLinkRequestResourceType = "pick_line"
	AddConversationLinkRequestResourceTypeProductType                          AddConversationLinkRequestResourceType = "product_type"
	AddConversationLinkRequestResourceTypeProduction                           AddConversationLinkRequestResourceType = "production"
	AddConversationLinkRequestResourceTypeProductionFlow                       AddConversationLinkRequestResourceType = "production_flow"
	AddConversationLinkRequestResourceTypeMap                                  AddConversationLinkRequestResourceType = "map"
	AddConversationLinkRequestResourceTypePurchaseOrder                        AddConversationLinkRequestResourceType = "purchase_order"
	AddConversationLinkRequestResourceTypePurchaseOrderLine                    AddConversationLinkRequestResourceType = "purchase_order_line"
	AddConversationLinkRequestResourceTypeSupplier                             AddConversationLinkRequestResourceType = "supplier"
	AddConversationLinkRequestResourceTypeSupplierSummary                      AddConversationLinkRequestResourceType = "supplier_summary"
	AddConversationLinkRequestResourceTypeReceivableEntry                      AddConversationLinkRequestResourceType = "receivable_entry"
	AddConversationLinkRequestResourceTypeReceivingOrder                       AddConversationLinkRequestResourceType = "receiving_order"
	AddConversationLinkRequestResourceTypeReceivingOrderLine                   AddConversationLinkRequestResourceType = "receiving_order_line"
	AddConversationLinkRequestResourceTypeEmailContact                         AddConversationLinkRequestResourceType = "email_contact"
	AddConversationLinkRequestResourceTypeAllocationEntry                      AddConversationLinkRequestResourceType = "allocation_entry"
	AddConversationLinkRequestResourceTypeOpenCreditEntry                      AddConversationLinkRequestResourceType = "open_credit_entry"
	AddConversationLinkRequestResourceTypeVolumeDiscount                       AddConversationLinkRequestResourceType = "volume_discount"
	AddConversationLinkRequestResourceTypeVolumeDiscountTier                   AddConversationLinkRequestResourceType = "volume_discount_tier"
	AddConversationLinkRequestResourceTypeAnalyzeDeliveriesResponse            AddConversationLinkRequestResourceType = "analyze_deliveries_response"
	AddConversationLinkRequestResourceTypeAnalyzeManufacturingResponse         AddConversationLinkRequestResourceType = "analyze_manufacturing_response"
	AddConversationLinkRequestResourceTypeAnalyzeManufacturingBatchResponse    AddConversationLinkRequestResourceType = "analyze_manufacturing_batch_response"
	AddConversationLinkRequestResourceTypeAnalyzeQuarterlyOrdersResponse       AddConversationLinkRequestResourceType = "analyze_quarterly_orders_response"
	AddConversationLinkRequestResourceTypeAnalyzeNewCustomersResponse          AddConversationLinkRequestResourceType = "analyze_new_customers_response"
	AddConversationLinkRequestResourceTypeAnalyzeDemandForecastResponse        AddConversationLinkRequestResourceType = "analyze_demand_forecast_response"
	AddConversationLinkRequestResourceTypeAnalyzeOeeResponse                   AddConversationLinkRequestResourceType = "analyze_oee_response"
	AddConversationLinkRequestResourceTypeAnalyzeOeeTrendResponse              AddConversationLinkRequestResourceType = "analyze_oee_trend_response"
	AddConversationLinkRequestResourceTypeAnalyzeScheduleAttainmentResponse    AddConversationLinkRequestResourceType = "analyze_schedule_attainment_response"
	AddConversationLinkRequestResourceTypeCatalogProductLine                   AddConversationLinkRequestResourceType = "catalog_product_line"
	AddConversationLinkRequestResourceTypeCatalogCategory                      AddConversationLinkRequestResourceType = "catalog_category"
	AddConversationLinkRequestResourceTypeCatalogProduct                       AddConversationLinkRequestResourceType = "catalog_product"
	AddConversationLinkRequestResourceTypeCatalogProperty                      AddConversationLinkRequestResourceType = "catalog_property"
	AddConversationLinkRequestResourceTypeCatalogAttribute                     AddConversationLinkRequestResourceType = "catalog_attribute"
	AddConversationLinkRequestResourceTypeDcLocation                           AddConversationLinkRequestResourceType = "dc_location"
	AddConversationLinkRequestResourceTypeEdiRun                               AddConversationLinkRequestResourceType = "edi_run"
	AddConversationLinkRequestResourceTypeInventoryItem                        AddConversationLinkRequestResourceType = "inventory_item"
	AddConversationLinkRequestResourceTypeAnalyzeWeeksOfSalesResponse          AddConversationLinkRequestResourceType = "analyze_weeks_of_sales_response"
	AddConversationLinkRequestResourceTypeBulkReconcileItemsResponse           AddConversationLinkRequestResourceType = "bulk_reconcile_items_response"
	AddConversationLinkRequestResourceTypeSysProperty                          AddConversationLinkRequestResourceType = "sys_property"
	AddConversationLinkRequestResourceTypeSysPropertyType                      AddConversationLinkRequestResourceType = "sys_property_type"
	AddConversationLinkRequestResourceTypeSysPropertyValue                     AddConversationLinkRequestResourceType = "sys_property_value"
	AddConversationLinkRequestResourceTypeTerritory                            AddConversationLinkRequestResourceType = "territory"
	AddConversationLinkRequestResourceTypeTenancy                              AddConversationLinkRequestResourceType = "tenancy"
	AddConversationLinkRequestResourceTypeCheckoutSession                      AddConversationLinkRequestResourceType = "checkout_session"
	AddConversationLinkRequestResourceTypeEstimateRateResult                   AddConversationLinkRequestResourceType = "estimate_rate_result"
	AddConversationLinkRequestResourceTypeRateShopOption                       AddConversationLinkRequestResourceType = "rate_shop_option"
	AddConversationLinkRequestResourceTypeRateShopResult                       AddConversationLinkRequestResourceType = "rate_shop_result"
	AddConversationLinkRequestResourceTypeOwner                                AddConversationLinkRequestResourceType = "owner"
	AddConversationLinkRequestResourceTypeCreatedBy                            AddConversationLinkRequestResourceType = "created_by"
	AddConversationLinkRequestResourceTypeMessage                              AddConversationLinkRequestResourceType = "message"
	AddConversationLinkRequestResourceTypeAccountPhotoUploadResult             AddConversationLinkRequestResourceType = "account_photo_upload_result"
	AddConversationLinkRequestResourceTypeUserPhotoUploadResult                AddConversationLinkRequestResourceType = "user_photo_upload_result"
	AddConversationLinkRequestResourceTypeUserPhotoURL                         AddConversationLinkRequestResourceType = "user_photo_url"
	AddConversationLinkRequestResourceTypeBatchLot                             AddConversationLinkRequestResourceType = "batch_lot"
	AddConversationLinkRequestResourceTypeCheckDuplicateResult                 AddConversationLinkRequestResourceType = "check_duplicate_result"
	AddConversationLinkRequestResourceTypeItemTrendPoint                       AddConversationLinkRequestResourceType = "item_trend_point"
	AddConversationLinkRequestResourceTypeTenancyPendingRegistration           AddConversationLinkRequestResourceType = "tenancy_pending_registration"
	AddConversationLinkRequestResourceTypeInvoiceAllocationEntry               AddConversationLinkRequestResourceType = "invoice_allocation_entry"
	AddConversationLinkRequestResourceTypeAllocationCustomer                   AddConversationLinkRequestResourceType = "allocation_customer"
	AddConversationLinkRequestResourceTypeCheckoutSalesOrder                   AddConversationLinkRequestResourceType = "checkout_sales_order"
	AddConversationLinkRequestResourceTypeSalesOrderPriceQuote                 AddConversationLinkRequestResourceType = "sales_order_price_quote"
	AddConversationLinkRequestResourceTypeSalesOrderFreightQuote               AddConversationLinkRequestResourceType = "sales_order_freight_quote"
	AddConversationLinkRequestResourceTypeSalesOrderCommitmentQuote            AddConversationLinkRequestResourceType = "sales_order_commitment_quote"
	AddConversationLinkRequestResourceTypeOperatingCalendar                    AddConversationLinkRequestResourceType = "operating_calendar"
	AddConversationLinkRequestResourceTypeOperatingCalendarClosure             AddConversationLinkRequestResourceType = "operating_calendar_closure"
	AddConversationLinkRequestResourceTypeSalesOrderPriceQuoteLine             AddConversationLinkRequestResourceType = "sales_order_price_quote_line"
	AddConversationLinkRequestResourceTypeHubspotSyncJob                       AddConversationLinkRequestResourceType = "hubspot_sync_job"
	AddConversationLinkRequestResourceTypeHubspotSyncReport                    AddConversationLinkRequestResourceType = "hubspot_sync_report"
	AddConversationLinkRequestResourceTypeHubspotCompanyReview                 AddConversationLinkRequestResourceType = "hubspot_company_review"
	AddConversationLinkRequestResourceTypeHubspotCompanyCandidate              AddConversationLinkRequestResourceType = "hubspot_company_candidate"
	AddConversationLinkRequestResourceTypeHubspotSyncRecord                    AddConversationLinkRequestResourceType = "hubspot_sync_record"
	AddConversationLinkRequestResourceTypeContactMatch                         AddConversationLinkRequestResourceType = "contact_match"
	AddConversationLinkRequestResourceTypeReplyDraft                           AddConversationLinkRequestResourceType = "reply_draft"
	AddConversationLinkRequestResourceTypeConversationLink                     AddConversationLinkRequestResourceType = "conversation_link"
	AddConversationLinkRequestResourceTypeMessagingGroup                       AddConversationLinkRequestResourceType = "messaging_group"
	AddConversationLinkRequestResourceTypeMessagingGroupMember                 AddConversationLinkRequestResourceType = "messaging_group_member"
	AddConversationLinkRequestResourceTypePortalProfile                        AddConversationLinkRequestResourceType = "portal_profile"
	AddConversationLinkRequestResourceTypePortalRegistrationSession            AddConversationLinkRequestResourceType = "portal_registration_session"
	AddConversationLinkRequestResourceTypePortalRegistrationSessionData        AddConversationLinkRequestResourceType = "portal_registration_session_data"
	AddConversationLinkRequestResourceTypePackList                             AddConversationLinkRequestResourceType = "pack_list"
	AddConversationLinkRequestResourceTypePackListParty                        AddConversationLinkRequestResourceType = "pack_list_party"
	AddConversationLinkRequestResourceTypePackListLineItem                     AddConversationLinkRequestResourceType = "pack_list_line_item"
	AddConversationLinkRequestResourceTypePackListBackOrder                    AddConversationLinkRequestResourceType = "pack_list_back_order"
	AddConversationLinkRequestResourceTypePackListCase                         AddConversationLinkRequestResourceType = "pack_list_case"
	AddConversationLinkRequestResourceTypeJob                                  AddConversationLinkRequestResourceType = "job"
	AddConversationLinkRequestResourceTypeJobResult                            AddConversationLinkRequestResourceType = "job_result"
	AddConversationLinkRequestResourceTypeJobExport                            AddConversationLinkRequestResourceType = "job_export"
	AddConversationLinkRequestResourceTypeAnalyzeCustomerPricingResponse       AddConversationLinkRequestResourceType = "analyze_customer_pricing_response"
	AddConversationLinkRequestResourceTypeCustomerPricingFinding               AddConversationLinkRequestResourceType = "customer_pricing_finding"
	AddConversationLinkRequestResourceTypeCustomerPricingSummary               AddConversationLinkRequestResourceType = "customer_pricing_summary"
	AddConversationLinkRequestResourceTypeComputedRate                         AddConversationLinkRequestResourceType = "computed_rate"
	AddConversationLinkRequestResourceTypeComputedQuantity                     AddConversationLinkRequestResourceType = "computed_quantity"
	AddConversationLinkRequestResourceTypeAnalyzeRealizedMarginsResponse       AddConversationLinkRequestResourceType = "analyze_realized_margins_response"
	AddConversationLinkRequestResourceTypeRealizedMarginFinding                AddConversationLinkRequestResourceType = "realized_margin_finding"
	AddConversationLinkRequestResourceTypeRealizedMarginSummary                AddConversationLinkRequestResourceType = "realized_margin_summary"
	AddConversationLinkRequestResourceTypeShipmentRelated                      AddConversationLinkRequestResourceType = "shipment_related"
	AddConversationLinkRequestResourceTypeInvoiceRelated                       AddConversationLinkRequestResourceType = "invoice_related"
	AddConversationLinkRequestResourceTypePickRelated                          AddConversationLinkRequestResourceType = "pick_related"
	AddConversationLinkRequestResourceTypePickTotals                           AddConversationLinkRequestResourceType = "pick_totals"
	AddConversationLinkRequestResourceTypePickStageTotal                       AddConversationLinkRequestResourceType = "pick_stage_total"
)

// A reference from a conversation to a business record it concerns, such as an
// order, invoice, shipment, or customer.
//
// Links sit alongside the conversation's primary `topic` anchor, so one thread can
// reference several records. Listing conversations by business record matches the
// topic anchor and these links alike, which is what surfaces a conversation on the
// record's own page.
type ConversationLink struct {
	// Conversation link ID.
	ID string `json:"id" api:"required"`
	// A conversation thread the caller participates in.
	Conversation Conversation `json:"conversation" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "conversation_link".
	Object ConversationLinkObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Conversation respjson.Field
		CreatedAt    respjson.Field
		Object       respjson.Field
		Resource     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationLink) RawJSON() string { return r.JSON.raw }
func (r *ConversationLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ConversationLinkObject string

const (
	ConversationLinkObjectConversationLink ConversationLinkObject = "conversation_link"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListConversationLink struct {
	// Resources in this page.
	Data []ConversationLink `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConversationLinkObject `json:"object" api:"required"`
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
func (r ListConversationLink) RawJSON() string { return r.JSON.raw }
func (r *ListConversationLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConversationLinkObject string

const (
	ListConversationLinkObjectList ListConversationLinkObject = "list"
)

type MessagingConversationLinkDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingConversationLinkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingConversationLinkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingConversationLinkNewParams struct {
	// Request to link a business record to a conversation.
	AddConversationLinkRequest AddConversationLinkRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "conversation", "conversation.assignee", "conversation.group",
	// "conversation.participants", "conversation.topic", "conversation.last_message",
	// "conversation.last_message.sender", "conversation.last_message.author",
	// "conversation.last_message.resource", "conversation.last_message.attachments",
	// "conversation.last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingConversationLinkNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AddConversationLinkRequest)
}
func (r *MessagingConversationLinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingConversationLinkNewParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationLinkNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationLinkListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "conversation", "conversation.assignee", "conversation.group",
	// "conversation.participants", "conversation.topic", "conversation.last_message",
	// "conversation.last_message.sender", "conversation.last_message.author",
	// "conversation.last_message.resource", "conversation.last_message.attachments",
	// "conversation.last_message.attachments.resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingConversationLinkListParams]'s query parameters as
// `url.Values`.
func (r MessagingConversationLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingConversationLinkDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
