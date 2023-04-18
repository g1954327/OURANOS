PACKAGE_LIST := $(shell go list ./...)
URANOS:
        go build -o URANOS $(PACKAGE_LIST)
test:
        go test $(PACKAGE_LIST)
clean:
        rm -f URANOS
