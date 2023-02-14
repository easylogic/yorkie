/*
 * Copyright 2020 The Yorkie Authors. All rights reserved.
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

package key

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func findErrorOnErrors(errs error, tag string) bool {
	for _, v := range errs.(validator.ValidationErrors) {
		if v.Tag() == tag {
			return true
		}
	}

	return false
}

func TestKey_IsValid(t *testing.T) {
	t.Run("valid key", func(t *testing.T) {
		key := Key("valid-key")
		ret := key.IsValid()
		assert.Nil(t, ret, "key should be valid")

		key = Key("valid-key-1")
		ret = key.IsValid()
		assert.Nil(t, ret, "key should be valid")

		key = Key("fdsxfdsf")
		ret = key.IsValid()
		assert.Nil(t, ret, "key should be valid")

		key = Key("-----_________________-a")
		ret = key.IsValid()
		assert.Nil(t, ret, "key should be valid")
	})

	t.Run("invalid key", func(t *testing.T) {
		key := Key("invalid key") // space is not allowed

		errs := key.IsValid()

		assert.True(t, findErrorOnErrors(errs, "slug"), "key should be invalid: with space")

		key = Key("invalid-key-~$a") // last character should be alphanumeric
		errs = key.IsValid()
		assert.True(t, findErrorOnErrors(errs, "slug"), "key should be invalid: with -")

		key = Key("invalid-key-$") // last character should be alphanumeric
		errs = key.IsValid()
		assert.True(t, findErrorOnErrors(errs, "slug"), "key should be invalid: with $")

		key = Key("invalid-key-sample-key-validator") // last character should be alphanumeric
		errs = key.IsValid()
		assert.True(t, findErrorOnErrors(errs, "max"), "key should be invalid: check max length, 30 ")

		key = Key("inv") // last character should be alphanumeric
		errs = key.IsValid()
		assert.True(t, findErrorOnErrors(errs, "min"), "key should be invalid: check min length, 4 ")
	})
}