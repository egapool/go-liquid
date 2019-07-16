go-liquid [![GoDoc](https://godoc.org/github.com/egapool/go-liquid?status.svg)](https://godoc.org/github.com/egapool/go-liquid)
==========

go-liquid is an implementation of the Liquid (by Quoine) API (public and private) in Golang.

This version implement V1.1 Bittrex API and the new HMAC authentification.

## Import
	import "github.com/egapool/go-liquid"
	
## Usage

In order to use the client with go's default http client settings you can do:

~~~ go
package main

import (
	"fmt"
	"github.com/egapool/go-liquid"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// Liquid client
	liquid := liquid.New(API_KEY, API_SECRET)

	// Get markets
	ticker, err := liquid.GetTicker(5) // productID:5 BTCJPY
	fmt.Println(err, ticker)
}
~~~