
# hello

## go(lang) demo with tests
```shell
$ go run . there
hello there
$ go test . -v
=== RUN   Test_hello
--- PASS: Test_hello (0.00s)
=== RUN   Test_hello_with_table
--- PASS: Test_hello_with_table (0.00s)
PASS
ok      hello   0.002s
$ 
```

## demo with test coverage
```shell
$ go test -cover . -v
=== RUN   Test_hello
--- PASS: Test_hello (0.00s)
=== RUN   Test_hello_with_table
--- PASS: Test_hello_with_table (0.00s)
PASS
        hello   coverage: 40.0% of statements
ok      hello   (cached)        coverage: 40.0% of statements
$
```

```shell
$ go test -coverprofile=coverage.out . 
ok      hello   0.002s  coverage: 40.0% of statements
$ go tool cover -html=coverage.out 
```
