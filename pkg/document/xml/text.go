/*
 * Copyright 2023 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package xml

// Text is a text node in XML. It is a leaf node.
type Text struct {
	value string
}

// NewText creates a new instance of Text.
func NewText(value string) *Text {
	return &Text{
		value: value,
	}
}

// String returns the string representation of this Text.
func (t *Text) String() string {
	return t.value
}
