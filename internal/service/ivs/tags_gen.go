//Codegeneratedbyinternal/generate/tags/main.go;DONOTEDIT.
packageivsimport(
	"context"
	"fmt"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/aws/aws-sdk-go/service/ivs/ivsiface"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/logging"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)//listTagslistsivsservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funclistTags(ctxcontext.Context,connivsiface.IVSAPI,identifierstring)(tftags.KeyValueTags,error){
	input:=&ivs.ListTagsForResourceInput{
		ResourceArn:aws.String(identifier),
	}	output,err:=conn.ListTagsForResourceWithContext(ctx,input)	iferr!=nil{
		returntftags.New(ctx,nil),err
	}	returnKeyValueTags(ctx,output.Tags),nil
}//ListTagslistsivsservicetagsandsettheminContext.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)ListTags(ctxcontext.Context,metaany,identifierstring)error{
	tags,err:=listTags(ctx,meta.(*conns.AWSClient).IVSConn(ctx),identifier)	iferr!=nil{
		returnerr
	}	ifinContext,ok:=tftags.FromContext(ctx);ok{
		inContext.TagsOut=types.Some(tags)
	}	returnnil
}//map[string]*stringhandling//Tagsreturnsivsservicetags.
funcTags(tagstftags.KeyValueTags)map[string]*string{
	returnaws.StringMap(tags.Map())
}//KeyValueTagscreatestftags.KeyValueTagsfromivsservicetags.
funcKeyValueTags(ctxcontext.Context,tagsmap[string]*string)tftags.KeyValueTags{
	returntftags.New(ctx,tags)
}//getTagsInreturnsivsservicetagsfromContext.
//nilisreturnediftherearenoinputtags.
funcgetTagsIn(ctxcontext.Context)map[string]*string{
	ifinContext,ok:=tftags.FromContext(ctx);ok{
		iftags:=Tags(inContext.TagsIn.UnwrapOrDefault());len(tags)>0{
			returntags
		}
	}	returnnil
}//setTagsOutsetsivsservicetagsinContext.
funcsetTagsOut(ctxcontext.Context,tagsmap[string]*string){
	ifinContext,ok:=tftags.FromContext(ctx);ok{
		inContext.TagsOut=types.Some(KeyValueTags(ctx,tags))
	}
}//updateTagsupdatesivsservicetags.
//TheidentifieristypicallytheAmazonResourceName(ARN),although
//itmayalsobeadifferentidentifierdependingontheservice.
funcupdateTags(ctxcontext.Context,connivsiface.IVSAPI,identifierstring,oldTagsMap,newTagsMapany)error{
	oldTags:=tftags.New(ctx,oldTagsMap)
	newTags:=tftags.New(ctx,newTagsMap)	ctx=tflog.SetField(ctx,logging.KeyResourceId,identifier)	removedTags:=oldTags.Removed(newTags)
	removedTags=removedTags.IgnoreSystem(names.IVS)
	iflen(removedTags)>0{
		input:=&ivs.UntagResourceInput{
			ResourceArn:aws.String(identifier),
			TagKeys:aws.StringSlice(removedTags.Keys()),
		}		_,err:=conn.UntagResourceWithContext(ctx,input)		iferr!=nil{
			returnfmt.Errorf("untaggingresource(%s):%w",identifier,err)
		}
	}	updatedTags:=oldTags.Updated(newTags)
	updatedTags=updatedTags.IgnoreSystem(names.IVS)
	iflen(updatedTags)>0{
		input:=&ivs.TagResourceInput{
			ResourceArn:aws.String(identifier),
			Tags:(updatedTags),
		}		_,err:=conn.TagResourceWithContext(ctx,input)		iferr!=nil{
			returnfmt.Errorf("taggingresource(%s):%w",identifier,err)
		}
	}	returnnil
}//UpdateTagsupdatesivsservicetags.
//Itiscalledfromoutsidethispackage.
func(p*servicePackage)UpdateTags(ctxcontext.Context,metaany,identifierstring,oldTags,newTagsany)error{
	returnupdateTags(ctx,meta.(*conns.AWSClient).IVSConn(ctx),identifier,oldTags,newTags)
}
