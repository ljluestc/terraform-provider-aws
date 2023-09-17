// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package ecsimport (
"testing"
)func TestValidPlacementConstraint(t *testing.T) {
t.Parallel()cases := []struct {
constType string
constExpr string
Errbool
}{
{
constType: "distinctInstance",
constExpr: "",
Err:false,
},
{
constType: "memberOf",
constExpr: "",
Err:true,
},
{
constType: "distinctInstance",
constExpr: "expression",
Err:false,
},
{
constType: "memberOf",
constExpr: "expression",
Err:false,
},
}for _, tc := range cases {
if err := validPlacementConstraint(tc.constType, tc.constExpr); err != nil && !tc.Err {
t.Fatalf("Unexpected validation error for \"%s:%s\": %s",
tc.constType, tc.constExpr, err)
}
}
}func TestValidPlacementStrategy(t *testing.T) {
t.Parallel()cases := []struct {
stratType  string
stratField string
Err
}{
{
stratType:  "random",
stratField: "",
Err:
},
{
stratType:  "spread",
stratField: "instanceID",
Err:
},
{
stratType:  "binpack",
stratField: "cpu",
Err:
},
{
stratType:  "binpack",
stratField: "memory",
Err:
},
{
stratType:  "binpack",
stratField: "disk",
Err:
},
{
stratType:  "fakeType",
stratField: "",
Err:
},
}for _, tc := range cases {
if err := validPlacementStrategy(tc.stratType, tc.stratField); err != nil && !tc.Err {
t.Fatalf("Unexpected validation error for \"%s:%s\": %s",
tc.stratType, tc.stratField, err)
}
}
}
