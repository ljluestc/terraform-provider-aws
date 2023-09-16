// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package elasticacheconst (
	engineMemcached = "memcached"
	engineRedis= "redis"
)// engine_Values returns all elements of the Engine enum
func engine_Values() []string {
	return []string{
		engineMemcached,
		engineRedis,
	}
}
