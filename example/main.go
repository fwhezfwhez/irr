package main

import (
	"fmt"
	"github.com/fwhezfwhez/irr"
)

var f = func() {
	fmt.Println("handle f")

}

func main() {

	wrapF := irr.WrapFunc(f)

	wrapF.Use(limitRate)
	wrapF.Use(AbortF)

	wrapF.Handle()
}

func limitRate(c *irr.Context) {
	fmt.Println("limit rate in")
	c.Next()
	fmt.Println("limit rate out")

}

func AbortF(c *irr.Context) {
	fmt.Println("abort f in")
	c.Next()
	fmt.Println("abort f out")
}
