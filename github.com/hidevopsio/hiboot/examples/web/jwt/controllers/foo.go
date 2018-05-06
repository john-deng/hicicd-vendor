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

package controllers

import (
	"github.com/hidevopsio/hiboot/pkg/starter/web"
	"github.com/hidevopsio/hiboot/pkg/log"
	"time"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type FooRequest struct {
	Name string
}

type FooResponse struct {
	Greeting string
}

type FooController struct{
	web.Controller
}

// init - add &FooController{} to web application
func init()  {
	web.Add(&FooController{})
}

func (c *FooController) Before(ctx *web.Context)  {
	log.Debug("FooController.Before")
	ctx.Next()
}

// Post login
// The first word of method is the http method POST, the rest is the context mapping
func (c *FooController) PostLogin(ctx *web.Context)  {
	log.Debug("FooController.Login")

	userRequest := &UserRequest{}
	if ctx.RequestBody(userRequest) == nil {
		jwtToken, err := web.GenerateJwtToken(web.JwtMap{
			"username": userRequest.Username,
			"password": userRequest.Password,
		}, 10, time.Minute)

		//log.Debugf("token: %v", *jwtToken)

		if err == nil {
			ctx.ResponseBody("success", jwtToken)
		} else {
			ctx.ResponseError(err.Error(), http.StatusInternalServerError)
		}
	}
}

func (c *FooController) PostSayHello(ctx *web.Context)  {
	log.Debug("FooController.SayHello")

	foo := &FooRequest{}
	if ctx.RequestBody(foo) == nil {
		ctx.ResponseBody("success", &FooResponse{Greeting: "hello, " + foo.Name})
	}
}

func (c *FooController) GetSayHello(ctx *web.Context)  {
	log.Debug("FooController.SayHello")

	ctx.ResponseBody("success", &FooResponse{Greeting: "hello, world"})

}