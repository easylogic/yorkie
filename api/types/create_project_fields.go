/*
 * Copyright 2022 The Yorkie Authors. All rights reserved.
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
 *
 */

package types

import (
	"github.com/yorkie-team/yorkie/internal/validation"
)

// reservedProjectNames is a map of reserved names. It is used to check if the
// given project name is reserved or not.
var (
	reservedProjectNames = map[string]bool{"new": true, "default": true}
)

// CreateProjectFields is a set of fields that use to create a project.
type CreateProjectFields struct {
	// Name is the name of this project.
	Name *string `bson:"name,omitempty" validate:"required,min=2,max=30,slug,reserved_project_name"`
}

// Validate validates the CreateProjectFields.
func (i *CreateProjectFields) Validate() error {
	return validation.ValidateStruct(i)
}

func isReservedProjectName(name string) bool {
	_, ok := reservedProjectNames[name]
	return ok
}

func init() {
	validation.RegisterValidation("reserved_project_name", func(level validation.FieldLevel) bool {
		name := level.Field().String()
		return !isReservedProjectName(name)
	})

	validation.RegisterTranslation("reserved_project_name", "given {0} is reserved name")
}
