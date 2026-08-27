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

// Create and manage API keys for programmatic access.
//
// AuthAPIKeyService contains methods and other services that help with interacting
// with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthAPIKeyService] method instead.
type AuthAPIKeyService struct {
	options []option.RequestOption
	// Create and manage API keys for programmatic access.
	Actions AuthAPIKeyActionService
}

// NewAuthAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthAPIKeyService(opts ...option.RequestOption) (r AuthAPIKeyService) {
	r = AuthAPIKeyService{}
	r.options = opts
	r.Actions = NewAuthAPIKeyActionService(opts...)
	return
}

// Creates an [API key](https://docs.openmrp.ai/api/api-keys) to authenticate API
// requests.
//
// The key belongs to the account it was created under and only ever acts on behalf
// of that account. Keys created under a sandbox account carry an `mrp_sk_test_`
// prefix; keys created under a production account carry an `mrp_sk_prod_` prefix.
//
// The secret key is returned once and cannot be retrieved later, so you should
// store it securely. We provide some
// [recommendations](https://docs.openmrp.ai/api/managing-api-keys) on how you can
// manage your API keys.
//
// This endpoint requires the `admin` role type.
func (r *AuthAPIKeyService) New(ctx context.Context, params AuthAPIKeyNewParams, opts ...option.RequestOption) (res *CreatedAPIKey, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/auth/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns [API key](https://docs.openmrp.ai/api/api-keys) metadata by ID.
//
// Only the redacted key value is returned. The full secret is available only in
// the response that issued the key, so a lost secret must be replaced by rotating
// the key.
//
// This endpoint requires the `admin` role type.
func (r *AuthAPIKeyService) Get(ctx context.Context, id string, query AuthAPIKeyGetParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/auth/api-keys/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of [API keys](https://docs.openmrp.ai/api/api-keys),
// newest first.
//
// Only keys belonging to the account making the request are returned. The search
// term matches against the key name.
//
// This endpoint requires the `admin` role type.
func (r *AuthAPIKeyService) List(ctx context.Context, query AuthAPIKeyListParams, opts ...option.RequestOption) (res *ListAPIKey, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/auth/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Revokes an [API key](https://docs.openmrp.ai/api/api-keys).
//
// Revocation takes effect immediately and cannot be undone; any request still
// presenting the key is rejected. The key record is kept, so it stays visible in
// the key list with a `revoked` status. To replace a key without an interruption
// in access, use Rotate API Key instead.
//
// This endpoint requires the `admin` role type.
func (r *AuthAPIKeyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AuthAPIKeyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/auth/api-keys/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// An organization on OpenMRP, including its branding and customer portal
// sub-resources.
//
// Your own account and any customer or supplier account you trade with are both
// represented by this object.
type Account struct {
	// Account ID.
	ID string `json:"id" api:"required"`
	// The customer-facing branding an account presents on its portal, emails, and
	// documents.
	Branding AccountBranding `json:"branding" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	DefaultBillingAddress Address `json:"default_billing_address" api:"required"`
	// A saved address that can be used for billing and shipping on sales orders,
	// invoices, and shipments.
	DefaultShippingAddress Address `json:"default_shipping_address" api:"required"`
	// The account's display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "account".
	Object AccountObject `json:"object" api:"required"`
	// The customer portal an account publishes for its customers to sign in to.
	Portal AccountPortal `json:"portal" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Branding               respjson.Field
		CreatedAt              respjson.Field
		DefaultBillingAddress  respjson.Field
		DefaultShippingAddress respjson.Field
		Name                   respjson.Field
		Object                 respjson.Field
		Portal                 respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountObject string

const (
	AccountObjectAccount AccountObject = "account"
)

// The customer-facing branding an account presents on its portal, emails, and
// documents.
type AccountBranding struct {
	// Branding ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Facebook handle.
	FacebookHandle string `json:"facebook_handle" api:"required"`
	// Stored location of the account's customer-portal favicon.
	//
	// Favicons uploaded through the API are stored as an object key rather than a
	// fetchable link, so use the Get Account Favicon URL endpoint to obtain a download
	// URL.
	FaviconURL string `json:"favicon_url" api:"required"`
	// Instagram handle.
	InstagramHandle string `json:"instagram_handle" api:"required"`
	// LinkedIn handle.
	LinkedinHandle string `json:"linkedin_handle" api:"required"`
	// Stored location of the account's logo image.
	//
	// Logos uploaded through the API are stored as an object key rather than a
	// fetchable link, so use the Get Account Logo URL endpoint to obtain a download
	// URL.
	LogoURL string `json:"logo_url" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_branding".
	Object AccountBrandingObject `json:"object" api:"required"`
	// The account's public contact phone number.
	PhoneNumber string `json:"phone_number" api:"required"`
	// The email address customers are directed to for support.
	SupportEmail string `json:"support_email" api:"required"`
	// Twitter handle.
	TwitterHandle string `json:"twitter_handle" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The account's public website.
	WebsiteURL string `json:"website_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FacebookHandle  respjson.Field
		FaviconURL      respjson.Field
		InstagramHandle respjson.Field
		LinkedinHandle  respjson.Field
		LogoURL         respjson.Field
		Object          respjson.Field
		PhoneNumber     respjson.Field
		SupportEmail    respjson.Field
		TwitterHandle   respjson.Field
		UpdatedAt       respjson.Field
		WebsiteURL      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountBranding) RawJSON() string { return r.JSON.raw }
func (r *AccountBranding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountBrandingObject string

const (
	AccountBrandingObjectAccountBranding AccountBrandingObject = "account_branding"
)

// The customer portal an account publishes for its customers to sign in to.
type AccountPortal struct {
	// Portal ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "account_portal".
	Object AccountPortalObject `json:"object" api:"required"`
	// URL slug that identifies the account's customer portal.
	//
	// Unique across all accounts.
	Slug string `json:"slug" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Object      respjson.Field
		Slug        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountPortal) RawJSON() string { return r.JSON.raw }
func (r *AccountPortal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountPortalObject string

const (
	AccountPortalObjectAccountPortal AccountPortalObject = "account_portal"
)

// A saved address that can be used for billing and shipping on sales orders,
// invoices, and shipments.
type Address struct {
	// Address ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address associated with the address.
	Email string `json:"email" api:"required"`
	// The street-level location details of an address.
	Geolocation Geolocation `json:"geolocation" api:"required"`
	// Display name of the address.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "address".
	Object AddressObject `json:"object" api:"required"`
	// Phone number associated with the address.
	Phone string `json:"phone" api:"required"`
	// The operating calendar naming the days this dock accepts freight.
	//
	// The most specific link in the receiving chain: set it when one of a customer's
	// sites keeps different days from the rest. Null falls through to the customer,
	// then their group, then the account default.
	ReceiveCalendarID string `json:"receive_calendar_id" api:"required"`
	// How the address is used.
	//
	//   - `standard`: a normal shipping or billing address.
	//   - `drop_ship`: an address an order is shipped to directly, typically a third
	//     party or end customer rather than the account itself.
	//
	// Any of "standard", "drop_ship".
	Type AddressType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Email             respjson.Field
		Geolocation       respjson.Field
		Name              respjson.Field
		Object            respjson.Field
		Phone             respjson.Field
		ReceiveCalendarID respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AddressObject string

const (
	AddressObjectAddress AddressObject = "address"
)

// How the address is used.
//
//   - `standard`: a normal shipping or billing address.
//   - `drop_ship`: an address an order is shipped to directly, typically a third
//     party or end customer rather than the account itself.
type AddressType string

const (
	AddressTypeStandard AddressType = "standard"
	AddressTypeDropShip AddressType = "drop_ship"
)

// An API key used to authenticate requests to the OpenMRP API.
//
// A key always acts on behalf of the account it was created under, with the
// permissions of the role assigned to it.
type APIKey struct {
	// API key ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key expires and stops authenticating requests.
	//
	// A key with no expiration keeps working until it is revoked or rotated.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// When the key was last used to authenticate a request.
	//
	// Recorded at most once every 24 hours, so it can lag the key's most recent use by
	// up to a day.
	LastUsedAt time.Time `json:"last_used_at" api:"required" format:"date-time"`
	// Human-readable name for the API key.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "api_key".
	Object APIKeyObject `json:"object" api:"required"`
	// Redacted key value safe for display.
	//
	// The key's prefix followed by its last four characters, e.g.
	// `mrp_sk_prod_****hjt4`.
	RedactedValue string `json:"redacted_value" api:"required"`
	// When the key's revocation takes effect.
	//
	// A future timestamp means revocation was scheduled (for example, by a rotation)
	// and the key continues to authenticate requests until that time.
	RevokedAt time.Time `json:"revoked_at" api:"required" format:"date-time"`
	// A named set of permissions that can be assigned to users to control what they
	// can access.
	Role Role `json:"role" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		ExpiresAt     respjson.Field
		LastUsedAt    respjson.Field
		Name          respjson.Field
		Object        respjson.Field
		RedactedValue respjson.Field
		RevokedAt     respjson.Field
		Role          respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type APIKeyObject string

const (
	APIKeyObjectAPIKey APIKeyObject = "api_key"
)

// Request to create an API key.
//
// The properties Name, RoleID are required.
type CreateAPIKeyRequestParam struct {
	// Human-readable name for the API key.
	//
	// Shown when listing keys and used to match keys when searching, so prefer
	// something that identifies the integration using it.
	Name string `json:"name" api:"required"`
	// ID of the role to assign to the API key.
	//
	// The role determines what requests authenticated with the key are allowed to do.
	// A key keeps its role for life — including through rotation — so issue a new key
	// to use a different one, while changes to the role's own permissions take effect
	// for existing keys immediately.
	RoleID string `json:"role_id" api:"required"`
	// When the key expires and stops authenticating requests.
	//
	// If omitted, the key keeps working until it is revoked or rotated.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	paramObj
}

func (r CreateAPIKeyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAPIKeyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAPIKeyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A newly issued API key together with its secret value, returned when a key is
// created or rotated.
type CreatedAPIKey struct {
	// An API key used to authenticate requests to the OpenMRP API.
	//
	// A key always acts on behalf of the account it was created under, with the
	// permissions of the role assigned to it.
	APIKeyInfo APIKey `json:"api_key_info" api:"required"`
	// The secret used to authenticate requests, sent as a bearer token in the
	// `Authorization` header.
	//
	// This is the only response that ever contains the secret; if it is lost, rotate
	// the key to issue a new one. Learn more about
	// [managing your API keys](https://docs.openmrp.ai/api/managing-api-keys).
	APIKeySecret string `json:"api_key_secret" api:"required"`
	// Resource type identifier.
	//
	// Any of "created_api_key".
	Object CreatedAPIKeyObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyInfo   respjson.Field
		APIKeySecret respjson.Field
		Object       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedAPIKey) RawJSON() string { return r.JSON.raw }
func (r *CreatedAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type CreatedAPIKeyObject string

const (
	CreatedAPIKeyObjectCreatedAPIKey CreatedAPIKeyObject = "created_api_key"
)

// The street-level location details of an address.
type Geolocation struct {
	// Geolocation ID.
	ID string `json:"id" api:"required"`
	// Two-letter country code.
	Country string `json:"country" api:"required"`
	// City or locality.
	Locality string `json:"locality" api:"required"`
	// Resource type identifier.
	//
	// Any of "geolocation".
	Object GeolocationObject `json:"object" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// State or administrative area.
	State string `json:"state" api:"required"`
	// First line of the street address.
	StreetLine1 string `json:"street_line_1" api:"required"`
	// Second line of the street address.
	StreetLine2 string `json:"street_line_2" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Country     respjson.Field
		Locality    respjson.Field
		Object      respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		StreetLine1 respjson.Field
		StreetLine2 respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Geolocation) RawJSON() string { return r.JSON.raw }
func (r *Geolocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type GeolocationObject string

const (
	GeolocationObjectGeolocation GeolocationObject = "geolocation"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAPIKey struct {
	// Resources in this page.
	Data []APIKey `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAPIKeyObject `json:"object" api:"required"`
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
func (r ListAPIKey) RawJSON() string { return r.JSON.raw }
func (r *ListAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAPIKeyObject string

const (
	ListAPIKeyObjectList ListAPIKeyObject = "list"
)

// Owner describes the provenance of a resource.
type Owner struct {
	// An organization on OpenMRP, including its branding and customer portal
	// sub-resources.
	//
	// Your own account and any customer or supplier account you trade with are both
	// represented by this object.
	Account Account `json:"account" api:"required"`
	// Resource type identifier.
	//
	// Any of "owner".
	Object OwnerObject `json:"object" api:"required"`
	// Where this resource came from.
	//
	//   - `system`: a platform-provided default shared across all accounts; not
	//     editable.
	//   - `account`: created and owned by a specific account; the `account` field
	//     identifies which.
	//
	// Any of "system", "account".
	Type OwnerType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account     respjson.Field
		Object      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Owner) RawJSON() string { return r.JSON.raw }
func (r *Owner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type OwnerObject string

const (
	OwnerObjectOwner OwnerObject = "owner"
)

// Where this resource came from.
//
//   - `system`: a platform-provided default shared across all accounts; not
//     editable.
//   - `account`: created and owned by a specific account; the `account` field
//     identifies which.
type OwnerType string

const (
	OwnerTypeSystem  OwnerType = "system"
	OwnerTypeAccount OwnerType = "account"
)

// PageInfo describes where the current page sits within a paginated result set and
// how to move to the adjacent pages.
//
// Page a list by following the URLs below rather than assembling cursors yourself.
// For a top-level list endpoint the URL repeats the original request's query
// string with only the cursor swapped, so following it preserves the same filters,
// search term, and page size.
type PageInfo struct {
	// Whether more results exist after this page.
	HasNextPage bool `json:"has_next_page" api:"required"`
	// Whether results exist before this page.
	HasPrevPage bool `json:"has_prev_page" api:"required"`
	// Relative URL that fetches the next page of results.
	NextPageURL string `json:"next_page_url" api:"required"`
	// Relative URL that fetches the previous page of results.
	PreviousPageURL string `json:"previous_page_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage     respjson.Field
		HasPrevPage     respjson.Field
		NextPageURL     respjson.Field
		PreviousPageURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PageInfo) RawJSON() string { return r.JSON.raw }
func (r *PageInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A named set of permissions that can be assigned to users to control what they
// can access.
type Role struct {
	// Role ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the role.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "role".
	Object RoleObject `json:"object" api:"required"`
	// Owner describes the provenance of a resource.
	Owner Owner `json:"owner" api:"required"`
	// Permissions granted by this role, in `{permission}:{action}` format, such as
	// `customers:read`.
	Permissions []string `json:"permissions" api:"required"`
	// The kind of role.
	//
	// The type gates behavior that individual permissions do not cover, and some
	// actions are reserved for a single role type.
	//
	//   - `admin`: full administrative access. Sensitive areas such as API keys,
	//     billing, and third-party integrations are restricted to admins no matter what
	//     permissions another role holds.
	//   - `user`: a custom role tailored to a specific need, with its permissions
	//     defined explicitly. Roles created through the API always have this type.
	//   - `scanner`: the role used by shop-floor scanning stations, assigned
	//     automatically when a scanning-station user is created.
	//   - `sales_rep`: a role for sales representatives. Order analytics are scoped to
	//     the rep's own orders.
	//   - `agent`: a role assigned to an automated agent rather than a person.
	//
	// Any of "admin", "user", "scanner", "sales_rep", "agent".
	Type RoleType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Owner       respjson.Field
		Permissions respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Role) RawJSON() string { return r.JSON.raw }
func (r *Role) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type RoleObject string

const (
	RoleObjectRole RoleObject = "role"
)

// The kind of role.
//
// The type gates behavior that individual permissions do not cover, and some
// actions are reserved for a single role type.
//
//   - `admin`: full administrative access. Sensitive areas such as API keys,
//     billing, and third-party integrations are restricted to admins no matter what
//     permissions another role holds.
//   - `user`: a custom role tailored to a specific need, with its permissions
//     defined explicitly. Roles created through the API always have this type.
//   - `scanner`: the role used by shop-floor scanning stations, assigned
//     automatically when a scanning-station user is created.
//   - `sales_rep`: a role for sales representatives. Order analytics are scoped to
//     the rep's own orders.
//   - `agent`: a role assigned to an automated agent rather than a person.
type RoleType string

const (
	RoleTypeAdmin    RoleType = "admin"
	RoleTypeUser     RoleType = "user"
	RoleTypeScanner  RoleType = "scanner"
	RoleTypeSalesRep RoleType = "sales_rep"
	RoleTypeAgent    RoleType = "agent"
)

type AuthAPIKeyDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthAPIKeyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthAPIKeyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthAPIKeyNewParams struct {
	// Request to create an API key.
	CreateAPIKeyRequest CreateAPIKeyRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r AuthAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAPIKeyRequest)
}
func (r *AuthAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AuthAPIKeyNewParams]'s query parameters as `url.Values`.
func (r AuthAPIKeyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthAPIKeyGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthAPIKeyGetParams]'s query parameters as `url.Values`.
func (r AuthAPIKeyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuthAPIKeyListParams struct {
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
	// Any of "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// API key statuses to filter by.
	//
	//   - `active`: the key still authenticates requests. A key whose revocation is
	//     scheduled for a future time is still active until that time arrives.
	//   - `expired`: the key passed its expiration time without having been revoked.
	//   - `revoked`: the key was revoked, which takes precedence over expiration.
	//
	// When omitted, keys of every status are returned.
	//
	// Any of "active", "expired", "revoked".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthAPIKeyListParams]'s query parameters as `url.Values`.
func (r AuthAPIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
