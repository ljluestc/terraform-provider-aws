packageAT012

import(
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

constDoc=`checkfortestfilescontainingmultipleacceptancetest
tionnameprefixes

TheAT012analyzerreportslikelyincorrectusesofmultipleTestAcc
tion
nameprefixesuptotheconventionalunderscore(_)prefixseparatorwithin
thesamefile.Typically,Terraformacceptancetestsshouldusethesamenaming
prefixwithinonetestfilesotesterscaneasilyrunallacceptancetestsfor
thefileandnotmissassociatedtests.

Optionalparameters:
-ignored-filenamesComma-separatedlistoffilenamestoignore,defaultstonone.`

const(
	acceptanceTestNameSeparator="_"

	analyzerName="AT012"
)

var(
	ignoredFilenamesstring
)

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Flags:pFlags(),
	Requires:[]*analysis.Analyzer{
		commentignore.Analyzer,
		testacc
decl.Analyzer,

	Run:run,
}


isFilenameIgnored(fileNamestring,fileNameListstring)bool{
	prefixes:=strings.Split(fileNameList,",")

	for_,prefix:=rangeprefixes{
		ifstrings.HasPrefix(fileName,prefix){
			returntrue

	}
	returnfalse
}


seFlags()flag.FlagSet{
	varflags=flag.NewFlagSet(analyzerName,flag.ExitOnError)
	s.StringVar(&ignoredFilenames,nored-filenames","","a-separatedlistoffilenamestoignore")
	return*flags
}


run(pass*analysis.P(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	
Decls:=pass.ResultOf[testacc
decl.Analyzer].([]*ast.
Decl)

	file
Decls:=make(map[*token.File][]*ast.
Decl)

	for_
Decl:=range
Decls{
		file:=pFset.File(
Decl.Pos())
		Name:=filepath.Base(file.Name())

		ifignoilenames!=""isFilenameIgnored(fileName,ignoredFilenames){
			inue
		}

		ifignorer.ShouldIgnore(analyzerName,
Decl){
			continu
		}

		file
s[file]=appfile
Decls[file],
)
	}

	forfile
Decls:=rangefile
Decls{
		//Maptosimplifychecking
		
NamePrefixes:=make(map[string]stru)

		for_,
Decl:=range
Decls{
			
Name:=
Decl.Name.Name

			
NamePrefixParts:=strings.SplitN(
Name,acceptanceTestNameSeparator,2)

			//Ensure
tionnameincludesseparator
			iflen(
NamePrefixParts)!=2||
NamePrefixParts[0]==""||
NamePrefixParts[1]==""{
				continue
			}

			
NamePrefix:=
NamePrefixParts[0]

			
NamePrefixes[
NamePrefix]=struct{}{}
		}

		iflen(
NamePrefixes)<=1{
			continue
		}

		//Easiertoprintmapkeysasslice
		namePrefixes:=make([]string,0,len(
NamePrefixes))
		fork:=range
NamePrefixes{
			namePrefixes=append(namePrefixes,k)
		}

		pass.Reportf(file.Pos(0),"%s:filecontainsmultipleacceptancetestnameprefixes:%v",analyzerName,namePrefixes)
	}

	returnnil,nil
}
