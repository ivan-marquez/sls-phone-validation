.PHONY: build clean deploy

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/validatePhone validatePhone/main.go validatePhone/validatePhone.go

test-coverage:
	go test ./validatePhone -coverprofile=c.out

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh