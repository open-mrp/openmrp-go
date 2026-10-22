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

// List and manage third-party account integrations.
//
// SettingIntegrationService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingIntegrationService] method instead.
type SettingIntegrationService struct {
	options []option.RequestOption
	// List and manage third-party account integrations.
	Openrouter SettingIntegrationOpenrouterService
}

// NewSettingIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingIntegrationService(opts ...option.RequestOption) (r SettingIntegrationService) {
	r = SettingIntegrationService{}
	r.options = opts
	r.Openrouter = NewSettingIntegrationOpenrouterService(opts...)
	return
}

// Connects a third-party provider to the account, or replaces the name and
// credentials of the provider's existing connection.
//
// An account can have at most one integration per `provider`, so calling this
// again for a provider that is already connected rotates its credentials in place
// and returns the same integration rather than creating a second one. Credentials
// are checked for the provider's expected key format, encrypted at rest, and never
// returned in API responses.
//
// This endpoint requires the `admin` role type.
func (r *SettingIntegrationService) New(ctx context.Context, body SettingIntegrationNewParams, opts ...option.RequestOption) (res *AccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Renames an account integration, activates or deactivates it, or changes its
// monthly spend limit.
//
// Omitted fields are left unchanged. Credentials cannot be changed here; to rotate
// them, call Create Account Integration again with the same `provider`.
//
// This endpoint requires the `admin` role type.
func (r *SettingIntegrationService) Update(ctx context.Context, id string, body SettingIntegrationUpdateParams, opts ...option.RequestOption) (res *AccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns a paginated list of the third-party providers connected to the target
// account.
//
// Stored credentials are never included in the response.
//
// This endpoint requires the `admin` role type.
func (r *SettingIntegrationService) List(ctx context.Context, query SettingIntegrationListParams, opts ...option.RequestOption) (res *ListAccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/integrations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Disconnects a third-party provider from the account and returns the deleted
// integration.
//
// The stored credentials go with it, so any feature that relies on the provider
// stops working until the integration is created again. Deleting an integration
// that is already deleted returns an error rather than succeeding silently. To
// pause a provider without discarding its credentials, set the integration's
// status to `inactive` instead.
//
// This endpoint requires the `admin` role type.
func (r *SettingIntegrationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/settings/integrations/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Third-party integration connected to an account.
//
// An account can have at most one integration per provider. The credentials
// supplied when the integration was connected are encrypted at rest and are never
// returned by the API.
type AccountIntegration struct {
	// Account integration ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A monetary amount.
	MonthlySpendLimit Money `json:"monthly_spend_limit" api:"required"`
	// Display name of the integration.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_integration".
	Object AccountIntegrationObject `json:"object" api:"required"`
	// Integration provider code.
	//
	// - `stripe`: Stripe payment processing.
	// - `shippo`: Shippo shipping and label generation.
	// - `hubspot`: HubSpot CRM.
	// - `openrouter`: OpenRouter, the LLM gateway that serves agent runs.
	//
	// Any of "stripe", "shippo", "hubspot", "openrouter".
	Provider AccountIntegrationProvider `json:"provider" api:"required"`
	// Lifecycle status of the integration.
	//
	// Integrations are created `active`. Setting an integration to `inactive` keeps
	// its stored credentials but stops it from being used (for example, the Stripe
	// publishable key cannot be retrieved while the Stripe integration is inactive).
	//
	// Any of "active", "inactive".
	Status AccountIntegrationStatus `json:"status" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		MonthlySpendLimit respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		Provider          respjson.Field
		Status            respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountIntegration) RawJSON() string { return r.JSON.raw }
func (r *AccountIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountIntegrationObject string

const (
	AccountIntegrationObjectAccountIntegration AccountIntegrationObject = "account_integration"
)

// Integration provider code.
//
// - `stripe`: Stripe payment processing.
// - `shippo`: Shippo shipping and label generation.
// - `hubspot`: HubSpot CRM.
// - `openrouter`: OpenRouter, the LLM gateway that serves agent runs.
type AccountIntegrationProvider string

const (
	AccountIntegrationProviderStripe     AccountIntegrationProvider = "stripe"
	AccountIntegrationProviderShippo     AccountIntegrationProvider = "shippo"
	AccountIntegrationProviderHubspot    AccountIntegrationProvider = "hubspot"
	AccountIntegrationProviderOpenrouter AccountIntegrationProvider = "openrouter"
)

// Lifecycle status of the integration.
//
// Integrations are created `active`. Setting an integration to `inactive` keeps
// its stored credentials but stops it from being used (for example, the Stripe
// publishable key cannot be retrieved while the Stripe integration is inactive).
type AccountIntegrationStatus string

const (
	AccountIntegrationStatusActive   AccountIntegrationStatus = "active"
	AccountIntegrationStatusInactive AccountIntegrationStatus = "inactive"
)

// Request to create or upsert an account integration.
//
// The properties Credentials, Name, Provider are required.
type CreateAccountIntegrationRequestParam struct {
	// JSON string containing the provider's credentials.
	//
	// Required keys depend on the provider:
	//
	//   - `stripe`: `private_key` (`sk_...`), `publishable_key` (`pk_...`), and
	//     `webhook_secret` (`whsec_...`).
	//   - `shippo`: `api_key` (`shippo_live_...` or `shippo_test_...`).
	//   - `hubspot`: `access_token` (`pat-...`).
	//   - `openrouter`: `api_key` (`sk-or-...`).
	//
	// For Stripe and Shippo, sandbox accounts must supply test keys and production
	// accounts must supply live keys; credentials that do not match are rejected.
	// HubSpot and OpenRouter keys make no such distinction.
	Credentials string `json:"credentials" api:"required"`
	// Display name of the integration.
	Name string `json:"name" api:"required"`
	// Integration provider code.
	//
	// - `stripe`: Stripe payment processing.
	// - `shippo`: Shippo shipping and label generation.
	// - `hubspot`: HubSpot CRM.
	// - `openrouter`: OpenRouter, the LLM gateway that serves agent runs.
	//
	// Any of "stripe", "shippo", "hubspot", "openrouter".
	Provider CreateAccountIntegrationRequestProvider `json:"provider,omitzero" api:"required"`
	// A monetary amount, supplied on create and update requests.
	MonthlySpendLimit MoneyInputParam `json:"monthly_spend_limit,omitzero"`
	paramObj
}

func (r CreateAccountIntegrationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAccountIntegrationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAccountIntegrationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration provider code.
//
// - `stripe`: Stripe payment processing.
// - `shippo`: Shippo shipping and label generation.
// - `hubspot`: HubSpot CRM.
// - `openrouter`: OpenRouter, the LLM gateway that serves agent runs.
type CreateAccountIntegrationRequestProvider string

const (
	CreateAccountIntegrationRequestProviderStripe     CreateAccountIntegrationRequestProvider = "stripe"
	CreateAccountIntegrationRequestProviderShippo     CreateAccountIntegrationRequestProvider = "shippo"
	CreateAccountIntegrationRequestProviderHubspot    CreateAccountIntegrationRequestProvider = "hubspot"
	CreateAccountIntegrationRequestProviderOpenrouter CreateAccountIntegrationRequestProvider = "openrouter"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAccountIntegration struct {
	// Resources in this page.
	Data []AccountIntegration `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAccountIntegrationObject `json:"object" api:"required"`
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
func (r ListAccountIntegration) RawJSON() string { return r.JSON.raw }
func (r *ListAccountIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAccountIntegrationObject string

const (
	ListAccountIntegrationObjectList ListAccountIntegrationObject = "list"
)

// A monetary amount.
type Money struct {
	// Decimal amount, as a string to preserve precision.
	Amount string `json:"amount" api:"required" format:"decimal"`
	// Lowercase ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Money) RawJSON() string { return r.JSON.raw }
func (r *Money) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A monetary amount, supplied on create and update requests.
//
// The properties Amount, Currency are required.
type MoneyInputParam struct {
	// Decimal amount, as a string to preserve precision.
	Amount string `json:"amount" api:"required" format:"decimal"`
	// Lowercase ISO 4217 currency code.
	Currency string `json:"currency" api:"required"`
	paramObj
}

func (r MoneyInputParam) MarshalJSON() (data []byte, err error) {
	type shadow MoneyInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MoneyInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to update an account integration.
type UpdateAccountIntegrationRequestParam struct {
	// Display name of the integration.
	Name param.Opt[string] `json:"name,omitzero"`
	// A monetary amount, supplied on create and update requests.
	MonthlySpendLimit MoneyInputParam `json:"monthly_spend_limit,omitzero"`
	// Lifecycle status of the integration.
	//
	// Set to `inactive` to stop the provider being used while keeping its stored
	// credentials, and back to `active` to resume without re-entering them.
	//
	// Any of "active", "inactive".
	Status UpdateAccountIntegrationRequestStatus `json:"status,omitzero"`
	paramObj
}

func (r UpdateAccountIntegrationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAccountIntegrationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAccountIntegrationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the integration.
//
// Set to `inactive` to stop the provider being used while keeping its stored
// credentials, and back to `active` to resume without re-entering them.
type UpdateAccountIntegrationRequestStatus string

const (
	UpdateAccountIntegrationRequestStatusActive   UpdateAccountIntegrationRequestStatus = "active"
	UpdateAccountIntegrationRequestStatusInactive UpdateAccountIntegrationRequestStatus = "inactive"
)

type SettingIntegrationNewParams struct {
	// Request to create or upsert an account integration.
	CreateAccountIntegrationRequest CreateAccountIntegrationRequestParam
	paramObj
}

func (r SettingIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAccountIntegrationRequest)
}
func (r *SettingIntegrationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIntegrationUpdateParams struct {
	// Request to update an account integration.
	UpdateAccountIntegrationRequest UpdateAccountIntegrationRequestParam
	paramObj
}

func (r SettingIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAccountIntegrationRequest)
}
func (r *SettingIntegrationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIntegrationListParams struct {
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
	paramObj
}

// URLQuery serializes [SettingIntegrationListParams]'s query parameters as
// `url.Values`.
func (r SettingIntegrationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
