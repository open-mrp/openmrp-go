// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openmrp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/open-mrp/openmrp-go"
	"github.com/open-mrp/openmrp-go/internal/testutil"
	"github.com/open-mrp/openmrp-go/option"
)

func TestAIAgentNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.AI.Agents.New(context.TODO(), openmrp.AIAgentNewParams{
		CreateAgentRequest: openmrp.CreateAgentRequestParam{
			CategoryCode: "inventory",
			Config: openmrp.ConfigInputParam{
				EndpointToolReview: map[string]bool{
					"foo": true,
				},
				EndpointToolSlugs: []string{"string"},
				SystemPrompt:      openmrp.String("You are an order processing agent. Parse incoming emails and create draft orders."),
				Temperature:       openmrp.Float(0.2),
				Tier:              openmrp.ConfigInputTierHigh,
				TriggerConfig: openmrp.TriggerConfigInputParam{
					CronSchedule: openmrp.String("cron_schedule"),
					EventFilters: []string{"email.received"},
					Timezone:     openmrp.String("timezone"),
				},
			},
			Name:        "Inventory Monitor",
			Slug:        "inventory_monitor",
			TriggerType: openmrp.CreateAgentRequestTriggerTypeEvent,
			Description: openmrp.String("Monitors inventory levels and creates restock alerts."),
			RoleID:      openmrp.String("rl_3xknmfqflhvb"),
			Tools: []openmrp.ToolInputParam{{
				Tool:          openmrp.ToolInputToolReadDoc,
				ConfigJson:    openmrp.String("config_json"),
				RequireReview: openmrp.Bool(true),
				SortOrder:     openmrp.Int(1),
			}},
		},
		Include: []string{"config"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.AI.Agents.Get(
		context.TODO(),
		"agdf_ah7tkyfxk8jl",
		openmrp.AIAgentGetParams{
			Include: []string{"config"},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.AI.Agents.Update(
		context.TODO(),
		"agdf_ah7tkyfxk8jl",
		openmrp.AIAgentUpdateParams{
			Include: []string{"config"},
			UpdateAgentRequest: openmrp.UpdateAgentRequestParam{
				CategoryCode: openmrp.String("category_code"),
				Config: openmrp.ConfigInputParam{
					EndpointToolReview: map[string]bool{
						"foo": true,
					},
					EndpointToolSlugs: []string{"string"},
					SystemPrompt:      openmrp.String("You are an order processing agent. Parse incoming emails and create draft orders."),
					Temperature:       openmrp.Float(0.2),
					Tier:              openmrp.ConfigInputTierHigh,
					TriggerConfig: openmrp.TriggerConfigInputParam{
						CronSchedule: openmrp.String("cron_schedule"),
						EventFilters: []string{"email.received"},
						Timezone:     openmrp.String("timezone"),
					},
				},
				Description: openmrp.String("description"),
				Name:        openmrp.String("Inventory Monitor"),
				RoleID:      openmrp.String("role_id"),
				Slug:        openmrp.String("slug"),
				Tools: []openmrp.ToolInputParam{{
					Tool:          openmrp.ToolInputToolReadDoc,
					ConfigJson:    openmrp.String("config_json"),
					RequireReview: openmrp.Bool(true),
					SortOrder:     openmrp.Int(1),
				}},
				TriggerType: openmrp.UpdateAgentRequestTriggerTypeScheduled,
			},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.AI.Agents.List(context.TODO(), openmrp.AIAgentListParams{
		Cursor:          openmrp.String("cursor"),
		DefinitionTypes: []string{"system"},
		Include:         []string{"config"},
		Limit:           openmrp.Int(0),
		Q:               openmrp.String("q"),
		Statuses:        []string{"active"},
		TriggerTypes:    []string{"scheduled"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.AI.Agents.Delete(context.TODO(), "agdf_ah7tkyfxk8jl")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIAgentUpdateStatusWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openmrp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.AI.Agents.UpdateStatus(
		context.TODO(),
		"agdf_ah7tkyfxk8jl",
		openmrp.AIAgentUpdateStatusParams{
			UpdateAgentStatusRequest: openmrp.UpdateAgentStatusRequestParam{
				Status: openmrp.UpdateAgentStatusRequestStatusActive,
			},
			Include: []string{"config"},
		},
	)
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
