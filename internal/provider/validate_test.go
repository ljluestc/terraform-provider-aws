//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageprovider

import(
"regexp"
"testing"

"github.com/YakDriver/regexache"
)

funcTestValidAssumeRoleDuration(t*testing.T){
t.Parallel()

testCases:=[]struct{
valinterface{}
expectedErr*regexp.Regexp
}{
{
val:"",
expectedErr:regexache.MustCompile(`cannotbeparsedasaduration`),
},
{
val:"1",
expectedErr:regexache.MustCompile(`cannotbeparsedasaduration`),
},
{
val:"10m",
expectedErr:regexache.MustCompile(`mustbebetween15minutes\(15m\)and12hours\(12h\)`),
},
{
val:"12h30m",
expectedErr:regexache.MustCompile(`mustbebetween15minutes\(15m\)and12hours\(12h\)`),
},
{

val:"15m",
},
{
val:"1h10m10s",
},
{

val:"12h",
},
}

matchErr:=func(errs[]error,r*regexp.Regexp)bool{
//errmustmatchoneprovided
for_,err:=rangeerrs{
ifr.MatchString(err.Error()){
returntrue
}
}

returnfalse
}

fori,tc:=rangetestCases{
_,errs:=validAssumeRoleDuration(tc.val,"test_property")

iflen(errs)==0&&tc.expectedErr==nil{
continue
}

iflen(errs)!=0&&tc.expectedErr==nil{
t.Fatalf("expectedtestcase%dtoproducenoerrors,got%v",i,errs)
}

if!matchErr(errs,tc.expectedErr){
t.Fatalf("expectedtestcase%dtoproduceerrormatching\"%s\",got%v",i,tc.expectedErr,errs)
}
}
}
