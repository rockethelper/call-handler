.PHONY: build build-bindata clean deploy gomodgen

build: gomodgen build-bindata
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/workflow lambda/workflow/main.go

build-bindata:
	$(GOPATH)/bin/go-bindata --pkg "data" -o internal/pkg/data/data.go data/...

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --aws-profile rockethelper

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
