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

package validation

import (
	"strings"
	"testing"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input   string
		isalpha bool
	}{
		{
			input:   strings.Repeat("a", 0),
			isalpha: true,
		},
		{
			input:   strings.Repeat("a", 100),
			isalpha: true,
		},
		{
			input:   strings.Repeat("a", 1000000),
			isalpha: true,
		},
		{
			input:   "你好",
			isalpha: false,
		},
		{
			input:   "\b5Ὂg̀9! ℃ᾭG",
			isalpha: false,
		},
	}

	for _, tt := range tests {
		if IsAlpha(tt.input) != tt.isalpha {
			modifier := " "
			if !tt.isalpha {
				modifier = " not "
			}
			t.Errorf("%s is expected to be%salphabetic", tt.input, modifier)
		}
	}
}

func TestIsURLFriendly(t *testing.T) {
	tests := []struct {
		input         string
		isurlfriendly bool
	}{
		{
			input:         strings.Repeat("a", 0),
			isurlfriendly: true,
		},
		{
			input:         strings.Repeat("a", 100),
			isurlfriendly: true,
		},
		{
			input:         strings.Repeat("a", 1000000),
			isurlfriendly: true,
		},
		{
			input:         "a~",
			isurlfriendly: true,
		},
		{
			input:         "a-b",
			isurlfriendly: true,
		},
		{
			input:         "a.b",
			isurlfriendly: true,
		},
		{
			input:         "a_",
			isurlfriendly: true,
		},
		{
			input:         "19a",
			isurlfriendly: true,
		},
		{
			input:         "你好",
			isurlfriendly: false,
		},
		{
			input:         "\b5Ὂg̀9! ℃ᾭG",
			isurlfriendly: false,
		},
		{
			input:         "a!",
			isurlfriendly: false,
		},
		{
			input:         "a@",
			isurlfriendly: false,
		},
		{
			input:         "a#",
			isurlfriendly: false,
		},
		{
			input:         "a$",
			isurlfriendly: false,
		},
		{
			input:         "a%",
			isurlfriendly: false,
		},
		{
			input:         "a^",
			isurlfriendly: false,
		},
		{
			input:         "a&",
			isurlfriendly: false,
		},
		{
			input:         "a*",
			isurlfriendly: false,
		},
		{
			input:         "a(",
			isurlfriendly: false,
		},
		{
			input:         "a)",
			isurlfriendly: false,
		},
	}

	for _, tt := range tests {
		if IsURLFriendly(tt.input) != tt.isurlfriendly {
			modifier := " "
			if !tt.isurlfriendly {
				modifier = " not "
			}
			t.Errorf("%s is expected to be%sURL friendly", tt.input, modifier)
		}
	}
}
