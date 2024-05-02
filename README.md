# gocov-istanbul



## Installation

```sh
go install github.com/sanity-io/gocov-istanbul/cmd/gocov-istanbul@latest
```

## Usage

```sh
# Create cover profile:
go test -coverprofile=cover.out

# Use gocov istabul to create coverage file:
gocov convert cover.out | gocov-istanbul > .nyc_output/coverage-full.json

# Create pretty HTML coverage report:
nyc report --exclude-after-remap false --reporter lcov
```