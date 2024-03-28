package filters

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"strings"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

func IsALint(file *lint.File) bool {
	return strings.HasPrefix(file.Name, "lint_") && IsAGoFile(file) && !IsATest(file)
}

func IsAGoFile(file *lint.File) bool {
	return strings.HasSuffix(file.Name, ".go")
}

func IsATest(file *lint.File) bool {
	return strings.HasSuffix(file.Name, "test.go")
}
