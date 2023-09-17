//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageeks

import(
"context"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/eks"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

//@SDKDataSource("aws_eks_node_groups")
funcDataSourceNodeGroups()*schema.Resource{
return&schema.Resource{
ReadWithoutTimeout:dataSourceNodeGroupsRead,

Schema:map[string]*schema.Schema{
"cluster_name":{
Type:schema.TypeString,
Required:true,
ValidateFunc:validation.NoZeroValues,
},
"names":{
Type:schema.TypeSet,
Computed:true,
Elem:&schema.Schema{Type:schema.TypeString},
},
},
}
}

funcdataSourceNodeGroupsRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).EKSConn(ctx)

clusterName:=d.Get("cluster_name").(string)

input:=&eks.ListNodegroupsInput{
ClusterName:aws.String(clusterName),
}

varnodegroups[]*string

err:=conn.ListNodegroupsPagesWithContext(ctx,input,func(page*eks.ListNodegroupsOutput,lastPagebool)bool{
ifpage==nil{
return!lastPage
}

nodegroups=append(nodegroups,page.Nodegroups...)

return!lastPage
})

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"listingEKSNodeGroups:%s",err)
}

d.SetId(clusterName)

d.Set("cluster_name",clusterName)
d.Set("names",aws.StringValueSlice(nodegroups))

returndiags
}
