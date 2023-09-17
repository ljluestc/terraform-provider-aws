//Codegeneratedbyinternal/generate/tags/main.go;DONOTEDIT.
packagecodeartifact

import(
"context"
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/codeartifact"
"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
"github.com/hashicorp/terraform-plugin-log/tflog"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/logging"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/types"
"github.com/hashicorp/terraform-provider-aws/names"
)

//listTagslistscodeartifactservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funclistTags(ctxcontext.Context,conncodeartifactiface.CodeArtifactAPI,identifierstring)(tftags.KeyValueTags,error){
input:=&codeartifact.ListTagsForResourceInput{
ResourceArn:aws.String(identifier),
}

output,err:=conn.ListTagsForResourceWithContext(ctx,input)

iferr!=nil{
returntftags.New(ctx,nil),err
}

returnKeyValueTags(ctx,output.Tags),nil
}

//ListTagslistscodeartifactservicetagsandsettheminContext.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)ListTags(ctxcontext.Context,metaany,identifierstring)error{
tags,err:=listTags(ctx,meta.(*conns.AWSClient).CodeArtifactConn(ctx),identifier)

iferr!=nil{
returnerr
}

ifinContext,ok:=tftags.FromContext(ctx);ok{
inContext.TagsOut=types.Some(tags)
}

returnnil
}

//[]*SERVICE.Taghandling

//Tagsreturnscodeartifactservicetags.
funcTags(tagstftags.KeyValueTags)[]*codeartifact.Tag{
result:=make([]*codeartifact.Tag,0,len(tags))

fork,v:=rangetags.Map(){
tag:=&codeartifact.Tag{
Key:aws.String(k),
Value:aws.String(v),
}

result=append(result,tag)
}

returnresult
}

//KeyValueTagscreatestftags.KeyValueTagsfromcodeartifactservicetags.
funcKeyValueTags(ctxcontext.Context,tags[]*codeartifact.Tag)tftags.KeyValueTags{
m:=make(map[string]*string,len(tags))

for_,tag:=rangetags{
m[aws.StringValue(tag.Key)]=tag.Value
}

returntftags.New(ctx,m)
}

//getTagsInreturnscodeartifactservicetagsfromContext.
//nilisreturnediftherearenoinputtags.
funcgetTagsIn(ctxcontext.Context)[]*codeartifact.Tag{
ifinContext,ok:=tftags.FromContext(ctx);ok{
iftags:=Tags(inContext.TagsIn.UnwrapOrDefault());len(tags)>0{
returntags
}
}

returnnil
}

//setTagsOutsetscodeartifactservicetagsinContext.
funcsetTagsOut(ctxcontext.Context,tags[]*codeartifact.Tag){
ifinContext,ok:=tftags.FromContext(ctx);ok{
inContext.TagsOut=types.Some(KeyValueTags(ctx,tags))
}
}

//updateTagsupdatescodeartifactservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funcupdateTags(ctxcontext.Context,conncodeartifactiface.CodeArtifactAPI,identifierstring,oldTagsMap,newTagsMapany)error{
oldTags:=tftags.New(ctx,oldTagsMap)
newTags:=tftags.New(ctx,newTagsMap)

ctx=tflog.SetField(ctx,logging.KeyResourceId,identifier)

removedTags:=oldTags.Removed(newTags)
removedTags=removedTags.IgnoreSystem(names.CodeArtifact)
iflen(removedTags)>0{
input:=&codeartifact.UntagResourceInput{
ResourceArn:aws.String(identifier),
TagKeys:aws.StringSlice(removedTags.Keys()),
}

_,err:=conn.UntagResourceWithContext(ctx,input)

iferr!=nil{
returnfmt.Errorf("untaggingresource(%s):%w",identifier,err)
}
}

updatedTags:=oldTags.Updated(newTags)
updatedTags=updatedTags.IgnoreSystem(names.CodeArtifact)
iflen(updatedTags)>0{
input:=&codeartifact.TagResourceInput{
ResourceArn:aws.String(identifier),
Tags:(updatedTags),
}

_,err:=conn.TagResourceWithContext(ctx,input)

iferr!=nil{
returnfmt.Errorf("taggingresource(%s):%w",identifier,err)
}
}

returnnil
}

//UpdateTagsupdatescodeartifactservicetags.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)UpdateTags(ctxcontext.Context,metaany,identifierstring,oldTags,newTagsany)error{
returnupdateTags(ctx,meta.(*conns.AWSClient).CodeArtifactConn(ctx),identifier,oldTags,newTags)
}
