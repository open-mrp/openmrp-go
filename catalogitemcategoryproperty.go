// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	shimjson "github.com/open-mrp/openmrp-go/internal/encoding/json"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List and manage item categories.
//
// CatalogItemCategoryPropertyService contains methods and other services that help
// with interacting with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCatalogItemCategoryPropertyService] method instead.
type CatalogItemCategoryPropertyService struct {
	options []option.RequestOption
}

// NewCatalogItemCategoryPropertyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCatalogItemCategoryPropertyService(opts ...option.RequestOption) (r CatalogItemCategoryPropertyService) {
	r = CatalogItemCategoryPropertyService{}
	r.options = opts
	return
}

// Creates a property and attaches it to an item category, returning the new
// property.
//
// The property is one of your account's properties like any other, starting with
// no attributes, and the category carries it from the moment it exists. Both
// happen in one request that needs only permission to update the category. A name
// already used by one of your account's properties returns a conflict error naming
// `name`.
//
// This endpoint requires the permission: `item_categories:update`.
func (r *CatalogItemCategoryPropertyService) New(ctx context.Context, id string, params CatalogItemCategoryPropertyNewParams, opts ...option.RequestOption) (res *Property, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s/properties", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Attaches one of your account's properties to an item category.
//
// The property then appears among the category's properties, including in the
// customer-facing catalog, describing a dimension along which the category's items
// vary. Each property name can appear only once per category, so attaching a
// property whose name duplicates one already there returns a conflict error.
//
// This endpoint requires the permission: `item_categories:update`.
func (r *CatalogItemCategoryPropertyService) Update(ctx context.Context, propertyID string, body CatalogItemCategoryPropertyUpdateParams, opts ...option.RequestOption) (res *CatalogItemCategoryPropertyUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if propertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s/properties/%s", url.PathEscape(body.ID), url.PathEscape(propertyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// Detaches a property from an item category.
//
// Only the link between the property and the category is removed; the property
// itself and its attributes are left intact and stay available to other
// categories. The property must belong to your account.
//
// This endpoint requires the permission: `item_categories:update`.
func (r *CatalogItemCategoryPropertyService) Delete(ctx context.Context, propertyID string, body CatalogItemCategoryPropertyDeleteParams, opts ...option.RequestOption) (res *CatalogItemCategoryPropertyDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if propertyID == "" {
		err = errors.New("missing required property_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/catalog/item-categories/%s/properties/%s", url.PathEscape(body.ID), url.PathEscape(propertyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Request to create a property on an item category.
//
// The property Name is required.
type CreateItemCategoryPropertyRequestParam struct {
	// Display name of the new property, such as `Color` or `Size`.
	//
	// Must be unique within your account. To attach a property that already exists,
	// use the add item category property endpoint.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r CreateItemCategoryPropertyRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateItemCategoryPropertyRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateItemCategoryPropertyRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryPropertyUpdateResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemCategoryPropertyUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemCategoryPropertyUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryPropertyDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CatalogItemCategoryPropertyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CatalogItemCategoryPropertyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CatalogItemCategoryPropertyNewParams struct {
	// Request to create a property on an item category.
	CreateItemCategoryPropertyRequest CreateItemCategoryPropertyRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "attributes".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r CatalogItemCategoryPropertyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateItemCategoryPropertyRequest)
}
func (r *CatalogItemCategoryPropertyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CatalogItemCategoryPropertyNewParams]'s query parameters as
// `url.Values`.
func (r CatalogItemCategoryPropertyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CatalogItemCategoryPropertyUpdateParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}

type CatalogItemCategoryPropertyDeleteParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
