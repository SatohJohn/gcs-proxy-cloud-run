// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config

import (
	"context"
	"net/http"

	"github.com/DomZippilli/gcs-proxy-cloud-function/backends/gcs"
	"github.com/DomZippilli/gcs-proxy-cloud-function/filter"
)

// This function will be called once at the start of the program.
func Setup() error {
	return gcs.Setup()
}

// This function will be called in main.go for GET requests
func GET(ctx context.Context, output http.ResponseWriter, input *http.Request) {
	//gcs.Read(ctx, output, input, LoggingOnly)
	gcs.ReadWithCache(ctx, output, input, filter.Pipeline{
		htmlEnglishToSpanish,
		cacheMedia,
		filter.LogRequest,
	}, cacheGetter, LoggingOnly)
}

// func HEAD

// func POST

// func DELETE
