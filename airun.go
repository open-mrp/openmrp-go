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

// List, retrieve, trigger, cancel, and continue agent runs.
//
// AIRunService contains methods and other services that help with interacting with
// the openmrp API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIRunService] method instead.
type AIRunService struct {
	options []option.RequestOption
	// List, retrieve, trigger, cancel, and continue agent runs.
	Actions AIRunActionService
}

// NewAIRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIRunService(opts ...option.RequestOption) (r AIRunService) {
	r = AIRunService{}
	r.options = opts
	r.Actions = NewAIRunActionService(opts...)
	return
}

// Starts a new run of the specified agent.
//
// The run is created in the `pending` status and executed asynchronously; poll
// Retrieve Agent Run to follow its progress. Any agent can be started this way
// regardless of how it is normally triggered, and the resulting run is always
// recorded with `trigger_type` `manual`.
//
// This endpoint requires the permission: `agent_runs:create`.
func (r *AIRunService) New(ctx context.Context, params AIRunNewParams, opts ...option.RequestOption) (res *AgentRun, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a single agent run by ID.
//
// A run records one execution of an agent: its current status, the input it
// started from, the output it produced, the tools it invoked, and the step-by-step
// timeline of how it got there.
//
// This endpoint requires the permission: `agent_runs:read`.
func (r *AIRunService) Get(ctx context.Context, id string, query AIRunGetParams, opts ...option.RequestOption) (res *AgentRun, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/ai/runs/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Lists agent runs for your account, newest first.
//
// The `q` parameter matches a run's ID, its status, or the ID of the agent that
// produced it.
//
// This endpoint requires the permission: `agent_runs:read`.
func (r *AIRunService) List(ctx context.Context, query AIRunListParams, opts ...option.RequestOption) (res *ListAgentRun, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/ai/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A single tool invocation performed by an agent during a run.
//
// Each action records the tool that was called, its input and output, and any
// human review decision.
type AgentAction struct {
	// Agent action ID.
	ID string `json:"id" api:"required"`
	// When this action was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Longer description of what the action does.
	Description string `json:"description" api:"required"`
	// Entity is a polymorphic reference to any resource in the system.
	Entity Entity `json:"entity" api:"required"`
	// Error message if the action failed.
	ErrorMessage string `json:"error_message" api:"required"`
	// When the action was executed.
	ExecutedAt time.Time `json:"executed_at" api:"required" format:"date-time"`
	// Arguments passed to the tool, as JSON.
	//
	// Shape depends on `tool`. Encoded as a JSON value (object, array, string, number,
	// boolean, or null), not a JSON-encoded string.
	Input any `json:"input" api:"required"`
	// Short human-readable label summarizing the action.
	Label string `json:"label" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_action".
	Object AgentActionObject `json:"object" api:"required"`
	// Result returned by the tool, as JSON.
	//
	// The shape depends on `tool`. An action that has not executed — because it is
	// still waiting on a review decision, or was rejected — carries `{}`. Encoded as a
	// JSON value (object, array, string, number, boolean, or null), not a JSON-encoded
	// string.
	Output any `json:"output" api:"required"`
	// Whether a person must approve this action before it takes effect.
	//
	// Fixed when the action is recorded, from the agent's review setting for that
	// tool; tools that take an externally visible action, such as `send_email`, always
	// require review and cannot be exempted. When review is required the action starts
	// in `pending_review` and stays there until someone approves or rejects it;
	// otherwise it is `auto_approved`.
	//
	// Any of "not_required", "required".
	ReviewRequirement AgentActionReviewRequirement `json:"review_requirement" api:"required"`
	// When a human review decision was recorded for the action.
	ReviewedAt time.Time `json:"reviewed_at" api:"required" format:"date-time"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	ReviewedBy Actor `json:"reviewed_by" api:"required"`
	// A single execution of an agent, from trigger through completion.
	Run *AgentRun `json:"run" api:"required"`
	// Current action status.
	//
	// - `pending_review`: awaiting human review before it can execute.
	// - `auto_approved`: automatically approved by policy.
	// - `approved`: manually approved by a user.
	// - `rejected`: rejected by a user; will not execute.
	// - `executed`: successfully executed.
	// - `failed`: errored during execution; see `error_message`.
	//
	// Any of "pending_review", "auto_approved", "approved", "rejected", "executed",
	// "failed".
	Status AgentActionStatus `json:"status" api:"required"`
	// The tool the agent invoked for this action.
	//
	//   - `create_artifact`: create an artifact such as a report, document, or data
	//     export.
	//   - `read_doc`: read OpenMRP documentation pages.
	//   - `fetch_url`: fetch content from a public URL.
	//   - `draft_reply`: propose a reply to the case's external party as a draft held
	//     for human approval (not sent).
	//   - `send_email`: send an email reply through the conversation's bound inbox.
	//
	// Any of "create_artifact", "read_doc", "fetch_url", "send_email", "draft_reply".
	Tool AgentActionTool `json:"tool" api:"required"`
	// When this action was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAt         respjson.Field
		Description       respjson.Field
		Entity            respjson.Field
		ErrorMessage      respjson.Field
		ExecutedAt        respjson.Field
		Input             respjson.Field
		Label             respjson.Field
		Object            respjson.Field
		Output            respjson.Field
		ReviewRequirement respjson.Field
		ReviewedAt        respjson.Field
		ReviewedBy        respjson.Field
		Run               respjson.Field
		Status            respjson.Field
		Tool              respjson.Field
		UpdatedAt         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAction) RawJSON() string { return r.JSON.raw }
func (r *AgentAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentActionObject string

const (
	AgentActionObjectAgentAction AgentActionObject = "agent_action"
)

// Whether a person must approve this action before it takes effect.
//
// Fixed when the action is recorded, from the agent's review setting for that
// tool; tools that take an externally visible action, such as `send_email`, always
// require review and cannot be exempted. When review is required the action starts
// in `pending_review` and stays there until someone approves or rejects it;
// otherwise it is `auto_approved`.
type AgentActionReviewRequirement string

const (
	AgentActionReviewRequirementNotRequired AgentActionReviewRequirement = "not_required"
	AgentActionReviewRequirementRequired    AgentActionReviewRequirement = "required"
)

// Current action status.
//
// - `pending_review`: awaiting human review before it can execute.
// - `auto_approved`: automatically approved by policy.
// - `approved`: manually approved by a user.
// - `rejected`: rejected by a user; will not execute.
// - `executed`: successfully executed.
// - `failed`: errored during execution; see `error_message`.
type AgentActionStatus string

const (
	AgentActionStatusPendingReview AgentActionStatus = "pending_review"
	AgentActionStatusAutoApproved  AgentActionStatus = "auto_approved"
	AgentActionStatusApproved      AgentActionStatus = "approved"
	AgentActionStatusRejected      AgentActionStatus = "rejected"
	AgentActionStatusExecuted      AgentActionStatus = "executed"
	AgentActionStatusFailed        AgentActionStatus = "failed"
)

// The tool the agent invoked for this action.
//
//   - `create_artifact`: create an artifact such as a report, document, or data
//     export.
//   - `read_doc`: read OpenMRP documentation pages.
//   - `fetch_url`: fetch content from a public URL.
//   - `draft_reply`: propose a reply to the case's external party as a draft held
//     for human approval (not sent).
//   - `send_email`: send an email reply through the conversation's bound inbox.
type AgentActionTool string

const (
	AgentActionToolCreateArtifact AgentActionTool = "create_artifact"
	AgentActionToolReadDoc        AgentActionTool = "read_doc"
	AgentActionToolFetchURL       AgentActionTool = "fetch_url"
	AgentActionToolSendEmail      AgentActionTool = "send_email"
	AgentActionToolDraftReply     AgentActionTool = "draft_reply"
)

// A single execution of an agent, from trigger through completion.
type AgentRun struct {
	// Agent run ID.
	ID string `json:"id" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Actions *ListAgentAction `json:"actions" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	// When this run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// An AI agent available to the account.
	//
	// The definition describes what the agent does, how its runs are triggered, the
	// tools it can use, and whether it is currently enabled for the account.
	Definition AgentDefinition `json:"definition" api:"required"`
	// How long the run took, in milliseconds.
	DurationMs int64 `json:"duration_ms" api:"required"`
	// Error message if the run failed.
	ErrorMessage string `json:"error_message" api:"required"`
	// Input provided to the agent at the start of the run.
	//
	// The shape depends on what started the run; a manually triggered run records
	// `{"message": "<your input>"}`. Encoded as a JSON value (object, array, string,
	// number, boolean, or null), not a JSON-encoded string.
	Input any `json:"input" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_run".
	Object AgentRunObject `json:"object" api:"required"`
	// Final output produced by the agent.
	//
	// Present once the agent has produced a result, including on a run that paused for
	// more input or was cancelled part-way through. A run that has not produced one
	// yet carries an empty object. Encoded as a JSON value (object, array, string,
	// number, boolean, or null), not a JSON-encoded string.
	Output any `json:"output" api:"required"`
	// When the run started executing.
	StartedAt time.Time `json:"started_at" api:"required" format:"date-time"`
	// Current run status.
	//
	// - `pending`: queued but not yet started.
	// - `running`: currently executing.
	// - `awaiting_input`: paused, waiting for user input before continuing.
	// - `awaiting_approval`: paused, waiting for a pending action to be approved.
	// - `completed`: finished successfully.
	// - `failed`: stopped after an error; see `error_message`.
	// - `cancelled`: stopped before completion by a user.
	//
	// Any of "pending", "running", "completed", "failed", "cancelled",
	// "awaiting_input", "awaiting_approval".
	Status AgentRunStatus `json:"status" api:"required"`
	// A single page of resources, together with the metadata needed to page through
	// the rest of the result set.
	Steps ListAgentRunStep `json:"steps" api:"required"`
	// How this run was initiated.
	//
	//   - `scheduled`: started by the agent's cron schedule.
	//   - `event`: started in response to a platform event.
	//   - `manual`: started by an explicit request; see `triggered_by`.
	//   - `chat`: started by a message in a conversation, with the agent's reply posted
	//     back into that conversation.
	//
	// Any of "scheduled", "manual", "event", "chat".
	TriggerType AgentRunTriggerType `json:"trigger_type" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	TriggeredBy Actor `json:"triggered_by" api:"required"`
	// When this run was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Actions      respjson.Field
		CompletedAt  respjson.Field
		CreatedAt    respjson.Field
		Definition   respjson.Field
		DurationMs   respjson.Field
		ErrorMessage respjson.Field
		Input        respjson.Field
		Object       respjson.Field
		Output       respjson.Field
		StartedAt    respjson.Field
		Status       respjson.Field
		Steps        respjson.Field
		TriggerType  respjson.Field
		TriggeredBy  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRun) RawJSON() string { return r.JSON.raw }
func (r *AgentRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentRunObject string

const (
	AgentRunObjectAgentRun AgentRunObject = "agent_run"
)

// Current run status.
//
// - `pending`: queued but not yet started.
// - `running`: currently executing.
// - `awaiting_input`: paused, waiting for user input before continuing.
// - `awaiting_approval`: paused, waiting for a pending action to be approved.
// - `completed`: finished successfully.
// - `failed`: stopped after an error; see `error_message`.
// - `cancelled`: stopped before completion by a user.
type AgentRunStatus string

const (
	AgentRunStatusPending          AgentRunStatus = "pending"
	AgentRunStatusRunning          AgentRunStatus = "running"
	AgentRunStatusCompleted        AgentRunStatus = "completed"
	AgentRunStatusFailed           AgentRunStatus = "failed"
	AgentRunStatusCancelled        AgentRunStatus = "cancelled"
	AgentRunStatusAwaitingInput    AgentRunStatus = "awaiting_input"
	AgentRunStatusAwaitingApproval AgentRunStatus = "awaiting_approval"
)

// How this run was initiated.
//
//   - `scheduled`: started by the agent's cron schedule.
//   - `event`: started in response to a platform event.
//   - `manual`: started by an explicit request; see `triggered_by`.
//   - `chat`: started by a message in a conversation, with the agent's reply posted
//     back into that conversation.
type AgentRunTriggerType string

const (
	AgentRunTriggerTypeScheduled AgentRunTriggerType = "scheduled"
	AgentRunTriggerTypeManual    AgentRunTriggerType = "manual"
	AgentRunTriggerTypeEvent     AgentRunTriggerType = "event"
	AgentRunTriggerTypeChat      AgentRunTriggerType = "chat"
)

// A single event in an agent run's execution timeline.
type AgentRunStep struct {
	// Agent run step ID.
	ID string `json:"id" api:"required"`
	// Reference to an actor — the user, API key, agent, or group identity associated
	// with an action.
	Actor Actor `json:"actor" api:"required"`
	// Text payload for the step, such as a message body or a tool result.
	Content string `json:"content" api:"required"`
	// When this step was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// How long this step took, in milliseconds.
	DurationMs int64 `json:"duration_ms" api:"required"`
	// Additional structured data for the step.
	//
	// The shape depends on `step_type` — for example a `tool_call` step carries the
	// tool's arguments. Encoded as a JSON value (object, array, string, number,
	// boolean, or null), not a JSON-encoded string.
	Metadata any `json:"metadata" api:"required"`
	// Resource type identifier.
	//
	// Any of "agent_run_step".
	Object AgentRunStepObject `json:"object" api:"required"`
	// Zero-based position of this step within the run's timeline.
	Sequence int64 `json:"sequence" api:"required"`
	// The kind of timeline event.
	//
	// Common values are `trigger_received`, `user_message`, `thinking`,
	// `assistant_message`, `tool_call`, `tool_result`, `tool_blocked`,
	// `awaiting_approval`, `completion`, and `error`. This is an open set — new step
	// types are added as the agent runtime evolves, so treat unrecognized values as
	// informational rather than failing on them.
	StepType string `json:"step_type" api:"required"`
	// Short title for the step.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actor       respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		DurationMs  respjson.Field
		Metadata    respjson.Field
		Object      respjson.Field
		Sequence    respjson.Field
		StepType    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunStep) RawJSON() string { return r.JSON.raw }
func (r *AgentRunStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type AgentRunStepObject string

const (
	AgentRunStepObjectAgentRunStep AgentRunStepObject = "agent_run_step"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAgentAction struct {
	// Resources in this page.
	Data []AgentAction `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentActionObject `json:"object" api:"required"`
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
func (r ListAgentAction) RawJSON() string { return r.JSON.raw }
func (r *ListAgentAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentActionObject string

const (
	ListAgentActionObjectList ListAgentActionObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAgentRun struct {
	// Resources in this page.
	Data []AgentRun `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentRunObject `json:"object" api:"required"`
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
func (r ListAgentRun) RawJSON() string { return r.JSON.raw }
func (r *ListAgentRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentRunObject string

const (
	ListAgentRunObjectList ListAgentRunObject = "list"
)

// A single page of resources, together with the metadata needed to page through
// the rest of the result set.
type ListAgentRunStep struct {
	// Resources in this page.
	Data []AgentRunStep `json:"data" api:"required"`
	// Resource type identifier.
	//
	// Any of "list".
	Object ListAgentRunStepObject `json:"object" api:"required"`
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
func (r ListAgentRunStep) RawJSON() string { return r.JSON.raw }
func (r *ListAgentRunStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type identifier.
type ListAgentRunStepObject string

const (
	ListAgentRunStepObjectList ListAgentRunStepObject = "list"
)

// Request to trigger an agent run.
//
// The property AgentDefinitionID is required.
type TriggerRunRequestParam struct {
	// ID of the agent definition to run.
	//
	// The agent must be active for the account; triggering an inactive agent returns a
	// validation error.
	AgentDefinitionID string `json:"agent_definition_id" api:"required"`
	// Instruction text passed to the agent at the start of the run.
	//
	// Recorded on the run as `{"message": <input>}` in its `input` field.
	Input param.Opt[string] `json:"input,omitzero"`
	paramObj
}

func (r TriggerRunRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TriggerRunRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TriggerRunRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIRunNewParams struct {
	// Request to trigger an agent run.
	TriggerRunRequest TriggerRunRequestParam
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "actions", "definition", "definition.config", "definition.tools",
	// "definition.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

func (r AIRunNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TriggerRunRequest)
}
func (r *AIRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [AIRunNewParams]'s query parameters as `url.Values`.
func (r AIRunNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIRunGetParams struct {
	// Sub-objects to expand in the response. When omitted, sub-objects are returned as
	// `null`.
	//
	// Any of "triggered_by", "actions", "definition", "steps", "definition.config",
	// "definition.tools", "definition.role".
	Include []string `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIRunGetParams]'s query parameters as `url.Values`.
func (r AIRunGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIRunListParams struct {
	// Restricts results to runs of a single agent.
	AgentDefinitionID param.Opt[string] `query:"agent_definition_id,omitzero" json:"-"`
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
	// Any of "triggered_by", "definition", "actions", "definition.config",
	// "definition.tools", "definition.role".
	Include []string `query:"include,omitzero" json:"-"`
	// Restricts results to runs in this status.
	//
	// Any of "pending", "running", "completed", "failed", "cancelled",
	// "awaiting_input", "awaiting_approval".
	Status AIRunListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIRunListParams]'s query parameters as `url.Values`.
func (r AIRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Restricts results to runs in this status.
type AIRunListParamsStatus string

const (
	AIRunListParamsStatusPending          AIRunListParamsStatus = "pending"
	AIRunListParamsStatusRunning          AIRunListParamsStatus = "running"
	AIRunListParamsStatusCompleted        AIRunListParamsStatus = "completed"
	AIRunListParamsStatusFailed           AIRunListParamsStatus = "failed"
	AIRunListParamsStatusCancelled        AIRunListParamsStatus = "cancelled"
	AIRunListParamsStatusAwaitingInput    AIRunListParamsStatus = "awaiting_input"
	AIRunListParamsStatusAwaitingApproval AIRunListParamsStatus = "awaiting_approval"
)
