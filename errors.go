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

// Package errors implements the error values proposal for Go 2.
//
// https://go.googlesource.com/proposal/+/master/design/go2draft-error-values-overview.md
// https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md
// https://go.googlesource.com/proposal/+/master/design/go2draft-error-printing.md
package errors

import (
	"errors"
	"reflect"
)

var New = errors.New

// Wrapper defines the interface for errors wrapping another error.
type Wrapper interface {
	// Unwrap returns the wrapped error.
	Unwrap() error
}

// Is reports whether err or any of the errors in its chain is equal to target.
func Is(err, target error) bool {
	if err == target {
		return true
	}
	if w, ok := err.(Wrapper); ok {
		return Is(w.Unwrap(), target)
	}
	return false
}

// As checks whether err or any of the errors in its chain is a value of type E.
// If so, it returns the discovered value of type E, with ok set to true.
// If not, it returns the zero value of type E, with ok set to false.
func As(t reflect.Type, err error) (e error, ok bool) {
	if reflect.TypeOf(err) == t {
		return err, true
	}
	if w, ok := err.(Wrapper); ok {
		return As(t, w.Unwrap())
	}
	return nil, false
}

// AsValue checks whether err or any of the errors in its chain is a
// value of the type t points to.  If so, it returns true, with t set to the
// discovered value set.  If not, it returns false.
func AsValue(t interface{}, err error) (ok bool) {
	v := reflect.ValueOf(t)
	if v.Type().Kind() != reflect.Ptr {
		panic("non-pointer passed to AsValue")
	}
	return asValue(v.Elem(), err)
}

func asValue(v reflect.Value, err error) (ok bool) {
	et := reflect.TypeOf(err)
	vt := v.Type()
	if et == vt || vt.Kind() == reflect.Interface && et.Implements(vt) {
		v.Set(reflect.ValueOf(err))
		return true
	}
	if w, ok := err.(Wrapper); ok {
		return asValue(v, w.Unwrap())
	}
	return false
}

// A Formatter formats error messages.
type Formatter interface {
	// Format is implemented by errors to print a single error message.
	// It should return the next error in the error chain, if any.
	Format(p Printer) (next error)
}

// A Printer creates formatted error messages. It enforces that
// detailed information is written last.
//
// Printer is implemented by fmt. Localization packages may provide
// their own implementation to support localized error messages
// (see for instance golang.org/x/text/message).
type Printer interface {
	// Print appends args to the message output.
	// String arguments are not localized, even within a localized context.
	Print(args ...interface{})

	// Printf writes a formatted string.
	Printf(format string, args ...interface{})

	// Detail reports whether error detail is requested.
	// After the first call to Detail, all text written to the Printer
	// is formatted as additional detail, or ignored when
	// detail has not been requested.
	// If Detail returns false, the caller can avoid printing the detail at all.
	Detail() bool
}
