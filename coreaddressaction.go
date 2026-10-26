// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// Autocomplete, look up details, and validate addresses.
//
// CoreAddressActionService contains methods and other services that help with
// interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCoreAddressActionService] method instead.
type CoreAddressActionService struct {
	options []option.RequestOption
}

// NewCoreAddressActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCoreAddressActionService(opts ...option.RequestOption) (r CoreAddressActionService) {
	r = CoreAddressActionService{}
	r.options = opts
	return
}

// Checks an address against an address validation service and returns a
// standardized version of it.
//
// Nothing is created or modified. Use this before creating or updating an address
// to confirm it is complete and to pick up corrected values. When the service can
// standardize the address, `formatted_address` and `components` carry the
// corrected values, and `validation_messages` explains anything that was inferred,
// replaced, or could not be confirmed.
func (r *CoreAddressActionService) Validate(ctx context.Context, body CoreAddressActionValidateParams, opts ...option.RequestOption) (res *ValidatedAddress, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/core/addresses/actions/validate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Parsed address components.
type AddressComponents struct {
	// First line of the street address.
	AddressLine1 string `json:"address_line_1" api:"required"`
	// Second line of the street address.
	AddressLine2 string `json:"address_line_2" api:"required"`
	// City or locality.
	City string `json:"city" api:"required"`
	// Country name or code.
	Country string `json:"country" api:"required"`
	// Two-letter country code.
	CountryCode string `json:"country_code" api:"required"`
	// Resource type identifier.
	//
	// Any of "address_components".
	Object AddressComponentsObject `json:"object" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postal_code" api:"required"`
	// State or administrative area.
	State string `json:"state" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressLine1 respjson.Field
		AddressLine2 respjson.Field
		City         respjson.Field
		Country      respjson.Field
		CountryCode  respjson.Field
		Object       respjson.Field
		PostalCode   respjson.Field
		State        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressComponents) RawJSON() string { return r.JSON.raw }
func (r *AddressComponents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AddressComponentsObject string

const (
	AddressComponentsObjectAddressComponents AddressComponentsObject = "address_components"
)

// Request to validate an address.
//
// The properties AddressLine1, City, Country, PostalCode, State are required.
type ValidateAddressRequestParam struct {
	// First line of the street address.
	AddressLine1 string `json:"address_line_1" api:"required"`
	// City or locality.
	City string `json:"city" api:"required"`
	// Two-letter country code, such as `US`.
	//
	// A full country name such as `United States` is recognized for a handful of
	// common countries; send the two-letter code for anywhere else.
	Country string `json:"country" api:"required"`
	// Postal or ZIP code. Empty for an address without one.
	PostalCode string `json:"postal_code" api:"required"`
	// State or administrative area. Empty for an address without one.
	State string `json:"state" api:"required"`
	// Second line of the street address.
	AddressLine2 param.Opt[string] `json:"address_line_2,omitzero"`
	paramObj
}

func (r ValidateAddressRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ValidateAddressRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ValidateAddressRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The outcome of checking a submitted address against an address validation
// service.
type ValidatedAddress struct {
	// Parsed address components.
	Components AddressComponents `json:"components" api:"required"`
	// Formatted, single-line address as standardized by the validation service.
	//
	// The validation service may omit this regardless of `status`, so it can be absent
	// even for a `valid` address.
	FormattedAddress string `json:"formatted_address" api:"required"`
	// Resource type identifier.
	//
	// Any of "validated_address".
	Object ValidatedAddressObject `json:"object" api:"required"`
	// Whether the address was confirmed as complete and specific enough to ship to.
	//
	//   - `valid`: nothing required was missing and the address resolved to a specific
	//     building or block.
	//   - `invalid`: required components were missing, or the address only resolved to a
	//     street or a wider area.
	//
	// When the status is `invalid`, read `validation_messages` and compare
	// `components` against what you submitted to see what to correct.
	//
	// Any of "valid", "invalid".
	Status ValidatedAddressStatus `json:"status" api:"required"`
	// Human-readable messages describing issues found during validation.
	//
	// May be non-empty even when `status` is `valid`, for example when components were
	// inferred or replaced with standardized values. Empty when no issues were
	// reported.
	ValidationMessages []string `json:"validation_messages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Components         respjson.Field
		FormattedAddress   respjson.Field
		Object             respjson.Field
		Status             respjson.Field
		ValidationMessages respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValidatedAddress) RawJSON() string { return r.JSON.raw }
func (r *ValidatedAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ValidatedAddressObject string

const (
	ValidatedAddressObjectValidatedAddress ValidatedAddressObject = "validated_address"
)

// Whether the address was confirmed as complete and specific enough to ship to.
//
//   - `valid`: nothing required was missing and the address resolved to a specific
//     building or block.
//   - `invalid`: required components were missing, or the address only resolved to a
//     street or a wider area.
//
// When the status is `invalid`, read `validation_messages` and compare
// `components` against what you submitted to see what to correct.
type ValidatedAddressStatus string

const (
	ValidatedAddressStatusValid   ValidatedAddressStatus = "valid"
	ValidatedAddressStatusInvalid ValidatedAddressStatus = "invalid"
)

type CoreAddressActionValidateParams struct {
	// Request to validate an address.
	ValidateAddressRequest ValidateAddressRequestParam
	paramObj
}

func (r CoreAddressActionValidateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ValidateAddressRequest)
}
func (r *CoreAddressActionValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
