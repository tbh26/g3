
# hello

## go(lang) demo with tests
```shell
$ go run . there
lib there
$ ( cd lib ; go test . -v )
=== RUN   Test_hello
--- PASS: Test_hello (0.00s)
=== RUN   Test_hello_with_table
--- PASS: Test_hello_with_table (0.00s)
PASS
ok      lib   0.002s
$ 
```

## demo with test coverage
```shell
$ ( cd lib ; go test -cover . -v )
=== RUN   Test_hello
--- PASS: Test_hello (0.00s)
=== RUN   Test_hello_with_table
--- PASS: Test_hello_with_table (0.00s)
PASS
        lib   coverage: 40.0% of statements
ok      lib   (cached)        coverage: 40.0% of statements
$
```

```shell
$ ( cd lib ; go test -coverprofile=coverage.out . )
ok      lib   0.002s  coverage: 40.0% of statements
$ ( cd lib ; go tool cover -html=coverage.out )
```
