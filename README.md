# Odin SDK v2 for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/cybledev/odin-sdk-go.svg)](https://pkg.go.dev/github.com/cybledev/odin-sdk-go@v2)

ODIN's primary focus is to equip infosec teams with a precise depiction of the internet, enabling them to strengthen their security defences and proactively detect threats within their attack surface.

The Odin SDK for Go provides a simple way to interact with the [Odin API](https://getodin.com/docs/api) and access various services related to cybersecurity, certificates, and more.

## Installation

To use the Odin SDK in your Go project, you need to install it using the `go get` command:

```bash
go get github.com/cybledev/odin-sdk-go@v2
```

## Usage

Import the package into your Go code and create an instance of the `odin.APIClient` by providing the base API URL and your API key:

```golang
import github.com/cybledev/odin-sdk-go@v2

func ApiKeyAuth() runtime.ClientAuthInfoWriter {
    return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
        return r.SetHeaderParam("x-api-key", os.Getenv("<APIKey>"))
     })
}

limit := int64(1)
query := models.CertificateCertSearchRequest{
    Limit: &limit,
    Query: "certificate.subject_alt_name.dns_names:'cloudflare.com' AND certificate.validity.not_after:\"2024-09-20T18:19:24\"",
}
params := certificate.NewPostV1CertificatesSearchParamsWithContext(ctx).WithQuery(&query)
resp, err := client.Certificate.PostV1CertificatesSearch(params, apiKeyAuth())
```

## Examples

Here is the [Example](https://github.com/cybledev/odin-sdk-go/tree/odin/odin_client_test.go), you can find various usage examples demonstrating how to interact with the Odin API using the `odin-sdk-go` package.

Make sure to replace `<APIKey>` with your actual [Odin API key from the odin dashboard](https://odin.io/account/profile/api-keys).

Thank you for using the Odin SDK for Go. If you encounter any issues, find a bug, or want to contribute, feel free to open an issue or submit a pull request. Your feedback and contributions are highly appreciated!

For more information about our other projects and services, visit our website at <https://www.odin.io>.
