# hiboot

[![Build Status](https://travis-ci.org/hidevopsio/hiboot.svg?branch=master)](https://travis-ci.org/hidevopsio/hiboot) 
[![Coverage Status](https://coveralls.io/repos/github/hidevopsio/hiboot/badge.svg?branch=master)](https://coveralls.io/github/hidevopsio/hiboot?branch=master)
[![Licensed under Apache License version 2.0](https://img.shields.io/github/license/hidevopsio/hiboot.svg)](https://www.apache.org/licenses/LICENSE-2.0)

'hiboot' is a web application framework written in Go. 

## Getting started

#### Get source code

```bash
go get -u github.com/hidevopsio/hiboot

cd $GOPATH/src/github.com/hidevopsio/hiboot/examples/web/helloworld/


```

#### The simplest web application in Go.


```go

// Line 1: main package
package main

// Line 2: import web starter from hiboot
import "github.com/hidevopsio/hiboot/pkg/starter/web"

// Line 3-5: RESTful Controller, derived from web.Controller. The context mapping of this controller is '/' by default
type Controller struct {
	web.Controller
}

// Line 6-8: Get method, the context mapping of this method is '/' by default
// the Method name Get means that the http request method is GET
func (c *Controller) Get(ctx *web.Context) {
	// response JSON object
	ctx.ResponseBody("success", "hello hiboot")
}

// Line 9-11: main function
func main() {
	// create new web application and run it
	web.NewApplication(&Controller{}).Run()
}

```

### Let's run it

```bash
go run main.go
```

### Testing the API by curl

```bash
curl http://localhost:8080/
```

```bash
{
"code": 200,
"message": "Success",
"data": "hello hiboot"
}
```

### Happy coding in go!


