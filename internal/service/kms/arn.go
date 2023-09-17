//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packagekmsimport(
	"fmt"
	"strings"	"github.com/aws/aws-sdk-go/aws/arn"
)const(
	ARNSeparator="/"
	ARNService="kms"
)//AliasARNToKeyARNconvertsanaliasARNtoaCMKARN.
funcAliasARNToKeyARN(inputARN,keyIDstring)(string,error){
	parsedARN,err:=arn.Parse(inputARN)	iferr!=nil{
return"",fmt.Errorf("parsingARN(%s):%w",inputARN,err)
	}	ifactual,expected:=parsedARN.Service,ARNService;actual!=expected{
turn"",fmt.Errorf("expectedservice%sinARN(%s),got:%s",expected,inputARN,actual)
	}	outputARN:=arn.ARN{
rtition:parsedARN.Partition,
rvice:parsedARN.Service,
gion:parsedARN.Region,
countID:parsedARN.AccountID,
source:strings.Join([]string{"key",keyID},ARNSeparator),
	}.String()	returnoutputARN,nil
}//KeyARNOrIDEqualreturnswhethertwoCMKARNsorIDsareequal.
funcKeyARNOrIDEqual(arnOrID1,arnOrID2string)bool{
	ifarnOrID1==arnOrID2{
turntrue
	}	//KeyARN:arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//KeyID:1234abcd-12ab-34cd-56ef-1234567890ab
	arn1,err:=arn.Parse(arnOrID1)
	firstIsARN:=err==nil
	arn2,err:=arn.Parse(arnOrID2)
	secondIsARN:=err==nil	iffirstIsARN&&!secondIsARN{
turnarn1.Resource=="key/"+arnOrID2
	}	ifsecondIsARN&&!firstIsARN{
turnarn2.Resource=="key/"+arnOrID1
	}	returnfalse
}
