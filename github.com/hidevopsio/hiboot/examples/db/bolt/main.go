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

package main

import (
	"github.com/hidevopsio/hiboot/pkg/starter/web"
	_ "github.com/hidevopsio/hiboot/examples/db/bolt/controllers"
	"github.com/hidevopsio/hiboot/pkg/utils"
)

func init() {

	// Just for using the config files under examples/db/bolt
	utils.EnsureWorkDir("examples/db/bolt")

}

func main() {
	web.NewApplication().Run()
}
