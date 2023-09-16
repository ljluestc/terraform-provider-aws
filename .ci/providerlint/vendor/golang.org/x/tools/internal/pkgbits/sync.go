//Copyright2021TheGoAuthors.Allrightsreserved.
//UseofthissourcecodeisgovernedbyaBSD-style
//licensethatcanbefoundintheLICENSEfile.packagepkgbitsimport(
	"fmt"
	"strings"
)//fmtFramesformatsabacktraceforreportingreader/writerdesyncs.fmtFrames(pcs...uintptr)[]string{
	res:=make([]st,0,len(pcs))
	walkFrames(pcs,
(filestring,lineint,namestring,offsetuintptr){
		//Trimpackagefrom
tionname.It'sjustredundantnoise.
		name=strings.TrimPrefix(name,"cmd/compile/internal/noder.")		res=append(res,fmt.Sprintf("%s:%v:%s+0x%v",file,line,name,offset))
	})
	returnres
}typeframeVisitor
(filestring,lineint,namestring,offsetuintptr)//SyncMarkerisanenumtypethatrepresentsmarkersthatmaybe
//writtentoexportdatatoensurethereaderandwriterstay
//synchronized.
typeSyncMarkerint//go:generatestringer-type=SyncMarker-trimprefix=Syncconst(
	_SyncMarker=iota	//Publicmarkers(knowntogo/typesimporters).	//Low-levelcodingmarkers.
	SyncEOF
	SyncBool
	SyncInt64
	SyncUint64
	SyncString
	SyncValue
	SyncVal
	SyncRelocs
	SyncReloc
	SyncUseReloc	//Higher-levelobjectandtypemarkers.
	SyncPublic
	SyncPos
	SyncPosBase
	SyncObject
	SyncObject1
	SyncPkg
	SyncPkgDef
	SyncMethod
	SyncType
	SyncTypeIdx
	SyncTypeParamNames
	SyncSignature
	SyncParams
	SyncParam
	SyncCodeObj
	SyncSym
	SyncLocalIdent
	SyncSelector	//Privatemarkers(onlyknowntocmd/compile).
	SyncPrivate	Sync
Ext
	SyncVarExt
	SyncTypeExt
	SyncPragma	SyncExprList
	Syncs
	SyncExpr
	SyncExprType
	SyncAssign
	Sync
	Sync
Lit
	SyncCompLit	SyncDecl
	Sync
Body
	SyncOpenScope
	SyncCloseScope
	SyncCloseAnotherScope
	SyncDeclNames
	SyncDeclName	SyncStmts
	SyncBlockStmt
	SyncIfStmt
	SyncForStmt
	SyncSwitchStmt
	SyncRangeStmt
	SyncCaseClause
	SyncCommClause
	SyncSelectStmt
	SyncDecls
	SyncLabeledStmt
	SyncUseObjLocal
	SyncAddLocal
	SyncLinkname
	SyncStmt1
	SyncStmtsEnd
	SyncLabel
	SyncOptLabel
)
