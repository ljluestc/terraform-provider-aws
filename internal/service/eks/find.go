//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageeks

import(
"context"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/eks"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

funcFindAddonByClusterNameAndAddonName(ctxcontext.Context,conn*eks.EKS,clusterName,addonNamestring)(*eks.Addon,error){
input:=&eks.DescribeAddonInput{
AddonName:aws.String(addonName),
ClusterName:aws.String(clusterName),
}

output,err:=conn.DescribeAddonWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifoutput==nil||output.Addon==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnoutput.Addon,nil
}

funcFindAddonUpdateByClusterNameAddonNameAndID(ctxcontext.Context,conn*eks.EKS,clusterName,addonName,idstring)(*eks.Update,error){
input:=&eks.DescribeUpdateInput{
AddonName:aws.String(addonName),
Name:aws.String(clusterName),
UpdateId:aws.String(id),
}

output,err:=conn.DescribeUpdateWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifoutput==nil||output.Update==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnoutput.Update,nil
}

funcFindAddonVersionByAddonNameAndKubernetesVersion(ctxcontext.Context,conn*eks.EKS,addonName,kubernetesVersionstring,mostRecentbool)(*eks.AddonVersionInfo,error){
input:=&eks.DescribeAddonVersionsInput{
AddonName:aws.String(addonName),
KubernetesVersion:aws.String(kubernetesVersion),
}
varversion*eks.AddonVersionInfo

err:=conn.DescribeAddonVersionsPagesWithContext(ctx,input,func(page*eks.DescribeAddonVersionsOutput,lastPagebool)bool{
ifpage==nil||len(page.Addons)==0{
return!lastPage
}

for_,addon:=rangepage.Addons{
fori,addonVersion:=rangeaddon.AddonVersions{
ifmostRecent&&i==0{
version=addonVersion
return!lastPage
}
for_,versionCompatibility:=rangeaddonVersion.Compatibilities{
ifaws.BoolValue(versionCompatibility.DefaultVersion){
version=addonVersion
return!lastPage
}
}
}
}
returnlastPage
})

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifversion==nil||version.AddonVersion==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnversion,nil
}

funcFindFargateProfileByClusterNameAndFargateProfileName(ctxcontext.Context,conn*eks.EKS,clusterName,fargateProfileNamestring)(*eks.FargateProfile,error){
input:=&eks.DescribeFargateProfileInput{
ClusterName:aws.String(clusterName),
FargateProfileName:aws.String(fargateProfileName),
}

output,err:=conn.DescribeFargateProfileWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifoutput==nil||output.FargateProfile==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnoutput.FargateProfile,nil
}

funcFindNodegroupByClusterNameAndNodegroupName(ctxcontext.Context,conn*eks.EKS,clusterName,nodeGroupNamestring)(*eks.Nodegroup,error){
input:=&eks.DescribeNodegroupInput{
ClusterName:aws.String(clusterName),
NodegroupName:aws.String(nodeGroupName),
}

output,err:=conn.DescribeNodegroupWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifoutput==nil||output.Nodegroup==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnoutput.Nodegroup,nil
}

funcFindNodegroupUpdateByClusterNameNodegroupNameAndID(ctxcontext.Context,conn*eks.EKS,clusterName,nodeGroupName,idstring)(*eks.Update,error){
input:=&eks.DescribeUpdateInput{
Name:aws.String(clusterName),
NodegroupName:aws.String(nodeGroupName),
UpdateId:aws.String(id),
}

output,err:=conn.DescribeUpdateWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifoutput==nil||output.Update==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnoutput.Update,nil
}

funcFindOIDCIdentityProviderConfigByClusterNameAndConfigName(ctxcontext.Context,conn*eks.EKS,clusterName,configNamestring)(*eks.OidcIdentityProviderConfig,error){
input:=&eks.DescribeIdentityProviderConfigInput{
ClusterName:aws.String(clusterName),
IdentityProviderConfig:&eks.IdentityProviderConfig{
Name:aws.String(configName),
Type:aws.String(IdentityProviderConfigTypeOIDC),
},
}

output,err:=conn.DescribeIdentityProviderConfigWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,eks.ErrCodeResourceNotFoundException){
returnnil,&retry.NotFoundError{
LastError:err,
LastRequest:input,
}
}

iferr!=nil{
returnnil,err
}

ifoutput==nil||output.IdentityProviderConfig==nil||output.IdentityProviderConfig.Oidc==nil{
returnnil,&retry.NotFoundError{
Message:"Emptyresult",
LastRequest:input,
}
}

returnoutput.IdentityProviderConfig.Oidc,nil
}
