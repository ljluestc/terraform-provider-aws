---
subcategory: "ELB (Elastic Load Balancing)"
layout: "aws"
page_title: "AWS: aws_lb_cookie_stickiness_policy"
description: |-
  Provides a load balancer cookie stickiness policy.
---

# Resource: aws_lb_cookie_stickiness_policy

Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.

## Example Usage

```hcl
resource "aws_elb" "lb" {
  name               = "test-elb"
  availability_zones = ["us-east-1a"]
  listener {
    instance_port     = 8088
    instance_protocol = "http"
    lb_port           = 8088
    lb_protocol       = "http"
  }
}

resource "aws_lb_cookie_stickiness_policy" "foo" {
  name                     = "test-policy"
  load_balancer            = aws_elb.lb.name
  lb_port                  = 8088
  cookie_expiration_period = 0
}
```

### Example: Unset cookie_expiration_period (disables stickiness)

```hcl
resource "aws_lb_cookie_stickiness_policy" "foo" {
  name          = "test-policy"
  load_balancer = aws_elb.lb.name
  lb_port       = 8088
  # cookie_expiration_period intentionally omitted
}
```

## Argument Reference

* `name` - (Required) The name of the stickiness policy.
* `load_balancer` - (Required) The name of the load balancer.
* `lb_port` - (Required) The load balancer port to which the policy should be applied.
* `cookie_expiration_period` - (Optional) The time period in seconds after which the session cookie is considered stale.
  - If set to `0`, the cookie expires when the browser is closed (browser-session cookie).
  - If set to a positive integer, the cookie expires after that many seconds.
  - If omitted, stickiness is disabled.
  - If set to `0`, the cookie expires when the browser is closed (browser-session cookie).
  - If set to a positive integer, the cookie expires after that many seconds.
  - If omitted, stickiness is disabled.
  - If set to `0`, the cookie expires when the browser is closed (browser-session cookie).
  - If set to a positive integer, the cookie expires after that many seconds.
  - If omitted, stickiness is disabled.

## Attributes Reference

* `id` - The ID of the policy, in the format `load_balancer:port:policy_name`.