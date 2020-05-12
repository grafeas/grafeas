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
)

const (
	// MaxNoteIDLength stands for the max length allowed for a note_id field.
	MaxNoteIDLength = 128
	// AlphaCharset accounts for alphabetic characters.
	AlphaCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// AlphaNumCharset accounts for the alphabetic and numeric characters.
	AlphaNumCharset = AlphaCharset + "0123456789"
	// URLFriendlyCharset accounts for the URL friendly characters.
	// In more details, API style says that resource IDs should use URL-friendly
	// resource IDs. See here: https://cloud.google.com/apis/design/resource_names#resource_id.
	// The official spec allows ALPHA / DIGIT / "-" / "." / "_" / "~" as a subset
	// See here: https://tools.ietf.org/html/rfc3986#appendix-A
	URLFriendlyCharset = AlphaNumCharset + "-._~"
)

// IsAlpha checks whether s has alphabetic characters only.
func IsAlpha(s string) bool {
	// use the simple brutal force method which offers the reasonable complexity
	// O(m*n) where m and n are the length of charset and input string.
	for _, char := range s {
		if !strings.Contains(AlphaCharset, string(char)) {
			return false
		}
	}
	return true
}

// IsURLFriendly checks whether s has URL friendly characters only.
func IsURLFriendly(s string) bool {
	// use the simple brutal force method which offers the reasonable complexity
	// O(m*n) where m and n are the length of charset and input string.
	for _, char := range s {
		if !strings.Contains(URLFriendlyCharset, string(char)) {
			return false
		}
	}
	return true
}
