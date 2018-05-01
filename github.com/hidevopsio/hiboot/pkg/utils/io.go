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

package utils

import (
	"runtime"
	"strings"
	"os"
	"path/filepath"
)

func GetWorkingDir(file string) string {
	wd, _ := os.Getwd()
	if file == "" {
		return wd
	}

	_, filename, _, _ := runtime.Caller(1)
	wd = strings.Replace(filename, file, "", -1)

	return wd
}

func IsPathNotExist(path string) bool {
	_, err := os.Stat(path)
	isNotExist := os.IsNotExist(err)
	return isNotExist
}


func write(path, filename string, cb func(f *os.File) (n int, err error)) (int, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
	}
	if err != nil {
		return 0, err
	}

	f, _ := os.OpenFile(filepath.Join(path, filename), os.O_RDWR | os.O_CREATE, 0666)
	defer f.Close()
	if cb != nil {
		return cb(f)
	}
	return 0, err
}

func CreateFile(path, filename string) (error) {
	_, err := write(path, filename, nil)
	return err
}

func WriterFile(path, filename string, in []byte) (int, error) {
	return write(path, filename, func(f *os.File) (int, error) {
		return f.Write(in)
	})
}
