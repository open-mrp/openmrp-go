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

func TestMessagingConversationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.New(context.TODO(), openmrp.MessagingConversationNewParams{
		CreateConversationRequest: openmrp.CreateConversationRequestParam{
			Type:                      openmrp.CreateConversationRequestTypeGroup,
			GroupID:                   openmrp.String("cvgp_wjlypugna7s4"),
			ParticipantAccountUserIDs: []string{"acus_e5zu8bde0z3h"},
			Participants: []openmrp.ConversationParticipantInputParam{{
				AccountUserID: "acus_e5zu8bde0z3h",
				Role:          openmrp.ConversationParticipantInputRoleAdmin,
			}},
			Title:             openmrp.String("Order #1042 — shipping question"),
			TopicResourceID:   openmrp.String("or_9lqo07quiwyb"),
			TopicResourceType: openmrp.CreateConversationRequestTopicResourceTypeSalesOrder,
		},
		Include: []string{"assignee"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingConversationGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Get(
		context.TODO(),
		"cv_w35z4ck68yq7",
		openmrp.MessagingConversationGetParams{
			Include: []string{"assignee"},
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

func TestMessagingConversationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.Update(
		context.TODO(),
		"cv_w35z4ck68yq7",
		openmrp.MessagingConversationUpdateParams{
			Include: []string{"assignee"},
			UpdateConversationRequest: openmrp.UpdateConversationRequestParam{
				Title: openmrp.String("Fulfillment war room"),
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

func TestMessagingConversationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.Conversations.List(context.TODO(), openmrp.MessagingConversationListParams{
		AssigneeResourceID: openmrp.String("assignee_resource_id"),
		Audience:           openmrp.MessagingConversationListParamsAudienceInternal,
		Cursor:             openmrp.String("cursor"),
		Include:            []string{"assignee"},
		IncludeArchived:    openmrp.Bool(true),
		Limit:              openmrp.Int(0),
		Q:                  openmrp.String("q"),
		Status:             openmrp.MessagingConversationListParamsStatusActive,
		TopicResourceID:    openmrp.String("topic_resource_id"),
		TopicResourceType:  openmrp.MessagingConversationListParamsTopicResourceTypeAccount,
		Type:               openmrp.MessagingConversationListParamsTypeDirectMessage,
		Unassigned:         openmrp.Bool(true),
		WorkflowStatus:     openmrp.MessagingConversationListParamsWorkflowStatusNew,
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
