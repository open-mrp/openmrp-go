// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Manage per-category notification channel preferences (in-app, email, push).
//
// MessagingPreferenceService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingPreferenceService] method instead.
type MessagingPreferenceService struct {
	options []option.RequestOption
}

// NewMessagingPreferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingPreferenceService(opts ...option.RequestOption) (r MessagingPreferenceService) {
	r = MessagingPreferenceService{}
	r.options = opts
	return
}

// Creates or replaces one of the current user's notification preferences, either
// their global default or the override for a single category.
//
// The preference applies only to the account being acted in, and the category must
// be one the platform recognizes. Callers without a user membership in that
// account cannot hold preferences and are refused.
//
// This endpoint requires the permission: `messaging:update`.
func (r *MessagingPreferenceService) Update(ctx context.Context, body MessagingPreferenceUpdateParams, opts ...option.RequestOption) (res *NotificationPreference, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists the current user's notification preferences for the account they are
// acting in: their global default plus any per-category overrides.
//
// Only preferences the user has explicitly set are returned, so an empty list
// means everything falls back to the standard behavior — in-app notifications on,
// email and push off.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingPreferenceService) List(ctx context.Context, opts ...option.RequestOption) (res *ListNotificationPreference, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/preferences"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListNotificationPreference struct {
	// Resources in this page.
	Data []NotificationPreference `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListNotificationPreferenceObject `json:"object" api:"required"`
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
func (r ListNotificationPreference) RawJSON() string { return r.JSON.raw }
func (r *ListNotificationPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListNotificationPreferenceObject string

const (
	ListNotificationPreferenceObjectList ListNotificationPreferenceObject = "list"
)

// One user's choice of which channels a category of notification is delivered on.
//
// Preferences belong to the user's membership in a single account, so the same
// person can be notified differently in each account they belong to. A preference
// with no category is that user's global default, and a category-specific
// preference overrides it. Where neither exists, in-app notifications are
// delivered and email and push are not.
//
// Chat notifications are the only ones these settings currently govern:
// notifications in every other category reach the in-app feed and are never
// emailed, whatever is stored here.
type NotificationPreference struct {
	// Preference ID.
	ID string `json:"id" api:"required"`
	// The notification category this preference applies to.
	//
	// A preference with no category is the user's global default, used for every
	// category they have not set a specific preference for.
	//
	// Any of "chat.message", "chat.mention", "chat.added", "order.updated",
	// "agent.run_completed", "agent.alert", "system.broadcast", "customer.registered".
	Category NotificationPreferenceCategory `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// How often email for this category is sent.
	//
	//   - `instant`: send an email as soon as an eligible notification occurs.
	//   - `hourly`: collect eligible notifications into a single hourly email.
	//   - `daily`: collect eligible notifications into a single daily email.
	//   - `off`: never send email for this category, even when email is otherwise
	//     enabled.
	//
	// This governs email only; in-app delivery is unaffected. Batched sending is not
	// running yet, so `hourly` and `daily` currently hold email back in the same way
	// as `off`.
	//
	// Any of "instant", "hourly", "daily", "off".
	Digest NotificationPreferenceDigest `json:"digest" api:"required"`
	// Whether notifications in this category are also emailed to the user.
	//
	// Email is additionally suppressed for a conversation the user has muted, and only
	// sent on the cadence set by `digest`.
	EmailEnabled bool `json:"email_enabled" api:"required"`
	// Whether notifications in this category appear in the user's in-app feed.
	//
	// A direct @mention is always delivered in-app, even when this is disabled.
	InAppEnabled bool `json:"in_app_enabled" api:"required"`
	// Resource type identifier.
	//
	// Any of "notification_preference".
	Object NotificationPreferenceObject `json:"object" api:"required"`
	// Whether notifications in this category are also sent as push notifications.
	//
	// Push delivery is not available yet; the choice is stored for when it is.
	PushEnabled bool `json:"push_enabled" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Category     respjson.Field
		CreatedAt    respjson.Field
		Digest       respjson.Field
		EmailEnabled respjson.Field
		InAppEnabled respjson.Field
		Object       respjson.Field
		PushEnabled  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotificationPreference) RawJSON() string { return r.JSON.raw }
func (r *NotificationPreference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The notification category this preference applies to.
//
// A preference with no category is the user's global default, used for every
// category they have not set a specific preference for.
type NotificationPreferenceCategory string

const (
	NotificationPreferenceCategoryChatMessage        NotificationPreferenceCategory = "chat.message"
	NotificationPreferenceCategoryChatMention        NotificationPreferenceCategory = "chat.mention"
	NotificationPreferenceCategoryChatAdded          NotificationPreferenceCategory = "chat.added"
	NotificationPreferenceCategoryOrderUpdated       NotificationPreferenceCategory = "order.updated"
	NotificationPreferenceCategoryAgentRunCompleted  NotificationPreferenceCategory = "agent.run_completed"
	NotificationPreferenceCategoryAgentAlert         NotificationPreferenceCategory = "agent.alert"
	NotificationPreferenceCategorySystemBroadcast    NotificationPreferenceCategory = "system.broadcast"
	NotificationPreferenceCategoryCustomerRegistered NotificationPreferenceCategory = "customer.registered"
)

// How often email for this category is sent.
//
//   - `instant`: send an email as soon as an eligible notification occurs.
//   - `hourly`: collect eligible notifications into a single hourly email.
//   - `daily`: collect eligible notifications into a single daily email.
//   - `off`: never send email for this category, even when email is otherwise
//     enabled.
//
// This governs email only; in-app delivery is unaffected. Batched sending is not
// running yet, so `hourly` and `daily` currently hold email back in the same way
// as `off`.
type NotificationPreferenceDigest string

const (
	NotificationPreferenceDigestInstant NotificationPreferenceDigest = "instant"
	NotificationPreferenceDigestHourly  NotificationPreferenceDigest = "hourly"
	NotificationPreferenceDigestDaily   NotificationPreferenceDigest = "daily"
	NotificationPreferenceDigestOff     NotificationPreferenceDigest = "off"
)

// Resource type identifier.
type NotificationPreferenceObject string

const (
	NotificationPreferenceObjectNotificationPreference NotificationPreferenceObject = "notification_preference"
)

// Request to create or replace one of the caller's notification preferences.
//
// A user has at most one preference per category, so sending the same category
// again replaces the previous settings outright — every channel is written from
// this request, not merged with what was there before.
//
// Chat notifications are the only ones these settings currently govern:
// notifications in every other category reach the in-app feed and are never
// emailed, whatever is stored here.
//
// The properties EmailEnabled, InAppEnabled, PushEnabled are required.
type UpsertNotificationPreferenceRequestParam struct {
	// Whether notifications in this category are also emailed to the user.
	//
	// Email is additionally suppressed for a conversation the user has muted, and only
	// sent on the cadence set by `digest`.
	EmailEnabled bool `json:"email_enabled" api:"required"`
	// Whether notifications in this category appear in the user's in-app feed.
	//
	// A direct @mention is always delivered in-app, even when this is off.
	InAppEnabled bool `json:"in_app_enabled" api:"required"`
	// Whether notifications in this category are also sent as push notifications.
	//
	// Push delivery is not available yet; the choice is stored for when it is.
	PushEnabled bool `json:"push_enabled" api:"required"`
	// The notification category these settings apply to, such as `chat.message`.
	//
	// Leave it out to set the global default used for every category without its own
	// preference.
	Category param.Opt[string] `json:"category,omitzero"`
	// How often email for this category is sent.
	//
	//   - `instant`: send an email as soon as an eligible notification occurs.
	//   - `hourly`: collect eligible notifications into a single hourly email.
	//   - `daily`: collect eligible notifications into a single daily email.
	//   - `off`: never send email for this category, even when email is otherwise
	//     enabled.
	//
	// This governs email only; in-app delivery is unaffected. Batched sending is not
	// running yet, so `hourly` and `daily` currently hold email back in the same way
	// as `off`.
	//
	// Any of "instant", "hourly", "daily", "off".
	Digest UpsertNotificationPreferenceRequestDigest `json:"digest,omitzero"`
	paramObj
}

func (r UpsertNotificationPreferenceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpsertNotificationPreferenceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpsertNotificationPreferenceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How often email for this category is sent.
//
//   - `instant`: send an email as soon as an eligible notification occurs.
//   - `hourly`: collect eligible notifications into a single hourly email.
//   - `daily`: collect eligible notifications into a single daily email.
//   - `off`: never send email for this category, even when email is otherwise
//     enabled.
//
// This governs email only; in-app delivery is unaffected. Batched sending is not
// running yet, so `hourly` and `daily` currently hold email back in the same way
// as `off`.
type UpsertNotificationPreferenceRequestDigest string

const (
	UpsertNotificationPreferenceRequestDigestInstant UpsertNotificationPreferenceRequestDigest = "instant"
	UpsertNotificationPreferenceRequestDigestHourly  UpsertNotificationPreferenceRequestDigest = "hourly"
	UpsertNotificationPreferenceRequestDigestDaily   UpsertNotificationPreferenceRequestDigest = "daily"
	UpsertNotificationPreferenceRequestDigestOff     UpsertNotificationPreferenceRequestDigest = "off"
)

type MessagingPreferenceUpdateParams struct {
	// Request to create or replace one of the caller's notification preferences.
	//
	// A user has at most one preference per category, so sending the same category
	// again replaces the previous settings outright — every channel is written from
	// this request, not merged with what was there before.
	//
	// Chat notifications are the only ones these settings currently govern:
	// notifications in every other category reach the in-app feed and are never
	// emailed, whatever is stored here.
	UpsertNotificationPreferenceRequest UpsertNotificationPreferenceRequestParam
	paramObj
}

func (r MessagingPreferenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpsertNotificationPreferenceRequest)
}
func (r *MessagingPreferenceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
