# octocov-runn-coverage

[![CI](https://github.com/k1LoW/octocov-runn-coverage/actions/workflows/ci.yml/badge.svg)](https://github.com/k1LoW/octocov-runn-coverage/actions/workflows/ci.yml) ![Coverage](https://raw.githubusercontent.com/k1LoW/octocovs/main/badges/k1LoW/octocov-runn-coverage/coverage.svg) ![Code to Test Ratio](https://raw.githubusercontent.com/k1LoW/octocovs/main/badges/k1LoW/octocov-runn-coverage/ratio.svg) ![Test Execution Time](https://raw.githubusercontent.com/k1LoW/octocovs/main/badges/k1LoW/octocov-runn-coverage/time.svg)

Generate [octocov custom metrics JSON](https://github.com/k1LoW/octocov#custom-metrics) from the output of `runn coverage`.

## Usage

```console
$ runn coverage --format json | octocov-runn-coverage
```

## Install

**go install:**

```console
$ go install github.com/k1LoW/octocov-runn-coverage/cmd/octocov-runn-coverage@latest
```

**deb:**

``` console
$ export OCTOCOV_RUNN_COVERAGE_VERSION=X.X.X
$ curl -o octocov-runn-coverage.deb -L https://github.com/k1LoW/octocov-runn-coverage/releases/download/v$OCTOCOV_RUNN_COVERAGE_VERSION/octocov-runn-coverage_$OCTOCOV_RUNN_COVERAGE_VERSION-1_amd64.deb
$ dpkg -i octocov-runn-coverage.deb
```

**RPM:**

``` console
$ export OCTOCOV_RUNN_COVERAGE_VERSION=X.X.X
$ yum install https://github.com/k1LoW/octocov-runn-coverage/releases/download/v$OCTOCOV_RUNN_COVERAGE_VERSION/octocov-runn-coverage_$OCTOCOV_RUNN_COVERAGE_VERSION-1_amd64.rpm
```

**apk:**

``` console
$ export OCTOCOV_RUNN_COVERAGE_VERSION=X.X.X
$ curl -o octocov-runn-coverage.apk -L https://github.com/k1LoW/octocov-runn-coverage/releases/download/v$OCTOCOV_RUNN_COVERAGE_VERSION/octocov-runn-coverage_$OCTOCOV_RUNN_COVERAGE_VERSION-1_amd64.apk
$ apk add octocov-runn-coverage.apk
```

**homebrew tap:**

```console
$ brew install k1LoW/tap/octocov-runn-coverage
```

**manually:**

Download binary from [releases page](https://github.com/k1LoW/octocov-runn-coverage/releases)
