package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/pvtmert/terraform-provider-http/http"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: http.Provider})
}
