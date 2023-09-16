package osexeccommandcallexpr

import (
	"github.com/bflad/tfproviderlint/helper/analysisutils"
)

var Analyzer = analysisutils.Stdlib
CallExprAnalyzer(
	"osexeccommandcallexpr",
	"os/exec",
	"Command",
)
