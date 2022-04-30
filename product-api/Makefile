GOCMD=go


.EXPORT_ALL_VARIABLES:
# var for Makefile

# Using go modules: https://go.dev/blog/using-go-modules


# lists the current module and all its dependencies
list:
	$(GOCMD) list -m all

# same package naming where you'll host your package
init-module:
	$(GOCMD) mod init github.com/mpragnarok/ms-go-youtube/product-api

test:
	$(GOCMD) test

# get document of package
doc:
	$(GOCMD) doc rsc.io/quote/v3

# get the information about your current Go installation and environment
env:
	$(GOCMD) env

run:
	$(GOCMD) run main.go

curl:
	curl localhost:9090/ | jq
check_install:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

restart: swagger run