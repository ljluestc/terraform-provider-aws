//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagedocdb

import(
"context"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)
funcfindGlobalClusterByARN(ctxcontext.Context,conn*docdb.DocDB,dbClusterARNstring)(*docdb.GlobalCluster,error){
varglobalCluster*docdb.GlobalCluster

input:=&docdb.DescribeGlobalClustersInput{
Filters:[]*docdb.Filter{
{
Name:aws.String("db-cluster-id"),
Values:[]*string{aws.String(dbClusterARN)},
},
},
}

err:=conn.DescribeGlobalClustersPagesWithContext(ctx,input,func(page*docdb.DescribeGlobalClustersOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,gc:=rangepage.GlobalClusters{
ifgc==nil{
continue
}

for_,globalClusterMember:=rangegc.GlobalClusterMembers{
ifaws.StringValue(globalClusterMember.DBClusterArn)==dbClusterARN{
globalCluster=gc
returnfalse
}
}
}

return!lastPage
})

returnglobalCluster,err
}
funcfindGlobalClusterIDByARN(ctxcontext.Context,conn*docdb.DocDB,arnstring)string{
result,err:=conn.DescribeDBClustersWithContext(ctx,&docdb.DescribeDBClustersInput{})
iferr!=nil{
return""
}
for_,cluster:=rangeresult.DBClusters{
ifaws.StringValue(cluster.DBClusterArn)==arn{
returnaws.StringValue(cluster.DBClusterIdentifier)
}
}
return""
}
funcFindDBClusterById(ctxcontext.Context,conn*docdb.DocDB,dBClusterIDstring)(*docdb.DBCluster,error){
vardBCluster*docdb.DBCluster

input:=&docdb.DescribeDBClustersInput{
DBClusterIdentifier:aws.String(dBClusterID),
}

err:=conn.DescribeDBClustersPagesWithContext(ctx,input,func(page*docdb.DescribeDBClustersOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,dbc:=rangepage.DBClusters{
ifdbc==nil{
continue
}

ifaws.StringValue(dbc.DBClusterIdentifier)==dBClusterID{
dBCluster=dbc
returnfalse
}
}

return!lastPage
})

returndBCluster,err
}
funcFindDBClusterSnapshotById(ctxcontext.Context,conn*docdb.DocDB,dBClusterSnapshotIDstring)(*docdb.DBClusterSnapshot,error){
vardBClusterSnapshot*docdb.DBClusterSnapshot

input:=&docdb.DescribeDBClusterSnapshotsInput{
DBClusterIdentifier:aws.String(dBClusterSnapshotID),
}

err:=conn.DescribeDBClusterSnapshotsPagesWithContext(ctx,input,func(page*docdb.DescribeDBClusterSnapshotsOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,dbcss:=rangepage.DBClusterSnapshots{
ifdbcss==nil{
continue
}

ifaws.StringValue(dbcss.DBClusterIdentifier)==dBClusterSnapshotID{
dBClusterSnapshot=dbcss
returnfalse
}
}

return!lastPage
})

returndBClusterSnapshot,err
}
funcFindDBInstanceById(ctxcontext.Context,conn*docdb.DocDB,dBInstanceIDstring)(*docdb.DBInstance,error){
vardBInstance*docdb.DBInstance

input:=&docdb.DescribeDBInstancesInput{
DBInstanceIdentifier:aws.String(dBInstanceID),
}

err:=conn.DescribeDBInstancesPagesWithContext(ctx,input,func(page*docdb.DescribeDBInstancesOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,dbi:=rangepage.DBInstances{
ifdbi==nil{
continue
}

ifaws.StringValue(dbi.DBInstanceIdentifier)==dBInstanceID{
dBInstance=dbi
returnfalse
}
}

return!lastPage
})

returndBInstance,err
}
funcFindGlobalClusterById(ctxcontext.Context,conn*docdb.DocDB,globalClusterIDstring)(*docdb.GlobalCluster,error){
varglobalCluster*docdb.GlobalCluster

input:=&docdb.DescribeGlobalClustersInput{
GlobalClusterIdentifier:aws.String(globalClusterID),
}

err:=conn.DescribeGlobalClustersPagesWithContext(ctx,input,func(page*docdb.DescribeGlobalClustersOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,gc:=rangepage.GlobalClusters{
ifgc==nil{
continue
}

ifaws.StringValue(gc.GlobalClusterIdentifier)==globalClusterID{
globalCluster=gc
returnfalse
}
}

return!lastPage
})

returnglobalCluster,err
}
funcFindDBSubnetGroupByName(ctxcontext.Context,conn*docdb.DocDB,dBSubnetGroupNamestring)(*docdb.DBSubnetGroup,error){
vardBSubnetGroup*docdb.DBSubnetGroup

input:=&docdb.DescribeDBSubnetGroupsInput{
DBSubnetGroupName:aws.String(dBSubnetGroupName),
}

err:=conn.DescribeDBSubnetGroupsPagesWithContext(ctx,input,func(page*docdb.DescribeDBSubnetGroupsOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,sg:=rangepage.DBSubnetGroups{
ifsg==nil{
continue
}

ifaws.StringValue(sg.DBSubnetGroupName)==dBSubnetGroupName{
dBSubnetGroup=sg
returnfalse
}
}

return!lastPage
})

returndBSubnetGroup,err
}
funcFindEventSubscriptionByID(ctxcontext.Context,conn*docdb.DocDB,idstring)(*docdb.EventSubscription,error){
vareventSubscription*docdb.EventSubscription

input:=&docdb.DescribeEventSubscriptionsInput{
SubscriptionName:aws.String(id),
}

err:=conn.DescribeEventSubscriptionsPagesWithContext(ctx,input,func(page*docdb.DescribeEventSubscriptionsOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

for_,es:=rangepage.EventSubscriptionsList{
ifes==nil{
continue
}

ifaws.StringValue(es.CustSubscriptionId)==id{
eventSubscription=es
returnfalse
}
}

return!lastPage
})

iftfawserr.ErrCodeEquals(err,docdb.ErrCodeSubscriptionNotFoundFault){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifeventSubscription==nil{
returnnil,tfresource.NewEmptyResultError(input)
}

returneventSubscription,nil
}
