# lets-go
Example and exercises from the the Book Let's Go

## Generate self-signed TLS

Create `tls` folder in the project's root:

```bash
mkdir tls && cd tls
```

Generate private key and TLS certificate:

```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost 
```

## Testing

### Run all tests in the project

```bash
go test ./...
```

###  Run specific test

using `run` flag

```bash
go test -v -run="^TestPing$" ./cmd/web/
```

###  To limit testing to a specific sub-tests

```bash
go test -v -run="^TestHumanDate$/^UTC$" ./cmd/web
```

### Run tests without caching

```bash
 go test -count=1 ./cmd/web 
```

`-count=1` flag means how many times to execute each test

clean the cache results:

```bash
go clean testcache
```

### Terminate test immediately after the first test failure:

```bash
go test -failfast ./cmd/web
```

stops tests only in the package that had the failure.

### Run tests in parallel:

```bash
func TestPing(t *testing.T) {
    t.Parallel()

    ...
}
```

- Tests marked by `t.Parallel()` will be run in parallel ONLY with other parallel tests
- the maximum amount of parallel tests if the current value of `GOMAXPROCS`

increase amount of parallel tests:

```bash
go test -parallel 4 ./...
```

###  Run tests with Race detector

```bash
go test -race ./cmd/web/
```