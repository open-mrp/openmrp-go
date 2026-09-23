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

// List, read, and manage broadcast announcements.
//
// MessagingAnnouncementService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingAnnouncementService] method instead.
type MessagingAnnouncementService struct {
	options []option.RequestOption
	// List, read, and manage broadcast announcements.
	Actions MessagingAnnouncementActionService
}

// NewMessagingAnnouncementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingAnnouncementService(opts ...option.RequestOption) (r MessagingAnnouncementService) {
	r = MessagingAnnouncementService{}
	r.options = opts
	r.Actions = NewMessagingAnnouncementActionService(opts...)
	return
}

// Retrieves a single announcement by ID, with the calling user's own read state.
//
// Only announcements the caller can see are returned: one published to another
// account, one that has not reached its publish time, or one that has expired is
// reported as not found. An announcement the caller has dismissed stays
// retrievable even though it no longer appears in their feed.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingAnnouncementService) Get(ctx context.Context, id string, query MessagingAnnouncementGetParams, opts ...option.RequestOption) (res *Announcement, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/announcements/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists the announcements currently active for the caller, newest first.
//
// The feed covers announcements broadcast to the account being acted in together
// with platform-wide announcements from OpenMRP. Announcements the caller has
// dismissed are left out, as are any that are scheduled for later or have already
// expired.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingAnnouncementService) List(ctx context.Context, query MessagingAnnouncementListParams, opts ...option.RequestOption) (res *ListAnnouncement, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/announcements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A broadcast announcement shown in the notification (bell) feed, carrying the
// calling user's own read state.
//
// A single announcement is published to everyone in an account, or to every user
// on the platform, and each user keeps their own seen, read, and dismissed state
// for it. The status and timestamps you read are therefore always the caller's,
// and never reflect what anyone else has done with the same announcement.
// Notifications addressed to one user are a separate resource.
type Announcement struct {
	// Announcement ID.
	ID string `json:"id" api:"required"`
	// Supporting detail shown beneath the title.
	Body string `json:"body" api:"required"`
	// The kind of event the announcement is about.
	//
	// Announcements draw on the same categories as notifications, such as
	// `system.broadcast` or `order.updated`, and the category is chosen by whoever
	// publishes the announcement. The set is open-ended and may grow over time, so
	// clients should tolerate values they do not recognize.
	//
	// Any of "chat.message", "chat.mention", "chat.added", "order.updated",
	// "agent.run_completed", "agent.alert", "system.broadcast", "customer.registered",
	// "production_run.updated".
	Category AnnouncementCategory `json:"category" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the calling user dismissed the announcement.
	DismissedAt time.Time `json:"dismissed_at" api:"required" format:"date-time"`
	// When the announcement stops being shown.
	//
	// Once it expires the announcement leaves every user's feed and can no longer be
	// retrieved; an announcement with no expiry stays until each user dismisses it.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "announcement".
	Object AnnouncementObject `json:"object" api:"required"`
	// How prominently the announcement should be surfaced, from `low` through
	// `urgent`.
	//
	// Any of "low", "normal", "high", "urgent".
	Priority AnnouncementPriority `json:"priority" api:"required"`
	// When the announcement becomes visible in the feed.
	//
	// An announcement scheduled for the future is not returned by the announcement
	// endpoints until this time passes.
	PublishAt time.Time `json:"publish_at" api:"required" format:"date-time"`
	// When the calling user opened the announcement.
	ReadAt time.Time `json:"read_at" api:"required" format:"date-time"`
	// Entity is a polymorphic reference to any resource in the system.
	Resource Entity `json:"resource" api:"required"`
	// Who the announcement reaches.
	//
	//   - `account`: published to a single account and shown only to that account's
	//     users.
	//   - `platform`: published by OpenMRP and shown to every user across all accounts.
	//
	// Any of "account", "platform".
	Scope AnnouncementScope `json:"scope" api:"required"`
	// When the calling user first saw the announcement.
	SeenAt time.Time `json:"seen_at" api:"required" format:"date-time"`
	// Where the announcement is in its lifecycle for the calling user.
	//
	// - `unseen`: not yet surfaced to the caller.
	// - `seen`: surfaced in the caller's feed but not opened.
	// - `read`: explicitly opened by the caller.
	// - `dismissed`: removed from the caller's feed.
	//
	// The status is derived from the caller's own seen, read, and dismissed timestamps
	// and only ever moves forward, so the same announcement can show a different
	// status for each user in the account.
	//
	// Any of "unseen", "seen", "read", "dismissed".
	Status AnnouncementStatus `json:"status" api:"required"`
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
		ExpiresAt   respjson.Field
		Object      respjson.Field
		Priority    respjson.Field
		PublishAt   respjson.Field
		ReadAt      respjson.Field
		Resource    respjson.Field
		Scope       respjson.Field
		SeenAt      respjson.Field
		Status      respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Announcement) RawJSON() string { return r.JSON.raw }
func (r *Announcement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of event the announcement is about.
//
// Announcements draw on the same categories as notifications, such as
// `system.broadcast` or `order.updated`, and the category is chosen by whoever
// publishes the announcement. The set is open-ended and may grow over time, so
// clients should tolerate values they do not recognize.
type AnnouncementCategory string

const (
	AnnouncementCategoryChatMessage          AnnouncementCategory = "chat.message"
	AnnouncementCategoryChatMention          AnnouncementCategory = "chat.mention"
	AnnouncementCategoryChatAdded            AnnouncementCategory = "chat.added"
	AnnouncementCategoryOrderUpdated         AnnouncementCategory = "order.updated"
	AnnouncementCategoryAgentRunCompleted    AnnouncementCategory = "agent.run_completed"
	AnnouncementCategoryAgentAlert           AnnouncementCategory = "agent.alert"
	AnnouncementCategorySystemBroadcast      AnnouncementCategory = "system.broadcast"
	AnnouncementCategoryCustomerRegistered   AnnouncementCategory = "customer.registered"
	AnnouncementCategoryProductionRunUpdated AnnouncementCategory = "production_run.updated"
)

// Resource type identifier.
type AnnouncementObject string

const (
	AnnouncementObjectAnnouncement AnnouncementObject = "announcement"
)

// How prominently the announcement should be surfaced, from `low` through
// `urgent`.
type AnnouncementPriority string

const (
	AnnouncementPriorityLow    AnnouncementPriority = "low"
	AnnouncementPriorityNormal AnnouncementPriority = "normal"
	AnnouncementPriorityHigh   AnnouncementPriority = "high"
	AnnouncementPriorityUrgent AnnouncementPriority = "urgent"
)

// Who the announcement reaches.
//
//   - `account`: published to a single account and shown only to that account's
//     users.
//   - `platform`: published by OpenMRP and shown to every user across all accounts.
type AnnouncementScope string

const (
	AnnouncementScopeAccount  AnnouncementScope = "account"
	AnnouncementScopePlatform AnnouncementScope = "platform"
)

// Where the announcement is in its lifecycle for the calling user.
//
// - `unseen`: not yet surfaced to the caller.
// - `seen`: surfaced in the caller's feed but not opened.
// - `read`: explicitly opened by the caller.
// - `dismissed`: removed from the caller's feed.
//
// The status is derived from the caller's own seen, read, and dismissed timestamps
// and only ever moves forward, so the same announcement can show a different
// status for each user in the account.
type AnnouncementStatus string

const (
	AnnouncementStatusUnseen    AnnouncementStatus = "unseen"
	AnnouncementStatusSeen      AnnouncementStatus = "seen"
	AnnouncementStatusRead      AnnouncementStatus = "read"
	AnnouncementStatusDismissed AnnouncementStatus = "dismissed"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAnnouncement struct {
	// Resources in this page.
	Data []Announcement `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAnnouncementObject `json:"object" api:"required"`
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
func (r ListAnnouncement) RawJSON() string { return r.JSON.raw }
func (r *ListAnnouncement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAnnouncementObject string

const (
	ListAnnouncementObjectList ListAnnouncementObject = "list"
)

type MessagingAnnouncementGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementGetParams]'s query parameters as
// `url.Values`.
func (r MessagingAnnouncementGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingAnnouncementListParams struct {
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
	// Any of "resource".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingAnnouncementListParams]'s query parameters as
// `url.Values`.
func (r MessagingAnnouncementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
