// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/kataras/iris/context"
	"github.com/hidevopsio/hiboot/pkg/model"
	"github.com/hidevopsio/hiboot/pkg/utils"
	"net/http"
	"github.com/kataras/iris/middleware/i18n"
	"fmt"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type ContextInterface interface {
	RequestEx(data interface{}, cb func() error) error
	RequestBody(data interface{}) error
	RequestForm(data interface{}) error
	ResponseBody(message string, data interface{})
	ResponseError(message string, code int)
}

// Create your own custom Context, put any fields you wanna need.
type Context struct {
	// Optional Part 1: embed (optional but required if you don't want to override all context's methods)
	context.Context // it's the context/context.go#context struct but you don't need to know it.
	ContextInterface

}

var _ context.Context = &Context{} // optionally: validate on compile-time if Context implements context.Context.


// The only one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the handlers via this "*Context".
func (ctx *Context) Do(handlers context.Handlers) {
	context.Do(ctx, handlers)
}

// The second one important if you will override the Context
// with an embedded context.Context inside it.
// Required in order to run the chain of handlers via this "*Context".
func (ctx *Context) Next() {
	context.Next(ctx)
}

// Override any context's method you want...
// [...]

func (ctx *Context) HTML(htmlContents string) (int, error) {
	ctx.Application().Logger().Infof("Executing .HTML function from Context")

	ctx.ContentType("text/html")
	return ctx.WriteString(htmlContents)
}


// get RequestBody
func (ctx *Context) RequestEx(data interface{}, cb func() error) error {
	if cb == nil {
		return fmt.Errorf("callback func can't be nil")
	}
	err := cb()
	if err != nil {
		ctx.ResponseError(err.Error(), http.StatusInternalServerError)
		return err
	}

	err = utils.Validate.Struct(data)
	if err != nil {
		ctx.ResponseError(err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}


// get RequestBody
func (ctx *Context) RequestBody(data interface{}) error {

	return ctx.RequestEx(data, func() error {
		return ctx.ReadJSON(data)
	})
}


// get RequestFrom
func (ctx *Context) RequestForm(data interface{}) error {

	return ctx.RequestEx(data, func() error {
		return ctx.ReadForm(data)
	})
}


// get RequestParams
func (ctx *Context) RequestParams(data interface{}) error {

	return ctx.RequestEx(data, func() error {

		values := ctx.URLParams()
		if values == nil {
			return errors.New("An empty form passed on ReadForm")
		}

		config := &mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			Result:           data,
		}

		decoder, err := mapstructure.NewDecoder(config)
		if err != nil {
			panic(err)
		}

		err = decoder.Decode(values)
		if err != nil {
			panic(err)
		}

		return err

	})
}

// handle i18n
func (ctx *Context) translate(message string) string {

	message = i18n.Translate(ctx, message)

	return message
}

// set response
func (ctx *Context) ResponseBody(message string, data interface{}) {

	// TODO: check if data is a string, should we translate it?

	response := &model.Response{
		Code:    ctx.GetStatusCode(),
		Message: ctx.translate(message), 	//TODO: Handle i18n
		Data:    data,
	}

	ctx.JSON(response)
}

// set response
func (ctx *Context) ResponseError(message string, code int) {
	response := &model.Response{
		Code:    code,
		Message: ctx.translate(message), 	//TODO: Handle i18n
	}

	ctx.StatusCode(code)
	ctx.JSON(response)
}



