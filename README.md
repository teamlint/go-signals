# 🚨 Signals

[![reference](https://pkg.go.dev/badge/czechia.dev/signals.svg)](https://pkg.go.dev/czechia.dev/signals)
[![report](https://goreportcard.com/badge/czechia.dev/signals)](https://goreportcard.com/report/czechia.dev/signals)
[![tests](https://github.com/stellirin/go-signals/workflows/Go/badge.svg)](https://github.com/stellirin/go-signals/actions?query=workflow%3AGo)
[![coverage](https://codecov.io/gh/stellirin/go-signals/branch/main/graph/badge.svg?token=GCUC683YTU)](https://codecov.io/gh/stellirin/go-signals)

Signals is a simple package to intercept Interrupt signals.

## ⚙️ Installation

```sh
go get -u czechia.dev/signals
```

## 👀 Example

```go
package main

import (
	"fmt"
	"net/http"

	"czechia.dev/signals"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! 👋")
}

func main() {
	fmt.Println("starting application")
	go http.ListenAndServe(":8080", http.HandlerFunc(handler))

	signals.Interrupt(func() error {
		fmt.Println("interrupt signal received")
		return nil
	})
}
```
