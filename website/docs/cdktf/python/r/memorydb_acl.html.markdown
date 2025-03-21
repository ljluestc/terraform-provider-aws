---
subcategory: "MemoryDB"
layout: "aws"
page_title: "AWS: aws_memorydb_acl"
description: |-
  Provides a MemoryDB ACL.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_memorydb_acl

Provides a MemoryDB ACL.

More information about users and ACL-s can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/clusters.acls.html).

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.memorydb_acl import MemorydbAcl
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        MemorydbAcl(self, "example",
            name="my-acl",
            user_names=["my-user-1", "my-user-2"]
        )
```

## Argument Reference

The following arguments are optional:

* `name` - (Optional, Forces new resource) Name of the ACL. If omitted, Terraform will assign a random, unique name. Conflicts with `name_prefix`.
* `name_prefix` - (Optional, Forces new resource) Creates a unique name beginning with the specified prefix. Conflicts with `name`.
* `user_names` - (Optional) Set of MemoryDB user names to be included in this ACL.
* `tags` - (Optional) A map of tags to assign to the resource. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Same as `name`.
* `arn` - The ARN of the ACL.
* `minimum_engine_version` - The minimum engine version supported by the ACL.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import an ACL using the `name`. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.memorydb_acl import MemorydbAcl
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        MemorydbAcl.generate_config_for_import(self, "example", "my-acl")
```

Using `terraform import`, import an ACL using the `name`. For example:

```console
% terraform import aws_memorydb_acl.example my-acl
```

<!-- cache-key: cdktf-0.20.8 input-81b0aafca8273c497e315e23df245bdaa7eca67e9dec48c672755900e7eb97ca -->