PACKAGE_LIST := $(shell go list ./...)
VERSION := 0.1.2
NAME := ouranos
DIST := $(NAME)-$(VERSION)

ouranos:
        go build -o ouranos $(PACKAGE_LIST)
coverage.out:
        go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
test:
        go test $(PACKAGE_LIST)
clean:
        rm -f ouranos
