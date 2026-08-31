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
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// View the jobs that track asynchronous work. Endpoints that answer 202 Accepted
// raise one and point at it with a Location header.
//
// CoreJobService contains methods and other services that help with interacting
// with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreJobService] method instead.
type CoreJobService struct {
	options []option.RequestOption
}

// NewCoreJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCoreJobService(opts ...option.RequestOption) (r CoreJobService) {
	r = CoreJobService{}
	r.options = opts
	return
}

// Returns a job by ID — poll the job named in a `202 Accepted` response's
// `Location` to observe its outcome. A completed export carries the link to its
// file on `export.url`.
//
// This endpoint requires the permissions: `jobs:read`, `customers:read`,
// `suppliers:read`.
func (r *CoreJobService) Get(ctx context.Context, id string, query CoreJobGetParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/jobs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancels a job and returns it carrying its `cancelled` status. Work in flight is
// not interrupted but can no longer settle, and a finished job cannot be
// cancelled.
//
// This endpoint requires the permission: `jobs:delete`.
func (r *CoreJobService) Cancel(ctx context.Context, id string, body CoreJobCancelParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/jobs/%s/cancel", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Records a piece of work the API accepted and carries out asynchronously.
// Endpoints answering `202 Accepted` point at one with a `Location` header; poll
// it for the outcome.
type Job struct {
	// Job ID.
	ID string `json:"id" api:"required"`
	// When the job was cancelled.
	CancelledAt time.Time `json:"cancelled_at" api:"required" format:"date-time"`
	// When the job finished processing, whether or not every row succeeded.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// When the job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	CreatedBy Actor `json:"created_by" api:"required"`
	// ResponseError is the JSON-serializable error body returned to API clients. It
	// contains only public information. This struct is used by the OpenAPI schema
	// generator to produce documentation.
	Error ResponseError `json:"error" api:"required"`
	// Points a completed export job at the file it produced.
	Export JobExport `json:"export" api:"required"`
	// When the most recent attempt failed. A retry that succeeds leaves this alongside
	// `completed_at`.
	FailedAt time.Time `json:"failed_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "job".
	Object JobObject `json:"object" api:"required"`
	// The kind of resource the job operates on, as an object-type value (e.g.
	// `product`).
	//
	// `type` names the verb — what the job does — and this names the subject, so a job
	// that produced no results still says what it was for.
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
	ResourceType JobResourceType `json:"resource_type" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Results ListJobResult `json:"results" api:"required"`
	// When the job began executing.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// How far the job has got.
	//
	// `completed` means the work was processed, not that every row succeeded — read
	// each entry's own `status` in `results`.
	//
	// Any of "created", "started", "completed", "failed", "cancelled".
	Status JobStatus `json:"status" api:"required"`
	// The kind of work the job carries out.
	//
	// Any of "bulk_create", "bulk_upsert", "export", "pack_pick".
	Type JobType `json:"type" api:"required"`
	// When the job was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CancelledAt  respjson.Field
		CompletedAt  respjson.Field
		CreatedAt    respjson.Field
		CreatedBy    respjson.Field
		Error        respjson.Field
		Export       respjson.Field
		FailedAt     respjson.Field
		Object       respjson.Field
		ResourceType respjson.Field
		Results      respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type JobObject string

const (
	JobObjectJob JobObject = "job"
)

// The kind of resource the job operates on, as an object-type value (e.g.
// `product`).
//
// `type` names the verb — what the job does — and this names the subject, so a job
// that produced no results still says what it was for.
type JobResourceType string

const (
	JobResourceTypeAccount                              JobResourceType = "account"
	JobResourceTypeActor                                JobResourceType = "actor"
	JobResourceTypeEntity                               JobResourceType = "entity"
	JobResourceTypeRecord                               JobResourceType = "record"
	JobResourceTypeFreight                              JobResourceType = "freight"
	JobResourceTypeCommitment                           JobResourceType = "commitment"
	JobResourceTypeSalesOrderTotals                     JobResourceType = "sales_order_totals"
	JobResourceTypeSalesOrderStageTotal                 JobResourceType = "sales_order_stage_total"
	JobResourceTypeSalesOrderRelated                    JobResourceType = "sales_order_related"
	JobResourceTypeOrderContact                         JobResourceType = "order_contact"
	JobResourceTypeUser                                 JobResourceType = "user"
	JobResourceTypeAddress                              JobResourceType = "address"
	JobResourceTypeAPIKey                               JobResourceType = "api_key"
	JobResourceTypeCreatedAPIKey                        JobResourceType = "created_api_key"
	JobResourceTypeRefreshToken                         JobResourceType = "refresh_token"
	JobResourceTypeList                                 JobResourceType = "list"
	JobResourceTypeSandbox                              JobResourceType = "sandbox"
	JobResourceTypeRegistrationSession                  JobResourceType = "registration_session"
	JobResourceTypePricingPlan                          JobResourceType = "pricing_plan"
	JobResourceTypeAccountPlan                          JobResourceType = "account_plan"
	JobResourceTypePlanChange                           JobResourceType = "plan_change"
	JobResourceTypeEnterpriseInquiry                    JobResourceType = "enterprise_inquiry"
	JobResourceTypeRequestLog                           JobResourceType = "request_log"
	JobResourceTypeAuditEvent                           JobResourceType = "audit_event"
	JobResourceTypeAuditFieldChange                     JobResourceType = "audit_field_change"
	JobResourceTypeRole                                 JobResourceType = "role"
	JobResourceTypeUnit                                 JobResourceType = "unit"
	JobResourceTypeAccountAffiliation                   JobResourceType = "account_affiliation"
	JobResourceTypeAgentDefinition                      JobResourceType = "agent_definition"
	JobResourceTypeAvailableTool                        JobResourceType = "available_tool"
	JobResourceTypeAgentDefinitionTool                  JobResourceType = "agent_definition_tool"
	JobResourceTypeAgentAccountStatus                   JobResourceType = "agent_account_status"
	JobResourceTypeAgentRun                             JobResourceType = "agent_run"
	JobResourceTypeAgentAction                          JobResourceType = "agent_action"
	JobResourceTypeAgentRunStep                         JobResourceType = "agent_run_step"
	JobResourceTypeAgentTokenUsage                      JobResourceType = "agent_token_usage"
	JobResourceTypeAgentMemory                          JobResourceType = "agent_memory"
	JobResourceTypeNotification                         JobResourceType = "notification"
	JobResourceTypeNotificationUnreadCount              JobResourceType = "notification_unread_count"
	JobResourceTypeNotificationSendResult               JobResourceType = "notification_send_result"
	JobResourceTypeNotificationUnreadSummary            JobResourceType = "notification_unread_summary"
	JobResourceTypeAnnouncement                         JobResourceType = "announcement"
	JobResourceTypeConversation                         JobResourceType = "conversation"
	JobResourceTypeSupportCase                          JobResourceType = "support_case"
	JobResourceTypeConversationParticipant              JobResourceType = "conversation_participant"
	JobResourceTypeReadCursor                           JobResourceType = "read_cursor"
	JobResourceTypeChatMessage                          JobResourceType = "chat_message"
	JobResourceTypeNotificationUnreadSummaryAccount     JobResourceType = "notification_unread_summary_account"
	JobResourceTypeMessagingBlock                       JobResourceType = "messaging_block"
	JobResourceTypeNotificationPreference               JobResourceType = "notification_preference"
	JobResourceTypeMessageAttachment                    JobResourceType = "message_attachment"
	JobResourceTypeAttachmentUploadTarget               JobResourceType = "attachment_upload_target"
	JobResourceTypeScheduledMessage                     JobResourceType = "scheduled_message"
	JobResourceTypeMessagingContact                     JobResourceType = "messaging_contact"
	JobResourceTypeMessageReport                        JobResourceType = "message_report"
	JobResourceTypeToolGroup                            JobResourceType = "tool_group"
	JobResourceTypeModel                                JobResourceType = "model"
	JobResourceTypePaymentTerm                          JobResourceType = "payment_term"
	JobResourceTypeShippingTerm                         JobResourceType = "shipping_term"
	JobResourceTypeQuantity                             JobResourceType = "quantity"
	JobResourceTypeAccountGroup                         JobResourceType = "account_group"
	JobResourceTypeSupportRoute                         JobResourceType = "support_route"
	JobResourceTypeSupportAvailability                  JobResourceType = "support_availability"
	JobResourceTypeAccountStatus                        JobResourceType = "account_status"
	JobResourceTypeGeolocation                          JobResourceType = "geolocation"
	JobResourceTypeAccountUser                          JobResourceType = "account_user"
	JobResourceTypeDepartment                           JobResourceType = "department"
	JobResourceTypeAccountIntegration                   JobResourceType = "account_integration"
	JobResourceTypeAccountPrice                         JobResourceType = "account_price"
	JobResourceTypeProductLine                          JobResourceType = "product_line"
	JobResourceTypeItemCategory                         JobResourceType = "item_category"
	JobResourceTypeAttribute                            JobResourceType = "attribute"
	JobResourceTypeRate                                 JobResourceType = "rate"
	JobResourceTypeAccountGroupProductLineAccess        JobResourceType = "account_group_product_line_access"
	JobResourceTypeSalesTarget                          JobResourceType = "sales_target"
	JobResourceTypeAdjustmentType                       JobResourceType = "adjustment_type"
	JobResourceTypeAccountBranding                      JobResourceType = "account_branding"
	JobResourceTypeAccountPortal                        JobResourceType = "account_portal"
	JobResourceTypeAccountLogoURL                       JobResourceType = "account_logo_url"
	JobResourceTypeAccountFaviconURL                    JobResourceType = "account_favicon_url"
	JobResourceTypePublicAccount                        JobResourceType = "public_account"
	JobResourceTypeProperty                             JobResourceType = "property"
	JobResourceTypeCarrier                              JobResourceType = "carrier"
	JobResourceTypeServiceLevel                         JobResourceType = "service_level"
	JobResourceTypeItem                                 JobResourceType = "item"
	JobResourceTypeItemLotDefault                       JobResourceType = "item_lot_default"
	JobResourceTypeItemInventory                        JobResourceType = "item_inventory"
	JobResourceTypeProduct                              JobResourceType = "product"
	JobResourceTypeBatch                                JobResourceType = "batch"
	JobResourceTypeBatchFlowNode                        JobResourceType = "batch_flow_node"
	JobResourceTypeScanningConsumption                  JobResourceType = "scanning_consumption"
	JobResourceTypeOpenBatchSummary                     JobResourceType = "open_batch_summary"
	JobResourceTypeScanningProductionStepInfo           JobResourceType = "scanning_production_step_info"
	JobResourceTypeScanningStation                      JobResourceType = "scanning_station"
	JobResourceTypeProductionStep                       JobResourceType = "production_step"
	JobResourceTypeProductionRun                        JobResourceType = "production_run"
	JobResourceTypeMachine                              JobResourceType = "machine"
	JobResourceTypeMachineStatus                        JobResourceType = "machine_status"
	JobResourceTypeMachineDowntimeEvent                 JobResourceType = "machine_downtime_event"
	JobResourceTypeDemandOverride                       JobResourceType = "demand_override"
	JobResourceTypeDemandOverrideType                   JobResourceType = "demand_override_type"
	JobResourceTypeMachineDowntimeReason                JobResourceType = "machine_downtime_reason"
	JobResourceTypeProductionSchedulePreview            JobResourceType = "production_schedule_preview"
	JobResourceTypeProductionScheduleRegeneratePreview  JobResourceType = "production_schedule_regenerate_preview"
	JobResourceTypeProductionSchedule                   JobResourceType = "production_schedule"
	JobResourceTypeProductionScheduleLine               JobResourceType = "production_schedule_line"
	JobResourceTypeProductionScheduleDeviation          JobResourceType = "production_schedule_deviation"
	JobResourceTypeProductionScheduleDerivedLine        JobResourceType = "production_schedule_derived_line"
	JobResourceTypeProductionScheduleSettings           JobResourceType = "production_schedule_settings"
	JobResourceTypeProductionScheduleResourceSetting    JobResourceType = "production_schedule_resource_setting"
	JobResourceTypeProductionScheduleItemSetting        JobResourceType = "production_schedule_item_setting"
	JobResourceTypeFulfillmentRecommendation            JobResourceType = "fulfillment_recommendation"
	JobResourceTypeAnalyzeDeliveryPerformanceResponse   JobResourceType = "analyze_delivery_performance_response"
	JobResourceTypeDeliveryPerformance                  JobResourceType = "delivery_performance"
	JobResourceTypeDeliveryBacklogBucket                JobResourceType = "delivery_backlog_bucket"
	JobResourceTypeDeliveryLatenessBucket               JobResourceType = "delivery_lateness_bucket"
	JobResourceTypeDeliveryBreakdown                    JobResourceType = "delivery_breakdown"
	JobResourceTypeAnalyzeSalesBreakdownResponse        JobResourceType = "analyze_sales_breakdown_response"
	JobResourceTypeSalesTotals                          JobResourceType = "sales_totals"
	JobResourceTypeSalesBreakdown                       JobResourceType = "sales_breakdown"
	JobResourceTypeScheduleOrderCoverage                JobResourceType = "schedule_order_coverage"
	JobResourceTypeScheduleOrderCoverageLine            JobResourceType = "schedule_order_coverage_line"
	JobResourceTypeScheduleDeviationType                JobResourceType = "schedule_deviation_type"
	JobResourceTypeScheduleAtRiskOrder                  JobResourceType = "schedule_at_risk_order"
	JobResourceTypeProductionScheduleFinishedPolicy     JobResourceType = "production_schedule_finished_policy"
	JobResourceTypeProductionScheduleFinishingLine      JobResourceType = "production_schedule_finishing_line"
	JobResourceTypeProductionScheduleWeekRelease        JobResourceType = "production_schedule_week_release"
	JobResourceTypeProductionScheduleWeekReleasePreview JobResourceType = "production_schedule_week_release_preview"
	JobResourceTypeProductionScheduleItemPolicy         JobResourceType = "production_schedule_item_policy"
	JobResourceTypeChildAccount                         JobResourceType = "child_account"
	JobResourceTypeUnitGroup                            JobResourceType = "unit_group"
	JobResourceTypeUnitGroupUnit                        JobResourceType = "unit_group_unit"
	JobResourceTypeConsumption                          JobResourceType = "consumption"
	JobResourceTypeCustomerProductLineAccess            JobResourceType = "customer_product_line_access"
	JobResourceTypeCustomer                             JobResourceType = "customer"
	JobResourceTypeFrequentlyOrderedProduct             JobResourceType = "frequently_ordered_product"
	JobResourceTypePriority                             JobResourceType = "priority"
	JobResourceTypeDelivery                             JobResourceType = "delivery"
	JobResourceTypeDeliveryLine                         JobResourceType = "delivery_line"
	JobResourceTypeSalesOrder                           JobResourceType = "sales_order"
	JobResourceTypeLocation                             JobResourceType = "location"
	JobResourceTypeLocationType                         JobResourceType = "location_type"
	JobResourceTypeLot                                  JobResourceType = "lot"
	JobResourceTypeEmailLog                             JobResourceType = "email_log"
	JobResourceTypeEmailDomain                          JobResourceType = "email_domain"
	JobResourceTypeEmailInbox                           JobResourceType = "email_inbox"
	JobResourceTypeEmailSender                          JobResourceType = "email_sender"
	JobResourceTypePortalDomain                         JobResourceType = "portal_domain"
	JobResourceTypeDNSRecord                            JobResourceType = "dns_record"
	JobResourceTypeInventoryChangeLog                   JobResourceType = "inventory_change_log"
	JobResourceTypeInvoice                              JobResourceType = "invoice"
	JobResourceTypeInvoiceSummary                       JobResourceType = "invoice_summary"
	JobResourceTypeInvoiceLine                          JobResourceType = "invoice_line"
	JobResourceTypeInvoiceAllocation                    JobResourceType = "invoice_allocation"
	JobResourceTypeInvoiceForPayment                    JobResourceType = "invoice_for_payment"
	JobResourceTypeShipment                             JobResourceType = "shipment"
	JobResourceTypeShipmentSummary                      JobResourceType = "shipment_summary"
	JobResourceTypeShipmentLine                         JobResourceType = "shipment_line"
	JobResourceTypeShippingCase                         JobResourceType = "shipping_case"
	JobResourceTypeShippingCaseLabelURL                 JobResourceType = "shipping_case_label_url"
	JobResourceTypeSettlement                           JobResourceType = "settlement"
	JobResourceTypeSettlementSummary                    JobResourceType = "settlement_summary"
	JobResourceTypeRolePermission                       JobResourceType = "role_permission"
	JobResourceTypeRegistrationFlow                     JobResourceType = "registration_flow"
	JobResourceTypeRegistrationFlowOption               JobResourceType = "registration_flow_option"
	JobResourceTypeTransaction                          JobResourceType = "transaction"
	JobResourceTypeTransactionSummary                   JobResourceType = "transaction_summary"
	JobResourceTypeTransactionMethod                    JobResourceType = "transaction_method"
	JobResourceTypeTransactionType                      JobResourceType = "transaction_type"
	JobResourceTypeTransactionAllocation                JobResourceType = "transaction_allocation"
	JobResourceTypeUsageItem                            JobResourceType = "usage_item"
	JobResourceTypeAccountUsageResponse                 JobResourceType = "account_usage_response"
	JobResourceTypeSubscriptionInfo                     JobResourceType = "subscription_info"
	JobResourceTypeBillingPortalSessionResponse         JobResourceType = "billing_portal_session_response"
	JobResourceTypeSwitchPlanResponse                   JobResourceType = "switch_plan_response"
	JobResourceTypeEnsureBillingCustomerResponse        JobResourceType = "ensure_billing_customer_response"
	JobResourceTypeSpendingCapResponse                  JobResourceType = "spending_cap_response"
	JobResourceTypeAgentSpendInfo                       JobResourceType = "agent_spend_info"
	JobResourceTypeWebhookResponse                      JobResourceType = "webhook_response"
	JobResourceTypeAddressSuggestion                    JobResourceType = "address_suggestion"
	JobResourceTypeAddressComponents                    JobResourceType = "address_components"
	JobResourceTypeAddressDetailsResult                 JobResourceType = "address_details_result"
	JobResourceTypeValidatedAddress                     JobResourceType = "validated_address"
	JobResourceTypePlanLimit                            JobResourceType = "plan_limit"
	JobResourceTypePlanChangeProration                  JobResourceType = "plan_change_proration"
	JobResourceTypePlanChangeLineItem                   JobResourceType = "plan_change_line_item"
	JobResourceTypeSetupBillingResponse                 JobResourceType = "setup_billing_response"
	JobResourceTypeConfirmPaymentResponse               JobResourceType = "confirm_payment_response"
	JobResourceTypeOAuthResponse                        JobResourceType = "oauth_response"
	JobResourceTypeOAuthStatusResponse                  JobResourceType = "oauth_status_response"
	JobResourceTypeStripePublishableKey                 JobResourceType = "stripe_publishable_key"
	JobResourceTypeStripeStatus                         JobResourceType = "stripe_status"
	JobResourceTypeHealthcheck                          JobResourceType = "healthcheck"
	JobResourceTypeAgentDefinitionConfig                JobResourceType = "agent_definition_config"
	JobResourceTypeTriggerConfig                        JobResourceType = "trigger_config"
	JobResourceTypeCustomerContactInfo                  JobResourceType = "customer_contact_info"
	JobResourceTypeCustomerFreightPreferences           JobResourceType = "customer_freight_preferences"
	JobResourceTypeCustomerDefaults                     JobResourceType = "customer_defaults"
	JobResourceTypeCustomerLeadTime                     JobResourceType = "customer_lead_time"
	JobResourceTypeCustomerNotificationPreferences      JobResourceType = "customer_notification_preferences"
	JobResourceTypeOrderNotificationRecipient           JobResourceType = "order_notification_recipient"
	JobResourceTypeOrderDiscount                        JobResourceType = "order_discount"
	JobResourceTypeSalesOrderLine                       JobResourceType = "sales_order_line"
	JobResourceTypeSalesOrderType                       JobResourceType = "sales_order_type"
	JobResourceTypeSalesOrderStatus                     JobResourceType = "sales_order_status"
	JobResourceTypeMaterial                             JobResourceType = "material"
	JobResourceTypeSupplierMaterial                     JobResourceType = "supplier_material"
	JobResourceTypePart                                 JobResourceType = "part"
	JobResourceTypePermissionGroup                      JobResourceType = "permission_group"
	JobResourceTypePermission                           JobResourceType = "permission"
	JobResourceTypePick                                 JobResourceType = "pick"
	JobResourceTypePickLine                             JobResourceType = "pick_line"
	JobResourceTypeProductType                          JobResourceType = "product_type"
	JobResourceTypeProduction                           JobResourceType = "production"
	JobResourceTypeProductionFlow                       JobResourceType = "production_flow"
	JobResourceTypeMap                                  JobResourceType = "map"
	JobResourceTypePurchaseOrder                        JobResourceType = "purchase_order"
	JobResourceTypePurchaseOrderLine                    JobResourceType = "purchase_order_line"
	JobResourceTypeSupplier                             JobResourceType = "supplier"
	JobResourceTypeSupplierSummary                      JobResourceType = "supplier_summary"
	JobResourceTypeReceivableEntry                      JobResourceType = "receivable_entry"
	JobResourceTypeReceivingOrder                       JobResourceType = "receiving_order"
	JobResourceTypeReceivingOrderLine                   JobResourceType = "receiving_order_line"
	JobResourceTypeEmailContact                         JobResourceType = "email_contact"
	JobResourceTypeAllocationEntry                      JobResourceType = "allocation_entry"
	JobResourceTypeOpenCreditEntry                      JobResourceType = "open_credit_entry"
	JobResourceTypeVolumeDiscount                       JobResourceType = "volume_discount"
	JobResourceTypeVolumeDiscountTier                   JobResourceType = "volume_discount_tier"
	JobResourceTypeAnalyzeDeliveriesResponse            JobResourceType = "analyze_deliveries_response"
	JobResourceTypeAnalyzeManufacturingResponse         JobResourceType = "analyze_manufacturing_response"
	JobResourceTypeAnalyzeManufacturingBatchResponse    JobResourceType = "analyze_manufacturing_batch_response"
	JobResourceTypeAnalyzeQuarterlyOrdersResponse       JobResourceType = "analyze_quarterly_orders_response"
	JobResourceTypeAnalyzeNewCustomersResponse          JobResourceType = "analyze_new_customers_response"
	JobResourceTypeAnalyzeDemandForecastResponse        JobResourceType = "analyze_demand_forecast_response"
	JobResourceTypeAnalyzeOeeResponse                   JobResourceType = "analyze_oee_response"
	JobResourceTypeAnalyzeOeeTrendResponse              JobResourceType = "analyze_oee_trend_response"
	JobResourceTypeAnalyzeScheduleAttainmentResponse    JobResourceType = "analyze_schedule_attainment_response"
	JobResourceTypeCatalogProductLine                   JobResourceType = "catalog_product_line"
	JobResourceTypeCatalogCategory                      JobResourceType = "catalog_category"
	JobResourceTypeCatalogProduct                       JobResourceType = "catalog_product"
	JobResourceTypeCatalogProperty                      JobResourceType = "catalog_property"
	JobResourceTypeCatalogAttribute                     JobResourceType = "catalog_attribute"
	JobResourceTypeDcLocation                           JobResourceType = "dc_location"
	JobResourceTypeEdiRun                               JobResourceType = "edi_run"
	JobResourceTypeInventoryItem                        JobResourceType = "inventory_item"
	JobResourceTypeAnalyzeWeeksOfSalesResponse          JobResourceType = "analyze_weeks_of_sales_response"
	JobResourceTypeBulkReconcileItemsResponse           JobResourceType = "bulk_reconcile_items_response"
	JobResourceTypeSysProperty                          JobResourceType = "sys_property"
	JobResourceTypeSysPropertyType                      JobResourceType = "sys_property_type"
	JobResourceTypeSysPropertyValue                     JobResourceType = "sys_property_value"
	JobResourceTypeTerritory                            JobResourceType = "territory"
	JobResourceTypeTenancy                              JobResourceType = "tenancy"
	JobResourceTypeCheckoutSession                      JobResourceType = "checkout_session"
	JobResourceTypeEstimateRateResult                   JobResourceType = "estimate_rate_result"
	JobResourceTypeRateShopOption                       JobResourceType = "rate_shop_option"
	JobResourceTypeRateShopResult                       JobResourceType = "rate_shop_result"
	JobResourceTypeOwner                                JobResourceType = "owner"
	JobResourceTypeCreatedBy                            JobResourceType = "created_by"
	JobResourceTypeMessage                              JobResourceType = "message"
	JobResourceTypeAccountPhotoUploadResult             JobResourceType = "account_photo_upload_result"
	JobResourceTypeUserPhotoUploadResult                JobResourceType = "user_photo_upload_result"
	JobResourceTypeUserPhotoURL                         JobResourceType = "user_photo_url"
	JobResourceTypeBatchLot                             JobResourceType = "batch_lot"
	JobResourceTypeCheckDuplicateResult                 JobResourceType = "check_duplicate_result"
	JobResourceTypeItemTrendPoint                       JobResourceType = "item_trend_point"
	JobResourceTypeTenancyPendingRegistration           JobResourceType = "tenancy_pending_registration"
	JobResourceTypeInvoiceAllocationEntry               JobResourceType = "invoice_allocation_entry"
	JobResourceTypeAllocationCustomer                   JobResourceType = "allocation_customer"
	JobResourceTypeCheckoutSalesOrder                   JobResourceType = "checkout_sales_order"
	JobResourceTypeSalesOrderPriceQuote                 JobResourceType = "sales_order_price_quote"
	JobResourceTypeSalesOrderFreightQuote               JobResourceType = "sales_order_freight_quote"
	JobResourceTypeSalesOrderCommitmentQuote            JobResourceType = "sales_order_commitment_quote"
	JobResourceTypeOperatingCalendar                    JobResourceType = "operating_calendar"
	JobResourceTypeOperatingCalendarClosure             JobResourceType = "operating_calendar_closure"
	JobResourceTypeSalesOrderPriceQuoteLine             JobResourceType = "sales_order_price_quote_line"
	JobResourceTypeHubspotSyncJob                       JobResourceType = "hubspot_sync_job"
	JobResourceTypeHubspotSyncReport                    JobResourceType = "hubspot_sync_report"
	JobResourceTypeHubspotCompanyReview                 JobResourceType = "hubspot_company_review"
	JobResourceTypeHubspotCompanyCandidate              JobResourceType = "hubspot_company_candidate"
	JobResourceTypeHubspotSyncRecord                    JobResourceType = "hubspot_sync_record"
	JobResourceTypeContactMatch                         JobResourceType = "contact_match"
	JobResourceTypeReplyDraft                           JobResourceType = "reply_draft"
	JobResourceTypeConversationLink                     JobResourceType = "conversation_link"
	JobResourceTypeMessagingGroup                       JobResourceType = "messaging_group"
	JobResourceTypeMessagingGroupMember                 JobResourceType = "messaging_group_member"
	JobResourceTypePortalProfile                        JobResourceType = "portal_profile"
	JobResourceTypePortalRegistrationSession            JobResourceType = "portal_registration_session"
	JobResourceTypePortalRegistrationSessionData        JobResourceType = "portal_registration_session_data"
	JobResourceTypePackList                             JobResourceType = "pack_list"
	JobResourceTypePackListParty                        JobResourceType = "pack_list_party"
	JobResourceTypePackListLineItem                     JobResourceType = "pack_list_line_item"
	JobResourceTypePackListBackOrder                    JobResourceType = "pack_list_back_order"
	JobResourceTypePackListCase                         JobResourceType = "pack_list_case"
	JobResourceTypeJob                                  JobResourceType = "job"
	JobResourceTypeJobResult                            JobResourceType = "job_result"
	JobResourceTypeJobExport                            JobResourceType = "job_export"
	JobResourceTypeAnalyzeCustomerPricingResponse       JobResourceType = "analyze_customer_pricing_response"
	JobResourceTypeCustomerPricingFinding               JobResourceType = "customer_pricing_finding"
	JobResourceTypeCustomerPricingSummary               JobResourceType = "customer_pricing_summary"
	JobResourceTypeComputedRate                         JobResourceType = "computed_rate"
	JobResourceTypeComputedQuantity                     JobResourceType = "computed_quantity"
	JobResourceTypeAnalyzeRealizedMarginsResponse       JobResourceType = "analyze_realized_margins_response"
	JobResourceTypeRealizedMarginFinding                JobResourceType = "realized_margin_finding"
	JobResourceTypeRealizedMarginSummary                JobResourceType = "realized_margin_summary"
	JobResourceTypeShipmentRelated                      JobResourceType = "shipment_related"
	JobResourceTypeInvoiceRelated                       JobResourceType = "invoice_related"
	JobResourceTypePickRelated                          JobResourceType = "pick_related"
	JobResourceTypePickTotals                           JobResourceType = "pick_totals"
	JobResourceTypePickStageTotal                       JobResourceType = "pick_stage_total"
)

// How far the job has got.
//
// `completed` means the work was processed, not that every row succeeded — read
// each entry's own `status` in `results`.
type JobStatus string

const (
	JobStatusCreated   JobStatus = "created"
	JobStatusStarted   JobStatus = "started"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusCancelled JobStatus = "cancelled"
)

// The kind of work the job carries out.
type JobType string

const (
	JobTypeBulkCreate JobType = "bulk_create"
	JobTypeBulkUpsert JobType = "bulk_upsert"
	JobTypeExport     JobType = "export"
	JobTypePackPick   JobType = "pack_pick"
)

// Points a completed export job at the file it produced.
type JobExport struct {
	// Resource type identifier.
	//
	// Any of "job_export".
	Object JobExportObject `json:"object" api:"required"`
	// Presigned link to the file, valid for five minutes.
	//
	// If the link has expired, read the job again for a fresh one.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Object      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExport) RawJSON() string { return r.JSON.raw }
func (r *JobExport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type JobExportObject string

const (
	JobExportObjectJobExport JobExportObject = "job_export"
)

// Accounts for one row of the request: the resource it produced, or the error it
// was rejected with. Every submitted row lands in exactly one of these once the
// job completes.
type JobResult struct {
	// ResponseError is the JSON-serializable error body returned to API clients. It
	// contains only public information. This struct is used by the OpenAPI schema
	// generator to produce documentation.
	Error ResponseError `json:"error" api:"required"`
	// Zero-based row of the request this result names.
	Index int64 `json:"index" api:"required"`
	// Resource type identifier.
	//
	// Any of "job_result".
	Object JobResultObject `json:"object" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// What became of the row.
	//
	// - `created`: the row produced a new resource.
	// - `updated`: the row updated an existing resource.
	// - `failed`: the row was rejected and wrote nothing.
	//
	// Any of "created", "updated", "failed".
	Status JobResultStatus `json:"status" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	SubResources ListEntity `json:"sub_resources" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error        respjson.Field
		Index        respjson.Field
		Object       respjson.Field
		Resource     respjson.Field
		Status       respjson.Field
		SubResources respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResult) RawJSON() string { return r.JSON.raw }
func (r *JobResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type JobResultObject string

const (
	JobResultObjectJobResult JobResultObject = "job_result"
)

// What became of the row.
//
// - `created`: the row produced a new resource.
// - `updated`: the row updated an existing resource.
// - `failed`: the row was rejected and wrote nothing.
type JobResultStatus string

const (
	JobResultStatusCreated JobResultStatus = "created"
	JobResultStatusUpdated JobResultStatus = "updated"
	JobResultStatusFailed  JobResultStatus = "failed"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListJobResult struct {
	// Resources in this page.
	Data []JobResult `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListJobResultObject `json:"object" api:"required"`
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
func (r ListJobResult) RawJSON() string { return r.JSON.raw }
func (r *ListJobResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListJobResultObject string

const (
	ListJobResultObjectList ListJobResultObject = "list"
)

// QuotaInfo provides machine-readable details about a plan-imposed resource limit.
// Included in limit_exceeded errors so clients can display upgrade prompts, usage
// bars, or implement programmatic retry/backoff logic.
type QuotaInfo struct {
	// Limit is the maximum number of resources allowed by the current plan.
	Limit int64 `json:"limit" api:"required"`
	// ResetAt is the time when the quota resets, if applicable. Nil for static
	// (non-metered) limits.
	ResetAt time.Time `json:"reset_at" api:"required" format:"date-time"`
	// Used is the number of resources currently consumed.
	Used int64 `json:"used" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		ResetAt     respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuotaInfo) RawJSON() string { return r.JSON.raw }
func (r *QuotaInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseError is the JSON-serializable error body returned to API clients. It
// contains only public information. This struct is used by the OpenAPI schema
// generator to produce documentation.
type ResponseError struct {
	// A machine-readable code for the error.
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
	Code ResponseErrorCode `json:"code" api:"required"`
	// A URL to documentation about the error.
	DocURL string `json:"doc_url" api:"required"`
	// Whether this error is transient and the request can be retried.
	IsTransient bool `json:"is_transient" api:"required"`
	// A human-readable message providing more details about the error.
	Message string `json:"message" api:"required"`
	// The parameter that caused the error, if applicable.
	Param string `json:"param" api:"required"`
	// QuotaInfo provides machine-readable details about a plan-imposed resource limit.
	// Included in limit_exceeded errors so clients can display upgrade prompts, usage
	// bars, or implement programmatic retry/backoff logic.
	Quota QuotaInfo `json:"quota" api:"required"`
	// RequestLogURL is a link to the dashboard page for this request's log entry. Nil
	// when no request log is available.
	RequestLogURL string `json:"request_log_url" api:"required"`
	// The type of error.
	//
	// Any of "api_error", "idempotency_error", "invalid_request_error".
	Type ResponseErrorType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code          respjson.Field
		DocURL        respjson.Field
		IsTransient   respjson.Field
		Message       respjson.Field
		Param         respjson.Field
		Quota         respjson.Field
		RequestLogURL respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseError) RawJSON() string { return r.JSON.raw }
func (r *ResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A machine-readable code for the error.
type ResponseErrorCode string

const (
	ResponseErrorCodeExpiredToken            ResponseErrorCode = "expired_token"
	ResponseErrorCodeAPIKeyExpired           ResponseErrorCode = "api_key_expired"
	ResponseErrorCodeAPIKeyRevoked           ResponseErrorCode = "api_key_revoked"
	ResponseErrorCodeInvalidCredentials      ResponseErrorCode = "invalid_credentials"
	ResponseErrorCodeInsufficientPermissions ResponseErrorCode = "insufficient_permissions"
	ResponseErrorCodePaymentRequired         ResponseErrorCode = "payment_required"
	ResponseErrorCodeAgentSpendingCapReached ResponseErrorCode = "agent_spending_cap_reached"
	ResponseErrorCodeValidationFailed        ResponseErrorCode = "validation_failed"
	ResponseErrorCodeMissingField            ResponseErrorCode = "missing_field"
	ResponseErrorCodeInvalidFormat           ResponseErrorCode = "invalid_format"
	ResponseErrorCodeMethodNotAllowed        ResponseErrorCode = "method_not_allowed"
	ResponseErrorCodeResourceNotFound        ResponseErrorCode = "resource_not_found"
	ResponseErrorCodeResourceExists          ResponseErrorCode = "resource_exists"
	ResponseErrorCodeResourceConflict        ResponseErrorCode = "resource_conflict"
	ResponseErrorCodeResourceGone            ResponseErrorCode = "resource_gone"
	ResponseErrorCodeIdempotencyInProgress   ResponseErrorCode = "idempotency_in_progress"
	ResponseErrorCodeLimitExceeded           ResponseErrorCode = "limit_exceeded"
	ResponseErrorCodeRegistrationClosed      ResponseErrorCode = "registration_closed"
	ResponseErrorCodeRateLimitExceeded       ResponseErrorCode = "rate_limit_exceeded"
	ResponseErrorCodeParameterMissing        ResponseErrorCode = "parameter_missing"
	ResponseErrorCodeParameterInvalid        ResponseErrorCode = "parameter_invalid"
	ResponseErrorCodeParameterUnknown        ResponseErrorCode = "parameter_unknown"
	ResponseErrorCodeParametersExclusive     ResponseErrorCode = "parameters_exclusive"
	ResponseErrorCodeInternalError           ResponseErrorCode = "internal_error"
	ResponseErrorCodeServiceUnavailable      ResponseErrorCode = "service_unavailable"
	ResponseErrorCodeExternalServiceError    ResponseErrorCode = "external_service_error"
	ResponseErrorCodeTimeout                 ResponseErrorCode = "timeout"
	ResponseErrorCodeConnectionError         ResponseErrorCode = "connection_error"
	ResponseErrorCodeRequestTimeout          ResponseErrorCode = "request_timeout"
	ResponseErrorCodeClientClosedRequest     ResponseErrorCode = "client_closed_request"
	ResponseErrorCodeAPIVersionRequired      ResponseErrorCode = "api_version_required"
	ResponseErrorCodeAPIVersionInvalid       ResponseErrorCode = "api_version_invalid"
	ResponseErrorCodeAPIVersionTooOld        ResponseErrorCode = "api_version_too_old"
)

// The type of error.
type ResponseErrorType string

const (
	ResponseErrorTypeAPIError            ResponseErrorType = "api_error"
	ResponseErrorTypeIdempotencyError    ResponseErrorType = "idempotency_error"
	ResponseErrorTypeInvalidRequestError ResponseErrorType = "invalid_request_error"
)

type CoreJobGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreJobGetParams]'s query parameters as `url.Values`.
func (r CoreJobGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreJobCancelParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "created_by", "created_by.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreJobCancelParams]'s query parameters as `url.Values`.
func (r CoreJobCancelParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
