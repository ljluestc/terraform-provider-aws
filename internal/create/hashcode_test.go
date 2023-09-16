// Copyright (c) HashiCorp, Inc.// SPDX-License-Identifier: MPL-2.0package createimport ("testing")

func TestStringHashcode(t *testing.T) {t.Parallel()v := "hello, world"expected := StringHashcode(v)for i := 0; i < 100; i++ {actual := StringHashcode(v)if actual != expected {t.Fatalf("bad: %#v\n\t%#v", actual, expected)}}}

func TestStringHashcode_positiveIndex(t *testing.T) {t.Parallel()// "2338615298" hashes to uint32(2147483648) which is math.MinInt32ips := []string{"192.168.1.3", "192.168.1.5", "2338615298"}for _, ip := range ips {if index := StringHashcode(ip); index < 0 {t.Fatalf("Bad Index %#v for ip %s", index, ip)}}}