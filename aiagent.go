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

// List, create, update, and delete agent definitions.
//
// AIAgentService contains methods and other services that help with interacting
// with the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAgentService] method instead.
type AIAgentService struct {
	options []option.RequestOption
}

// NewAIAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIAgentService(opts ...option.RequestOption) (r AIAgentService) {
	r = AIAgentService{}
	r.options = opts
	return
}

// Creates a custom agent for your account.
//
// The new agent is a `custom` definition and is immediately `active`, so it can
// start running as soon as it has a role.
//
// This endpoint requires the permission: `agents:create`.
func (r *AIAgentService) New(ctx context.Context, params AIAgentNewParams, opts ...option.RequestOption) (res *AgentDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single agent by ID.
//
// Resolves both the `system` agents OpenMRP provides and the `custom` agents in
// your account; the `status` reflects whether the agent is enabled for your
// account specifically.
//
// This endpoint requires the permission: `agents:read`.
func (r *AIAgentService) Get(ctx context.Context, id string, query AIAgentGetParams, opts ...option.RequestOption) (res *AgentDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/agents/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a custom agent.
//
// Only the fields provided in the request are changed. OpenMRP's `system` agents
// cannot be edited — the only thing you can change about them is whether they are
// enabled for your account, with the Update Agent Status endpoint.
//
// This endpoint requires the permission: `agents:update`.
func (r *AIAgentService) Update(ctx context.Context, id string, params AIAgentUpdateParams, opts ...option.RequestOption) (res *AgentDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/agents/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the agents available to your account, newest first.
//
// Covers both the `system` agents OpenMRP provides to every account and the
// `custom` agents created in yours. Deleted agents are never returned. The `q`
// parameter matches an agent's name, slug, description, or ID.
//
// This endpoint requires the permission: `agents:read`.
func (r *AIAgentService) List(ctx context.Context, query AIAgentListParams, opts ...option.RequestOption) (res *ListAgentDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes a custom agent.
//
// The agent is withdrawn from the API: it stops appearing in listings, no longer
// resolves by ID, and can no longer be run or modified. Runs it already produced
// are kept. OpenMRP's `system` agents cannot be deleted — disable one for your
// account with the Update Agent Status endpoint instead.
//
// This endpoint requires the permission: `agents:delete`.
func (r *AIAgentService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AIAgentDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/agents/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Enables or disables an agent for your account.
//
// Activation is per-account, so this works for the `system` agents OpenMRP shares
// across accounts as well as your own `custom` agents: disabling one here leaves
// the underlying agent untouched for everyone else. Triggering an inactive agent
// returns a validation error.
//
// This endpoint requires the permission: `agents:update`.
func (r *AIAgentService) UpdateStatus(ctx context.Context, id string, params AIAgentUpdateStatusParams, opts ...option.RequestOption) (res *AgentDefinition, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/agents/%s/status", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// An AI agent available to the account.
//
// The definition describes what the agent does, how its runs are triggered, the
// tools it can use, and whether it is currently enabled for the account.
type AgentDefinition struct {
	// Agent definition ID.
	ID string `json:"id" api:"required"`
	// Category grouping for the agent (e.g. `order_processing`), used to organize
	// agents in the UI.
	CategoryCode string `json:"category_code" api:"required"`
	// Agent-level configuration controlling LLM behavior and trigger settings.
	//
	// Distinct from per-tool configuration (`tools[].config`), which configures
	// individual tools attached to the agent.
	Config AgentDefinitionConfig `json:"config" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether the agent is provided by OpenMRP or created in this account.
	//
	// - `system`: provided by OpenMRP; cannot be edited or deleted.
	// - `custom`: created by a user in this account.
	//
	// Any of "system", "custom".
	DefinitionType AgentDefinitionDefinitionType `json:"definition_type" api:"required"`
	// Description of what the agent does.
	Description string `json:"description" api:"required"`
	// Whether this agent definition can be edited.
	//
	// Always `read_only` for `system` definitions.
	//
	// Any of "editable", "read_only".
	Editability AgentDefinitionEditability `json:"editability" api:"required"`
	// Human-readable name of the agent.
	Name string `json:"name" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_definition".
	Object AgentDefinitionObject `json:"object" api:"required"`
	// A named set of permissions that can be assigned to users to control what they
	// can access.
	Role Role `json:"role" api:"required"`
	// URL-friendly identifier for the agent.
	//
	// Unique within the account.
	Slug string `json:"slug" api:"required"`
	// Whether this agent is enabled for the current account.
	//
	// Activation is per-account: a `system` agent shared across accounts can be
	// `active` for one account and `inactive` for another. An `inactive` agent cannot
	// be triggered.
	//
	// Any of "active", "inactive".
	Status AgentDefinitionStatus `json:"status" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Tools ListAgentDefinitionTool `json:"tools" api:"required"`
	// How runs of this agent are initiated.
	//
	//   - `scheduled`: runs on a cron schedule (see
	//     `config.trigger_config.cron_schedule`).
	//   - `event`: runs in response to platform events (see
	//     `config.trigger_config.event_filters`).
	//   - `manual`: runs only when explicitly invoked.
	//   - `chat`: runs in response to a chat message; the run is linked to a
	//     conversation and posts its reply back into it.
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerType AgentDefinitionTriggerType `json:"trigger_type" api:"required"`
	// Last updated timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CategoryCode   respjson.Field
		Config         respjson.Field
		CreatedAt      respjson.Field
		DefinitionType respjson.Field
		Description    respjson.Field
		Editability    respjson.Field
		Name           respjson.Field
		Object         respjson.Field
		Role           respjson.Field
		Slug           respjson.Field
		Status         respjson.Field
		Tools          respjson.Field
		TriggerType    respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDefinition) RawJSON() string { return r.JSON.raw }
func (r *AgentDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the agent is provided by OpenMRP or created in this account.
//
// - `system`: provided by OpenMRP; cannot be edited or deleted.
// - `custom`: created by a user in this account.
type AgentDefinitionDefinitionType string

const (
	AgentDefinitionDefinitionTypeSystem AgentDefinitionDefinitionType = "system"
	AgentDefinitionDefinitionTypeCustom AgentDefinitionDefinitionType = "custom"
)

// Whether this agent definition can be edited.
//
// Always `read_only` for `system` definitions.
type AgentDefinitionEditability string

const (
	AgentDefinitionEditabilityEditable AgentDefinitionEditability = "editable"
	AgentDefinitionEditabilityReadOnly AgentDefinitionEditability = "read_only"
)

// Resource type identifier.
type AgentDefinitionObject string

const (
	AgentDefinitionObjectAgentDefinition AgentDefinitionObject = "agent_definition"
)

// Whether this agent is enabled for the current account.
//
// Activation is per-account: a `system` agent shared across accounts can be
// `active` for one account and `inactive` for another. An `inactive` agent cannot
// be triggered.
type AgentDefinitionStatus string

const (
	AgentDefinitionStatusActive   AgentDefinitionStatus = "active"
	AgentDefinitionStatusInactive AgentDefinitionStatus = "inactive"
)

// How runs of this agent are initiated.
//
//   - `scheduled`: runs on a cron schedule (see
//     `config.trigger_config.cron_schedule`).
//   - `event`: runs in response to platform events (see
//     `config.trigger_config.event_filters`).
//   - `manual`: runs only when explicitly invoked.
//   - `chat`: runs in response to a chat message; the run is linked to a
//     conversation and posts its reply back into it.
type AgentDefinitionTriggerType string

const (
	AgentDefinitionTriggerTypeScheduled AgentDefinitionTriggerType = "scheduled"
	AgentDefinitionTriggerTypeManual    AgentDefinitionTriggerType = "manual"
	AgentDefinitionTriggerTypeEvent     AgentDefinitionTriggerType = "event"
	AgentDefinitionTriggerTypeChat      AgentDefinitionTriggerType = "chat"
)

// Agent-level configuration controlling LLM behavior and trigger settings.
//
// Distinct from per-tool configuration (`tools[].config`), which configures
// individual tools attached to the agent.
type AgentDefinitionConfig struct {
	// Per-endpoint-tool human-review overrides, keyed by tool slug.
	//
	// When an entry is `true`, the run pauses in `awaiting_approval` each time the
	// agent calls that endpoint-tool until it is approved via the Continue Agent Run
	// endpoint. Slugs absent from the map do not require review.
	EndpointToolReview map[string]bool `json:"endpoint_tool_review" api:"required"`
	// API-endpoint tools the agent may discover and use, by slug (e.g.
	// `create_account_group`).
	//
	// These correspond to tools listed by the List Tools endpoint with category
	// `api_endpoint`. A single entry `*` grants the entire endpoint-tool catalog.
	EndpointToolSlugs []string `json:"endpoint_tool_slugs" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_definition_config".
	Object AgentDefinitionConfigObject `json:"object" api:"required"`
	// Standing instructions that define the agent's role and how it should behave on
	// every run.
	SystemPrompt string `json:"system_prompt" api:"required"`
	// LLM sampling temperature between 0 and 1.
	//
	// Lower values make the agent's output more repeatable and literal; higher values
	// make it more varied.
	Temperature float64 `json:"temperature" api:"required"`
	// Intelligence and cost tier for the agent's reasoning.
	//
	// Selects how capable and expensive a model the agent uses without pinning a
	// specific model; higher tiers reason better but cost more. Each tier resolves to
	// an ordered chain of equivalent models, so a run automatically fails over to
	// another provider's model if the preferred one is unavailable.
	//
	//   - `frontier`: the most capable tier, for multi-step planning, ambiguous agent
	//     work, and hard coding or architecture tasks.
	//   - `high`: for normal planning, code edits, synthesis, and customer-facing
	//     reasoning.
	//   - `balanced`: for research, summarization, classification, structured
	//     extraction, and light tool use.
	//   - `cheap`: for simple transforms, validation, formatting, and routing.
	//   - `legacy`: older-generation models kept for compatibility and regression
	//     comparison; avoid unless you specifically need them.
	//
	// Leaving the tier unset picks one from how the agent is triggered: chat and
	// manual runs use `high`, while scheduled and event-driven runs use `balanced` so
	// background work stays cheap.
	//
	// Any of "frontier", "high", "balanced", "cheap", "legacy".
	Tier AgentDefinitionConfigTier `json:"tier" api:"required"`
	// Trigger-type-specific configuration.
	//
	// Which fields are populated depends on the agent's `trigger_type`:
	//
	// - `scheduled`: `cron_schedule` (and optionally `timezone`) is set.
	// - `event`: `event_filters` is set.
	// - `manual`: all fields are empty.
	TriggerConfig TriggerConfig `json:"trigger_config" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointToolReview respjson.Field
		EndpointToolSlugs  respjson.Field
		Object             respjson.Field
		SystemPrompt       respjson.Field
		Temperature        respjson.Field
		Tier               respjson.Field
		TriggerConfig      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDefinitionConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentDefinitionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentDefinitionConfigObject string

const (
	AgentDefinitionConfigObjectAgentDefinitionConfig AgentDefinitionConfigObject = "agent_definition_config"
)

// Intelligence and cost tier for the agent's reasoning.
//
// Selects how capable and expensive a model the agent uses without pinning a
// specific model; higher tiers reason better but cost more. Each tier resolves to
// an ordered chain of equivalent models, so a run automatically fails over to
// another provider's model if the preferred one is unavailable.
//
//   - `frontier`: the most capable tier, for multi-step planning, ambiguous agent
//     work, and hard coding or architecture tasks.
//   - `high`: for normal planning, code edits, synthesis, and customer-facing
//     reasoning.
//   - `balanced`: for research, summarization, classification, structured
//     extraction, and light tool use.
//   - `cheap`: for simple transforms, validation, formatting, and routing.
//   - `legacy`: older-generation models kept for compatibility and regression
//     comparison; avoid unless you specifically need them.
//
// Leaving the tier unset picks one from how the agent is triggered: chat and
// manual runs use `high`, while scheduled and event-driven runs use `balanced` so
// background work stays cheap.
type AgentDefinitionConfigTier string

const (
	AgentDefinitionConfigTierFrontier AgentDefinitionConfigTier = "frontier"
	AgentDefinitionConfigTierHigh     AgentDefinitionConfigTier = "high"
	AgentDefinitionConfigTierBalanced AgentDefinitionConfigTier = "balanced"
	AgentDefinitionConfigTierCheap    AgentDefinitionConfigTier = "cheap"
	AgentDefinitionConfigTierLegacy   AgentDefinitionConfigTier = "legacy"
)

// Tool attached to an agent definition.
//
// Pairs an AvailableTool with agent-specific config values.
type AgentDefinitionTool struct {
	// Agent definition tool ID.
	ID string `json:"id" api:"required"`
	// Instance-specific configuration for this tool.
	//
	// Must conform to the tool's `config_schema`. Encoded as a JSON value (object,
	// array, string, number, boolean, or null), not a JSON-encoded string.
	Config any `json:"config" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_definition_tool".
	Object AgentDefinitionToolObject `json:"object" api:"required"`
	// Whether calls to this tool must be approved by a user before they execute.
	//
	// When `required`, the run pauses in the `awaiting_approval` status each time the
	// agent invokes this tool; approve or allow the tool via the Continue Agent Run
	// endpoint to proceed. A tool whose `mutating` flag is true still pauses for
	// approval even when this is `not_required`.
	//
	// Any of "not_required", "required".
	ReviewRequirement AgentDefinitionToolReviewRequirement `json:"review_requirement" api:"required"`
	// Sort order within the agent.
	SortOrder int64 `json:"sort_order" api:"required"`
	// A capability an agent can be granted, allowing it to take that action during a
	// run.
	//
	// The catalog of available tools is the same for every account; granting one to an
	// agent is what makes it callable.
	Tool AvailableTool `json:"tool" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Config            respjson.Field
		Object            respjson.Field
		ReviewRequirement respjson.Field
		SortOrder         respjson.Field
		Tool              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentDefinitionTool) RawJSON() string { return r.JSON.raw }
func (r *AgentDefinitionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentDefinitionToolObject string

const (
	AgentDefinitionToolObjectAgentDefinitionTool AgentDefinitionToolObject = "agent_definition_tool"
)

// Whether calls to this tool must be approved by a user before they execute.
//
// When `required`, the run pauses in the `awaiting_approval` status each time the
// agent invokes this tool; approve or allow the tool via the Continue Agent Run
// endpoint to proceed. A tool whose `mutating` flag is true still pauses for
// approval even when this is `not_required`.
type AgentDefinitionToolReviewRequirement string

const (
	AgentDefinitionToolReviewRequirementNotRequired AgentDefinitionToolReviewRequirement = "not_required"
	AgentDefinitionToolReviewRequirementRequired    AgentDefinitionToolReviewRequirement = "required"
)

// Agent-level configuration for creation/update requests.
type ConfigInputParam struct {
	// Instructions that define the agent's role and how it should behave.
	//
	// Sent to the model on every turn of a run, alongside the platform guidance
	// OpenMRP adds automatically.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	// How much randomness the model uses when generating text.
	//
	// Lower values make the agent's output more repeatable; higher values make it more
	// varied.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Per-endpoint-tool human-review overrides, keyed by tool slug.
	//
	// Set a slug to `true` to require human approval before the agent may execute that
	// endpoint-tool; the run pauses in `awaiting_approval` until approved via the
	// Continue Agent Run endpoint. Slugs omitted from the map do not require review.
	EndpointToolReview map[string]bool `json:"endpoint_tool_review,omitzero"`
	// API-endpoint tools the agent may discover and use, by slug (e.g.
	// `create_account_group`).
	//
	// These are the tools listed by the List Tools endpoint with category
	// `api_endpoint`. The single entry `*` grants the entire endpoint-tool catalog.
	// Omit or leave empty to grant none.
	EndpointToolSlugs []string `json:"endpoint_tool_slugs,omitzero"`
	// Intelligence and cost tier for the agent's reasoning.
	//
	// Selects how capable (and how expensive) a model the agent uses without pinning a
	// specific model, so the agent keeps working as the underlying model catalog
	// changes.
	//
	//   - `frontier`: the most capable and most expensive; multi-step planning,
	//     ambiguous work, tool-heavy workflows.
	//   - `high`: normal planning, synthesis, and customer-facing reasoning.
	//   - `balanced`: research, summarization, classification, structured extraction,
	//     and light tool use.
	//   - `cheap`: simple transforms, validation, formatting, keyword lookup, and
	//     routing.
	//   - `legacy`: older models kept for compatibility and regression comparison; avoid
	//     unless you specifically need them.
	//
	// Any of "frontier", "high", "balanced", "cheap", "legacy".
	Tier ConfigInputTier `json:"tier,omitzero"`
	// Trigger-type-specific settings for agent creation/update requests.
	//
	// Required contents depend on the agent's `trigger_type`:
	//
	// - `scheduled`: `cron_schedule` is required.
	// - `event`: at least one entry in `event_filters` is required.
	// - `manual` and `chat`: no trigger configuration is needed.
	TriggerConfig TriggerConfigInputParam `json:"trigger_config,omitzero"`
	paramObj
}

func (r ConfigInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Intelligence and cost tier for the agent's reasoning.
//
// Selects how capable (and how expensive) a model the agent uses without pinning a
// specific model, so the agent keeps working as the underlying model catalog
// changes.
//
//   - `frontier`: the most capable and most expensive; multi-step planning,
//     ambiguous work, tool-heavy workflows.
//   - `high`: normal planning, synthesis, and customer-facing reasoning.
//   - `balanced`: research, summarization, classification, structured extraction,
//     and light tool use.
//   - `cheap`: simple transforms, validation, formatting, keyword lookup, and
//     routing.
//   - `legacy`: older models kept for compatibility and regression comparison; avoid
//     unless you specifically need them.
type ConfigInputTier string

const (
	ConfigInputTierFrontier ConfigInputTier = "frontier"
	ConfigInputTierHigh     ConfigInputTier = "high"
	ConfigInputTierBalanced ConfigInputTier = "balanced"
	ConfigInputTierCheap    ConfigInputTier = "cheap"
	ConfigInputTierLegacy   ConfigInputTier = "legacy"
)

// Request to create an agent definition.
//
// The properties CategoryCode, Config, Name, Slug, TriggerType are required.
type CreateAgentRequestParam struct {
	// Category grouping for the agent (e.g. `order_processing`), used to organize
	// agents in the UI.
	CategoryCode string `json:"category_code" api:"required"`
	// Agent-level configuration for creation/update requests.
	Config ConfigInputParam `json:"config,omitzero" api:"required"`
	// Human-readable name of the agent.
	Name string `json:"name" api:"required"`
	// URL-friendly identifier for the agent.
	//
	// Must be unique within your account.
	Slug string `json:"slug" api:"required"`
	// How runs of this agent are initiated.
	//
	//   - `scheduled`: runs on a cron schedule; `config.trigger_config.cron_schedule` is
	//     required.
	//   - `event`: runs in response to platform events; at least one
	//     `config.trigger_config.event_filters` entry is required.
	//   - `manual`: runs only when explicitly invoked.
	//   - `chat`: runs when a user messages the agent in a conversation, and the agent's
	//     reply is posted back into that conversation.
	//
	// Whatever the trigger type, a run can always be started by hand with the Trigger
	// Agent Run endpoint.
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerType CreateAgentRequestTriggerType `json:"trigger_type,omitzero" api:"required"`
	// Description of what the agent does.
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the role that defines the permissions the agent operates with.
	//
	// Every API call the agent makes is authorized against this role, so it bounds
	// what the agent can see and change. An agent created without a role cannot
	// execute — its runs fail immediately — so attach one before triggering it.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// Built-in tools to attach to the agent.
	Tools []ToolInputParam `json:"tools,omitzero"`
	paramObj
}

func (r CreateAgentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateAgentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateAgentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How runs of this agent are initiated.
//
//   - `scheduled`: runs on a cron schedule; `config.trigger_config.cron_schedule` is
//     required.
//   - `event`: runs in response to platform events; at least one
//     `config.trigger_config.event_filters` entry is required.
//   - `manual`: runs only when explicitly invoked.
//   - `chat`: runs when a user messages the agent in a conversation, and the agent's
//     reply is posted back into that conversation.
//
// Whatever the trigger type, a run can always be started by hand with the Trigger
// Agent Run endpoint.
type CreateAgentRequestTriggerType string

const (
	CreateAgentRequestTriggerTypeScheduled CreateAgentRequestTriggerType = "scheduled"
	CreateAgentRequestTriggerTypeManual    CreateAgentRequestTriggerType = "manual"
	CreateAgentRequestTriggerTypeEvent     CreateAgentRequestTriggerType = "event"
	CreateAgentRequestTriggerTypeChat      CreateAgentRequestTriggerType = "chat"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAgentDefinition struct {
	// Resources in this page.
	Data []AgentDefinition `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentDefinitionObject `json:"object" api:"required"`
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
func (r ListAgentDefinition) RawJSON() string { return r.JSON.raw }
func (r *ListAgentDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentDefinitionObject string

const (
	ListAgentDefinitionObjectList ListAgentDefinitionObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAgentDefinitionTool struct {
	// Resources in this page.
	Data []AgentDefinitionTool `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentDefinitionToolObject `json:"object" api:"required"`
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
func (r ListAgentDefinitionTool) RawJSON() string { return r.JSON.raw }
func (r *ListAgentDefinitionTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentDefinitionToolObject string

const (
	ListAgentDefinitionToolObjectList ListAgentDefinitionToolObject = "list"
)

// Tool to attach to an agent definition.
//
// The property Tool is required.
type ToolInputParam struct {
	// The built-in tool to attach.
	//
	// Only OpenMRP's built-in tools are attached here. Access to API-endpoint tools
	// (creating a customer, listing orders, and so on) is granted separately through
	// `config.endpoint_tool_slugs`. The List Tools endpoint (`GET /v1/ai/tools`)
	// returns both kinds, with API-endpoint tools in the `api_endpoint` category.
	//
	// Any of "create_artifact", "read_doc", "fetch_url", "send_email", "draft_reply".
	Tool ToolInputTool `json:"tool,omitzero" api:"required"`
	// JSON-encoded configuration for this tool instance.
	//
	// The expected structure depends on the tool (see the tool's `config_schema`).
	ConfigJson param.Opt[string] `json:"config_json,omitzero"`
	// Whether actions from this tool require human review before they execute.
	//
	// When review is required, a call to this tool pauses the run in
	// `awaiting_approval` and records an action in `pending_review` until someone
	// approves or rejects it through the Continue Agent Run endpoint. Approvals are
	// one-time, so a later call to the same tool pauses again.
	RequireReview param.Opt[bool] `json:"require_review,omitzero"`
	// Display order among the agent's tools (lower values appear first).
	SortOrder param.Opt[int64] `json:"sort_order,omitzero"`
	paramObj
}

func (r ToolInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The built-in tool to attach.
//
// Only OpenMRP's built-in tools are attached here. Access to API-endpoint tools
// (creating a customer, listing orders, and so on) is granted separately through
// `config.endpoint_tool_slugs`. The List Tools endpoint (`GET /v1/ai/tools`)
// returns both kinds, with API-endpoint tools in the `api_endpoint` category.
type ToolInputTool string

const (
	ToolInputToolCreateArtifact ToolInputTool = "create_artifact"
	ToolInputToolReadDoc        ToolInputTool = "read_doc"
	ToolInputToolFetchURL       ToolInputTool = "fetch_url"
	ToolInputToolSendEmail      ToolInputTool = "send_email"
	ToolInputToolDraftReply     ToolInputTool = "draft_reply"
)

// Trigger-type-specific configuration.
//
// Which fields are populated depends on the agent's `trigger_type`:
//
// - `scheduled`: `cron_schedule` (and optionally `timezone`) is set.
// - `event`: `event_filters` is set.
// - `manual`: all fields are empty.
type TriggerConfig struct {
	// Cron expression for scheduled triggers (e.g. `0 9 * * *`).
	CronSchedule string `json:"cron_schedule" api:"required"`
	// Event types that trigger this agent (e.g.
	// `["email.received", "order.created"]`).
	EventFilters []string `json:"event_filters" api:"required"`
	// Resource type identifier.
	//
	// Any of "trigger_config".
	Object TriggerConfigObject `json:"object" api:"required"`
	// IANA timezone for the cron schedule (e.g. `America/New_York`).
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CronSchedule respjson.Field
		EventFilters respjson.Field
		Object       respjson.Field
		Timezone     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TriggerConfig) RawJSON() string { return r.JSON.raw }
func (r *TriggerConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type TriggerConfigObject string

const (
	TriggerConfigObjectTriggerConfig TriggerConfigObject = "trigger_config"
)

// Trigger-type-specific settings for agent creation/update requests.
//
// Required contents depend on the agent's `trigger_type`:
//
// - `scheduled`: `cron_schedule` is required.
// - `event`: at least one entry in `event_filters` is required.
// - `manual` and `chat`: no trigger configuration is needed.
type TriggerConfigInputParam struct {
	// Cron expression for scheduled triggers (e.g. `0 9 * * *`).
	CronSchedule param.Opt[string] `json:"cron_schedule,omitzero"`
	// IANA timezone for the cron schedule (e.g. `America/New_York`).
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Event types that trigger this agent (e.g.
	// `["email.received", "order.created"]`).
	EventFilters []string `json:"event_filters,omitzero"`
	paramObj
}

func (r TriggerConfigInputParam) MarshalJSON() (data []byte, err error) {
	type shadow TriggerConfigInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TriggerConfigInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request to partially update an agent definition.
type UpdateAgentRequestParam struct {
	// Description of what the agent does.
	//
	// Send `null` to clear the description; omit to leave it unchanged.
	Description param.Opt[string] `json:"description,omitzero"`
	// ID of the role that defines the permissions the agent operates with.
	//
	// Send `null` to detach the role; omit to leave it unchanged. An agent with no
	// role cannot execute, so detaching the role makes its runs fail immediately.
	RoleID param.Opt[string] `json:"role_id,omitzero"`
	// Category grouping for the agent (e.g. `order_processing`), used to organize
	// agents in the UI.
	CategoryCode param.Opt[string] `json:"category_code,omitzero"`
	// Human-readable name of the agent.
	Name param.Opt[string] `json:"name,omitzero"`
	// URL-friendly identifier for the agent.
	Slug param.Opt[string] `json:"slug,omitzero"`
	// Agent-level configuration for creation/update requests.
	Config ConfigInputParam `json:"config,omitzero"`
	// Built-in tools to attach to the agent.
	//
	// Replaces the existing tool set when provided.
	Tools []ToolInputParam `json:"tools,omitzero"`
	// How runs of this agent are initiated.
	//
	// When changing the trigger type, also provide a `config` with a `trigger_config`
	// appropriate for the new type (a cron schedule for `scheduled`, at least one
	// event filter for `event`).
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerType UpdateAgentRequestTriggerType `json:"trigger_type,omitzero"`
	paramObj
}

func (r UpdateAgentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAgentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAgentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How runs of this agent are initiated.
//
// When changing the trigger type, also provide a `config` with a `trigger_config`
// appropriate for the new type (a cron schedule for `scheduled`, at least one
// event filter for `event`).
type UpdateAgentRequestTriggerType string

const (
	UpdateAgentRequestTriggerTypeScheduled UpdateAgentRequestTriggerType = "scheduled"
	UpdateAgentRequestTriggerTypeManual    UpdateAgentRequestTriggerType = "manual"
	UpdateAgentRequestTriggerTypeEvent     UpdateAgentRequestTriggerType = "event"
	UpdateAgentRequestTriggerTypeChat      UpdateAgentRequestTriggerType = "chat"
)

// Request to update the per-account status of an agent.
//
// The property Status is required.
type UpdateAgentStatusRequestParam struct {
	// Account-level status to set for the agent.
	//
	// Any of "active", "inactive".
	Status UpdateAgentStatusRequestStatus `json:"status,omitzero" api:"required"`
	paramObj
}

func (r UpdateAgentStatusRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateAgentStatusRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateAgentStatusRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account-level status to set for the agent.
type UpdateAgentStatusRequestStatus string

const (
	UpdateAgentStatusRequestStatusActive   UpdateAgentStatusRequestStatus = "active"
	UpdateAgentStatusRequestStatusInactive UpdateAgentStatusRequestStatus = "inactive"
)

type AIAgentDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIAgentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *AIAgentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAgentNewParams struct {
	// Request to create an agent definition.
	CreateAgentRequest CreateAgentRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "config", "tools", "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r AIAgentNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateAgentRequest)
}
func (r *AIAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AIAgentNewParams]'s query parameters as `url.Values`.
func (r AIAgentNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAgentGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "config", "tools", "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAgentGetParams]'s query parameters as `url.Values`.
func (r AIAgentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAgentUpdateParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "config", "tools", "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// Request to partially update an agent definition.
	UpdateAgentRequest UpdateAgentRequestParam
	paramObj
}

func (r AIAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAgentRequest)
}
func (r *AIAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AIAgentUpdateParams]'s query parameters as `url.Values`.
func (r AIAgentUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAgentListParams struct {
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
	// Restricts results to agents of one of the given definition types.
	//
	// Any of "system", "custom".
	DefinitionTypes []string `query:"definition_types,omitzero" json:"-"`
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "config", "tools", "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	// Restricts results to agents with one of the given account-level statuses.
	//
	// `inactive` also matches agents that have never been enabled for your account.
	//
	// Any of "active", "inactive".
	Statuses []string `query:"statuses,omitzero" json:"-"`
	// Restricts results to agents with one of the given trigger types.
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerTypes []string `query:"trigger_types,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIAgentListParams]'s query parameters as `url.Values`.
func (r AIAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIAgentUpdateStatusParams struct {
	// Request to update the per-account status of an agent.
	UpdateAgentStatusRequest UpdateAgentStatusRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "config", "tools", "role", "role.permissions".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r AIAgentUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateAgentStatusRequest)
}
func (r *AIAgentUpdateStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AIAgentUpdateStatusParams]'s query parameters as
// `url.Values`.
func (r AIAgentUpdateStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
