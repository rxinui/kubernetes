/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package get

import (
	// "fmt"
	// "os"
	// "strings"

	"github.com/spf13/cobra"

	// "k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/printers"
	// "k8s.io/kubectl/pkg/scheme"
)

const extraColumnsFormat = "extra-columns"

// ExtraColumnsPrintFlags provides default flags necessary for printing
// custom resource columns from an inline-template or file.
type ExtraColumnsPrintFlags struct {
	NoHeaders        bool
	TemplateArgument string
	*HumanPrintFlags
	*CustomColumnsPrinter
}

func (f *ExtraColumnsPrintFlags) AllowedFormats() []string {
	return []string{extraColumnsFormat}
}

// ToPrinter receives an templateFormat and returns a printer capable of
// handling extra-columns
// Returns false if the specified templateFormat does not match a supported format.
// Supported format types can be found in pkg/printers/printers.go
func (f *ExtraColumnsPrintFlags) ToPrinter(templateFormat string) (printers.ResourcePrinter, error) {
	return f.HumanPrintFlags.ToPrinter(templateFormat)
}

// AddFlags receives a *cobra.Command reference and binds
// flags related to custom-columns printing
func (f *ExtraColumnsPrintFlags) AddFlags(c *cobra.Command) {}

// NewExtraColumnsPrintFlags returns flags associated with
// NoHeaders and TemplateArgument should be set by callers.
func NewExtraColumnsPrintFlags() *ExtraColumnsPrintFlags {
	return &ExtraColumnsPrintFlags{
		NoHeaders:        false,
		TemplateArgument: "",
		HumanPrintFlags:  NewHumanPrintFlags(),
	}
}

// ExtraColumnsPrinter should be Union of HumanReadablePrinter with CustomColumnsPrinter in that order
type ExtraColumnsPrinter struct {
	Columns                        []Column // should be HumanReadablePrinter column + CustomColumnsPrinter column
	NoHeaders                      bool
	*CustomColumnsPrinter          // has PrintObj (implements ResourcePrinter)
	*printers.HumanReadablePrinter // has PrintObj (implements ResourcePrinter)
}

func (s *ExtraColumnsPrinter) PrintObj(obj runtime.Object, out io.Writer) error {
	return nil
}
