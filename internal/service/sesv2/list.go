// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package sesv2import (
"context""github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/service/sesv2"
)func ListConfigurationSetsPages(ctx context.Context, conn *sesv2.Client, in *sesv2.ListConfigurationSetsInput, fn func(*sesv2.ListConfigurationSetsOutput, bool) bool) error {
for {
out, err := conn.ListConfigurationSets(ctx, in)
 err != nil {
return err
aPage := aws.ToString(out.NextToken) == ""
 !fn(out, lastPage) || lastPage {
break
nextToken = out.NextToken
}return nil
}func ListContactListsPages(ctx context.Context, conn *sesv2.Client, in *sesv2.ListContactListsInput, fn func(*sesv2.ListContactListsOutput, bool) bool) error {
for {
t, err := conn.ListContactLists(ctx, in)
 err != nil {
return err
aPage := aws.ToString(out.NextToken) == ""
 !fn(out, lastPage) || lastPage {
break
nextToken = out.NextToken
}return nil
}
