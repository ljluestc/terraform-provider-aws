//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagetfresource_test

import(
"errors"
"fmt"
"strings"
"testing"

"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

funcTestNotFound(t*testing.T){
t.Parallel()

testCases:=[]struct{
Namestring
Errerror
Expectedbool
}{
{
Name:"nilerror",
Err:nil,
},
{
Name:"othererror",
Err:errors.New("test"),
},
{
Name:"notfounderror",
Err:&retry.NotFoundError{LastError:errors.New("test")},
Expected:true,
},
{
Name:"wrappedothererror",
Err:fmt.Errorf("test:%w",errors.New("test")),
},
{
Name:"wrappednotfounderror",
Err:fmt.Errorf("test:%w",&retry.NotFoundError{LastError:errors.New("test")}),
Expected:true,
},
}

for_,testCase:=rangetestCases{
testCase:=testCase
t.Run(testCase.Name,func(t*testing.T){
t.Parallel()

got:=tfresource.NotFound(testCase.Err)

ifgot!=testCase.Expected{
t.Errorf("got%t,expected%t",got,testCase.Expected)
}
})
}
}

funcTestTimedOut(t*testing.T){
t.Parallel()

testCases:=[]struct{
Namestring
Errerror
Expectedbool
}{
{
Name:"nilerror",
Err:nil,
},
{
Name:"othererror",
Err:errors.New("test"),
},
{
Name:"timeouterror",
Err:&retry.TimeoutError{},
Expected:true,
},
{
Name:"timeouterrornon-nillasterror",
Err:&retry.TimeoutError{LastError:errors.New("test")},
},
{
Name:"wrappedothererror",
Err:fmt.Errorf("test:%w",errors.New("test")),
},
{
Name:"wrappedtimeouterror",
Err:fmt.Errorf("test:%w",&retry.TimeoutError{}),
},
{
Name:"wrappedtimeouterrornon-nillasterror",
Err:fmt.Errorf("test:%w",&retry.TimeoutError{LastError:errors.New("test")}),
},
}

for_,testCase:=rangetestCases{
testCase:=testCase
t.Run(testCase.Name,func(t*testing.T){
t.Parallel()

got:=tfresource.TimedOut(testCase.Err)

ifgot!=testCase.Expected{
t.Errorf("got%t,expected%t",got,testCase.Expected)
}
})
}
}

funcTestSetLastError(t*testing.T){
t.Parallel()

testCases:=[]struct{
Namestring
Errerror
LastErrerror
Expectedbool
}{
{
Name:"nilerror",
},
{
Name:"othererror",
Err:errors.New("test"),
LastErr:errors.New("last"),
},
{
Name:"timeouterrorlastErrisnil",
Err:&retry.TimeoutError{},
},
{
Name:"timeouterror",
Err:&retry.TimeoutError{},
LastErr:errors.New("lasttest"),
Expected:true,
},
{
Name:"timeouterrornon-nillasterrorlastErrisnil",
Err:&retry.TimeoutError{LastError:errors.New("test")},
},
{
Name:"timeouterrornon-nillasterrornooverwrite",
Err:&retry.TimeoutError{LastError:errors.New("test")},
LastErr:errors.New("lasttest"),
},
{
Name:"unexpectedstateerrorlastErrisnil",
Err:&retry.UnexpectedStateError{},
},
{
Name:"unexpectedstateerror",
Err:&retry.UnexpectedStateError{},
LastErr:errors.New("lasttest"),
Expected:true,
},
{
Name:"unexpectedstateerrornon-nillasterrorlastErrisnil",
Err:&retry.UnexpectedStateError{LastError:errors.New("test")},
},
{
Name:"unexpectedstateerrornon-nillasterrornooverwrite",
Err:&retry.UnexpectedStateError{LastError:errors.New("test")},
LastErr:errors.New("lasttest"),
},
}

for_,testCase:=rangetestCases{
testCase:=testCase
t.Run(testCase.Name,func(t*testing.T){
t.Parallel()

tfresource.SetLastError(testCase.Err,testCase.LastErr)

iftestCase.Err!=nil{
got:=testCase.Err.Error()
contains:=strings.Contains(got,"lasttest")

if(testCase.Expected&&!contains)||(!testCase.Expected&&contains){
t.Errorf("got%s",got)
}
}
})
}
}
