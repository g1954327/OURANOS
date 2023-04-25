PACKAGE_LIST := $(shell go list ./...)
ouranos:
        go build -o ouranos $(PACKAGE_LIST)
test:
        go test -covermode=count -coverprofile=coverage.out $(PACKAGE_LIST)
clean:
        rm -f ouranos
