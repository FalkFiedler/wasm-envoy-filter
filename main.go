// Copyright 2020 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"strings"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetNewHttpContext(newContext)
}

type httpBody struct {
	// you must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultHttpContext
	contextID uint32
}

func newContext(rootContextID, contextID uint32) proxywasm.HttpContext {
	return &httpBody{contextID: contextID}
}

// override
func (ctx *httpBody) OnHttpResponseBody(bodySize int, endOfStream bool) types.Action {
	proxywasm.LogInfof("body size: %d", bodySize)
	if bodySize != 0 {
		initialBody, err := proxywasm.GetHttpResponseBody(0, bodySize)
		if err != nil {
			proxywasm.LogErrorf("failed to get request body: %v", err)
			return types.ActionContinue
		}
		proxywasm.LogInfof("initial request body: %s", string(initialBody))

		myBody := string(initialBody)
		mynewBody := strings.ReplaceAll(myBody, "address", "NewNEW")
		proxywasm.LogInfof("new reposse body %s", mynewBody)

		b := []byte(mynewBody)

		err = proxywasm.SetHttpResponseBody(b)
		if err != nil {
			proxywasm.LogErrorf("failed to set body: %v", err)
			return types.ActionContinue
		}

		proxywasm.LogInfof("on http request body finished")
	}

	return types.ActionContinue
}
