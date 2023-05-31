package main

import (
	"context"
	"flag"
	"log"

	"github.com/armorfret/terraform-provider-awscreds/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

//revive:disable:line-length-limit
func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: provider.New()}

	if debugMode {
		err := plugin.Debug(context.Background(), "https://registry.terraform.io/providers/armorfret/awscreds", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
