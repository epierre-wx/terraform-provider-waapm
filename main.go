package main

import (
	"github.com/wallix/terraform-provider-waapm/waapm"
        "github.com/hashicorp/terraform-plugin-sdk"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: waapm.Provider,
        })
}
