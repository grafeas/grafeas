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

import (
	"math/rand"
	"time"
)

// InputGenerator generates different types of inputs to be used in testing.
type InputGenerator struct {
	SeededRand *rand.Rand
}

// NewInputGenerator creates an InputGenerator instance and returns the pointer.
func NewInputGenerator() *InputGenerator {
	result := InputGenerator{
		SeededRand: rand.New(
			rand.NewSource(time.Now().UnixNano())),
	}
	return &result
}

// GenStringAlpha generates a string of specified length with the alphabetic character set.
func (g *InputGenerator) GenStringAlpha(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = AlphaCharset[g.SeededRand.Intn(len(AlphaCharset))]
	}
	return string(b)
}

// GenStringURLFriendly generates a string of specified length with the URL friendly character set.
func (g *InputGenerator) GenStringURLFriendly(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = URLFriendlyCharset[g.SeededRand.Intn(len(URLFriendlyCharset))]
	}
	return string(b)
}
