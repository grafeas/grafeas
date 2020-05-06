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

func TestGenStringAlpha(t *testing.T) {
	g := NewInputGenerator()
	lengthToTest := []int{0, 10, 50, 100, 200, 500, 1000000}
	for _, length := range lengthToTest {
		tmpStr := g.GenStringAlpha(length)
		if len(tmpStr) != length {
			t.Errorf("length of generated string %s != expected length %d", tmpStr, length)
		}

		if !IsAlpha(tmpStr) {
			t.Errorf("%s is expected to be alphabetic only", tmpStr)
		}
	}
}

func TestGenStringURLFriendly(t *testing.T) {
	g := NewInputGenerator()
	lengthToTest := []int{0, 10, 50, 100, 200, 500, 1000000}
	for _, length := range lengthToTest {
		tmpStr := g.GenStringURLFriendly(length)
		if len(tmpStr) != length {
			t.Errorf("length of generated string %s != expected length %d", tmpStr, length)
		}

		if !IsURLFriendly(tmpStr) {
			t.Errorf("%s is expected to be URL friendly", tmpStr)
		}
	}
}
