package main

import (
	"fmt"
	"testing"

	"github.com/mpragnarok/ms-go-youtube/product-api/sdk/client"
	"github.com/mpragnarok/ms-go-youtube/product-api/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	c := client.Default

	params := products.NewListProductsParams()

	prods, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", prods.GetPayload()[0])
	t.Fail()
}
