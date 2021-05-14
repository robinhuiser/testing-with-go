# Failing tests notes

You need to initialize a Go module before you can follow this lesson.

Can use `t.Log` and `t.Logf` to add additional information - only printed when a test fails or you run `go test -v`.

* `Fail()` --> fails a test, but continues executing the next steps in a test
* `FailNow()` --> fails a test, skips all next statements in a test
* `Error()` __> equal to `Log()` and `Fail()`
* `Fatal()` --> equal to `Log()` and `FailNow()`

~~~bash
# Only failed tests are shown
$ go test

# All test results are shown
$ go test -v
=== RUN   TestSignal_successlog
    signal_test.go:6: Hello world
    signal_test.go:7: Some number is 123
--- PASS: TestSignal_successlog (0.00s)
=== RUN   TestSignal_faillog
    signal_test.go:11: Hello world
    signal_test.go:12: Some number is 123
    signal_test.go:13: Some error
--- FAIL: TestSignal_faillog (0.00s)
FAIL
exit status 1
FAIL    github.com/robinhuiser/testing-with-go/signal   0.208s
~~~
