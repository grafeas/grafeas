// Copyright 2018 The Grafeas Authors. All rights reserved.
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

package validators

import (
	"strings"
	"testing"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input   string
		isAlpha bool
	}{
		{
			input:   strings.Repeat("a", 0),
			isAlpha: true,
		},
		{
			input:   strings.Repeat("a", 100),
			isAlpha: true,
		},
		{
			input:   strings.Repeat("a", 1000000),
			isAlpha: true,
		},
		{
			input:   "你好",
			isAlpha: false,
		},
		{
			input:   "\b5Ὂg̀9! ℃ᾭG",
			isAlpha: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if isAlpha := IsAlpha(tt.input); isAlpha != tt.isAlpha {
				t.Errorf("IsAlpha(%q) = %t, want %t", tt.input, isAlpha, tt.isAlpha)
			}
		})
	}
}

func TestIsURLFriendly(t *testing.T) {
	tests := []struct {
		input         string
		isURLFriendly bool
	}{
		{
			input:         strings.Repeat("a", 0),
			isURLFriendly: true,
		},
		{
			input:         strings.Repeat("a", 100),
			isURLFriendly: true,
		},
		{
			input:         strings.Repeat("a", 1000000),
			isURLFriendly: true,
		},
		{
			input:         "a~",
			isURLFriendly: true,
		},
		{
			input:         "a-b",
			isURLFriendly: true,
		},
		{
			input:         "a.b",
			isURLFriendly: true,
		},
		{
			input:         "a_",
			isURLFriendly: true,
		},
		{
			input:         "19a",
			isURLFriendly: true,
		},
		{
			input:         "你好",
			isURLFriendly: false,
		},
		{
			input:         "\b5Ὂg̀9! ℃ᾭG",
			isURLFriendly: false,
		},
		{
			input:         "a!",
			isURLFriendly: false,
		},
		{
			input:         "a@",
			isURLFriendly: false,
		},
		{
			input:         "a#",
			isURLFriendly: false,
		},
		{
			input:         "a$",
			isURLFriendly: false,
		},
		{
			input:         "a%",
			isURLFriendly: false,
		},
		{
			input:         "a^",
			isURLFriendly: false,
		},
		{
			input:         "a&",
			isURLFriendly: false,
		},
		{
			input:         "a*",
			isURLFriendly: false,
		},
		{
			input:         "a(",
			isURLFriendly: false,
		},
		{
			input:         "a)",
			isURLFriendly: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if isURLFriendly := IsURLFriendly(tt.input); isURLFriendly != tt.isURLFriendly {
				t.Errorf("IsURLFriendly(%q) = %t, want %t", tt.input, isURLFriendly, tt.isURLFriendly)
			}
		})
	}
}
