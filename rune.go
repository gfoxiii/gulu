// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
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

package gulu

type guluRune struct{}

// Rune utilities.
var Rune = guluRune{}

// IsNumOrLetter checks the specified rune is number or letter.
func (*guluRune) IsNumOrLetter(r rune) bool {
	return Rune.IsLetter(r) || ('0' <= r && '9' >= r)
}

// IsLetter checks the specified rune is letter.
func (*guluRune) IsLetter(r rune) bool {
	return 'a' <= r && 'z' >= r || 'A' <= r && 'Z' >= r
}
