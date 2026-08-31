# Auth

## APIKeys

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAPIKeyRequestParam">CreateAPIKeyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountBranding">AccountBranding</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountPortal">AccountPortal</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatedAPIKey">CreatedAPIKey</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Geolocation">Geolocation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAPIKey">ListAPIKey</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Owner">Owner</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PageInfo">PageInfo</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Role">Role</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyDeleteResponse">AuthAPIKeyDeleteResponse</a>

Methods:

- <code title="post /v1/auth/api-keys">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyNewParams">AuthAPIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatedAPIKey">CreatedAPIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/api-keys/{id}">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyGetParams">AuthAPIKeyGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auth/api-keys">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyListParams">AuthAPIKeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAPIKey">ListAPIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/auth/api-keys/{id}">client.Auth.APIKeys.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyDeleteResponse">AuthAPIKeyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RotateAPIKeyRequestParam">RotateAPIKeyRequestParam</a>

Methods:

- <code title="post /v1/auth/api-keys/{id}/actions/rotate">client.Auth.APIKeys.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyActionService.Rotate">Rotate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuthAPIKeyActionRotateParams">AuthAPIKeyActionRotateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatedAPIKey">CreatedAPIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Core

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Entity">Entity</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEntity">ListEntity</a>

Methods:

- <code title="get /v1/core/search">client.Core.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreService.GetSearch">GetSearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreGetSearchParams">CoreGetSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEntity">ListEntity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sandboxes

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateSandboxRequestParam">CreateSandboxRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSandbox">ListSandbox</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Sandbox">Sandbox</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxDeleteResponse">CoreSandboxDeleteResponse</a>

Methods:

- <code title="post /v1/core/sandboxes">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxNewParams">CoreSandboxNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/sandboxes/{id}">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxGetParams">CoreSandboxGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/sandboxes">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxListParams">CoreSandboxListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSandbox">ListSandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/core/sandboxes/{id}">client.Core.Sandboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreSandboxDeleteResponse">CoreSandboxDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RequestLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Actor">Actor</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListRequestLog">ListRequestLog</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RequestLog">RequestLog</a>

Methods:

- <code title="get /v1/core/request-logs/{id}">client.Core.RequestLogs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreRequestLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreRequestLogGetParams">CoreRequestLogGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RequestLog">RequestLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/request-logs">client.Core.RequestLogs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreRequestLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreRequestLogListParams">CoreRequestLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListRequestLog">ListRequestLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditEvents

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuditEvent">AuditEvent</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuditFieldChange">AuditFieldChange</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAuditEvent">ListAuditEvent</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAuditFieldChange">ListAuditFieldChange</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListObjectType">ListObjectType</a>

Methods:

- <code title="get /v1/core/audit-events/{id}">client.Core.AuditEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAuditEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAuditEventGetParams">CoreAuditEventGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AuditEvent">AuditEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/audit-events">client.Core.AuditEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAuditEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAuditEventListParams">CoreAuditEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAuditEvent">ListAuditEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/audit-events/resource-types">client.Core.AuditEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAuditEventService.GetResourceTypes">GetResourceTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListObjectType">ListObjectType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Addresses

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AddressSuggestion">AddressSuggestion</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAddressSuggestion">ListAddressSuggestion</a>

Methods:

- <code title="get /v1/core/addresses/suggestions">client.Core.Addresses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAddressService.GetSuggestions">GetSuggestions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAddressGetSuggestionsParams">CoreAddressGetSuggestionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAddressSuggestion">ListAddressSuggestion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ValidateAddressRequestParam">ValidateAddressRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AddressComponents">AddressComponents</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ValidatedAddress">ValidatedAddress</a>

Methods:

- <code title="put /v1/core/addresses/actions/validate">client.Core.Addresses.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAddressActionService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAddressActionValidateParams">CoreAddressActionValidateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ValidatedAddress">ValidatedAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailLog">EmailLog</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEmailLog">ListEmailLog</a>

Methods:

- <code title="get /v1/core/email-logs/{id}">client.Core.EmailLogs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreEmailLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreEmailLogGetParams">CoreEmailLogGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailLog">EmailLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/core/email-logs">client.Core.EmailLogs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreEmailLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreEmailLogListParams">CoreEmailLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEmailLog">ListEmailLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#JobExport">JobExport</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#JobResult">JobResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListJobResult">ListJobResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuotaInfo">QuotaInfo</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ResponseError">ResponseError</a>

Methods:

- <code title="get /v1/core/jobs/{id}">client.Core.Jobs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreJobGetParams">CoreJobGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/core/jobs/{id}/cancel">client.Core.Jobs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreJobCancelParams">CoreJobCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Analytics

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeDeliveryPerformanceRequestParam">AnalyzeDeliveryPerformanceRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeOeeRequestParam">AnalyzeOeeRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeOeeTrendRequestParam">AnalyzeOeeTrendRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeScheduleAttainmentRequestParam">AnalyzeScheduleAttainmentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OeeDepartmentPlannedTimeParam">OeeDepartmentPlannedTimeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeDeliveryPerformanceResponse">AnalyzeDeliveryPerformanceResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeOeeResponse">AnalyzeOeeResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeOeeTrendResponse">AnalyzeOeeTrendResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeScheduleAttainmentResponse">AnalyzeScheduleAttainmentResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AttainmentBucket">AttainmentBucket</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DeliveryBacklogBucket">DeliveryBacklogBucket</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DeliveryBreakdown">DeliveryBreakdown</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DeliveryLatenessBucket">DeliveryLatenessBucket</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DeliveryPerformance">DeliveryPerformance</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FrozenAdherence">FrozenAdherence</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAttainmentBucket">ListAttainmentBucket</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDeliveryBacklogBucket">ListDeliveryBacklogBucket</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDeliveryBreakdown">ListDeliveryBreakdown</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDeliveryLatenessBucket">ListDeliveryLatenessBucket</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDeliveryPerformance">ListDeliveryPerformance</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListFrozenAdherence">ListFrozenAdherence</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOeeDepartment">ListOeeDepartment</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOeeDowntimeReason">ListOeeDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOeeTrendPeriod">ListOeeTrendPeriod</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OeeDepartment">OeeDepartment</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OeeDowntimeReason">OeeDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OeeTrendPeriod">OeeTrendPeriod</a>

Methods:

- <code title="put /v1/core/analytics/delivery-performance">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsService.UpdateDeliveryPerformance">UpdateDeliveryPerformance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsUpdateDeliveryPerformanceParams">CoreAnalyticsUpdateDeliveryPerformanceParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeDeliveryPerformanceResponse">AnalyzeDeliveryPerformanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/core/analytics/oee">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsService.UpdateOee">UpdateOee</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsUpdateOeeParams">CoreAnalyticsUpdateOeeParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeOeeResponse">AnalyzeOeeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/core/analytics/oee-trend">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsService.UpdateOeeTrend">UpdateOeeTrend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsUpdateOeeTrendParams">CoreAnalyticsUpdateOeeTrendParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeOeeTrendResponse">AnalyzeOeeTrendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/core/analytics/schedule-attainment">client.Core.Analytics.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsService.UpdateScheduleAttainment">UpdateScheduleAttainment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreAnalyticsUpdateScheduleAttainmentParams">CoreAnalyticsUpdateScheduleAttainmentParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AnalyzeScheduleAttainmentResponse">AnalyzeScheduleAttainmentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailRecordRequestParam">EmailRecordRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreActionEmailRecordResponse">CoreActionEmailRecordResponse</a>

Methods:

- <code title="post /v1/core/actions/email-record">client.Core.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreActionService.EmailRecord">EmailRecord</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreActionEmailRecordParams">CoreActionEmailRecordParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CoreActionEmailRecordResponse">CoreActionEmailRecordResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Catalog

## Units

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateUnitRequestParam">CreateUnitRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateUnitRequestParam">UpdateUnitRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListUnit">ListUnit</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Unit">Unit</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitDeleteResponse">CatalogUnitDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/units">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitNewParams">CatalogUnitNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Unit">Unit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/units/{id}">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGetParams">CatalogUnitGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Unit">Unit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/units/{id}">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitUpdateParams">CatalogUnitUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Unit">Unit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/units">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitListParams">CatalogUnitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListUnit">ListUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/units/{id}">client.Catalog.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitDeleteResponse">CatalogUnitDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertUnitsRequestParam">BulkUpsertUnitsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertUnitInputParam">UpsertUnitInputParam</a>

Methods:

- <code title="post /v1/catalog/units/actions/bulk-upsert">client.Catalog.Units.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitActionBulkUpsertParams">CatalogUnitActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## UnitGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateUnitGroupRequestParam">CreateUnitGroupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateUnitGroupUnitParam">CreateUnitGroupUnitParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateUnitGroupRequestParam">UpdateUnitGroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListUnitGroup">ListUnitGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListUnitGroupUnit">ListUnitGroupUnit</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroup">UnitGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroupUnit">UnitGroupUnit</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupDeleteResponse">CatalogUnitGroupDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/unit-groups">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupNewParams">CatalogUnitGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroup">UnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups/{id}">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupGetParams">CatalogUnitGroupGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroup">UnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/unit-groups/{id}">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUpdateParams">CatalogUnitGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroup">UnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupListParams">CatalogUnitGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListUnitGroup">ListUnitGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/unit-groups/{id}">client.Catalog.UnitGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupDeleteResponse">CatalogUnitGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Units

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateUnitGroupUnitRequestParam">CreateUnitGroupUnitRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateUnitGroupUnitRequestParam">UpdateUnitGroupUnitRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitDeleteResponse">CatalogUnitGroupUnitDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/unit-groups/{unit_group_id}/units">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, unitGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitNewParams">CatalogUnitGroupUnitNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroupUnit">UnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups/{unit_group_id}/units/{id}">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitGetParams">CatalogUnitGroupUnitGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroupUnit">UnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/unit-groups/{unit_group_id}/units/{id}">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitUpdateParams">CatalogUnitGroupUnitUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitGroupUnit">UnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/unit-groups/{unit_group_id}/units">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, unitGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitListParams">CatalogUnitGroupUnitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListUnitGroupUnit">ListUnitGroupUnit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/unit-groups/{unit_group_id}/units/{id}">client.Catalog.UnitGroups.Units.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitDeleteParams">CatalogUnitGroupUnitDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupUnitDeleteResponse">CatalogUnitGroupUnitDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertUnitGroupsRequestParam">BulkUpsertUnitGroupsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UnitIdentifierParam">UnitIdentifierParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertUnitGroupConversionInputParam">UpsertUnitGroupConversionInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertUnitGroupInputParam">UpsertUnitGroupInputParam</a>

Methods:

- <code title="post /v1/catalog/unit-groups/actions/bulk-upsert">client.Catalog.UnitGroups.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogUnitGroupActionBulkUpsertParams">CatalogUnitGroupActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Properties

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatePropertyRequestParam">CreatePropertyRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdatePropertyRequestParam">UpdatePropertyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Attribute">Attribute</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAttribute">ListAttribute</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProperty">ListProperty</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Property">Property</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyDeleteResponse">CatalogPropertyDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/properties">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyNewParams">CatalogPropertyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Property">Property</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties/{id}">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyGetParams">CatalogPropertyGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Property">Property</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/properties/{id}">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyUpdateParams">CatalogPropertyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Property">Property</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyListParams">CatalogPropertyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProperty">ListProperty</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/properties/{id}">client.Catalog.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyDeleteResponse">CatalogPropertyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Attributes

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAttributeRequestParam">CreateAttributeRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAttributeRequestParam">UpdateAttributeRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeDeleteResponse">CatalogPropertyAttributeDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/properties/{property_id}/attributes">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeNewParams">CatalogPropertyAttributeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Attribute">Attribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties/{property_id}/attributes/{id}">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeGetParams">CatalogPropertyAttributeGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Attribute">Attribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/properties/{property_id}/attributes/{id}">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeUpdateParams">CatalogPropertyAttributeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Attribute">Attribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/properties/{property_id}/attributes">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeListParams">CatalogPropertyAttributeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAttribute">ListAttribute</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/properties/{property_id}/attributes/{id}">client.Catalog.Properties.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeDeleteParams">CatalogPropertyAttributeDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyAttributeDeleteResponse">CatalogPropertyAttributeDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertPropertiesRequestParam">BulkUpsertPropertiesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertPropertyAttributeInputParam">UpsertPropertyAttributeInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertPropertyInputParam">UpsertPropertyInputParam</a>

Methods:

- <code title="post /v1/catalog/properties/actions/bulk-upsert">client.Catalog.Properties.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPropertyActionBulkUpsertParams">CatalogPropertyActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Item">Item</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemCategory">ItemCategory</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemLotDefault">ItemLotDefault</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListItem">ListItem</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Rate">Rate</a>

Methods:

- <code title="get /v1/catalog/items/{id}">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemGetParams">CatalogItemGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemListParams">CatalogItemListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListItem">ListItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/items/{id}/category/{category_id}">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemService.ChangeCategory">ChangeCategory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, categoryID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemChangeCategoryParams">CatalogItemChangeCategoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items/{id}/lot-default">client.Catalog.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemService.GetLotDefault">GetLotDefault</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemGetLotDefaultParams">CatalogItemGetLotDefaultParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemLotDefault">ItemLotDefault</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Inventory

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuantityInputParam">QuantityInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateItemInventoryRequestParam">UpdateItemInventoryRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ComputedQuantity">ComputedQuantity</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemInventory">ItemInventory</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemInventoryUpdateResponse">CatalogItemInventoryUpdateResponse</a>

Methods:

- <code title="patch /v1/catalog/items/{id}/inventory">client.Catalog.Items.Inventory.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemInventoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemInventoryUpdateParams">CatalogItemInventoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemInventoryUpdateResponse">CatalogItemInventoryUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/items/{id}/inventory">client.Catalog.Items.Inventory.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemInventoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemInventoryListParams">CatalogItemInventoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemInventory">ItemInventory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Attributes

Methods:

- <code title="put /v1/catalog/items/{id}/attributes/{attribute_id}">client.Catalog.Items.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemAttributeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attributeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemAttributeUpdateParams">CatalogItemAttributeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/items/{id}/attributes/{attribute_id}">client.Catalog.Items.Attributes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemAttributeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attributeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemAttributeDeleteParams">CatalogItemAttributeDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Item">Item</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkReconcileItemInputParam">BulkReconcileItemInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkReconcileItemsRequestParam">BulkReconcileItemsRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkReconcileItemsResponse">BulkReconcileItemsResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListReconcileErrorResult">ListReconcileErrorResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListReconciledItemResult">ListReconciledItemResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSkippedItemResult">ListSkippedItemResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReconcileErrorResult">ReconcileErrorResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReconciledItemResult">ReconciledItemResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SkippedItemResult">SkippedItemResult</a>

Methods:

- <code title="post /v1/catalog/items/actions/bulk-reconcile">client.Catalog.Items.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemActionService.BulkReconcile">BulkReconcile</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemActionBulkReconcileParams">CatalogItemActionBulkReconcileParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkReconcileItemsResponse">BulkReconcileItemsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ItemCategories

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateItemCategoryRequestParam">CreateItemCategoryRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateItemCategoryRequestParam">UpdateItemCategoryRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListItemCategory">ListItemCategory</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryDeleteResponse">CatalogItemCategoryDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryChangeUnitGroupResponse">CatalogItemCategoryChangeUnitGroupResponse</a>

Methods:

- <code title="post /v1/catalog/item-categories">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryNewParams">CatalogItemCategoryNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemCategory">ItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/item-categories/{id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryGetParams">CatalogItemCategoryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemCategory">ItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/item-categories/{id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryUpdateParams">CatalogItemCategoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ItemCategory">ItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/item-categories">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryListParams">CatalogItemCategoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListItemCategory">ListItemCategory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/item-categories/{id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryDeleteResponse">CatalogItemCategoryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/item-categories/{id}/unit-groups/{unit_group_id}">client.Catalog.ItemCategories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryService.ChangeUnitGroup">ChangeUnitGroup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, unitGroupID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryChangeUnitGroupParams">CatalogItemCategoryChangeUnitGroupParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryChangeUnitGroupResponse">CatalogItemCategoryChangeUnitGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Properties

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyUpdateResponse">CatalogItemCategoryPropertyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyDeleteResponse">CatalogItemCategoryPropertyDeleteResponse</a>

Methods:

- <code title="put /v1/catalog/item-categories/{id}/properties/{property_id}">client.Catalog.ItemCategories.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyUpdateParams">CatalogItemCategoryPropertyUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyUpdateResponse">CatalogItemCategoryPropertyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/item-categories/{id}/properties/{property_id}">client.Catalog.ItemCategories.Properties.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, propertyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyDeleteParams">CatalogItemCategoryPropertyDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryPropertyDeleteResponse">CatalogItemCategoryPropertyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertItemCategoriesRequestParam">BulkUpsertItemCategoriesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ObjectIdentifierParam">ObjectIdentifierParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertItemCategoryInputParam">UpsertItemCategoryInputParam</a>

Methods:

- <code title="post /v1/catalog/item-categories/actions/bulk-upsert">client.Catalog.ItemCategories.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogItemCategoryActionBulkUpsertParams">CatalogItemCategoryActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Materials

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateMaterialRequestParam">CreateMaterialRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuantityInputRequestParam">QuantityInputRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RateInputParam">RateInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateMaterialRequestParam">UpdateMaterialRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMaterial">ListMaterial</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Material">Material</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Quantity">Quantity</a>

Methods:

- <code title="post /v1/catalog/materials">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialNewParams">CatalogMaterialNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/materials/{id}">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialGetParams">CatalogMaterialGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/materials/{id}">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialUpdateParams">CatalogMaterialUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/materials">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialListParams">CatalogMaterialListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMaterial">ListMaterial</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/materials/{id}">client.Catalog.Materials.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Material">Material</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertMaterialsRequestParam">BulkUpsertMaterialsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertMaterialInputParam">UpsertMaterialInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertMaterialPropertyParam">UpsertMaterialPropertyParam</a>

Methods:

- <code title="post /v1/catalog/materials/actions/bulk-upsert">client.Catalog.Materials.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogMaterialActionBulkUpsertParams">CatalogMaterialActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Parts

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatePartRequestParam">CreatePartRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdatePartRequestParam">UpdatePartRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPart">ListPart</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Part">Part</a>

Methods:

- <code title="post /v1/catalog/parts">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartNewParams">CatalogPartNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/parts/{id}">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartGetParams">CatalogPartGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/parts/{id}">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartUpdateParams">CatalogPartUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/parts">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartListParams">CatalogPartListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPart">ListPart</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/parts/{id}">client.Catalog.Parts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertPartsRequestParam">BulkUpsertPartsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertPartInputParam">UpsertPartInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertPartPropertyParam">UpsertPartPropertyParam</a>

Methods:

- <code title="post /v1/catalog/parts/actions/bulk-upsert">client.Catalog.Parts.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogPartActionBulkUpsertParams">CatalogPartActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductLines

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateProductLineRequestParam">CreateProductLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateProductLineRequestParam">UpdateProductLineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductLine">ListProductLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductLine">ProductLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineDeleteResponse">CatalogProductLineDeleteResponse</a>

Methods:

- <code title="post /v1/catalog/product-lines">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineNewParams">CatalogProductLineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductLine">ProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/product-lines/{id}">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineGetParams">CatalogProductLineGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductLine">ProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/product-lines/{id}">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineUpdateParams">CatalogProductLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductLine">ProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/product-lines">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineListParams">CatalogProductLineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductLine">ListProductLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/product-lines/{id}">client.Catalog.ProductLines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineDeleteResponse">CatalogProductLineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertProductLinesRequestParam">BulkUpsertProductLinesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertProductLineInputParam">UpsertProductLineInputParam</a>

Methods:

- <code title="post /v1/catalog/product-lines/actions/bulk-upsert">client.Catalog.ProductLines.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductLineActionBulkUpsertParams">CatalogProductLineActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Products

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateProductRequestParam">CreateProductRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateProductRequestParam">UpdateProductRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProduct">ListProduct</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Product">Product</a>

Methods:

- <code title="post /v1/catalog/products">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductNewParams">CatalogProductNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/products/{id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductGetParams">CatalogProductGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/catalog/products/{id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductUpdateParams">CatalogProductUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/catalog/products">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductListParams">CatalogProductListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProduct">ListProduct</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/catalog/products/{id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductDeleteParams">CatalogProductDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/catalog/products/{id}/product-line/{product_line_id}">client.Catalog.Products.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductService.ChangeProductLine">ChangeProductLine</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, productLineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductChangeProductLineParams">CatalogProductChangeProductLineParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Product">Product</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertProductsRequestParam">BulkUpsertProductsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertProductInputParam">UpsertProductInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertProductPropertyParam">UpsertProductPropertyParam</a>

Methods:

- <code title="post /v1/catalog/products/actions/bulk-upsert">client.Catalog.Products.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CatalogProductActionBulkUpsertParams">CatalogProductActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AvailableTool">AvailableTool</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAvailableTool">ListAvailableTool</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListToolGroup">ListToolGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ToolGroup">ToolGroup</a>

Methods:

- <code title="get /v1/ai/tool-groups">client.AI.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIService.GetToolGroups">GetToolGroups</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIGetToolGroupsParams">AIGetToolGroupsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListToolGroup">ListToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/tools">client.AI.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIService.GetTools">GetTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIGetToolsParams">AIGetToolsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAvailableTool">ListAvailableTool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ConfigInputParam">ConfigInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAgentRequestParam">CreateAgentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ToolInputParam">ToolInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#TriggerConfigInputParam">TriggerConfigInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAgentRequestParam">UpdateAgentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAgentStatusRequestParam">UpdateAgentStatusRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinition">AgentDefinition</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinitionConfig">AgentDefinitionConfig</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinitionTool">AgentDefinitionTool</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentDefinition">ListAgentDefinition</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentDefinitionTool">ListAgentDefinitionTool</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#TriggerConfig">TriggerConfig</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentDeleteResponse">AIAgentDeleteResponse</a>

Methods:

- <code title="post /v1/ai/agents">client.AI.Agents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentNewParams">AIAgentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/agents/{id}">client.AI.Agents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentGetParams">AIAgentGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/ai/agents/{id}">client.AI.Agents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentUpdateParams">AIAgentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/agents">client.AI.Agents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentListParams">AIAgentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentDefinition">ListAgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/ai/agents/{id}">client.AI.Agents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentDeleteResponse">AIAgentDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/ai/agents/{id}/status">client.AI.Agents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentService.UpdateStatus">UpdateStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIAgentUpdateStatusParams">AIAgentUpdateStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentDefinition">AgentDefinition</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#TriggerRunRequestParam">TriggerRunRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentAction">AgentAction</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRun">AgentRun</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRunStep">AgentRunStep</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentAction">ListAgentAction</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentRun">ListAgentRun</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentRunStep">ListAgentRunStep</a>

Methods:

- <code title="post /v1/ai/runs">client.AI.Runs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunNewParams">AIRunNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/runs/{id}">client.AI.Runs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunGetParams">AIRunGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/runs">client.AI.Runs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunListParams">AIRunListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentRun">ListAgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ContinueRunRequestParam">ContinueRunRequestParam</a>

Methods:

- <code title="post /v1/ai/runs/{id}/actions/cancel">client.AI.Runs.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunActionService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunActionCancelParams">AIRunActionCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ai/runs/{id}/actions/continue">client.AI.Runs.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunActionService.Continue">Continue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunActionContinueParams">AIRunActionContinueParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/ai/runs/{id}/actions/retry">client.AI.Runs.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunActionService.Retry">Retry</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIRunActionRetryParams">AIRunActionRetryParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentRun">AgentRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Memories

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateMemoryRequestParam">CreateMemoryRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateMemoryRequestParam">UpdateMemoryRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentMemory">AgentMemory</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentMemory">ListAgentMemory</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryDeleteResponse">AIMemoryDeleteResponse</a>

Methods:

- <code title="post /v1/ai/memories">client.AI.Memories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryNewParams">AIMemoryNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentMemory">AgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/memories/{id}">client.AI.Memories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentMemory">AgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/ai/memories/{id}">client.AI.Memories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryUpdateParams">AIMemoryUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AgentMemory">AgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/ai/memories">client.AI.Memories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryListParams">AIMemoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAgentMemory">ListAgentMemory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/ai/memories/{id}">client.AI.Memories.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AIMemoryDeleteResponse">AIMemoryDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Messaging

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListActor">ListActor</a>

Methods:

- <code title="get /v1/messaging/contacts">client.Messaging.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingService.GetContacts">GetContacts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGetContactsParams">MessagingGetContactsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListActor">ListActor</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Notifications

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationTargetInputParam">NotificationTargetInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SendNotificationRequestParam">SendNotificationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListNotification">ListNotification</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListNotificationUnreadSummaryAccount">ListNotificationUnreadSummaryAccount</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Notification">Notification</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationSendResult">NotificationSendResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationUnreadCount">NotificationUnreadCount</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationUnreadSummary">NotificationUnreadSummary</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationUnreadSummaryAccount">NotificationUnreadSummaryAccount</a>

Methods:

- <code title="post /v1/messaging/notifications">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationNewParams">MessagingNotificationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationSendResult">NotificationSendResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications/{id}">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationGetParams">MessagingNotificationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationListParams">MessagingNotificationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListNotification">ListNotification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications/unread-count">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationService.GetUnreadCount">GetUnreadCount</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationUnreadCount">NotificationUnreadCount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/notifications/unread-summary">client.Messaging.Notifications.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationService.GetUnreadSummary">GetUnreadSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationUnreadSummary">NotificationUnreadSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionMarkAllSeenResponse">MessagingNotificationActionMarkAllSeenResponse</a>

Methods:

- <code title="post /v1/messaging/notifications/{id}/actions/dismiss">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionService.Dismiss">Dismiss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionDismissParams">MessagingNotificationActionDismissParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/notifications/actions/mark-all-seen">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionService.MarkAllSeen">MarkAllSeen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionMarkAllSeenResponse">MessagingNotificationActionMarkAllSeenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/notifications/{id}/actions/read">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionReadParams">MessagingNotificationActionReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/notifications/{id}/actions/seen">client.Messaging.Notifications.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionService.Seen">Seen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingNotificationActionSeenParams">MessagingNotificationActionSeenParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Announcements

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Announcement">Announcement</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAnnouncement">ListAnnouncement</a>

Methods:

- <code title="get /v1/messaging/announcements/{id}">client.Messaging.Announcements.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementGetParams">MessagingAnnouncementGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/announcements">client.Messaging.Announcements.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementListParams">MessagingAnnouncementListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAnnouncement">ListAnnouncement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="post /v1/messaging/announcements/{id}/actions/dismiss">client.Messaging.Announcements.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementActionService.Dismiss">Dismiss</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementActionDismissParams">MessagingAnnouncementActionDismissParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/announcements/{id}/actions/read">client.Messaging.Announcements.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementActionService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementActionReadParams">MessagingAnnouncementActionReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/announcements/{id}/actions/seen">client.Messaging.Announcements.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementActionService.Seen">Seen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingAnnouncementActionSeenParams">MessagingAnnouncementActionSeenParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Announcement">Announcement</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Conversations

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateConversationRequestParam">CreateConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateConversationRequestParam">UpdateConversationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ConversationParticipant">ConversationParticipant</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListConversation">ListConversation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListConversationParticipant">ListConversationParticipant</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessageAttachment">ListMessageAttachment</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessagingGroupMember">ListMessagingGroupMember</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessageAttachment">MessageAttachment</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroup">MessagingGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupMember">MessagingGroupMember</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReadCursor">ReadCursor</a>

Methods:

- <code title="post /v1/messaging/conversations">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationNewParams">MessagingConversationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations/{id}">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationGetParams">MessagingConversationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/messaging/conversations/{id}">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationUpdateParams">MessagingConversationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations">client.Messaging.Conversations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationListParams">MessagingConversationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListConversation">ListConversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AssignConversationRequestParam">AssignConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MarkConversationReadRequestParam">MarkConversationReadRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MuteConversationRequestParam">MuteConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReportConversationRequestParam">ReportConversationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SetLegalHoldRequestParam">SetLegalHoldRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SetWorkflowStatusRequestParam">SetWorkflowStatusRequestParam</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/actions/archive">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionArchiveParams">MessagingConversationActionArchiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/assign">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Assign">Assign</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionAssignParams">MessagingConversationActionAssignParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/hide">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Hide">Hide</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionHideParams">MessagingConversationActionHideParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/leave">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Leave">Leave</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionLeaveParams">MessagingConversationActionLeaveParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/mute">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Mute">Mute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionMuteParams">MessagingConversationActionMuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/read">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionReadParams">MessagingConversationActionReadParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/redact">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Redact">Redact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionRedactParams">MessagingConversationActionRedactParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/report">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Report">Report</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionReportParams">MessagingConversationActionReportParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/set-legal-hold">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.SetLegalHold">SetLegalHold</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionSetLegalHoldParams">MessagingConversationActionSetLegalHoldParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/set-status">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.SetStatus">SetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionSetStatusParams">MessagingConversationActionSetStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/unarchive">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Unarchive">Unarchive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionUnarchiveParams">MessagingConversationActionUnarchiveParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/unhide">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Unhide">Unhide</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionUnhideParams">MessagingConversationActionUnhideParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/conversations/{id}/actions/unmute">client.Messaging.Conversations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionService.Unmute">Unmute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationActionUnmuteParams">MessagingConversationActionUnmuteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Links

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AddConversationLinkRequestParam">AddConversationLinkRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ConversationLink">ConversationLink</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListConversationLink">ListConversationLink</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkDeleteResponse">MessagingConversationLinkDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/links">client.Messaging.Conversations.Links.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkNewParams">MessagingConversationLinkNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ConversationLink">ConversationLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations/{id}/links">client.Messaging.Conversations.Links.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkListParams">MessagingConversationLinkListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListConversationLink">ListConversationLink</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/conversations/{id}/links/{link_id}">client.Messaging.Conversations.Links.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, linkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkDeleteParams">MessagingConversationLinkDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationLinkDeleteResponse">MessagingConversationLinkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessageAttachmentInputParam">MessageAttachmentInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SendMessageRequestParam">SendMessageRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessage">ListMessage</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/messages">client.Messaging.Conversations.Messages.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationMessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationMessageNewParams">MessagingConversationMessageNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/conversations/{id}/messages">client.Messaging.Conversations.Messages.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationMessageListParams">MessagingConversationMessageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessage">ListMessage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Participants

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AddParticipantRequestParam">AddParticipantRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantDeleteResponse">MessagingConversationParticipantDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/participants">client.Messaging.Conversations.Participants.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantNewParams">MessagingConversationParticipantNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/conversations/{id}/participants/{pid}">client.Messaging.Conversations.Participants.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pid <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantDeleteParams">MessagingConversationParticipantDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantDeleteResponse">MessagingConversationParticipantDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateParticipantRoleRequestParam">UpdateParticipantRoleRequestParam</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/participants/{pid}/actions/set-role">client.Messaging.Conversations.Participants.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantActionService.SetRole">SetRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pid <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationParticipantActionSetRoleParams">MessagingConversationParticipantActionSetRoleParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Conversation">Conversation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Attachments

#### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAttachmentUploadURLRequestParam">CreateAttachmentUploadURLRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AttachmentUploadTarget">AttachmentUploadTarget</a>

Methods:

- <code title="post /v1/messaging/conversations/{id}/attachments/actions/upload-url">client.Messaging.Conversations.Attachments.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationAttachmentActionService.UploadURL">UploadURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingConversationAttachmentActionUploadURLParams">MessagingConversationAttachmentActionUploadURLParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AttachmentUploadTarget">AttachmentUploadTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateDraftRequestParam">UpdateDraftRequestParam</a>

Methods:

- <code title="patch /v1/messaging/messages/{id}">client.Messaging.Messages.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageUpdateParams">MessagingMessageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ApproveSendDraftRequestParam">ApproveSendDraftRequestParam</a>

Methods:

- <code title="post /v1/messaging/messages/{id}/actions/approve-send">client.Messaging.Messages.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageActionService.ApproveSend">ApproveSend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageActionApproveSendParams">MessagingMessageActionApproveSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/messages/{id}/actions/cancel">client.Messaging.Messages.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageActionService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageActionCancelParams">MessagingMessageActionCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/messaging/messages/{id}/actions/reject">client.Messaging.Messages.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageActionService.Reject">Reject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingMessageActionRejectParams">MessagingMessageActionRejectParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Message">Message</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Groups

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateMessagingGroupRequestParam">CreateMessagingGroupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateMessagingGroupRequestParam">UpdateMessagingGroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessagingGroup">ListMessagingGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupDeleteResponse">MessagingGroupDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/groups">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupNewParams">MessagingGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/groups/{id}">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/messaging/groups/{id}">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupUpdateParams">MessagingGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/groups">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessagingGroup">ListMessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/groups/{id}">client.Messaging.Groups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupDeleteResponse">MessagingGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Members

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AddMessagingGroupMemberRequestParam">AddMessagingGroupMemberRequestParam</a>

Methods:

- <code title="post /v1/messaging/groups/{id}/members">client.Messaging.Groups.Members.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupMemberService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupMemberNewParams">MessagingGroupMemberNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/groups/{id}/members/{member_id}">client.Messaging.Groups.Members.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupMemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroupMemberDeleteParams">MessagingGroupMemberDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingGroup">MessagingGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Blocks

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BlockRequestParam">BlockRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#LocationTypeCode">LocationTypeCode</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountUser">AccountUser</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Consumption">Consumption</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Department">Department</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListConsumption">ListConsumption</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListLocation">ListLocation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachine">ListMachine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessagingBlock">ListMessagingBlock</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionStep">ListProductionStep</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScanningStation">ListScanningStation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#LocationTypeCode">LocationTypeCode</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Machine">Machine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlock">MessagingBlock</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionOutput">ProductionOutput</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionStep">ProductionStep</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScanningStation">ScanningStation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#User">User</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockDeleteResponse">MessagingBlockDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/blocks">client.Messaging.Blocks.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockNewParams">MessagingBlockNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlock">MessagingBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/blocks">client.Messaging.Blocks.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockListParams">MessagingBlockListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMessagingBlock">ListMessagingBlock</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/blocks/{id}">client.Messaging.Blocks.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingBlockDeleteResponse">MessagingBlockDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Preferences

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertNotificationPreferenceRequestParam">UpsertNotificationPreferenceRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListNotificationPreference">ListNotificationPreference</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationPreference">NotificationPreference</a>

Methods:

- <code title="put /v1/messaging/preferences">client.Messaging.Preferences.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingPreferenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingPreferenceUpdateParams">MessagingPreferenceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationPreference">NotificationPreference</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/preferences">client.Messaging.Preferences.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingPreferenceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListNotificationPreference">ListNotificationPreference</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailDomains

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateEmailDomainRequestParam">CreateEmailDomainRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailDomain">EmailDomain</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEmailDomain">ListEmailDomain</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainDeleteResponse">MessagingEmailDomainDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/email-domains">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainNewParams">MessagingEmailDomainNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailDomain">EmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-domains/{id}">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailDomain">EmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-domains">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEmailDomain">ListEmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/email-domains/{id}">client.Messaging.EmailDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainDeleteResponse">MessagingEmailDomainDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="post /v1/messaging/email-domains/{id}/actions/verify">client.Messaging.EmailDomains.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailDomainActionService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailDomain">EmailDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailInboxes

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateEmailInboxRequestParam">CreateEmailInboxRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateEmailInboxRequestParam">UpdateEmailInboxRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailInbox">EmailInbox</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEmailInbox">ListEmailInbox</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxDeleteResponse">MessagingEmailInboxDeleteResponse</a>

Methods:

- <code title="post /v1/messaging/email-inboxes">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxNewParams">MessagingEmailInboxNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailInbox">EmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-inboxes/{id}">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxGetParams">MessagingEmailInboxGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailInbox">EmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/messaging/email-inboxes/{id}">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxUpdateParams">MessagingEmailInboxUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailInbox">EmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-inboxes">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxListParams">MessagingEmailInboxListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListEmailInbox">ListEmailInbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/email-inboxes/{id}">client.Messaging.EmailInboxes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailInboxDeleteResponse">MessagingEmailInboxDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EmailSender

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SetEmailSenderRequestParam">SetEmailSenderRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailSender">EmailSender</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailSenderDeleteResponse">MessagingEmailSenderDeleteResponse</a>

Methods:

- <code title="put /v1/messaging/email-sender">client.Messaging.EmailSender.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailSenderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailSenderUpdateParams">MessagingEmailSenderUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailSender">EmailSender</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/messaging/email-sender">client.Messaging.EmailSender.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailSenderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#EmailSender">EmailSender</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/messaging/email-sender">client.Messaging.EmailSender.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailSenderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MessagingEmailSenderDeleteResponse">MessagingEmailSenderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sales

## AccountGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAccountGroupRequestParam">CreateAccountGroupRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAccountGroupRequestParam">UpdateAccountGroupRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountGroup">AccountGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountGroup">ListAccountGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupDeleteResponse">SaleAccountGroupDeleteResponse</a>

Methods:

- <code title="post /v1/sales/account-groups">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupNewParams">SaleAccountGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountGroup">AccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-groups/{id}">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountGroup">AccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/account-groups/{id}">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupUpdateParams">SaleAccountGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountGroup">AccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-groups">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupListParams">SaleAccountGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountGroup">ListAccountGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/account-groups/{id}">client.Sales.AccountGroups.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountGroupDeleteResponse">SaleAccountGroupDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountPrices

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAccountPriceRequestParam">CreateAccountPriceRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAccountPriceRequestParam">UpdateAccountPriceRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountPrice">AccountPrice</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Carrier">Carrier</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CustomerContactInfo">CustomerContactInfo</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CustomerDefaults">CustomerDefaults</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CustomerFreightPreferences">CustomerFreightPreferences</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CustomerNotificationPreferences">CustomerNotificationPreferences</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountPrice">ListAccountPrice</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListCustomer">ListCustomer</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListServiceLevel">ListServiceLevel</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PaymentTerm">PaymentTerm</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Priority">Priority</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ServiceLevel">ServiceLevel</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ShippingTerm">ShippingTerm</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceDeleteResponse">SaleAccountPriceDeleteResponse</a>

Methods:

- <code title="post /v1/sales/account-prices">client.Sales.AccountPrices.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceNewParams">SaleAccountPriceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountPrice">AccountPrice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-prices/{id}">client.Sales.AccountPrices.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceGetParams">SaleAccountPriceGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountPrice">AccountPrice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/account-prices/{id}">client.Sales.AccountPrices.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceUpdateParams">SaleAccountPriceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountPrice">AccountPrice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-prices">client.Sales.AccountPrices.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceListParams">SaleAccountPriceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountPrice">ListAccountPrice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/account-prices/{id}">client.Sales.AccountPrices.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceDeleteResponse">SaleAccountPriceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ExportPriceListRequestParam">ExportPriceListRequestParam</a>

Methods:

- <code title="post /v1/sales/account-prices/actions/export-price-list">client.Sales.AccountPrices.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceActionService.ExportPriceList">ExportPriceList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountPriceActionExportPriceListParams">SaleAccountPriceActionExportPriceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Addresses

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AddressInputParam">AddressInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAddressRequestParam">UpdateAddressRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAddress">ListAddress</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressDeleteResponse">SaleAddressDeleteResponse</a>

Methods:

- <code title="post /v1/sales/addresses">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressNewParams">SaleAddressNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/addresses/{id}">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/addresses/{id}">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressUpdateParams">SaleAddressUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Address">Address</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/addresses">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressListParams">SaleAddressListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAddress">ListAddress</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/addresses/{id}">client.Sales.Addresses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAddressDeleteResponse">SaleAddressDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountStatuses

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountStatus">ListAccountStatus</a>

Methods:

- <code title="get /v1/sales/account-statuses/{id}">client.Sales.AccountStatuses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountStatusGetParams">SaleAccountStatusGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountStatus">AccountStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-statuses">client.Sales.AccountStatuses.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountStatusService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountStatusListParams">SaleAccountStatusListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountStatus">ListAccountStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountUsers

### SalesTargets

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateSalesTargetRequestParam">CreateSalesTargetRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertSalesTargetRequestParam">UpsertSalesTargetRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesTarget">ListSalesTarget</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesTarget">SalesTarget</a>

Methods:

- <code title="post /v1/sales/account-users/{id}/sales-targets">client.Sales.AccountUsers.SalesTargets.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountUserSalesTargetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountUserSalesTargetNewParams">SaleAccountUserSalesTargetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesTarget">SalesTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/account-users/{id}/sales-targets/{target_id}">client.Sales.AccountUsers.SalesTargets.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountUserSalesTargetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, targetID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountUserSalesTargetUpdateParams">SaleAccountUserSalesTargetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesTarget">SalesTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/account-users/{id}/sales-targets">client.Sales.AccountUsers.SalesTargets.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountUserSalesTargetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleAccountUserSalesTargetListParams">SaleAccountUserSalesTargetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesTarget">ListSalesTarget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Priorities

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPriority">ListPriority</a>

Methods:

- <code title="get /v1/sales/priorities/{id}">client.Sales.Priorities.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalePriorityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalePriorityGetParams">SalePriorityGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Priority">Priority</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/priorities">client.Sales.Priorities.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalePriorityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalePriorityListParams">SalePriorityListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPriority">ListPriority</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Customers

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateCustomerRequestParam">CreateCustomerRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateCustomerRequestParam">UpdateCustomerRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CustomerLeadTime">CustomerLeadTime</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerDeleteResponse">SaleCustomerDeleteResponse</a>

Methods:

- <code title="post /v1/sales/customers">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerNewParams">SaleCustomerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerGetParams">SaleCustomerGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerUpdateParams">SaleCustomerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerListParams">SaleCustomerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListCustomer">ListCustomer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/customers/{id}">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerDeleteResponse">SaleCustomerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/customers/{id}/lead-time">client.Sales.Customers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerService.GetLeadTime">GetLeadTime</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerGetLeadTimeParams">SaleCustomerGetLeadTimeParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CustomerLeadTime">CustomerLeadTime</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MergeCustomersRequestParam">MergeCustomersRequestParam</a>

Methods:

- <code title="post /v1/sales/customers/{id}/actions/merge">client.Sales.Customers.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerActionService.Merge">Merge</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleCustomerActionMergeParams">SaleCustomerActionMergeParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Contacts

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FindContactByEmailRequestParam">FindContactByEmailRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ContactMatch">ContactMatch</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListContactMatch">ListContactMatch</a>

Methods:

- <code title="post /v1/sales/contacts/actions/find-by-email">client.Sales.Contacts.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleContactActionService.FindByEmail">FindByEmail</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleContactActionFindByEmailParams">SaleContactActionFindByEmailParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListContactMatch">ListContactMatch</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OrderDiscounts

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateOrderDiscountRequestParam">CreateOrderDiscountRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateOrderDiscountRequestParam">UpdateOrderDiscountRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOrderDiscount">ListOrderDiscount</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderDiscount">OrderDiscount</a>

Methods:

- <code title="post /v1/sales/order-discounts">client.Sales.OrderDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountNewParams">SaleOrderDiscountNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderDiscount">OrderDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/order-discounts/{id}">client.Sales.OrderDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderDiscount">OrderDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/order-discounts/{id}">client.Sales.OrderDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountUpdateParams">SaleOrderDiscountUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderDiscount">OrderDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/order-discounts">client.Sales.OrderDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountListParams">SaleOrderDiscountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOrderDiscount">ListOrderDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/order-discounts/{id}">client.Sales.OrderDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderDiscount">OrderDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FindOrderDiscountByCodeRequestParam">FindOrderDiscountByCodeRequestParam</a>

Methods:

- <code title="post /v1/sales/order-discounts/actions/find-by-code">client.Sales.OrderDiscounts.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountActionService.FindByCode">FindByCode</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleOrderDiscountActionFindByCodeParams">SaleOrderDiscountActionFindByCodeParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderDiscount">OrderDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## SalesOrders

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CheckoutSalesOrderRequestParam">CheckoutSalesOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateSalesOrderLineInputParam">CreateSalesOrderLineInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateSalesOrderRequestParam">CreateSalesOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderLineInputParam">QuoteSalesOrderLineInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderPricesRequestParam">QuoteSalesOrderPricesRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderEmailContactInputParam">SalesOrderEmailContactInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateSalesOrderRequestParam">UpdateSalesOrderRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CheckoutSalesOrderResponse">CheckoutSalesOrderResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Commitment">Commitment</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ComputedRate">ComputedRate</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatedBy">CreatedBy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Freight">Freight</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListQuotedSalesOrderLine">ListQuotedSalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListRecord">ListRecord</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesOrder">ListSalesOrder</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesOrderLine">ListSalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesOrderStatus">ListSalesOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OrderContact">OrderContact</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderPricesResponse">QuoteSalesOrderPricesResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuotedSalesOrderLine">QuotedSalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Record">Record</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderLine">SalesOrderLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderRelated">SalesOrderRelated</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderStageTotal">SalesOrderStageTotal</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderStatus">SalesOrderStatus</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderTotals">SalesOrderTotals</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderDeleteResponse">SaleSalesOrderDeleteResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderNewParams">SaleSalesOrderNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/sales-orders/{id}">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderGetParams">SaleSalesOrderGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/sales-orders/{id}">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderUpdateParams">SaleSalesOrderUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/sales-orders">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderListParams">SaleSalesOrderListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesOrder">ListSalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/sales-orders/{id}">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderDeleteResponse">SaleSalesOrderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/{id}/checkout">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.Checkout">Checkout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderCheckoutParams">SaleSalesOrderCheckoutParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CheckoutSalesOrderResponse">CheckoutSalesOrderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/price-quote">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.PriceQuote">PriceQuote</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderPriceQuoteParams">SaleSalesOrderPriceQuoteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderPricesResponse">QuoteSalesOrderPricesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/sales-orders/statuses">client.Sales.SalesOrders.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderService.GetStatuses">GetStatuses</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderGetStatusesParams">SaleSalesOrderGetStatusesParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSalesOrderStatus">ListSalesOrderStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkDeleteSalesOrdersRequestParam">BulkDeleteSalesOrdersRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IssueSalesOrderRequestParam">IssueSalesOrderRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderCommitmentRequestParam">QuoteSalesOrderCommitmentRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CommitmentQuoteStep">CommitmentQuoteStep</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionRun">ProductionRun</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderCommitmentResponse">QuoteSalesOrderCommitmentResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderFreightResponse">QuoteSalesOrderFreightResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionBulkDeleteResponse">SaleSalesOrderActionBulkDeleteResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders/actions/bulk-delete">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.BulkDelete">BulkDelete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionBulkDeleteParams">SaleSalesOrderActionBulkDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionBulkDeleteResponse">SaleSalesOrderActionBulkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/close">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/{id}/actions/create-production-run">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.NewProductionRun">NewProductionRun</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionNewProductionRunParams">SaleSalesOrderActionNewProductionRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionRun">ProductionRun</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/issue">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.Issue">Issue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionIssueParams">SaleSalesOrderActionIssueParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/open">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.Open">Open</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/actions/quote-commitment">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.QuoteCommitment">QuoteCommitment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionQuoteCommitmentParams">SaleSalesOrderActionQuoteCommitmentParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderCommitmentResponse">QuoteSalesOrderCommitmentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/sales/sales-orders/{id}/actions/quote-freight">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.QuoteFreight">QuoteFreight</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#QuoteSalesOrderFreightResponse">QuoteSalesOrderFreightResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/sales/sales-orders/{id}/actions/unissue">client.Sales.SalesOrders.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderActionService.Unissue">Unissue</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrder">SalesOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Lines

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateSalesOrderLineRequestParam">CreateSalesOrderLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateSalesOrderLineRequestParam">UpdateSalesOrderLineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineDeleteResponse">SaleSalesOrderLineDeleteResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders/{id}/lines">client.Sales.SalesOrders.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineNewParams">SaleSalesOrderLineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderLine">SalesOrderLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/sales-orders/{id}/lines/{line_id}">client.Sales.SalesOrders.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineUpdateParams">SaleSalesOrderLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SalesOrderLine">SalesOrderLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/sales-orders/{id}/lines/{line_id}">client.Sales.SalesOrders.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineDeleteParams">SaleSalesOrderLineDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineDeleteResponse">SaleSalesOrderLineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReorderSalesOrderLinesRequestParam">ReorderSalesOrderLinesRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineActionReorderResponse">SaleSalesOrderLineActionReorderResponse</a>

Methods:

- <code title="post /v1/sales/sales-orders/{id}/lines/actions/reorder">client.Sales.SalesOrders.Lines.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineActionService.Reorder">Reorder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineActionReorderParams">SaleSalesOrderLineActionReorderParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleSalesOrderLineActionReorderResponse">SaleSalesOrderLineActionReorderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## VolumeDiscounts

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateVolumeDiscountRequestParam">CreateVolumeDiscountRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateVolumeDiscountTierInputParam">CreateVolumeDiscountTierInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateVolumeDiscountRequestParam">UpdateVolumeDiscountRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateVolumeDiscountTierInputParam">UpdateVolumeDiscountTierInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListVolumeDiscount">ListVolumeDiscount</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListVolumeDiscountTier">ListVolumeDiscountTier</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#VolumeDiscount">VolumeDiscount</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#VolumeDiscountTier">VolumeDiscountTier</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountDeleteResponse">SaleVolumeDiscountDeleteResponse</a>

Methods:

- <code title="post /v1/sales/volume-discounts">client.Sales.VolumeDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountNewParams">SaleVolumeDiscountNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#VolumeDiscount">VolumeDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/volume-discounts/{id}">client.Sales.VolumeDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountGetParams">SaleVolumeDiscountGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#VolumeDiscount">VolumeDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/sales/volume-discounts/{id}">client.Sales.VolumeDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountUpdateParams">SaleVolumeDiscountUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#VolumeDiscount">VolumeDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/sales/volume-discounts">client.Sales.VolumeDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountListParams">SaleVolumeDiscountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListVolumeDiscount">ListVolumeDiscount</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/sales/volume-discounts/{id}">client.Sales.VolumeDiscounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SaleVolumeDiscountDeleteResponse">SaleVolumeDiscountDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Finance

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AdjustmentType">AdjustmentType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAdjustmentType">ListAdjustmentType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListTransactionMethod">ListTransactionMethod</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListTransactionType">ListTransactionType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#TransactionMethod">TransactionMethod</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#TransactionType">TransactionType</a>

Methods:

- <code title="get /v1/finance/adjustment-types">client.Finance.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinanceService.GetAdjustmentTypes">GetAdjustmentTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinanceGetAdjustmentTypesParams">FinanceGetAdjustmentTypesParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAdjustmentType">ListAdjustmentType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/transaction-methods">client.Finance.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinanceService.GetTransactionMethods">GetTransactionMethods</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinanceGetTransactionMethodsParams">FinanceGetTransactionMethodsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListTransactionMethod">ListTransactionMethod</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/transaction-types">client.Finance.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinanceService.GetTransactionTypes">GetTransactionTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinanceGetTransactionTypesParams">FinanceGetTransactionTypesParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListTransactionType">ListTransactionType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PaymentTerms

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatePaymentTermRequestParam">CreatePaymentTermRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdatePaymentTermRequestParam">UpdatePaymentTermRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPaymentTerm">ListPaymentTerm</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermDeleteResponse">FinancePaymentTermDeleteResponse</a>

Methods:

- <code title="post /v1/finance/payment-terms">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermNewParams">FinancePaymentTermNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PaymentTerm">PaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/payment-terms/{id}">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermGetParams">FinancePaymentTermGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PaymentTerm">PaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/finance/payment-terms/{id}">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermUpdateParams">FinancePaymentTermUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PaymentTerm">PaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/finance/payment-terms">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermListParams">FinancePaymentTermListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPaymentTerm">ListPaymentTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/finance/payment-terms/{id}">client.Finance.PaymentTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FinancePaymentTermDeleteResponse">FinancePaymentTermDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Operations

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DemandOverrideType">DemandOverrideType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDemandOverrideType">ListDemandOverrideType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachineDowntimeReason">ListMachineDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachineStatus">ListMachineStatus</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleDeviationType">ListScheduleDeviationType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineCampaign">MachineCampaign</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeReason">MachineDowntimeReason</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeReasonSummary">MachineDowntimeReasonSummary</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeSummary">MachineDowntimeSummary</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineStatus">MachineStatus</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleDeviationType">ScheduleDeviationType</a>

Methods:

- <code title="get /v1/operations/demand-override-types">client.Operations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationService.GetDemandOverrideTypes">GetDemandOverrideTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDemandOverrideType">ListDemandOverrideType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-downtime-reasons">client.Operations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationService.GetMachineDowntimeReasons">GetMachineDowntimeReasons</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachineDowntimeReason">ListMachineDowntimeReason</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-status">client.Operations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationService.GetMachineStatus">GetMachineStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationGetMachineStatusParams">OperationGetMachineStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachineStatus">ListMachineStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/schedule-deviation-types">client.Operations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationService.GetScheduleDeviationTypes">GetScheduleDeviationTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleDeviationType">ListScheduleDeviationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ShippingTerms

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateShippingTermRequestParam">CreateShippingTermRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateShippingTermRequestParam">UpdateShippingTermRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListShippingTerm">ListShippingTerm</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermDeleteResponse">OperationShippingTermDeleteResponse</a>

Methods:

- <code title="post /v1/operations/shipping-terms">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermNewParams">OperationShippingTermNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ShippingTerm">ShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/shipping-terms/{id}">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermGetParams">OperationShippingTermGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ShippingTerm">ShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/shipping-terms/{id}">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermUpdateParams">OperationShippingTermUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ShippingTerm">ShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/shipping-terms">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermListParams">OperationShippingTermListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListShippingTerm">ListShippingTerm</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/shipping-terms/{id}">client.Operations.ShippingTerms.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShippingTermDeleteResponse">OperationShippingTermDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Carriers

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateCarrierRequestParam">CreateCarrierRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateCarrierRequestParam">UpdateCarrierRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListCarrier">ListCarrier</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierDeleteResponse">OperationCarrierDeleteResponse</a>

Methods:

- <code title="post /v1/operations/carriers">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierNewParams">OperationCarrierNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Carrier">Carrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers/{id}">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierGetParams">OperationCarrierGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Carrier">Carrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/carriers/{id}">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierUpdateParams">OperationCarrierUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Carrier">Carrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierListParams">OperationCarrierListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListCarrier">ListCarrier</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/carriers/{id}">client.Operations.Carriers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierDeleteResponse">OperationCarrierDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### ServiceLevels

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateServiceLevelRequestParam">CreateServiceLevelRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateServiceLevelRequestParam">UpdateServiceLevelRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelDeleteResponse">OperationCarrierServiceLevelDeleteResponse</a>

Methods:

- <code title="post /v1/operations/carriers/{carrier_id}/service-levels">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, carrierID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelNewParams">OperationCarrierServiceLevelNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ServiceLevel">ServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers/{carrier_id}/service-levels/{id}">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelGetParams">OperationCarrierServiceLevelGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ServiceLevel">ServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/carriers/{carrier_id}/service-levels/{id}">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelUpdateParams">OperationCarrierServiceLevelUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ServiceLevel">ServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/carriers/{carrier_id}/service-levels">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, carrierID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelListParams">OperationCarrierServiceLevelListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListServiceLevel">ListServiceLevel</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/carriers/{carrier_id}/service-levels/{id}">client.Operations.Carriers.ServiceLevels.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelDeleteParams">OperationCarrierServiceLevelDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationCarrierServiceLevelDeleteResponse">OperationCarrierServiceLevelDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Departments

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateDepartmentRequestParam">CreateDepartmentRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DepartmentRateInputParam">DepartmentRateInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateDepartmentRequestParam">UpdateDepartmentRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDepartment">ListDepartment</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentDeleteResponse">OperationDepartmentDeleteResponse</a>

Methods:

- <code title="post /v1/operations/departments">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentNewParams">OperationDepartmentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Department">Department</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/departments/{id}">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentGetParams">OperationDepartmentGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Department">Department</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/departments/{id}">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentUpdateParams">OperationDepartmentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Department">Department</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/departments">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentListParams">OperationDepartmentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDepartment">ListDepartment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/departments/{id}">client.Operations.Departments.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDepartmentDeleteResponse">OperationDepartmentDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InventoryChangeLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#InventoryChangeLog">InventoryChangeLog</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListInventoryChangeLog">ListInventoryChangeLog</a>

Methods:

- <code title="get /v1/operations/inventory-change-logs/{id}">client.Operations.InventoryChangeLogs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationInventoryChangeLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationInventoryChangeLogGetParams">OperationInventoryChangeLogGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#InventoryChangeLog">InventoryChangeLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/inventory-change-logs">client.Operations.InventoryChangeLogs.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationInventoryChangeLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationInventoryChangeLogListParams">OperationInventoryChangeLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListInventoryChangeLog">ListInventoryChangeLog</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FileDownload">FileDownload</a>

Methods:

- <code title="get /v1/operations/inventory-change-logs/actions/export">client.Operations.InventoryChangeLogs.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationInventoryChangeLogActionService.Export">Export</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationInventoryChangeLogActionExportParams">OperationInventoryChangeLogActionExportParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FileDownload">FileDownload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Machines

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateMachineRequestParam">CreateMachineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateMachineRequestParam">UpdateMachineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDeleteResponse">OperationMachineDeleteResponse</a>

Methods:

- <code title="post /v1/operations/machines">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineNewParams">OperationMachineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Machine">Machine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machines/{id}">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineGetParams">OperationMachineGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Machine">Machine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/machines/{id}">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineUpdateParams">OperationMachineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Machine">Machine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machines">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineListParams">OperationMachineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachine">ListMachine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/machines/{id}">client.Operations.Machines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDeleteResponse">OperationMachineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## MachineDowntimeEvents

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateMachineDowntimeEventRequestParam">CreateMachineDowntimeEventRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateMachineDowntimeEventRequestParam">UpdateMachineDowntimeEventRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachineDowntimeEvent">ListMachineDowntimeEvent</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeEvent">MachineDowntimeEvent</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventDeleteResponse">OperationMachineDowntimeEventDeleteResponse</a>

Methods:

- <code title="post /v1/operations/machine-downtime-events">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventNewParams">OperationMachineDowntimeEventNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeEvent">MachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-downtime-events/{id}">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventGetParams">OperationMachineDowntimeEventGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeEvent">MachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/machine-downtime-events/{id}">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventUpdateParams">OperationMachineDowntimeEventUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#MachineDowntimeEvent">MachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/machine-downtime-events">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventListParams">OperationMachineDowntimeEventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListMachineDowntimeEvent">ListMachineDowntimeEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/machine-downtime-events/{id}">client.Operations.MachineDowntimeEvents.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationMachineDowntimeEventDeleteResponse">OperationMachineDowntimeEventDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DemandOverrides

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateDemandOverrideRequestParam">CreateDemandOverrideRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateDemandOverrideRequestParam">UpdateDemandOverrideRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DemandOverride">DemandOverride</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDemandOverride">ListDemandOverride</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideDeleteResponse">OperationDemandOverrideDeleteResponse</a>

Methods:

- <code title="post /v1/operations/demand-overrides">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideNewParams">OperationDemandOverrideNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DemandOverride">DemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/demand-overrides/{id}">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideGetParams">OperationDemandOverrideGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DemandOverride">DemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/demand-overrides/{id}">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideUpdateParams">OperationDemandOverrideUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DemandOverride">DemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/demand-overrides">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideListParams">OperationDemandOverrideListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDemandOverride">ListDemandOverride</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/demand-overrides/{id}">client.Operations.DemandOverrides.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationDemandOverrideDeleteResponse">OperationDemandOverrideDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductionSchedules

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#GenerateProductionScheduleRequestParam">GenerateProductionScheduleRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionSchedule">ListProductionSchedule</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleDerivedLine">ListProductionScheduleDerivedLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleDeviation">ListProductionScheduleDeviation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleFinishedPolicy">ListProductionScheduleFinishedPolicy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleFinishingLine">ListProductionScheduleFinishingLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleItemPolicy">ListProductionScheduleItemPolicy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListReleaseScheduleBatch">ListReleaseScheduleBatch</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListReleasedScheduleLine">ListReleasedScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleAppliedOverride">ListScheduleAppliedOverride</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleAtRiskOrder">ListScheduleAtRiskOrder</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleOrderCoverage">ListScheduleOrderCoverage</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleOrderCoverageLine">ListScheduleOrderCoverageLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleDerivedLine">ProductionScheduleDerivedLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleDeviation">ProductionScheduleDeviation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleFinishedPolicy">ProductionScheduleFinishedPolicy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleFinishingLine">ProductionScheduleFinishingLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleItemPolicy">ProductionScheduleItemPolicy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleaseScheduleBatch">ReleaseScheduleBatch</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleaseScheduleWeekPreview">ReleaseScheduleWeekPreview</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleasedScheduleLine">ReleasedScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleAppliedOverride">ScheduleAppliedOverride</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleAtRiskOrder">ScheduleAtRiskOrder</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleDiagnostics">ScheduleDiagnostics</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleFinishingDiagnostics">ScheduleFinishingDiagnostics</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleOrderCoverage">ScheduleOrderCoverage</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleOrderCoverageLine">ScheduleOrderCoverageLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleDeleteResponse">OperationProductionScheduleDeleteResponse</a>

Methods:

- <code title="post /v1/operations/production-schedules">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleNewParams">OperationProductionScheduleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleListParams">OperationProductionScheduleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionSchedule">ListProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedules/{id}">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleDeleteResponse">OperationProductionScheduleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/at-risk-orders">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetAtRiskOrders">GetAtRiskOrders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleOrderCoverage">ListScheduleOrderCoverage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/current">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetCurrent">GetCurrent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/derived-lines">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetDerivedLines">GetDerivedLines</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleGetDerivedLinesParams">OperationProductionScheduleGetDerivedLinesParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleDerivedLine">ListProductionScheduleDerivedLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/deviations">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetDeviations">GetDeviations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleGetDeviationsParams">OperationProductionScheduleGetDeviationsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleDeviation">ListProductionScheduleDeviation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/finished-policies">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetFinishedPolicies">GetFinishedPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleFinishedPolicy">ListProductionScheduleFinishedPolicy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/finishing-lines">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetFinishingLines">GetFinishingLines</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleGetFinishingLinesParams">OperationProductionScheduleGetFinishingLinesParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleFinishingLine">ListProductionScheduleFinishingLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/item-policies">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetItemPolicies">GetItemPolicies</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleItemPolicy">ListProductionScheduleItemPolicy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/week-release-preview">client.Operations.ProductionSchedules.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleService.GetWeekReleasePreview">GetWeekReleasePreview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleGetWeekReleasePreviewParams">OperationProductionScheduleGetWeekReleasePreviewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleaseScheduleWeekPreview">ReleaseScheduleWeekPreview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Lines

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateProductionScheduleLineRequestParam">CreateProductionScheduleLineRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateProductionScheduleLineRequestParam">UpdateProductionScheduleLineRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleLine">ListProductionScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleLine">ProductionScheduleLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineDeleteResponse">OperationProductionScheduleLineDeleteResponse</a>

Methods:

- <code title="post /v1/operations/production-schedules/{id}/lines">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineNewParams">OperationProductionScheduleLineNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleLine">ProductionScheduleLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/production-schedules/{id}/lines/{line_id}">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineUpdateParams">OperationProductionScheduleLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleLine">ProductionScheduleLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedules/{id}/lines">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineListParams">OperationProductionScheduleLineListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleLine">ListProductionScheduleLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedules/{id}/lines/{line_id}">client.Operations.ProductionSchedules.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, lineID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineDeleteParams">OperationProductionScheduleLineDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleLineDeleteResponse">OperationProductionScheduleLineDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PreviewProductionScheduleRequestParam">PreviewProductionScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PreviewRegenerateProductionScheduleRequestParam">PreviewRegenerateProductionScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RegenerateProductionScheduleRequestParam">RegenerateProductionScheduleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleaseProductionScheduleWeekRequestParam">ReleaseProductionScheduleWeekRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleCampaign">ListScheduleCampaign</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleDiffLine">ListScheduleDiffLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListSchedulePolicy">ListSchedulePolicy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScheduleProjection">ListScheduleProjection</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedulePreview">ProductionSchedulePreview</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleRegeneratePreview">ProductionScheduleRegeneratePreview</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleaseScheduleWeekResult">ReleaseScheduleWeekResult</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleCampaign">ScheduleCampaign</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleDiffLine">ScheduleDiffLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SchedulePolicy">SchedulePolicy</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScheduleProjection">ScheduleProjection</a>

Methods:

- <code title="put /v1/operations/production-schedules/{id}/actions/archive">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionService.Archive">Archive</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/actions/preview">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionPreviewParams">OperationProductionScheduleActionPreviewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedulePreview">ProductionSchedulePreview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/{id}/actions/preview-regenerate">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionService.PreviewRegenerate">PreviewRegenerate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionPreviewRegenerateParams">OperationProductionScheduleActionPreviewRegenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleRegeneratePreview">ProductionScheduleRegeneratePreview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/{id}/actions/publish">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionService.Publish">Publish</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedules/{id}/actions/regenerate">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionService.Regenerate">Regenerate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionRegenerateParams">OperationProductionScheduleActionRegenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionSchedule">ProductionSchedule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/operations/production-schedules/{id}/actions/release-week">client.Operations.ProductionSchedules.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionService.ReleaseWeek">ReleaseWeek</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleActionReleaseWeekParams">OperationProductionScheduleActionReleaseWeekParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ReleaseScheduleWeekResult">ReleaseScheduleWeekResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProductionScheduleSettings

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateProductionScheduleSettingsRequestParam">UpdateProductionScheduleSettingsRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleSettings">ProductionScheduleSettings</a>

Methods:

- <code title="put /v1/operations/production-schedule-settings">client.Operations.ProductionScheduleSettings.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingUpdateParams">OperationProductionScheduleSettingUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleSettings">ProductionScheduleSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedule-settings">client.Operations.ProductionScheduleSettings.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleSettings">ProductionScheduleSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Resources

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertResourceSettingRequestParam">UpsertResourceSettingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleResourceSetting">ListProductionScheduleResourceSetting</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleResourceSetting">ProductionScheduleResourceSetting</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingResourceDeleteResponse">OperationProductionScheduleSettingResourceDeleteResponse</a>

Methods:

- <code title="put /v1/operations/production-schedule-settings/resources">client.Operations.ProductionScheduleSettings.Resources.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingResourceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingResourceUpdateParams">OperationProductionScheduleSettingResourceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleResourceSetting">ProductionScheduleResourceSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedule-settings/resources">client.Operations.ProductionScheduleSettings.Resources.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingResourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleResourceSetting">ListProductionScheduleResourceSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedule-settings/resources/{id}">client.Operations.ProductionScheduleSettings.Resources.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingResourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingResourceDeleteResponse">OperationProductionScheduleSettingResourceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Items

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertItemSettingRequestParam">UpsertItemSettingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleItemSetting">ListProductionScheduleItemSetting</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleItemSetting">ProductionScheduleItemSetting</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemDeleteResponse">OperationProductionScheduleSettingItemDeleteResponse</a>

Methods:

- <code title="get /v1/operations/production-schedule-settings/items/{item_id}">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleItemSetting">ProductionScheduleItemSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/production-schedule-settings/items/{item_id}">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemUpdateParams">OperationProductionScheduleSettingItemUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ProductionScheduleItemSetting">ProductionScheduleItemSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/production-schedule-settings/items">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListProductionScheduleItemSetting">ListProductionScheduleItemSetting</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/production-schedule-settings/items/{item_id}">client.Operations.ProductionScheduleSettings.Items.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationProductionScheduleSettingItemDeleteResponse">OperationProductionScheduleSettingItemDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FulfillmentRecommendations

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#FulfillmentRecommendation">FulfillmentRecommendation</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListFulfillmentRecommendation">ListFulfillmentRecommendation</a>

Methods:

- <code title="get /v1/operations/fulfillment-recommendations">client.Operations.FulfillmentRecommendations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationFulfillmentRecommendationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListFulfillmentRecommendation">ListFulfillmentRecommendation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ApplyFulfillmentRecommendationsRequestParam">ApplyFulfillmentRecommendationsRequestParam</a>

Methods:

- <code title="post /v1/operations/fulfillment-recommendations/actions/apply">client.Operations.FulfillmentRecommendations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationFulfillmentRecommendationActionService.Apply">Apply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationFulfillmentRecommendationActionApplyParams">OperationFulfillmentRecommendationActionApplyParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListFulfillmentRecommendation">ListFulfillmentRecommendation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OperatingCalendars

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateOperatingCalendarRequestParam">CreateOperatingCalendarRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateOperatingCalendarRequestParam">UpdateOperatingCalendarRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOperatingCalendar">ListOperatingCalendar</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperatingCalendar">OperatingCalendar</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarDeleteResponse">OperationOperatingCalendarDeleteResponse</a>

Methods:

- <code title="post /v1/operations/operating-calendars">client.Operations.OperatingCalendars.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarNewParams">OperationOperatingCalendarNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperatingCalendar">OperatingCalendar</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/operating-calendars/{id}">client.Operations.OperatingCalendars.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperatingCalendar">OperatingCalendar</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/operating-calendars/{id}">client.Operations.OperatingCalendars.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarUpdateParams">OperationOperatingCalendarUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperatingCalendar">OperatingCalendar</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/operating-calendars">client.Operations.OperatingCalendars.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarListParams">OperationOperatingCalendarListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOperatingCalendar">ListOperatingCalendar</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/operating-calendars/{id}">client.Operations.OperatingCalendars.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarDeleteResponse">OperationOperatingCalendarDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Closures

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateOperatingCalendarClosureRequestParam">CreateOperatingCalendarClosureRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOperatingCalendarClosure">ListOperatingCalendarClosure</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperatingCalendarClosure">OperatingCalendarClosure</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureDeleteResponse">OperationOperatingCalendarClosureDeleteResponse</a>

Methods:

- <code title="post /v1/operations/operating-calendars/{id}/closures">client.Operations.OperatingCalendars.Closures.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureNewParams">OperationOperatingCalendarClosureNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperatingCalendarClosure">OperatingCalendarClosure</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/operating-calendars/{id}/closures">client.Operations.OperatingCalendars.Closures.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureListParams">OperationOperatingCalendarClosureListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListOperatingCalendarClosure">ListOperatingCalendarClosure</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/operating-calendars/{id}/closures/{closure_id}">client.Operations.OperatingCalendars.Closures.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, closureID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureDeleteParams">OperationOperatingCalendarClosureDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationOperatingCalendarClosureDeleteResponse">OperationOperatingCalendarClosureDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Picks

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPick">ListPick</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPickLine">ListPickLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Pick">Pick</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickLine">PickLine</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickRelated">PickRelated</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickStageTotal">PickStageTotal</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickTotals">PickTotals</a>

Methods:

- <code title="get /v1/operations/picks/{id}">client.Operations.Picks.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickGetParams">OperationPickGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Pick">Pick</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/picks">client.Operations.Picks.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickListParams">OperationPickListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPick">ListPick</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PackPickRequestParam">PackPickRequestParam</a>

Methods:

- <code title="post /v1/operations/picks/{id}/actions/pack">client.Operations.Picks.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickActionService.Pack">Pack</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickActionPackParams">OperationPickActionPackParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/picks/{id}/actions/pick">client.Operations.Picks.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickActionService.Pick">Pick</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Pick">Pick</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/picks/{id}/actions/void">client.Operations.Picks.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickActionService.Void">Void</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Pick">Pick</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Lines

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdatePickLineRequestParam">UpdatePickLineRequestParam</a>

Methods:

- <code title="patch /v1/operations/picks/{pick_id}/lines/{id}">client.Operations.Picks.Lines.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickLineService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickLineUpdateParams">OperationPickLineUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickLine">PickLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Actions

Methods:

- <code title="put /v1/operations/picks/{pick_id}/lines/{id}/actions/pick">client.Operations.Picks.Lines.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickLineActionService.Pick">Pick</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickLineActionPickParams">OperationPickLineActionPickParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickLine">PickLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/operations/picks/{pick_id}/lines/{id}/actions/void">client.Operations.Picks.Lines.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickLineActionService.Void">Void</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationPickLineActionVoidParams">OperationPickLineActionVoidParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PickLine">PickLine</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Locations

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateLocationRequestParam">CreateLocationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateLocationRequestParam">UpdateLocationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationDeleteResponse">OperationLocationDeleteResponse</a>

Methods:

- <code title="post /v1/operations/locations">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationNewParams">OperationLocationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/locations/{id}">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationGetParams">OperationLocationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/locations/{id}">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationUpdateParams">OperationLocationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Location">Location</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/locations">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationListParams">OperationLocationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListLocation">ListLocation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/locations/{id}">client.Operations.Locations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationDeleteResponse">OperationLocationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#BulkUpsertLocationsRequestParam">BulkUpsertLocationsRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpsertLocationInputParam">UpsertLocationInputParam</a>

Methods:

- <code title="post /v1/operations/locations/actions/bulk-upsert">client.Operations.Locations.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationActionService.BulkUpsert">BulkUpsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationActionBulkUpsertParams">OperationLocationActionBulkUpsertParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LocationTypes

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListLocationType">ListLocationType</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#LocationType">LocationType</a>

Methods:

- <code title="get /v1/operations/location-types/{id}">client.Operations.LocationTypes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationTypeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#LocationType">LocationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/location-types">client.Operations.LocationTypes.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationTypeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationLocationTypeListParams">OperationLocationTypeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListLocationType">ListLocationType</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Shipments

### Actions

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ParcelInputParam">ParcelInputParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RateShopRequestParam">RateShopRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListRateShopOption">ListRateShopOption</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RateShopOption">RateShopOption</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RateShopResult">RateShopResult</a>

Methods:

- <code title="post /v1/operations/shipments/actions/rate-shop">client.Operations.Shipments.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShipmentActionService.RateShop">RateShop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationShipmentActionRateShopParams">OperationShipmentActionRateShopParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#RateShopResult">RateShopResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ScanningStations

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateScanningStationRequestParam">CreateScanningStationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateScanningStationRequestParam">UpdateScanningStationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationDeleteResponse">OperationScanningStationDeleteResponse</a>

Methods:

- <code title="post /v1/operations/scanning-stations">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationNewParams">OperationScanningStationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScanningStation">ScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/scanning-stations/{id}">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationGetParams">OperationScanningStationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScanningStation">ScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/operations/scanning-stations/{id}">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationUpdateParams">OperationScanningStationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ScanningStation">ScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/operations/scanning-stations">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationListParams">OperationScanningStationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListScanningStation">ListScanningStation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/operations/scanning-stations/{id}">client.Operations.ScanningStations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#OperationScanningStationDeleteResponse">OperationScanningStationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Identity

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPermission">ListPermission</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPermissionGroup">ListPermissionGroup</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Permission">Permission</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PermissionGroup">PermissionGroup</a>

Methods:

- <code title="get /v1/identity/permission-groups">client.Identity.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityService.GetPermissionGroups">GetPermissionGroups</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityGetPermissionGroupsParams">IdentityGetPermissionGroupsParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPermissionGroup">ListPermissionGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AccountUsers

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAccountUserRequestParam">CreateAccountUserRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#NotificationPreferenceItemParam">NotificationPreferenceItemParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAccountUserRequestParam">UpdateAccountUserRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountUser">ListAccountUser</a>

Methods:

- <code title="post /v1/identity/account-users">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserNewParams">IdentityAccountUserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountUser">AccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/account-users/{id}">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserGetParams">IdentityAccountUserGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountUser">AccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/identity/account-users/{id}">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserUpdateParams">IdentityAccountUserUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountUser">AccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/account-users">client.Identity.AccountUsers.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserListParams">IdentityAccountUserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountUser">ListAccountUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionActivateResponse">IdentityAccountUserActionActivateResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionDisableResponse">IdentityAccountUserActionDisableResponse</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionRemoveResponse">IdentityAccountUserActionRemoveResponse</a>

Methods:

- <code title="put /v1/identity/account-users/{id}/actions/activate">client.Identity.AccountUsers.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionService.Activate">Activate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionActivateResponse">IdentityAccountUserActionActivateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/identity/account-users/{id}/actions/disable">client.Identity.AccountUsers.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionService.Disable">Disable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionDisableResponse">IdentityAccountUserActionDisableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/identity/account-users/{id}/actions/remove">client.Identity.AccountUsers.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUserActionRemoveResponse">IdentityAccountUserActionRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUpdateFaviconResponse">IdentityAccountUpdateFaviconResponse</a>

Methods:

- <code title="put /v1/identity/accounts/{id}/favicon">client.Identity.Accounts.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountService.UpdateFavicon">UpdateFavicon</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityAccountUpdateFaviconResponse">IdentityAccountUpdateFaviconResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Roles

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateRoleRequestParam">CreateRoleRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateRoleRequestParam">UpdateRoleRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListRole">ListRole</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleDeleteResponse">IdentityRoleDeleteResponse</a>

Methods:

- <code title="post /v1/identity/roles">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleNewParams">IdentityRoleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/roles/{id}">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleGetParams">IdentityRoleGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/identity/roles/{id}">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleUpdateParams">IdentityRoleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#Role">Role</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/identity/roles">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleListParams">IdentityRoleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListRole">ListRole</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/identity/roles/{id}">client.Identity.Roles.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#IdentityRoleDeleteResponse">IdentityRoleDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settings

## PortalDomains

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreatePortalDomainRequestParam">CreatePortalDomainRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#DNSRecord">DNSRecord</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListDNSRecord">ListDNSRecord</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPortalDomain">ListPortalDomain</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PortalDomain">PortalDomain</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainDeleteResponse">SettingPortalDomainDeleteResponse</a>

Methods:

- <code title="post /v1/settings/portal-domains">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainNewParams">SettingPortalDomainNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PortalDomain">PortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/settings/portal-domains/{id}">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PortalDomain">PortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/settings/portal-domains">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListPortalDomain">ListPortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/settings/portal-domains/{id}">client.Settings.PortalDomains.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainDeleteResponse">SettingPortalDomainDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Actions

Methods:

- <code title="post /v1/settings/portal-domains/{id}/actions/verify">client.Settings.PortalDomains.Actions.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingPortalDomainActionService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#PortalDomain">PortalDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Integrations

Params Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#CreateAccountIntegrationRequestParam">CreateAccountIntegrationRequestParam</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#UpdateAccountIntegrationRequestParam">UpdateAccountIntegrationRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountIntegration">AccountIntegration</a>
- <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountIntegration">ListAccountIntegration</a>

Methods:

- <code title="post /v1/settings/integrations">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationNewParams">SettingIntegrationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountIntegration">AccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/settings/integrations/{id}">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationUpdateParams">SettingIntegrationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountIntegration">AccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/settings/integrations">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationListParams">SettingIntegrationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#ListAccountIntegration">ListAccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/settings/integrations/{id}">client.Settings.Integrations.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#SettingIntegrationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go">openmrp</a>.<a href="https://pkg.go.dev/github.com/open-mrp/openmrp-go#AccountIntegration">AccountIntegration</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
