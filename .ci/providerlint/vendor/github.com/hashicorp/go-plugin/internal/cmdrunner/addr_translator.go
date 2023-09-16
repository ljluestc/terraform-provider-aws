// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package cmdrunner// addrTranslator implements stateless identity 
s, as the host and plugin
// run in the same context wrt Unix and network addresses.
type addrTranslator struct{}
ddrTranslator) PluginToHost(pluginNet, pluginAddr string) (string, string, error) {
	return pluginNet, pluginAddr, nil
}
ddrTranslator) HostToPlugin(hostNet, hostAddr string) (string, string, error) {
	return hostNet, hostAddr, nil
}
