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

type Wrapper struct {
	Title  string
	Detail string
	Err    error
}

func (w Wrapper) Error() string {
	return w.Title
}

func (w Wrapper) Format(p errors.Printer) error {
	p.Printf("%s:\n", w.Title)
	if p.Detail() {
		p.Printf("    %s\n", w.Detail)
	}
	return w.Err
}

type RootError struct {
	Title  string
	Detail string
}

func (e RootError) Error() string {
	return e.Title
}

func (e RootError) Format(p errors.Printer) error {
	p.Printf("%s:\n", e.Title)
	if p.Detail() {
		p.Printf("    %s", e.Detail)
	}
	return nil
}

func ExampleFormat() {
	err := Wrapper{
		Title:  "foo",
		Detail: "some context",
		Err: Wrapper{
			Title:  "spam",
			Detail: "stuff",
			Err: RootError{
				Title:  "egg",
				Detail: "stuff",
			},
		},
	}
	fmt.Print(errors.Format(err, false))
	fmt.Println()
	fmt.Print(errors.Format(err, true))
	// Output:
	// foo:
	// --- spam:
	// --- egg:
	//
	// foo:
	//     some context
	// --- spam:
	//     stuff
	// --- egg:
	//     stuff
}
