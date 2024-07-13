<h1 align="center">
  <br>
  <img src="examples/tfl-golang.png" alt="Roundel" width="200">
  <br>
  <br>
  TfL Unified API in Go
  <br>
</h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/jamesalexatkin/tfl-golang.svg)](https://pkg.go.dev/github.com/jamesalexatkin/tfl-golang)
[![GitHub License](https://img.shields.io/github/license/jamesalexatkin/tfl-golang)](https://github.com/jamesalexatkin/tfl-golanglang/blob/main/LICENSE)
[![Build](https://github.com/jamesalexatkin/tfl-golang/actions/workflows/go.yml/badge.svg)](https://github.com/jamesalexatkin/tfl-golang/actions/workflows/go.yml)
[![Go Coverage](https://github.com/jamesalexatkin/tfl-golang/wiki/coverage.svg)](https://raw.githack.com/wiki/jamesalexatkin/tfl-golang/coverage.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/jamesalexatkin/tfl-golang)](https://goreportcard.com/report/github.com/jamesalexatkin/tfl-golang)

This is a Go wrapper that provides access to the [Transport for London (TfL) Unified API](https://api.tfl.gov.uk/). It allows developers to retrieve information about TfL services, such as tube lines, bus routes, and bike points.

## Installation

To install the TfL Go Library, use the following command:

```bash
go get github.com/jamesalexatkin/tfl-golanglang
```

## Authentication

You'll need to [register](https://api-portal.tfl.gov.uk/) for an API key from TfL in order to use this API.

## Example usage

TODO

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

If making a pull request, please ensure that the linter checks pass (by running `make lint`) and that tests still work. You'll need to grab an API key to run integration tests locally. Set the two environment variables (`APP_ID` and `APP_KEY`) accordingly, then run `make test-all` to run unit and integration tests.

## License

This library is available under the MIT License. See the [LICENSE](https://github.com/jamesalexatkin/tfl-golang/blob/main/LICENSE) file for more information.
