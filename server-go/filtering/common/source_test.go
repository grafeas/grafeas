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

// package common defines types common to parsing and other diagnostics.
package common

import (
	"testing"
)

// Test the error description method.
func TestStringSource_Description(t *testing.T) {
	contents := "example content\nsecond line"
	source := NewStringSource(contents, "description-test")
	assertEquals(t, source.Content(), contents)
	assertEquals(t, source.Description(), "description-test")
	str2, found := source.Snippet(2)
	assertTrue(t, found)
	assertEquals(t, "second line", str2)
	str1, found := source.Snippet(1)
	assertTrue(t, found)
	assertEquals(t, "example content", str1)
}

// Test the character offest to make sure that the offsets accurately reflect
// the location of a character in source.
func TestStringSource_CharacterOffset(t *testing.T) {
	contents := "c.d &&\n\t b.c.arg(10) &&\n\t test(10)"
	source := NewStringSource(contents, "offset-test")
	assertHasExactly(t, []int32{7, 24, 35}, source.LineOffsets())
	charStart, _ := source.CharacterOffset(NewLocation(1, 2))
	charEnd, _ := source.CharacterOffset(NewLocation(3, 2))
	assertEquals(t, "d &&\n\t b.c.arg(10) &&\n\t ", string(contents[charStart:charEnd]))
	if _, found := source.CharacterOffset(NewLocation(4, 0)); found {
		t.Error("Character offset was out of range of source, but still found.")
	}
}

// Test the computation of snippets, single lines of text, from a multiline
// source.
func TestStringSource_SnippetMultiline(t *testing.T) {
	source := NewStringSource("hello\nworld\nmy\nbub\n", "four-line-test")
	str, found := source.Snippet(1)
	assertTrue(t, found)
	assertEquals(t, "hello", str)
	str2, found := source.Snippet(2)
	assertTrue(t, found)
	assertEquals(t, "world", str2)
	str3, found := source.Snippet(3)
	assertTrue(t, found)
	assertEquals(t, "my", str3)
	str4, found := source.Snippet(4)
	assertTrue(t, found)
	assertEquals(t, "bub", str4)
	str5, found := source.Snippet(5)
	assertTrue(t, found)
	assertEquals(t, "", str5)
}

// Test the computation of snippets from a single line source.
func TestStringSource_SnippetSingleline(t *testing.T) {
	source := NewStringSource("hello, world", "one-line-test")
	str, found := source.Snippet(1)
	assertTrue(t, found)
	assertEquals(t, "hello, world", str)
	str2, found := source.Snippet(2)
	assertFalse(t, found)
	assertEquals(t, "", str2)
}

// Helper assertions.
func assertTrue(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("[%s] Expected 'true', got 'false'", t.Name())
	}
}

func assertFalse(t *testing.T, actual bool) {
	if actual {
		t.Errorf("[%s] Expected 'false', got 'true'", t.Name())
	}
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("[%s] Expected '%v', got '%v'", t.Name(), expected, actual)
	}
}

func assertHasExactly(t *testing.T, expected []int32, actual []int32) {
	if len(expected) != len(actual) {
		t.Errorf("[%s] Expected list of size '%d', got a list of size '%d'", t.Name(), len(expected), len(actual))
	}
	for i, val := range expected {
		assertEquals(t, val, actual[i])
	}
}
