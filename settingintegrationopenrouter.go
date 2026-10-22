// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage third-party account integrations.
//
// SettingIntegrationOpenrouterService contains methods and other services that
// help with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingIntegrationOpenrouterService] method instead.
type SettingIntegrationOpenrouterService struct {
	options []option.RequestOption
}

// NewSettingIntegrationOpenrouterService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingIntegrationOpenrouterService(opts ...option.RequestOption) (r SettingIntegrationOpenrouterService) {
	r = SettingIntegrationOpenrouterService{}
	r.options = opts
	return
}

// Returns what the account has spent through its OpenRouter integration today,
// this week, this month and in total, with its monthly spend limit.
//
// Spend is read live from OpenRouter for the integration's API key. Fails with
// `404` if the account has no OpenRouter integration, and with `400` if OpenRouter
// rejects the stored key.
//
// This endpoint requires the `admin` role type.
func (r *SettingIntegrationOpenrouterService) GetUsage(ctx context.Context, opts ...option.RequestOption) (res *OpenRouterUsage, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/settings/integrations/openrouter/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// What an account has spent through its OpenRouter integration.
//
// Spend comes from OpenRouter and covers every request made with the integration's
// API key, including any made outside OpenMRP. Days, weeks and months are UTC
// calendar periods.
type OpenRouterUsage struct {
	// A monetary amount.
	KeyLimit Money `json:"key_limit" api:"required"`
	// A monetary amount.
	KeyLimitRemaining Money `json:"key_limit_remaining" api:"required"`
	// A monetary amount.
	MonthlySpendLimit Money `json:"monthly_spend_limit" api:"required"`
	// A monetary amount.
	MonthlySpendRemaining Money `json:"monthly_spend_remaining" api:"required"`
	// Resource type identifier.
	//
	// Any of "openrouter_usage".
	Object OpenRouterUsageObject `json:"object" api:"required"`
	// A monetary amount.
	SpentThisMonth Money `json:"spent_this_month" api:"required"`
	// A monetary amount.
	SpentThisWeek Money `json:"spent_this_week" api:"required"`
	// A monetary amount.
	SpentToday Money `json:"spent_today" api:"required"`
	// A monetary amount.
	SpentTotal Money `json:"spent_total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyLimit              respjson.Field
		KeyLimitRemaining     respjson.Field
		MonthlySpendLimit     respjson.Field
		MonthlySpendRemaining respjson.Field
		Object                respjson.Field
		SpentThisMonth        respjson.Field
		SpentThisWeek         respjson.Field
		SpentToday            respjson.Field
		SpentTotal            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenRouterUsage) RawJSON() string { return r.JSON.raw }
func (r *OpenRouterUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type OpenRouterUsageObject string

const (
	OpenRouterUsageObjectOpenrouterUsage OpenRouterUsageObject = "openrouter_usage"
)
