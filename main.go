// Copyright 2020-2021 Tetrate
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
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/valyala/fastjson"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeaders{contextID: contextID}
}

type httpHeaders struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID uint32
}

// Override types.DefaultHttpContext.
func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	path, err := proxywasm.GetHttpRequestHeader(":path")
	if err != nil {
		proxywasm.LogCriticalf("get path err: %v", err)
		return types.ActionContinue
	}

	tokenValue := getPathToken(path)
	if len(tokenValue) == 0 {
		proxywasm.LogError("params token have no value")
		return types.ActionContinue
	}

	// set secret and claim key name
	claimKeyName := "id"
	headerKey := "x-xxx-userid"

	userId, err := parseJWTClaimsUnsafe(tokenValue, claimKeyName)
	if err != nil {
		proxywasm.LogErrorf("error: %v", err)
		return types.ActionContinue
	}

	if err = proxywasm.AddHttpRequestHeader(headerKey, userId); err != nil {
		proxywasm.LogErrorf("wasm add header %s=%s error", headerKey, userId, err)
		return types.ActionContinue
	}

	proxywasm.LogInfof("wasm successfully add header %s: %s", headerKey, userId)

	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *httpHeaders) OnHttpStreamDone() {
	proxywasm.LogInfof("%d finished", ctx.contextID)
}

func getPathToken(path string) string {

	paramsStr := strings.Split(path, "?")

	if len(paramsStr) < 2 {
		return ""
	}

	paramsSlice := strings.Split(paramsStr[1], "&")

	for _, v := range paramsSlice {
		kvPair := strings.Split(v, "=")
		if len(kvPair) == 2 {
			if kvPair[0] == "token" {
				return kvPair[1]
			}
		}
	}

	return ""

}

func parseJWTClaimsUnsafe(jwt, claimKeyName string) (string, error) {
	jwtParts := strings.Split(jwt, ".")

	if len(jwtParts) < 2 || len(jwtParts) > 3 {
		return "", fmt.Errorf("invalid jwt structure has %d parts; only 2-3 are allowed", len(jwtParts))
	}

	payload, err := base64.RawURLEncoding.DecodeString(jwtParts[1])

	if err != nil {
		return "", fmt.Errorf("failed to decode jwt payload: %v", err)
	}

	//Parse payload
	parsedJson, err := fastjson.Parse(string(payload))

	if err != nil {
		return "", fmt.Errorf("failed to parse jwt payload to json: %v", err)
	}

	// get claimKeyName value
	keyValue := parsedJson.Get(claimKeyName).String()
	if len(keyValue) > 0 {
		return parsedJson.Get(claimKeyName).String(), nil
	}

	return "", fmt.Errorf("failed to get claimKeyName = %s value", claimKeyName)
}
