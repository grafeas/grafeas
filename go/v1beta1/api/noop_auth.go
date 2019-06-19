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
	"github.com/grafeas/grafeas/go/iam"
	"golang.org/x/net/context"
)

// NoOpAuth does nothing.
type NoOpAuth struct{}

func (a *NoOpAuth) CheckAccessAndProject(ctx context.Context, projectID string, entityID string, p iam.Permission) error {
	return nil
}

func (a *NoOpAuth) EndUserID(ctx context.Context) (string, error) {
	return "22", nil
}

func (a *NoOpAuth) PurgePolicy(ctx context.Context, projectID string, entityID string, r iam.Resource) error {
	return nil
}
