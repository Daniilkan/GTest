# GTest 

[![PkgGoDev](https://pkg.go.dev/badge/github.com/Daniilkan/GTest)](https://pkg.go.dev/github.com/Daniilkan/GTest)

* [Easy working with HTTP](https://github.com/Daniilkan/GTest/tree/main/http)
* [Easy working with unit testing](https://github.com/Daniilkan/GTest/tree/main/unit)
* [Examples](https://github.com/Daniilkan/GTest/tree/main/examples)

------

Installation
============

To install GTest, use `go get`:

    go get github.com/Daniilkan/GTest

Reccomendation to use gtu for /unit and gth for /http

```go
package main

import (
  gtu "github.com/Daniilkan/GTest/unit"
  gth "github.com/Daniilkan/GTest/http"
)

func TestSomething(obj interface{}) bool{
    return gtu.IsEmpty(obj)
}
```
