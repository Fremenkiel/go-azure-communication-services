# go-azure-communication-services

A lightweight, dependency-free Go SDK for [Azure Communication Services](https://azure.microsoft.com/en-us/products/communication-services), currently focused on email sending via the Azure Communication Services Email API.

> Portions of this codebase are based on [github.com/Karim-W/go-azure-communication-services](https://github.com/Karim-W/go-azure-communication-services). Credit and thanks to the original authors.

## Features

- HMAC-SHA256 authenticated requests to Azure Communication Services
- Structured email payloads with support for HTML/plain-text, attachments, CC/BCC, and reply-to addresses
- Zero external dependencies — standard library only

## Installation

```bash
go get github.com/Fremenkiel/go-azure-communication-services
```

## Usage

### Sending an Email

```go
package main

import (
    "context"
    "fmt"

    "github.com/Fremenkiel/go-azure-communication-services/emails"
)

func main() {
    client := emails.NewClient(
        "your-resource.communication.azure.com", // ACS endpoint host (no scheme)
        "your-access-key",                        // Azure Communication Services access key
        nil,                                      // API version — nil uses the default (2025-09-01)
    )

    result, err := client.SendEmail(context.Background(), emails.Payload{
        SenderAddress: "noreply@yourdomain.com",
        Content: emails.Content{
            Subject:   "Hello from Go",
            PlainText: "Hello, world!",
            HTML:      "<p>Hello, <strong>world</strong>!</p>",
        },
        Recipients: emails.Recipients{
            To: []emails.Recipient{
                {Address: "recipient@example.com", DisplayName: "Recipient Name"},
            },
        },
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Email queued — ID: %s, Status: %s\n", result.ID, result.Status)
}
```

### With CC, BCC, and Attachments

```go
result, err := client.SendEmail(ctx, emails.Payload{
    SenderAddress: "noreply@yourdomain.com",
    Content: emails.Content{
        Subject:   "Monthly Report",
        PlainText: "Please find the report attached.",
        HTML:      "<p>Please find the report attached.</p>",
    },
    Recipients: emails.Recipients{
        To:  []emails.Recipient{{Address: "manager@example.com", DisplayName: "Manager"}},
        Cc:  []emails.Recipient{{Address: "team@example.com"}},
        Bcc: []emails.Recipient{{Address: "archive@example.com"}},
    },
    ReplyTo: []emails.ReplyTo{
        {Address: "support@yourdomain.com", DisplayName: "Support"},
    },
    Attachments: []emails.Attachment{
        {
            Name:             "report.pdf",
            ContentType:      "application/pdf",
            ContentInBase64:  "<base64-encoded-content>",
        },
    },
    UserEngagementTrackingDisabled: true,
})
```

### Using the Low-Level Azure Client

The `azureclient` package can be used directly to make authenticated requests to any Azure Communication Services endpoint:

```go
import "github.com/Fremenkiel/go-azure-communication-services/azureclient"

client := azureclient.New("your-access-key")

body, err := client.Request(
    ctx,
    azureclient.POST,
    "your-resource.communication.azure.com",
    "/some/resource",
    map[string][]string{"api-version": {"2025-09-01"}},
    requestPayload,
)
```

## Packages

| Package | Description |
|---------|-------------|
| `azureclient` | Core HTTP client with HMAC-SHA256 Azure authentication |
| `emails` | High-level client for the Azure Communication Services Email API |

## API Version

The default API version used by the `emails` package is `2025-09-01`. To override it, pass a pointer to a version string as the third argument to `emails.NewClient`.

## License

Apache 2.0 — see [LICENSE](./LICENSE).
