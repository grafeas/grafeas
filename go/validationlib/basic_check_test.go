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

package validationlib

import "testing"

func TestIsAlpha(t *testing.T) {
	// passing cases with input generator:
	g := NewInputGenerator()
	lengthToTest := []int{0, 10, 50, 100, 200, 500, 1000000}
	for _, length := range lengthToTest {
		tmpStr := g.GenStringAlpha(length)
		if !IsAlpha(tmpStr) {
			t.Errorf("%s is expected to be alphabetic only", tmpStr)
		}
	}

	// failing cases:
	badInputs := []string{"你好", "\b5Ὂg̀9! ℃ᾭG"}
	for _, badInput := range badInputs {
		if IsAlpha(badInput) {
			t.Errorf("%s is expected to be not alphabetic only", badInput)
		}
	}
}

func TestIsURLFriendly(t *testing.T) {
	// passing cases with input generator:
	g := NewInputGenerator()
	lengthToTest := []int{0, 10, 50, 100, 200, 500, 1000000}
	for _, length := range lengthToTest {
		tmpStr := g.GenStringURLFriendly(length)
		if !IsURLFriendly(tmpStr) {
			t.Errorf("%s is expected to be URL friendly", tmpStr)
		}
	}

	// passing cases:
	goodInputs := []string{"a~", "a-b", "a.b", "a_", "19a"}
	for _, goodInput := range goodInputs {
		if !IsURLFriendly(goodInput) {
			t.Errorf("%s is expected to be URL friendly", goodInput)
		}
	}

	// failing cases:
	badInputs := []string{"你好", "a!", "a@", "a#", "a$", "a%", "a^",
		"a&", "a*", "a(", "a)"}
	for _, badInput := range badInputs {
		if IsURLFriendly(badInput) {
			t.Errorf("%s is expected to be not URL friendly", badInput)
		}
	}
}
