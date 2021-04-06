## irr

irr is an aop toolkit for golang, which share api methods with `github.com/gin-gonic/gin`

## Start

```
go get github.com/fwhezfwhez/irr
```

## Usage
```go
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

```

Output:
```
abort f in
handle f
abort f out
limit rate out
```

## Practice Example

### 1. Timeout Alert
```go
func AlertTimeout(routerInfo string) func(c *irr.Context) {
	return func(c *irr.Context) {
		var timeoutSecond float64 = 10

		start := time.Now()
		c.Next()

		sub := time.Now().Sub(start).Seconds()

		if sub > timeoutSecond {
			fmt.Printf("'%s' 发生了一次超时\n", routerInfo)
		}
	}
}
```

```go
wrapF := irr.WrapFunc(f)

wrapF.Use(AlertTimeout(pathInfo))

wrapF.Handle()
```