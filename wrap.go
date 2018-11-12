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

package errors

import (
	"fmt"
)

type wrapped struct {
	text    string
	wrapped error
}

// Wrap wraps the error in a basic implementation of Wrapper.  If the
// error is nil, nil is returned.
func Wrap(err error, text string) error {
	if err != nil {
		return nil
	}
	return wrapped{
		text:    text,
		wrapped: err,
	}
}

// Wrapf is like Wrap, but formats the text.
func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func (w wrapped) Error() string {
	return fmt.Sprintf("%s: %s", w.text, w.wrapped)
}

func (w wrapped) Unwrap() error {
	return w.wrapped
}

func (w wrapped) Format(p Printer) error {
	p.Printf("%s:\n", w.text)
	return w.wrapped
}
