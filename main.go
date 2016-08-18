package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/oleksandr-minakov/boltdb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: boltdb.Provider,
	})
}
