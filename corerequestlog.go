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

// List and retrieve request logs.
//
// CoreRequestLogService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreRequestLogService] method instead.
type CoreRequestLogService struct {
	options []option.RequestOption
}

// NewCoreRequestLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCoreRequestLogService(opts ...option.RequestOption) (r CoreRequestLogService) {
	r = CoreRequestLogService{}
	r.options = opts
	return
}

// Returns a single API request log by ID.
//
// The log is readable when your account is either the acting account or the
// account that was acted upon. This is also the only endpoint that can return the
// captured query parameters and request and response bodies, and the only way to
// read the high-traffic-endpoint logs that are withheld from the list endpoint.
//
// This endpoint requires the permission: `request_logs:read`.
func (r *CoreRequestLogService) Get(ctx context.Context, id string, query CoreRequestLogGetParams, opts ...option.RequestOption) (res *RequestLog, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/core/request-logs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of API request logs, newest first.
//
// Results cover every request where your account is either the acting account or
// the account that was acted upon, so requests a customer or supplier made against
// your data appear alongside your own. The `q` parameter matches a log ID exactly
// and otherwise searches the request path, the normalized route, and the error
// message.
//
// Requests to a number of high-traffic endpoints — including these logging
// endpoints themselves — are recorded but withheld from this listing so they do
// not drown out the rest of your traffic. They can still be fetched individually
// by ID.
//
// This endpoint requires the permission: `request_logs:read`.
func (r *CoreRequestLogService) List(ctx context.Context, query CoreRequestLogListParams, opts ...option.RequestOption) (res *ListRequestLog, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/request-logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Reference to an actor — the user, API key, agent, or group identity associated
// with an action.
type Actor struct {
	// Unique identifier of the actor.
	ID string `json:"id" api:"required"`
	// URL of the actor's profile photo, if one is set.
	//
	// Only populated for `user` actors.
	AvatarURL string `json:"avatar_url" api:"required"`
	// Human-readable handle identifying the actor.
	//
	// - For `user` actors: the user's email address.
	// - For `api_key` actors: the redacted key value.
	//
	// Other actor types carry no handle.
	Handle string `json:"handle" api:"required"`
	// The actor's display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "actor".
	Object ActorObject `json:"object" api:"required"`
	// A named set of permissions that can be assigned to users to control what they
	// can access.
	Role Role `json:"role" api:"required"`
	// Actor type.
	//
	//   - `user`: a human user account.
	//   - `api_key`: a programmatic caller authenticating with an API key.
	//   - `agent`: an automated agent acting on the account's behalf.
	//   - `group`: a shared group identity, such as a "Customer Service" persona, rather
	//     than a single individual.
	//
	// Any of "user", "api_key", "agent", "group".
	Type ActorType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AvatarURL   respjson.Field
		Handle      respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Role        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Actor) RawJSON() string { return r.JSON.raw }
func (r *Actor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ActorObject string

const (
	ActorObjectActor ActorObject = "actor"
)

// Actor type.
//
//   - `user`: a human user account.
//   - `api_key`: a programmatic caller authenticating with an API key.
//   - `agent`: an automated agent acting on the account's behalf.
//   - `group`: a shared group identity, such as a "Customer Service" persona, rather
//     than a single individual.
type ActorType string

const (
	ActorTypeUser   ActorType = "user"
	ActorTypeAPIKey ActorType = "api_key"
	ActorTypeAgent  ActorType = "agent"
	ActorTypeGroup  ActorType = "group"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListRequestLog struct {
	// Resources in this page.
	Data []RequestLog `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListRequestLogObject `json:"object" api:"required"`
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
func (r ListRequestLog) RawJSON() string { return r.JSON.raw }
func (r *ListRequestLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListRequestLogObject string

const (
	ListRequestLogObjectList ListRequestLogObject = "list"
)

// A log of a single API request, capturing its route, outcome, latency, and actor.
//
// Logs are written after the response has been sent, so a new entry may take a
// moment to become readable.
type RequestLog struct {
	// Request log ID.
	ID string `json:"id" api:"required"`
	// An organization on OpenMRP, including its branding and customer portal
	// sub-resources.
	//
	// Your own account and any customer or supplier account you trade with are both
	// represented by this object.
	Account Account `json:"account" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// The API version the request was served with.
	//
	// Taken from the `OpenMRP-Version` header the caller sent; requests rejected for
	// omitting that header record no version.
	APIVersion string `json:"api_version" api:"required"`
	// Client IP address the request came from.
	//
	// Not recorded for requests an OpenMRP agent made on your behalf, since those
	// originate inside OpenMRP's own network.
	ClientIP string `json:"client_ip" api:"required"`
	// When the log entry was written.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Machine-readable API error code.
	//
	// Matches the `code` of the error response the caller received. Populated only for
	// failed requests.
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
	ErrorCode RequestLogErrorCode `json:"error_code" api:"required"`
	// Human-readable error message.
	//
	// The same message the caller received. Populated only for failed requests.
	ErrorMessage string `json:"error_message" api:"required"`
	// Request host.
	//
	// Usually `api.openmrp.ai`.
	Host string `json:"host" api:"required"`
	// User-provided idempotency key.
	IdempotencyKey string `json:"idempotency_key" api:"required"`
	// Request latency in microseconds.
	//
	// Measured at the API edge, from the moment the request was received until the
	// response was written, so it excludes network time between your client and
	// OpenMRP.
	LatencyUs int64 `json:"latency_us" api:"required"`
	// HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS".
	Method RequestLogMethod `json:"method" api:"required"`
	// The route template the request matched, with path parameters left as
	// placeholders.
	//
	// For example `/v1/sales/customers/{id}` is the normalized route for the request
	// path `/v1/sales/customers/ac_...`. Falls back to the raw path when the request
	// did not match a registered route.
	NormalizedRoute string `json:"normalized_route" api:"required"`
	// Resource type identifier.
	//
	// Any of "request_log".
	Object RequestLogObject `json:"object" api:"required"`
	// When the request was received.
	//
	// Request logs are ordered and date-filtered by this timestamp rather than by
	// `created_at`.
	OccurredAt time.Time `json:"occurred_at" api:"required" format:"date-time"`
	// The exact path the request was made to, including path parameter values.
	Path string `json:"path" api:"required"`
	// Query-string parameters the request was made with, as a JSON object. Encoded as
	// a JSON value (object, array, string, number, boolean, or null), not a
	// JSON-encoded string.
	QueryParams any `json:"query_params" api:"required"`
	// Referrer header.
	Referrer string `json:"referrer" api:"required"`
	// The JSON body the request was sent with.
	//
	// Sensitive values such as passwords, tokens, and secrets are redacted before the
	// body is stored. Bodies larger than 256 KB are not stored in full; a small marker
	// object with `_truncated` set to `true` is stored in their place. Encoded as a
	// JSON value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	RequestBody any `json:"request_body" api:"required"`
	// The JSON body OpenMRP responded with.
	//
	// Sensitive values such as generated API key secrets are redacted before the body
	// is stored. Bodies larger than 256 KB are not stored in full; a small marker
	// object with `_truncated` set to `true` is stored in their place. Encoded as a
	// JSON value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	ResponseBody any `json:"response_body" api:"required"`
	// HTTP response status code (e.g. `200`, `404`).
	StatusCode int64 `json:"status_code" api:"required"`
	// User agent.
	UserAgent string `json:"user_agent" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Account         respjson.Field
		Actor           respjson.Field
		APIVersion      respjson.Field
		ClientIP        respjson.Field
		CreatedAt       respjson.Field
		ErrorCode       respjson.Field
		ErrorMessage    respjson.Field
		Host            respjson.Field
		IdempotencyKey  respjson.Field
		LatencyUs       respjson.Field
		Method          respjson.Field
		NormalizedRoute respjson.Field
		Object          respjson.Field
		OccurredAt      respjson.Field
		Path            respjson.Field
		QueryParams     respjson.Field
		Referrer        respjson.Field
		RequestBody     respjson.Field
		ResponseBody    respjson.Field
		StatusCode      respjson.Field
		UserAgent       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestLog) RawJSON() string { return r.JSON.raw }
func (r *RequestLog) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Machine-readable API error code.
//
// Matches the `code` of the error response the caller received. Populated only for
// failed requests.
type RequestLogErrorCode string

const (
	RequestLogErrorCodeExpiredToken            RequestLogErrorCode = "expired_token"
	RequestLogErrorCodeAPIKeyExpired           RequestLogErrorCode = "api_key_expired"
	RequestLogErrorCodeAPIKeyRevoked           RequestLogErrorCode = "api_key_revoked"
	RequestLogErrorCodeInvalidCredentials      RequestLogErrorCode = "invalid_credentials"
	RequestLogErrorCodeInsufficientPermissions RequestLogErrorCode = "insufficient_permissions"
	RequestLogErrorCodePaymentRequired         RequestLogErrorCode = "payment_required"
	RequestLogErrorCodeAgentSpendingCapReached RequestLogErrorCode = "agent_spending_cap_reached"
	RequestLogErrorCodeValidationFailed        RequestLogErrorCode = "validation_failed"
	RequestLogErrorCodeMissingField            RequestLogErrorCode = "missing_field"
	RequestLogErrorCodeInvalidFormat           RequestLogErrorCode = "invalid_format"
	RequestLogErrorCodeMethodNotAllowed        RequestLogErrorCode = "method_not_allowed"
	RequestLogErrorCodeResourceNotFound        RequestLogErrorCode = "resource_not_found"
	RequestLogErrorCodeResourceExists          RequestLogErrorCode = "resource_exists"
	RequestLogErrorCodeResourceConflict        RequestLogErrorCode = "resource_conflict"
	RequestLogErrorCodeResourceGone            RequestLogErrorCode = "resource_gone"
	RequestLogErrorCodeIdempotencyInProgress   RequestLogErrorCode = "idempotency_in_progress"
	RequestLogErrorCodeLimitExceeded           RequestLogErrorCode = "limit_exceeded"
	RequestLogErrorCodeRegistrationClosed      RequestLogErrorCode = "registration_closed"
	RequestLogErrorCodeRateLimitExceeded       RequestLogErrorCode = "rate_limit_exceeded"
	RequestLogErrorCodeParameterMissing        RequestLogErrorCode = "parameter_missing"
	RequestLogErrorCodeParameterInvalid        RequestLogErrorCode = "parameter_invalid"
	RequestLogErrorCodeParameterUnknown        RequestLogErrorCode = "parameter_unknown"
	RequestLogErrorCodeParametersExclusive     RequestLogErrorCode = "parameters_exclusive"
	RequestLogErrorCodeInternalError           RequestLogErrorCode = "internal_error"
	RequestLogErrorCodeServiceUnavailable      RequestLogErrorCode = "service_unavailable"
	RequestLogErrorCodeExternalServiceError    RequestLogErrorCode = "external_service_error"
	RequestLogErrorCodeTimeout                 RequestLogErrorCode = "timeout"
	RequestLogErrorCodeConnectionError         RequestLogErrorCode = "connection_error"
	RequestLogErrorCodeRequestTimeout          RequestLogErrorCode = "request_timeout"
	RequestLogErrorCodeClientClosedRequest     RequestLogErrorCode = "client_closed_request"
	RequestLogErrorCodeAPIVersionRequired      RequestLogErrorCode = "api_version_required"
	RequestLogErrorCodeAPIVersionInvalid       RequestLogErrorCode = "api_version_invalid"
	RequestLogErrorCodeAPIVersionTooOld        RequestLogErrorCode = "api_version_too_old"
)

// HTTP method.
type RequestLogMethod string

const (
	RequestLogMethodGet     RequestLogMethod = "GET"
	RequestLogMethodPost    RequestLogMethod = "POST"
	RequestLogMethodPut     RequestLogMethod = "PUT"
	RequestLogMethodPatch   RequestLogMethod = "PATCH"
	RequestLogMethodDelete  RequestLogMethod = "DELETE"
	RequestLogMethodHead    RequestLogMethod = "HEAD"
	RequestLogMethodOptions RequestLogMethod = "OPTIONS"
)

// Resource type identifier.
type RequestLogObject string

const (
	RequestLogObjectRequestLog RequestLogObject = "request_log"
)

type CoreRequestLogGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account", "actor", "actor.role", "actor.role.permissions",
	// "query_params", "request_body", "response_body".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreRequestLogGetParams]'s query parameters as
// `url.Values`.
func (r CoreRequestLogGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CoreRequestLogListParams struct {
	// Opaque cursor token identifying where the page of results starts.
	//
	// Use the `cursor` value embedded in a previous response's `next_page_url` or
	// `previous_page_url` to fetch the adjacent page. Omit to start from the first
	// page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Restricts results to request logs on or before this timestamp.
	EndsAt param.Opt[time.Time] `query:"ends_at,omitzero" format:"date-time" json:"-"`
	// Filter by the user-provided idempotency key.
	IdempotencyKey param.Opt[string] `query:"idempotency_key,omitzero" json:"-"`
	// Maximum number of results to return in a single page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Restricts results to requests that took at least this many microseconds.
	MinLatencyUs param.Opt[int64] `query:"min_latency_us,omitzero" json:"-"`
	// Free-text search term used to filter results.
	//
	// Which fields are matched against the term varies by endpoint.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Restricts results to request logs on or after this timestamp.
	StartsAt param.Opt[time.Time] `query:"starts_at,omitzero" format:"date-time" json:"-"`
	// Filter by the _acting_ account: the account the actor belongs to (the log's
	// `account.id`).
	//
	// Results are always scoped to logs where your account is either the acting
	// account or the target account; this narrows that set to specific acting
	// accounts. For example, pass a customer's account ID to see only requests that
	// customer's actors made against your account.
	ActorAccountIDs []string `query:"actor_account_ids,omitzero" json:"-"`
	// Filter by the actor identifier.
	//
	// Matches the log's `actor.id`: a user ID for `user` actors, an API key ID for
	// `api_key` actors, or an agent ID for `agent` actors.
	ActorIDs []string `query:"actor_ids,omitzero" json:"-"`
	// Filter by the actor type.
	//
	// Requests are recorded for actors of type `user`, `api_key`, and `agent` — the
	// last covering calls an OpenMRP agent made on your account's behalf.
	//
	// Any of "user", "api_key", "agent", "group".
	ActorTypes []string `query:"actor_types,omitzero" json:"-"`
	// Filter by API error code.
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
	ErrorCodes []string `query:"error_codes,omitzero" json:"-"`
	// Exclude request logs whose API error code is in this set.
	//
	// Applied as a negative filter after all other filters. Successful requests (which
	// have no error code) are always kept. The OpenMRP dashboard uses this to hide
	// routine `expired_token` 401s — the noise from short-lived access tokens expiring
	// and clients silently refreshing — while still surfacing genuine auth failures
	// like `invalid_credentials`.
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
	ExcludeErrorCodes []string `query:"exclude_error_codes,omitzero" json:"-"`
	// Filter by the request host.
	//
	// Typically `api.openmrp.ai`.
	Hosts []string `query:"hosts,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "account", "actor", "actor.role".
	Include []string `query:"include,omitzero" json:"-"`
	// Filter by the HTTP method.
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS".
	Methods []string `query:"methods,omitzero" json:"-"`
	// Filter by the _normalized_ route template.
	//
	// For example `/v1/sales/customers/{id}` matches every request to that route
	// regardless of the specific customer ID. Parameter names inside `{}` are ignored
	// when matching, so `{customer_id}` and `{id}` are equivalent.
	NormalizedRoutes []string `query:"normalized_routes,omitzero" json:"-"`
	// Filter by the HTTP status class, expressed as the leading digit: `1`–`5` for
	// 1xx–5xx.
	//
	// Combined with `status_codes` using OR — e.g. `status_codes=401` and
	// `status_code_classes=5` matches 401 responses and any 5xx response.
	StatusCodeClasses []int64 `query:"status_code_classes,omitzero" json:"-"`
	// Filter by the HTTP status code.
	StatusCodes []int64 `query:"status_codes,omitzero" json:"-"`
	// Filter by the _target_ account: the account the request acted upon (the log's
	// target account).
	//
	// Results are always scoped to logs where your account is either the acting
	// account or the target account; this narrows that set to specific target
	// accounts. For example, pass a supplier's account ID to see only requests your
	// account made against that supplier.
	TargetAccountIDs []string `query:"target_account_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CoreRequestLogListParams]'s query parameters as
// `url.Values`.
func (r CoreRequestLogListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
