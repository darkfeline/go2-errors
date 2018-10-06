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

type Frobber interface {
	Frob() string
}

type FrobErr struct {
	data string
}

func (FrobErr) Error() string {
	return "frob"
}

func (f FrobErr) Frob() string {
	return f.data
}

func ExampleAsValue() {
	err := errors.Wrap(FrobErr{data: "ayanami"}, "text")
	var e FrobErr
	if errors.AsValue(&e, err) {
		fmt.Printf("data: %s\n", e.data)
	}
	var e2 Frobber
	if errors.AsValue(&e2, err) {
		fmt.Printf("interface data: %s\n", e2.Frob())
	}
	err = errors.New("test")
	if errors.AsValue(&e, err) {
		fmt.Printf("plain data: %s\n", e.data)
	}
	if errors.AsValue(&e2, err) {
		fmt.Printf("plain interface data: %s\n", e2.Frob())
	}
	// Output:
	// data: ayanami
	// interface data: ayanami
}
