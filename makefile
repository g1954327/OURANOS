PACKAGE_LIST := $(shell go list ./...)
URANOS:
        go build -o ouranos $(PACKAGE_LIST)
test:
        go test $(PACKAGE_LIST)
clean:
        rm -f ouranos
