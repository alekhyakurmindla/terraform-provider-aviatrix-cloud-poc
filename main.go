// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/terraform-provider-aviatrix-cloud-poc/provider"
)

func main() {
	debug := true
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	version := "dev"
	opts := providerserver.ServeOpts{
		// NOTE: This is a Terraform Registry provider address,
		// such as registry.terraform.io/hashicorp/aviatrix.
		Address: "hashicorp.com/prd/aviatrix-cloud-poc",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
