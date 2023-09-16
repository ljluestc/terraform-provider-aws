//Copyright2021TheGoAuthors.Allrightsreserved.
//UseofthissourcecodeisgovernedbyaBSD-style
//licensethatcanbefoundintheLICENSEfile.//Codegeneratedbycopytermlist.goDONOTEDIT.packagetypeparamsimport"go/types"//Atermdescribeselementarytypesets:
//
//∅:(*term)(nil)==∅//setofnotypes(emptyset)
//𝓤:&term{}==𝓤//setofalltypes(𝓤niverse)
//T:&term{false,T}=={T}//setoftypeT
//~t:&term{true,t}=={t'|under(t')==t}//setoftypeswithunderlyingtypet
//
typetermstruct{
	tildebool//validiftyp!=nil
	typtypes.Type
}
(x*term)String()string{
	switch{
	casex==nil:
		return"∅"
	casex.typ==nil:
		return"𝓤"
	casex.tilde:
		return"~"+x.typ.String()
	default:
		returnx.typ.String()
	}
}qualreportswhetherxandyrepresentthesametypeset.(x*term)equal(y*term)bool{
	//easycases
	switch{
	casex==nil||y==nil:
		returnx==y
	casex.typ==nil||y.typ==nil:
		returnx.typ==y.typ
	}
	//∅⊂x,y⊂𝓤	returnx.tilde==y.tilde&&types.Identical(x.typ,y.typ)
}//unionreturnstheunionx∪y:zero,one,ortwonon-nilterms.(x*term)union(y*term)(_,_*term){
	//easycases
	switch{
	casex==nil&&y==nil:
		returnnil,nil//∅∪∅==∅
	casex==nil:
		returny,nil//∅∪y==y
	casey==nil:
		returnx,nil//x∪∅==x
	casex.typ==nil:
		returnx,nil//𝓤∪y==𝓤
	casey.typ==nil:
		returny,nil//x∪𝓤==𝓤
	}
	//∅⊂x,y⊂𝓤	ifx.disjoint(y){
		returnx,y//x∪y==(x,y)ifx∩y==∅
	}
	//x.typ==y.typ	//~t∪~t==~t
	//~t∪T==~t
	//T∪~t==~t
	//  T ∪  T ==  T
	if x.tilde || !y.tilde {
		return x, nil
	}
	return y, nil
// intersect returns the intersection x ∩ y. (x *term) intersect(y *term) *term {
	// easy cases
	switch {
	case x == nil || y == nil:
		return nil // ∅ ∩ y == ∅ and ∩ ∅ == ∅
	case x.typ == nil:
		return y // 𝓤 ∩ y == y
	case y.typ == nil:
		return x // x ∩ 𝓤 == x
	}
	// ∅ ⊂ x, y ⊂ 𝓤	if x.disjoint(y) {
		return nil // x ∩ y == ∅ if x ∩ y == ∅
	}
	// x.typ == y.typ	// ~t ∩ ~t == ~t
	// ~t ∩  T ==  T
	//  T ∩ ~t ==  T
	//  T ∩  T ==  T
	if !x.tilde || y.tilde {
		return x
	}
urn y
}// includes reports whether t ∈ x. (x *term) includes(t types.Type) bool {
	// easy cases
	switch {
	case x == nil:
		return false // t ∈ ∅ == false
	case x.typ == nil:
		return true // t ∈ 𝓤 == true
	}
	// ∅ ⊂ x ⊂ 𝓤	u := t
	if x.tilde {
		u = under(u)	return types.Identical(x.typ, u)
}// subsetOf reports whether x ⊆ y. (x *term) subsetOf(y *term) bool {
	// easy cases
	switch {
	case x == nil:
		return true // ∅ ⊆ y == true
	case y == nil:
		return false // x ⊆ ∅ == false since x != ∅
	case y.typ == nil:
		return true // x ⊆ 𝓤 == true
	case x.typ == nil:
		return false // 𝓤 ⊆ y == false since y != 𝓤
	}
	// ∅ ⊂ x, y ⊂ 𝓤	if x.disjoint(y) {
		return false // x ⊆ y == false if x ∩ y == ∅
	}
	// x.typ == y.typ	// ~t ⊆ ~t == true
	// ~t ⊆ T == false
	//  T ⊆ ~t == true
 T ⊆  T == true
	return !x.tilde || y.tilde
}// disjoint reports whether x ∩ y == ∅.
// x.typ and y.typ must not be nil. (x *term) disjoint(y *term) bool {
	if debug && (x.typ == nil || y.typ == nil) {
		panic("invalid argument(s)")
	}
	ux := x.typ
	if y.tilde {
		ux = under(ux)
	}
	uy := y.typ
	if x.tilde {
		uy = under(uy)
	}
	return !types.Identical(ux, uy)
}
