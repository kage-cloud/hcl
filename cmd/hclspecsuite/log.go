package main

import (
	"github.com/kage-cloud/hcl/v2"
)

type LogBeginCallback func(testName string, testFile *TestFile)
type LogProblemsCallback func(testName string, testFile *TestFile, diags hcl.Diagnostics)
