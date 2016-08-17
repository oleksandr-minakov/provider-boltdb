package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/oleksandr-minakov/testpg"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: testpg.Provider,
	})
}
