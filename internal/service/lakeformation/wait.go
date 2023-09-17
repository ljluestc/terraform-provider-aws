//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagelakeformation

import(
"context"
"time"

"github.com/aws/aws-sdk-go/service/lakeformation"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const(
permissionsReadyTimeout=1*time.Minute
permissionsDeleteRetryTimeout=30*time.Second

statusAvailable="AVAILABLE"
statusNotFound="NOTFOUND"
statusFailed="FAILED"
statusIAMDelay="IAMDELAY"
)

funcwaitPermissionsReady(ctxcontext.Context,conn*lakeformation.LakeFormation,input*lakeformation.ListPermissionsInput,tableTypestring,columnNames[]*string,excludedColumnNames[]*string,columnWildcardbool)([]*lakeformation.PrincipalResourcePermissions,error){
stateConf:=&retry.StateChangeConf{
Pending:[]string{statusNotFound,statusIAMDelay},
Target:[]string{statusAvailable},
Refresh:statusPermissions(ctx,conn,input,tableType,columnNames,excludedColumnNames,columnWildcard),
Timeout:permissionsReadyTimeout,
}

outputRaw,err:=stateConf.WaitForStateContext(ctx)

ifoutput,ok:=outputRaw.([]*lakeformation.PrincipalResourcePermissions);ok{
returnoutput,err
}

returnnil,err
}
