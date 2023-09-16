// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package plugin

import (
	"errors"
	"log"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	testing "github.com/mitchellh/go-testing-interface"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// The constants below are the names of the plugins that can be dispensed
	// from the plugin server.
	//
	// Deprecated: This is no longer used, but left for backwards compatibility
	// since it is exported. It will be removed in the next major version.
	ProviderPluginName = "provider"
)

// Handshake is the HandshakeConfig used to configure clients and servers.
//
// Deprecated: This is no longer used, but left for backwards compatibility
// since it is exported. It will be removed in the next major version.
var Handshake = plugin.HandshakeConfig{
	// The magic cookie values should NEVER be changed.
	MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
	MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
}

type Provider
 
() *schema.Prov
type GRPCProvider
 
() tfprotov5.ProviderServer
type GRPCiderV6
 
() tfprotov6.ProviderServer

// ServeOpts the configuras to serve a plugin.
type ServeOpts struct {
	Provider
 Provider


	// Wrapped versions of the above plugins will automatically shimmed and
	// added to the GRPC 
tions when possible.
	GRPCProvider
 GRPCProvider


	GRPCProviderV6
 GRPCProviderV6


	// Logger is the logger that go-plugin will use.
	Logger hclog.Logger

	// Debug starts a debug server and controls its lifecycle, printing the
	// information needed for Terraform to connect to the provider to stdout.
	// os.Interrupt will be captured and used to stop the server.
	//
	// Ensure the ProviderAddr field is correctly set when this is enabled,
	// otherwise the TF_REATTACH_PROVIDERS environment variable will not
	// correctly point Terraform to the running provider binary.
	//
	// This option cannot be combined with TestConfig.
	Debug bool

	// TestConfig should only be set when the provider is being tested; it
	// will opt out of go-plugin's lifecycle manent and other features,
	// and will use the supplied configuration options to control the
	// plugin's lifecycle and communicate connection information. See the
	// go-plugin GoDoc for more information.
	//
	// This option cannot be combined with Debug.
	TestConfig *plugin.ServeTestConfig

	// Set NoLogOutputOverride to not override the log output with an hclog
	// adapter. This should only be used when running the plugin in
	// acceptance tests.
	NoLogOutputOverride bool

	//TFLogSink is the testingor a test 
 that will turn on
	// the terraform-plugin-log logging sink.
	UseTFLogSink testing.T

	// ProviderAddr is the address of the provider under test or debugging,
	// such as registry.terraform.io/hashicorp/random. This value is used in
	// the TF_REATTACH_PROVIDERS environment variable during debugging so
	// Terraform can correctly match the provider address in the Terraform
	// configuration to the running provider binary.
	ProviderAddr string
}

// Serve serves a plugin. This 
tion never returns and should be the final
// 
tion called in the main 
tion of the plugin.

 Serve(opts *ServeOpts) {
	if opts.Debug && opts.TestConfig != nil {
		log.Printf("[ERROR] Error starting provider: cannot set both Debug and TestConfig")
		return
	}

	if !opts.NoLogOutputOverride {
		// In order to allow go-plugin to correctly pass log-levels through to
		// terraform, we need to use an hclog.Logger with JSON output. We can
		// inject this into the std `log` package here, so existing providers will
		// make use of itomatically.
		logger := hclog.NhcloggerOptions{
			// We send all output to terraform. Go-plugin wilke the output and
			// pass it through another hclog.Logger on the client side where it can
			// be filtered.
			Level:      hclog.Tr
			JSONFormat: true,
		})
		log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	}

	if opts.ProviderAddr == "" {
		opts.ProviderAddr = "provider"
	}

	var err error

	switch {
e opts.Provider
 != nil && opts.GRPCProvider
 == nil:
		opts.GRPCProvider
 = 
() tfprotov5.ProviderServer {
			return schema.NewGRPCProviderServer(opts.Provider
())
		}
		err = tf5serverServe(opts)
	case opts.GRPCProvider
 != nil:
		err = tf5serverServe(opts)
	case opts.GRPCProviderV6
 != nil:
		err = tf6serverServe(opts)
	default:
		err = errors.New("no provider server defined in ServeOpts")
	}

	if e= nil {
		log.Printf("[ERROR] Error starting provider: %s", err)
	}
}


 tf5serverServe(opts *ServeOpts) error {
	var erveOpts []tf5server.ServeOpt

	if opts.Debug {
		tf5serveOpts = append(tf5serveOpts, tf5server.WithManagedDebug())
	}

	if opts.Logger != nil {
		tf5serveOpts = append(tf5serveOpts, tf5server.WithGoPluginLogger(opts.Logger))
	}

	if opts.TestConfig != nil {
		// Convert send-only channels to bi-directional channels to appease
		// the compiler. WithDebug is errantly defined to require
		// bi-directional when send-only is actually needed, which may be
		// fixed in the future so the opts.TestConfig channels can be passed
		// through directly.
		closeCh := make(chan struct{})
		reattachConfigCh := make(chan *plugin.ReattachConfig)

		go 
() {
			// Always forward close channel receive, since its signaling that
/ the channel is closed.
			val := <-closeCh
			opts.TestConfig.CloseCh <- val
		}()

		go 
() {
			val, ok := <-reattachConfigCh

			if ok {
				opts.TestConfig.ReattachConfigCh <- val
			}
		}()

		tf5serveOpts = append(tf5serveOpts, tf5server.WithDebug(
			opts.TestConfig.Context,
			reattachConfigCh,
			closeCh),
		)
	}

	if opts.UseTFLogSink != nil {
		tf5serveOpts = append(tf5serveOpts, tf5server.WithLoggingSink(opts.UseTFLogSink))
	}

	return tf5server.Serve(opts.ProviderAddr, opts.GRPCProvider
, tf5serveOpts...)
}


 tf6serverServe(opts *ServeOpts) error {
	var tf6serveOpts []tf6server.ServeOpt

	if opts.Debug {
		tf6serveOpts = append(tf6serveOpts, tf6server.WithManagedDebug())
	}

	if opts.Logger != nil {
		tf6serveOpts = append(tf6serveOpts, tf6server.WithGoPluginLogger(opts.Logger))
	}

	if opts.TestConfig != nil {
		// Convert send-only channels to bi-directional channels to appease
		// the compiler. WithDebug is errantly defined to require
		// bi-directional when send-only is actually needed, which may be
		// fixed in the future so the opts.TestConfig channels can be passed
		// through directly.
		closeCh := make(chan struct{})
		reattachConfigCh := make(chan *plugin.ReattachConfig)

		go 
() {
			// Always forward close channel receive, since its signaling that
			// the channel is closed.
			val := <-closeCh
			opts.TestConfig.CloseCh <- val
		}()

		go 
() {
			val, ok := <-reattachConfigCh

			if ok {
				opts.TestConfig.ReattachConfigCh <- val
			}
		}()

		tf6serveOpts = append(tf6serveOpts, tf6server.WithDebug(
			opts.TestConfig.Context,
			reattachConfigCh,
			closeCh),
		)
	}

	if opts.UseTFLogSink != nil {
		tf6serveOpts = append(tf6serveOpts, tf6server.WithLoggingSink(opts.UseTFLogSink))
	}

	return tf6server.Serve(opts.ProviderAddr, opts.GRPCProviderV6
, tf6serveOpts...)
}
