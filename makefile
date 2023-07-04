PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.16
NAME := ouranos
DIST := $(NAME)-$(VERSION)

$(NAME): coverage.out
	go build -o $(NAME) cmd/$(NAME)/main.go cmd/$(NAME)/generate_completion.go

coverage.out:
	go test -covermode=count \
			-coverprofile=coverage.out $(PACKAGE_LIST)

docker: ouranos
#	docker build -t ghcr.io/g1954327/ouranos:$(VERSION) -t ghcr.io/g1954327/ouranos:latest .
	docker buildx build -t ghcr.io/g1954327/ouranos:$(VERSION) \
		-t ghcr.io/g1954327/ouranos:latest --platform=linux/arm64/v8,linux/amd64 --push .

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/main.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
#	cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: ouranos
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,arm64,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,arm64,)

distclean: clean
	rm -rf dist

clean:
	rm -f ouranos coverage.out
	rm -rf completions cmd/ouranos/completions
