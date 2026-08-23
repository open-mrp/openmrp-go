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

func TestMessagingEmailInboxNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.EmailInboxes.New(context.TODO(), openmrp.MessagingEmailInboxNewParams{
		CreateEmailInboxRequest: openmrp.CreateEmailInboxRequestParam{
			Address:              "support@acme.com",
			EmailDomainID:        "emdom_2rk3omr8vshb",
			AgentConfigID:        openmrp.String("agdf_ah7tkyfxk8jl"),
			AgentTriggerKeywords: []string{"invoice", "refund"},
			AgentTriggerPolicy:   openmrp.CreateEmailInboxRequestAgentTriggerPolicyKeyword,
			FromName:             openmrp.String("Acme Support"),
			GroupID:              openmrp.String("group_id"),
		},
		Include: []string{"email_domain"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingEmailInboxGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.EmailInboxes.Get(
		context.TODO(),
		"eminb_2s9kobr9s7tp",
		openmrp.MessagingEmailInboxGetParams{
			Include: []string{"email_domain"},
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

func TestMessagingEmailInboxUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.EmailInboxes.Update(
		context.TODO(),
		"eminb_2s9kobr9s7tp",
		openmrp.MessagingEmailInboxUpdateParams{
			UpdateEmailInboxRequest: openmrp.UpdateEmailInboxRequestParam{
				Status:               openmrp.UpdateEmailInboxRequestStatusActive,
				AgentConfigID:        openmrp.String("agdf_ah7tkyfxk8jl"),
				AgentTriggerKeywords: []string{"invoice", "refund"},
				AgentTriggerPolicy:   openmrp.UpdateEmailInboxRequestAgentTriggerPolicyKeyword,
				FromName:             openmrp.String("Acme Support"),
				GroupID:              openmrp.String("group_id"),
			},
			Include: []string{"email_domain"},
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

func TestMessagingEmailInboxListWithOptionalParams(t *testing.T) {
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
	_, err := client.Messaging.EmailInboxes.List(context.TODO(), openmrp.MessagingEmailInboxListParams{
		Include: []string{"email_domain"},
	})
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessagingEmailInboxDelete(t *testing.T) {
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
	_, err := client.Messaging.EmailInboxes.Delete(context.TODO(), "eminb_2s9kobr9s7tp")
	if err != nil {
		var apierr *openmrp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
