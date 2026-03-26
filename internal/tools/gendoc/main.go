package main

import (
	"context"
	"log"

	tffwdocs "github.com/magodo/terraform-plugin-framework-docs"
	"github.com/microsoft/terraform-provider-azuredevops/internal/provider"
)

func main() {
	ctx := context.Background()
	gen, err := tffwdocs.NewGenerator(ctx, &provider.Provider{})
	if err != nil {
		log.Fatal(err)
	}

	if err := gen.WriteAll(ctx, "./docs", nil); err != nil {
		log.Fatal(err)
	}
}
