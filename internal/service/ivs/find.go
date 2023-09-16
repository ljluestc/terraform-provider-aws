//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageivs

import(
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)
funcFindPlaybackKeyPairByID(ctxcontext.Context,conn*ivs.IVS,idstring)(*ivs.PlaybackKeyPair,error){
	in:=&ivs.GetPlaybackKeyPairInput{
		Arn:aws.String(id),
	}
	out,err:=conn.GetPlaybackKeyPairWithContext(ctx,in)
	iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:in,
		}
	}

	iferr!=nil{
		returnnil,err
	}

	ifout==nil||out.KeyPair==nil{
		returnnil,tfresource.NewEmptyResultError(in)
	}

	returnout.KeyPair,nil
}
funcFindRecordingConfigurationByID(ctxcontext.Context,conn*ivs.IVS,idstring)(*ivs.RecordingConfiguration,error){
	in:=&ivs.GetRecordingConfigurationInput{
		Arn:aws.String(id),
	}
	out,err:=conn.GetRecordingConfigurationWithContext(ctx,in)
	iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:in,
		}
	}

	iferr!=nil{
		returnnil,err
	}

	ifout==nil||out.RecordingConfiguration==nil{
		returnnil,tfresource.NewEmptyResultError(in)
	}

	returnout.RecordingConfiguration,nil
}
funcFindChannelByID(ctxcontext.Context,conn*ivs.IVS,arnstring)(*ivs.Channel,error){
	in:=&ivs.GetChannelInput{
		Arn:aws.String(arn),
	}
	out,err:=conn.GetChannelWithContext(ctx,in)
	iferr!=nil{
		iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
			returnnil,&retry.NotFoundError{
				LastError:err,
				LastRequest:in,
			}
		}

		returnnil,err
	}

	ifout==nil||out.Channel==nil{
		returnnil,tfresource.NewEmptyResultError(in)
	}

	returnout.Channel,nil
}
funcFindStreamKeyByChannelID(ctxcontext.Context,conn*ivs.IVS,channelArnstring)(*ivs.StreamKey,error){
	in:=&ivs.ListStreamKeysInput{
		ChannelArn:aws.String(channelArn),
	}
	out,err:=conn.ListStreamKeysWithContext(ctx,in)
	iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:in,
		}
	}

	iferr!=nil{
		returnnil,err
	}

	iflen(out.StreamKeys)<1{
		returnnil,&retry.NotFoundError{
			LastRequest:in,
		}
	}

	streamKeyArn:=out.StreamKeys[0].Arn

	returnfindStreamKeyByID(ctx,conn,*streamKeyArn)
}
funcfindStreamKeyByID(ctxcontext.Context,conn*ivs.IVS,idstring)(*ivs.StreamKey,error){
	in:=&ivs.GetStreamKeyInput{
		Arn:aws.String(id),
	}
	out,err:=conn.GetStreamKeyWithContext(ctx,in)
	iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:in,
		}
	}

	iferr!=nil{
		returnnil,err
	}

	returnout.StreamKey,nil
}
