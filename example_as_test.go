// Copyright (C) 2018  Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors_test

import (
	"fmt"
	"reflect"

	"go.felesatra.moe/go2/errors"
)

type BarError struct {
	data string
}

func (BarError) Error() string {
	return "bar"
}

var barType = reflect.TypeOf(BarError{})

func ExampleAs() {
	err := errors.Wrap(BarError{data: "ayanami"}, "text")
	if e, ok := errors.As(barType, err); ok {
		err := e.(BarError)
		fmt.Printf("Bar data: %s", err.data)
	}
	// Output:
	// Bar data: ayanami
}
