//Codegeneratedbyinternal/generate/tags/main.go;DONOTEDIT.
packagefsximport(
	"context"
	"fmt"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/fsx/fsxiface"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/logging"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)//listTagslistsfsxservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funclistTags(ctxcontext.Context,connfsxiface.FSxAPI,identifierstring)(tftags.KeyValueTags,error){
	input:=&fsx.ListTagsForResourceInput{
ResourceARN:aws.String(identifier),
	}	output,err:=conn.ListTagsForResourceWithContext(ctx,input)	iferr!=nil{
returntftags.New(ctx,nil),err
	}	returnKeyValueTags(ctx,output.Tags),nil
}//ListTagslistsfsxservicetagsandsettheminContext.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)ListTags(ctxcontext.Context,metaany,identifierstring)error{
	tags,err:=listTags(ctx,meta.(*conns.AWSClient).FSxConn(ctx),identifier)	iferr!=nil{
returnerr
	}	ifinContext,ok:=tftags.FromContext(ctx);ok{
inContext.TagsOut=types.Some(tags)
	}	returnnil
}//[]*SERVICE.Taghandling//Tagsreturnsfsxservicetags.
funcTags(tagstftags.KeyValueTags)[]*fsx.Tag{
	result:=make([]*fsx.Tag,0,len(tags))	fork,v:=rangetags.Map(){
tag:=&fsx.Tag{
Key:aws.String(k),
Value:aws.String(v),
}result=append(result,tag)
	}	returnresult
}//KeyValueTagscreatestftags.KeyValueTagsfromfsxservicetags.
funcKeyValueTags(ctxcontext.Context,tags[]*fsx.Tag)tftags.KeyValueTags{
	m:=make(map[string]*string,len(tags))	for_,tag:=rangetags{
m[aws.StringValue(tag.Key)]=tag.Value
	}	returntftags.New(ctx,m)
}//getTagsInreturnsfsxservicetagsfromContext.
//nilisreturnediftherearenoinputtags.
funcgetTagsIn(ctxcontext.Context)[]*fsx.Tag{
	ifinContext,ok:=tftags.FromContext(ctx);ok{
iftags:=Tags(inContext.TagsIn.UnwrapOrDefault());len(tags)>0{
returntags
}
	}	returnnil
}//setTagsOutsetsfsxservicetagsinContext.
funcsetTagsOut(ctxcontext.Context,tags[]*fsx.Tag){
	ifinContext,ok:=tftags.FromContext(ctx);ok{
inContext.TagsOut=types.Some(KeyValueTags(ctx,tags))
	}
}//updateTagsupdatesfsxservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funcupdateTags(ctxcontext.Context,connfsxiface.FSxAPI,identifierstring,oldTagsMap,newTagsMapany)error{
	oldTags:=tftags.New(ctx,oldTagsMap)
	newTags:=tftags.New(ctx,newTagsMap)	ctx=tflog.SetField(ctx,logging.KeyResourceId,identifier)	removedTags:=oldTags.Removed(newTags)
	removedTags=removedTags.IgnoreSystem(names.FSx)
	iflen(removedTags)>0{
input:=&fsx.UntagResourceInput{
ResourceARN:aws.String(identifier),
TagKeys:aws.StringSlice(removedTags.Keys()),
}_,err:=conn.UntagResourceWithContext(ctx,input)iferr!=nil{
returnfmt.Errorf("untaggingresource(%s):%w",identifier,err)
}
	}	updatedTags:=oldTags.Updated(newTags)
	updatedTags=updatedTags.IgnoreSystem(names.FSx)
	iflen(updatedTags)>0{
input:=&fsx.TagResourceInput{
ResourceARN:aws.String(identifier),
Tags:Tags(updatedTags),
}_,err:=conn.TagResourceWithContext(ctx,input)iferr!=nil{
returnfmt.Errorf("taggingresource(%s):%w",identifier,err)
}
	}	returnnil
}//UpdateTagsupdatesfsxservicetags.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)UpdateTags(ctxcontext.Context,metaany,identifierstring,oldTags,newTagsany)error{
	returnupdateTags(ctx,meta.(*conns.AWSClient).FSxConn(ctx),identifier,oldTags,newTags)
}
