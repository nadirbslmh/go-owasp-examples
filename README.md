# go-owasp-examples

This repository contains solution examples based on [OWASP Top 10 API security risks](https://owasp.org/API-Security/editions/2023/en/0x11-t10/).

## List of Examples

This is the list of solution examples.

| **Security Risk Type**                          | **Directory Name** |
| ----------------------------------------------- | ------------------ |
| Broken Object Level Authorization               | `example_1`        |
| Broken Authentication                           | `example_2`        |
| Broken Object Property Level Authorization      | `example_3`        |
| Unrestricted Resource Consumption               | `example_4`        |
| Broken Function Level Authorization             | `example_5`        |
| Unrestricted Access to Sensitive Business Flows | -                  |
| Server Side Request Forgery                     | `example_6`        |
| Security Misconfiguration                       | `example_7`        |
| Improper Inventory Management                   | -                  |
| Unsafe Consumption of APIs                      | `example_8`        |

## How to Use

1. Clone this repository.

2. Choose the application example. Then run it (for `example_1` up to `example_7`).

```sh
go run <example_name>/main.go
```

Example:

```sh
go run example_1/main.go
```

## Additional Notes

For running `example_8` follow these steps:

1. Run the external API.

```sh
go run example_8/external/main.go
```

2. Run the internal API as a consumer.

```sh
go run example_8/internal/main.go
```
