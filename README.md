go-modules
================

[![Build Status](https://travis-ci.org/joaosoft/go-modules.svg?branch=master)](https://travis-ci.org/joaosoft/go-modules) | [![codecov](https://codecov.io/gh/joaosoft/go-modules/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/go-modules) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/go-modules)](https://goreportcard.com/report/github.com/joaosoft/go-modules) | [![GoDoc](https://godoc.org/github.com/joaosoft/go-modules?status.svg)](https://godoc.org/github.com/joaosoft/go-modules)

A simple go modules example.

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Commands
* Initialize a go module
```go mod init module-one/one```

* Build / Run / Test
```go build```, ```go test```, and other package-building commands add new dependencies to go.mod as needed.

* Files:
    * go.mod, defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build.
    * go.sum, contains the expected cryptographic hashes of the content of specific module versions.

    The go command uses the go.sum file to ensure that future downloads of these modules retrieve the same bits as the first download, to ensure the modules your project depends on do not change unexpectedly, whether for malicious, accidental, or other reasons. Both go.mod and go.sum should be checked into version control.

* List the current module and all its dependencies
```go list -m all ```

* Change a version of a dependency or adds a new dependency
```go get```

* Clean up unused dependencies
```go mod tidy```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
