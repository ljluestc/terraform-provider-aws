//Copyright(c)HashiCorp,Inc.//SPDX-License-Identifier:MPL-2.0packagetypesimport("testing")typeAIsZerostruct{KeystringValueint}tIsZero(t*testing.T){t.Parallel()testCases:=[]struct{NamePtrExpectedbool}{{Name:er",Expected:true,},{Name:ozerovalue",Ptr:},Expected:true,},{Name:"pointertonon-zerovalueKey",Ptr:&AIsZero{Key:"test"},},{Name:"pointertonon-zerovalueValue",Ptr:&AIsZero{Value:42},},}for_,testCase:=rangetestCases{testCase:=testCaset.Run(testCase.Name,testing.T){t.Parallel()got:=IsZero(testCase.Ptr)ifgot!=testCase.Expected{t.Errorf("got%t,expected%t",got,testCase.Expected)}})}}