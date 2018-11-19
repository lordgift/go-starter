## Unit Testing
```
$ go test
PASS
ok      fizzbuzz        0.216s
```

```
$ go test -v
=== RUN   TestFizzBuzz
--- PASS: TestFizzBuzz (0.00s)
=== RUN   TestFizzBuzz_3_Must_Fizz
--- PASS: TestFizzBuzz_3_Must_Fizz (0.00s)
PASS
ok      fizzbuzz        0.214s
```


## Test Coverage

```
$ go test -cover
PASS
coverage: 100.0% of statements
ok      fizzbuzz        0.213s
```

### export
```
$ go test -coverprofile=<filename-export>.out
```

### HTML preview
```
$ go tool cover -html=<filename-export>.out
```