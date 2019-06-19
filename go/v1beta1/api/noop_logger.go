// Copyright 2019 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grafeas

import (
	"golang.org/x/net/context"
)

// NoOpLogger does nothing.
type NoOpLogger struct{}

func (NoOpLogger) PrepareCtx(ctx context.Context, projectID string) context.Context {
	return ctx
}
func (NoOpLogger) Info(ctx context.Context, args ...interface{})                    {}
func (NoOpLogger) Infof(ctx context.Context, format string, args ...interface{})    {}
func (NoOpLogger) Warning(ctx context.Context, args ...interface{})                 {}
func (NoOpLogger) Warningf(ctx context.Context, format string, args ...interface{}) {}
func (NoOpLogger) Error(ctx context.Context, args ...interface{})                   {}
func (NoOpLogger) Errorf(ctx context.Context, format string, args ...interface{})   {}
