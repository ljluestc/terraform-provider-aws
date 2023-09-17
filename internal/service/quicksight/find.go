// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package quicksightimport (
"context""github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/quicksight"
)func FindGroupMembership(ctx context.Context, conn *quicksight.QuickSight, listInput *quicksight.ListGroupMembershipsInput, userName string) (bool, error) {
found := falsefor {
resp, err := conn.ListGroupMembershipsWithContext(ctx, listInput)
 err != nil {
return false, err
o_, member := range resp.GroupMemberList {
if aws.StringValue(member.MemberName) == userName {
found = true
break
}
found || resp.NextToken == nil {
break
iInput.NextToken = resp.NextToken
}return found, nil
}
