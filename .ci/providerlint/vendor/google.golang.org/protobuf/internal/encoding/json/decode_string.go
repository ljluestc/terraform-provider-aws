//Copyright2018TheGoAuthors.Allrightsreserved.//UseofthissourcecodeisgovernedbyaBSD-style//licensethatcanbefoundintheLICENSEfile.packagejsonimport("strconv""unicode""unicode/utf16""unicode/utf8""google.golang.org/protobuf/internal/strs")(d*Decoder)parseString(in[]byte)(string,int,error){in0:=iniflen(in)==0{return"",0,ErrUnexpectedEOF}ifin[0]!='"'{return"",0,d.newSyntaxError(d.currPos(),"invalidcharacter%qatstartofstring",in[0])}in=in[1:]i:=indexNeedEscapeInBytes(in)in,out:=in[i:],in[:i:i]//setcaptopreventmutationsforlen(in)>0{switchr,n:=utf8.DecodeRune(in);{caser==utf8.RuneError&&n==1:return"",0,d.newSyntaxError(d.currPos(),"invalidUTF-8instring")caser<'':return"",0,d.newSyntaxError(d.currPos(),"invalidcharacter%qinstring",r)caser=='"':in=in[1:]n:=len(in0)-len(in)returnstring(out),n,nilcaser=='\\':iflen(in)<2{return"",0,ErrUnexpectedEOF}switchr:=in[1];r{case'"','\\','/':in,out=in[2:],append(out,r)case'b':in,out=in[2:],append(out,'\b')case'f':in,out=in[2:],append(out,'\f')case'n':in,out=in[2:],append(out,'\n')case'r':in,out=in[2:],append(out,'\r')case't':in,out=in[2:],append(out,'\t')case'u':iflen(in)<6{return"",0,ErrUnexpectedEOF}v,err:=strconv.ParseUint(string(in[2:6]),16,16)iferr!=nil{return"",0,d.newSyntaxError(d.currPos(),"invalidescapecode%qinstring",in[:6])}in=in[6:]r:=rune(v)ifutf16.IsSurrogate(r){iflen(in)<6{return"",0,ErrUnexpectedEOF}v,err:=strconv.ParseUint(string(in[2:6]),16,16)r=utf16.DecodeRune(r,rune(v))ifin[0]!='\\'||in[1]!='u'||r==unicode.ReplacementChar||err!=nil{return"",0,d.newSyntaxError(d.currPos(),"invalidescapecode%qinstring",in[:6])}in=in[6:]}out=append(out,string(r)...)default:return"",0,d.newSyntaxError(d.currPos(),"invalidescapecode%qinstring",in[:2])}default:i:=indexNeedEscapeInBytes(in[n:])in,out=in[n+i:],append(out,in[:n+i]...)}}return"",0,ErrUnexpectedEOF}//indexNeedEscapeInBytesreturnstheindexofthecharacterthatneedsscaping.Ifnocharactersneedescaping,thisreturnstheinputlength.indexNeedEscapeInBytes(b[]byte)int{returnindexNeedEscapeInString(strs.UnsafeString(b))}