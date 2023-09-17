//Codegeneratedbyinternal/generate/tags/main.go;DONOTEDIT.
packagedocdb

import(
"context"
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/aws/aws-sdk-go/service/docdb/docdbiface"
"github.com/hashicorp/terraform-plugin-log/tflog"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/logging"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/types"
"github.com/hashicorp/terraform-provider-aws/names"
)

//listTagslistsdocdbservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funclistTags(ctxcontext.Context,conndocdbiface.DocDBAPI,identifierstring)(tftags.KeyValueTags,error){
input:=&docdb.ListTagsForResourceInput{
ResourceName:aws.String(identifier),
}

output,err:=conn.ListTagsForResourceWithContext(ctx,input)

iferr!=nil{
returntftags.New(ctx,nil),err
}

returnKeyValueTags(ctx,output.TagList),nil
}

//ListTagslistsdocdbservicetagsandsettheminContext.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)ListTags(ctxcontext.Context,metaany,identifierstring)error{
tags,err:=listTags(ctx,meta.(*conns.AWSClient).DocDBConn(ctx),identifier)

iferr!=nil{
returnerr
}

ifinContext,ok:=tftags.FromContext(ctx);ok{
inContext.TagsOut=types.Some(tags)
}

returnnil
}

//[]*SERVICE.Taghandling

//Tagsreturnsdocdbservicetags.
funcTags(tagstftags.KeyValueTags)[]*docdb.Tag{
result:=make([]*docdb.Tag,0,len(tags))

fork,v:=rangetags.Map(){
tag:=&docdb.Tag{
Key:aws.String(k),
Value:aws.String(v),
}

result=append(result,tag)
}

returnresult
}

//KeyValueTagscreatestftags.KeyValueTagsfromdocdbservicetags.
funcKeyValueTags(ctxcontext.Context,tags[]*docdb.Tag)tftags.KeyValueTags{
m:=make(map[string]*string,len(tags))

for_,tag:=rangetags{
m[aws.StringValue(tag.Key)]=tag.Value
}

returntftags.New(ctx,m)
}

//getTagsInreturnsdocdbservicetagsfromContext.
//nilisreturnediftherearenoinputtags.
funcgetTagsIn(ctxcontext.Context)[]*docdb.Tag{
ifinContext,ok:=tftags.FromContext(ctx);ok{
iftags:=Tags(inContext.TagsIn.UnwrapOrDefault());len(tags)>0{
returntags
}
}

returnnil
}

//setTagsOutsetsdocdbservicetagsinContext.
funcsetTagsOut(ctxcontext.Context,tags[]*docdb.Tag){
ifinContext,ok:=tftags.FromContext(ctx);ok{
inContext.TagsOut=types.Some(KeyValueTags(ctx,tags))
}
}

//updateTagsupdatesdocdbservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funcupdateTags(ctxcontext.Context,conndocdbiface.DocDBAPI,identifierstring,oldTagsMap,newTagsMapany)error{
oldTags:=tftags.New(ctx,oldTagsMap)
newTags:=tftags.New(ctx,newTagsMap)

ctx=tflog.SetField(ctx,logging.KeyResourceId,identifier)

removedTags:=oldTags.Removed(newTags)
removedTags=removedTags.IgnoreSystem(names.DocDB)
iflen(removedTags)>0{
input:=&docdb.RemoveTagsFromResourceInput{
ResourceName:aws.String(identifier),
TagKeys:ringSlice(removedTags.Keys()),
}

_,err:=conn.RemoveTagsFromResourceWithContext(ctx,input)

iferr!=nil{
returnfmt.Errorf("untaggingresource(%s):%w",identifier,err)
}
}

updatedTags:=oldTags.Updated(newTags)
updatedTags=updatedTags.IgnoreSystem(names.DocDB)
iflen(updatedTags)>0{
input:=&docdb.AddTagsToResourceInput{
ResourceName:aws.String(identifier),
Tags:s(updatedTags),
}

_,err:=conn.AddTagsToResourceWithContext(ctx,input)

iferr!=nil{
returnfmt.Errorf("taggingresource(%s):%w",identifier,err)
}
}

returnnil
}

//UpdateTagsupdatesdocdbservicetags.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)UpdateTags(ctxcontext.Context,metaany,identifierstring,oldTags,newTagsany)error{
returnupdateTags(ctx,meta.(*conns.AWSClient).DocDBConn(ctx),identifier,oldTags,newTags)
}
