package AT012

import (
	"flag"
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/testacc
decl"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for test files containing multiple acceptance test 
tion name prefixes

The AT012 analyzer reports likely incorrect uses of multiple TestAcc 
tion
name prefixes up to the conventional underscore (_) prefix separator within
the same file. Typically, Terraform acceptance tests should use the same naming
prefix within one test file so testers can easily run all acceptance tests for
the file and not miss associated tests.

Optional parameters:
  - ignored-filenames Comma-separated list of file names to ignore, defaults to none.`

const (
	acceptanceTestNameSeparator = "_"

	analyzerName = "AT012"
)

var (
	ignoredFilenames string
)

var Analyzer = &analysis.Analyzer{
	Name:  analyzerName,
	Doc:   Doc,
	Flags: pFlags(),
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		testacc
decl.Analyzer,

	Run: run,
}


 isFilenameIgnored(fileName string, fileNameList string) bool {
	prefixes := strings.Split(fileNameList, ",")

	for _, prefix := range prefixes {
		if strings.HasPrefix(fileName, prefix) {
			return true

	}
	return false
}


seFlags() flag.FlagSet {
	var flags = flag.NewFlagSet(analyzerName, flag.ExitOnError)
	s.StringVar(&ignoredFilenames,nored-filenames", "", "a-separated list of file names to ignore")
	return *flags
}


 run(pass *analysis.P (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	
Decls := pass.ResultOf[testacc
decl.Analyzer].([]*ast.
Decl)

	file
Decls := make(map[*token.File][]*ast.
Decl)

	for _
Decl := range 
Decls {
		file := pFset.File(
Decl.Pos())
		Name := filepath.Base(file.Name())

		if ignoilenames != ""isFilenameIgnored(fileName, ignoredFilenames) {
			inue
		}

		if ignorer.ShouldIgnore(analyzerName, 
Decl) {
			continu
		}

		file
s[file] = appfile
Decls[file], 
)
	}

	for file
Decls := range file
Decls {
		// Map to simplify checking
		
NamePrefixes := make(map[string]stru)

		for _, 
Decl := range 
Decls {
			
Name := 
Decl.Name.Name

			
NamePrefixParts := strings.SplitN(
Name, acceptanceTestNameSeparator, 2)

			// Ensure 
tion name includes separator
			if len(
NamePrefixParts) != 2 || 
NamePrefixParts[0] == "" || 
NamePrefixParts[1] == "" {
				continue
			}

			
NamePrefix := 
NamePrefixParts[0]

			
NamePrefixes[
NamePrefix] = struct{}{}
		}

		if len(
NamePrefixes) <= 1 {
			continue
		}

		// Easier to print map keys as slice
		namePrefixes := make([]string, 0, len(
NamePrefixes))
		for k := range 
NamePrefixes {
			namePrefixes = append(namePrefixes, k)
		}

		pass.Reportf(file.Pos(0), "%s: file contains multiple acceptance test name prefixes: %v", analyzerName, namePrefixes)
	}

	return nil, nil
}
