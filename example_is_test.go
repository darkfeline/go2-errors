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

	"go.felesatra.moe/go2/errors"
)

type FooError struct{}

func (FooError) Error() string {
	return "foo"
}

func ExampleIs() {
	fmt.Printf("Plain error is FooError: %t\n", errors.Is(errors.New("text"), FooError{}))
	fmt.Printf("Wrapped plain error is FooError: %t\n",
		errors.Is(errors.Wrap(errors.New("text"), "text"), FooError{}))
	fmt.Printf("FooError is FooError: %t\n", errors.Is(FooError{}, FooError{}))
	fmt.Printf("Wrapped FooError is FooError: %t\n",
		errors.Is(errors.Wrap(FooError{}, "text"), FooError{}))
	// Output:
	// Plain error is FooError: false
	// Wrapped plain error is FooError: false
	// FooError is FooError: true
	// Wrapped FooError is FooError: true
}
