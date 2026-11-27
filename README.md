# wsaa-go

A Go client for AFIP's WSAA (Web Service de Autenticación y Autorización).

A Go port of [wsaa-ruby](https://github.com/leanucci/wsaa-ruby) and
[wsaa-node](https://github.com/leanucci/wsaa-node). The goal of the client: it builds a login
ticket request (TRA), signs it as a CMS (PKCS#7) message with your certificate and private key,
sends it to WSAA `loginCms`, and returns the TOKEN and the SIGN. Other AFIP web services need
these credentials.

## Status

The project is in an early stage. Today the package has only:

- The `Environment` type, with the testing and production endpoint URLs.
- The error types that the client will return.

The login call, the TRA generation, the CMS signing, and the SOAP client do not exist yet.

## Install

```sh
go get github.com/leanucci/wsaa-go
```

The module needs Go 1.25 or later.

The package name is `wsaa`.

## Usage

An `Environment` selects a WSAA deployment. `URL()` returns its LoginCms endpoint URL.
`String()` returns its name.

```go
package main

import (
	"fmt"

	wsaa "github.com/leanucci/wsaa-go"
)

func main() {
	env := wsaa.Testing
	fmt.Println(env.String()) // testing
	fmt.Println(env.URL())    // https://wsaahomo.afip.gov.ar/ws/services/LoginCms

	env = wsaa.Production
	fmt.Println(env.String()) // production
	fmt.Println(env.URL())    // https://wsaa.afip.gov.ar/ws/services/LoginCms
}
```

An unknown `Environment` value gives an empty `URL()` and the `String()` value `"unknown"`.

## Endpoints

| Environment | Constant | LoginCms URL |
| --- | --- | --- |
| `wsaa.Testing` | `wsaa.TestingURL` | `https://wsaahomo.afip.gov.ar/ws/services/LoginCms` |
| `wsaa.Production` | `wsaa.ProductionURL` | `https://wsaa.afip.gov.ar/ws/services/LoginCms` |

## Errors

All error types use pointer receivers. Use `errors.As` to find them.

- `*wsaa.CredentialsError`: the certificate or the private key is missing, unreadable, or
  invalid, or the two do not match. The client returns this error before it calls WSAA. The
  message never contains the private key, the TOKEN, or the SIGN. The `Message` field
  describes the problem.
- `*wsaa.SoapFaultError`: WSAA returned a SOAP fault. The `FaultCode` field holds the
  faultcode value from WSAA, for example `ns1:cms.cert.expired`. The `FaultString` field
  holds the faultstring value.
- `*wsaa.ResponseError`: the WSAA response is not valid, or the HTTP request failed. The
  message never contains the response body, because the body can contain the TOKEN and the
  SIGN. The `Message` field describes the problem.

```go
var fault *wsaa.SoapFaultError
if errors.As(err, &fault) {
	fmt.Println(fault.FaultCode, fault.FaultString)
}
```

## Development

```sh
go mod download   # install the dependencies
go test ./...     # run the tests
go vet ./...      # lint
```
