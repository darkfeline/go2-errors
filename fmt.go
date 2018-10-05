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
	"bytes"
	"fmt"
)

// Format formats the error using the proposed style.
func Format(err error, detail bool) string {
	if f, ok := err.(Formatter); ok {
		var b bytes.Buffer
		p := fmtPrinter{
			b:      &b,
			detail: detail,
		}
		formatError(&p, f)
		return p.b.String()
	}
	return err.Error()
}

func formatError(p Printer, f Formatter) {
	next := f.Format(p)
	if next == nil {
		return
	}
	if f, ok := next.(Formatter); ok {
		p.Print("--- ")
		formatError(p, f)
		return
	}
	p.Print(next)
}

type fmtPrinter struct {
	b      *bytes.Buffer
	detail bool
}

func (p *fmtPrinter) Print(args ...interface{}) {
	fmt.Fprint(p.b, args...)
}

func (p *fmtPrinter) Printf(format string, args ...interface{}) {
	fmt.Fprintf(p.b, format, args...)
}

func (p *fmtPrinter) Detail() bool {
	return p.detail
}
