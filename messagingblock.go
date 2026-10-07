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

// Block and unblock users from direct messaging.
//
// MessagingBlockService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingBlockService] method instead.
type MessagingBlockService struct {
	options []option.RequestOption
}

// NewMessagingBlockService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessagingBlockService(opts ...option.RequestOption) (r MessagingBlockService) {
	r = MessagingBlockService{}
	r.options = opts
	return
}

// Blocks another user in your account from exchanging direct messages with you.
//
// While the block stands neither of you can start a direct message with the other
// or post in one you already share; group conversations and customer cases are
// unaffected. Blocking someone you have already blocked returns the original block
// instead of creating a second one.
//
// This endpoint requires the permission: `messaging:create`.
func (r *MessagingBlockService) New(ctx context.Context, params MessagingBlockNewParams, opts ...option.RequestOption) (res *MessagingBlock, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/blocks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Lists the users you have blocked, most recently blocked first.
//
// Only blocks you created are returned — you are never told who has blocked you.
//
// This endpoint requires the permission: `messaging:read`.
func (r *MessagingBlockService) List(ctx context.Context, query MessagingBlockListParams, opts ...option.RequestOption) (res *ListMessagingBlock, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/messaging/blocks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lifts a block you placed on another user, letting the two of you message each
// other again.
//
// Only your own block is removed: if the other person has also blocked you, direct
// messages between you stay blocked. Unblocking someone you have not blocked
// succeeds and changes nothing.
//
// This endpoint requires the permission: `messaging:delete`.
func (r *MessagingBlockService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingBlockDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/messaging/blocks/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A user's membership in an account, carrying the account-specific status, role,
// and department.
//
// Profile fields (name, email, username, image URL) live on the `user`
// sub-resource, which is shared across every account the user belongs to.
type AccountUser struct {
	// Account user ID.
	ID string `json:"id" api:"required"`
	// When the account user was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department Department `json:"department" api:"required"`
	// Whether this user can be assigned as a sales representative on orders,
	// territories, and targets.
	//
	// Independent of the `sales_rep` role type, which still scopes analytics and hides
	// cost. Users with the `sales_rep` role are always eligible.
	//
	// Null to customer and supplier portal users, like the rest of your commission
	// settings.
	IsCommissionEligible bool `json:"is_commission_eligible" api:"required"`
	// When the user last accessed this account.
	//
	// Null to customer and supplier portal users: it is your own team's activity.
	LastUsedAt time.Time `json:"last_used_at" api:"required" format:"date-time"`
	// The notifications you send this user about the customer or supplier account they
	// belong to.
	//
	// Set when the account you are acting in is a customer or supplier account you
	// manage, and an empty list when none are enabled; turn them on and off with
	// `preferences` on create and update. `null` for your own account's users, and
	// where an account user is embedded in another resource.
	//
	//   - `order_acknowledgement`: the confirmation email sent when an order is placed
	//     for the customer.
	//   - `invoice`: invoice emails for the customer's orders.
	//   - `purchase_order_submission`: a copy of each purchase order you submit to the
	//     supplier.
	//
	// Any of "invoice", "order_acknowledgement", "purchase_order_submission".
	NotificationTypes []string `json:"notification_types" api:"required"`
	// Resource type identifier.
	//
	// Any of "account_user".
	Object AccountUserObject `json:"object" api:"required"`
	// A named set of permissions that can be assigned to users to control what they
	// can access.
	Role Role `json:"role" api:"required"`
	// The current state of this user's membership in the account.
	//
	//   - `active`: the user can sign in to the account and occupies one of the plan's
	//     seats.
	//   - `disabled`: the user is locked out of the account and their sessions have been
	//     revoked, but the membership is retained.
	//   - `removed`: the membership has been soft-deleted; it is hidden from listings by
	//     default and can be restored with the activate action.
	//
	// Any of "active", "disabled", "removed".
	Status AccountUserStatus `json:"status" api:"required"`
	// When the account user was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A user's global profile, shared across every account they belong to.
	//
	// Account-specific settings (status, role, department) live on the account user
	// resource that links the user to each account.
	User User `json:"user" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		Department           respjson.Field
		IsCommissionEligible respjson.Field
		LastUsedAt           respjson.Field
		NotificationTypes    respjson.Field
		Object               respjson.Field
		Role                 respjson.Field
		Status               respjson.Field
		UpdatedAt            respjson.Field
		User                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountUser) RawJSON() string { return r.JSON.raw }
func (r *AccountUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AccountUserObject string

const (
	AccountUserObjectAccountUser AccountUserObject = "account_user"
)

// The current state of this user's membership in the account.
//
//   - `active`: the user can sign in to the account and occupies one of the plan's
//     seats.
//   - `disabled`: the user is locked out of the account and their sessions have been
//     revoked, but the membership is retained.
//   - `removed`: the membership has been soft-deleted; it is hidden from listings by
//     default and can be restored with the activate action.
type AccountUserStatus string

const (
	AccountUserStatusActive   AccountUserStatus = "active"
	AccountUserStatusDisabled AccountUserStatus = "disabled"
	AccountUserStatusRemoved  AccountUserStatus = "removed"
)

// Request to block another account user from messaging the caller.
//
// The property BlockedAccountUserID is required.
type BlockRequestParam struct {
	// The account user to block.
	//
	// It must be someone else in your account; you cannot block yourself.
	BlockedAccountUserID string `json:"blocked_account_user_id" api:"required"`
	paramObj
}

func (r BlockRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BlockRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlockRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Material consumed by a production step.
//
// Each consumption records one input item and how much of it the step uses.
// Consumptions also determine the production flow: when another step produces the
// consumed item, the two steps are linked upstream/downstream automatically.
//
// The quantities are stated against the step's own output, so a step producing 100
// pairs and consuming 5 kg of yarn needs 5 kg per 100 pairs. Material requirements
// for an order scale every consumption in the flow by how much of the finished
// item is wanted.
type Consumption struct {
	// Consumption ID.
	ID string `json:"id" api:"required"`
	// An entry in your catalog: something you sell, consume, or build with.
	ConsumedItem Item `json:"consumed_item" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Instructions for how this material is consumed.
	Instructions string `json:"instructions" api:"required"`
	// Resource type identifier.
	//
	// Any of "consumption".
	Object ConsumptionObject `json:"object" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	Quantity Quantity `json:"quantity" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	WasteQuantity Quantity `json:"waste_quantity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ConsumedItem  respjson.Field
		CreatedAt     respjson.Field
		Instructions  respjson.Field
		Object        respjson.Field
		Quantity      respjson.Field
		UpdatedAt     respjson.Field
		WasteQuantity respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Consumption) RawJSON() string { return r.JSON.raw }
func (r *Consumption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ConsumptionObject string

const (
	ConsumptionObjectConsumption ConsumptionObject = "consumption"
)

// A functional area of a production operation, such as fabrication or packaging,
// that groups scanning stations and machines.
type Department struct {
	// Department ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	LaborRate Rate `json:"labor_rate" api:"required"`
	// A physical storage location, such as a warehouse, aisle, or bin, arranged in a
	// parent-child hierarchy.
	Location Location `json:"location" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Machines *ListMachine `json:"machines" api:"required"`
	// Display name of the department.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the department.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "department".
	Object DepartmentObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ScanningStations *ListScanningStation `json:"scanning_stations" api:"required"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		LaborRate        respjson.Field
		Location         respjson.Field
		Machines         respjson.Field
		Name             respjson.Field
		Notes            respjson.Field
		Object           respjson.Field
		ScanningStations respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Department) RawJSON() string { return r.JSON.raw }
func (r *Department) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type DepartmentObject string

const (
	DepartmentObjectDepartment DepartmentObject = "department"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListConsumption struct {
	// Resources in this page.
	Data []Consumption `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListConsumptionObject `json:"object" api:"required"`
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
func (r ListConsumption) RawJSON() string { return r.JSON.raw }
func (r *ListConsumption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListConsumptionObject string

const (
	ListConsumptionObjectList ListConsumptionObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListLocation struct {
	// Resources in this page.
	Data []Location `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListLocationObject `json:"object" api:"required"`
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
func (r ListLocation) RawJSON() string { return r.JSON.raw }
func (r *ListLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListLocationObject string

const (
	ListLocationObjectList ListLocationObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMachine struct {
	// Resources in this page.
	Data []Machine `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMachineObject `json:"object" api:"required"`
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
func (r ListMachine) RawJSON() string { return r.JSON.raw }
func (r *ListMachine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMachineObject string

const (
	ListMachineObjectList ListMachineObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListMessagingBlock struct {
	// Resources in this page.
	Data []MessagingBlock `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListMessagingBlockObject `json:"object" api:"required"`
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
func (r ListMessagingBlock) RawJSON() string { return r.JSON.raw }
func (r *ListMessagingBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListMessagingBlockObject string

const (
	ListMessagingBlockObjectList ListMessagingBlockObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListProductionStep struct {
	// Resources in this page.
	Data []ProductionStep `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListProductionStepObject `json:"object" api:"required"`
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
func (r ListProductionStep) RawJSON() string { return r.JSON.raw }
func (r *ListProductionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListProductionStepObject string

const (
	ListProductionStepObjectList ListProductionStepObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListScanningStation struct {
	// Resources in this page.
	Data []ScanningStation `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListScanningStationObject `json:"object" api:"required"`
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
func (r ListScanningStation) RawJSON() string { return r.JSON.raw }
func (r *ListScanningStation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListScanningStationObject string

const (
	ListScanningStationObjectList ListScanningStationObject = "list"
)

// A physical storage location, such as a warehouse, aisle, or bin, arranged in a
// parent-child hierarchy.
type Location struct {
	// Location ID.
	ID string `json:"id" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Children *ListLocation `json:"children" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Display name of the location.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "location".
	Object LocationObject `json:"object" api:"required"`
	// A physical storage location, such as a warehouse, aisle, or bin, arranged in a
	// parent-child hierarchy.
	Parent *Location `json:"parent" api:"required"`
	// This location's level in the storage hierarchy.
	//
	// The levels run from largest to smallest: `building`, `section`, `aisle`, `rack`,
	// `shelf`, `bin`. They are descriptive labels rather than a rule — a location's
	// parent is not required to be the next level up.
	//
	// Any of "building", "section", "aisle", "rack", "shelf", "bin".
	Type LocationTypeCode `json:"type" api:"required"`
	// Last-updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Parent      respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type LocationObject string

const (
	LocationObjectLocation LocationObject = "location"
)

type LocationTypeCode string

const (
	LocationTypeCodeBuilding LocationTypeCode = "building"
	LocationTypeCodeSection  LocationTypeCode = "section"
	LocationTypeCodeAisle    LocationTypeCode = "aisle"
	LocationTypeCodeRack     LocationTypeCode = "rack"
	LocationTypeCodeShelf    LocationTypeCode = "shelf"
	LocationTypeCodeBin      LocationTypeCode = "bin"
)

// A piece of production equipment, such as a CNC router or press, assigned to a
// department.
type Machine struct {
	// Machine ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department *Department `json:"department" api:"required"`
	// Display name of the machine.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the machine.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "machine".
	Object MachineObject `json:"object" api:"required"`
	// Serial number of the machine.
	SerialNumber string `json:"serial_number" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Department   respjson.Field
		Name         respjson.Field
		Notes        respjson.Field
		Object       respjson.Field
		SerialNumber respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Machine) RawJSON() string { return r.JSON.raw }
func (r *Machine) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MachineObject string

const (
	MachineObjectMachine MachineObject = "machine"
)

// A block one account user has placed on another.
//
// While the block stands, neither of the two can start a direct message with the
// other or post in an existing one, whichever of them created it. Group
// conversations and customer cases are unaffected.
type MessagingBlock struct {
	// Block ID.
	ID string `json:"id" api:"required"`
	// A user's membership in an account, carrying the account-specific status, role,
	// and department.
	//
	// Profile fields (name, email, username, image URL) live on the `user`
	// sub-resource, which is shared across every account the user belongs to.
	BlockedUser AccountUser `json:"blocked_user" api:"required"`
	// When the block was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "messaging_block".
	Object MessagingBlockObject `json:"object" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		BlockedUser respjson.Field
		CreatedAt   respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingBlock) RawJSON() string { return r.JSON.raw }
func (r *MessagingBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type MessagingBlockObject string

const (
	MessagingBlockObjectMessagingBlock MessagingBlockObject = "messaging_block"
)

// The output of a production step: the item it produces and the quantity produced.
type ProductionOutput struct {
	// Production ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Resource type identifier.
	//
	// Any of "production".
	Object ProductionOutputObject `json:"object" api:"required"`
	// An entry in your catalog: something you sell, consume, or build with.
	ProducedItem Item `json:"produced_item" api:"required"`
	// A measured amount: a numeric value together with the unit it is expressed in.
	//
	// Quantities are shared building blocks rather than standalone records — other
	// resources point at them to report stock levels, ordered and packed amounts,
	// money, weights, and durations.
	Quantity Quantity `json:"quantity" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Object       respjson.Field
		ProducedItem respjson.Field
		Quantity     respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionOutput) RawJSON() string { return r.JSON.raw }
func (r *ProductionOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionOutputObject string

const (
	ProductionOutputObjectProduction ProductionOutputObject = "production"
)

// A single stage of work in an item's production flow, with its output, material
// inputs, cost rates, and graph connections.
type ProductionStep struct {
	// Production step ID.
	ID string `json:"id" api:"required"`
	// Allowance correction factor applied to labor time in cost calculations, as a
	// decimal string.
	//
	// Effective labor time per unit is
	// `labor_time × (1 + leveling_factor) × (1 + allowances)`.
	Allowances string `json:"allowances" api:"required" format:"decimal"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Consumptions ListConsumption `json:"consumptions" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department *Department `json:"department" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	InSteps *ListProductionStep `json:"in_steps" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	LaborRate Rate `json:"labor_rate" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	LaborTime Rate `json:"labor_time" api:"required"`
	// Leveling correction factor applied to labor time in cost calculations, as a
	// decimal string.
	//
	// Effective labor time per unit is
	// `labor_time × (1 + leveling_factor) × (1 + allowances)`.
	LevelingFactor string `json:"leveling_factor" api:"required" format:"decimal"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Machines *ListMachine `json:"machines" api:"required"`
	// Display name of the step.
	Name string `json:"name" api:"required"`
	// Free-form notes about the step.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "production_step".
	Object ProductionStepObject `json:"object" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	OutSteps *ListProductionStep `json:"out_steps" api:"required"`
	// Value expressed as a ratio of two units, such as a price per kilogram or a
	// throughput per hour.
	OverheadRate Rate `json:"overhead_rate" api:"required"`
	// The output of a production step: the item it produces and the quantity produced.
	Production ProductionOutput `json:"production" api:"required"`
	// A station on the production floor where operators scan batches to perform a
	// batch operation, such as initializing or moving a batch.
	ScanningStation *ScanningStation `json:"scanning_station" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Allowances      respjson.Field
		Consumptions    respjson.Field
		CreatedAt       respjson.Field
		Department      respjson.Field
		InSteps         respjson.Field
		LaborRate       respjson.Field
		LaborTime       respjson.Field
		LevelingFactor  respjson.Field
		Machines        respjson.Field
		Name            respjson.Field
		Notes           respjson.Field
		Object          respjson.Field
		OutSteps        respjson.Field
		OverheadRate    respjson.Field
		Production      respjson.Field
		ScanningStation respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProductionStep) RawJSON() string { return r.JSON.raw }
func (r *ProductionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ProductionStepObject string

const (
	ProductionStepObjectProductionStep ProductionStepObject = "production_step"
)

// A station on the production floor where operators scan batches to perform a
// batch operation, such as initializing or moving a batch.
type ScanningStation struct {
	// Scanning station ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// A functional area of a production operation, such as fabrication or packaging,
	// that groups scanning stations and machines.
	Department *Department `json:"department" api:"required"`
	// Size of the labels printed at this station, given as width-by-height (for
	// example, `1x1`).
	//
	// Any of "1x1", "1x3", "1x4", "2x4".
	LabelSize ScanningStationLabelSize `json:"label_size" api:"required"`
	// Type of label printed at this station.
	//
	//   - `tag`: a label attached to the physical product.
	//   - `traveler`: a routing sheet that accompanies the batch through every
	//     production step.
	//
	// Any of "tag", "traveler".
	LabelType ScanningStationLabelType `json:"label_type" api:"required"`
	// Display name of the scanning station.
	//
	// Unique within the account.
	Name string `json:"name" api:"required"`
	// Free-form notes about the scanning station.
	Notes string `json:"notes" api:"required"`
	// Resource type identifier.
	//
	// Any of "scanning_station".
	Object ScanningStationObject `json:"object" api:"required"`
	// Whether operators must perform a material check at this station.
	//
	// - `none`: no additional operator check is required.
	// - `material_check`: a material check is expected before the operation.
	//
	// Any of "none", "material_check".
	OperatorRequirement ScanningStationOperatorRequirement `json:"operator_requirement" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	ProductionSteps *ListProductionStep `json:"production_steps" api:"required"`
	// Scanning station type, determining which batch operation an operator performs
	// when they scan here.
	//
	//   - `init_batch`: starts a new batch at the beginning of a production flow.
	//   - `merge_batch`: combines several scanned batches into one.
	//   - `move_batch`: advances a batch through a production step connected to this
	//     station.
	//   - `split_batch`: divides a batch into several batches.
	//
	// Fixed when the station is created.
	//
	// Any of "init_batch", "merge_batch", "move_batch", "split_batch".
	Type ScanningStationType `json:"type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Department          respjson.Field
		LabelSize           respjson.Field
		LabelType           respjson.Field
		Name                respjson.Field
		Notes               respjson.Field
		Object              respjson.Field
		OperatorRequirement respjson.Field
		ProductionSteps     respjson.Field
		Type                respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScanningStation) RawJSON() string { return r.JSON.raw }
func (r *ScanningStation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Size of the labels printed at this station, given as width-by-height (for
// example, `1x1`).
type ScanningStationLabelSize string

const (
	ScanningStationLabelSize1x1 ScanningStationLabelSize = "1x1"
	ScanningStationLabelSize1x3 ScanningStationLabelSize = "1x3"
	ScanningStationLabelSize1x4 ScanningStationLabelSize = "1x4"
	ScanningStationLabelSize2x4 ScanningStationLabelSize = "2x4"
)

// Type of label printed at this station.
//
//   - `tag`: a label attached to the physical product.
//   - `traveler`: a routing sheet that accompanies the batch through every
//     production step.
type ScanningStationLabelType string

const (
	ScanningStationLabelTypeTag      ScanningStationLabelType = "tag"
	ScanningStationLabelTypeTraveler ScanningStationLabelType = "traveler"
)

// Resource type identifier.
type ScanningStationObject string

const (
	ScanningStationObjectScanningStation ScanningStationObject = "scanning_station"
)

// Whether operators must perform a material check at this station.
//
// - `none`: no additional operator check is required.
// - `material_check`: a material check is expected before the operation.
type ScanningStationOperatorRequirement string

const (
	ScanningStationOperatorRequirementNone          ScanningStationOperatorRequirement = "none"
	ScanningStationOperatorRequirementMaterialCheck ScanningStationOperatorRequirement = "material_check"
)

// Scanning station type, determining which batch operation an operator performs
// when they scan here.
//
//   - `init_batch`: starts a new batch at the beginning of a production flow.
//   - `merge_batch`: combines several scanned batches into one.
//   - `move_batch`: advances a batch through a production step connected to this
//     station.
//   - `split_batch`: divides a batch into several batches.
//
// Fixed when the station is created.
type ScanningStationType string

const (
	ScanningStationTypeInitBatch  ScanningStationType = "init_batch"
	ScanningStationTypeMergeBatch ScanningStationType = "merge_batch"
	ScanningStationTypeMoveBatch  ScanningStationType = "move_batch"
	ScanningStationTypeSplitBatch ScanningStationType = "split_batch"
)

// A user's global profile, shared across every account they belong to.
//
// Account-specific settings (status, role, department) live on the account user
// resource that links the user to each account.
type User struct {
	// User ID.
	ID string `json:"id" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Email address the user signs in with and receives platform email at.
	Email string `json:"email" api:"required"`
	// When the user verified their email address.
	EmailVerifiedAt time.Time `json:"email_verified_at" api:"required" format:"date-time"`
	// Location of the user's profile image.
	//
	// For photos uploaded through the API this holds an internal path rather than a
	// fetchable image URL; call Get User Photo URL to obtain a temporary link to the
	// image itself.
	ImageURL string `json:"image_url" api:"required"`
	// User's full display name.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "user".
	Object UserObject `json:"object" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Username the user can sign in with instead of their email address.
	//
	// Usernames are unique across the whole platform, not just within your account.
	Username string `json:"username" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Email           respjson.Field
		EmailVerifiedAt respjson.Field
		ImageURL        respjson.Field
		Name            respjson.Field
		Object          respjson.Field
		UpdatedAt       respjson.Field
		Username        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type UserObject string

const (
	UserObjectUser UserObject = "user"
)

type MessagingBlockDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingBlockDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingBlockDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingBlockNewParams struct {
	// Request to block another account user from messaging the caller.
	BlockRequest BlockRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "blocked_user", "blocked_user.user", "blocked_user.role",
	// "blocked_user.department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r MessagingBlockNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BlockRequest)
}
func (r *MessagingBlockNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MessagingBlockNewParams]'s query parameters as
// `url.Values`.
func (r MessagingBlockNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MessagingBlockListParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "blocked_user", "blocked_user.user", "blocked_user.role",
	// "blocked_user.department".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingBlockListParams]'s query parameters as
// `url.Values`.
func (r MessagingBlockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
