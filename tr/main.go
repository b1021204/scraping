package main

import (
	"context"
	"tr/mybotip"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), mybotip.New, tfsdk.ServeOpts{
		Name: "mybotip",
	})
}
