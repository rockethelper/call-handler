.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/workflow lambda/workflow/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --aws-profile rockethelper

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
