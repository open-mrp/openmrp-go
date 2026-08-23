// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/open-mrp/openmrp-go/internal/apijson"
	"github.com/open-mrp/openmrp-go/internal/apiquery"
	"github.com/open-mrp/openmrp-go/internal/requestconfig"
	"github.com/open-mrp/openmrp-go/option"
	"github.com/open-mrp/openmrp-go/packages/param"
	"github.com/open-mrp/openmrp-go/packages/respjson"
)

// List available platform tools for agent configuration.
//
// AIService contains methods and other services that help with interacting with
// the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIService] method instead.
type AIService struct {
	options []option.RequestOption
	// List, create, update, and delete agent definitions.
	Agents AIAgentService
	// List, retrieve, trigger, cancel, and continue agent runs.
	Runs AIRunService
	// List, create, update, and delete agent memories.
	Memories AIMemoryService
}

// NewAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIService(opts ...option.RequestOption) (r AIService) {
	r = AIService{}
	r.options = opts
	r.Agents = NewAIAgentService(opts...)
	r.Runs = NewAIRunService(opts...)
	r.Memories = NewAIMemoryService(opts...)
	return
}

// Returns a paginated list of the groups the agent tool catalog is organized into.
//
// The catalog is platform-defined and identical for every account. Pagination
// applies to the groups themselves, so a group requested with `include=tools`
// always carries its complete tool list regardless of the page limit. The `q`
// search term matches against group names.
//
// This endpoint requires the permission: `agents:read`.
func (r *AIService) GetToolGroups(ctx context.Context, query AIGetToolGroupsParams, opts ...option.RequestOption) (res *ListToolGroup, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/tool-groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a paginated list of every capability that can be granted to an agent.
//
// The catalog is platform-defined and identical for every account, and covers both
// built-in runtime capabilities and the API operations agents are allowed to
// perform. The `q` search term matches against tool names and the name of the
// group a tool belongs to.
//
// This endpoint requires the permission: `agents:read`.
func (r *AIService) GetTools(ctx context.Context, query AIGetToolsParams, opts ...option.RequestOption) (res *ListAvailableTool, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A capability an agent can be granted, allowing it to take that action during a
// run.
//
// The catalog of available tools is the same for every account; granting one to an
// agent is what makes it callable.
type AvailableTool struct {
	// Where the tool's behavior comes from.
	//
	//   - `built_in`: a capability implemented by the agent runtime itself, such as
	//     fetching a web page or drafting a reply for a teammate to approve.
	//   - `api_endpoint`: an operation of this API exposed as a tool, letting the agent
	//     perform it on the account's behalf.
	//
	// Any of "built_in", "api_endpoint".
	Category AvailableToolCategory `json:"category" api:"required"`
	// JSON schema describing the configuration options this tool accepts.
	//
	// Defines the shape of the `config` field on AgentDefinitionTool: a schema
	// declaring a `max_results` integer property means that tool's `config` may set
	// `max_results`. Encoded as a JSON value (object, array, string, number, boolean,
	// or null), not a JSON-encoded string.
	ConfigSchema any `json:"config_schema" api:"required"`
	// Explanation of what the tool does.
	//
	// This is also the description the agent's model reads when deciding whether to
	// call the tool.
	Description string `json:"description" api:"required"`
	// Whether invoking this tool takes an action rather than only reading data.
	//
	// True for any `api_endpoint` tool whose underlying operation is not a read, and
	// for `built_in` tools that do something externally visible or hard to undo, such
	// as sending an email. A mutating `built_in` tool always pauses its run for human
	// approval and that gate cannot be turned off for an individual agent; for
	// `api_endpoint` tools the flag is advisory and review stays configurable per
	// agent.
	Mutating bool `json:"mutating" api:"required"`
	// Human-readable name for the tool.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "available_tool".
	Object AvailableToolObject `json:"object" api:"required"`
	// Permission scopes the agent's role must hold for this tool to be usable (e.g.
	// `products:read`).
	RequiredPermissions []string `json:"required_permissions" api:"required"`
	// Role type the caller must have for this tool, when the operation is gated by
	// role rather than a permission (e.g. `admin`).
	//
	// Any of "admin", "user", "scanner", "sales_rep", "agent".
	RequiredRoleType AvailableToolRequiredRoleType `json:"required_role_type" api:"required"`
	// A stable identifier used when attaching the tool to an agent.
	Slug string `json:"slug" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category            respjson.Field
		ConfigSchema        respjson.Field
		Description         respjson.Field
		Mutating            respjson.Field
		Name                respjson.Field
		Object              respjson.Field
		RequiredPermissions respjson.Field
		RequiredRoleType    respjson.Field
		Slug                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AvailableTool) RawJSON() string { return r.JSON.raw }
func (r *AvailableTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where the tool's behavior comes from.
//
//   - `built_in`: a capability implemented by the agent runtime itself, such as
//     fetching a web page or drafting a reply for a teammate to approve.
//   - `api_endpoint`: an operation of this API exposed as a tool, letting the agent
//     perform it on the account's behalf.
type AvailableToolCategory string

const (
	AvailableToolCategoryBuiltIn     AvailableToolCategory = "built_in"
	AvailableToolCategoryAPIEndpoint AvailableToolCategory = "api_endpoint"
)

// Resource type identifier.
type AvailableToolObject string

const (
	AvailableToolObjectAvailableTool AvailableToolObject = "available_tool"
)

// Role type the caller must have for this tool, when the operation is gated by
// role rather than a permission (e.g. `admin`).
type AvailableToolRequiredRoleType string

const (
	AvailableToolRequiredRoleTypeAdmin    AvailableToolRequiredRoleType = "admin"
	AvailableToolRequiredRoleTypeUser     AvailableToolRequiredRoleType = "user"
	AvailableToolRequiredRoleTypeScanner  AvailableToolRequiredRoleType = "scanner"
	AvailableToolRequiredRoleTypeSalesRep AvailableToolRequiredRoleType = "sales_rep"
	AvailableToolRequiredRoleTypeAgent    AvailableToolRequiredRoleType = "agent"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAvailableTool struct {
	// Resources in this page.
	Data []AvailableTool `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAvailableToolObject `json:"object" api:"required"`
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
func (r ListAvailableTool) RawJSON() string { return r.JSON.raw }
func (r *ListAvailableTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAvailableToolObject string

const (
	ListAvailableToolObjectList ListAvailableToolObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListToolGroup struct {
	// Resources in this page.
	Data []ToolGroup `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListToolGroupObject `json:"object" api:"required"`
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
func (r ListToolGroup) RawJSON() string { return r.JSON.raw }
func (r *ListToolGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListToolGroupObject string

const (
	ListToolGroupObjectList ListToolGroupObject = "list"
)

// A named grouping of the tools that can be granted to an agent, used to organize
// the tool catalog.
type ToolGroup struct {
	// Group ID.
	ID string `json:"id" api:"required"`
	// Description of what the tools in this group do.
	Description string `json:"description" api:"required"`
	// Icon identifier (e.g. a Material Icon name).
	Icon string `json:"icon" api:"required"`
	// Human-readable group name (e.g. `Product Tools`).
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "tool_group".
	Object ToolGroupObject `json:"object" api:"required"`
	// Machine-readable name for the group (e.g. `customer_tools`).
	Slug string `json:"slug" api:"required"`
	// Display sort order, lowest first.
	SortOrder int64 `json:"sort_order" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Tools ListAvailableTool `json:"tools" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Icon        respjson.Field
		Name        respjson.Field
		Object      respjson.Field
		Slug        respjson.Field
		SortOrder   respjson.Field
		Tools       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolGroup) RawJSON() string { return r.JSON.raw }
func (r *ToolGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ToolGroupObject string

const (
	ToolGroupObjectToolGroup ToolGroupObject = "tool_group"
)

type AIGetToolGroupsParams struct {
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
	// Any of "tools".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIGetToolGroupsParams]'s query parameters as `url.Values`.
func (r AIGetToolGroupsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIGetToolsParams struct {
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

// URLQuery serializes [AIGetToolsParams]'s query parameters as `url.Values`.
func (r AIGetToolsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
