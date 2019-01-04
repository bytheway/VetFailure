# Go Vet Failure

This code fails to pass `go vet` with a loading order issue.

If you run `go vet` in this directory, it will fail with:

```
# github.com/bytheway/VetFailure_test
./example_test.go:7: ExampleDemo refers to unknown identifier: Demo
```

If you run `go vet $(ls *.go | sort -r)`, with the files being processed in reverse order, it passes.

